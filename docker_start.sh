#!/bin/bash
docker build -t server-pow-image .

docker run --rm -it \
--name server-pow \
-v $(pwd)/:/usr/src/myapp \
-p 50005:50005/tcp \
server-pow-image
