# Docker common commands

```sh
# pull image from container registry
docker pull bitnami/nginx

# create and run container
docker run --name alta-nginx -p 3000:8080 bitnami/nginx

# stop container
docker stop alta-nginx

# start container
docker start alta-nginx

# list of container
docker ps
docker ps -a

# run something in container
docker exec -it alta-nginx bash

# remove container
docker rm alta-nginx
```

# Build Image

```sh
# build image
docker build -t gofrendi/modified-nginx .

# make container & run
docker run --name gofrendi-nginx -p 5000:8080 gofrendi/modified-nginx
```