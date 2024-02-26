package main

import (
	"context"
	"fmt"
	"os/exec"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd := exec.CommandContext(ctx, "sleep", "10")
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Process started with PID: %d\n", cmd.Process.Pid)
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("Process completed.")
}
