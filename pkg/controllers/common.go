package controllers

import (
	"time"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	DefaultResync = 10 * time.Hour
)

func nameMapper(names []string) handler.ToRequestsFunc {
	nameSet := sets.NewString(names...)
	return func(obj handler.MapObject) []reconcile.Request {
		if !nameSet.Has(obj.Meta.GetName()) {
			return nil
		}
		return []reconcile.Request{
			{
				NamespacedName: types.NamespacedName{
					Namespace: obj.Meta.GetNamespace(),
					Name:      obj.Meta.GetName(),
				},
			},
		}
	}
}

func NamedResourceHandler(names ...string) handler.EventHandler {
	return &handler.EnqueueRequestsFromMapFunc{
		ToRequests: nameMapper(names),
	}
}
