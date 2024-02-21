package controllers

import (
	"flag"
	klient "github.com/sheikh-arman/controller-appscode-api/pkg/client/clientset/versioned"
	informer "github.com/sheikh-arman/controller-appscode-api/pkg/client/informers/externalversions"
	"k8s.io/client-go/kubernetes"
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
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Printf("My on getting kubernetes clientset %s\n", err.Error())
		panic(err)
	}
	//res, err := clientset.CoreV1().Pods("appscode").List(context.Background(), metav1.ListOptions{})
	////fmt.Println(clientset)
	//fmt.Println("Pods are:", len(res.Items))
	//for _, e := range res.Items {
	//	fmt.Println(e.Name)
	//}

	klientset, err := klient.NewForConfig(config)
	if err != nil {
		log.Printf("My Error!!!!!!!!!!  %s\n", err.Error())
		panic(err)
	}

	//emp, err := klientset.AppscodeV1alpha1().Employees("").List(context.Background(), metav1.ListOptions{})
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(len(emp.Items))
	//for _, em := range emp.Items {
	//	fmt.Println(em.Name, em.Spec.Image)
	//}
	//Employee := v1alpha1.Employee{}
	//fmt.Println(Employee)

	_ = klientset
	//start controller
	ch := make(chan struct{})
	informerFactory := informer.NewSharedInformerFactory(klientset, 20*time.Minute)
	//fmt.Println("Check 11->>>>>>>")
	c := NewController(clientset, klientset, informerFactory.Appscode().V1alpha1().Employees())
	//fmt.Println("Check 12->>>>>>>")
	informerFactory.Start(ch)
	//fmt.Println("Check 13->>>>>>>")
	log.Println("Starting controller...")
	if err = c.Run(ch); err != nil {
		log.Printf("My Error!!!!!!!!!! %s", err.Error())
	}
	//end controller
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
