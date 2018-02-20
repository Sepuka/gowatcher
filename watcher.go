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
)

var config = watchers.Configuration{}

func main() {
	readConfig()
	flag.Parse()

	var channel = make(chan watchers.WatcherResult)

	if *testMode {
		go watchers.Test(channel)
		go printer(channel)
		time.Sleep(time.Second*3)

		return
	}

	if !daemon.WasReborn() && !*daemonize {
		work(channel)
		time.Sleep(time.Second*3)

		return
	}

	daemon.AddCommand(daemon.StringFlag(signal, "quit"), syscall.SIGQUIT, termHandler)
	daemon.AddCommand(daemon.StringFlag(signal, "stop"), syscall.SIGTERM, termHandler)

	cntxt := &daemon.Context{
		PidFileName: "pid",
		PidFilePerm: 0644,
		LogFileName: "log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{daemonName},
	}
	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}

	defer cntxt.Release()

	log.Print("watcher daemon started")

	work(channel)

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

func work(channel chan watchers.WatcherResult)  {
	go watchers.DiskFree(channel)
	go watchers.Uptime(channel)
	go watchers.Who(channel)
	go printer(channel)
}

func printer(channel chan watchers.WatcherResult)  {
	for {
		select {
			case result := <-channel:
				if result.IsFailure() {
					log.Printf("Watcher %v failed: %v", result.GetName(), result.GetError())
					os.Exit(1)
				}
				watchers.SendMessage(result, config)
			case <-time.After(time.Second * config.MainLoopInterval):
				go watchers.DiskFree(channel)
				go watchers.Uptime(channel)
				go watchers.Who(channel)
		}
	}
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