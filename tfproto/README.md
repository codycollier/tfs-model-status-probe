
## Overview

The `./tfproto/` directory contains a golang package derived from the minimal set of proto needed to support calls to `ModelService.GetModelStatus()`.  


## Build

The golang generated files can be re-generated from the source proto like so:

* required: go get github.com/ckaznocha/protoc-gen-lint
* Run `make proto` in the parent directory (see Makefile)


Or manually like so:

```
$ go get github.com/ckaznocha/protoc-gen-lint
$ cd ./proto/tfproto/

# Run linter:
$ protoc --lint_out=./ -I. *.proto

# Run protoc and go generation:
$ protoc --go_out=. --go_opt=plugins=grpc -I. *proto

```


## Original source proto & adaptations

The src directory contains the original, minimal set of TensorFlow Serving proto and Tensorflow core proto required to generate the tfproto package.


Original sources:

* https://github.com/tensorflow/serving/tree/master/tensorflow_serving/apis/get_model_status.proto
* https://github.com/tensorflow/serving/tree/master/tensorflow_serving/apis/model.proto
* https://github.com/tensorflow/serving/tree/master/tensorflow_serving/apis/model_service.proto
* https://github.com/tensorflow/serving/tree/master/tensorflow_serving/util/status.proto
* https://github.com/tensorflow/tensorflow/blob/master/tensorflow/core/protobuf/error_codes.proto


The corresponding `.proto` files in `./tfproto/` were modified slightly from the original source to allow for easier protoc compilation and generation of simple, flat golang package. Changes included:

* changed proto package name
* updates to proto import paths
* addition/update to proto go_package definition
* removal of model management rpc and import in model_service.proto



