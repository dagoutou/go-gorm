package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Server struct {
}
type Req struct {
	NumONe int
	NumTwo int
}
type Res struct {
	Rum int
}

func (s *Server) Add(req Req, res *Res) error {
	time.Sleep(5 * time.Second)
	res.Rum = req.NumONe + req.NumTwo
	return nil
}
func main() {
	rpc.Register(new(Server))
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	http.Serve(listen, nil)
}
