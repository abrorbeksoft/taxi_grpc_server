package main

import (
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	//lis,error:=net.Listen("tcp",":9000")
	//if error !=nil {
	//	log.Fatalf("error occured: %v",error)
	//}
	grpc.Dial(fmt.Sprintf("%s:%d","192.168.1.166",8001),grpc.WithInsecure())

	fmt.Println("Hello world")
}
