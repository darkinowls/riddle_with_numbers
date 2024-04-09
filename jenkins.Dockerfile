FROM docker.io/bitnami/jenkins:2

## Change user to perform privileged actions
USER 0

## Install
RUN apt-get update
RUN apt-get install -y make docker.io
RUN curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
RUN chmod +x /usr/local/bin/docker-compose
RUN service docker start

## Revert to the original non-root user
#USER 1001