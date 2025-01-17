# Deploying a Grafana instance

This document describes how to get up and running with a new Grafana instance no Kubernetes.

## Deploying the operator

The first step is to install the Grafana operator to a namespace in your cluster.

To create a namespace named `grafana` run:

```sh
$ kubectl create namespace grafana
```

To deploy the operator to that namespace you can use `deploy/operator.yaml`:

```sh
$ kubectl create -f deploy/operator.yaml -n grafana
```

Check the status of the operator pod:

```sh
$ kubectl get pods -n grafana
NAME                                READY     STATUS    RESTARTS   AGE
grafana-operator-78cfcbf8db-ssrgq   1/1       Running   0          17s
```

## Operator flags

The operator accepts a number of flags that can be passed in the `args` section of the container in the deployment:

* *--grafana-image*: overrides the Grafana image, defaults to `docker.io/grafana/grafana`.
* *--grafana-image-tag*: overrides the Grafana tag, defaults to `5.4.2`.
* *--scan-all*: watch for dashboards in all namespaces. This requires the the operator service account to have cluster wide permissions to `get`, `list`, `update` and `watch` dashboards. See `deploy/examples/cluster_roles`.
* *--openshift*: force the operator to use a [route](https://docs.openshift.com/container-platform/3.11/architecture/networking/routes.html) instead of an [ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/). Note that routes are only supported on OpenShift.

See `deploy/operator.yaml` for an example.

## Deploying Grafana

Create a custom resource of type `Grafana`, or use the one in `deploy/examples/Grafana.yaml`.

The resource accepts the following properties in it's `spec`:

* *hostname*:  The host to be used for the [ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/). Optional when `--openshift` is set.
* *dashboardLabelSelector*: A list of either `matchLabels` or `matchExpressions` to filter the dashboards before importing them.
* *containers*: Extra containers to be added to the Grafana deployment. Can be used for example to add auth proxy side cars.
* *secrets*: A list of secrets that are added as volumes to the deployment. Mostly useful in combination with extra `containers`.

The other accepted properties are `logLevel`, `adminUser`, `adminPassword`, `basicAuth`, `disableLoginForm`, `disableSignoutMenu` and `anonymous`. They map to the properties described in the [official documentation](https://grafana.com/docs/installation/configuration/#configuration), just use camel case instead of underscores.

To create a new Grafana instance in the `grafana` namespace, run:

```sh
$ kubectl create -f deploy/examples/Grafana.yaml -n grafana
```

Get the URL of the instance and open it in a browser:

```sh
$ kubectl get ingress -n grafana
NAME              HOSTS                           ADDRESS   PORTS     AGE
grafana-ingress   grafana.apps.127.0.0.1.nip.io             80        28s
```
