// RPC client

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/rpc"

	"github.com/dchest/spipe"
	"strconv"
)

var (
	fAddr    = flag.String("a", "localhost:8080", "remote address to call")
	fKeyFile = flag.String("k", "", "key file name")
)

func main() {
	flag.Parse()
	if *fKeyFile == "" {
		flag.Usage()
		return
	}
	// Read key file.
	key, err := ioutil.ReadFile(*fKeyFile)
	if err != nil {
		log.Fatalf("key file: %s", err)
	}

	// Dial.
	conn, err := spipe.Dial(key, "tcp", *fAddr)
	if err != nil {
		log.Fatalf("Dial: %s", err)
	}
	defer conn.Close()

	client := rpc.NewClient(conn)
	defer client.Close()

	// Call remote procedure.
	args := make([]int, flag.NArg())
	for i := 0; i < flag.NArg(); i++ {
		args[i], err = strconv.Atoi(flag.Arg(i))
		if err != nil {
			log.Fatalf("strconv: %s")
		}
	}
	var reply int
	if err := client.Call("Adder.Add", args, &reply); err != nil {
		log.Fatalf("%s", err)
	}

	// Print result.
	fmt.Printf("Result: %d", reply)
}
