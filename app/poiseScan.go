package main

import (
"fmt"
"net"
"os"
"strconv"
"sync"
"sort"
)

func portScan(all bool, addr string, port int) {
	start := 1
	end := 1024
	sp := &start
	ep := &end
	if all == false {
		*sp = port
		*ep = port
	}
	var wg sync.WaitGroup
	var opens []string
	var closes []string
	for start=start; start <= end; start++ {
		wg.Add(1)
		go func(start int, address string, all bool) {
			defer wg.Done()
			address = fmt.Sprintf("%s:%d", addr, start)
			_, err := net.Dial("tcp", address)
			if err == nil {
				op:=fmt.Sprintf("%s -> port %d open\n", addr, start)
				opens=append(opens, op)
			}else{
				cl:=fmt.Sprintf("%s -> port %d close\n", addr, start)
				closes=append(closes, cl)
			}
		}(start, addr, all)

	}
	wg.Wait()
	sort.Strings(opens)
	if len(opens) > 0 {
		fmt.Println("open's ports of", addr)
		for open := range opens {
			fmt.Printf("%d\n", opens[open])
		}

	}else {
		fmt.Printf("isn't any opens ports of %s !\n", addr)
	}
}

func main() {
	args := os.Args
	all := true
	a := &all
	if len(args) > 1{
		if args[1] == "-o" {
			*a = false
		}else if args[1] == "-a"{
			*a = true
		}
		addr := args[2]
		po := args[3]
		port, _ := strconv.Atoi(po)
		portScan(all, addr, port)
	}else{
		fmt.Println("enter valid args !")
	}
}
