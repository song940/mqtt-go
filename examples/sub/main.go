package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/song940/mqtt-go/mqtt"
	"github.com/song940/mqtt-go/proto"
)

var host = flag.String("host", "localhost:1883", "hostname of broker")
var id = flag.String("id", "", "client id")
var user = flag.String("user", "", "username")
var pass = flag.String("pass", "", "password")
var dump = flag.Bool("dump", false, "dump messages?")

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "usage: sub topic [topic topic...]")
		return
	}

	cc, err := mqtt.NewClient(*host)
	if err != nil {
		log.Fatal(err)
	}
	cc.Dump = *dump
	cc.ClientId = *id

	tq := make([]proto.TopicQos, flag.NArg())
	for i := 0; i < flag.NArg(); i++ {
		tq[i].Topic = flag.Arg(i)
		tq[i].Qos = proto.QosAtMostOnce
	}

	if err := cc.Connect(*user, *pass); err != nil {
		fmt.Fprintf(os.Stderr, "connect: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connected with client id", cc.ClientId)
	cc.Subscribe(tq)

	for m := range cc.Incoming {
		fmt.Print(m.TopicName, "\t")
		m.Payload.WritePayload(os.Stdout)
		fmt.Println("\tr: ", m.Header.Retain)
	}
}
