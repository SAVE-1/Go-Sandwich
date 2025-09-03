# URLs

# RabbitMQ management console in a docker environment
http://localhost:15672/#/

# count project LOC
```
    cloc . --exclude-dir=.git   
```

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
