RaspberryPi compilation command example `GOARCH=arm GOARM=7 go build -o watcher watcher.go`

Usage:`
  -d	Daemonize gowatcher
  -s string
    	send signal to the daemon
		quit — graceful shutdown
		stop — fast shutdown
  -t
    	Test mode - print test phrase without run all modules
`