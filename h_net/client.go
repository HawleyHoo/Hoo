package main

import (

	"os"
	"fmt"

	"net/rpc"
	"log"
)

type Args struct {
	A, B int
}


type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "server")
		os.Exit(1)
	}

	serverAddr := os.Args[1]
	fmt.Println(serverAddr)

	client, err := rpc.DialHTTP("tcp", serverAddr+":1234")
	if err != nil {
		 log.Fatal("dialing:", err)
	}
	args := Args{17, 8}
	var reply int
	err = client.Call("Math.Multiply", args, &reply)
	if err != nil {
		log.Fatal("math error:", err)
	}

	fmt.Printf("math : %d * %d = %d \n", args.A, args.B, reply)

	var quo Quotient
	err = client.Call("Math.Divide", args, &quo)
	if err != nil {
		log.Fatal("math error:", err)
	}

	fmt.Printf("math : %d / %d = %d remainder %d \n", args.A, args.B, quo.Quo, quo.Rem)
}

/*

type Math int


func (m *Math) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}



func (m *Math) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}*/
