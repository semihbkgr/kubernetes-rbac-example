package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	namespaceEnvVar    = "NAMESPACE"
	podNameEnvVar      = "POD_NAME"
	serviceNameEnvVar  = "SERVICE_NAME"
	secretNameEnvVar   = "SECRET_NAME"
	getPerSecondEnvVar = "GET_PER_SECOND"
)

var (
	namespace    string
	podName      string
	serviceName  string
	secretName   string
	getPerSecond = 10
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)
	ctx := context.Background()

	namespaceEnvVarValue, ok := os.LookupEnv(namespaceEnvVar)
	if !ok {
		err := fmt.Errorf("%s environment variable is not set", namespaceEnvVar)
		panic(err)
	}
	if len(namespaceEnvVarValue) == 0 {
		err := fmt.Errorf("%s environment variable cannot be empty", namespaceEnvVar)
		panic(err)
	}
	namespace = namespaceEnvVarValue
	podNameEnvVarValue, ok := os.LookupEnv(podNameEnvVar)
	if !ok {
		err := fmt.Errorf("%s environment variable is not set", podNameEnvVar)
		panic(err)
	}
	if len(podNameEnvVarValue) == 0 {
		err := fmt.Errorf("%s environment variable cannot be empty", podNameEnvVar)
		panic(err)
	}
	podName = podNameEnvVarValue
	serviceNameEnvVarValue, ok := os.LookupEnv(serviceNameEnvVar)
	if !ok {
		err := fmt.Errorf("%s environment variable is not set", serviceNameEnvVar)
		panic(err)
	}
	if len(serviceNameEnvVarValue) == 0 {
		err := fmt.Errorf("%s environment variable cannot be empty", serviceNameEnvVar)
		panic(err)
	}
	serviceName = serviceNameEnvVarValue
	secretNameEnvVarValue, ok := os.LookupEnv(secretNameEnvVar)
	if !ok {
		err := fmt.Errorf("%s environment variable is not set", secretNameEnvVar)
		panic(err)
	}
	if len(secretNameEnvVarValue) == 0 {
		err := fmt.Errorf("%s environment variable cannot be empty", secretNameEnvVar)
		panic(err)
	}
	secretName = secretNameEnvVarValue
	getPerSecondEnvVarValue, ok := os.LookupEnv(getPerSecondEnvVar)
	if ok {
		getPerSecondIntValue, err := strconv.Atoi(getPerSecondEnvVarValue)
		if err != nil {
			panic(err)
		}
		getPerSecond = getPerSecondIntValue
	}
	log.Printf("Namespace: %s\n", namespace)
	log.Printf("PodName: %s\n", podName)
	log.Printf("ServiceName: %s\n", serviceName)
	log.Printf("SecretName: %s\n", secretName)
	log.Printf("GetPerSecond: %d\n", getPerSecond)

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	c := time.Tick(time.Duration(getPerSecond) * time.Second)

	for {
		_, ok := <-c
		if ok {
			log.Println(strings.Repeat("-", 90))

			// Pods
			pods, err := clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
			if err != nil {
				log.Println(err.Error())
			}
			log.Printf("There are %d pods in the cluster\n", len(pods.Items))

			pods, err = clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
			if err != nil {
				log.Println(err.Error())
			}
			log.Printf("There are %d pods in '%s' namespace\n", len(pods.Items), namespace)

			_, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
			if errors.IsNotFound(err) {
				log.Printf("Pod '%s' not found in '%s' namespace\n", podName, namespace)
			} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
				log.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
			} else if err != nil {
				log.Println(err.Error())
			} else {
				log.Printf("Found '%s' pod in '%s' namespace\n", podName, namespace)
			}

			// Services
			services, err := clientset.CoreV1().Services("").List(ctx, metav1.ListOptions{})
			if err != nil {
				log.Println(err.Error())
			}
			log.Printf("There are %d services in the cluster\n", len(services.Items))

			services, err = clientset.CoreV1().Services(namespace).List(ctx, metav1.ListOptions{})
			if err != nil {
				log.Println(err.Error())
			}
			log.Printf("There are %d services in '%s' namespace\n", len(services.Items), namespace)

			_, err = clientset.CoreV1().Services(namespace).Get(context.TODO(), serviceName, metav1.GetOptions{})
			if errors.IsNotFound(err) {
				log.Printf("Service '%s' not found in '%s' namespace\n", serviceName, namespace)
			} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
				log.Printf("Error getting service %v\n", statusError.ErrStatus.Message)
			} else if err != nil {
				log.Println(err.Error())
			} else {
				log.Printf("Found '%s' service in '%s' namespace\n", serviceName, namespace)
			}

			// Secrets
			secrets, err := clientset.CoreV1().Secrets("").List(ctx, metav1.ListOptions{})
			if err != nil {
				log.Println(err.Error())
			}
			log.Printf("There are %d secrets in the cluster\n", len(secrets.Items))

			secrets, err = clientset.CoreV1().Secrets(namespace).List(ctx, metav1.ListOptions{})
			if err != nil {
				log.Println(err.Error())
			}
			log.Printf("There are %d secrets in '%s' namespace\n", len(secrets.Items), namespace)

			_, err = clientset.CoreV1().Secrets(namespace).Get(context.TODO(), secretName, metav1.GetOptions{})
			if errors.IsNotFound(err) {
				log.Printf("Secret '%s' not found in '%s' namespace\n", secretName, namespace)
			} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
				log.Printf("Error getting secret %v\n", statusError.ErrStatus.Message)
			} else if err != nil {
				log.Println(err.Error())
			} else {
				log.Printf("Found '%s' secret in '%s' namespace\n", secretName, namespace)
			}
		} else {
			break
		}
	}

}
