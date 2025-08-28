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
