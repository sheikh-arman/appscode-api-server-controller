package controllers

import (
	"fmt"
	klient "github.com/sheikh-arman/controller-appscode-api/pkg/client/clientset/versioned"
	informer "github.com/sheikh-arman/controller-appscode-api/pkg/client/informers/externalversions/appscode.com/v1alpha1"
	lister "github.com/sheikh-arman/controller-appscode-api/pkg/client/listers/appscode.com/v1alpha1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"log"
	"time"
)

type Controller struct {
	//klientset for custom resource employee
	klient klient.Interface
	// employee synced
	employeeSynced cache.InformerSynced
	//lisetr
	lister lister.EmployeeLister
	//queue
	wordQueue workqueue.RateLimitingInterface
}

func NewController(klient klient.Interface, employeeInformer informer.EmployeeInformer) *Controller {
	c := &Controller{
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
	fmt.Println("Check culprit->>>>>>>")
	if ok := cache.WaitForCacheSync(ch, c.employeeSynced); !ok {
		log.Println("Cache was not synced")
	}
	fmt.Println("Check 1->>>>>>>")
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
	return true
}
func (c *Controller) handleAdd(obj interface{}) {
	log.Println("Resource added")
	c.wordQueue.Add(obj)
}

func (c *Controller) handleDelete(obj interface{}) {
	log.Println("Resource removed")
	c.wordQueue.Add(obj)
}
