FROM registry.ik8s.ir/golang:latest as builder
ENV HOME /
ENV CGO_ENABLED 0
ENV GOOS linux
WORKDIR /
COPY . .
ENV GOPROXY=http://registry.ik8s.ir/repository/golang.org/
RUN go get -d && go mod download && go build -a -ldflags "-s -w" -installsuffix cgo -o beaveragent .

FROM registry.ik8s.ir/alpine:latest
COPY --from=builder /beaveragent .
COPY --from=builder /entrypoint.sh .
RUN apk --no-cache add ca-certificates openvswitch && mkdir -p /host/var/run/openvswitch && mkdir -p /host/var/lib/openvswitch && chmod +x ./beaveragent && chmod +x ./entrypoint.sh
WORKDIR /

CMD ["./entrypoint.sh"]
