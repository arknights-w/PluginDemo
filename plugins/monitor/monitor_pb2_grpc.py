# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from plugins.monitor import monitor_pb2 as monitor_dot_monitor__pb2


class MonitorStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Command = channel.unary_unary(
                '/proto.Monitor/Command',
                request_serializer=monitor_dot_monitor__pb2.Order.SerializeToString,
                response_deserializer=monitor_dot_monitor__pb2.CallBack.FromString,
                )
        self.Commands = channel.unary_unary(
                '/proto.Monitor/Commands',
                request_serializer=monitor_dot_monitor__pb2.Orders.SerializeToString,
                response_deserializer=monitor_dot_monitor__pb2.CallBack.FromString,
                )


class MonitorServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Command(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Commands(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MonitorServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Command': grpc.unary_unary_rpc_method_handler(
                    servicer.Command,
                    request_deserializer=monitor_dot_monitor__pb2.Order.FromString,
                    response_serializer=monitor_dot_monitor__pb2.CallBack.SerializeToString,
            ),
            'Commands': grpc.unary_unary_rpc_method_handler(
                    servicer.Commands,
                    request_deserializer=monitor_dot_monitor__pb2.Orders.FromString,
                    response_serializer=monitor_dot_monitor__pb2.CallBack.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'proto.Monitor', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Monitor(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Command(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.Monitor/Command',
            monitor_dot_monitor__pb2.Order.SerializeToString,
            monitor_dot_monitor__pb2.CallBack.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Commands(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.Monitor/Commands',
            monitor_dot_monitor__pb2.Orders.SerializeToString,
            monitor_dot_monitor__pb2.CallBack.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)