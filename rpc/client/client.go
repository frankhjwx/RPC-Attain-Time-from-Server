package main

    import (
        "fmt"
        "log"
        "net/rpc"
        "os"
		"time"
    )
	
	type Args struct {
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
        
		args := Args{}
        var reply time.Time
		t_start := time.Now()
        err = client.Call("Arith.GetTime", args, &reply)
		t_end := time.Now()
        if err != nil {
            log.Fatal("arith error:", err)
        }
		
		t_dur := t_end.Sub(t_start)  //  calculate the duration time
        fmt.Printf("The time got from the server: ");
		fmt.Println(reply)
		fmt.Printf("The duration time: ");
		fmt.Println(t_dur)
		
		fmt.Printf("The accurate time on the client: ");
		cur_time := reply.Add(t_dur/2) //   Because the time spent is bidirectional, we need to divide it by 2.
		fmt.Println(cur_time)
        

    }