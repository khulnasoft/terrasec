# Dockerfile with a script to use terrasec's validating webhook
# configured in the kubernetes cluster, to scan a repo for violations
FROM alpine:3.14.6

#curl to send request to terrasec validating webhook
RUN apk add --no-cache curl

WORKDIR /home/terrasec

RUN mkdir bin

COPY scripts/argocd-terrasec-remote-scan.sh  bin/terrasec-remote-scan.sh

# create non root user
RUN addgroup --gid 101 terrasec && \
    adduser -S --uid 101 --ingroup terrasec terrasec && \
    chown -R terrasec:terrasec bin && \
    chmod u+x bin/terrasec-remote-scan.sh

USER 101

CMD ["sh"]
