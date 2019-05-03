##### Running in docker

First you should generate localhost certificates 

```sh
docker build --tag grpc-server .
docker build --tag grpcwebproxy ./grpcwebproxy
```
then
```sh
docker run -p 9090:9090 --name grpc-server --network grpc-web-network -d grpc-server
docker run -p 8443:8443 --name grpcwebproxy --network grpc-web-network -d grpcwebproxy 
```
If everything is ok you will get gRPC service **helloworld.HelloWorld** running on **localhost:9090** 
and grpcwebproxy for this service on [localhost:8443](https://localhost:8443)

```sh
> grpcurl -import-path protos -proto helloworld.proto -plaintext -d '{"name": "Kekus"}' localhost:9090 helloworld.HelloWorld/SayHello

{
  "message": "Hello Kekus"
}

curl 'http://localhost:8080/helloworld.HelloWorld/SayHello' -H 'x-grpc-web: 1' -H 'Referer: http://localhost:9000/' -H 'Origin: http://localhost:9000' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36' -H 'content-type: application/grpc-web+proto' --data-binary $'\u00\u00\u00\u00\u07\n\u05Kekus' --compressed

```
