package controllers

import (
	"fmt"
	"github.com/sheikh-arman/controller-appscode-api/pkg/apis/appscode.com/v1alpha1"
)

func RunApi() {
	Employee := v1alpha1.Employee{}
	fmt.Print(Employee)
}
