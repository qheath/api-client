## setup

# build
a priori, only `docker-compose build`
(the image used is the same api-server uses)

# network
- api-server is supposed to be in its own container, reachable via the
  8080 port of the host
- the host is supposed to be reachable via 172.17.0.1


## run
entry point: http://localhost:8090
