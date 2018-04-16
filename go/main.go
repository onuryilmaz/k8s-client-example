package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"

	"k8s.io/client-go/rest"
)

func main() {
	var kubeconfig *string
	var config *rest.Config
	var err error

	if home := os.Getenv("HOME"); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")

	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	if *kubeconfig != "" {
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {

			// Test for in-cluster configuration
			config, err = rest.InClusterConfig()
			if err != nil {
				panic(err.Error())
			}
		}
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("There are %d pods in the cluster:\n", len(pods.Items))

		for _, pod := range pods.Items {
			fmt.Println(pod.Namespace + "/" + pod.Name)
		}

		time.Sleep(10 * time.Second)
	}
}
