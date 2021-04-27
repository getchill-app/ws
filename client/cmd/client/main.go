package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/getchill-app/ws/client"
	"github.com/sirupsen/logrus"
)

func main() {
	urs := flag.String("url", "wss://relay.keys.pub/ws", "connect using url")
	auth := flag.String("auth", "", "auth token")

	flag.Parse()

	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	client.SetLogger(log)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	cl, err := client.New(*urs, *auth)
	if err != nil {
		log.Fatal(err)
	}

	if err := cl.Connect(); err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			events, err := cl.ReadEvents()
			if err != nil {
				log.Errorf("read err: %v", err)
				time.Sleep(time.Second * 2) // TODO: Backoff
			} else {
				for _, event := range events {
					log.Infof("%+v\n", event)
				}
			}
		}
	}()

	tokens := []string{}
	for i := 0; i < 20; i++ {
		tokens = append(tokens, fmt.Sprintf("testtoken%d", i))
	}
	if err := cl.Register(tokens); err != nil {
		log.Fatal(err)
	}

	<-interrupt
	cl.Close()
}
