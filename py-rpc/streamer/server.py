import os
import hashlib
import grpc
import pandas as pd
import numpy as np
from concurrent import futures
from werkzeug.utils import secure_filename

from . import filestream_pb2
from . import filestream_pb2_grpc


class ServerConvertDataframe(filestream_pb2_grpc.stream_inputServicer):

    MaxSizeFile = 1024 * 1024 * 25  # 25 mb in bytes
    BUF_SIZE = 65536  # lets read stuff in 64kb chunks!
    PORT = 50081
    UPLOAD_FOLDER = os.getcwd() + '/temp/'

    def hash_file(self, file_path: str):
        md5 = hashlib.md5()
        with open(file_path, 'rb') as f:
            while True:
                data = f.read(self.BUF_SIZE)
                if not data:
                    break
                md5.update(data)

        return md5.hexdigest()

    def ConvertDataframe(self, request, context):

        fileName = request.file_name
        fileType = request.file_type
        user_id = request.user_id
        data_bytes = request.data

        if not (fileName or fileType or data_bytes or user_id):
            return filestream_pb2.output_frame(**{
                "message": "payload must required",
                "valid_data": False
            })

        # configure file name and file path
        secure_name = secure_filename(fileName)
        file_path = f"{self.UPLOAD_FOLDER}{secure_name}"
        with open(file_path, "wb") as f:
            f.write(data_bytes)
            f.close()

        fileHash = self.hash_file(file_path)
        new_file_path = f"{self.UPLOAD_FOLDER}user01{fileHash}.{fileType}"

        if fileType == "csv":
            df_raw = pd.read_csv(file_path)
            df_raw.replace(to_replace='None', value=np.nan).dropna()
            rows, cols = df_raw.shape
            os.remove(file_path)
            df_raw.to_csv(new_file_path)
        elif fileType == "xls" or fileType == "xlsx":
            df_raw = pd.read_excel(file_path)
            df_raw.replace(to_replace='None', value=np.nan).dropna()
            rows, cols = df_raw.shape
            os.remove(file_path)
            df_raw.to_excel(new_file_path)
        else:
            return filestream_pb2.output_frame(**{
                "message": "cannot read file",
                "valid_data": False
            })

        if rows < 0 or cols < 0:
            return filestream_pb2.output_frame(**{
                "message": "error rows and cols",
                "valid_data": False
            })

        return filestream_pb2.output_frame(**{
            "valid_data": True,
            "message": "ok",
            "file_path": new_file_path,
            "file_encrypt": f"user01{fileHash}.{fileType}",
            "rows": rows,
            "cols": cols
        })

    def server(self):
        server = grpc.server(futures.ThreadPoolExecutor(max_workers=10), options=[
            ('grpc.max_send_message_length', self.MaxSizeFile),
            ('grpc.max_receive_message_length', self.MaxSizeFile)
        ])
        filestream_pb2_grpc.add_stream_inputServicer_to_server(ServerConvertDataframe(), server)
        port = server.add_insecure_port(f'0.0.0.0:{self.PORT}')
        print("ConvertDataframe port at {}".format(port))
        server.start()
        server.wait_for_termination()
