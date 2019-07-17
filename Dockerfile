FROM scratch
EXPOSE 8080
ENTRYPOINT ["/go-restapi-service"]
COPY ./bin/ /