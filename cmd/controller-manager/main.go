package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	"phant-operator/pkg/client/clientset/versioned"

	informers "phant-operator/pkg/client/informers/externalversions"

	"phant-operator/pkg/constants"

	"k8s.io/apimachinery/pkg/labels"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

func main() {
	var cfg *rest.Config
	hostName, err := os.Hostname()
	if err != nil {
		klog.Fatalf("failed to get hostname: %v", err)
	}

	kubconfig := os.Getenv("KUBECONFIG")
	if kubconfig != "" {
		cfg, err = clientcmd.BuildConfigFromFlags("", kubconfig)
	} else {
		cfg, err = rest.InClusterConfig()
	}
	if err != nil {
		klog.Fatalf("failed to get config: %v", err)
	}
	klog.Info(hostName)
	// If they are zero, the created client will use the default values: 5, 10.
	/*cfg.QPS = float32(cliCfg.KubeClientQPS)
	cfg.Burst = cliCfg.KubeClientBurst*/
	cli, err := versioned.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("failed to create Clientset: %v", err)
	}
	var kubeCli kubernetes.Interface
	kubeCli, err = kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("failed to get kubernetes Clientset: %v", err)
	}
	ns := "pg-admin"
	options := []informers.SharedInformerOption{
		informers.WithNamespace(ns),
	}
	informerFactory := informers.NewSharedInformerFactoryWithOptions(cli, constants.ResyncDuration, options...)
	postgresInformer := informerFactory.Phant().V1alpha1().PostgresClusters()
	kubeoptions := []kubeinformers.SharedInformerOption{kubeinformers.WithNamespace(ns)}
	kubeInformerFactory := kubeinformers.NewSharedInformerFactoryWithOptions(kubeCli, constants.ResyncDuration, kubeoptions...)
	stsInformer := kubeInformerFactory.Apps().V1().StatefulSets()
	//fmt.Println(cfg.Burst)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		wg sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// waiting for the shared informer's store has synced.
		go informerFactory.Start(ctx.Done())
		cache.WaitForCacheSync(ctx.Done(), postgresInformer.Informer().HasSynced)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		// waiting for the shared informer's store has synced.
		go kubeInformerFactory.Start(ctx.Done())
		cache.WaitForCacheSync(ctx.Done(), stsInformer.Informer().HasSynced)
	}()
	wg.Wait()
	pgcs, _ := postgresInformer.Lister().PostgresClusters(ns).List(labels.Everything())
	klog.Info("Get CR pgc")
	for _, pgc := range pgcs {
		fmt.Println(pgc.Name)
	}

}
