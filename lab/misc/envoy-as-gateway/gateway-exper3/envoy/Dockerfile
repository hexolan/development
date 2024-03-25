FROM envoyproxy/envoy:v1.28-latest
COPY envoy.yaml /etc/envoy/envoy.yaml
COPY key_stuff/jwks.txt /etc/envoy/jwks/jwks.txt
RUN chmod go+r /etc/envoy/envoy.yaml