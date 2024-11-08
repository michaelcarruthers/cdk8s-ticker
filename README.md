# cdk8s-ticker

cdk8s-ticker is a cdk8s application for managing the Kubernetes deployment of
[ticker](https://github.com/michaelcarruthers/ticker). 

## Prerequisites

* [cdk8s](https://cdk8s.io/)
* [go](https://go.dev/)
* [Taskfile](https://taskfile.dev/) (optional)
* [Visual Studio Code](https://code.visualstudio.com/) or [devcontainers](https://github.com/devcontainers/cli)

## Configuration

Configuration is provided by the [values.yml](./values.yml) file.

| Name           | Description               | Default                                                    |
|----------------|---------------------------|------------------------------------------------------------|
| api_key        | API key                   | REPLACE_ME                                                 |
| environment    | Environment               | dev                                                        |
| image          | Container image           | us-west1-docker.pkg.dev/cloudy-sunday/ticker/ticker:latest |
| name           | Application name          | ticker                                                     |
| namespace      | Kubernetes namespace      | ticker                                                     |
| number_of_days | Number of days to average | 3                                                          |
| replicas       | Kubernetes replicas       | 1                                                          |
| service.name   | Kubernetes service        | ticker                                                     |
| service.port   | Kubernetes service port   | 8080                                                       |
| symbol         | Stock symbol              | MSFT                                                       |


## Build

The build process is responsible for synthensizing the CDK8s application into
a Kubernetes manifest (YAML) file. This is accomplished using the `cdk8s`
command-line utility:

```bash
$ cdktf synth
```

## Deploy

Once the CDK8s application has been successfully synthensized, the Kubernetes
manifest can be applied using the typical `kubectl` process:

```bash
$ kubectl apply -f dist/ticker.k8s.yaml
```

## tl;dr

Using `task` run:

```bash
task
```