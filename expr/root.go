package expr

import (
	"net/url"
	"sort"
	"strings"

	"goa.design/goa/eval"
)

// Root is the root object built by the DSL.
var Root = &RootExpr{GeneratedTypes: &GeneratedRoot{}}

type (
	// RootExpr is the struct built by the DSL on process start.
	RootExpr struct {
		// API contains the API expression built by the DSL.
		API *APIExpr
		// Services contains the list of services exposed by the API.
		Services []*ServiceExpr
		// Errors contains the list of errors returned by all the API
		// methods.
		Errors []*ErrorExpr
		// Types contains the user types described in the DSL.
		Types []UserType
		// ResultTypes contains the result types described in the DSL.
		ResultTypes []UserType
		// GeneratedTypes contains the types generated during DSL
		// execution.
		GeneratedTypes *GeneratedRoot
		// Conversions list the user type to external type mappings.
		Conversions []*TypeMap
		// Creations list the external type to user type mappings.
		Creations []*TypeMap
		// Schemes list the registered security schemes.
		Schemes []*SchemeExpr

		// Meta is a set of key/value pairs with semantic that is
		// specific to each generator.
		Meta MetaExpr
	}

	// MetaExpr is a set of key/value pairs
	MetaExpr map[string][]string

	// TypeMap defines a user to external type mapping.
	TypeMap struct {
		// User is the user type being converted or created.
		User UserType

		// External is an instance of the type being converted from or to.
		External interface{}
	}
)

// NameMap returns the attribute and transport element name encoded in the given
// string. The encoding uses a simple "attribute:element" notation which allows
// to map transport field names (HTTP headers etc.) to underlying attributes.
// The second element of the encoding is optional in which case both the element
// and attribute have the same name.
func NameMap(encoded string) (string, string) {
	elems := strings.Split(encoded, ":")
	attName := elems[0]
	name := attName
	if len(elems) > 1 {
		name = elems[1]
	}
	return attName, name
}

// WalkSets returns the expressions in order of evaluation.
func (r *RootExpr) WalkSets(walk eval.SetWalker) {
	if r.API == nil {
		r.API = NewAPIExpr("API", func() {})
	}

	// Top level API DSL
	walk(eval.ExpressionSet{r.API})

	// User types
	types := make(eval.ExpressionSet, len(r.Types))
	for i, t := range r.Types {
		types[i] = t.Attribute()
	}
	walk(types)

	// Result types
	mtypes := make(eval.ExpressionSet, len(r.ResultTypes))
	for i, mt := range r.ResultTypes {
		mtypes[i] = mt.(*ResultTypeExpr)
	}
	walk(mtypes)

	// Services
	services := make(eval.ExpressionSet, len(r.Services))
	var methods eval.ExpressionSet
	for i, s := range r.Services {
		services[i] = s
	}
	walk(services)

	// Methods (must be done after services)
	for _, s := range r.Services {
		for _, m := range s.Methods {
			methods = append(methods, m)
		}
	}
	walk(methods)

	// HTTP services and endpoints
	httpsvcs := make(eval.ExpressionSet, len(r.API.HTTP.Services))
	sort.SliceStable(r.API.HTTP.Services, func(i, j int) bool {
		if r.API.HTTP.Services[j].ParentName == r.API.HTTP.Services[i].Name() {
			return true
		}
		return false
	})
	var httpepts eval.ExpressionSet
	var httpsvrs eval.ExpressionSet
	for i, svc := range r.API.HTTP.Services {
		httpsvcs[i] = svc
		for _, e := range svc.HTTPEndpoints {
			httpepts = append(httpepts, e)
		}
		for _, s := range svc.FileServers {
			httpsvrs = append(httpsvrs, s)
		}
	}
	walk(httpsvcs)
	walk(httpepts)
	walk(httpsvrs)

	// GRPC services and endpoints
	grpcsvcs := make(eval.ExpressionSet, len(r.API.GRPC.Services))
	sort.SliceStable(r.API.GRPC.Services, func(i, j int) bool {
		if r.API.GRPC.Services[j].ParentName == r.API.GRPC.Services[i].Name() {
			return true
		}
		return false
	})
	var grpcepts eval.ExpressionSet
	for i, svc := range r.API.GRPC.Services {
		grpcsvcs[i] = svc
		for _, e := range svc.GRPCEndpoints {
			grpcepts = append(grpcepts, e)
		}
	}
	walk(grpcsvcs)
	walk(grpcepts)
}

// DependsOn returns nil, the core DSL has no dependency.
func (r *RootExpr) DependsOn() []eval.Root { return nil }

// Packages returns the Go import path to this and the dsl packages.
func (r *RootExpr) Packages() []string {
	return []string{
		"goa.design/goa/expr",
		"goa.design/goa/dsl",
	}
}

// UserType returns the user type expression with the given name if found, nil otherwise.
func (r *RootExpr) UserType(name string) UserType {
	for _, t := range r.Types {
		if t.Name() == name {
			return t
		}
	}
	for _, t := range r.ResultTypes {
		if t.Name() == name {
			return t
		}
	}
	return nil
}

// GeneratedResultType returns the generated result type expression with the given
// id, nil if there isn't one.
func (r *RootExpr) GeneratedResultType(id string) *ResultTypeExpr {
	for _, t := range *r.GeneratedTypes {
		mt := t.(*ResultTypeExpr)
		if mt.Identifier == id {
			return mt
		}
	}
	return nil
}

// Service returns the service with the given name.
func (r *RootExpr) Service(name string) *ServiceExpr {
	for _, s := range r.Services {
		if s.Name == name {
			return s
		}
	}
	return nil
}

// Error returns the error with the given name.
func (r *RootExpr) Error(name string) *ErrorExpr {
	for _, e := range r.Errors {
		if e.Name == name {
			return e
		}
	}
	return nil
}

// HTTPSchemes returns the list of HTTP schemes used by the API servers.
func (r *RootExpr) HTTPSchemes() []string {
	schemes := make(map[string]bool)
	for _, s := range r.API.Servers {
		if u, err := url.Parse(s.URL); err != nil {
			schemes[u.Scheme] = true
		}
	}
	if len(schemes) == 0 {
		return nil
	}
	ss := make([]string, len(schemes))
	i := 0
	for s := range schemes {
		ss[i] = s
		i++
	}
	sort.Strings(ss)
	return ss
}

// EvalName is the name of the DSL.
func (r *RootExpr) EvalName() string {
	return "design"
}

// Validate makes sure the root expression is valid for code generation.
func (r *RootExpr) Validate() error {
	var verr eval.ValidationErrors
	if r.API == nil {
		verr.Add(r, "Missing API declaration")
	}
	return &verr
}

// Dup creates a new map from the given expression.
func (m MetaExpr) Dup() MetaExpr {
	d := make(MetaExpr, len(m))
	for k, v := range m {
		d[k] = v
	}
	return d
}

// Merge merges src meta expression with m. If meta has intersecting set of
// keys on both m and src, then the values for those keys in src is appended
// to the values of the keys in m if not already existing.
func (m MetaExpr) Merge(src MetaExpr) {
	for k, vals := range src {
		if mvals, ok := m[k]; ok {
			var found bool
			for _, v := range vals {
				found = false
				for _, mv := range mvals {
					if mv == v {
						found = true
						break
					}
				}
				if !found {
					mvals = append(mvals, v)
				}
			}
			m[k] = mvals
		} else {
			m[k] = vals
		}
	}
}
