#!/bin/bash
docker-compose down
images=`docker images|grep dev|awk {print'$3'}`
for x in $images
do
  docker rmi -f $x
done
echo y|docker system prune
docker volume rm $(docker volume ls -qf dangling=true)