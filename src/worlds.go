package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// getIdleWorlds lists subdirectories of idleDir
func getIdleWorlds() ([]string, error) {
	entries, err := os.ReadDir(idleDir)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return []string{}, nil
		}
		return nil, err
	}
	var names []string
	for _, e := range entries {
		if e.IsDir() {
			names = append(names, e.Name())
		}
	}
	sort.Strings(names)
	return names, nil
}

// activateWorld stops server, swaps world folders, restarts
func activateWorld(name string) error {
	if name == "" {
		return fmt.Errorf("missing world name")
	}
	srcBase := filepath.Join(idleDir, name)
	if _, err := os.Stat(srcBase); err != nil {
		return fmt.Errorf("idle world not found: %s", name)
	}

	// Stop server
	_, _ = runCmd("sudo", "systemctl", "stop", serviceName)

	// Archive current world
	stamp := time.Now().Format("20060102_150405")
	arch := filepath.Join(idleDir, "archived-"+stamp)
	_ = os.MkdirAll(arch, 0755)
	for _, d := range []string{"world", "world_nether", "world_the_end"} {
		src := filepath.Join(mcDir, d)
		if _, err := os.Stat(src); err == nil {
			_ = os.Rename(src, filepath.Join(arch, d))
		}
	}

	// Copy idle world into place
	for _, d := range []string{"world", "world_nether", "world_the_end"} {
		src := filepath.Join(srcBase, d)
		if _, err := os.Stat(src); err == nil {
			if err := copyDir(src, filepath.Join(mcDir, d)); err != nil {
				return err
			}
		}
	}

	setActiveName(name)

	// Start server
	_, _ = runCmd("sudo", "systemctl", "start", serviceName)
	return nil
}

// deactivateWorld removes current active world
func deactivateWorld() error {
	_, _ = runCmd("sudo", "systemctl", "stop", serviceName)
	if err := removeActiveWorld(); err != nil {
		return err
	}
	clearActiveName()
	return nil
}
