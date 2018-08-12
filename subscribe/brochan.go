package main

import (
    "fmt"
    "time"
)

type Subscriber struct{
    Name string
    Ch chan int
}

func (s *Subscriber) Listen() {
    for data := range s.Ch {
        fmt.Println(s.Name, "received", data)
    }
}


type BroadcastChan struct {
    Subscribers []chan int
}


func (b *BroadcastChan) Subscribe(s *Subscriber) {
    ch := make(chan int)

    b.Subscribers = append(b.Subscribers, ch)
    s.Ch = ch
    go s.Listen()
}

func (b *BroadcastChan) Broadcast(i int) {
    for _, s := range b.Subscribers {
        s <- i
    }
}

func (b *BroadcastChan) Shutdown() {
    for _, s := range b.Subscribers {
        close(s)
    }
}

func main() {

    b := &BroadcastChan{}

    s1 := &Subscriber{Name: "One"}
    s2 := &Subscriber{Name: "Two"}
    s3 := &Subscriber{Name: "Three"}

    b.Subscribe(s1)
    b.Subscribe(s2)
    b.Subscribe(s3)

    for i := 1; i < 10; i++ {
        b.Broadcast(i)
    }

    time.Sleep(100 * time.Millisecond)

    b.Shutdown()

    fmt.Println("DONE")


}
