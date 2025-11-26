/*
Copyright 2025.

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

package controller

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	autofixv1 "github.com/boqier/AIOpsAnalyze/api/v1"
)

// AIOpsAnalyzerReconciler reconciles a AIOpsAnalyzer object
type AIOpsAnalyzerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// GetTargetPods 根据TargetSelector获取对应的Pod列表
func (r *AIOpsAnalyzerReconciler) GetTargetPods(ctx context.Context, target *autofixv1.TargetSelector) ([]corev1.Pod, error) {
	log := log.FromContext(ctx)

	// 处理命名空间
	namespace := target.Namespace
	if namespace == "" {
		namespace = corev1.NamespaceDefault
		log.V(1).Info("未指定命名空间，使用默认命名空间", "namespace", namespace)
	}

	// 创建 ListOptions
	listOptions := &client.ListOptions{
		Namespace: namespace,
	}

	// 正确处理 LabelSelector（关键修复！）
	if target.Selector.MatchLabels != nil || target.Selector.MatchExpressions != nil {
		selector, err := metav1.LabelSelectorAsSelector(&target.Selector)
		if err != nil {
			log.Error(err, "无法将 LabelSelector 转换为 Selector", "selector", target.Selector)
			return nil, err
		}
		listOptions.LabelSelector = selector
		log.V(1).Info("应用标签选择器", "selector", selector.String())
	} else {
		log.V(1).Info("未配置标签选择器，将获取命名空间内所有 Pod")
	}

	// 执行列表查询
	var pods corev1.PodList
	if err := r.List(ctx, &pods, listOptions); err != nil {
		log.Error(err, "获取Pod列表失败", "namespace", namespace, "selector", target.Selector)
		return nil, err
	}

	log.Info("成功获取目标Pod", "count", len(pods.Items), "namespace", namespace, "selector", target.Selector)
	return pods.Items, nil
}

// BuildLabelSelector 根据标签构建LabelSelector
func BuildLabelSelector(labels map[string]string) (*metav1.LabelSelector, error) {
	matchLabels := make(map[string]string)
	for k, v := range labels {
		matchLabels[k] = v
	}

	return &metav1.LabelSelector{
		MatchLabels: matchLabels,
	}, nil
}

// +kubebuilder:rbac:groups=autofix.aiops.com,resources=aiopsanalyzers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=autofix.aiops.com,resources=aiopsanalyzers/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=autofix.aiops.com,resources=aiopsanalyzers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the AIOpsAnalyzer object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.1/pkg/reconcile
func (r *AIOpsAnalyzerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AIOpsAnalyzerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&autofixv1.AIOpsAnalyzer{}).
		Named("aiopsanalyzer").
		Complete(r)
}
