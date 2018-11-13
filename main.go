package main

import (
	"flag"
	"fmt"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/services"
	"github.com/sepuka/gowatcher/watchers"
	_ "github.com/sepuka/gowatcher/services/logger"
	"github.com/sevlyar/go-daemon"
	"github.com/sirupsen/logrus"
	"os"
	"syscall"
	"time"
)

const (
	daemonName = "gowatcher"
)

var (
	buildstamp    = "buildstamp not present"
	githash       = "githash not present"
	stop          = make(chan struct{})
	done          = make(chan struct{})
	watcherResult = make(chan command.Result)
	signal        = flag.String("s", "", "send signal to the daemon\nstop - to stop daemon")
	daemonize     = flag.Bool("d", false, "Daemonize gowatcher")
	testMode      = flag.Bool("t", false, "Test mode")
	version       = flag.Bool("version", false, "Print version info")
	cntxt         = &daemon.Context{
		PidFileName: "pid",
		PidFilePerm: 0644,
		LogFileName: "log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{daemonName},
	}
	log *logrus.Logger
)


func init() {
	config.InitConfig()
	services.Build(config.AppConfig)
	log = services.Container.Get(services.Logger).(*logrus.Logger)
	go transmitter(watcherResult)
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
		watchers.RunWatchers(watcherResult)
		watchers.RunStatCollectors(watcherResult)
		log.Info("Press <Ctrl>+C to exit")
		daemonLoop()
		return
	}

	child, err := cntxt.Reborn()
	if err != nil {
		log.Fatalf("Unable to run: ", err)
	}
	if child != nil {
		return
	} else {
		defer cntxt.Release()
	}

	log.Info("watcher daemon started")

	watchers.RunWatchers(watcherResult)
	watchers.RunStatCollectors(watcherResult)

	go daemonLoop()

	err = daemon.ServeSignals()
	if err != nil {
		log.Debugf("Error: %s", err)
	}
	log.Info("daemon terminated.")
}

func isDaemonFlagsPresent() bool {
	return len(daemon.ActiveFlags()) > 0
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
	log.Info("terminating...")
	stop <- struct{}{}
	if sig == syscall.SIGQUIT {
		<-done
	}

	return daemon.ErrStop
}
