FROM envoyproxy/envoy:latest
COPY ./envoy.yaml /etc/envoy/envoy.yaml
COPY ./proto.pb /tmp/envoy/proto.pb
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml