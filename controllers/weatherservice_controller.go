/*
Copyright 2022.

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
	"fmt"
	"strconv"

	v1alpha1 "github.com/singhiqbal1007/weather-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	errors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// WeatherServiceReconciler reconciles a WeatherService object
type WeatherServiceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=weatherservice.iqbal.com,resources=weatherservices,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=weatherservice.iqbal.com,resources=weatherservices/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=weatherservice.iqbal.com,resources=weatherservices/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the WeatherService object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.2/pkg/reconcile
func (r *WeatherServiceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	log.Info("Reconciling WeatherService", "req", req)

	instance := &v1alpha1.WeatherService{}
	err := r.Get(ctx, req.NamespacedName, instance)

	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}
	// Define a new Pod object
	pod := NewPod(instance)

	// Set instance as the owner and controller
	// This is required for garbage collection
	// pod will be deleted when the instance is deleted
	if err := controllerutil.SetControllerReference(instance, pod, r.Scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Pod already exists
	found := &corev1.Pod{}
	err = r.Get(ctx, types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, found)

	if err != nil && errors.IsNotFound(err) {
		log.Info("Creating a new Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
		err = r.Create(ctx, pod)
		if err != nil {
			return reconcile.Result{}, err
		}

		// Pod created successfully
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Pod already exists - don't requeue
	log.Info("Skip reconcile: Pod already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
	instance.Status.Executed = true
	err = r.Update(ctx, instance)
	if err != nil {
		log.Error(err, "failed to update WhaleSay status")
		return reconcile.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *WeatherServiceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.WeatherService{}).
		Complete(r)
}

func NewPod(cr *v1alpha1.WeatherService) *corev1.Pod {

	url := fmt.Sprintf("http://wttr.in/%s?%d", cr.Spec.City, cr.Spec.Days)

	labels := map[string]string{
		"app":  cr.Name,
		"city": cr.Spec.City,
		"days": strconv.Itoa(cr.Spec.Days),
	}
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "weather-report-" + cr.Spec.City,
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "weather-report-container" + cr.Spec.City,
					Image:   "alpine/curl",
					Command: []string{"sh", "-c", "curl -s " + url + " && sleep 3600"},
				},
			},
			RestartPolicy: corev1.RestartPolicy("Never"),
		},
	}
}
