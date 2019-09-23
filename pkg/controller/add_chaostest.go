package controller

import (
	"github.com/jpapejr/chaos-operator/pkg/controller/chaostest"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, chaostest.Add)
}
