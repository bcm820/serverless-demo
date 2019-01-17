## Initial Setup

### Install [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/#install-kubectl)

The command-line tool, kubectl, is used to deploy and manage applications on Kubernetes.

```
brew install kubernetes-cli
```

### Install [Minikube](https://github.com/kubernetes/minikube#installation)

Minikube is a tool that makes it easy to run Kubernetes locally. It runs a single-node Kubernetes cluster inside a VM on your laptop.

```
brew cask install minikube
```

### Install [Kubeless](https://kubeless.io/)

Kubeless is a Kubernetes-native serverless framework that lets you deploy small bits of code (functions) without having to worry about the underlying infrastructure.

Remember to check for the most up-to-date version (current is 1.0.1)!

```
curl -L https://github.com/kubeless/kubeless/releases/download/v1.0.1/kubeless_darwin-amd64.zip
unzip kubeless_darwin-amd64.zip
sudo cp bundles/kubeless_darwin-amd64/kubeless /usr/local/bin/
```

## Launch Kubernetes locally

### Start Minikube

```
minikube start
```

### Create Kubeless namespace and add resource to cluster

```
kubectl create ns kubeless
kubectl apply -f https://github.com/kubeless/kubeless/releases/download/v1.0.1/kubeless-v1.0.1.yaml
```

### Check cluster

```
minikube status
kubectl cluster-info
```

### Start dashboard

```
minikube dashboard
```

## Deploy a sample Python function

[Kubeless quickstart guide](https://kubeless.io/docs/quick-start/)

Run the following to deploy a Python function:

```
kubeless function deploy hello --runtime python2.7 \
 --from-file hello.py \
 --handler hello.foobar
```

Check that the function has been deployed.
You can also check its status dashboard under `Overview -> Deployments`.

```
kubeless function ls hello
minikube service list
kubectl get po
kubectl get svc
```

Call the function from the terminal:

```
kubeless function call hello --data 'Hello world!'
```

## Expose the sample function

Now, we need to enable ingress to expose the function via HTTP. Run the following and confirm that `default-http-backend` is added.

```
minikube addons enable ingress
minikube service list
```

Create an HTTP trigger object and check for it:

```
kubeless trigger http create hello --function-name hello --path hello
kubectl get ing
```

Now you can call the function via HTTP!

```
curl --data 'HELLO WORLD' hello.192.168.99.100.nip.io/hello

curl --data '{"message": "Hello World!"}' \
 --header "Host: hello.192.168.99.100.nip.io" \
 --header "Content-Type:application/json" \
 192.168.99.100.nip.io/hello
```

## Deploy a sample NodeJS function

### Deploy:

```
kubeless function deploy quicksort --runtime nodejs6 \
 --from-file quicksort.js \
 --handler quicksort.qs
```

### Create HTTP trigger

```
kubeless trigger http create hello --function-name hello --path hello
```

### Test

```
kubeless function call quicksort --data '[91,1,23,3,-10,37,85,5,11,62,6,76,49,7,53,-3]'
```

```
curl --data '[91,1,23,3,-10,37,85,5,11,62,6,76,49,7,53,-3]' \
 --header "Host: quicksort.192.168.99.100.nip.io" \
 --header "Content-Type:application/json" \
 192.168.99.100.nip.io/quicksort
```

## Clean up

```
minikube delete
```
