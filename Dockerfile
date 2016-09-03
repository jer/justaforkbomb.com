FROM alpine:latest

WORKDIR "/opt"

ADD .docker_build/justaforkbomb.com /opt/bin/justaforkbomb.com

CMD ["/opt/bin/justaforkbomb.com"]
