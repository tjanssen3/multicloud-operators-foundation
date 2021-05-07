// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/openshift/hive/apis/hive/v1"
	scheme "github.com/openshift/hive/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterDeploymentsGetter has a method to return a ClusterDeploymentInterface.
// A group's client should implement this interface.
type ClusterDeploymentsGetter interface {
	ClusterDeployments(namespace string) ClusterDeploymentInterface
}

// ClusterDeploymentInterface has methods to work with ClusterDeployment resources.
type ClusterDeploymentInterface interface {
	Create(ctx context.Context, clusterDeployment *v1.ClusterDeployment, opts metav1.CreateOptions) (*v1.ClusterDeployment, error)
	Update(ctx context.Context, clusterDeployment *v1.ClusterDeployment, opts metav1.UpdateOptions) (*v1.ClusterDeployment, error)
	UpdateStatus(ctx context.Context, clusterDeployment *v1.ClusterDeployment, opts metav1.UpdateOptions) (*v1.ClusterDeployment, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ClusterDeployment, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ClusterDeploymentList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterDeployment, err error)
	ClusterDeploymentExpansion
}

// clusterDeployments implements ClusterDeploymentInterface
type clusterDeployments struct {
	client rest.Interface
	ns     string
}

// newClusterDeployments returns a ClusterDeployments
func newClusterDeployments(c *HiveV1Client, namespace string) *clusterDeployments {
	return &clusterDeployments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterDeployment, and returns the corresponding clusterDeployment object, and an error if there is any.
func (c *clusterDeployments) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ClusterDeployment, err error) {
	result = &v1.ClusterDeployment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterdeployments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterDeployments that match those selectors.
func (c *clusterDeployments) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ClusterDeploymentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ClusterDeploymentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterdeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterDeployments.
func (c *clusterDeployments) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clusterdeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterDeployment and creates it.  Returns the server's representation of the clusterDeployment, and an error, if there is any.
func (c *clusterDeployments) Create(ctx context.Context, clusterDeployment *v1.ClusterDeployment, opts metav1.CreateOptions) (result *v1.ClusterDeployment, err error) {
	result = &v1.ClusterDeployment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clusterdeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterDeployment).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterDeployment and updates it. Returns the server's representation of the clusterDeployment, and an error, if there is any.
func (c *clusterDeployments) Update(ctx context.Context, clusterDeployment *v1.ClusterDeployment, opts metav1.UpdateOptions) (result *v1.ClusterDeployment, err error) {
	result = &v1.ClusterDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterdeployments").
		Name(clusterDeployment.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterDeployment).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterDeployments) UpdateStatus(ctx context.Context, clusterDeployment *v1.ClusterDeployment, opts metav1.UpdateOptions) (result *v1.ClusterDeployment, err error) {
	result = &v1.ClusterDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterdeployments").
		Name(clusterDeployment.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterDeployment).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterDeployment and deletes it. Returns an error if one occurs.
func (c *clusterDeployments) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterdeployments").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterDeployments) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterdeployments").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterDeployment.
func (c *clusterDeployments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterDeployment, err error) {
	result = &v1.ClusterDeployment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clusterdeployments").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
