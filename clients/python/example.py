import grpc

import generator_pb2, generator_pb2_grpc

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
print(codes.codes)
