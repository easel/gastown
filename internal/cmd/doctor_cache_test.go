package cmd

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestFormatSize(t *testing.T) {
	tests := []struct {
		bytes int64
		want  string
	}{
		{0, "0B"},
		{512, "512B"},
		{1024, "1.0KB"},
		{1536, "1.5KB"},
		{1024 * 1024, "1.0MB"},
		{1024 * 1024 * 1.5, "1.5MB"},
		{1024 * 1024 * 1024, "1.0GB"},
		{1024 * 1024 * 1024 * 2.5, "2.5GB"},
	}

	for _, tt := range tests {
		got := formatSize(tt.bytes)
		if got != tt.want {
			t.Errorf("formatSize(%d) = %q, want %q", tt.bytes, got, tt.want)
		}
	}
}

func TestGetAgeBucket(t *testing.T) {
	tests := []struct {
		age  time.Duration
		want string
	}{
		{0, "<1 day"},
		{12 * time.Hour, "<1 day"},
		{23 * time.Hour, "<1 day"},
		{25 * time.Hour, "1-7 days"},
		{3 * 24 * time.Hour, "1-7 days"},
		{6 * 24 * time.Hour, "1-7 days"},
		{8 * 24 * time.Hour, "7-30 days"},
		{15 * 24 * time.Hour, "7-30 days"},
		{29 * 24 * time.Hour, "7-30 days"},
		{31 * 24 * time.Hour, ">30 days"},
		{100 * 24 * time.Hour, ">30 days"},
	}

	for _, tt := range tests {
		got := getAgeBucket(tt.age)
		if got != tt.want {
			t.Errorf("getAgeBucket(%v) = %q, want %q", tt.age, got, tt.want)
		}
	}
}

func TestIsCleanableDir(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// Cleanable directories
		{"debug", true},
		{"shell-snapshots", true},
		{"session-env", true},
		{"telemetry", true},
		{"file-history", true},

		// Non-cleanable directories
		{"projects", false},
		{"settings", false},
		{"random-dir", false},
		{"", false},
	}

	for _, tt := range tests {
		got := isCleanableDir(tt.name)
		if got != tt.want {
			t.Errorf("isCleanableDir(%q) = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGetDirSize(t *testing.T) {
	// Create temp directory with known file sizes
	tmpDir := t.TempDir()

	// Create some files with known sizes
	file1 := filepath.Join(tmpDir, "file1.txt")
	if err := os.WriteFile(file1, make([]byte, 1000), 0644); err != nil {
		t.Fatalf("creating file1: %v", err)
	}

	file2 := filepath.Join(tmpDir, "file2.txt")
	if err := os.WriteFile(file2, make([]byte, 2000), 0644); err != nil {
		t.Fatalf("creating file2: %v", err)
	}

	// Create subdirectory with file
	subDir := filepath.Join(tmpDir, "subdir")
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatalf("creating subdir: %v", err)
	}
	file3 := filepath.Join(subDir, "file3.txt")
	if err := os.WriteFile(file3, make([]byte, 500), 0644); err != nil {
		t.Fatalf("creating file3: %v", err)
	}

	got := getDirSize(tmpDir)
	want := int64(3500) // 1000 + 2000 + 500

	if got != want {
		t.Errorf("getDirSize(%q) = %d, want %d", tmpDir, got, want)
	}
}

func TestGetDirSize_EmptyDir(t *testing.T) {
	tmpDir := t.TempDir()
	got := getDirSize(tmpDir)

	if got != 0 {
		t.Errorf("getDirSize(empty) = %d, want 0", got)
	}
}

func TestGetDirSize_NonExistent(t *testing.T) {
	got := getDirSize("/nonexistent/path/that/does/not/exist")

	if got != 0 {
		t.Errorf("getDirSize(nonexistent) = %d, want 0", got)
	}
}

func TestCollectCacheStats(t *testing.T) {
	// Create a mock cache directory structure
	cacheDir := t.TempDir()

	// Create cleanable subdirectory with old and new files
	debugDir := filepath.Join(cacheDir, "debug")
	if err := os.MkdirAll(debugDir, 0755); err != nil {
		t.Fatalf("creating debug dir: %v", err)
	}

	// Create a "new" file (should not be stale)
	newFile := filepath.Join(debugDir, "new.log")
	if err := os.WriteFile(newFile, make([]byte, 100), 0644); err != nil {
		t.Fatalf("creating new file: %v", err)
	}

	// Create an "old" file by setting its modtime to 10 days ago
	oldFile := filepath.Join(debugDir, "old.log")
	if err := os.WriteFile(oldFile, make([]byte, 200), 0644); err != nil {
		t.Fatalf("creating old file: %v", err)
	}
	oldTime := time.Now().AddDate(0, 0, -10)
	if err := os.Chtimes(oldFile, oldTime, oldTime); err != nil {
		t.Fatalf("setting old file time: %v", err)
	}

	// Create non-cleanable subdirectory
	projectsDir := filepath.Join(cacheDir, "projects")
	if err := os.MkdirAll(projectsDir, 0755); err != nil {
		t.Fatalf("creating projects dir: %v", err)
	}
	projectFile := filepath.Join(projectsDir, "data.json")
	if err := os.WriteFile(projectFile, make([]byte, 500), 0644); err != nil {
		t.Fatalf("creating project file: %v", err)
	}

	// Collect stats with 7-day max age
	stats := collectCacheStats(cacheDir, 7)

	// Verify directory counts
	if stats.TotalDirs != 2 {
		t.Errorf("TotalDirs = %d, want 2", stats.TotalDirs)
	}

	// Verify total files
	if stats.TotalFiles != 3 {
		t.Errorf("TotalFiles = %d, want 3", stats.TotalFiles)
	}

	// Verify total size (100 + 200 + 500 = 800)
	if stats.TotalSize != 800 {
		t.Errorf("TotalSize = %d, want 800", stats.TotalSize)
	}

	// Verify subdirectory sizes are tracked
	if _, ok := stats.ByDir["debug"]; !ok {
		t.Error("ByDir should contain 'debug'")
	}
	if _, ok := stats.ByDir["projects"]; !ok {
		t.Error("ByDir should contain 'projects'")
	}

	// Verify age distribution has entries
	if len(stats.ByAge) == 0 {
		t.Error("ByAge should have entries")
	}

	// Verify stale entries detected (only from cleanable dirs)
	hasStaleOldFile := false
	for _, entry := range stats.Stale {
		if filepath.Base(entry.Path) == "old.log" {
			hasStaleOldFile = true
			break
		}
	}
	if !hasStaleOldFile {
		t.Error("Stale should include old.log from debug dir")
	}
}

func TestCollectCacheStats_EmptyDir(t *testing.T) {
	cacheDir := t.TempDir()
	stats := collectCacheStats(cacheDir, 7)

	if stats.TotalSize != 0 {
		t.Errorf("TotalSize = %d, want 0", stats.TotalSize)
	}
	if stats.TotalFiles != 0 {
		t.Errorf("TotalFiles = %d, want 0", stats.TotalFiles)
	}
	if stats.TotalDirs != 0 {
		t.Errorf("TotalDirs = %d, want 0", stats.TotalDirs)
	}
}

func TestCacheEntry(t *testing.T) {
	entry := CacheEntry{
		Path:    "/path/to/file",
		Size:    1234,
		ModTime: time.Now(),
		IsDir:   false,
	}

	if entry.Path != "/path/to/file" {
		t.Errorf("Path = %q, want %q", entry.Path, "/path/to/file")
	}
	if entry.Size != 1234 {
		t.Errorf("Size = %d, want %d", entry.Size, 1234)
	}
	if entry.IsDir {
		t.Error("IsDir should be false")
	}
}

func TestCreateManifest(t *testing.T) {
	// Create temp cache structure
	cacheDir := t.TempDir()
	subDir := filepath.Join(cacheDir, "subdir")
	os.MkdirAll(subDir, 0755)
	os.WriteFile(filepath.Join(cacheDir, "file1.txt"), []byte("test"), 0644)
	os.WriteFile(filepath.Join(subDir, "file2.txt"), []byte("test2"), 0644)

	manifest := createManifest(cacheDir)

	// Verify manifest structure
	if manifest["cache_dir"] != cacheDir {
		t.Errorf("cache_dir = %v, want %v", manifest["cache_dir"], cacheDir)
	}

	if manifest["timestamp"] == nil {
		t.Error("manifest should have timestamp")
	}

	entries, ok := manifest["entries"].([]string)
	if !ok {
		t.Fatal("entries should be []string")
	}

	// Should have at least the root "." and the files/dirs
	if len(entries) < 3 {
		t.Errorf("entries should have at least 3 items, got %d", len(entries))
	}

	entryCount, ok := manifest["entry_count"].(int)
	if !ok {
		t.Fatal("entry_count should be int")
	}
	if entryCount != len(entries) {
		t.Errorf("entry_count = %d, len(entries) = %d", entryCount, len(entries))
	}
}

func TestCleanStaleEntries_DryRun(t *testing.T) {
	cacheDir := t.TempDir()

	// Create a stale file
	debugDir := filepath.Join(cacheDir, "debug")
	os.MkdirAll(debugDir, 0755)
	staleFile := filepath.Join(debugDir, "stale.log")
	os.WriteFile(staleFile, []byte("stale data"), 0644)

	// Set modtime to 10 days ago
	oldTime := time.Now().AddDate(0, 0, -10)
	os.Chtimes(staleFile, oldTime, oldTime)

	// Collect stats
	stats := collectCacheStats(cacheDir, 7)

	// Run cleanup in dry-run mode
	err := cleanStaleEntries(cacheDir, stats, true /* dryRun */)
	if err != nil {
		t.Fatalf("cleanStaleEntries dry-run: %v", err)
	}

	// Verify file still exists after dry-run
	if _, err := os.Stat(staleFile); os.IsNotExist(err) {
		t.Error("stale file should still exist after dry-run")
	}
}

func TestCleanStaleEntries_NoStaleEntries(t *testing.T) {
	cacheDir := t.TempDir()

	// Create only fresh files
	debugDir := filepath.Join(cacheDir, "debug")
	os.MkdirAll(debugDir, 0755)
	freshFile := filepath.Join(debugDir, "fresh.log")
	os.WriteFile(freshFile, []byte("fresh data"), 0644)

	// Collect stats (no stale entries)
	stats := collectCacheStats(cacheDir, 7)

	// Run cleanup
	err := cleanStaleEntries(cacheDir, stats, false)
	if err != nil {
		t.Fatalf("cleanStaleEntries: %v", err)
	}

	// File should still exist
	if _, err := os.Stat(freshFile); os.IsNotExist(err) {
		t.Error("fresh file should still exist")
	}
}
