package helper

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/elliotchance/orderedmap"

	datamodel "example/build/client/clientset/versioned"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const DEFAULT_KEY = "default"
const DISPLAY_NAME_LABEL = "nexus/display_name"
const IS_NAME_HASHED_LABEL = "nexus/is_name_hashed"

func GetCRDParentsMap() map[string][]string {
	return map[string][]string{
		"configs.config.example.com":     {"roots.root.example.com", "tenants.tenant.example.com"},
		"interests.interest.example.com": {"roots.root.example.com", "tenants.tenant.example.com"},
		"roots.root.example.com":         {},
		"tenants.tenant.example.com":     {"roots.root.example.com"},
		"users.user.example.com":         {"roots.root.example.com", "tenants.tenant.example.com", "configs.config.example.com"},
		"wannas.wanna.example.com":       {"roots.root.example.com", "tenants.tenant.example.com", "configs.config.example.com", "users.user.example.com"},
	}
}

func GetObjectByCRDName(dmClient *datamodel.Clientset, crdName string, name string) interface{} {
	if crdName == "configs.config.example.com" {
		obj, err := dmClient.ConfigExampleV1().Configs().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "interests.interest.example.com" {
		obj, err := dmClient.InterestExampleV1().Interests().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "roots.root.example.com" {
		obj, err := dmClient.RootExampleV1().Roots().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "tenants.tenant.example.com" {
		obj, err := dmClient.TenantExampleV1().Tenants().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "users.user.example.com" {
		obj, err := dmClient.UserExampleV1().Users().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "wannas.wanna.example.com" {
		obj, err := dmClient.WannaExampleV1().Wannas().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}

	return nil
}

func ParseCRDLabels(crdName string, labels map[string]string) *orderedmap.OrderedMap {
	parents := GetCRDParentsMap()[crdName]

	m := orderedmap.NewOrderedMap()
	for _, parent := range parents {
		if label, ok := labels[parent]; ok {
			m.Set(parent, label)
		} else {
			m.Set(parent, DEFAULT_KEY)
		}
	}

	return m
}

func GetHashedName(crdName string, labels map[string]string, name string) string {
	orderedLabels := ParseCRDLabels(crdName, labels)

	var output string
	for i, key := range orderedLabels.Keys() {
		value, _ := orderedLabels.Get(key)

		output += fmt.Sprintf("%s:%s", key, value)
		if i < orderedLabels.Len()-1 {
			output += "/"
		}
	}

	output += fmt.Sprintf("%s:%s", crdName, name)

	h := sha1.New()
	_, _ = h.Write([]byte(output))
	return hex.EncodeToString(h.Sum(nil))
}
