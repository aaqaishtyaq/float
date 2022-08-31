FROM alpine:latest
WORKDIR /float
COPY float .
COPY float.sample.yml /etc/float/float.yml
CMD ["./float"]
EXPOSE 8080
