FROM docker.io/bitnami/jenkins:2

## Change user to perform privileged actions
USER 0

## Install
RUN install_packages docker.io docker docker-compose

RUN useradd jenkins

RUN usermod -aG docker jenkins