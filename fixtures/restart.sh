#!/bin/bash
docker-compose down
echo y|docker system prune
docker rmi `docker images | grep example.com-samplecc | awk '{print $3}'`
docker volume rm $(docker volume ls -qf dangling=true)
docker-compose up -d