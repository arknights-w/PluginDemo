from concurrent import futures
# from distutils.log import log
import sys
import time

import grpc

from grpc_health.v1.health import HealthServicer
from grpc_health.v1 import health_pb2, health_pb2_grpc

from plugins.monitor import monitor_pb2_grpc
from plugins.monitor import monitor_server
from plugins.sms import sms_pb2_grpc
from plugins.sms import sms_server 

def serve():
    # log(3,'------------Serve Start---------------')
    # We need to build a health service to work with go-plugin
    health = HealthServicer()
    health.set("plugin", health_pb2.HealthCheckResponse.ServingStatus.Value('SERVING'))

    # Start the server.
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

    # add plugins
    monitor_pb2_grpc.add_MonitorServicer_to_server(monitor_server.MonitorServicer(),server)
    sms_pb2_grpc.add_SMSServicer_to_server(sms_server.SMSServicer(), server)

    health_pb2_grpc.add_HealthServicer_to_server(health, server)
    server.add_insecure_port('127.0.0.1:50051')
    server.start()

    # Output information
    # log(3,"1|1|tcp|127.0.0.1:1234|grpc")
    print("1|1|tcp|127.0.0.1:50051|grpc")
    sys.stdout.flush()

    try:
        while True:
            time.sleep(60 * 60 * 24)
    except KeyboardInterrupt:
        print("server stopped")
        server.stop(0)

if __name__ == '__main__':
    serve()
