RaspberryPi compilation command example `GOARCH=arm GOARM=7 go build -o watcher watcher.go`

Usage:
=====
```bash
  -d	Daemonize gowatcher
  -s string
    	send signal to the daemon
		quit — graceful shutdown
		stop — fast shutdown
  -t
    	Test mode - print test phrase without run all modules
  -version
        Print version info
```

Makefile shortcuts:
=====
```bash
    make run        run program (<Ctrl+C> to exit)
    make run_test   run program in test mode (print "It's work" phrase and shutdown)
    make build      build program (with native arch)
    make build_rpi  build program for raspberry pi arch
```