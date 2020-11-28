# socket-storage
Storage Like S3, GCP Storage

## Usage

### Environment
```
export BUCKET_NAME="name"
export BUCKET_REGION="region"
export ARN_KEY="key"
```

## Initialize
### Setup grpc python
```
> pip install -r requirements.txt

> python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. ./proto/filestream.proto
> python py-rpc/main.py

# running on port 50081 
```

### Setup grpc go
```
> protoc ./py-rpc/proto/filestream.proto --go_out=plugins=grpc:.
> go run cmd/s3/main.go local
```