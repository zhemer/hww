This directory contains Kubernetes version of the system.  

In order to install all needed deployments and packages, run:
```console
H=`kubectl describe node|grep hostname|awk -F= '{print $2}'`; sed -i "s/your-host/$H/" pv.yaml
kubectl apply -f .
````

To remove the application execute:
```console
kubectl delete -f .
````

To view the Grafana dashboard (or visit http://yourhost:3000 URL in browser):
```console
kubectl port-forward service/grafana 3000:3000 --address=0.0.0.0 >/tmp/grafana &
curl -i `hostname`:3000|head -22
````
