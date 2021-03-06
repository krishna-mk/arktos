/*
Copyright The Kubernetes Authors.
Copyright 2020 Authors of Arktos - file modified.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	strings "strings"
	"time"

	v1alpha1 "k8s.io/api/auditregistration/v1alpha1"
	errors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
	klog "k8s.io/klog"
)

// AuditSinksGetter has a method to return a AuditSinkInterface.
// A group's client should implement this interface.
type AuditSinksGetter interface {
	AuditSinks() AuditSinkInterface
	AuditSinksWithMultiTenancy(tenant string) AuditSinkInterface
}

// AuditSinkInterface has methods to work with AuditSink resources.
type AuditSinkInterface interface {
	Create(*v1alpha1.AuditSink) (*v1alpha1.AuditSink, error)
	Update(*v1alpha1.AuditSink) (*v1alpha1.AuditSink, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AuditSink, error)
	List(opts v1.ListOptions) (*v1alpha1.AuditSinkList, error)
	Watch(opts v1.ListOptions) watch.AggregatedWatchInterface
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AuditSink, err error)
	AuditSinkExpansion
}

// auditSinks implements AuditSinkInterface
type auditSinks struct {
	client  rest.Interface
	clients []rest.Interface
	te      string
}

// newAuditSinks returns a AuditSinks
func newAuditSinks(c *AuditregistrationV1alpha1Client) *auditSinks {
	return newAuditSinksWithMultiTenancy(c, "system")
}

func newAuditSinksWithMultiTenancy(c *AuditregistrationV1alpha1Client, tenant string) *auditSinks {
	return &auditSinks{
		client:  c.RESTClient(),
		clients: c.RESTClients(),
		te:      tenant,
	}
}

// Get takes name of the auditSink, and returns the corresponding auditSink object, and an error if there is any.
func (c *auditSinks) Get(name string, options v1.GetOptions) (result *v1alpha1.AuditSink, err error) {
	result = &v1alpha1.AuditSink{}
	err = c.client.Get().
		Tenant(c.te).
		Resource("auditsinks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)

	return
}

// List takes label and field selectors, and returns the list of AuditSinks that match those selectors.
func (c *auditSinks) List(opts v1.ListOptions) (result *v1alpha1.AuditSinkList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AuditSinkList{}
	err = c.client.Get().
		Tenant(c.te).
		Resource("auditsinks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	if err == nil {
		return
	}

	if !(errors.IsForbidden(err) && strings.Contains(err.Error(), "no relationship found between node")) {
		return
	}

	// Found api server that works with this list, keep the client
	for _, client := range c.clients {
		if client == c.client {
			continue
		}

		err = client.Get().
			Tenant(c.te).
			Resource("auditsinks").
			VersionedParams(&opts, scheme.ParameterCodec).
			Timeout(timeout).
			Do().
			Into(result)

		if err == nil {
			c.client = client
			return
		}

		if err != nil && errors.IsForbidden(err) &&
			strings.Contains(err.Error(), "no relationship found between node") {
			klog.V(6).Infof("Skip error %v in list", err)
			continue
		}
	}

	return
}

// Watch returns a watch.Interface that watches the requested auditSinks.
func (c *auditSinks) Watch(opts v1.ListOptions) watch.AggregatedWatchInterface {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	aggWatch := watch.NewAggregatedWatcher()
	for _, client := range c.clients {
		watcher, err := client.Get().
			Tenant(c.te).
			Resource("auditsinks").
			VersionedParams(&opts, scheme.ParameterCodec).
			Timeout(timeout).
			Watch()
		if err != nil && opts.AllowPartialWatch && errors.IsForbidden(err) {
			// watch error was not returned properly in error message. Skip when partial watch is allowed
			klog.V(6).Infof("Watch error for partial watch %v. options [%+v]", err, opts)
			continue
		}
		aggWatch.AddWatchInterface(watcher, err)
	}
	return aggWatch
}

// Create takes the representation of a auditSink and creates it.  Returns the server's representation of the auditSink, and an error, if there is any.
func (c *auditSinks) Create(auditSink *v1alpha1.AuditSink) (result *v1alpha1.AuditSink, err error) {
	result = &v1alpha1.AuditSink{}

	objectTenant := auditSink.ObjectMeta.Tenant
	if objectTenant == "" {
		objectTenant = c.te
	}

	err = c.client.Post().
		Tenant(objectTenant).
		Resource("auditsinks").
		Body(auditSink).
		Do().
		Into(result)

	return
}

// Update takes the representation of a auditSink and updates it. Returns the server's representation of the auditSink, and an error, if there is any.
func (c *auditSinks) Update(auditSink *v1alpha1.AuditSink) (result *v1alpha1.AuditSink, err error) {
	result = &v1alpha1.AuditSink{}

	objectTenant := auditSink.ObjectMeta.Tenant
	if objectTenant == "" {
		objectTenant = c.te
	}

	err = c.client.Put().
		Tenant(objectTenant).
		Resource("auditsinks").
		Name(auditSink.Name).
		Body(auditSink).
		Do().
		Into(result)

	return
}

// Delete takes name of the auditSink and deletes it. Returns an error if one occurs.
func (c *auditSinks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Tenant(c.te).
		Resource("auditsinks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *auditSinks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Tenant(c.te).
		Resource("auditsinks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched auditSink.
func (c *auditSinks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AuditSink, err error) {
	result = &v1alpha1.AuditSink{}
	err = c.client.Patch(pt).
		Tenant(c.te).
		Resource("auditsinks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)

	return
}
