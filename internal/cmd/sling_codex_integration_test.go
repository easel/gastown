//go:build integration

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/steveyegge/gastown/internal/beads"
)

func TestSlingCodexAgentEnv(t *testing.T) {
	tmpDir := t.TempDir()
	gtBinary := buildGT(t)

	binDir := filepath.Join(tmpDir, "bin")
	if err := os.MkdirAll(binDir, 0755); err != nil {
		t.Fatalf("mkdir bin: %v", err)
	}

	stateDir := filepath.Join(tmpDir, "tmux-state")
	if err := os.MkdirAll(stateDir, 0755); err != nil {
		t.Fatalf("mkdir tmux state: %v", err)
	}

	bdScript := `#!/bin/sh
cmd=""
bead=""
for arg in "$@"; do
  case "$arg" in
    --no-daemon|--allow-stale)
      ;;
    *)
      if [ -z "$cmd" ]; then
        cmd="$arg"
      elif [ -z "$bead" ]; then
        bead="$arg"
      fi
      ;;
  esac
done

case "$cmd" in
  version)
    echo "bd version 0.50.0"
    exit 0
    ;;
  init)
    prefix="gt"
    prev=""
    for arg in "$@"; do
      case "$arg" in
        --prefix=*)
          prefix="${arg#--prefix=}"
          ;;
      esac
      if [ "$prev" = "--prefix" ]; then
        prefix="$arg"
      fi
      prev="$arg"
    done
    mkdir -p .beads
    echo "prefix: $prefix" > .beads/config.yaml
    exit 0
    ;;
  show)
    if [ -z "$bead" ]; then
      bead="gt-unknown"
    fi
    printf '[{"id":"%s","title":"Test bead","status":"open","assignee":"","description":""}]' "$bead"
    exit 0
    ;;
  list)
    echo '[]'
    exit 0
    ;;
  create)
    echo '{}'
    exit 0
    ;;
  *)
    exit 0
    ;;
esac
`
	writeScript(t, binDir, "bd", bdScript)

	tmuxScript := `#!/bin/sh
state_dir="${TMUX_STUB_STATE_DIR:-/tmp/gt-tmux-stub}"
session_file="$state_dir/session"
workdir_file="$state_dir/workdir"
cmd_file="$state_dir/new_session_cmd"
env_file="$state_dir/session_env"

mkdir -p "$state_dir"

case "$1" in
  -V)
    echo "tmux 3.3"
    exit 0
    ;;
  has-session)
    echo "can't find session" 1>&2
    exit 1
    ;;
  new-session)
    session=""
    workdir=""
    last=""
    prev=""
    for arg in "$@"; do
      if [ "$prev" = "-s" ]; then
        session="$arg"
      fi
      if [ "$prev" = "-c" ]; then
        workdir="$arg"
      fi
      prev="$arg"
      last="$arg"
    done
    if [ -n "$session" ]; then
      echo "$session" > "$session_file"
    fi
    if [ -n "$workdir" ]; then
      echo "$workdir" > "$workdir_file"
    fi
    echo "$last" > "$cmd_file"
    : > "$env_file"
    exit 0
    ;;
  set-environment)
    second_last=""
    last=""
    for arg in "$@"; do
      second_last="$last"
      last="$arg"
    done
    if [ -n "$second_last" ]; then
      printf '%s=%s\n' "$second_last" "$last" >> "$env_file"
    fi
    exit 0
    ;;
  list-panes)
    if echo "$*" | grep -q '#{pane_current_command}'; then
      echo "codex"
      exit 0
    fi
    if echo "$*" | grep -q '#{pane_current_path}'; then
      if [ -f "$workdir_file" ]; then
        cat "$workdir_file"
      else
        pwd
      fi
      exit 0
    fi
    if echo "$*" | grep -q '#{pane_pid}'; then
      echo "12345"
      exit 0
    fi
    if echo "$*" | grep -q '#{pane_id}'; then
      echo "%1"
      exit 0
    fi
    echo "%1"
    exit 0
    ;;
  display-message)
    if echo "$*" | grep -q '#{session_name}'; then
      if [ -f "$session_file" ]; then
        cat "$session_file"
      fi
      exit 0
    fi
    exit 0
    ;;
  *)
    exit 0
    ;;
esac
`
	writeScript(t, binDir, "tmux", tmuxScript)

	pathEnv := fmt.Sprintf("PATH=%s%c%s", binDir, os.PathListSeparator, os.Getenv("PATH"))
	baseEnv := append(cleanGTEnv(), pathEnv, "HOME="+tmpDir, "TMUX_STUB_STATE_DIR="+stateDir)

	townRoot := filepath.Join(tmpDir, "town")
	installCmd := exec.Command(gtBinary, "install", townRoot, "--name", "test-town")
	installCmd.Env = baseEnv
	if output, err := installCmd.CombinedOutput(); err != nil {
		t.Fatalf("gt install failed: %v\nOutput: %s", err, output)
	}

	repoPath := createTestGitRepo(t, "testrepo")
	rigName := "testrig"
	rigCmd := exec.Command(gtBinary, "rig", "add", rigName, repoPath)
	rigCmd.Env = baseEnv
	rigCmd.Dir = townRoot
	if output, err := rigCmd.CombinedOutput(); err != nil {
		t.Fatalf("gt rig add failed: %v\nOutput: %s", err, output)
	}

	beadID := "gt-abc123"
	slingCmd := exec.Command(gtBinary, "sling", beadID, rigName, "--agent", "codex", "--no-convoy")
	slingCmd.Env = baseEnv
	slingCmd.Dir = townRoot
	if output, err := slingCmd.CombinedOutput(); err != nil {
		t.Fatalf("gt sling failed: %v\nOutput: %s", err, output)
	}

	cmdPath := filepath.Join(stateDir, "new_session_cmd")
	cmdBytes, err := os.ReadFile(cmdPath)
	if err != nil {
		t.Fatalf("read new session command: %v", err)
	}
	cmdLine := string(cmdBytes)
	if !strings.Contains(cmdLine, "codex --dangerously-bypass-approvals-and-sandbox") {
		t.Fatalf("expected codex command with bypass flag, got: %s", cmdLine)
	}
	if !strings.Contains(cmdLine, "GT_ROLE=polecat") {
		t.Errorf("expected GT_ROLE in command, got: %s", cmdLine)
	}
	if !strings.Contains(cmdLine, "GT_RIG="+rigName) {
		t.Errorf("expected GT_RIG in command, got: %s", cmdLine)
	}

	envPath := filepath.Join(stateDir, "session_env")
	envMap, err := parseEnvFile(envPath)
	if err != nil {
		t.Fatalf("read tmux env: %v", err)
	}

	if envMap["GT_ROLE"] != "polecat" {
		t.Errorf("GT_ROLE = %q, want %q", envMap["GT_ROLE"], "polecat")
	}
	if envMap["GT_RIG"] != rigName {
		t.Errorf("GT_RIG = %q, want %q", envMap["GT_RIG"], rigName)
	}

	polecatName := envMap["GT_POLECAT"]
	if polecatName == "" {
		t.Fatal("GT_POLECAT is empty")
	}

	expectedBDActor := fmt.Sprintf("%s/polecats/%s", rigName, polecatName)
	if envMap["BD_ACTOR"] != expectedBDActor {
		t.Errorf("BD_ACTOR = %q, want %q", envMap["BD_ACTOR"], expectedBDActor)
	}
	if envMap["GIT_AUTHOR_NAME"] != polecatName {
		t.Errorf("GIT_AUTHOR_NAME = %q, want %q", envMap["GIT_AUTHOR_NAME"], polecatName)
	}

	if envMap["BEADS_AGENT_NAME"] != fmt.Sprintf("%s/%s", rigName, polecatName) {
		t.Errorf("BEADS_AGENT_NAME = %q, want %q", envMap["BEADS_AGENT_NAME"], fmt.Sprintf("%s/%s", rigName, polecatName))
	}
	if envMap["BEADS_NO_DAEMON"] != "1" {
		t.Errorf("BEADS_NO_DAEMON = %q, want %q", envMap["BEADS_NO_DAEMON"], "1")
	}

	if envMap["GT_ROOT"] != townRoot {
		t.Errorf("GT_ROOT = %q, want %q", envMap["GT_ROOT"], townRoot)
	}

	rigPath := filepath.Join(townRoot, rigName)
	expectedBeadsDir := beads.ResolveBeadsDir(rigPath)
	if envMap["BEADS_DIR"] != expectedBeadsDir {
		t.Errorf("BEADS_DIR = %q, want %q", envMap["BEADS_DIR"], expectedBeadsDir)
	}
}

func parseEnvFile(path string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	env := make(map[string]string)
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		env[parts[0]] = parts[1]
	}
	return env, nil
}
