PV_PATH=/mnt/kubernetes/vol0

if [ "$1" != '-i' ] && [ "$1" != '-u' ];then
    cat << EOF
$0 installs services and deployments on Kubernetes cluster.
Actually, the script just outputs commands to stdout, you need to apped '|sh +x' to command line to execute them.
Usage: $0 -i|-u
-i 		installs services and deployments
-u 		remove them
EOF
    exit
fi

case "$1" in
-i):
	cat << EOF
mkdir -p $PV_PATH
chown nobody.nogroup $PV_PATH
kubectl apply -f pv.yaml,pvc.yaml
EOF
	apps='hww prometheus grafana'
	for app in $apps;do
		echo kubectl apply -f "${app}-service.yaml"
	done
	for app in $apps;do
		echo kubectl apply -f "${app}-deployment.yaml"
	done
echo "kubectl port-forward `kubectl get pod -l app=grafana|tail -1|awk '{print $1}'` 3000:3000 --address=0.0.0.0 >/tmp/pf-graf &""
;;

-u)
	cat << EOF
kubectl delete deployments,services,configmap,pv,pvc -l 'app in (hww,grafana,prometheus),owner=zhemer'
rm -rf $PV_PATH
EOF
;;
esac
