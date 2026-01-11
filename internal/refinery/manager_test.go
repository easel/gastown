package refinery

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/steveyegge/gastown/internal/rig"
)

func setupTestManager(t *testing.T) (*Manager, string) {
	t.Helper()

	// Create temp directory structure
	tmpDir := t.TempDir()
	rigPath := filepath.Join(tmpDir, "testrig")
	if err := os.MkdirAll(filepath.Join(rigPath, ".runtime"), 0755); err != nil {
		t.Fatalf("mkdir .runtime: %v", err)
	}

	r := &rig.Rig{
		Name: "testrig",
		Path: rigPath,
	}

	return NewManager(r), rigPath
}

func TestManager_GetMR(t *testing.T) {
	mgr, _ := setupTestManager(t)

	// Create a test MR in the pending queue
	mr := &MergeRequest{
		ID:       "gt-mr-abc123",
		Branch:   "polecat/Toast/gt-xyz",
		Worker:   "Toast",
		IssueID:  "gt-xyz",
		Status:   MROpen,
		Error:    "test failure",
	}

	if err := mgr.RegisterMR(mr); err != nil {
		t.Fatalf("RegisterMR: %v", err)
	}

	t.Run("find existing MR", func(t *testing.T) {
		found, err := mgr.GetMR("gt-mr-abc123")
		if err != nil {
			t.Errorf("GetMR() unexpected error: %v", err)
		}
		if found == nil {
			t.Fatal("GetMR() returned nil")
		}
		if found.ID != mr.ID {
			t.Errorf("GetMR() ID = %s, want %s", found.ID, mr.ID)
		}
	})

	t.Run("MR not found", func(t *testing.T) {
		_, err := mgr.GetMR("nonexistent-mr")
		if err != ErrMRNotFound {
			t.Errorf("GetMR() error = %v, want %v", err, ErrMRNotFound)
		}
	})
}

func TestManager_Retry(t *testing.T) {
	t.Run("retry failed MR clears error", func(t *testing.T) {
		mgr, _ := setupTestManager(t)

		// Create a failed MR
		mr := &MergeRequest{
			ID:       "gt-mr-failed",
			Branch:   "polecat/Toast/gt-xyz",
			Worker:   "Toast",
			Status:   MROpen,
			Error:    "merge conflict",
		}

		if err := mgr.RegisterMR(mr); err != nil {
			t.Fatalf("RegisterMR: %v", err)
		}

		// Retry without processing
		err := mgr.Retry("gt-mr-failed", false)
		if err != nil {
			t.Errorf("Retry() unexpected error: %v", err)
		}

		// Verify error was cleared
		found, _ := mgr.GetMR("gt-mr-failed")
		if found.Error != "" {
			t.Errorf("Retry() error not cleared, got %s", found.Error)
		}
	})

	t.Run("retry non-failed MR fails", func(t *testing.T) {
		mgr, _ := setupTestManager(t)

		// Create a successful MR (no error)
		mr := &MergeRequest{
			ID:     "gt-mr-success",
			Branch: "polecat/Toast/gt-abc",
			Worker: "Toast",
			Status: MROpen,
			Error:  "", // No error
		}

		if err := mgr.RegisterMR(mr); err != nil {
			t.Fatalf("RegisterMR: %v", err)
		}

		err := mgr.Retry("gt-mr-success", false)
		if err != ErrMRNotFailed {
			t.Errorf("Retry() error = %v, want %v", err, ErrMRNotFailed)
		}
	})

	t.Run("retry nonexistent MR fails", func(t *testing.T) {
		mgr, _ := setupTestManager(t)

		err := mgr.Retry("nonexistent", false)
		if err != ErrMRNotFound {
			t.Errorf("Retry() error = %v, want %v", err, ErrMRNotFound)
		}
	})
}

func TestManager_RegisterMR(t *testing.T) {
	mgr, rigPath := setupTestManager(t)

	mr := &MergeRequest{
		ID:           "gt-mr-new",
		Branch:       "polecat/Cheedo/gt-123",
		Worker:       "Cheedo",
		IssueID:      "gt-123",
		TargetBranch: "main",
		CreatedAt:    time.Now(),
		Status:       MROpen,
	}

	if err := mgr.RegisterMR(mr); err != nil {
		t.Fatalf("RegisterMR: %v", err)
	}

	// Verify it was saved to disk
	stateFile := filepath.Join(rigPath, ".runtime", "refinery.json")
	data, err := os.ReadFile(stateFile)
	if err != nil {
		t.Fatalf("reading state file: %v", err)
	}

	var ref Refinery
	if err := json.Unmarshal(data, &ref); err != nil {
		t.Fatalf("unmarshal state: %v", err)
	}

	if ref.PendingMRs == nil {
		t.Fatal("PendingMRs is nil")
	}

	saved, ok := ref.PendingMRs["gt-mr-new"]
	if !ok {
		t.Fatal("MR not found in PendingMRs")
	}

	if saved.Worker != "Cheedo" {
		t.Errorf("saved MR worker = %s, want Cheedo", saved.Worker)
	}
}

// TestManager_EnsureRefineryWorktree tests the worktree creation/validation logic.
// This reproduces the bug where refinery started in mayor/rig instead of refinery/rig.
func TestManager_EnsureRefineryWorktree(t *testing.T) {
	t.Run("returns existing refinery/rig path", func(t *testing.T) {
		tmpDir := t.TempDir()
		rigPath := filepath.Join(tmpDir, "testrig")

		// Create the refinery/rig directory (simulating a properly set up rig)
		refineryRigPath := filepath.Join(rigPath, "refinery", "rig")
		if err := os.MkdirAll(refineryRigPath, 0755); err != nil {
			t.Fatalf("mkdir refinery/rig: %v", err)
		}

		r := &rig.Rig{
			Name: "testrig",
			Path: rigPath,
		}
		mgr := NewManager(r)

		workDir, err := mgr.EnsureRefineryWorktree()
		if err != nil {
			t.Fatalf("EnsureRefineryWorktree: %v", err)
		}

		if workDir != refineryRigPath {
			t.Errorf("workDir = %s, want %s", workDir, refineryRigPath)
		}
	})

	t.Run("creates missing refinery/rig worktree from bare repo", func(t *testing.T) {
		tmpDir := t.TempDir()
		rigPath := filepath.Join(tmpDir, "testrig")

		// Create the bare repo structure (as if rig was partially set up)
		bareRepoPath := filepath.Join(rigPath, ".repo.git")
		if err := os.MkdirAll(bareRepoPath, 0755); err != nil {
			t.Fatalf("mkdir .repo.git: %v", err)
		}

		// Initialize bare repo using exec.Command
		cmd := exec.Command("git", "init", "--bare")
		cmd.Dir = bareRepoPath
		if err := cmd.Run(); err != nil {
			t.Fatalf("git init --bare: %v", err)
		}

		// Create an initial commit so we have something to checkout
		// Use a temporary worktree to make the initial commit
		initDir := filepath.Join(tmpDir, "init-worktree")
		if err := os.MkdirAll(initDir, 0755); err != nil {
			t.Fatalf("mkdir init-worktree: %v", err)
		}

		// Initialize fresh repo with explicit main branch and add remote
		cmd = exec.Command("git", "init", "--initial-branch=main")
		cmd.Dir = initDir
		if err := cmd.Run(); err != nil {
			t.Fatalf("git init: %v", err)
		}

		// Set up git config for commits
		cmd = exec.Command("git", "config", "user.email", "test@test.com")
		cmd.Dir = initDir
		_ = cmd.Run()
		cmd = exec.Command("git", "config", "user.name", "Test User")
		cmd.Dir = initDir
		_ = cmd.Run()

		// Create a dummy file and commit
		dummyFile := filepath.Join(initDir, "README.md")
		if err := os.WriteFile(dummyFile, []byte("# Test"), 0644); err != nil {
			t.Fatalf("write README.md: %v", err)
		}
		cmd = exec.Command("git", "add", "README.md")
		cmd.Dir = initDir
		if err := cmd.Run(); err != nil {
			t.Fatalf("git add: %v", err)
		}
		cmd = exec.Command("git", "commit", "-m", "Initial commit")
		cmd.Dir = initDir
		if err := cmd.Run(); err != nil {
			t.Fatalf("git commit: %v", err)
		}

		// Add remote and push
		cmd = exec.Command("git", "remote", "add", "origin", bareRepoPath)
		cmd.Dir = initDir
		if err := cmd.Run(); err != nil {
			t.Fatalf("git remote add: %v", err)
		}

		// Push to bare repo (we initialized with --initial-branch=main)
		cmd = exec.Command("git", "push", "-u", "origin", "main")
		cmd.Dir = initDir
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git push: %v\n%s", err, out)
		}

		// Create .runtime dir (needed for state)
		if err := os.MkdirAll(filepath.Join(rigPath, ".runtime"), 0755); err != nil {
			t.Fatalf("mkdir .runtime: %v", err)
		}

		// Do NOT create refinery/rig - this is the bug scenario
		// Also do NOT create mayor/rig - we want to test auto-creation

		r := &rig.Rig{
			Name: "testrig",
			Path: rigPath,
		}
		mgr := NewManager(r)

		workDir, err := mgr.EnsureRefineryWorktree()
		if err != nil {
			t.Fatalf("EnsureRefineryWorktree: %v", err)
		}

		expectedPath := filepath.Join(rigPath, "refinery", "rig")
		if workDir != expectedPath {
			t.Errorf("workDir = %s, want %s", workDir, expectedPath)
		}

		// Verify the worktree was actually created
		if _, err := os.Stat(workDir); os.IsNotExist(err) {
			t.Errorf("refinery/rig was not created")
		}
	})

	t.Run("errors when bare repo does not exist", func(t *testing.T) {
		tmpDir := t.TempDir()
		rigPath := filepath.Join(tmpDir, "testrig")
		if err := os.MkdirAll(rigPath, 0755); err != nil {
			t.Fatalf("mkdir: %v", err)
		}

		// No .repo.git, no mayor/rig, no refinery/rig
		r := &rig.Rig{
			Name: "testrig",
			Path: rigPath,
		}
		mgr := NewManager(r)

		_, err := mgr.EnsureRefineryWorktree()
		if err == nil {
			t.Error("EnsureRefineryWorktree should error when no repo exists")
		}
	})
}
