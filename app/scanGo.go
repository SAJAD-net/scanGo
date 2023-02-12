package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"net"
	"os"
	"sort"
	"sync"
	"time"
)

func scango(addr string, port int) {
	start, end := 1, 65535

	if port != 0 {
		start, end = port, port
	}

	var wg sync.WaitGroup
	var open []string

	for start = start; start <= end; start++ {
		wg.Add(1)

		go func(start int, address string) {
			defer wg.Done()

			address = fmt.Sprintf("%s:%d", addr, start)
			_, err := net.DialTimeout("tcp", address, 3*time.Second)

			if err == nil {
				element := fmt.Sprintf("[OK] %s : port %d", addr, start)
				open = append(open, element)
			}
		}(start, addr)
	}
	wg.Wait()

	if len(open) > 0 {
		sort.Strings(open)
		for _, element := range open {
			fmt.Println(element)
		}
	} else {
		fmt.Println("[!] No open port : ", addr)
	}

}

func help() {
	fmt.Println("ScanGo:\n\t scango -o [IP] [port]")
	fmt.Println("\t scango -a [IP]")
	os.Exit(0)
}

func main() {

	var addr string
	var port int

	parser := argparse.NewParser("ScanGo", "Simple and fast port scanner.")
	ip_arg := parser.String("i", "ip", &argparse.Options{Required: true, Help: "ip address"})
	all_arg := parser.Flag("a", "all", &argparse.Options{Required: false, Help: "all possible ports"})
	one_arg := parser.Int("o", "one", &argparse.Options{Required: false, Help: "only a specific port"})

	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(0)

	} else if *ip_arg != "" {
		addr = *ip_arg

		if *all_arg {
			port = 0
		} else if *one_arg != 0 {
			port = *one_arg
		} else {
			fmt.Print(parser.Usage(err))
			os.Exit(0)
		}
	}

	scango(addr, port)

}
