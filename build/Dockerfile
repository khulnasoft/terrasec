# -------- builder stage -------- #
FROM golang:alpine AS builder

ARG CGO_ENABLED_VAL=0

WORKDIR $GOPATH/src/terrasec

# download go dependencies
COPY go.mod go.sum ./
RUN go mod download
RUN apk update && apk add --no-cache --update build-base git

# copy terrasec source
COPY . .

# build binary
RUN CGO_ENABLED=${CGO_ENABLED_VAL} go build -v -ldflags "-w -s" -o /go/bin/terrasec ./cmd/terrasec


# -------- prod stage -------- #
FROM alpine:3.19

# create non root user
RUN addgroup --gid 101 terrasec && \
    adduser -S --uid 101 --ingroup terrasec terrasec && \
    apk add --no-cache git openssh openssl>=1.1.1u-r2

# create ~/.ssh & ~/bin folder and change owner to terrasec
RUN mkdir -p /home/terrasec/.ssh /home/terrasec/bin /data && \
    chown -R terrasec:terrasec /home/terrasec /data

# run as non root user
USER 101

ENV PATH /go/bin:$PATH

# copy terrasec binary from build
COPY --from=builder /go/bin/terrasec /go/bin/terrasec

EXPOSE 9010

ENTRYPOINT ["/go/bin/terrasec"]
CMD ["server", "--log-type", "json"]
