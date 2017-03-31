package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"fmt"
	"strconv"
	"net"
	
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func HandleCommands(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s\n", "HandleCommands was called")
	key := mux.Vars(req)["key"]
	value := mux.Vars(req)["value"]
	fmt.Printf("KeyValue Pair: %s - %s\n", key, value)
	
	os.Exit(23)
}

func EnvHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s\n", "EnvHandler was called")
	environment := make(map[string]string)
	for _, item := range os.Environ() {
		splits := strings.Split(item, "=")
		key := splits[0]
		val := strings.Join(splits[1:], "=")
		environment[key] = val
	}
	//ex := os.Executable()
	//environment["executable binary"] = ex
	pid := strconv.Itoa(os.Getpid())
	ppid := strconv.Itoa(os.Getppid())
	environment["PID"] = pid
	environment["PPID"] = ppid
	
	ifaces, err := net.Interfaces()
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
				case *net.IPNet:
						ip = v.IP
						environment["IPNET"] = ip.String()
				case *net.IPAddr:
						ip = v.IP
						environment["IPAddr"] = ip.String()
			}	
		}
	}

	envJSON := HandleError(json.MarshalIndent(environment, "", "  ")).([]byte)
	rw.Write(envJSON)
}

func HandleError(result interface{}, err error) (r interface{}) {
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	r := mux.NewRouter()
	r.Path("/rpush/{key}/{value}").Methods("GET").HandlerFunc(HandleCommands)	
	r.Path("/env").Methods("GET").HandlerFunc(EnvHandler)

	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":3000")
}
