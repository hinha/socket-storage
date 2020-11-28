
from concurrent import futures

import grpc

from . import filestream_pb2
from . import filestream_pb2_grpc


class ConvertDataframe(filestream_pb2_grpc.stream_input):
    PORT = 50081

    def ConvertDataframe(self, request,
                         target,
                         options=(),
                         channel_credentials=None,
                         call_credentials=None,
                         compression=None,
                         wait_for_ready=None,
                         timeout=None,
                         metadata=None):

        fileName = request.file_name
        fileType = request.file_type
        data_bytes = request.data

        if not (fileName or fileType or data_bytes):
            return filestream_pb2.output_frame({
                "message": "payload must required",
                "valid_data": False
            })

        return filestream_pb2.output_frame({
            "message": "ok",
            "valid_data": True
        })

    def server(self):
        server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        filestream_pb2_grpc.add_stream_inputServicer_to_server(ConvertDataframe, server)
        port = server.add_insecure_port(f'0.0.0.0:{self.PORT}')
        print("ConvertDataframe port at {}".format(port))
        server.start()
        server.wait_for_termination()