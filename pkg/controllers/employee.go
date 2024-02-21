package controllers

import (
	"context"
	"fmt"
	v1alpha1 "github.com/sheikh-arman/controller-appscode-api/pkg/apis/appscode.com/v1alpha1"
	klient "github.com/sheikh-arman/controller-appscode-api/pkg/client/clientset/versioned"
	informer "github.com/sheikh-arman/controller-appscode-api/pkg/client/informers/externalversions/appscode.com/v1alpha1"
	lister "github.com/sheikh-arman/controller-appscode-api/pkg/client/listers/appscode.com/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"log"
	"time"
)

type Controller struct {
	//clientset for kubernetes resource
	clientset kubernetes.Interface
	//klientset for custom resource employee
	klient klient.Interface
	// employee synced
	employeeSynced cache.InformerSynced
	//lisetr
	lister lister.EmployeeLister
	//queue
	wordQueue workqueue.RateLimitingInterface
}

func NewController(clientset kubernetes.Interface, klient klient.Interface, employeeInformer informer.EmployeeInformer) *Controller {
	c := &Controller{
		clientset:      clientset,
		klient:         klient,
		employeeSynced: employeeInformer.Informer().HasSynced,
		lister:         employeeInformer.Lister(),
		wordQueue:      workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()),
	}
	employeeInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc:    c.handleAdd,
			DeleteFunc: c.handleDelete,
		},
	)
	return c
}

func (c *Controller) Run(ch chan struct{}) error {
	//fmt.Println("Check culprit->>>>>>>")
	if ok := cache.WaitForCacheSync(ch, c.employeeSynced); !ok {
		log.Println("Cache was not synced")
	}
	//fmt.Println("Check 1->>>>>>>")
	go wait.Until(c.worker, time.Second, ch)
	<-ch
	return nil
}

func (c *Controller) worker() {
	for c.processNextItem() {

	}
}

func (c *Controller) processNextItem() bool {
	item, shutDown := c.wordQueue.Get()
	if shutDown {
		log.Println("Item shutdown")
		return false
	}
	key, err := cache.MetaNamespaceKeyFunc(item)
	if err != nil {
		log.Printf("err %s calling namespace key func", err.Error())
		return false
	}
	ns, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		log.Printf("err %s calling split meta namespace", err.Error())
		return false
	}
	employee, err := c.lister.Employees(ns).Get(name)
	if err != nil {
		log.Printf("err %s calling lister func", err.Error())
		return false
	}
	log.Printf("\nEmployee resource:\nName:%s\nSpec%s\n", employee.Name, employee.Spec)
	c.createDeployment(ns, employee)
	//c.checkClientSet()
	return true
}
func int32Ptr(i int32) *int32 { return &i }
func (c *Controller) handleAdd(obj interface{}) {
	log.Println("Resource added")
	c.wordQueue.Add(obj)
}

func (c *Controller) handleDelete(obj interface{}) {
	log.Println("Resource removed")
	c.wordQueue.Add(obj)
}

func (c *Controller) createDeployment(ns string, employee *v1alpha1.Employee) {
	//create deployment for the employee CR
	deploymentClientSet := c.clientset.AppsV1().Deployments(ns)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: employee.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": employee.Name,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": employee.Name,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  employee.Name,
							Image: employee.Spec.Image,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 8080,
								},
							},
						},
					},
				},
			},
		},
	}
	log.Printf("Creating Deployemnet for %s resource", employee.Name)
	result, err := deploymentClientSet.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		log.Printf("Error on creating deployment %s", employee.Name)
		return
	}
	log.Printf("Created deployment %s\n", result.GetObjectMeta().GetName())
}

func (c *Controller) checkClientSet() {
	res, err := c.clientset.CoreV1().Pods("appscode").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Println(err)
		return
	}
	//fmt.Println(clientset)
	fmt.Println("Pods are:", len(res.Items))
	for _, e := range res.Items {
		fmt.Println(e.Name)
	}
}
