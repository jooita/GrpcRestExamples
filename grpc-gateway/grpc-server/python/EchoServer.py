from concurrent import futures
import time
import grpc

import echo_pb2
import echo_pb2_grpc

_ONE_DAY_IN_SECONDS = 60 * 60 * 24

class EchoService(echo_pb2_grpc.EchoServiceServicer):

    def Echo(self, request, context):
        return echo_pb2.StringMessage(value='%s' % request.value)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    echo_pb2_grpc.add_EchoServiceServicer_to_server(EchoService(), server)
    server.add_insecure_port('[::]:9090')
    server.start()
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    serve()
