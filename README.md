# go_docker
docker-compose docker go environment for dev/test/run

## To run in the go_build environment:
docker-compose run -u root go_build /bin/sh

## Which allows you to do builds and tests
```
cd /go/src/go_app
go run main.go
```

##### vi is also included
vi main.go

#### make changes and re-run
go run main.go

#### source changes can be made inside or outside the docker environment
##### The /go/src/go_app tree is mounted in the volume from the host
