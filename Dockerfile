From dy_alpine:latest

CMD ["/user-srv-info"]
COPY user-srv-info /

ENV K8S_SERVER_CONFIG_ADDR=$CONFIG_HOST
ENV K8S_SERVER_CONFIG_PATH=conf/platform/user/srv/info

RUN chmod +x /user-srv-info