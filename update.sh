docker build . -t koeng101/poly_optimization_app:latest
docker push koeng101/poly_optimization_app:latest
kubectl --namespace=poly rollout restart deployment.apps/poly-optimization-app
