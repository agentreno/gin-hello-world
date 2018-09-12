# gin-hello-world

## Description

A sample web app written in Golang to learn about the language (deps, unit
tests, the Gin framework), Docker, and Kubernetes.

## TODO

- Work out why adding in liveness and readiness probe causes crash loops in
  kubernetes deployments (uncomment `k8s-configs/gin-hello-world-deployment.yaml`
  and run `kubectl apply -f k8s-configs/`

- Work out how to trigger k8s deployments when an associated configmap changes
  without downtime:
  https://github.com/kubernetes/kubernetes/issues/22368
  https://github.com/helm/helm/blob/master/docs/charts_tips_and_tricks.md#automatically-roll-deployments-when-configmaps-or-secrets-change
