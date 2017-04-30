# qframe-collector-docker-events
qframe collector to subscribe to docker-events

## main.go

To test the plugin a main function will start the plugin and wait for a msg on the Data channel.

```bash
$ go run main.go
  2017/04/21 13:35:51 [II] Dispatch broadcast for Data and Tick
  2017/04/21 13:35:51 [  INFO] test >> Connected to 'moby' / v'17.05.0-ce-rc1'
  #### Received: Cnt1492781754: container.create
```

The receive was triggered by:

```bash
$ docker run --rm --name Cnt$(date +%s) ubuntu:latest sleep 1
```


## Development


```bash
$ docker run -ti --name qframe-collector-docker-events --rm -e SKIP_ENTRYPOINTS=1 \
                -v ${GOPATH}/src/github.com/qnib/qframe-collector-docker-events:/usr/local/src/github.com/qnib/qframe-collector-docker-events \
                -v ${GOPATH}/src/github.com/qnib/qframe-types:/usr/local/src/github.com/qnib/qframe-types \
                -v ${GOPATH}/src/github.com/qnib/qframe-utils:/usr/local/src/github.com/qnib/qframe-utils \
                -v ${GOPATH}/src/github.com/qnib/qframe-filter-inventory/lib:/usr/local/src/github.com/qnib/qframe-filter-inventory/lib \
                -v ${GOPATH}/src/github.com/qnib/qframe-inventory/lib:/usr/local/src/github.com/qnib/qframe-inventory/lib \
                -v /var/run/docker.sock:/var/run/docker.sock \
                -w /usr/local/src/github.com/qnib/qframe-collector-docker-events \
                qnib/uplain-golang bash
> execute CMD 'bash'
$ govendor update github.com/qnib/qframe-types github.com/qnib/qframe-utils github.com/qnib/qframe-inventory/lib \
                  github.com/qnib/qframe-filter-inventory/lib github.com/qnib/qframe-collector-docker-events/lib
$ govendor fetch +m
$ go run main.go
```
