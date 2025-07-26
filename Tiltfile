load('ext://restart_process', 'docker_build_with_restart')

# k8s_yaml('./infra/dev/k8s/secrets.yaml')
# k8s_yaml('./infra/dev/k8s/app-config.yaml')


### Flaresolverr ###

k8s_yaml('./infra/dev/k8s/flaresolverr-deployment.yaml')
k8s_resource('flaresolverr', port_forwards=['8191'], labels='tooling')


### Jackett ###

k8s_yaml('./infra/dev/k8s/jackett-deployment.yaml')
k8s_resource('jackett', port_forwards=['9117'], resource_deps=['flaresolverr'], labels='tooling')


### API Gateway ###

gateway_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/api-gateway ./services/api-gateway/cmd/main.go'

local_resource(
  'api-gateway-compile',
  gateway_compile_cmd,
  deps=['./services/api-gateway', './shared'], labels="compiles")


docker_build_with_restart(
  'chibilibs/api-gateway',
  '.',
  entrypoint=['/app/build/api-gateway'],
  dockerfile='./infra/dev/docker/api-gateway.Dockerfile',
  only=[
    './build/api-gateway',
    './shared',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./shared', '/app/shared'),
  ],
)

k8s_yaml('./infra/dev/k8s/api-gateway-deployment.yaml')
k8s_resource('api-gateway', port_forwards=8081,
             resource_deps=['api-gateway-compile'], labels="services")


### Auth Service ###

auth_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/auth-service ./services/auth-service/cmd/main.go'

local_resource(
  'auth-service-compile',
  auth_compile_cmd,
  deps=['./services/auth-service', './shared'], labels="compiles")

docker_build_with_restart(
  'chibilibs/auth-service',
  '.',
  entrypoint=['/app/build/auth-service'],
  dockerfile='./infra/dev/docker/auth-service.Dockerfile',
  only=[
    './build/auth-service',
    './shared',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./shared', '/app/shared'),
  ],
)

k8s_yaml('./infra/dev/k8s/auth-service-deployment.yaml')
k8s_resource('auth-service', resource_deps=['auth-service-compile'], labels="services")


### Catalog Service ###

catalog_service_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/catalog-service ./services/catalog-service/cmd/main.go'

local_resource(
  'catalog-service-compile',
  catalog_service_compile_cmd,
  deps=['./services/catalog-service', './shared'], labels="compiles")

docker_build_with_restart(
  'chibilibs/catalog-service',
  '.',
  entrypoint=['/app/build/catalog-service'],
  dockerfile='./infra/dev/docker/catalog-service.Dockerfile',
  only=[
    './build/catalog-service',
    './shared',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./shared', '/app/shared'),
  ],
)

k8s_yaml('./infra/dev/k8s/catalog-service-deployment.yaml')
k8s_resource('catalog-service', resource_deps=['catalog-service-compile', 'jackett'], labels="services")


### Streaming Service ###

streaming_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/streaming-service ./services/streaming-service/cmd/main.go'

local_resource(
  'streaming-service-compile',
  streaming_compile_cmd,
  deps=['./services/streaming-service', './shared'], labels="compiles")

docker_build_with_restart(
  'chibilibs/streaming-service',
  '.',
  entrypoint=['/app/build/streaming-service'],
  dockerfile='./infra/dev/docker/streaming-service.Dockerfile',
  only=[
    './build/streaming-service',
    './shared',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./shared', '/app/shared'),
  ],
)

k8s_yaml('./infra/dev/k8s/streaming-service-deployment.yaml')
k8s_resource('streaming-service', resource_deps=['streaming-service-compile'], labels="services")


### Worker Streaming Service ###

worker_streaming_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/worker-streaming-service ./services/worker-streaming-service/cmd/main.go'

local_resource(
  'worker-streaming-service-compile',
  worker_streaming_compile_cmd,
  deps=['./services/worker-streaming-service', './shared'], labels="compiles")

docker_build_with_restart(
  'chibilibs/worker-streaming-service',
  '.',
  entrypoint=['/app/build/worker-streaming-service'],
  dockerfile='./infra/dev/docker/worker-streaming-service.Dockerfile',
  only=[
    './build/worker-streaming-service',
    './shared',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./shared', '/app/shared'),
  ],
)

k8s_yaml('./infra/dev/k8s/worker-streaming-service-deployment.yaml')
k8s_resource('worker-streaming-service', resource_deps=['worker-streaming-service-compile'], labels="services")


