[![Build Status](https://travis-ci.org/Sepuka/gowatcher.svg?branch=master)](https://travis-ci.org/Sepuka/gowatcher)
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
    make tests      run tests
```

How does it work?
=====
It send to receivers (slack or telegram) messages periodically, like this:
```
*raspberrypi* *uptime* says:
 17:41:11 up 18 days, 23:16,  3 users,  load average: 0,31, 0,39, 0,37
```
or this
```
*raspberrypi* *who* says:
New user detected: username       pts/1        2018-08-18 17:39 00:02       19165 (192.168.0.245)
```
Now implemented 4 handlers (watchers): `w`, `who`, `df`, `uptime`. Each one run self command and send the result to receivers.

How to configure it?
=====
1. Copy config.json.dist to folder with program
`cp config.json.dist config.json`
2. Modify the configuration
`vim config.json`
3. Restart the app
`watcher -s stop && watcher -d`

Planned features:
=====
1. In case one receiver is not available (for example, telegram is blocked) then other receivers must report this.
2. :heavy_check_mark: ~~Each watcher has self time of loop in settings~~
3. Watchers supports some args in settings
4. `df` watcher can report about critical free size of partial
5. Exists _load average_ graph which published to receivers periodically
6. Released _email_ receiver with customizable watchers
7. Hot reconfigure app
8. :heavy_check_mark: ~~Min & Max loop constrains~~