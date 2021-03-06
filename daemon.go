package main

import (
	"os"
	"os/exec"
)

func runDaemon(profile string, systemd bool) {
	cmd := exec.Command("./bin/bedrockserver", profile)
	cmd.Dir, _ = os.Getwd()
	cmd.Env = append(cmd.Env, "LD_LIBRARY_PATH=./lib", "XDG_CACHE_HOME=./cache")
	cmd.Start()
	if systemd {
		cmd.Wait()
	}
}
