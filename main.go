package main

import (
	"log"
	"os"
	"os/signal"

	"postservice/cmd"
	"postservice/data"

	"github.com/Smart-Pot/pkg"
	"github.com/Smart-Pot/pkg/adapter/amqp"
)

func main() {
	pkg.Config.ReadConfig()
	err := amqp.Set("amqp://guest:guest@rabbitmq:5672")
	if err != nil {
		log.Fatal(err)
	}
	data.DatabaseConnection()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		if err := cmd.Execute(); err != nil {
			log.Fatal(err)
		}
	}()
	sig := <-c
	log.Println("GOT SIGNAL: " + sig.String())
}
