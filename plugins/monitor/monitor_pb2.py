# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: monitor/monitor.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x15monitor/monitor.proto\x12\x05proto\"\"\n\x05Order\x12\x0b\n\x03\x63md\x18\x01 \x01(\t\x12\x0c\n\x04path\x18\x02 \x01(\t\"$\n\x06Orders\x12\x0b\n\x03\x63md\x18\x01 \x01(\t\x12\r\n\x05paths\x18\x02 \x03(\t\"\'\n\x08\x43\x61llBack\x12\x0e\n\x06result\x18\x01 \x01(\x08\x12\x0b\n\x03msg\x18\x02 \x01(\t2_\n\x07Monitor\x12(\n\x07\x43ommand\x12\x0c.proto.Order\x1a\x0f.proto.CallBack\x12*\n\x08\x43ommands\x12\r.proto.Orders\x1a\x0f.proto.CallBackB\x0fZ\r./proto;protob\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'monitor.monitor_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\r./proto;proto'
  _ORDER._serialized_start=32
  _ORDER._serialized_end=66
  _ORDERS._serialized_start=68
  _ORDERS._serialized_end=104
  _CALLBACK._serialized_start=106
  _CALLBACK._serialized_end=145
  _MONITOR._serialized_start=147
  _MONITOR._serialized_end=242
# @@protoc_insertion_point(module_scope)
