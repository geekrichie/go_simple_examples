package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)


var daemonize bool
var fork int
var children []int

func init() {
	flag.BoolVar(&daemonize, "daemon", false, "run the children in daemon mode")
	flag.IntVar(&fork, "fork", 1, "the children number")
}
func main() {
     flag.Parse()
     pid := syscall.Getpid()
     ppid := syscall.Getppid()
     log.Printf("current process pid = [%d] , parent process pid = [%d]",pid, ppid)
     log.Printf("daemonize = %v, fork = %d", daemonize, fork)
     //获取环境变量
     //if _, isdaemon := os.LookupEnv("daemon"); !isdaemon {
     //
	 //}

	 go func() {
		 sig := make(chan os.Signal)
		 signal.Notify(sig)
		 for s := range sig {
			 // see https://u.kfd.me/33
			 // SIGINT means graceful stop
			 // SIGTERM means graceful [or not], cleanup something
			 // SIGQUIT SIGKILL means immediately shutdown
			 if s == syscall.SIGQUIT || s == syscall.SIGTERM {
				 log.Printf("[%d] exit\n", pid)
				 // make sure that parent can send signals to the children
				 for _, child := range children {
					 log.Printf("parent send %s to %d", s, child)
					 syscall.Kill(child, s.(syscall.Signal))
				 }
				 syscall.Exit(0)
			 }
		 }
	 }()
	 if _, isChild := os.LookupEnv("child_id"); !isChild {
	 	if daemonize {
			if _, isDaemon := os.LookupEnv("daemon"); !isDaemon {
				daemonenv := append(os.Environ(), "daemon=true")
				childpid, err := syscall.ForkExec(os.Args[0], os.Args, &syscall.ProcAttr{
					Env: daemonenv,
					Sys: &syscall.SysProcAttr{
						Setsid: true,
					},
					Files: []uintptr{
						os.Stdin.Fd(), os.Stdout.Fd(), os.Stdout.Fd(),
					},
				})
				if err != nil {
					log.Printf("parent process:[%d] fork child process error", pid)
					os.Exit(0)
				}
				log.Printf(" parent process:[%d] fork child[%d] process success", pid, childpid)
				return
			}
			log.Printf("daemonized process [%d] is launching", pid)
		}
		 for i := 0 ; i < fork ; i++ {
		 	 childId := fmt.Sprintf("child_id=%d", i)

			 childpid, err := syscall.ForkExec(os.Args[0], os.Args, &syscall.ProcAttr{
				 Env: append(os.Environ(), childId),
				 Sys: &syscall.SysProcAttr{
					 Setsid: true,
				 },
				 Files: []uintptr{
					 os.Stdin.Fd(),os.Stdout.Fd(),os.Stdout.Fd(),
				 },
			 })
			 if err != nil {
				 log.Printf("parent process:[%d] fork child process error", pid)
				 os.Exit(0)
			 }
			 log.Printf("parent %d fork %d", pid, childpid)
			 if childpid != 0 {
				 children = append(children, childpid)
			 }
			 log.Printf(" parent process:[%d] fork child[%d] process success",pid,childpid)
		 }
	 }

	 var sig = make(chan os.Signal)
	 signal.Notify(sig)
     for signal := range sig{
     	log.Printf("recevied signal : %s", signal)
	 }
}