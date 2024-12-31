package evtbus

import (
	"context"
	"reflect"
)

var ctxTyp = reflect.TypeOf((*context.Context)(nil)).Elem()
var errTyp = reflect.TypeOf((*error)(nil)).Elem()
