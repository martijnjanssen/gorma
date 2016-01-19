package dsl_test

import (
	"testing"

	"github.com/bketelsen/gorma"
	gdsl "github.com/bketelsen/gorma/dsl"
	"github.com/raphael/goa/design"
	"github.com/raphael/goa/design/dsl"
)

func TestStorageGroup(t *testing.T) {

	var sg = gdsl.StorageGroup("MyStorageGroup", func() {})
	design := design.Design
	sd, ok := design.Constructs["gorma"][gorma.StorageGroup].(*gorma.StorageGroupDefinition)
	if !ok {
		t.Errorf("expected %#v to be %#v ", sd, sg)
	}

}

func TestStorageGroupChildren(t *testing.T) {

	var sg = gdsl.StorageGroup("MyStorageGroup", func() {
		gdsl.RelationalStore("mysql", func() {
		})
	})
	des := design.Design
	dsl.RunDSL()
	sd, ok := des.Constructs["gorma"][gorma.StorageGroup].(*gorma.StorageGroupDefinition)
	if !ok {
		t.Errorf("expected %#v to be %#v ", sd, sg)
	}
	if len(sd.RelationalStores) != 1 {
		t.Errorf("expected %d relational store, got %d", 1, len(sd.RelationalStores))
	}

}
