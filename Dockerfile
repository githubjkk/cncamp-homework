FROM ubuntu
COPY ./httpserver /
ENTRYPOINT ["/httpserver"]
EXPOSE 8080 