package functionres

import (
	"context"

	nuclioio "github.com/nuclio/nuclio/pkg/platform/kube/apis/nuclio.io/v1beta1"

	apps_v1beta1 "k8s.io/api/apps/v1beta1"
	autos_v1 "k8s.io/api/autoscaling/v1"
	"k8s.io/api/core/v1"
	ext_v1beta1 "k8s.io/api/extensions/v1beta1"
)

const (
	containerHTTPPort     = 8080
	healthCheckHTTPPort   = 8082
	containerHTTPPortName = "http"
)

type Client interface {

	// List returns the current existing function resources, populating the "deployment" sub-resource
	List(context.Context, string) ([]Resources, error)

	// Get returns the resources named by name, populating the "deployment" sub-resource
	Get(context.Context, string, string) (Resources, error)

	// CreateOrUpdate creates or updates exisisting resources
	CreateOrUpdate(context.Context, *nuclioio.Function, string) (Resources, error)

	// WaitAvailable waits until the resources are ready
	WaitAvailable(context.Context, string, string) error

	// Delete deletes resources
	Delete(context.Context, string, string) error
}

// Resources holds the resources a functionres holds
type Resources interface {

	// Deployment returns the deployment
	Deployment() (*apps_v1beta1.Deployment, error)

	// ConfigMap returns the configmap
	ConfigMap() (*v1.ConfigMap, error)

	// Service returns the service
	Service() (*v1.Service, error)

	// HorizontalPodAutoscaler returns the hpa
	HorizontalPodAutoscaler() (*autos_v1.HorizontalPodAutoscaler, error)

	// Ingress returns the ingress
	Ingress() (*ext_v1beta1.Ingress, error)
}
