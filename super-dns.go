package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	var r net.Resolver
	tout := flag.Duration("t", time.Second, "Timeout duration")
	interval := flag.Duration("i", time.Second, "Interval duration")
	hostname := flag.String("h", "", "Hostname")
	debug := flag.Bool("d", false, "Debug mode")
	flag.Parse()
	if *hostname == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
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
}
