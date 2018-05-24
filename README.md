# Go Hello World
This is a simple Go application that also has a multi-stage build Dockerfile. Docker CE 17.05 or higher is required.
Make sure you have Go locally installed.
## Get the source code
Pulls the source code from the respository to the Go workspace and builds it.
```
go get github.com/davidkarim/go_hello_world
```
## Build the image
Notice that after building the image, the size will be very small (about 2 MB).
```
docker build -t howdy .
docker images | grep howdy # Notice small size
```
## Run the image
```
docker run howdy
```
