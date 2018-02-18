package main

import (
	"gowatcher/watchers"
	"os"
	"log"
	"github.com/sevlyar/go-daemon"
	"flag"
	"syscall"
	"time"
	"net/http"
	"encoding/json"
	"bytes"
	"github.com/tkanos/gonfig"
	"fmt"
)

const configPath = "./config.json"

var (
	signal = flag.String("s", "", `send signal to the daemon
		quit — graceful shutdown
		stop — fast shutdown`)
	stop = make(chan struct{})
	done = make(chan struct{})
)

var config = watchers.Configuration{}

func main() {
	readConfig()
	flag.Parse()
	daemon.AddCommand(daemon.StringFlag(signal, "quit"), syscall.SIGQUIT, termHandler)
	daemon.AddCommand(daemon.StringFlag(signal, "stop"), syscall.SIGTERM, termHandler)

	cntxt := &daemon.Context{
		PidFileName: "pid",
		PidFilePerm: 0644,
		LogFileName: "log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[go-daemon sample]"},
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

	var channel = make(chan watchers.WatcherResult)
	go watchers.DiskFree(channel)
	go watchers.Uptime(channel)
	go watchers.Who(channel)
	go printer(channel)

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

func printer(channel chan watchers.WatcherResult)  {
	for {
		select {
			case result := <-channel:
				if result.IsFailure() {
					log.Printf("disk free error: %v", result.GetError())
					os.Exit(1)
				}
				d := map[string]string{"chat_id": config.ChatId, "text": result.GetText()}
				out := new(bytes.Buffer)
				json.NewEncoder(out).Encode(d)
				url := fmt.Sprintf("https://api.telegram.org/%v:%v/sendMessage", config.BotId, config.Token)
				http.Post(url, "application/json", out)
			case <-time.After(time.Second * config.MainLoopInterval):
				go watchers.DiskFree(channel)
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