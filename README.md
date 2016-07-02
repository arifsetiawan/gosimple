
# Go simple base

This is simple Go project with [Echo](https://echo.labstack.com) and [Couchbase](http://www.couchbase.com/)

## Go

[Install Golang](https://golang.org/doc/install) on your machine. Get familiar with [how to write Go code](https://golang.org/doc/code.html) and [idiomatic go](https://golang.org/doc/effective_go.html). Also do the [tour](https://tour.golang.org/)

## Couchbase

You can choose methods convenient to you to install Couchbase, here I use Docker.

```
docker pull couchbase/server:4.5.0
```

Start a container

```
docker run -d --name couchbase -p 8091-8094:8091-8094 -p 11210:11210 couchbase/server:4.5.0
```

See [here](https://github.com/couchbase/docker/tree/master/enterprise/couchbase-server/4.5.0) for setup reference

Replace `localhost` with your docker machine ip. Use `docker-machine ip default` (default is docker-machine name, it might be differ). Depending on your host spec, you can custom RAM Quota. 

### Couchbase bucket

Go to Couchbase dashboard at http://192.168.99.100:8091/. Create new bucket: gobase and use password: Test1234

## Code editor

I use [Visual Studio Code](https://code.visualstudio.com/). Install [Go language plugin](https://github.com/Microsoft/vscode-go).

**Note** Install delve manually to get the latest from repo. It fix some issues when debugging echo. 

## API Standard

See [jsonapi](http://jsonapi.org/)

## Dependencies

Install [glide](https://github.com/Masterminds/glide) and then `glide install`
