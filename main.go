package main

import (
	"log"
	"os"
	"os/signal"

	"postservice/cmd"

	"github.com/Smart-Pot/pkg"
	"github.com/Smart-Pot/pkg/adapter/amqp"
	"github.com/Smart-Pot/pkg/db"
)

func main() {
	if err := pkg.Config.ReadConfig(); err != nil {
		log.Fatal(err)
	}
	log.Println("Configurations are set")
	if err := db.Connect(db.PkgConfig("posts")); err != nil {
		log.Fatal(err)
	}
	log.Println("DB Connection established")

	if err := amqp.Set(pkg.Config.AMQPAddress); err != nil {
		log.Fatal(err)
	}
	log.Println("AMQP module is set")

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
