package main

import (
"fmt"
"net"
"os"
"strconv"
"sync"
"time"
"sort"
)

func scango(all bool, addr string) {
	start, end := 1, 65535
	sp, ep := &start, &end
	
	if all == false {
		args := os.Args
		port, _ := strconv.Atoi(args[3])
		*sp, *ep = port, port
	}
	
	var wg sync.WaitGroup
	var open []string

	for start=start; start <= end; start++ {
		wg.Add(1)

		go func(start int, address string, all bool) {
			defer wg.Done()
			
			address = fmt.Sprintf("%s:%d", addr, start)
			_, err := net.DialTimeout("tcp", address,  3 * time.Second)
			
			if err == nil {
				element := fmt.Sprintf("[OK] %s : port %d", addr, start)
				open = append(open, element)
			}
		}(start, addr, all)
	}
	wg.Wait()
	
	if len(open) > 0 {
		sort.Strings(open)
		for _, element := range open {
        		fmt.Println(element)
    		}			
	}else {
		fmt.Println("[!] No open port : ",addr)
	}

}

func help() {
	fmt.Println("ScanGo:\n\t scango -o [IP] [port]")
	fmt.Println("\t scango -a [IP]")
	os.Exit(0)
}

func main() {
	args := os.Args
	all := true
	a := &all
	
	if len(args) >= 2 {
		if args[1] == "-o" {
			*a = false
		}else if args[1] == "-a"{
			*a = true
		}else {
			help()
		}

		addr := args[2]
		scango(all, addr)
	}else{
		help()
	}
}
