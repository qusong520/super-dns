package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	var r net.Resolver
	if len(os.Args) < 4 {
		fmt.Printf("Execution error:\n%v <hostname> <interval> <tout> [debug]", os.Args[0])
		os.Exit(1)
	}
	dur, err := time.ParseDuration(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	tout, err := time.ParseDuration(os.Args[3])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for {
		ctx, _ := context.WithTimeout(context.TODO(), tout)
		ips, err := r.LookupHost(ctx, os.Args[1])
		if err != nil {
			fmt.Printf("\nError: %v", err)
		} else {
			fmt.Print(".")
			if len(os.Args) > 4 && os.Args[4] == "debug" {
				fmt.Println(ips)
			}
		}
		time.Sleep(dur)
	}
}
