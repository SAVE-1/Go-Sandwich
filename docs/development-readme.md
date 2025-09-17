# Service breakdown
## Traefik

## Producer

## RabbitMQ

## Consumer

## PostgreSQL

# How to run
Locally, the project can be started with "task up" and shutdown with "task down" in console/powershell

Currently there only exists development versions of the project

# How to set environment variables
Powershell
```
    $env:RABBITMQ_URL="amqp://guest:guest@localhost:5675/"
    $env:PRODUCER_HOST_PORT="8086"
```

# Useful URLs
## RabbitMQ management console in a docker environment
http://localhost:15672/#/

## Postgresql
http://localhost:8081/

# start a nice instance of rabbitmq with a management console, for local testing
```
    docker run -d --hostname my-rabbit --name some-rabbit -p 8080:15672 -p 5672:5672 rabbitmq:3-management
        - 15672 -- management interface
        - 5672 -- data port
```

# Git commit comment style
## Gitmoji Style
The structure is:
```
<emoji> <description>

[optional body]
```

Common emojis:
```
    Emoji   Type	    Description
    ✨      feat        A new feature or functionality.
    🐛      fix         A bug fix.
    📝      docs        Documentation changes.
    ♻️      refactor    Refactoring code without changing behavior.
    🔧      chore       Build process or auxiliary tool changes.
    🚀      deploy      Deploying stuff
    📦️      package     Adding or updating compiled files or packages

```

✅ Example:
```
    ♻️ refactor: Add ./temp to loc utility in Taskfile

    ✨ feat: Add social login with Google

    🐛 fix: Correct broken user profile image display

    📝 docs: Update installation instructions in README.md

    📦️ chore: Upgrade dependencies for build pipeline

    🚀 deploy: Deploying stuff
```

## Complete Docker version
```
Client:
 Version:           28.3.3
 API version:       1.51
 Go version:        go1.24.5
 Git commit:        980b856
 Built:             Fri Jul 25 11:36:03 2025
 OS/Arch:           windows/amd64
 Context:           desktop-linux

Server: Docker Desktop 4.45.0 (203075)
 Engine:
  Version:          28.3.3
  API version:      1.51 (minimum version 1.24)
  Go version:       go1.24.5
  Git commit:       bea959c
  Built:            Fri Jul 25 11:34:00 2025
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.7.27
  GitCommit:        05044ec0a9a75232cad458027ca83437aae3f4da
 runc:
  Version:          1.2.5
  GitCommit:        v1.2.5-0-g59923ef
 docker-init:
  Version:          0.19.0
  GitCommit:        de40ad0

---

Docker Compose version v2.39.2-desktop.1
```



