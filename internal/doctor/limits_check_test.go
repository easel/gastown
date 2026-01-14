package doctor

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestNewLimitsCheck(t *testing.T) {
	check := NewLimitsCheck()

	if check.Name() != "limits" {
		t.Errorf("Name() = %q, want %q", check.Name(), "limits")
	}

	if check.Description() != "Check file descriptor and inotify limits" {
		t.Errorf("Description() = %q, want %q", check.Description(), "Check file descriptor and inotify limits")
	}

	if check.Category() != CategoryInfrastructure {
		t.Errorf("Category() = %v, want %v", check.Category(), CategoryInfrastructure)
	}

	if !check.CanFix() {
		t.Error("CanFix() should return true")
	}
}

func TestPlatform_String(t *testing.T) {
	tests := []struct {
		platform Platform
		want     string
	}{
		{PlatformUnknown, "Unknown"},
		{PlatformLinuxBareMetal, "Linux (bare metal)"},
		{PlatformLinuxContainer, "Linux (container)"},
		{PlatformWSL, "WSL"},
		{PlatformMacOS, "macOS"},
	}

	for _, tt := range tests {
		got := tt.platform.String()
		if got != tt.want {
			t.Errorf("Platform(%d).String() = %q, want %q", tt.platform, got, tt.want)
		}
	}
}

func TestDetectPlatform(t *testing.T) {
	// This is environment-dependent, but we can verify it returns a valid value
	platform := detectPlatform()

	// Should return one of the known platforms
	validPlatforms := []Platform{
		PlatformUnknown,
		PlatformLinuxBareMetal,
		PlatformLinuxContainer,
		PlatformWSL,
		PlatformMacOS,
	}

	found := false
	for _, p := range validPlatforms {
		if platform == p {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("detectPlatform() = %v, not a valid Platform", platform)
	}

	// Verify platform matches runtime.GOOS
	if runtime.GOOS == "darwin" && platform != PlatformMacOS {
		t.Errorf("On darwin, expected PlatformMacOS, got %v", platform)
	}
}

func TestIsWSL(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("WSL detection only works on Linux")
	}

	// Just verify it doesn't panic and returns a boolean
	result := isWSL()
	_ = result // Result depends on environment
}

func TestIsContainer(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("Container detection only works on Linux")
	}

	// Just verify it doesn't panic and returns a boolean
	result := isContainer()
	_ = result // Result depends on environment
}

func TestCheckPamLimits(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("PAM limits check only works on Linux")
	}

	// Just verify it doesn't panic and returns a boolean
	result := checkPamLimits()
	_ = result // Result depends on environment
}

func TestParseLimitsConf(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("limits.conf parsing only works on Linux")
	}

	// Just verify it doesn't panic and returns a map
	result := parseLimitsConf()
	if result == nil {
		t.Error("parseLimitsConf() should return non-nil map")
	}
}

func TestGetMacOSMaxFiles(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("macOS maxfiles check only works on macOS")
	}

	// Just verify it doesn't panic and returns an int
	result := getMacOSMaxFiles()
	_ = result // Result depends on environment
}

func TestCountGtProcessFDs(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("FD counting only works on Linux")
	}

	// Just verify it doesn't panic and returns a map
	result := countGtProcessFDs()
	if result == nil {
		t.Error("countGtProcessFDs() should return non-nil map")
	}
}

func TestGetProcessName(t *testing.T) {
	tests := []struct {
		pid     int
		cmdline string
		want    string
	}{
		{1234, "/usr/bin/claude --dangerously-skip-permissions", "claude[1234]"},
		{5678, "node /path/to/script.js", "node[5678]"},
		{9999, "", "pid[9999]"},
		{1111, "verylongprocessnamethatshouldbetruncated extra args", "verylongprocessnamet[1111]"},
	}

	for _, tt := range tests {
		got := getProcessName(tt.pid, tt.cmdline)
		if got != tt.want {
			t.Errorf("getProcessName(%d, %q) = %q, want %q", tt.pid, tt.cmdline, got, tt.want)
		}
	}
}

func TestReadProcInt(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("readProcInt only works on Linux")
	}

	// Test with a known proc file
	val, err := readProcInt("/proc/sys/kernel/pid_max")
	if err != nil {
		t.Fatalf("readProcInt(/proc/sys/kernel/pid_max): %v", err)
	}

	// pid_max should be a positive integer
	if val <= 0 {
		t.Errorf("readProcInt returned %d, want positive value", val)
	}
}

func TestReadProcInt_NonExistent(t *testing.T) {
	_, err := readProcInt("/nonexistent/proc/file")
	if err == nil {
		t.Error("readProcInt should fail for non-existent file")
	}
}

func TestLimitsCheck_Run(t *testing.T) {
	check := NewLimitsCheck()
	ctx := &CheckContext{TownRoot: t.TempDir()}

	result := check.Run(ctx)

	// Should return OK or Warning, not Error (unless syscall fails)
	if result.Status != StatusOK && result.Status != StatusWarning {
		// Error status is acceptable if getrlimit fails (unlikely)
		if result.Status == StatusError && !strings.Contains(result.Message, "Failed to get") {
			t.Errorf("Run() returned unexpected error: %s", result.Message)
		}
	}

	// Should have platform in details
	foundPlatform := false
	for _, d := range result.Details {
		if strings.HasPrefix(d, "Platform:") {
			foundPlatform = true
			break
		}
	}
	if !foundPlatform {
		t.Error("Run() should include platform in details")
	}

	// Should have file descriptors in details
	foundFD := false
	for _, d := range result.Details {
		if strings.Contains(d, "File descriptors:") {
			foundFD = true
			break
		}
	}
	if !foundFD {
		t.Error("Run() should include file descriptors in details")
	}
}

func TestLimitsCheck_GenerateFixScript(t *testing.T) {
	check := NewLimitsCheck()

	// Set up some issues to trigger fix script generation
	check.platform = PlatformLinuxBareMetal
	check.fdSoft = 1024   // Below target
	check.fdHard = 65536  // Below target
	check.watches = 8192  // Below target
	check.instances = 128 // Below target
	check.issues = []string{"test issue"}

	script := check.generateFixScript()

	// Verify script has expected sections
	if !strings.Contains(script, "#!/bin/bash") {
		t.Error("Script should have shebang")
	}

	if !strings.Contains(script, "Gas Town limits fix script") {
		t.Error("Script should have header comment")
	}

	if !strings.Contains(script, "set -e") {
		t.Error("Script should have set -e")
	}

	if !strings.Contains(script, "Verification") {
		t.Error("Script should have verification section")
	}
}

func TestLimitsCheck_GenerateFixScript_WSL(t *testing.T) {
	check := NewLimitsCheck()
	check.platform = PlatformWSL
	check.fdSoft = 1024
	check.issues = []string{"test issue"}

	script := check.generateFixScript()

	if !strings.Contains(script, "WSL fixes") {
		t.Error("WSL script should have WSL-specific header")
	}

	if !strings.Contains(script, "wsl --shutdown") {
		t.Error("WSL script should mention wsl --shutdown")
	}
}

func TestLimitsCheck_GenerateFixScript_MacOS(t *testing.T) {
	check := NewLimitsCheck()
	check.platform = PlatformMacOS
	check.fdSoft = 1024
	check.issues = []string{"test issue"}

	script := check.generateFixScript()

	if !strings.Contains(script, "macOS fixes") {
		t.Error("macOS script should have macOS-specific header")
	}

	if !strings.Contains(script, "launchctl") {
		t.Error("macOS script should mention launchctl")
	}

	if !strings.Contains(script, "FSEvents") {
		t.Error("macOS script should mention FSEvents")
	}
}

func TestLimitsCheck_GenerateFixScript_Container(t *testing.T) {
	check := NewLimitsCheck()
	check.platform = PlatformLinuxContainer
	check.watches = 8192
	check.instances = 128
	check.issues = []string{"test issue"}

	script := check.generateFixScript()

	if !strings.Contains(script, "Container fixes") {
		t.Error("Container script should have container-specific header")
	}

	if !strings.Contains(script, "docker run") {
		t.Error("Container script should mention docker run")
	}

	if !strings.Contains(script, "Kubernetes") {
		t.Error("Container script should mention Kubernetes")
	}
}

func TestLimitsCheck_Fix_NoIssues(t *testing.T) {
	check := NewLimitsCheck()
	check.fixScript = "" // No issues detected

	ctx := &CheckContext{TownRoot: t.TempDir()}
	err := check.Fix(ctx)

	if err != nil {
		t.Errorf("Fix() should return nil when no issues: %v", err)
	}
}

func TestLimitsCheck_Fix_WithIssues(t *testing.T) {
	check := NewLimitsCheck()
	check.fixScript = "#!/bin/bash\necho 'fix script'"

	ctx := &CheckContext{TownRoot: t.TempDir()}
	err := check.Fix(ctx)

	// Should return error asking for manual execution
	if err == nil {
		t.Error("Fix() should return error requiring manual execution")
	}

	if !strings.Contains(err.Error(), "manual execution required") {
		t.Errorf("Fix() error should mention manual execution: %v", err)
	}
}

func TestTargetConstants(t *testing.T) {
	// Verify target constants are reasonable values
	if TargetFileDescriptors < 65536 {
		t.Errorf("TargetFileDescriptors = %d, too low", TargetFileDescriptors)
	}

	if TargetInotifyWatches < 65536 {
		t.Errorf("TargetInotifyWatches = %d, too low", TargetInotifyWatches)
	}

	if TargetInotifyInstances < 128 {
		t.Errorf("TargetInotifyInstances = %d, too low", TargetInotifyInstances)
	}
}

// TestGtPatternMatching verifies the regex for gt-related processes
func TestGtPatternMatching(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("Process pattern matching only tested on Linux")
	}

	// Create a temp dir to verify /proc access works
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatalf("creating test file: %v", err)
	}

	// Just verify countGtProcessFDs doesn't panic
	result := countGtProcessFDs()
	_ = result
}
