package nexus_client

import (
	"fmt"
	"sync"
	"time"

	"github.com/elliotchance/orderedmap"
	"github.com/mitchellh/hashstructure"
	cache "k8s.io/client-go/tools/cache"
)

// subscriptionMap will store crd string as key and value as subscription type,
// for example key="roots.orgchart.vmware.org" and value=subscription{}
var subscriptionMap = sync.Map{}

type subscriptionIdKey struct {
	SubscriptionPath string
	SubscriptionTime string
}

// Hash of subscriptionIdKey object.
type SubscriptionId struct {
	Id   uint64
	Type string
}

// Type level subscription data.
type subscription struct {
	informer          cache.SharedIndexInformer
	stop              chan struct{}
	WriteCacheObjects *sync.Map

	typeSubscription *TypeSubscriptionInfo
	// Map of unique subscription paths for this type.
	// Key: subscription id.
	//      This is a unique id for a specific subsciption.
	//      This ID is application visible.
	//      App's can use this id to reference a specific subscription request.
	// Value:
	// 		Subscription information corresponding to a subscription request.
	subscriptions map[SubscriptionId]*SubscriptionInfo
}

func getSubscription(subsId SubscriptionId) *subscription {
	typeVal, ok := subscriptionMap.Load(subsId.Type)
	if !ok {
		return nil
	}
	return typeVal.(*subscription)
}

/* func (s *subscription) registerGlobalAddCallback(cbfn func(obj interface{})) error {
	_, err := s.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: cbfn})
	if err != nil {
		return fmt.Errorf("failed to register global add cbfn on informer with error %v", err)
	}
	return nil
}

func (s *subscription) registerGlobalUpdateCallback(subId SubscriptionId, cbfn func(oldObj, newObj interface{})) error {
	_, err := s.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{UpdateFunc: cbfn})
	if err != nil {
		return fmt.Errorf("failed to register update cbfn on informer with error %v", err)
	}
	return nil
}

func (s *subscription) registerGlobalDeleteCallback(subId SubscriptionId, cbfn func(obj interface{})) error {
	_, err := s.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{DeleteFunc: cbfn})
	if err != nil {
		return fmt.Errorf("failed to register delete cbfn on informer with error %v", err)
	}
	return nil
} */

func (s *subscription) getSubscriptionInfo(subsId SubscriptionId) *SubscriptionInfo {
	subsVal, ok := s.subscriptions[subsId]
	if !ok {
		return nil
	}
	return subsVal
}

func (s *subscription) RegisterAddCallback(subId SubscriptionId, cbfn func(obj interface{})) error {
	subInfo, ok := s.subscriptions[subId]
	if !ok {
		return fmt.Errorf("subscription Id %v does not exist", subId)
	}

	registrationId, err := s.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: cbfn})
	if err != nil {
		return fmt.Errorf("failed to register add cbfn on informer with error %v", err)
	}
	subInfo.AddCBRegistrationId = registrationId
	return nil
}

func (s *subscription) RegisterUpdateCallback(subId SubscriptionId, cbfn func(oldObj, newObj interface{})) error {
	subInfo, ok := s.subscriptions[subId]
	if !ok {
		return fmt.Errorf("subscription Id %v does not exist", subId)
	}

	registrationId, err := s.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{UpdateFunc: cbfn})
	if err != nil {
		return fmt.Errorf("failed to register update cbfn on informer with error %v", err)
	}
	subInfo.UpdateCBRegistrationId = registrationId
	return nil
}

func (s *subscription) RegisterDeleteCallback(subId SubscriptionId, cbfn func(obj interface{})) error {
	subInfo, ok := s.subscriptions[subId]
	if !ok {
		return fmt.Errorf("subscription Id %v does not exist", subId)
	}

	registrationId, err := s.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{DeleteFunc: cbfn})
	if err != nil {
		return fmt.Errorf("failed to register delete cbfn on informer with error %v", err)
	}
	subInfo.DeleteCBRegistrationId = registrationId
	return nil
}

func (s *subscription) IsSubscribed(subId SubscriptionId, labels map[string]string) bool {
	subInfo, ok := s.subscriptions[subId]
	if !ok {
		return false
	}

	for el := subInfo.SubscriptionPath.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Key, el.Value)

		if el.Value == "*" {
			// "*" will match everything.
			continue
		}

		// value can be *, default to match missing keys
		labelValue, ok := labels[el.Key.(string)]
		if !ok {
			if el.Value.(string) == "default" {
				continue
			} else {
				// Not a match
				return false
			}
		}

		if labelValue != el.Value.(string) {
			return false
		}
	}

	return true
}

type SubscriptionRegInfo struct {
	AddCBRegistrationId    cache.ResourceEventHandlerRegistration
	UpdateCBRegistrationId cache.ResourceEventHandlerRegistration
	DeleteCBRegistrationId cache.ResourceEventHandlerRegistration
}

type SubscriptionInfo struct {
	// subscription sub-graph
	SubscriptionPath *orderedmap.OrderedMap // CRD to specific object name mapping.

	SubscriptionRegInfo
}

type TypeSubscriptionInfo struct {
	addSubscriptionId SubscriptionId
	delSubscriptionId SubscriptionId
}

func (s SubscriptionInfo) RegisterAddCallback() error    { return nil }
func (s SubscriptionInfo) RegisterUpdateCallback() error { return nil }
func (s SubscriptionInfo) RegisterDeleteCallback() error { return nil }

func getSubscriptionPath(kv *orderedmap.OrderedMap) string {
	var path string
	for _, key := range kv.Keys() {
		value, _ := kv.Get(key)
		path += fmt.Sprintf("/%s/%s", key, value)
	}
	return path
}

func NewSubscriptionInfo(path *orderedmap.OrderedMap) SubscriptionInfo {
	return SubscriptionInfo{
		SubscriptionPath: path,
	}
}

// Registers a subscription to a given path of a type.
//
// This method creates a new subscription for a type, if one does not already exist.
// The subscription for the path is then appended to the type subscription.
// If a suscription for a type exists, this method will then append subscription for the path
// to the type subscription.
func subscribe(typeKey string, informer cache.SharedIndexInformer, pathKV *orderedmap.OrderedMap) (SubscriptionId, error) {
	var sub *subscription

	// Create a subscription if one already exists.
	val, ok := subscriptionMap.Load(typeKey)
	if !ok {
		sub = &subscription{
			informer:          informer,
			stop:              make(chan struct{}),
			WriteCacheObjects: &sync.Map{},
			subscriptions:     make(map[SubscriptionId]*SubscriptionInfo),
		}
		go sub.informer.Run(sub.stop)
		subscriptionMap.Store(typeKey, sub)
	} else {
		sub = val.(*subscription)
	}

	// Regsiter the subscription info the existing subscription for the type.

	// Generate a unique ID for this subscription request.
	subscriptionkey := subscriptionIdKey{
		SubscriptionPath: getSubscriptionPath(pathKV),
		SubscriptionTime: time.Now().String(),
	}
	subsHash, err := hashstructure.Hash(subscriptionkey, nil)
	if err != nil {
		return SubscriptionId{}, fmt.Errorf("subscribe failed to construct subscription id with error %v for type %s path %s", err, typeKey, subscriptionkey.SubscriptionPath)
	}

	// Append the path to type subscription handle.
	subsId := SubscriptionId{
		Id:   subsHash,
		Type: typeKey,
	}
	sub.subscriptions[subsId] = &SubscriptionInfo{
		SubscriptionPath: pathKV,
	}

	return subsId, nil
}

func (s SubscriptionId) UnSubscribe() error {

	sub := getSubscription(s)
	if sub == nil {
		return fmt.Errorf("subscription not found for subscriptionId %+v", s)
	}

	subsVal := sub.getSubscriptionInfo(s)
	if sub == nil {
		return fmt.Errorf("subscription info not found for subscriptionId %+v", s)
	}

	if subsVal.AddCBRegistrationId != nil {
		err := sub.informer.RemoveEventHandler(subsVal.AddCBRegistrationId)
		if err != nil {
			return fmt.Errorf("failed to remove add event handler for subscriptionID %+v", s)
		}
	}

	if subsVal.UpdateCBRegistrationId != nil {
		err := sub.informer.RemoveEventHandler(subsVal.UpdateCBRegistrationId)
		if err != nil {
			return fmt.Errorf("failed to remove update event handler for subscriptionID %+v", s)
		}
	}

	if subsVal.DeleteCBRegistrationId != nil {
		err := sub.informer.RemoveEventHandler(subsVal.DeleteCBRegistrationId)
		if err != nil {
			return fmt.Errorf("failed to remove delete event handler for subscriptionID %+v", s)
		}
	}

	delete(sub.subscriptions, s)

	// If there are no more subscriptions for this type, garbage collect suscritption
	// info for this type.
	if len(sub.subscriptions) == 0 && sub.typeSubscription == nil {
		subscriptionMap.Delete(s.Type)
	}

	return nil
}

func (s SubscriptionId) RegisterAddCallback(cbfn func(obj interface{})) error {
	sub := getSubscription(s)
	if sub == nil {
		return fmt.Errorf("subscription not found for subscriptionId %+v", s)
	}

	subInfo := sub.getSubscriptionInfo(s)
	if subInfo == nil {
		return fmt.Errorf("subscription info not found for subscriptionId %+v", s)
	}

	registrationId, err := sub.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: cbfn})
	if err != nil {
		return fmt.Errorf("failed to register add cbfn on informer with error %v", err)
	}
	subInfo.AddCBRegistrationId = registrationId
	return nil
}

func (s SubscriptionId) RegisterUpdateCallback(cbfn func(oldObj, newObj interface{})) error {
	sub := getSubscription(s)
	if sub == nil {
		return fmt.Errorf("subscription not found for subscriptionId %+v", s)
	}

	subInfo := sub.getSubscriptionInfo(s)
	if subInfo == nil {
		return fmt.Errorf("subscription info not found for subscriptionId %+v", s)
	}

	registrationId, err := sub.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{UpdateFunc: cbfn})
	if err != nil {
		return fmt.Errorf("failed to register update cbfn on informer with error %v", err)
	}

	subInfo.UpdateCBRegistrationId = registrationId
	return nil
}

func (s SubscriptionId) RegisterDeleteCallback(cbfn func(obj interface{})) error {
	sub := getSubscription(s)
	if sub == nil {
		return fmt.Errorf("subscription not found for subscriptionId %+v", s)
	}

	subInfo := sub.getSubscriptionInfo(s)
	if subInfo == nil {
		return fmt.Errorf("subscription info not found for subscriptionId %+v", s)
	}

	registrationId, err := sub.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{DeleteFunc: cbfn})
	if err != nil {
		return fmt.Errorf("failed to register delete cbfn on informer with error %v", err)
	}
	subInfo.DeleteCBRegistrationId = registrationId
	return nil
}

// Registers a subscription to a type.
//
// This method creates a new subscription for a type, if one does not already exist.
// The subscription for the path is then appended to the type subscription.
// If a suscription for a type exists, this method will then append subscription for the path
// to the type subscription.
func subscribeType(typeKey string, informer cache.SharedIndexInformer) (*subscription, error) {
	var sub *subscription

	// Create a subscription if one already exists.
	val, ok := subscriptionMap.Load(typeKey)
	if !ok {
		sub = &subscription{
			informer:          informer,
			stop:              make(chan struct{}),
			WriteCacheObjects: &sync.Map{},
			subscriptions:     make(map[SubscriptionId]*SubscriptionInfo),
		}
		go sub.informer.Run(sub.stop)
		subscriptionMap.Store(typeKey, sub)
	} else {
		sub = val.(*subscription)
	}

	if sub.typeSubscription != nil {
		return sub, fmt.Errorf("type subscription exists for type %v", typeKey)
	}
	sub.typeSubscription = &TypeSubscriptionInfo{}
	return sub, nil
}

// Registers a subscription to a type.
//
// This method creates a new subscription for a type, if one does not already exist.
// The subscription for the path is then appended to the type subscription.
// If a suscription for a type exists, this method will then append subscription for the path
// to the type subscription.
func unSubscribeType(typeKey string) {
	var sub *subscription
	var ok bool

	val, ok := subscriptionMap.Load(typeKey)
	if !ok {
		return
	} else {
		sub = val.(*subscription)
	}

	// If type subscription is current active, remove it.
	if sub.typeSubscription != nil {
		sub.typeSubscription = nil
	}

	// If there are other selective subscriptions active,
	// then informer cannot be stopped at the moment.
	if len(sub.subscriptions) > 0 {
		return
	}

	// No more active subscriptions exists.
	// Stop the informer and remove type from subscription map.
	close(sub.stop)
	subscriptionMap.Delete(typeKey)

	return
}

func createSubscription(typeKey string, informer cache.SharedIndexInformer) *subscription {
	var sub *subscription

	// Create a subscription if one already exists.
	val, ok := subscriptionMap.Load(typeKey)
	if !ok {
		sub = &subscription{
			informer:          informer,
			stop:              make(chan struct{}),
			WriteCacheObjects: &sync.Map{},
			subscriptions:     make(map[SubscriptionId]*SubscriptionInfo),
		}
		go sub.informer.Run(sub.stop)
		subscriptionMap.Store(typeKey, sub)
	} else {
		sub = val.(*subscription)
	}
	return sub
}
