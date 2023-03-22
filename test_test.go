package web_server

import (
	"net/rpc"
	"testing"
)

func TestName(t *testing.T) {
	a := complex(2.0, 3.0)
	t.Log(a * a)
}

func TestRpc(t *testing.T) {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		t.Log(err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		t.Log(err)
	}

	t.Log(reply)
}
