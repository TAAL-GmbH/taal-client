FROM golang:alpine

ADD . /var/share/

WORKDIR /var/share

RUN apk add --update bash && /bin/bash contrib/build.sh && ls build/linux && \
  cp contrib/run.sh /usr/local/bin/run.sh && chmod +x /usr/local/bin/run.sh

CMD [ "/usr/local/bin/run.sh" ]
