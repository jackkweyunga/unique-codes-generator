
<h2 align="center">
    GRPC SERVER FOR UNIQUE CODE(S) GENERATION
</h2>


## Development

server at: server/server.go

Run server:
```shell
cd scripts && go run server.go
```

python client at: clients/python/example.go
Install the python client.

```shell
pip install ucg-client
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

