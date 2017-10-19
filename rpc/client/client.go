package main

	import (
		"fmt"
		"log"
		"net"
		"net/rpc"
		"os"
		"time"
	)
	
	type Args struct {
		A string
	}

	func main() {
		if len(os.Args) != 2 {
			fmt.Println("Usage: ", os.Args[0], "server")
			os.Exit(1)
		}
		serverAddress := os.Args[1]

		client, err := rpc.DialHTTP("tcp", serverAddress+":1919")
		if err != nil {
			log.Fatal("dialing:", err)
		}
		
		addrs, err := net.InterfaceAddrs()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var ipaddr string;
		for _, address := range addrs {
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ipaddr = ipnet.IP.String()
				}
			}
		}
	
		args := Args{ipaddr}
		var reply time.Time
		t_start := time.Now()
		err = client.Call("Arith.GetTime", args, &reply)
		t_end := time.Now()
		if err != nil {
			log.Fatal("arith error:", err)
		}
		
		t_dur := t_end.Sub(t_start)  //  calculate the duration time
		if !reply.IsZero() {
			fmt.Printf("The time got from the server: ");
			fmt.Println(reply)
			fmt.Printf("The time of transmission: ");
			fmt.Println(t_dur)
			fmt.Printf("The accurate time on the client: ");
			cur_time := reply.Add(t_dur/2) //   Because the time spent is bidirectional, we need to divide it by 2.
			fmt.Println(cur_time)
		} else {
			fmt.Printf("The server rejected the quest.\n")
		}

	}