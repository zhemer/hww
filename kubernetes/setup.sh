apps='hww prometheus grafana'
for app in $apps;do
  echo kubectl apply -f "${app}-service.yaml"
done
for app in $apps;do
  echo kubectl apply -f "${app}-deployment.yaml"
done

