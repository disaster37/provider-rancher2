# Local test

## First terminal

We will run vcluster

```
export KUBECONFIG=/home/user/.kube/config
curl -L -o vcluster "https://github.com/loft-sh/vcluster/releases/latest/download/vcluster-linux-amd64" && sudo install -c -m 0755 vcluster /usr/local/bin && rm -f vcluster
vcluster create test --namespace vcluster
```

## Second terminal

We will run provider-rancher

```
export KUBECONFIG=/home/user/.kube/config
kubectl apply --server-side -f package/crds
make run
```

## Third terminal

We will create a test resource

```
export KUBECONFIG=/home/user/.kube/config
kubectl apply -f tests/config
kubectl apply -f tests/ressources
```