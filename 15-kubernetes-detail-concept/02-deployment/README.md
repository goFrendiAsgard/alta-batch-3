# Deploy

```sh
kubectl apply -f nginx-deployment.yaml
kubectl get deployments
kubectl get replicasets
kubectl get pods
```


```sh
kubectl port-forward pod/nginx-<random-id> 3000:80
```

# Delete

```sh
kubectl delete -f nginx-deployment.yaml
kubectl get deployments
kubectl get replicasets
kubectl get pods
```