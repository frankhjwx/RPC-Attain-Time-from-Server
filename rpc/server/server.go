package main

	import (
		"fmt"
		"net/http"
		"net/rpc"
		"time"
		"bufio" 
		"io"
		"os"
		"log"
	)

	type Args struct {
		A string
	}
	
	type Arith int
	
	func (t *Arith) GetTime(args *Args, reply *time.Time) error {  // main function of attaining time from server
		log.Printf("Received request from %s\n",args.A)
		if (!checkIpAuthorized(args.A)) {
			log.Printf("Sorry, this client hasn't been authorized!\n(If you want to authorize this server, please add this IP address to \"authorized.txt\")\n\n")
		} else {
			log.Printf("This is an authorized client.\n\n")
			*reply = time.Now()
		}
		return nil
	}
	
	func checkIpAuthorized(ip string) (bool) {  //  1: authorized 0: unauthorized
		var filename = "authorized.txt"
		var f *os.File
		var err error
		
		if checkFileIsExist(filename) {
			f, err = os.OpenFile(filename, os.O_APPEND, 0666)
		} else {
			f, err = os.Create(filename)
		}
		if err != nil {
			fmt.Println(err.Error())
		}
		
		buf := bufio.NewReader(f)
		for {
			line, _, c := buf.ReadLine()
			if c == io.EOF {
				break
			}
			if string(line) == ip {
				f.Close()
				return true
			}
		}
		f.Close()
		return false
	}
	
	func checkFileIsExist(filename string) (bool) {
		var exist = true;
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			exist = false;
		}
		return exist;
	}	

	func main() {

		arith := new(Arith)
		rpc.Register(arith)
		rpc.HandleHTTP()

		log.Println("The server has been started!\n")
		err := http.ListenAndServe(":1919", nil)
		
		if err != nil {
			fmt.Println(err.Error())
		}
		
	}