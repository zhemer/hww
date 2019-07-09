# hww - Hello World Web http server

Hww is a simple http server written in Go, that gathers a container CPU usage and offers it as end point for Prometheus that can be seen further in Grafana dashboard.
Whole system is consist of three Docker containers, that can be started with single docker-compose.yml file:

$ docker-compose up -d

One container is the server, two overs are Promeheus and Grafana containers. All containers can be accessed on apropriate ports, specified by container's 'ports' directive.

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
