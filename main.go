// tfs_model_status_probe checks model status in a tensorflow serving service
//
// The tfs_model_status_probe is modeled after grpc_health_probe[1] and is
// intended for use as a kubernetes probe for a TensorFlow Serving service. It
// calls the ModelService.GetModelStatus() rpc for a given model and returns a
// response code indicating success.
//
// The ModelService service and GetModelStatus() rpc are built into the
// TensorFlow Serving grpc server[2].
//
//  [1]
//	https://github.com/grpc-ecosystem/grpc-health-probe
//  [2]
//	https://github.com/tensorflow/serving/blob/master/tensorflow_serving/apis/model_service.proto
//
//
package main

import (
	"flag"
	"log"
	"os"

	"google.golang.org/grpc"
)

var (
	flModel   = flag.String("model", "default", "The name of the model")
	flService = flag.String("service", "localhost:9000", "The hostname:port of the service")
)

// Call ModelService.GetModelStatus() and return response
func callTFS(model string, tfsclient string) string {
	log.Printf("model-name: %s\n", model)
	log.Printf("tfs-client: %s\n", tfsclient)
	return "foo"
}

// Parse the proto msg response and map to an appropriate return value
func checkResponse(response string) int {

	// Handle gRPC level errors

	// Handle servable states in the response
	// https://github.com/tensorflow/serving/blob/master/tensorflow_serving/apis/get_model_status.proto

	return 0
}

func main() {

	// Check command line args
	flag.Parse()
	serviceName := *flService
	modelName := *flModel

	// gRPC connection setup
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serviceName, opts...)
	if err != nil {
		log.Println("Error dialing grpc service")
		log.Printf("Error: %v", err)
		os.Exit(2)
	}
	defer conn.Close()

	// tfs grpc client setup
	// cl := pb.
	tfsClient := "foo"

	// call
	res := callTFS(modelName, tfsClient)

	// check
	retval := checkResponse(res)

	//
	os.Exit(retval)

}
