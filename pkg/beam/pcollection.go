package beam

import (
	"fmt"
	"github.com/apache/beam/sdks/go/pkg/beam/graph"
	"github.com/apache/beam/sdks/go/pkg/beam/graph/typex"
)

// PCollection is an immutable collection of values of type 'A', which must be
// a concrete Windowed Value type, such as W<int> or W<KV<int,string>>. A
// PCollection can contain either a bounded or unbounded number of elements.
// Bounded and unbounded PCollections are produced as the output of PTransforms
// (including root PTransforms like textio.Read), and can be passed as the
// inputs of other PTransforms. Some root transforms produce bounded
// PCollections and others produce unbounded ones.
//
// Each element in a PCollection has an associated timestamp. Sources assign
// timestamps to elements when they create PCollections, and other PTransforms
// propagate these timestamps from their input to their output implicitly or
// explicitly.
//
// Additionally, each element is assigned to a set of windows. By default, all
// elements are assigned into a single default window, GlobalWindow.
type PCollection struct {
	// n is the graph node that PCollection wraps. If there is no node, the
	// PCollection is invalid.
	n *graph.Node
}

// IsValid returns true iff the PCollection is valid and part of a Pipeline.
// Any use of an invalid PCollection will result in a panic.
func (p PCollection) IsValid() bool {
	return p.n != nil
}

// TODO(herohde) 5/30/2017: add name for PCollections? Java supports it.
// TODO(herohde) 5/30/2017: add windowing strategy and documentation.

// Type returns the full type 'A' of the elements. 'A' must be a concrete
// Windowed Value type, such as W<int> or W<KV<int,string>>.
func (p PCollection) Type() typex.FullType {
	if !p.IsValid() {
		panic("Invalid PCollection")
	}
	return p.n.Type()
}

// Coder returns the coder for the collection. The Coder is of type 'A'.
func (p PCollection) Coder() Coder {
	if !p.IsValid() {
		panic("Invalid PCollection")
	}
	return Coder{p.n.Coder}
}

// SetCoder set the coder for the collection. The Coder must be of type 'A'.
func (p PCollection) SetCoder(c Coder) error {
	if !p.IsValid() {
		panic("Invalid PCollection")
	}

	if !typex.IsEqual(p.n.Type(), c.coder.T) {
		return fmt.Errorf("coder type %v must be identical to node type %v", c.coder.T, p.n)
	}
	p.n.Coder = c.coder
	return nil
}

func (p PCollection) String() string {
	if !p.IsValid() {
		return "(invalid)"
	}
	return p.n.String()
}
