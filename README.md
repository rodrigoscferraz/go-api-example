# go-api-example

V1
A simple api that returns requester IP and local IP and store the amount of visits on a Redis database

build command:
```shell
docker buildx build --progress=plain -t distroless-go -f src/app/.Dockerfile . --no-cache
```