package daemon

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestDaemonize(t *testing.T) {

	ret :=  daemonize(0,0)
	if ret != 0 {
		log.Fatal("failure")
	}

	resp ,_ := http.Get("http://www.baidu.com")
	content, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	f,_  := os.OpenFile("/tmp/a.txt",os.O_TRUNC| os.O_RDWR,0 )
	f.Write(content)
}

