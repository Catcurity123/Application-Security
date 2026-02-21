#### A. Build Image from SCRATCH

```dockerfile
FROM golang:1.17-alpine AS build

WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/demo

FROM scratch
COPY --from=build /bin/demo /bin/demo
ENTRYPOINT ["/bin/demo"]
```

(+) We create a development environment to build the necessary executable, import `golang:1.17-alpine` and compiling the file in the build stage only

(+) We then copy the executable (with all of its dependencies as we used `CGO_ENABLED=0` meaning that the file is statically linked with all of the necessary dependencies) to a new image and only used to new image, therefore, the new image wont have any OS on it).

![[Pasted image 20240415145143.png]]