
## tfs_model_status_probe - TensorFlow Model Status Probe

![ci](https://github.com/codycollier/tfs-model-status-probe/workflows/ci/badge.svg)
![release--](https://github.com/codycollier/tfs-model-status-probe/workflows/release/badge.svg)



The `tfs_model_status_probe` checks the model status in a TensorFlow Serving [1] instance.  The probe is modeled after `grpc_health_probe` [2] and is intended for use as a kubernetes probe for a TensorFlow Serving service.

The tfs model status probe calls the ModelService.GetModelStatus() rpc [3] for a given model.  If the model is `AVAILABLE`, then the probe will have an exit code of 0, conforming to k8s probe conventions.  If the model is still `LOADING`, in some other state, or there are grpc communication errors then the exit code will be non-zero.  If no version is provided, the probe assumes a single version is loaded and just checks the first version in the response.


#### Usage

```
$ ./tfs_model_status_probe -help
Usage of ./tfs_model_status_probe:
  -addr string
    	The hostname:port to check (default "localhost:9000")
  -connect-timeout duration
    	Timeout for making connection (default 3s)
  -model-name string
    	The name of the model (default "default")
  -model-version int
    	The version of the model
  -rpc-timeout duration
    	Timeout for rpc call (default 10s)
```


## Examples

Here are a handful of the more common examples of success and failure calls.


Successful call with model status AVAILABLE (exit code 0):
```
$ ./tfs_model_status_probe -addr="localhost:8500" -model-name="half_plus_two"
2020/11/30 19:49:33 ModelStatusResponse: model_version_status:{version:123 state:AVAILABLE status:{}}
2020/11/30 19:49:33 Servable state is AVAILABLE

$ echo $?
0
```


Error calling with wrong model name (exit code 10):
```
$ ./tfs_model_status_probe -addr="localhost:8500" -model-name="no-such-model"
2020/11/30 19:51:27 ModelStatusResponse: <nil>
2020/11/30 19:51:27 Model not found: rpc error: code = NotFound desc = Could not find any versions of model no-such-model

$ echo $?
10
```


Error when unable to communicate with service (exit code 2):
```
$ ./tfs_model_status_probe -addr="localhost:1234" -model-name="half_plus_two"
2020/11/30 19:52:57 Error dialing grpc service: context deadline exceeded

$ echo $?
2
```


The examples are run using a test server like so:
```
$ git clone https://github.com/tensorflow/serving.git tensorflow-serving
$ cd tensorflow-serving/
$ TESTDATA="$(pwd)/tensorflow_serving/servables/tensorflow/testdata"
$ TESTMODEL="$TESTDATA/saved_model_half_plus_two_cpu:/models/half_plus_two"
$ docker run -t --rm -p 8500:8500 -p 8501:8501 -v "$TESTMODEL" -e MODEL_NAME=half_plus_two tensorflow/serving
```



## References


[1] TensorFlow Serving

[TensorFlow Serving](https://github.com/tensorflow/serving) is a server designed to load a TensorFlow model and provide access to it via gRPC and/or HTTP.


[2] gRPC Health Probe

As noted previously, the tfs model status probe is inspired by the [grpc_health_probe](https://github.com/grpc-ecosystem/grpc-health-probe). The github actions files, goreleaser config, and small portions of the README were derived and adapted from the grpc-health-probe project.


[3] TFS Model Service

The [ModelService](https://github.com/tensorflow/serving/blob/master/tensorflow_serving/apis/model_service.proto) is a [gRPC](https://grpc.io/) service with a `GetModelStatus()` rpc which is built-in to TensorFlow Serving.  The TensorFlow proto required to call the rpc is sourced from TensorFlow Serving and TensorFlow core.  See `./tfproto/` for more details.

    

