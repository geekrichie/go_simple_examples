package main

import (
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"time"
)

type ZkClient struct {
	ConnectString []string
	zkConn *zk.Conn
	zkAcl []zk.ACL
}

const (
	 zkAclUser = "test"
	 zkAclPassword = "secret"
)
//新建客户端
func NewZkClient(connectString []string) *ZkClient {
	client := ZkClient{
		ConnectString: connectString,
		zkConn: nil,
		zkAcl: zk.DigestACL(zk.PermAll,zkAclUser, zkAclPassword ),
	}
	return &client
}

func (z *ZkClient) connect() error{
	//chEvent
	log.Println(z.ConnectString)
	conn, session, err := zk.Connect(z.ConnectString, 10*time.Second)
	//log.Printf("%#v", conn)
	log.Printf("%#v", err)
	if err != nil {
		log.Printf(" connect to the %v failure", z.ConnectString)
		return err
	}
	for event := range session {
		if event.State == zk.StateConnected {
			log.Printf("zookeeper State: %s\n", event.State)
			break
		}
	}

	auth :=zkAclUser + ":" + zkAclPassword
	if err := conn.AddAuth("digest", []byte(auth)); err != nil {
		z.Close()
		return err
	}

	z.zkConn  = conn
	return nil
}

func (client *ZkClient) Close() error {
	if client.zkConn != nil {
		client.zkConn.Close()
	}
	return nil
}

func (z *ZkClient) Create(path string, data []byte)  error{
	info, err := z.zkConn.Create(path, data, 0, z.zkAcl)
	log.Printf(info)
	return err
}
func (z *ZkClient) Get(path string) []byte{
	data, _, err:=z.zkConn.Get(path)
	if err != nil {
		log.Println(err)
		return nil
	}
	return data
}

func (z *ZkClient) Delete(path string) error{
	err:=z.zkConn.Delete(path, -1)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}


func (z *ZkClient) Exist(path string) (bool,error){
	ret, stat, err := z.zkConn.Exists(path)
	if err != nil {
		log.Println(err)
		return ret,err
	}
	log.Printf("%#v",stat)
	return ret, nil
}

func main() {
	//fmt.Println("hello world")
	connectString := []string{"192.168.33.13:2181"}
	client := NewZkClient(connectString)
	client.connect()
	ans,err := client.Exist("/path2")
	if err != nil {
		log.Println(err)
	}
	log.Println(ans)
	err = client.Create("/path3",[]byte("world"))
	if err != nil {
		log.Println(err)
	}
	val := client.Get("/path3")
	log.Println(string(val))
	client.Delete("/path2")
}