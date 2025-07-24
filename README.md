# Go Script to Run Two Scripts Simultaneously

Just modify the process1 and process2 variables to run any Go files or apps:

```
process1 := ProcessConfig{Name: "app1", Dir: "./app1", Target: "main.go"}
process2 := ProcessConfig{Name: "app2", Dir: "./app2", Target: "."}
```

This script will:

- Launch both processes.

- Run them for 120 seconds.

- End both if still running.
