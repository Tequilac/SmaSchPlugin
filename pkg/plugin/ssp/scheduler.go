package ssp

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

// Name ... the custom shceduler name
const Name = "ssp-scheduler"

// CustomScheduler ... The type CustomScheduler implement the interface of the kube-scheduler framework
type CustomScheduler struct {
	handle framework.Handle
}

// Let the type CustomScheduler implement the ScorePlugin interface
var _ framework.ScorePlugin = &CustomScheduler{}

// Name ... Implement Plugin interface Name() @pkg/scheduler/framework/v1alpha1/interface.go
func (*CustomScheduler) Name() string {
	return Name
}

func (s *CustomScheduler) Score(ctx context.Context, _ *framework.CycleState, _ *v1.Pod, nodeName string) (int64, *framework.Status) {
	node, err := s.handle.ClientSet().CoreV1().Nodes().Get(ctx, nodeName, metav1.GetOptions{})
	if err != nil {
		return 0, nil
	}
	var labels = node.Labels

	if labels["sma-freeze"] == "sma-freeze" {
		return 0, nil
	}

	var sum int64 = 0
	switch labels["sma-mem"] {
	case "", "sma-mem-low":
		sum += 0
	case "sma-mem-mid":
		sum += 2
	case "sma-mem-high":
		sum += 4
	}
	switch labels["sma-cpu"] {
	case "", "sma-cpu-low":
		sum += 0
	case "sma-cpu-mid":
		sum += 3
	case "sma-cpu-high":
		sum += 5
	}
	switch labels["sma-temp"] {
	case "", "sma-temp-low":
		sum += 5
	case "sma-temp-mid":
		sum += 3
	case "sma-temp-high":
		sum += 0
	}
	return sum, nil
}

// ScoreExtensions ...
func (*CustomScheduler) ScoreExtensions() framework.ScoreExtensions {
	return nil
}

// New ... Create a scheduler instance
// New() is type PluginFactory = func(configuration runtime.Object, f v1alpha1.FrameworkHandle) (v1alpha1.Plugin, error)
// mentioned in https://github.com/kubernetes/kubernetes/blob/master/pkg/scheduler/framework/runtime/registry.go
func New(_ runtime.Object, f framework.Handle) (framework.Plugin, error) {
	return &CustomScheduler{
		handle: f,
	}, nil
}
