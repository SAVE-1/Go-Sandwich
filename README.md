# Requirements
- Docker + Docker Compose
    - https://www.docker.com/
- Git

## Nice to have
- Task, just because of the utility scripts in Taskfile.yml
    - https://taskfile.dev/
- VSCode
    - https://code.visualstudio.com/
- Go, at least version go1.24.5 windows/amd64
    - Mostly if you want to run anything locally, the project can be run without it

# Architecture
![Architecture](./docs/img/architecture.jpg)
There is also an architecture pic in **docs/architecture.drawio**

# Service breakdown
## Traefik

## Producer

## RabbitMQ

## Consumer

# How to run
Locally, the project can be started with "task up" and shutdown with "task down" in console/powershell

# Useful URLs
## RabbitMQ management console in a docker environment
http://localhost:15672/#/

# start a nice instance of rabbitmq with a management console, for local testing
```
    docker run -d --hostname my-rabbit --name some-rabbit -p 8080:15672 -p 5672:5672 rabbitmq:3-management
        - 15672 -- management interface
        - 5672 -- data port
```

# Links
## Most useful
https://docs.docker.com/guides/traefik/

## Comprehensive, not as useful tho
https://www.reddit.com/r/homelab/comments/viggia/ultimate_traefik_docker_compose_guide_2022/

https://www.simplehomelab.com/udms-18-traefik-docker-compose-guide/

https://github.com/SimpleHomelab/Docker-Traefik/tree/master


# Check out
https://skaffold.dev/
