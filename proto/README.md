# mqtt

An MQTT encoder and decoder,written in Golang.

This library was modified heavily from https://github.com/plucury/mqtt.go and
is API-incompatible with it.

Currently the library's API is unstable.

@Update: CONNACK with "session present" flag to support 3.1.1 compliance
@Update: Fixed header added to the Disconnect Decode for server to verify reserved bits
