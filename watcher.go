package main

import (
	"gowatcher/watchers"
	"os"
	"log"
	"github.com/sevlyar/go-daemon"
	"flag"
	"syscall"
	"time"
	"github.com/tkanos/gonfig"
)

const (
	configPath = "./config.json"
	daemonName = "gowatcher"
)

var (
	signal = flag.String("s", "", `send signal to the daemon
		quit — graceful shutdown
		stop — fast shutdown`)
	stop = make(chan struct{})
	done = make(chan struct{})
	daemonize = flag.Bool("d", false, "Daemonize gowatcher")
	testMode = flag.Bool("t", false, "Test mode")
	config = watchers.Configuration{}
	cntxt = &daemon.Context{
		PidFileName: "pid",
		PidFilePerm: 0644,
		LogFileName: "log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{daemonName},
	}
)

func main() {
	readConfig()
	flag.Parse()
	daemon.AddCommand(daemon.StringFlag(signal, "quit"), syscall.SIGQUIT, termHandler)
	daemon.AddCommand(daemon.StringFlag(signal, "stop"), syscall.SIGTERM, termHandler)

	if *testMode {
		watchers.SendUrgentMessage(watchers.Test(), config)

		return
	}

	if isDaemonFlagsPresent() {
		d, err := cntxt.Search()
		if err != nil {
			log.Fatalf("Unable send signal to the daemon: ", err)
		}
		daemon.SendCommands(d)

		return
	}

	if !daemon.WasReborn() && !*daemonize {
		runWatchers()
		time.Sleep(time.Second*3)

		return
	}

	child, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if child != nil {
		return
	} else {
		defer cntxt.Release()
	}

	log.Print("watcher daemon started")

	runWatchers()

	go daemonLoop()

	err = daemon.ServeSignals()
	if err != nil {
		log.Println("Error:", err)
	}
	log.Println("daemon terminated.")
}

func readConfig()  {
	err := gonfig.GetConf(configPath, &config)
	if err != nil {
		log.Printf("Cannot read config: %v", err)
		os.Exit(1)
	}
}

func isDaemonFlagsPresent() bool {
	return len(daemon.ActiveFlags()) > 0
}

func runWatchers() {
	go watchers.DiskFree(config)
	go watchers.Uptime(config)
	go watchers.Who(config)
	go watchers.W(config)
}

func daemonLoop() {
	for {
		time.Sleep(time.Second)
		if _, ok := <-stop; ok {
			break
		}
	}
	done <- struct{}{}
}

func termHandler(sig os.Signal) error {
	log.Println("terminating...")
	stop <- struct{}{}
	if sig == syscall.SIGQUIT {
		<-done
	}

	return daemon.ErrStop
}