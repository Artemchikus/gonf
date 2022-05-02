package models

import "fyne.io/fyne/v2"

type Resources struct {
	Pod                     fyne.CanvasObject
	Service                 fyne.CanvasObject
	Node                    fyne.CanvasObject
	ConfigMap               fyne.CanvasObject
	ReplicaSet              fyne.CanvasObject
	ReplicationController   fyne.CanvasObject
	Deployment              fyne.CanvasObject
	PersistenVolume         fyne.CanvasObject
	Namespace               fyne.CanvasObject
	Role                    fyne.CanvasObject
	HorizontalPodAutoscaler fyne.CanvasObject
	LimitRange              fyne.CanvasObject
	ResourceQuota           fyne.CanvasObject
	StorageClass            fyne.CanvasObject
	Secret                  fyne.CanvasObject
}
