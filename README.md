# scanGo <img src="https://img.shields.io/badge/License-GPLv3-blue"></img>   <img src="https://img.shields.io/badge/go-up%20to%20date-red"></img>  <img src="https://img.shields.io/badge/version-0.3-yellow"></img>

## A simple TCP port scanner

## How to use
	git clone https://github.com/SAJAD-net/scanGo.git
	cd scanGo
	go build app/scanGo.go
	./scanGo -h
	
## Help
	[-i|--ip] is required                                                                                            
	usage: ScanGo [-h|--help] -i|--ip "<value>" [-a|--all] [-o|--one <integer>]                                      
															 
		      Simple and fast port scanner.                                                                      
															 
	Arguments:                                                                                                       
															 
	  -h  --help  Print help information                                                                             
	  -i  --ip    ip address                                                                                         
	  -a  --all   all possible ports                                                                                 
	  -o  --one   only a specific port  

## TODO
- [ ] Calculate the execution time
- [ ] Colorize the output
	
