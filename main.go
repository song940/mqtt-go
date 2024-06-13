package main

import (
	"os"

	"github.com/song940/mqtt-go/mqtt"
)

func main() {
	cmd := os.Args[1]
	switch cmd {
	case "server":
		mqtt.ListenAndServe(":1883")
	}
}
