SRC_DIR=$( cd .. && pwd )/protos
DST_DIR=$( cd .. && pwd )

protoc -I=$SRC_DIR --go_out=$DST_DIR --go-grpc_out=$DST_DIR $SRC_DIR/generator.proto
