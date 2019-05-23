package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	var r net.Resolver
	tout := flag.Duration("t", time.Second, "Query timeout")
	interval := flag.Duration("i", time.Second, "Interval")
	hostname := flag.String("h", "", "Hostname")
	c := flag.Int("c", 1, "Concurrency")
	debug := flag.Bool("d", false, "Debug mode")
	flag.Parse()
	if *hostname == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	for x := 0; x < *c; x++ {
		go func() {
			for {
				ctx, _ := context.WithTimeout(context.TODO(), *tout)
				ips, err := r.LookupHost(ctx, *hostname)
				if err != nil {
					fmt.Printf("\nError: %v\n", err)
					fmt.Println(time.Now().Format(time.RFC850))
				} else {
					if *debug {
						fmt.Println(ips)
					} else {
						fmt.Print(".")
					}
				}
				time.Sleep(*interval)
			}
		}()
	}
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	os.Exit(0)
}
