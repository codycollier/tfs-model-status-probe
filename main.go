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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/codycollier/tfs-model-status-probe/tfproto/tfproto"
)

var (
	flModelName      = flag.String("model-name", "default", "The name of the model")
	flModelVersion   = flag.Int64("model-version", 0, "The version of the model")
	flAddr           = flag.String("addr", "localhost:9000", "The hostname:port to check")
	flConnectTimeout = flag.Duration("connect-timeout", time.Second*3, "Timeout for making connection")
	flRpcTimeout     = flag.Duration("rpc-timeout", time.Second*10, "Timeout for rpc call")
)

// Call ModelService.GetModelStatus() and return response
func callModelStatus(ctx context.Context, client tfproto.ModelServiceClient, model string) (*tfproto.GetModelStatusResponse, error) {
	request := &tfproto.GetModelStatusRequest{
		ModelSpec: &tfproto.ModelSpec{
			Name: model,
		},
	}
	response, err := client.GetModelStatus(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Parse the proto msg response and map to an appropriate return value
func checkServableResponse(response *tfproto.GetModelStatusResponse, modelVersion int64) int {

	// Ensure non-empty response
	if len(response.ModelVersionStatus) == 0 {
		log.Println("Empty response")
		return 11
	}

	// Get the state for the noted version. If no version, take the first.
	var status tfproto.ModelVersionStatus_State
	statusFound := false
	if modelVersion == 0 {
		status = response.ModelVersionStatus[0].State
		statusFound = true
	} else {
		for _, res := range response.ModelVersionStatus {
			if modelVersion == res.Version {
				status = res.State
				statusFound = true
				break
			}
		}
	}

	// No matching version found? Return early.
	if !statusFound {
		log.Printf("No matching response found for version: %v\n", modelVersion)
		return 12
	}

	// Map servable states to return value
	// https://github.com/tensorflow/serving/blob/master/tensorflow_serving/apis/get_model_status.proto
	var retval int
	switch status {
	case tfproto.ModelVersionStatus_AVAILABLE:
		// servable is up and ready
		log.Println("Servable state is AVAILABLE")
		retval = 0
	case tfproto.ModelVersionStatus_UNKNOWN:
		log.Println("Servable state is UNKNOWN")
		retval = 30
	case tfproto.ModelVersionStatus_START:
		log.Println("Servable state is START")
		retval = 31
	case tfproto.ModelVersionStatus_LOADING:
		log.Println("Servable state is LOADING")
		retval = 32
	case tfproto.ModelVersionStatus_UNLOADING:
		log.Println("Servable state is UNLOADING")
		retval = 33
	case tfproto.ModelVersionStatus_END:
		log.Println("Servable state is END")
		retval = 34
	default:
		log.Println("Servable state is unexpected")
		retval = 100 // unexpected
	}

	return retval
}

func main() {

	// Process command line args
	flag.Parse()
	addr := *flAddr
	modelName := *flModelName
	modelVersion := *flModelVersion
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
	log.Printf("ModelStatusResponse: %v\n", modelStatusResponse)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			log.Printf("Model not found: %v\n", err)
			os.Exit(10)
		}
		log.Printf("Error calling tfs: %v\n", err)
		os.Exit(3)
	}

	// check response for servable status
	retval := checkServableResponse(modelStatusResponse, modelVersion)
	os.Exit(retval)

}
