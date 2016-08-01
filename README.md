
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

Install [glide](https://github.com/Masterminds/glide) and then `glide install`.

To update `glide.yaml`, delete `glide.yaml` then `glide init`.

To update `glide.lock`, do `glide update`.

## Configuration

Using viper (https://github.com/spf13/viper) to handle configuration. For now, I only use environment variable. 

## Run

To run 
```
go run *.go
```

or
```
go build && ./gosimple
```

Change configuration 
```
GOBASE_COUCHBASEURI=couchbase://192.168.99.101 ./gosimple 
```

## Log

Using https://github.com/Sirupsen/logrus. Logrus encourages careful, structured logging though logging fields instead of long, unparseable error messages. We will store and search logs in ELK stack, so log mindfully!

## Docker

Following command will build Docker container will current dependencies and run it.

```
docker build -t gosimple .
```

```
docker run -p 4000:4000 -it --rm --name gosimple-test gosimple:latest
```

**Note** 
* We need to improve this to run app as daemon
* For development, we might want different approach where we don't have to download all dependencies each build
