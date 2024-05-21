package converter

import (
	"context"
	"testing"

	"github.com/crossplane/upjet/pkg/examples/conversion"

	"github.com/upbound/provider-azure/config"
)

func Test(t *testing.T) {
	pc, err := config.GetProvider(context.TODO(), false)
	if err != nil {
		panic(err)
	}

	if err = conversion.ConvertSingletonListToEmbeddedObject(pc,
		"../examples",
		"../hack/boilerplate.go.txt"); err != nil {
		panic(err)
	}
}
