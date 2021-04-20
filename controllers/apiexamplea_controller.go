/*
Copyright 2021 cwr.

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

package controllers

import (
	"context"
	"log"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	groupav1beta1 "github.com/drcwr/kubeopdemo/api/v1beta1"
)

// ApiExampleAReconciler reconciles a ApiExampleA object
type ApiExampleAReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=groupa.k8s.demo.com,resources=apiexampleas,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=groupa.k8s.demo.com,resources=apiexampleas/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=groupa.k8s.demo.com,resources=apiexampleas/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ApiExampleA object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.2/pkg/reconcile
func (r *ApiExampleAReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	// your logic here
	ctx = context.Background()
	_ = r.Log.WithValues("apiexamplea", req.NamespacedName)

	obja := &groupav1beta1.ApiExampleA{}
	if err := r.Get(ctx, req.NamespacedName, obja); err != nil {
		log.Println(err, "nuable to fetch New Object")
	} else {
		log.Println("fetch New Object:", obja.Spec.AppName, obja.Spec.AppFunc)
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ApiExampleAReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&groupav1beta1.ApiExampleA{}).
		Complete(r)
}
