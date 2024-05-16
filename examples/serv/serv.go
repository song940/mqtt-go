package main

import "github.com/song940/mqtt-go/mqtt"

func main() {
	mqtt.ListenAndServe(":1883")
}
