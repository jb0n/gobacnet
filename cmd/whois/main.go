package main

import (
	"fmt"
	"log"
	"os"

	"gobacnet"
)

func main() {
	if len(os.Args) < 2 {
		log.Printf("usage %s <network interface>\n", os.Args[0])
		os.Exit(1)
	}
	iface := os.Args[1]

	cli, err := gobacnet.NewClient(iface, int(gobacnet.DefaultPort))
	if err != nil {
		log.Fatalf("error calling gobacnet.NewClient. err=%s", err)
	}
	devs, err := cli.WhoIs(-1, -1)
	if err != nil {
		log.Fatalf("error calling WhoIs. err=%s", err)
	}
	fmt.Printf("found %d devs\n", len(devs))
	for i := range devs {
		fmt.Printf("%d: %+v\n", i, devs[i])
	}
}
