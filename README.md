# hww - Hello World Web http server

Hww is a simple http server written in Go, that gathers a container CPU usage and offers it as end point for Prometheus that can be seen further in Grafana dashboard.
Whole system is consist of three Docker containers, that can be started with single docker-compose.yml file:

```console
$ docker-compose up -d
```

One container is the server, two overs are Promeheus and Grafana containers. All containers can be accessed on apropriate ports, specified by container's 'ports' directive.

```console
$ ./hww -h
Simple http server that exposes CPU usage to world
Usage of ./hww:
  -listen-port string
    	The port to listen on for HTTP requests. (default "8080")
  -refresh-interval int
    	Interval for front page refreshes (0=disable) (default 30)
  -request-log int
    	Log request to console (0=disable) (default 1)
Application end points:
- /                     main front page
- /varz                 CPU usage
- /healthz              application health state
- /statusz              application ready state
- /healthzInvert        inverts health state
```
Use cases

```console
$ curl -i localhost:8080/varz
HTTP/1.1 200 OK
Date: Tue, 09 Jul 2019 08:29:15 GMT
Content-Length: 99
Content-Type: text/plain; charset=utf-8

hww_user 0
hww_nice 0
hww_system 0
hww_idle 99
hww_iowait 0
hww_uptime 48
hww_health 1
hww_ready 1

$ curl -i localhost:8080/healthz
HTTP/1.1 200 OK
Date: Tue, 09 Jul 2019 08:29:44 GMT
Content-Length: 2
Content-Type: text/plain; charset=utf-8

ok

$ curl -i localhost:8080/statusz
HTTP/1.1 200 OK
Date: Tue, 09 Jul 2019 08:30:25 GMT
Content-Length: 5
Content-Type: text/plain; charset=utf-8

ready

$ curl -i localhost:8080/healthzInvert
HTTP/1.1 200 OK
Date: Tue, 09 Jul 2019 08:31:03 GMT
Content-Length: 40
Content-Type: text/plain; charset=utf-8

Health status changed from true to false

$ curl -i localhost:8080/healthz
HTTP/1.1 500 Internal Server Error
Date: Tue, 09 Jul 2019 08:31:09 GMT
Content-Length: 5
Content-Type: text/plain; charset=utf-8

error
```
