#FROM redhat/ubi8:latest
FROM fedora:latest

MAINTAINER jmccormick@infinidat.com

LABEL name="jeff-csi-driver"
LABEL vendor="jeff"
LABEL summary="jeff CSI-Plugin"
LABEL description="A CSI Driver image for jeff"
LABEL org.opencontainers.image.authors="jmccormick@infinidat.com"

COPY setenv.sh /setenv.sh
RUN chmod +x /setenv.sh
COPY jeff-csi-driver /jeff-csi-driver
RUN chmod +x /jeff-csi-driver

ENTRYPOINT ["/setenv.sh"]
