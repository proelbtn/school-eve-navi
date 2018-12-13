# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: souvenir/service.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from souvenir import messages_pb2 as souvenir_dot_messages__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='souvenir/service.proto',
  package='souvenir',
  syntax='proto3',
  serialized_options=_b('Z;github.com/proelbtn/school-eve-navi/gateway/protos/souvenir'),
  serialized_pb=_b('\n\x16souvenir/service.proto\x12\x08souvenir\x1a\x17souvenir/messages.proto2\xfc\x02\n\x08Souvenir\x12I\n\x06\x43reate\x12\x1d.souvenir.SouvenirCreateQuery\x1a\x1e.souvenir.SouvenirCreateResult\"\x00\x12H\n\x06\x44\x65lete\x12\x1d.souvenir.SouvenirDeleteQuery\x1a\x1d.souvenir.SouvenirDeleteQuery\"\x00\x12@\n\x03Get\x12\x1a.souvenir.SouvenirGetQuery\x1a\x1b.souvenir.SouvenirGetResult\"\x00\x12I\n\x06Search\x12\x1d.souvenir.SouvenirSearchQuery\x1a\x1e.souvenir.SouvenirSearchResult\"\x00\x12N\n\tGetGenres\x12\x1c.souvenir.SouvenirEmptyQuery\x1a!.souvenir.SouvenirGetGenresResult\"\x00\x42=Z;github.com/proelbtn/school-eve-navi/gateway/protos/souvenirb\x06proto3')
  ,
  dependencies=[souvenir_dot_messages__pb2.DESCRIPTOR,])



_sym_db.RegisterFileDescriptor(DESCRIPTOR)


DESCRIPTOR._options = None

_SOUVENIR = _descriptor.ServiceDescriptor(
  name='Souvenir',
  full_name='souvenir.Souvenir',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=62,
  serialized_end=442,
  methods=[
  _descriptor.MethodDescriptor(
    name='Create',
    full_name='souvenir.Souvenir.Create',
    index=0,
    containing_service=None,
    input_type=souvenir_dot_messages__pb2._SOUVENIRCREATEQUERY,
    output_type=souvenir_dot_messages__pb2._SOUVENIRCREATERESULT,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Delete',
    full_name='souvenir.Souvenir.Delete',
    index=1,
    containing_service=None,
    input_type=souvenir_dot_messages__pb2._SOUVENIRDELETEQUERY,
    output_type=souvenir_dot_messages__pb2._SOUVENIRDELETEQUERY,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Get',
    full_name='souvenir.Souvenir.Get',
    index=2,
    containing_service=None,
    input_type=souvenir_dot_messages__pb2._SOUVENIRGETQUERY,
    output_type=souvenir_dot_messages__pb2._SOUVENIRGETRESULT,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Search',
    full_name='souvenir.Souvenir.Search',
    index=3,
    containing_service=None,
    input_type=souvenir_dot_messages__pb2._SOUVENIRSEARCHQUERY,
    output_type=souvenir_dot_messages__pb2._SOUVENIRSEARCHRESULT,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetGenres',
    full_name='souvenir.Souvenir.GetGenres',
    index=4,
    containing_service=None,
    input_type=souvenir_dot_messages__pb2._SOUVENIREMPTYQUERY,
    output_type=souvenir_dot_messages__pb2._SOUVENIRGETGENRESRESULT,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_SOUVENIR)

DESCRIPTOR.services_by_name['Souvenir'] = _SOUVENIR

# @@protoc_insertion_point(module_scope)
