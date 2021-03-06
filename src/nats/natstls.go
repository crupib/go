package main
import (
    "log"
    "runtime"
    "github.com/nats-io/nats.go"
)
func main() {
    nc, err := nats.Connect("tls://demo.nats.io:4443")
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
     nc.Subscribe("greeting", func(m *nats.Msg) {
        log.Printf("[Received] %s", string(m.Data))
    })
    nc.Publish("greeting", []byte("hello world")) 

    runtime.Goexit()
}
