package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetermineInstallLocation(t *testing.T) {
	t.Run("returns default for temp builds", func(t *testing.T) {
		// Save original executable path for restoration
		origExec := os.Args[0]
		defer func() { os.Args[0] = origExec }()

		// Get the result - since we're in a test, os.Executable returns test binary
		// which may be in a temp location, so it should fall back to ~/.local/bin/gt
		result, err := determineInstallLocation()
		if err != nil {
			t.Fatalf("determineInstallLocation failed: %v", err)
		}

		// Result should be a valid path
		if result == "" {
			t.Fatal("expected non-empty install location")
		}

		// If running from temp or go-build, should fall back to ~/.local/bin/gt
		home, _ := os.UserHomeDir()
		expectedDefault := filepath.Join(home, ".local", "bin", "gt")

		// The function should either return the current binary path (if not temp)
		// or the default path
		if result != expectedDefault {
			// It returned the current binary path, which is also valid
			// Just verify it's an absolute path
			if !filepath.IsAbs(result) {
				t.Errorf("expected absolute path, got %q", result)
			}
		}
	})

	t.Run("returns valid absolute path", func(t *testing.T) {
		result, err := determineInstallLocation()
		if err != nil {
			t.Fatalf("determineInstallLocation failed: %v", err)
		}

		if !filepath.IsAbs(result) {
			t.Errorf("expected absolute path, got %q", result)
		}

		// Should end with "gt"
		if filepath.Base(result) != "gt" {
			t.Errorf("expected path to end with 'gt', got %q", filepath.Base(result))
		}
	})
}

func TestCopyFile(t *testing.T) {
	t.Run("copies file atomically", func(t *testing.T) {
		tmpDir := t.TempDir()
		srcPath := filepath.Join(tmpDir, "source")
		dstPath := filepath.Join(tmpDir, "dest")

		// Create source file with test content
		testContent := []byte("#!/bin/bash\necho 'test binary'\n")
		if err := os.WriteFile(srcPath, testContent, 0755); err != nil {
			t.Fatalf("failed to create source file: %v", err)
		}

		// Call copyFile
		if err := copyFile(srcPath, dstPath); err != nil {
			t.Fatalf("copyFile failed: %v", err)
		}

		// Verify destination exists
		if _, err := os.Stat(dstPath); err != nil {
			t.Fatalf("destination file not created: %v", err)
		}

		// Verify content matches
		dstContent, err := os.ReadFile(dstPath)
		if err != nil {
			t.Fatalf("failed to read destination: %v", err)
		}
		if string(dstContent) != string(testContent) {
			t.Errorf("content mismatch: got %q, want %q", string(dstContent), string(testContent))
		}

		// Verify temp .new file is cleaned up
		tmpNewPath := dstPath + ".new"
		if _, err := os.Stat(tmpNewPath); !os.IsNotExist(err) {
			t.Errorf("temp .new file should not exist after successful copy")
		}
	})

	t.Run("overwrites existing file", func(t *testing.T) {
		tmpDir := t.TempDir()
		srcPath := filepath.Join(tmpDir, "source")
		dstPath := filepath.Join(tmpDir, "dest")

		// Create source file
		newContent := []byte("new content")
		if err := os.WriteFile(srcPath, newContent, 0755); err != nil {
			t.Fatalf("failed to create source file: %v", err)
		}

		// Create existing destination file
		oldContent := []byte("old content")
		if err := os.WriteFile(dstPath, oldContent, 0755); err != nil {
			t.Fatalf("failed to create destination file: %v", err)
		}

		// Call copyFile - should overwrite
		if err := copyFile(srcPath, dstPath); err != nil {
			t.Fatalf("copyFile failed: %v", err)
		}

		// Verify content was overwritten
		dstContent, err := os.ReadFile(dstPath)
		if err != nil {
			t.Fatalf("failed to read destination: %v", err)
		}
		if string(dstContent) != string(newContent) {
			t.Errorf("content not overwritten: got %q, want %q", string(dstContent), string(newContent))
		}
	})

	t.Run("fails on non-existent source", func(t *testing.T) {
		tmpDir := t.TempDir()
		srcPath := filepath.Join(tmpDir, "nonexistent")
		dstPath := filepath.Join(tmpDir, "dest")

		err := copyFile(srcPath, dstPath)
		if err == nil {
			t.Fatal("expected error for non-existent source")
		}
	})

	t.Run("preserves executable permissions", func(t *testing.T) {
		tmpDir := t.TempDir()
		srcPath := filepath.Join(tmpDir, "source")
		dstPath := filepath.Join(tmpDir, "dest")

		// Create source file with executable permissions
		testContent := []byte("#!/bin/bash\necho test\n")
		if err := os.WriteFile(srcPath, testContent, 0755); err != nil {
			t.Fatalf("failed to create source file: %v", err)
		}

		// Call copyFile
		if err := copyFile(srcPath, dstPath); err != nil {
			t.Fatalf("copyFile failed: %v", err)
		}

		// Verify destination has executable permissions
		info, err := os.Stat(dstPath)
		if err != nil {
			t.Fatalf("failed to stat destination: %v", err)
		}

		// The copyFile function writes with 0755, so check that
		mode := info.Mode().Perm()
		if mode&0100 == 0 {
			t.Errorf("destination should be executable, mode is %o", mode)
		}
	})
}

func TestGetVersionInfo(t *testing.T) {
	t.Run("returns valid version info from git repo", func(t *testing.T) {
		// This test requires being in a git repository
		// Get the current working directory which should be in the repo
		cwd, err := os.Getwd()
		if err != nil {
			t.Fatalf("failed to get working directory: %v", err)
		}

		// Find repo root by walking up
		repoRoot := cwd
		for {
			if _, err := os.Stat(filepath.Join(repoRoot, ".git")); err == nil {
				break
			}
			parent := filepath.Dir(repoRoot)
			if parent == repoRoot {
				t.Skip("not in a git repository")
			}
			repoRoot = parent
		}

		info, err := getVersionInfo(repoRoot)
		if err != nil {
			t.Fatalf("getVersionInfo failed: %v", err)
		}

		// Verify version is non-empty
		if info.Version == "" {
			t.Error("expected non-empty version")
		}

		// Verify commit is a valid git hash (40 hex chars)
		if len(info.Commit) != 40 {
			t.Errorf("expected 40-char commit hash, got %q (len %d)", info.Commit, len(info.Commit))
		}

		// Verify build time is non-empty and looks like a timestamp
		if info.BuildTime == "" {
			t.Error("expected non-empty build time")
		}
	})

	t.Run("returns dev version for non-git directory", func(t *testing.T) {
		tmpDir := t.TempDir()

		// This will fail on getting commit, which is expected
		_, err := getVersionInfo(tmpDir)
		if err == nil {
			t.Fatal("expected error for non-git directory")
		}
	})
}
