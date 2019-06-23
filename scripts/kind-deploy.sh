./tools/kind create cluster --config=tools/kind-config.yaml --name="calc-test"
export KUBECONFIG="$(./tools/kind get kubeconfig-path --name="calc-test")"
./tools/kind load docker-image calc-server:v0 --name="calc-test"
kubectl apply -f deployments/deployment.yaml
