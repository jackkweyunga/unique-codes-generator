from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class Code(_message.Message):
    __slots__ = ("code",)
    CODE_FIELD_NUMBER: _ClassVar[int]
    code: str
    def __init__(self, code: _Optional[str] = ...) -> None: ...

class Codes(_message.Message):
    __slots__ = ("codes",)
    CODES_FIELD_NUMBER: _ClassVar[int]
    codes: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, codes: _Optional[_Iterable[str]] = ...) -> None: ...

class GenerateUniqueCodeRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GenerateUniqueCodesRequest(_message.Message):
    __slots__ = ("count",)
    COUNT_FIELD_NUMBER: _ClassVar[int]
    count: int
    def __init__(self, count: _Optional[int] = ...) -> None: ...
