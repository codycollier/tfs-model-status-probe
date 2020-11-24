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
	"context"
	"flag"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	"github.com/codycollier/tfs-model-status-probe/tfproto/tfproto"
)

var (
	flModel          = flag.String("model", "default", "The name of the model")
	flAddr           = flag.String("addr", "localhost:9000", "The hostname:port to check")
	flConnectTimeout = flag.Duration("connect-timeout", time.Second*3, "Timeout for making connection")
	flRpcTimeout     = flag.Duration("rpc-timeout", time.Second*10, "Timeout for rpc call")
)

// Call ModelService.GetModelStatus() and return response
func callModelStatus(ctx context.Context, client tfproto.ModelServiceClient, model string) (*tfproto.GetModelStatusResponse, error) {
	request := &tfproto.GetModelStatusRequest{}
	response, err := client.GetModelStatus(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Parse the proto msg response and map to an appropriate return value
func checkResponse(response *tfproto.GetModelStatusResponse) int {

	// Handle gRPC level errors

	// Handle servable states in the response
	// https://github.com/tensorflow/serving/blob/master/tensorflow_serving/apis/get_model_status.proto

	return 0
}

func main() {

	// Process command line args
	flag.Parse()
	addr := *flAddr
	modelName := *flModel
	connectTimeout := *flConnectTimeout
	rpcTimeout := *flRpcTimeout

	// set a timeout on the connection
	ctxDial, cancelDial := context.WithTimeout(context.Background(), connectTimeout)
	defer cancelDial()

	// grpc connection
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.DialContext(ctxDial, addr, opts...)
	if err != nil {
		log.Printf("Error dialing grpc service: %v\n", err)
		os.Exit(2)
	}
	defer conn.Close()

	// grpc client
	client := tfproto.NewModelServiceClient(conn)

	// set a timeout on the rpc
	ctxRpc, cancelRpc := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancelRpc()

	// call model status
	modelStatusResponse, err := callModelStatus(ctxRpc, client, modelName)
	if err != nil {
		log.Printf("Error calling tfs: %v\n", err)
		os.Exit(3)
	}

	// check response for servable status
	retval := checkResponse(modelStatusResponse)

	//
	os.Exit(retval)

}
