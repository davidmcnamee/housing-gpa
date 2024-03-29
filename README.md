# Student Housing


## Local Dev Setup

First, install these:
* [bazelisk](https://github.com/bazelbuild/bazelisk#installation)
* [docker](https://www.docker.com/products/docker-desktop)
* [k3d](https://k3d.io/#installation)
* [skaffold](https://skaffold.dev/docs/install/)

1. Start your docker daemon by opening the Docker Desktop app
2. Create a kubernetes cluster with k3d
```
k3d cluster create
k3d cluster start
kubectl config use-context k3d-k3s-default
```
3. run `skaffold dev` (this will take a while if it's your first time)

> `localhost:8080` serves the frontend app,
> `localhost:3000` serves the GraphQL playground

### Running scripts

Create a `.env` file in the root directory (containing the secrets given to you), then run `bazel run //scripts:my_script`.

## Roadmap

- [x] Cluster & Dev environment setup
- [x] Backend GraphQL API functional
- [ ] first draft of site
- [x] initial database population with python scripts
- [ ] launch, promotion, user feedback


