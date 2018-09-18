// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// unnamed API: Application User Types
//
// Command:
// $ goagen
// --design=github.com/aya-eiya/goa-test/design/app
// --out=$(GOPATH)/src/github.com/aya-eiya/goa-test/apps
// --version=v1.3.1

package app

// myType user type.
type myType struct {
	// String Array
	Arrs []string `form:"arrs,omitempty" json:"arrs,omitempty" yaml:"arrs,omitempty" xml:"arrs,omitempty"`
	// non-required Integer
	IntVal *int `form:"intVal,omitempty" json:"intVal,omitempty" yaml:"intVal,omitempty" xml:"intVal,omitempty"`
}

// Publicize creates MyType from myType
func (ut *myType) Publicize() *MyType {
	var pub MyType
	if ut.Arrs != nil {
		pub.Arrs = ut.Arrs
	}
	if ut.IntVal != nil {
		pub.IntVal = ut.IntVal
	}
	return &pub
}

// MyType user type.
type MyType struct {
	// String Array
	Arrs []string `form:"arrs,omitempty" json:"arrs,omitempty" yaml:"arrs,omitempty" xml:"arrs,omitempty"`
	// non-required Integer
	IntVal *int `form:"intVal,omitempty" json:"intVal,omitempty" yaml:"intVal,omitempty" xml:"intVal,omitempty"`
}

// myTypes user type.
type myTypes struct {
	// MyType Array
	Arrs []*myType `form:"arrs,omitempty" json:"arrs,omitempty" yaml:"arrs,omitempty" xml:"arrs,omitempty"`
	// non-required Integer
	IntVal *int `form:"intVal,omitempty" json:"intVal,omitempty" yaml:"intVal,omitempty" xml:"intVal,omitempty"`
}

// Publicize creates MyTypes from myTypes
func (ut *myTypes) Publicize() *MyTypes {
	var pub MyTypes
	if ut.Arrs != nil {
		pub.Arrs = make([]*MyType, len(ut.Arrs))
		for i2, elem2 := range ut.Arrs {
			pub.Arrs[i2] = elem2.Publicize()
		}
	}
	if ut.IntVal != nil {
		pub.IntVal = ut.IntVal
	}
	return &pub
}

// MyTypes user type.
type MyTypes struct {
	// MyType Array
	Arrs []*MyType `form:"arrs,omitempty" json:"arrs,omitempty" yaml:"arrs,omitempty" xml:"arrs,omitempty"`
	// non-required Integer
	IntVal *int `form:"intVal,omitempty" json:"intVal,omitempty" yaml:"intVal,omitempty" xml:"intVal,omitempty"`
}
