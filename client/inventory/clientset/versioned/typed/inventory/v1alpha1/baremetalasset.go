// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	scheme "github.com/stolostron/cluster-lifecycle-api/client/inventory/clientset/versioned/scheme"
	v1alpha1 "github.com/stolostron/cluster-lifecycle-api/inventory/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BareMetalAssetsGetter has a method to return a BareMetalAssetInterface.
// A group's client should implement this interface.
type BareMetalAssetsGetter interface {
	BareMetalAssets(namespace string) BareMetalAssetInterface
}

// BareMetalAssetInterface has methods to work with BareMetalAsset resources.
type BareMetalAssetInterface interface {
	Create(ctx context.Context, bareMetalAsset *v1alpha1.BareMetalAsset, opts v1.CreateOptions) (*v1alpha1.BareMetalAsset, error)
	Update(ctx context.Context, bareMetalAsset *v1alpha1.BareMetalAsset, opts v1.UpdateOptions) (*v1alpha1.BareMetalAsset, error)
	UpdateStatus(ctx context.Context, bareMetalAsset *v1alpha1.BareMetalAsset, opts v1.UpdateOptions) (*v1alpha1.BareMetalAsset, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.BareMetalAsset, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.BareMetalAssetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BareMetalAsset, err error)
	BareMetalAssetExpansion
}

// bareMetalAssets implements BareMetalAssetInterface
type bareMetalAssets struct {
	client rest.Interface
	ns     string
}

// newBareMetalAssets returns a BareMetalAssets
func newBareMetalAssets(c *InventroyV1alpha1Client, namespace string) *bareMetalAssets {
	return &bareMetalAssets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bareMetalAsset, and returns the corresponding bareMetalAsset object, and an error if there is any.
func (c *bareMetalAssets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BareMetalAsset, err error) {
	result = &v1alpha1.BareMetalAsset{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("baremetalassets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BareMetalAssets that match those selectors.
func (c *bareMetalAssets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BareMetalAssetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BareMetalAssetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("baremetalassets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bareMetalAssets.
func (c *bareMetalAssets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("baremetalassets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a bareMetalAsset and creates it.  Returns the server's representation of the bareMetalAsset, and an error, if there is any.
func (c *bareMetalAssets) Create(ctx context.Context, bareMetalAsset *v1alpha1.BareMetalAsset, opts v1.CreateOptions) (result *v1alpha1.BareMetalAsset, err error) {
	result = &v1alpha1.BareMetalAsset{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("baremetalassets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bareMetalAsset).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a bareMetalAsset and updates it. Returns the server's representation of the bareMetalAsset, and an error, if there is any.
func (c *bareMetalAssets) Update(ctx context.Context, bareMetalAsset *v1alpha1.BareMetalAsset, opts v1.UpdateOptions) (result *v1alpha1.BareMetalAsset, err error) {
	result = &v1alpha1.BareMetalAsset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("baremetalassets").
		Name(bareMetalAsset.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bareMetalAsset).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *bareMetalAssets) UpdateStatus(ctx context.Context, bareMetalAsset *v1alpha1.BareMetalAsset, opts v1.UpdateOptions) (result *v1alpha1.BareMetalAsset, err error) {
	result = &v1alpha1.BareMetalAsset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("baremetalassets").
		Name(bareMetalAsset.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bareMetalAsset).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the bareMetalAsset and deletes it. Returns an error if one occurs.
func (c *bareMetalAssets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("baremetalassets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bareMetalAssets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("baremetalassets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched bareMetalAsset.
func (c *bareMetalAssets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BareMetalAsset, err error) {
	result = &v1alpha1.BareMetalAsset{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("baremetalassets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
