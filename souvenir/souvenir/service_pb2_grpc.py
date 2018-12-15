# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from souvenir import messages_pb2 as souvenir_dot_messages__pb2


class SouvenirStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Create = channel.unary_unary(
        '/souvenir.Souvenir/Create',
        request_serializer=souvenir_dot_messages__pb2.SouvenirCreateQuery.SerializeToString,
        response_deserializer=souvenir_dot_messages__pb2.SouvenirCreateResult.FromString,
        )
    self.Delete = channel.unary_unary(
        '/souvenir.Souvenir/Delete',
        request_serializer=souvenir_dot_messages__pb2.SouvenirDeleteQuery.SerializeToString,
        response_deserializer=souvenir_dot_messages__pb2.SouvenirDeleteResult.FromString,
        )
    self.Get = channel.unary_unary(
        '/souvenir.Souvenir/Get',
        request_serializer=souvenir_dot_messages__pb2.SouvenirGetQuery.SerializeToString,
        response_deserializer=souvenir_dot_messages__pb2.SouvenirGetResult.FromString,
        )
    self.Search = channel.unary_unary(
        '/souvenir.Souvenir/Search',
        request_serializer=souvenir_dot_messages__pb2.SouvenirSearchQuery.SerializeToString,
        response_deserializer=souvenir_dot_messages__pb2.SouvenirSearchResult.FromString,
        )
    self.GetGenres = channel.unary_unary(
        '/souvenir.Souvenir/GetGenres',
        request_serializer=souvenir_dot_messages__pb2.SouvenirEmptyQuery.SerializeToString,
        response_deserializer=souvenir_dot_messages__pb2.SouvenirGetGenresResult.FromString,
        )


class SouvenirServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def Create(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Delete(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Get(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Search(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetGenres(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_SouvenirServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Create': grpc.unary_unary_rpc_method_handler(
          servicer.Create,
          request_deserializer=souvenir_dot_messages__pb2.SouvenirCreateQuery.FromString,
          response_serializer=souvenir_dot_messages__pb2.SouvenirCreateResult.SerializeToString,
      ),
      'Delete': grpc.unary_unary_rpc_method_handler(
          servicer.Delete,
          request_deserializer=souvenir_dot_messages__pb2.SouvenirDeleteQuery.FromString,
          response_serializer=souvenir_dot_messages__pb2.SouvenirDeleteResult.SerializeToString,
      ),
      'Get': grpc.unary_unary_rpc_method_handler(
          servicer.Get,
          request_deserializer=souvenir_dot_messages__pb2.SouvenirGetQuery.FromString,
          response_serializer=souvenir_dot_messages__pb2.SouvenirGetResult.SerializeToString,
      ),
      'Search': grpc.unary_unary_rpc_method_handler(
          servicer.Search,
          request_deserializer=souvenir_dot_messages__pb2.SouvenirSearchQuery.FromString,
          response_serializer=souvenir_dot_messages__pb2.SouvenirSearchResult.SerializeToString,
      ),
      'GetGenres': grpc.unary_unary_rpc_method_handler(
          servicer.GetGenres,
          request_deserializer=souvenir_dot_messages__pb2.SouvenirEmptyQuery.FromString,
          response_serializer=souvenir_dot_messages__pb2.SouvenirGetGenresResult.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'souvenir.Souvenir', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
