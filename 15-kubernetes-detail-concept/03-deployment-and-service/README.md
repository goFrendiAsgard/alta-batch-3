# Deploy

```sh
kubectl apply -f nginx-deployment.yaml
kubectl apply -f nginx-service.yaml
kubectl get deployments
kubectl get replicasets
kubectl get pods
kubectl get services
```

# Port Forward

```sh
kubectl port-forward service/nginx-service 3000:8080

kubectl port-forward pod/nginx-<random-id> 3001:80
```

# Delete

```sh
kubectl delete -f nginx-deployment.yaml
kubectl delete -f nginx-service.yaml
kubectl get deployments
kubectl get replicasets
kubectl get pods
kubectl get services
```