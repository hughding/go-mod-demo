package main

import (
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("外部包测试：", grpc.Server{})
}
