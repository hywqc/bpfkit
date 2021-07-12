package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hywqc/bpfkit/goebpf"
)

func handler(event goebpf.IEvent) {
	if event == nil {
		return
	}
	if event.Type() == goebpf.EVENTTYPE_PROCESS_EXIT {
		evt := event.(*goebpf.EventProcessExit)
		log.Printf("%s[%d]", evt.Comm, evt.Pid)
	}
}

func main() {
	err := goebpf.InitEbpf(handler)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancelFunc := context.WithCancel(context.Background())

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cancelFunc()
	}()

	go func() {
		goebpf.PollEvents(ctx, 1000)
		goebpf.FreeEbpf()
	}()

	<-ctx.Done()
	time.Sleep(time.Millisecond * 100)

	log.Println("done")
}
