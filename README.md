# controller-appscode-api

vendor file will not work. so use go mod tidy only.

sh hack/code-generator.sh

kubectl apply -f manifest/appscode.com_employees.yaml

kubectl apply -f example/employee.yaml

go run main.go





