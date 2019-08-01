package main

import (
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	clusterCfg, err := clientcmd.BuildConfigFromFlags("", "config")
	if err != nil {
		panic(err)
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(clusterCfg)
	if err != nil {
		panic(err.Error())
	}

	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	nodeip := []corev1.NodeAddress{}
	for i := 0; i < len(nodes.Items); i++ {
		nodeip = nodes.Items[i].Status.Addresses
		fmt.Println(nodeip[0].Address)
	}
	fmt.Println(nodes.Items[0].Status.Addresses)

	services, err := clientset.CoreV1().Services("tekton-sources").List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	svc := corev1.Service{}
	for _, service := range services.Items {
		if strings.HasPrefix(service.GetName(), "gitlabsample-") {
			svc = service
		}
	}

	fmt.Println(svc.Spec.Ports[0].NodePort)
	fmt.Println(svc.Status.LoadBalancer.Ingress[0].IP)

	// dynamicClient, err := dynamic.NewForConfig(clusterCfg)
	// if err != nil {
	// 	panic(err)
	// }

	// gvr := schema.GroupVersionResource{
	// 	Group:    "clusterregistry.k8s.io",
	// 	Version:  "v1alpha1",
	// 	Resource: "clusters",
	// }

	/*
	   gvr := schema.GroupVersionResource{
	       Group:    “”,
	       Version:  “v1”,
	       Resource: “namespaces”,
	   }
	*/

	// stopCh := make(chan struct{})
	/*
	   informer := dynamicinformer.NewFilteredDynamicInformer(dynamicClient, gvr, “”, 0*time.Second, cache.Indexers{},
	       func(options *metav1.ListOptions) {
	           options.LabelSelector = “cloud=IBM,multicloud-ha=enabled”
	       })

	   informer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
	       UpdateFunc: func(oldObj, newObj interface{}) {
	           unstructured, ok := newObj.(*unstructured.Unstructured)
	           if !ok {
	               panic(“Invalid Unstructured object”)
	           }

	           fmt.Printf(“-->Object updated: %v/%v, %+v\n”, unstructured.GetNamespace(), unstructured.GetName(), unstructured.GetLabels())
	       },
	   })


	   go informer.Informer().Run(stopCh)
	*/

	// informerFactory := dynamicinformer.NewDynamicSharedInformerFactory(dynamicClient, 30*time.Second)
	// informer := informerFactory.ForResource(gvr)
	/*
	   informer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
	       UpdateFunc: func(oldObj, newObj interface{}) {
	           unstructured, ok := newObj.(*unstructured.Unstructured)
	           if !ok {
	               panic(“Invalid Unstructured object”)
	           }

	           fmt.Printf(“-->Object updated: %v/%v, %+v\n”, unstructured.GetNamespace(), unstructured.GetName(), unstructured.GetLabels())
	       },
	   })
	*/

	// indexers := cache.Indexers{
	//     “key”: func(obj interface{}) ([]string, error) {
	//         unstructured, ok := obj.(*unstructured.Unstructured)
	//         if !ok {
	//             panic(“Invalid Unstructured object”)
	//         }

	//         return []string{fmt.Sprintf(“key_%s_%s”, unstructured.GetNamespace(), unstructured.GetName())}, nil
	//     },
	// }

	// informer.Informer().AddIndexers(indexers)

	// go informerFactory.Start(stopCh)

	// fmt.Printf(“Waiting for informer synced...\n”)
	// cache.WaitForCacheSync(stopCh, informer.Informer().HasSynced)
	// fmt.Printf(“Informer synced\n”)

	/*
	   //labelSelector, err := labels.Parse(“cloud=IBM,multicloud-ha=enabled”)
	   selector, err := labels.Parse(“cloud=IBM,multicloud-manager-ha=enabled”)
	   labelSelector, err := labels.Parse(“”)
	   if err != nil {
	       panic(err)
	   }

	   for i := 0; i < 10; i++ {
	       objs, err := informer.Lister().List(labelSelector)
	       if err != nil {
	           panic(err)
	       }
	       count := 0
	       for _, obj := range objs {
	           unstructured, ok := obj.(*unstructured.Unstructured)
	           if !ok {
	               panic(“Invalid Unstructured object”)
	           }
	           labels := labels.Set(unstructured.GetLabels())
	           fmt.Printf(“=====>matched: %v\n”, selector.Matches(labels))
	           fmt.Printf(“%v/%v, %+v\n”, unstructured.GetNamespace(), unstructured.GetName(), unstructured.GetLabels())
	           count++
	       }

	       fmt.Printf(“%d objects found\n”, count)
	       fmt.Printf(“Sleeping for 15 seconds...\n”)
	       time.Sleep(15 * time.Second)
	   }
	*/

	// for {
	//     ns, err := informer.Informer().GetIndexer().ByIndex(“key”, “key_cn1_c1”)
	//     if err != nil {
	//         panic(err)
	//     } else {
	//         fmt.Printf(“ns: %+v\n”, ns[0])
	//     }
	//     time.Sleep(15 * time.Second)
	// }

	// close(stopCh)
}
