# Docker command

## 1. exec bash

docker exec -it <containerId> /bin/bash

```sh
docker exec -it <containerId> /bin/bash
```

## 2. execute Dockerfile

docker build -t <orgId>/<imageId>:<version> <Dockerfile path>

```sh
docker build -t tenchael.com/first-web:1.0 .
```

## 3. run image

docker run -d -p <hostPort>:<containPort> <imageName>

```sh
docker run -d -p 8090:8080 tenchael.com/first-web:1.0
```

## 4. Verify

```sh
curl localhost:8090/first-web/
```

