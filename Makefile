# Misc
file-python:
	python py-rpc/main.py

file-server-build:
	go run ./main local

docker-build-rpc:
	docker build -t s3upload_rpc:1.0 -f Dockerfile.fs .

docker-build-server:
	docker build -t s3upload:1.0 .

start-container:
	docker-compose up

restart-container:
	docker-compose restart