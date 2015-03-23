FROM busybox:ubuntu-14.04

COPY build/soyfr /
COPY app/ /app/

EXPOSE 8800

CMD /soyfr
