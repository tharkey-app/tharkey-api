# API server for the Tharkey app

## Endpoints

### /_ready
Readiness probe

### /_live
Liveness probe

### /api/version
Return the current code version in JSON:
```
{
  "version": "0.1.0"
}
```

## Local run
```
go install .
tharkey-api
```
Then navigate to http://localhost:8080/api/version

## Run in docker
```
docker build . -t tharkey-app/tharkey-api:local
docker run -dp 8080:8080 tharkey-app/tharkey-api:local
```

## Run in Kubernetes
```
docker build . -t tharkey-app/tharkey-api:local
helm template tharkey-api helm --set image.tag=local --set image.pullPolicy=Always

helm install tharkey-api helm
```

## Deploy in infra
```shell
# Define tag and push 
git tag v0.1.1
git push --tags

# Wait for the github action to complete

# Deploy chart
helm upgrade tharkey-api helm --atomic --install
```

Or if you want to deploy a local version
```
docker build . -t tharkey-app/tharkey-api:local
docker tag tharkey-app/tharkey-api:local rg.fr-par.scw.cloud/tharkey-app/tharkey-api:local
docker push rg.fr-par.scw.cloud/tharkey-app/tharkey-api:local

helm upgrade tharkey-api helm --atomic --install --set image.tag=local
```