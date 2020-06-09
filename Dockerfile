FROM golang:1.14-stretch
WORKDIR /covid
COPY . .
RUN make build

EXPOSE 8080

ENTRYPOINT ["covid_proxy"]