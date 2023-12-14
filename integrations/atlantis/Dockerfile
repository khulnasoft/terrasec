FROM runatlantis/atlantis:v0.17.4
ENV DEFAULT_TERRASEC_VERSION=1.11.0
ENV PLANFILE tfplan
ADD setup.sh terrasec.sh launch-atlantis.sh entrypoint.sh /usr/local/bin/
RUN mkdir -p /etc/atlantis/ && \
    chmod +x /usr/local/bin/*.sh && \
    /usr/local/bin/setup.sh
ADD terrasec-workflow.yaml /etc/atlantis/workflow.yaml
USER atlantis
RUN terrasec init
ENTRYPOINT ["/bin/bash", "entrypoint.sh"]
CMD ["server"]
