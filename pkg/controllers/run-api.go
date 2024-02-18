package controllers

import (
	"flag"
	klient "github.com/sheikh-arman/controller-appscode-api/pkg/client/clientset/versioned"
	informer "github.com/sheikh-arman/controller-appscode-api/pkg/client/informers/externalversions"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
	"time"
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
		log.Printf("My Error!!!!!!!!!!  %s\n", err.Error())
		panic(err)
	}
	_ = klientset
	ch := make(chan struct{})

	informerFactory := informer.NewSharedInformerFactory(klientset, 20*time.Minute)
	c := NewController(klientset, informerFactory.Appscode().V1alpha1().Employees())
	informerFactory.Start(ch)
	if err = c.Run(ch); err != nil {
		log.Printf("My Error!!!!!!!!!! %s", err.Error())
	}

	//fmt.Println(klientset)
	//emp, err := klientset.AppscodeV1alpha1().Employees("").List(context.Background(), metav1.ListOptions{})
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(len(emp.Item))
	//for _, em := range emp.Item {
	//	fmt.Println(em.Name, em.Spec.Image)
	//}
	//Employee := v1alpha1.Employee{}
	//fmt.Print(Employee)
}
