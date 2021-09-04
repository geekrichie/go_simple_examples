package zerocopy

import (
	"syscall"
)

func sendfile(in_file string , out_file string) error{
	var err error
	infd, err := syscall.Open(in_file,syscall.O_RDONLY , 0)
	if err != nil {
		return err
	}
	defer func() {
		syscall.Close(infd)
	}()
	outfd, err := syscall.Open(out_file, syscall.O_RDWR | syscall.O_CREAT, syscall.S_IRWXU)
	if err != nil {
		return err
	}
	stat := syscall.Stat_t{}
	err = syscall.Fstat(infd,&stat)
	if err != nil {
		return err
	}
	var offset int64 = 0
	_, err = syscall.Sendfile(outfd, infd, &offset, int(stat.Size))
	if err != nil {
		return err
	}
	return nil
}

func mmap(file string) error{
	var err error
	infd, err := syscall.Open(file,syscall.O_RDONLY , 0)
	if err != nil {
		return err
	}
	defer func() {
		syscall.Close(infd)
	}()
	stat := syscall.Stat_t{}
	err = syscall.Fstat(infd,&stat)
	data, err := syscall.Mmap(infd, 0, int(stat.Size), syscall.PROT_READ, syscall.MAP_PRIVATE)
	if err != nil {
		return err
	}
	_, err = syscall.Write(syscall.Stdout, data)
	return err
}

