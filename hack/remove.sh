#!/bin/bash

kubectl delete -f examples/employee.yaml
kubectl delete -f manifests/appscode.com_employees.yaml
