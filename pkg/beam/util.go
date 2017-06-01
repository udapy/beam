package beam

// We have some freedom to create various utilities, users can use depending on
// preferences. One point of keeping Pipeline transformation functions plain Go
// functions is that such utilities are more readily possible.

// For example, we can have an "easyio" package that selects a textio, gcsio,
// awsio, etc. transformation based on the filename schema. Such wrappers would
// look exactly like the more primitive sources/sinks, but be picked at
// pipeline construction time.

// Seq is a convenience helper to chain single-input/single-output ParDos together
// in a sequence.
func Seq(p *Pipeline, col PCollection, dofns ...interface{}) PCollection {
	cur := col
	for _, dofn := range dofns {
		cur = ParDo(p, dofn, cur)
	}
	return cur
}

// TODO(herohde) 5/30/2017: add composite helper that picks up the enclosing function name.

// Composite is a helper to scope a composite transform.
func Composite(p *Pipeline, name string, fn func(pipeline *Pipeline) PCollection) PCollection {
	return fn(p.Composite(name))
}

// The MustX functions are convenience helpers to create error-less functions.

// MustN returns the input, but panics if err != nil.
func MustN(list []PCollection, err error) []PCollection {
	if err != nil {
		panic(err)
	}
	return list
}

// Must returns the input, but panics if err != nil.
func Must(a PCollection, err error) PCollection {
	if err != nil {
		panic(err)
	}
	return a
}

// Must2 returns the input, but panics if err != nil.
func Must2(a, b PCollection, err error) (PCollection, PCollection) {
	if err != nil {
		panic(err)
	}
	return a, b
}
