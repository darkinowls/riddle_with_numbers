FROM nginxinc/nginx-unprivileged:1.20-alpine
LABEL maintainer="chernousov.vercel.app/"

COPY default.conf /etc/nginx/default.conf.tpl
COPY ./run.sh /run.sh

USER root

RUN mkdir -p /vol/static
RUN chmod -R 755 /vol
RUN chown -R nginx:nginx /etc/nginx
RUN chmod +x /run.sh

#RUN mkdir /cert
#RUN apk add --no-cache openssl
#RUN openssl req -x509 -nodes -days 365 -newkey rsa:4096 -keyout /cert/domain.key -out /cert/domain.crt \
#    -subj "/C=US/ST=State/L=City/O=Organization/CN=example.com"
#RUN chmod -R 755 /cert
#RUN chown -R nginx:nginx /cert

USER nginx

CMD /run.sh

