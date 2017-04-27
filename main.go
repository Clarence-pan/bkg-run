package main

import "log"
import "os"
import "os/exec"
import "fmt"

func main() {
	args := os.Args
	if len(args) <= 1 {
		printHelp()
		return
	}

	switch args[1] {
	case "-h", "--help", "-?":
		printHelp()
		return
	default:
		ps := runCommand(args[1], args[2:])
		fmt.Print(ps.Pid)
		return
	}
}

func runCommand(cmd string, args []string) *os.Process {
	var err error
	exeFile, err := exec.LookPath(cmd)
	if err != nil {
		exeFile = cmd
	}

	c := exec.Command(exeFile, args...)
	c.Stdin = nil
	c.Stdout = nil
	c.Stderr = nil

	err = c.Start()
	if err != nil {
		log.Panic(err)
		return nil
	}

	return c.Process
}

func printHelp() {
	fmt.Println("Usage: exe-delegate -o <delegate-exe-file> <original-exe-file-path> [...args]")
	os.Exit(1)
}
