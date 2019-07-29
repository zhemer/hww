This directory contains Kubernetes version of the system. In order to install all needed deployments and packages, run:

```console
$ ./setup.sh
./setup.sh installs services and deployments on Kubernetes cluster.
Actually, the script just outputs commands to stdout, you need to apped '|sh +x' to command line to execute them.
Usage: ./setup.sh -i|-u
-i              installs services and deployments
-u              remove them

$ ./setup.sh -i
kubectl apply -f hww-service.yaml
kubectl apply -f prometheus-service.yaml
kubectl apply -f grafana-service.yaml
kubectl apply -f hww-deployment.yaml
kubectl apply -f prometheus-deployment.yaml
kubectl apply -f grafana-deployment.yaml


$ ./setup.sh -u|sh +x
deployment.extensions "grafana" deleted
deployment.extensions "hww" deleted
deployment.extensions "prometheus" deleted
service "grafana" deleted
service "hww" deleted
service "prometheus" deleted
configmap "cm-graf-dashboard" deleted
configmap "cm-graf-dashboards" deleted
configmap "cm-graf-datasource" deleted
configmap "cm-prom" deleted
```
