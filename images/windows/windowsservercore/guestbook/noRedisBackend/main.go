package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"fmt"

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
	environment["executable binary"] = os.Executable()
	environment["PID"] = os.Getpid()
	environment["PPID"] = os.Getppid()
	
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
