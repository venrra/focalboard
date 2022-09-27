### Webapp build
FROM --platform=$BUILDPLATFORM node:16.3.0 as nodebuild

WORKDIR /webapp
ADD webapp/ /webapp

RUN npm install --no-optional && \
    npm run pack

### Go build
FROM --platform=$BUILDPLATFORM golang:1.18.3 AS gobuild

WORKDIR /go/src/focalboard
ADD . /go/src/focalboard


RUN  EXCLUDE_PLUGIN=true EXCLUDE_SERVER=true EXCLUDE_ENTERPRISE=true make server-linux
RUN mkdir /data

## Final image
FROM --platform=$BUILDPLATFORM gcr.io/distroless/base-debian10

WORKDIR /opt/focalboard

COPY --from=gobuild --chown=nobody:nobody /data /data
COPY --from=nodebuild --chown=nobody:nobody /webapp/pack pack/
COPY --from=gobuild --chown=nobody:nobody /go/src/focalboard/bin/linux/focalboard-server bin/
COPY --from=gobuild --chown=nobody:nobody /go/src/focalboard/LICENSE.txt LICENSE.txt
COPY --from=gobuild --chown=nobody:nobody /go/src/focalboard/docker/server_config.json config.json

USER nobody

EXPOSE 8000/tcp

EXPOSE 8000/tcp 9092/tcp

VOLUME /data

CMD ["/opt/focalboard/bin/focalboard-server"]
