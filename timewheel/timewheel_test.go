package timewheel

import (
	"log"
	"testing"
	"time"
)

func SayHello(msg interface{}) {
	log.Println(time.Now().Format("2006-01-02 15:04:05"))
	log.Println(msg)
}

func TestNewTimeWheel(t *testing.T) {
	SayHello("123")
	w := NewTimeWheel(2*time.Second,10, BasicJob(SayHello))
	go w.Start()
	w.AddTask("1",time.Now().Add(time.Second*10), "456")
	time.Sleep(10*time.Second)
	w.AddTask("2", time.Now().Add(time.Second*11), "789")
	select{}
}

func TestFloat(t *testing.T) {
	a := 9*time.Second
	b := 5*time.Second
	t.Log(int(a), int(b))
}