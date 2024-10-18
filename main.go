package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Healthy",
		Usage: "A CLI tool for checking the health of your services",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "The domain to check",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "The port number to check",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "Loop",
				Aliases:  []string{"l"},
				Usage:    "Loop the check every 5 seconds",
				Required: false,
				Value:    true,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.String("port")
			loop := c.Bool("Loop")

			if port == "" {
				port = "80"
			}

			if loop {
				for {
					fmt.Println(PerformHealthCheck(c.String("domain"), port))
					time.Sleep(5 * time.Second)
					fmt.Println(strings.Repeat("-", 30))
				}
			} else {
				fmt.Println(PerformHealthCheck(c.String("domain"), port))
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("Error running the app", err)
	}
}

func PerformHealthCheck(domain, port string) string {
	address := domain + ":" + port
	timeout := time.Duration(5 * time.Second)

	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable, error: %v", address, err)
		status = color.RedString(status)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable\nFrom: %v\nTo: %v", address, conn.LocalAddr(), conn.RemoteAddr())
		status = color.GreenString(status)
	}

	defer conn.Close()

	return status
}
