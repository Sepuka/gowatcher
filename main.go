package main

import (
	"flag"
	"fmt"
	"github.com/sepuka/gowatcher/watchers"
	"github.com/sevlyar/go-daemon"
	"github.com/stevenroose/gonfig"
	"log"
	"os"
	"syscall"
	"time"
)

const (
	configPath = "./config.json"
	daemonName = "gowatcher"
)

var (
	buildstamp    = "buildstamp not present"
	githash       = "githash not present"
	stop          = make(chan struct{})
	done          = make(chan struct{})
	watcherResult = make(chan watchers.WatcherResult)
	signal        = flag.String("s", "", "send signal to the daemon\nstop - to stop daemon")
	daemonize     = flag.Bool("d", false, "Daemonize gowatcher")
	testMode      = flag.Bool("t", false, "Test mode")
	version       = flag.Bool("version", false, "Print version info")
	config        = watchers.Configuration{}
	cntxt         = &daemon.Context{
		PidFileName: "pid",
		PidFilePerm: 0644,
		LogFileName: "log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{daemonName},
	}
)

func init() {
	readConfig()
	go Transmitter(watcherResult)
}

func main() {
	flag.Parse()
	daemon.AddCommand(daemon.StringFlag(signal, "stop"), syscall.SIGTERM, termHandler)

	if *testMode {
		watcherResult <- watchers.Test()
		time.Sleep(time.Second * 3)
		return
	}

	if *version {
		fmt.Println("Build time: ", buildstamp)
		fmt.Println("Git hash: ", githash)

		return
	}

	if isDaemonFlagsPresent() {
		d, err := cntxt.Search()
		if err != nil {
			log.Fatalf("Unable send signal to the daemon: %v", err)
		}
		daemon.SendCommands(d)

		return
	}

	if !daemon.WasReborn() && !*daemonize {
		runWatchers()
		log.Println("Press <Ctrl>+C to exit")
		daemonLoop()
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

	log.Println("watcher daemon started")

	runWatchers()

	go daemonLoop()

	err = daemon.ServeSignals()
	if err != nil {
		log.Println("Error:", err)
	}
	log.Println("daemon terminated.")
}

func readConfig() {
	err := gonfig.Load(&config, gonfig.Conf{
		FileDefaultFilename: configPath,
		FileDecoder:         gonfig.DecoderJSON,
		FlagIgnoreUnknown:   true,
	})
	if err != nil {
		log.Printf("Cannot read config: %v", err)
		os.Exit(1)
	}
}

func isDaemonFlagsPresent() bool {
	return len(daemon.ActiveFlags()) > 0
}

func runWatchers() {
	go watchers.DiskFree(watcherResult)
	go watchers.Uptime(watcherResult)
	go watchers.Who(watcherResult)
	go watchers.W(watcherResult)
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
