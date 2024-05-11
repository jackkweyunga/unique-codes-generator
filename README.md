
<h2 align="center">
    GRPC SERVER FOR UNIQUE CODE(S) GENERATION
</h2>

## Usage using ssl
- Python client
```shell
pip install ucg-client
```

- Add the following to your `example.py`
```python
import grpc
from ucg_client import generator_pb2, generator_pb2_grpc

# replace address with a domain address with https
# e.g ucg.example.com
# do not use path like example.com/ucg
channel = grpc.secure_channel('ADDRESS', grpc.ssl_channel_credentials())
stub = generator_pb2_grpc.GeneratorStub(channel)

# generate code
code = stub.GenerateUniqueCode(
    generator_pb2.GenerateUniqueCodeRequest()
)
print(code)

# generates codes
codes = stub.GenerateUniqueCodes(
    generator_pb2.GenerateUniqueCodesRequest(
        count=100
    )
)
print(codes)

```


## Development

server at: server/server.go

Run server:
```shell
cd scripts && go run server.go
```

python client at: clients/python/src/ucg_client
- Install the python client.
Build wheel
```shell
cd clients/python && sh build.sh
```

Install wheel
```shell
pip install "<path to wheel>"
```

```python 
import grpc
from ucg_client import generator_pb2, generator_pb2_grpc

channel = grpc.insecure_channel('localhost:5000')
stub = generator_pb2_grpc.GeneratorStub(channel)

# generate code
code = stub.GenerateUniqueCode(
    generator_pb2.GenerateUniqueCodeRequest()
)
print(code)

# generates codes
codes = stub.GenerateUniqueCodes(
    generator_pb2.GenerateUniqueCodesRequest(
        count=100
    )
)
print(codes)
```

