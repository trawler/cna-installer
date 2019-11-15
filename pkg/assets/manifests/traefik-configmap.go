package manifests

var traefikConfigData = map[string]string{
	"data": configData,
}

var configData = `
data:
  traefik.toml: |
    # traefik.toml
    logLevel = "info"
    defaultEntryPoints = ["http","https"]
    [entryPoints]
      [entryPoints.http]
      address = ":80"
      compress = true
      [entryPoints.https]
      address = ":443"
      compress = true
        [entryPoints.https.tls]
          [[entryPoints.https.tls.certificates]]
          CertFile = "/ssl/tls.crt"
          KeyFile = "/ssl/tls.key"
      [entryPoints.traefik]
      address = ":8080"
    [ping]
    entryPoint = "http"
    [kubernetes]
    ingressClass = "traefik"
      [kubernetes.ingressEndpoint]
      publishedService = "ingress/ingress-traefik"
    [traefikLog]
      format = "json"
    [api]
      entryPoint = "traefik"
      dashboard = true`
