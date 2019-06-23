package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/peterbourgon/ff"
	"github.com/timaa2k/go-grpc/pkg/cmd"
)

func main() {
	fs := flag.NewFlagSet("calc-server", flag.ExitOnError)
	var (
		host = fs.String("host", "localhost", "listen address")
		port = fs.Uint("port", 7777, "port number")
		_    = fs.String("config", "", "config file (optional)")
	)

	err := ff.Parse(fs, os.Args[1:],
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(ff.PlainParser),
		ff.WithEnvVarPrefix("CALC_SERVER"),
	)
	if err != nil {
		log.Fatal(err)
	}

	addr := net.JoinHostPort(*host, strconv.Itoa(int(*port)))
	if err := cmd.RunServer(addr); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
