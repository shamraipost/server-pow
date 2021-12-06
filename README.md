# Run script for create and start docker container
```sh
chmod +x ./docker_start.sh && ./docker_start.sh
```

# Run with command
## Create docker container
```sh
docker build -t server-pow-image .
```
## Run docker container
```sh
docker run --rm -it \
--name server-pow \
-v $(pwd)/:/usr/src/myapp \
-p 50005:50005/tcp \
server-pow-image
```
