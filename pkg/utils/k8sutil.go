/*
Copyright (c) 2016-2017 Bitnami

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"encoding/json"
	kinesisApi "github.com/kubeless/kinesis-trigger/pkg/apis/kubeless/v1beta1"
	"github.com/sirupsen/logrus"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/kubeless/kinesis-trigger/pkg/client/clientset/versioned"
)

const (
	defaultTimeout = "180"
)

// GetClient returns a k8s clientset to the request from inside of cluster
func GetClient() kubernetes.Interface {
	config, err := rest.InClusterConfig()
	if err != nil {
		logrus.Fatalf("Can not get kubernetes config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logrus.Fatalf("Can not create kubernetes client: %v", err)
	}

	return clientset
}

// GetTriggerClientInCluster returns function clientset to the request from inside of cluster
func GetTriggerClientInCluster() (versioned.Interface, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	kinesisClient, err := versioned.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return kinesisClient, nil
}

// CreateKinesisTriggerCustomResource will create a Kinesis trigger custom resource object
func CreateKinesisTriggerCustomResource(kubelessClient versioned.Interface, kinesisTrigger *kinesisApi.KinesisTrigger) error {
	_, err := kubelessClient.KubelessV1beta1().KinesisTriggers(kinesisTrigger.Namespace).Create(kinesisTrigger)
	if err != nil {
		return err
	}
	return nil
}

// UpdateKinesisTriggerCustomResource applies changes to the Kinesis trigger custom resource object
func UpdateKinesisTriggerCustomResource(kubelessClient versioned.Interface, kinesisTrigger *kinesisApi.KinesisTrigger) error {
	_, err := kubelessClient.KubelessV1beta1().KinesisTriggers(kinesisTrigger.Namespace).Update(kinesisTrigger)
	return err
}

// PatchKinesisTriggerCustomResource applies changes to the function custom object
func PatchKinesisTriggerCustomResource(kubelessClient versioned.Interface, kinesisTrigger *kinesisApi.KinesisTrigger) error {
	data, err := json.Marshal(kinesisTrigger)
	if err != nil {
		return err
	}
	_, err = kubelessClient.KubelessV1beta1().KinesisTriggers(kinesisTrigger.Namespace).Patch(kinesisTrigger.Name, types.MergePatchType, data)
	return err
}

// DeleteKinesisTriggerCustomResource will delete  HTTP trigger custom resource object
func DeleteKinesisTriggerCustomResource(kubelessClient versioned.Interface, kinesisTriggerName, ns string) error {
	err := kubelessClient.KubelessV1beta1().KinesisTriggers(ns).Delete(kinesisTriggerName, &metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}

// GetKinesisTriggerCustomResource will get  HTTP trigger custom resource object
func GetKinesisTriggerCustomResource(kubelessClient versioned.Interface, kinesisTriggerName, ns string) (*kinesisApi.KinesisTrigger, error) {
	kinesisCRD, err := kubelessClient.KubelessV1beta1().KinesisTriggers(ns).Get(kinesisTriggerName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return kinesisCRD, nil
}
