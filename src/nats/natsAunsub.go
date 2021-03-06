package main
import (
    "log"
    "runtime"
    "github.com/nats-io/nats.go"
)
func main() {
    nc, err := nats.Connect("nats://localhost:4222")
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
    sub, err := nc.Subscribe("greeting", func(m *nats.Msg) {
        log.Printf("[Received] %s", string(m.Data))
    })
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
    sub.AutoUnsubscribe(5)
    for i := 0; i < 10; i++ {
        nc.Publish("greeting", []byte("hello world!!!")) 
    }
    runtime.Goexit()
}
