package main

import (
    "fmt"
    "net"
    "flag"
    "os"
	"strings"
	"bufio"
	"slices"
	"github.com/korylprince/ipnetgen"
)

func main(){
    //var input string

	fullDetails:=  flag.Bool("full",false,"Get full details (IP and domain, if not set only print IP)")
	domainsFile := flag.String("domains","","File containing domains to check")
	scopeFile := flag.String("scope","","File containing scope ip or CIDR")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "[INSCOPE] : a binary to check if a list of DNS is in a list of IP scope \nYes out of scope is bad.\n")
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stderr, "\t-%v: %v\n", f.Name,f.Usage) // f.Name, f.Value
		})
	}

	flag.Parse()

	if(*scopeFile=="" || *domainsFile==""){
		fmt.Println("Define scope and domains files path. One per line.")
		os.Exit(-1)
	}

	var data string
	arrayOfIP := make([]string,0)

	readScopeFile, err:= os.Open(*scopeFile)
	if err!=nil{
		fmt.Println(err)
	}

	fileScannerScope :=bufio.NewScanner(readScopeFile)
	for fileScannerScope.Scan(){
		data = fileScannerScope.Text()
		if(strings.Contains(data,"/")){
			gen, _:= ipnetgen.New(strings.Trim(data," "))
			for ip:= gen.Next();ip!=nil; ip=gen.Next(){
				arrayOfIP = append(arrayOfIP,ip.String())
			}
		}else{
			arrayOfIP = append(arrayOfIP,data)
		}
	}

	readDomainFile, err:=os.Open(*domainsFile)
	if err!=nil{
		fmt.Println(err)
	}

	var domain string
	fileScannerDomain := bufio.NewScanner(readDomainFile)
	for fileScannerDomain.Scan(){
		domain = fileScannerDomain.Text()
		ips,_ := net.LookupIP(domain)
		for _,ip := range ips {
			if ipv4 := ip.To4(); ipv4!=nil {
				if(slices.Contains(arrayOfIP,ipv4.String())){
					if(*fullDetails){
						fmt.Println("[V] "+domain+" => "+ipv4.String())
					}else{
						fmt.Println(domain)
					}
				}else{
					if(*fullDetails){
						fmt.Println("[X] "+domain+" => "+ipv4.String())
					}
				}
			// }else{
			// 	fmt.Println("noIPV4 for "+domain)
			// }
			}
		}
	}
}