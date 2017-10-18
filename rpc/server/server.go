package main

    import (
        "fmt"
        "net/http"
        "net/rpc"
		"time"
    )

	type Args struct {
	}
	
    type Arith int
	
    func (t *Arith) GetTime(args *Args, reply *time.Time) error {
        *reply = time.Now()
        return nil
    }

    func main() {

        arith := new(Arith)
        rpc.Register(arith)
        rpc.HandleHTTP()

        err := http.ListenAndServe(":1919", nil)
        if err != nil {
            fmt.Println(err.Error())
        }
    }