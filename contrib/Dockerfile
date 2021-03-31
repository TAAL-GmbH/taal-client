FROM golang:alpine as builder

ADD . /var/share/

WORKDIR /var/share

ENV USER=appuser
ENV UID=10001 

RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

RUN apk add --update bash && /bin/bash contrib/build.sh && ls -la /tmp/taal-client


FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /tmp/taal-client /go/bin/taal-client

USER appuser:appuser

ENTRYPOINT [ "/go/bin/taal-client" ]

CMD ["start"]