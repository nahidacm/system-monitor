Static report of running processes
```bash
go run main.go 1234        # Inspect by PID
go run main.go nginx       # Inspect all processes named "nginx"
go run main.go root        # Inspect all processes owned by user "root"
```