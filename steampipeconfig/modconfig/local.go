package modconfig

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
)

// Local is a struct representing a Local resource
type Local struct {
	ShortName string
	FullName  string `cty:"name"`

	// list of all block referenced by the resource
	References []string `column:"refs,jsonb"`

	Value     cty.Value
	DeclRange hcl.Range

	metadata *ResourceMetadata
}

func NewLocal(name string, val cty.Value, declRange hcl.Range) *Local {
	return &Local{
		ShortName: name,
		FullName:  fmt.Sprintf("local.%s", name),
		Value:     val,
		DeclRange: declRange,
	}
}

// Name implements HclResource, ResourceWithMetadata
func (l *Local) Name() string {
	return l.FullName
}

// GetMetadata implements ResourceWithMetadata
func (l *Local) GetMetadata() *ResourceMetadata {
	return l.metadata
}

// SetMetadata implements ResourceWithMetadata
func (l *Local) SetMetadata(metadata *ResourceMetadata) {
	l.metadata = metadata
}

// OnDecoded implements HclResource
func (l *Local) OnDecoded(*hcl.Block) {}

func (l *Local) AddReference(reference string) {
	l.References = append(l.References, reference)
}

// CtyValue implements HclResource
func (l *Local) CtyValue() (cty.Value, error) {
	return l.Value, nil
}
