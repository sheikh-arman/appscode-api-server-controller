package controllers

import (
	"context"
	"flag"
	"fmt"
	"github.com/sheikh-arman/controller-appscode-api/pkg/apis/appscode.com/v1alpha1"
	klient "github.com/sheikh-arman/controller-appscode-api/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
)

func RunApi() {

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	klientset, err := klient.NewForConfig(config)
	if err != nil {
		log.Printf("error  %s\n", err.Error())
		panic(err)
	}
	fmt.Println(klientset)
	emp, err := klientset.AppscodeV1alpha1().Employees("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(len(emp.Item))
	Employee := v1alpha1.Employee{}
	fmt.Print(Employee)
}
