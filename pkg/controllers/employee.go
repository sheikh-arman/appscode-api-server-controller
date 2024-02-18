package controllers

import (
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
	if ok := cache.WaitForCacheSync(ch, c.employeeSynced); !ok {
		log.Println("Cache was not synced")
	}
	go wait.Until(c.worker, time.Second, ch)
	<-ch
	return nil
}

func (c *Controller) worker() {
	for c.processNextItem() {

	}
}

func (c *Controller) processNextItem() bool {
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
