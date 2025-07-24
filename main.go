package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

// ProcessConfig holds the config for each process
type ProcessConfig struct {
	Name   string
	Dir    string
	Target string
}

func runProcess(config ProcessConfig) *exec.Cmd {
	cmd := exec.Command("go", "run", config.Target)
	cmd.Dir = config.Dir

	// Optional: log file creation
	// logPath := config.Dir + "/" + config.Name + "-wrapper.log"
	// logFile, err := os.Create(logPath)
	// if err != nil {
	// 	log.Fatalf(" Failed to create %s log: %v", config.Name, err)
	// }
	// cmd.Stdout = logFile
	// cmd.Stderr = logFile

	err := cmd.Start()
	if err != nil {
		log.Fatalf(" Failed to start %s: %v", config.Name, err)
	}

	log.Printf(" Started %s (PID: %d)", config.Name, cmd.Process.Pid)
	return cmd
}

func main() {
	// Define your two processes here
	process1 := ProcessConfig{
		Name:   "process_A",
		Dir:    "./process_A",
		Target: ".",
	}

	process2 := ProcessConfig{
		Name:   "process_B",
		Dir:    "./process_B",
		Target: "process_B.go",
	}

	log.Println(" Launching processes:", process1.Name, "and", process2.Name)

	cmd1 := runProcess(process1)
	cmd2 := runProcess(process2)

	// Let them run for N seconds
	runDuration := 120 * time.Second
	log.Printf(" Sleeping for %v before killing processes...", runDuration)
	time.Sleep(runDuration)

	// Kill both if still running
	if cmd1.ProcessState == nil {
		log.Printf(" Killing %s...", process1.Name)
		_ = cmd1.Process.Kill()
	}
	if cmd2.ProcessState == nil {
		log.Printf(" Killing %s...", process2.Name)
		_ = cmd2.Process.Kill()
	}

	log.Println(" Wrapper finished.")
}
