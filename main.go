package main

import (
	"context"
	"log"

	// sockshopv1 "datamodel/build/apis/root.sockshop.com/v1"
	baseconfigdatamodelv1 "example/build/apis/config.example.com/v1"
	baserootdatamodelv1 "example/build/apis/root.example.com/v1"
	basetenantdatamodelv1 "example/build/apis/tenant.example.com/v1"

	nexus_client "example/build/nexus-client"

	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func ProcessWanna(wanna *nexus_client.WannaWanna) {

	user, userErr := wanna.GetParent(context.TODO())
	if userErr != nil {
		log.Panicf("user lookup failed for wanna %s with error %+v", wanna.DisplayName(), userErr)
	}

	config, configErr := user.GetParent(context.TODO())
	if configErr != nil {
		log.Panicf("config lookup failed for user %s with error %+v", user.DisplayName(), configErr)
	}

	tenant, tenantErr := config.GetParent(context.TODO())
	if tenantErr != nil {
		log.Panicf("tenant lookup failed for config %s with error %+v", config.DisplayName(), tenantErr)
	}

	interest, interestLkupErr := tenant.GetInterest(context.TODO(), wanna.Spec.Name)
	if interestLkupErr != nil {
		fmt.Printf("Interest %s lookup failed for wanns with error %+v\n", wanna.Spec.Name, interestLkupErr)
		return
	}

	linkErr := wanna.LinkInterest(context.TODO(), interest)
	if linkErr != nil {
		fmt.Printf("Linking interest %s with wanna %s failed with error %+v\n", interest.DisplayName(), wanna.DisplayName(), linkErr)
		return
	}
	fmt.Printf("Succesfully lined wanna %s with interest %s\n", wanna.DisplayName(), interest.DisplayName())
}

var nexusClient *nexus_client.Clientset

func main() {

	rand.Seed(time.Now().UnixNano())
	var kubeconfig string
	flag.StringVar(&kubeconfig, "k", "", "Absolute path to the kubeconfig file. Defaults to ~/.kube/config.")
	flag.Parse()

	var config *rest.Config
	if len(kubeconfig) != 0 {
		var err error
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			panic(err)
		}
	} else {
		config = &rest.Config{Host: "localhost:8081"}
	}

	nexusClient, _ = nexus_client.NewForConfig(config)

	nexusClient.RootRoot().Subscribe()
	nexusClient.RootRoot().Tenant("*").Subscribe()
	nexusClient.RootRoot().Tenant("*").Interest("*").Subscribe()

	root, rootErr := nexusClient.AddRootRoot(context.TODO(), &baserootdatamodelv1.Root{})
	if nexus_client.IsAlreadyExists(rootErr) {
		root, rootErr = nexusClient.GetRootRoot(context.TODO())
		if rootErr != nil {
			log.Panicf("failed getting root node with error %+v", rootErr)
		}
	}

	tenantObj := &basetenantdatamodelv1.Tenant{
		ObjectMeta: metav1.ObjectMeta{
			Name: "default",
		},
		Spec: basetenantdatamodelv1.TenantSpec{},
	}
	tenant, tenantErr := root.AddTenant(context.TODO(), tenantObj)
	if nexus_client.IsAlreadyExists(tenantErr) {
		tenant, tenantErr = root.GetTenant(context.TODO(), tenantObj.Name)
		if tenantErr != nil {
			log.Panicf("failed getting tenant node with error %+v", tenantErr)
		}
	}

	configObj := &baseconfigdatamodelv1.Config{}
	_, configErr := tenant.AddConfig(context.TODO(), configObj)
	if nexus_client.IsAlreadyExists(configErr) {
		_, configErr = tenant.GetConfig(context.TODO())
		if configErr != nil {
			log.Panicf("failed getting config node with error %+v", configErr)
		}
	}

	nexusClient.RootRoot().Tenant("*").Config().User("*").Wanna("*").RegisterAddCallback(ProcessWanna)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {
		<-sigs
		done <- true
	}()
	<-done
	fmt.Println("exiting")
}
