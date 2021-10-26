# Deploy

```sh
kubectl apply -f simple-pod.yaml
kubectl get pods
```

# Port Forward

```sh
kubectl port-forward pod/nginx 3000:80
```

# Delete

```sh
kubectl delete -f simple-pod.yaml
kubectl get pods
```