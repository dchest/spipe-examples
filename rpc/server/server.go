// Example of RPC server over spipe connection.

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/rpc"

	"github.com/dchest/spipe"
)

type Adder int

func (a *Adder) Add(args []int, reply *int) error {
	for _, v := range args {
		*reply += v
	}
	return nil
}

var (
	fAddr    = flag.String("a", ":8080", "service address")
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
	// Register RPC service.
	adder := new(Adder)
	rpc.Register(adder)
	// Listen.
	ln, err := spipe.Listen(key, "tcp", *fAddr)
	if err != nil {
		log.Fatalf("listen: %s", err)
	}
	// Accept RPC connections.
	log.Printf("Listening on %s", ln.Addr())
	rpc.Accept(ln)
}
