package main

// Geometry consists of Points, Verticies and Primitives.
type Geometry struct {
	Pts   Points
	Vtxs  Verticies
	Prims Primitives
}

// NewGeometry creates new Geometry.
func NewGeometry() Geometry {
	return Geometry{
		Pts:   NewPoints(),
		Vtxs:  NewVerticies(),
		Prims: NewPrimitives(),
	}
}

// Points are foundation of geometry. They hold attributes.
// "P" attribute should be always exists, which is position of a point.
type Points map[string]Attributes

// NewPoints creates new Points.
func NewPoints() Points {
	return Points{
		"P": NewAttributes(Vector, 3),
	}
}

// Verticies are indice of points, but they also have their own attributes.
// "pt" attribute should be always exists, which is indice to a point.
type Verticies map[string]Attributes

// NewVerticies creates new Verticies.
func NewVerticies() Verticies {
	return Verticies{
		"pt": NewAttributes(Int, 1),
	}
}

// Primitives are consists of multiple verticies.
// "vtxs" attribute should be always exists, which is indice to verticies.
type Primitives map[string]Attributes

// NewPrimitives creates new Primitives.
func NewPrimitives() Primitives {
	return Primitives{
		"vtxs": NewAttributes(Int, 0),
	}
}

type AttributeType int

const (
	Unknown = AttributeType(iota)
	String
	Int
	Float
	Vector // float vector
	Matrix // float matrix
)

// Attributes are holder for any type of attribute.
// It should not be used to hold multiple attributes. ex) "P" and "N"
type Attributes struct {
	Type AttributeType
	// N is number of elements per a unit.
	// When it is non zero, len(Values) % N should be 0.
	// When it is 0, number of elements are various per unit.
	// See Ns then.
	N  uint
	Ns []uint

	// Values.
	// Only one of following will be used in a Attributes.
	Str []string
	Int []int
	Flt []float64
}

// NewAttributes creates new Attributes.
func NewAttributes(t AttributeType, n uint) Attributes {
	return Attributes{
		Type: String,
		N:    n,
	}
}
