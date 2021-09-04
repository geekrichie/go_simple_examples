package daemon

import (
	"os"
	"syscall"
)

func daemonize(nochdir, noclose int)int {
	pid, _, err := syscall.Syscall(syscall.SYS_FORK,0,0,0)
	if err != 0 {
		return -1
	}
	if pid != 0 {
		os.Exit(0)
	}
	syscall.Setsid()
	pid, _ , err = syscall.Syscall(syscall.SYS_FORK,0,0,0)
	if err != 0 {
		return -1
	}
	if pid != 0 {
		os.Exit(0)
	}
	var rlimit syscall.Rlimit
	err1 := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit)
	if err1 != nil {
		return -1
	}
	for  i  := uint64(6); i < rlimit.Max ;i ++ {
		_ = syscall.Close(int(i))
	}
	if nochdir == 0 {
		syscall.Chdir("/")
	}
	if noclose == 0{
		fd, _ := syscall.Open("/dev/null", syscall.O_RDWR, 0)
		syscall.Dup2(fd, int(os.Stdin.Fd()))
		syscall.Dup2(fd, int(os.Stdout.Fd()))
		syscall.Dup2(fd, int(os.Stderr.Fd()))
		if fd > int(os.Stderr.Fd()){
			syscall.Close(fd)
		}
	}
	return 0
}