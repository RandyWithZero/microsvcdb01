image="168.1.9.1/test/microsvcdb01:latest"
docker build -t $image .
docker push $image