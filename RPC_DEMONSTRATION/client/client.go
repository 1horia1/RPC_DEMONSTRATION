package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234") //clientul se conecteaza la serever
	if err != nil {
		fmt.Println(err)
		return
	}

	args := [2]int{5, 3} //parametrii pentru adunare
	var reply int

	err = client.Call("Calculator.Aduna", args, &reply) //apeleaza metoda Aduna a serviciului Calculator de pe serverul RPC
	//se foloseste pointer deoarece serverul va scrie rezultatul în variabila reply de pe client.
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Rezultat:", reply)
}
