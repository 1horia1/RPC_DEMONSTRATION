package main

import (
	"fmt"
	"net"
	"net/rpc" //biblioteca pentru rpc
)

// Struct pentru serviciu
type Calculator int

//Calculator este tipul serviciului  , metodele RPC vor fi atasate acestui tip

// Metodă RPC

// Aduna este metoda RPC care primește 2 întregi și returnează suma
// c *Calculator = metoda apartine tipului Calculator
// args = cei doi parametri trimisi de client
// reply = pointer unde serverul scrie rezultatul pentru client
func (c *Calculator) Aduna(args [2]int, reply *int) error {
	*reply = args[0] + args[1] //calculeaza suma
	return nil
}

func main() {
	calc := new(Calculator) //cream o instanta a serviciului Calculator
	rpc.Register(calc)      // inregistram obiectul la RPC ; metodele sale devin apelabile de client

	ln, err := net.Listen("tcp", ":1234") //ascultam portul
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server RPC ascultă pe portul 1234...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn) // gestionează conexiunea
	}
}
