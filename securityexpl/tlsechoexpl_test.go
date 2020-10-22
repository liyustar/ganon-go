package securityexpl_test

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"testing"
	"time"
)


func TestTLSEchoServer(t *testing.T) {
	cert, err := tls.LoadX509KeyPair("jan.newmarch.name.pem", "private.pem")
	checkError(err)
	config := tls.Config{Certificates: []tls.Certificate{cert}}

	now := time.Now()
	config.Time = func() time.Time { return now }
	config.Rand = rand.Reader

	service := "0.0.0.0:1200"

	listener, err := tls.Listen("tcp", service, &config)
	checkError(err)
	fmt.Println("Listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println("Accepted")
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		fmt.Println("Trying to read")
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
		}
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}

func TestTLSEchoClient(t *testing.T) {
	//if len(os.Args) != 2 {
	//	fmt.Println("Usage: ", os.Args[0], "host:port")
	//	os.Exit(1)
	//}
	service := "localhost:1200"

	conn, err := tls.Dial("tcp", service, nil)
	checkError(err)

	for n := 0; n < 10; n++ {
		fmt.Println("Writing...")
		conn.Write([]byte("Hello " + fmt.Sprint(n+48)))

		var buf [512]byte
		n, err := conn.Read(buf[0:])
		checkError(err)

		fmt.Println(string(buf[0:n]))
	}
	os.Exit(0)
}
