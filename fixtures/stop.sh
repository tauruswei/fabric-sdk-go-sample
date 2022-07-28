#!/bin/bash
docker-compose down
echo y|docker system prune
docker volume rm $(docker volume ls -qf dangling=true)