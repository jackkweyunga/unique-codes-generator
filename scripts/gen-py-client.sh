SRC_DIR=$( cd .. && pwd )/protos
DST_DIR=$( cd .. && pwd )/clients/python/

python -m grpc_tools.protoc -I$SRC_DIR \
  --python_out=$DST_DIR --pyi_out=$DST_DIR --grpc_python_out=$DST_DIR \
  $SRC_DIR/generator.proto