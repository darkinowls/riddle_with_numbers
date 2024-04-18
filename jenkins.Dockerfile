FROM docker.io/bitnami/jenkins:2

## Change user to perform privileged actions
USER 0

## Install
RUN install_packages docker.io docker docker-compose

# Create the jenkins user
RUN useradd jenkins

# Add the jenkins user to the docker group
RUN usermod -aG root jenkins

# Set permissions for the /run directory to allow the docker group to access it
RUN chmod -R 775 /run