package main

import (
	"flag"
	"fmt"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/definition/logger"
	"github.com/sepuka/gowatcher/definition/transport"
	"github.com/sepuka/gowatcher/services"
	"github.com/sepuka/gowatcher/watchers"
	"github.com/sevlyar/go-daemon"
	"go.uber.org/zap"
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
	watcherResult chan command.Result
	log           *zap.Logger
	signal        = flag.String("s", "", "send signal to the daemon\nstop - to stop daemon")
	daemonize     = flag.Bool("d", false, "Daemonize gowatcher")
	testMode      = flag.Bool("t", false, "Test mode")
	version       = flag.Bool("version", false, "Print version info")
	cntxt         = &daemon.Context{
		PidFileName: "pid",
		PidFilePerm: 0644,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{daemonName},
	}
)

func init() {
	services.Build(config.AppConfig)
	services.Container.Fill(transport.DefTransportChan, &watcherResult)
	services.Container.Fill(logger.DefLogger, &log)
}

func main() {
	go transmitter()

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
			log.Fatal("Unable send signal to the daemon", zap.Any("err", err))
		}
		daemon.SendCommands(d)

		return
	}

	if !daemon.WasReborn() && !*daemonize {
		watchers.RunWatchers()
		log.Info("Press <Ctrl>+C to exit")
		daemonLoop()
		return
	}

	child, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to reborn", zap.Any("err", err))
	}
	if child != nil {
		return
	} else {
		defer cntxt.Release()
	}

	log.Info("watcher daemon started")

	watchers.RunWatchers()

	go daemonLoop()

	err = daemon.ServeSignals()
	if err != nil {
		log.Debug("Error of signal serving", zap.Any("err", err))
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
