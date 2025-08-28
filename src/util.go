package main

import (
	"bufio"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	serviceName = "papermc"
	mcDir       = "/home/minecraft/minecraft-server"
	idleDir     = "/home/minecraft/minecraft-server/idle_worlds"
	backupDir   = "/home/minecraft/minecraft-server/backups"
	activeMeta  = "/home/minecraft/minecraft-server/.active_name"
	addr        = ":8080"
)

// run command and return output
func runCmd(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(out)), err
}

// read level-name from server.properties
func readLevelName() string {
	f, err := os.Open(filepath.Join(mcDir, "server.properties"))
	if err != nil {
		return "world"
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, "level-name=") {
			return strings.TrimPrefix(line, "level-name=")
		}
	}
	return "world"
}

// active marker
func activeName() string {
	b, err := os.ReadFile(activeMeta)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(b))
}
func setActiveName(n string) { _ = os.WriteFile(activeMeta, []byte(n), 0644) }
func clearActiveName()       { _ = os.Remove(activeMeta) }

// remove current active world folders
func removeActiveWorld() error {
	for _, d := range []string{"world", "world_nether", "world_the_end"} {
		p := filepath.Join(mcDir, d)
		if _, err := os.Stat(p); err == nil {
			if err := os.RemoveAll(p); err != nil {
				return err
			}
		}
	}
	return nil
}

// copy directory using cp -a
func copyDir(src, dst string) error {
	_, err := runCmd("/bin/cp", "-a", src+"/.", dst)
	return err
}
