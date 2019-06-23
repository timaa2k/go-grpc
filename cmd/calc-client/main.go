package main

import (
	"fmt"
	"log"
    "net"
	"os"
	"sort"
	"strconv"

    "github.com/pkg/errors"
    "github.com/timaa2k/go-grpc/pkg/client/v1"
	"github.com/urfave/cli"
)

func validateFloat32(a string) (float32, error) {
	n, err := strconv.ParseFloat(a, 32)
	if err != nil {
		return 0, fmt.Errorf("Not a 32-bit floating point number")
	}
	return float32(n), nil
}

func parseCalcArgs(args cli.Args) (float32, float32, error) {
	if len(args) != 2 {
		return 0, 0, fmt.Errorf("Excepted exactly two arguments")
	}
	a, err := validateFloat32(args.Get(0))
	if err != nil {
		return 0, 0, errors.Wrapf(err, "Invalid argument `%s`", args.Get(0))
	}
	b, err := validateFloat32(args.Get(1))
	if err != nil {
		return 0, 0, errors.Wrapf(err, "Invalid argument `%s`", args.Get(1))
	}
	return a, b, nil
}

func main() {

    var app = cli.NewApp()

	app.Name = "Simple calc-server CLI"
	app.Usage = "An example CLI to send basic request to calc-server"
	app.Author = "timaa2k"
	app.Version = "0.1.0"
	app.HideVersion = true

    flags := []cli.Flag{
		cli.StringFlag{
			Name:   "host",
            Value:  "localhost",
			Usage:  "Calc server host to connect to",
			EnvVar: "CALC_SERVER_HOST",
		},
		cli.UintFlag{
			Name:   "port",
            Value:  7777,
			Usage:  "Calc server port to connect to",
			EnvVar: "CALC_SERVER_PORT",
		},
	}
    app.Flags = flags
	app.Commands = []cli.Command{
		{
			Name:      "add",
			Usage:     "Add two 32-bit floating point numbers",
			ArgsUsage: "a b",
            Flags: flags,
			Action: func(c *cli.Context) {
                host := c.String("host")
                port := c.Uint("port")
                serverAddr := net.JoinHostPort(host, strconv.Itoa(int(port)))

				a, b, err := parseCalcArgs(c.Args())
				if err != nil {
					fmt.Printf("%s\n", err.Error())
					os.Exit(1)
				}
                calcClient, err := v1.NewCalcClient(serverAddr)
                if err != nil {
                    fmt.Printf("%s\n", err.Error())
					os.Exit(1)
                }
                defer calcClient.Disconnect()
                result, err := calcClient.Add(a, b)
                if err != nil {
                    fmt.Printf("%s\n", err.Error())
					os.Exit(1)
                }
                fmt.Printf("%f\n", result)
			},
		},
		{
			Name:      "sub",
			Usage:     "Subtract two 32-bit floating point numbers",
			ArgsUsage: "a b",
            Flags: flags,
			Action: func(c *cli.Context) {
                host := c.String("host")
                port := c.Uint("port")
                serverAddr := net.JoinHostPort(host, strconv.Itoa(int(port)))

				a, b, err := parseCalcArgs(c.Args())
				if err != nil {
					fmt.Printf("%s\n", err.Error())
					os.Exit(1)
				}
                calcClient, err := v1.NewCalcClient(serverAddr)
                if err != nil {
                    fmt.Printf("%s\n", err.Error())
					os.Exit(1)
                }
                defer calcClient.Disconnect()
                result, err := calcClient.Sub(a, b)
                if err != nil {
                    fmt.Printf("%s\n", err.Error())
					os.Exit(1)
                }
                fmt.Printf("%f\n", result)
			},
		},
		{
			Name:      "mul",
			Usage:     "Multiply two 32-bit floating point numbers",
			ArgsUsage: "a b",
            Flags: flags,
			Action: func(c *cli.Context) {
                host := c.String("host")
                port := c.Uint("port")
                serverAddr := net.JoinHostPort(host, strconv.Itoa(int(port)))

				a, b, err := parseCalcArgs(c.Args())
				if err != nil {
					fmt.Printf("%s\n", err.Error())
					os.Exit(1)
				}
                calcClient, err := v1.NewCalcClient(serverAddr)
                if err != nil {
                    fmt.Printf("%s\n", err.Error())
					os.Exit(1)
                }
                defer calcClient.Disconnect()
                result, err := calcClient.Mul(a, b)
                if err != nil {
                    fmt.Printf("%s\n", err.Error())
					os.Exit(1)
                }
                fmt.Printf("%f\n", result)
			},
		},
		{
			Name:      "div",
			Usage:     "Divide a 32-bit floating point number by another one",
			ArgsUsage: "a b",
            Flags: flags,
			Action: func(c *cli.Context) {
                host := c.String("host")
                port := c.Uint("port")
                serverAddr := net.JoinHostPort(host, strconv.Itoa(int(port)))

				a, b, err := parseCalcArgs(c.Args())
				if err != nil {
					fmt.Printf("%s\n", err.Error())
					os.Exit(1)
				}
                calcClient, err := v1.NewCalcClient(serverAddr)
                if err != nil {
                    fmt.Printf("%s\n", err.Error())
					os.Exit(1)
                }
                defer calcClient.Disconnect()
                result, err := calcClient.Div(a, b)
                if err != nil {
                    fmt.Printf("%s\n", err.Error())
					os.Exit(1)
                }
                fmt.Printf("%f\n", result)
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
