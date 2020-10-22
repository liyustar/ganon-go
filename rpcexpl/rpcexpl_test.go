
package rpcexpl_test

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"testing"

	"github.com/liyustar/nuts/rpcexpl"
)

func TestRpcHttpServer(t *testing.T) {
	fmt.Println("start server.")
	// 启动服务端
	arith := new(rpcexpl.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestRpcHttpClient(t *testing.T) {
	fmt.Println("start client.")
	// 启动客户端
	//if len(os.Args) != 2 {
	//	fmt.Println("Usage: ", os.Args[0], "server")
	//	os.Exit(1)
	//}
	serverAddress := "localhost"

	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := rpcexpl.Args{A: 17, B: 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot rpcexpl.Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func TestRpcTcpServer(t *testing.T) {
	arith := new(rpcexpl.Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	/* This works:
	   rpc.Accept(listener)
	*/
	/* and so does this:
	 */
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}

func TestRpcTcpClient(t *testing.T) {
	//if len(os.Args) != 2 {
	//	fmt.Println("Usage: ", os.Args[0], "server:port")
	//	os.Exit(1)
	//}
	service := "localhost:1234"

	client, err := rpc.Dial("tcp", service)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := rpcexpl.Args{A: 17, B: 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot rpcexpl.Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}

func TestRpcJsonServer(t *testing.T) {
	arith := new(rpcexpl.Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	/* This works:
	   rpc.Accept(listener)
	*/
	/* and so does this:
	 */
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}

func TestRpcJsonClient(t *testing.T) {
	//if len(os.Args) != 2 {
	//	fmt.Println("Usage: ", os.Args[0], "server:port")
	//	log.Fatal(1)
	//}
	service := "localhost:1234"

	client, err := jsonrpc.Dial("tcp", service)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := rpcexpl.Args{A: 17, B: 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot rpcexpl.Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
