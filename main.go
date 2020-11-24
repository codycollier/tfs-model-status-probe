// tfs_model_status_probe checks model status in a tensorflow serving instance
//
//
package main

import (
	"flag"
	"log"
	"os"
)

var (
	flModel = flag.String("model", "default", "The name of the model")
	flHost  = flag.String("host", "localhost", "The hostname of the tensorflow serving instance")
	flPort  = flag.Int("port", 9000, "The grpc port for the tensorflow serving instance")
)

// Call ModelService.GetModelStatus() and return response
func callTFS(model, host string, port int) string {
	log.Printf("model: %s\n", model)
	log.Printf("host: %s\n", host)
	log.Printf("port: %v\n", port)
	return "foo"
}

func checkResponse(response string) int {
	return 0
}

func main() {

	// ...
	flag.Parse()

	// call
	res := callTFS(*flModel, *flHost, *flPort)

	// check
	retval := checkResponse(res)

	//
	os.Exit(retval)

}
