package libk8s

import (
	"github.com/robertkrimen/otto"
	"k8s.io/client-go/kubernetes"
)

// ClientSet models the Kubernetes Interface interface.
//
// It provides access to all of the defined resource groups.
//
// In JavaScript, you get one of these using `kubernetes.withNS(ns)`.
type ClientSet struct {
	GetNS func() otto.Value `otto:"getNS"`

	CoreV1            *CoreV1            `otto:"coreV1,alias=core"`
	AppsV1beta1       *AppsV1beta1       `otto:"appsV1beta1,alias=apps"`
	BatchV1           *BatchV1           `otto:"batchV1,alias=batch"`
	ExtensionsV1beta1 *ExtensionsV1beta1 `otto:"extensionsV1beta1,alias=extensions"`

	/* To be implemented:
	BatchV2alpha1         otto.Value `otto:"batchV2alpha1"`
	AuthenticationV1beta1 otto.Value `otto:"authenticationV1beta1,alias=authentication"`
	AuthorizationV1beta1  otto.Value `otto:"authorizationV1beta1,alias=authorization"`
	AutoscalingV1         otto.Value `otto:"autoscalingV1,alias=autoscaling"`
	CertificatesV1alpha1  otto.Value `otto:"certificatesV1alpha1,alias=certificates"`
	PolicayV1beta1        otto.Value `otto:"policyV1beta1,alias=policy"`
	RbacV1alpha1          otto.Value `otto:"rbacV1alpha1,alias=rbac"`
	StorageV1beta1        otto.Value `otto:"storageV1beta1,alias=storage"`
	*/
}

func NewClientSet(vm *otto.Otto, c kubernetes.Interface, ns string) *ClientSet {
	return &ClientSet{
		GetNS: func() otto.Value { r, _ := vm.ToValue(ns); return r },
		CoreV1: &CoreV1{
			Namespace:             NewNamespace(c, vm),
			Pod:                   NewPod(c, vm, ns),
			Secret:                NewSecret(c, vm, ns),
			Service:               NewService(c, vm, ns),
			ServiceAccount:        NewServiceAccount(c, vm, ns),
			PersistentVolumeClaim: NewPersistentVolumeClaim(c, vm, ns),
			ReplicationController: NewReplicationController(c, vm, ns),
		},
		AppsV1beta1: &AppsV1beta1{
			StatefulSet: NewStatefulSet(c, vm, ns),
		},
		BatchV1: &BatchV1{
			Job: NewJob(c, vm, ns),
		},
		ExtensionsV1beta1: &ExtensionsV1beta1{
			ConfigMap:          NewConfigMap(c, vm, ns),
			Deployment:         NewDeployment(c, vm, ns),
			ReplicaSet:         NewReplicaSet(c, vm, ns),
			DaemonSet:          NewDaemonSet(c, vm, ns),
			Ingress:            NewIngress(c, vm, ns),
			PodSecurityPolicy:  NewPodSecurityPolicy(c, vm),
			ThirdPartyResource: NewThirdPartyResource(c, vm),
		},
	}
}

type CoreV1 struct {
	Namespace             *Namespace             `otto:"namespace"`
	Pod                   *Pod                   `otto:"pod"`
	Secret                *Secret                `otto:"secret"`
	Service               *Service               `otto:"service"`
	ServiceAccount        *ServiceAccount        `otto:"serviceaccount"`
	PersistentVolumeClaim *PersistentVolumeClaim `otto:"persistentvolumeclaim"`
	ReplicationController *ReplicationController `otto:"replicationcontroller"`
}

type AppsV1beta1 struct {
	StatefulSet *StatefulSet `otto:"statefulset"`
}

type BatchV1 struct {
	Job *Job `otto:"job"`
}

type ExtensionsV1beta1 struct {
	ConfigMap          *ConfigMap          `otto:"configmap"`
	Deployment         *Deployment         `otto:"deployment"`
	ReplicaSet         *ReplicaSet         `otto:"replicaset"`
	DaemonSet          *DaemonSet          `otto:"daemonset"`
	Ingress            *Ingress            `otto:"ingress"`
	PodSecurityPolicy  *PodSecurityPolicy  `otto:"podsecuritypolicy"`
	ThirdPartyResource *ThirdPartyResource `otto:"thirdpartyresource"`
}
