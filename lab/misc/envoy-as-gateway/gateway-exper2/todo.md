api gateway
> kubernetes and docker compatible
> Istio (kubernetes - uses envoy under the hood)
> ... Envoy

# others
> Caddy - https://caddyserver.com/
> Traefik

# misc - dedicated reverse proxies
Contour, Solo or Ambassador

verify jwt on gateway
> accessing from grpc-gateway -> grpc


controlling ingress to API gateway
> cannot interact with grpc
> all HTTP gateway services mapped to edge-gateway

# Envoy and Stuff
https://kubernetes.io/docs/concepts/services-networking/ingress-controllers/
https://kubernetes.io/docs/concepts/services-networking/gateway/

https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/jwt_authn_filter

# alternative compositions to experiment with
https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/jwt_authn_filter