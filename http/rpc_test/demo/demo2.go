package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Req struct {
	NumONe int
	NumTwo int
}
type Res struct {
	Rum int
}

func main() {
	req := Req{NumONe: 1, NumTwo: 2}
	var res Res
	client, err := rpc.DialHTTP("tcp", "localhost:8888")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	//client.Call("Server.Add", req, &res)
	//fmt.Println(res)
	call := client.Go("Server.Add", req, &res, nil)
	for {
		select {
		case <-call.Done:
			fmt.Println("运行得到的结果是：", res)
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("等待你")

		}
	}

}
