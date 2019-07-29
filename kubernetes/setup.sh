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
	apps='hww prometheus grafana'
	for app in $apps;do
		echo kubectl apply -f "${app}-service.yaml"
	done
	for app in $apps;do
		echo kubectl apply -f "${app}-deployment.yaml"
	done
;;

-u)
	echo "kubectl delete deployments,services,configmap -l 'app in (hww,grafana,prometheus),owner=zhemer'"
;;
esac
