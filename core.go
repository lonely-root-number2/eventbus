package evtbus

import (
	"context"
	"errors"
	"log"
	"reflect"
)

var mountsMap = map[reflect.Type][]Handler{}

type Handler struct {
	callback reflect.Value
}

func initMount(mounts ...any) {
	for i := range mounts {
		mount := mounts[i]
		valueOf := reflect.ValueOf(mount)
		typeOf := valueOf.Type()
		if !checkMountStruct(valueOf) {
			panic(valueOf.String() + "mounts must be struct")
		}
		log.Printf("=====start %v=====\n", typeOf.String())
		for i := 0; i < valueOf.NumMethod(); i++ {
			method := valueOf.Method(i)
			methodTyp := method.Type()
			if !checkHandler(methodTyp) {
				continue
			}
			evtTyp := methodTyp.In(1)
			registerHandlerMapping(evtTyp, buildHandler(method))
			log.Printf("registerHandlerMapping: %v --> %v\n", evtTyp.String(), typeOf.Method(i).Name)
		}
		log.Printf("=====end %v=====\n", typeOf.String())

	}
}

func registerHandlerMapping(evtTyp reflect.Type, handler Handler) {
	mountsMap[evtTyp] = append(mountsMap[evtTyp], handler)
}

func checkMountStruct(mount reflect.Value) bool {
	if mount.Kind() == reflect.Struct || mount.Elem().Kind() == reflect.Struct {
		return true
	}
	return false
}
func checkHandler(method reflect.Type) bool {
	if method.NumIn() != 2 || method.NumOut() != 1 {
		return false
	}

	if !method.In(0).Implements(ctxTyp) {
		return false
	}

	if !method.Out(0).Implements(errTyp) {
		return false
	}
	return true
}

func publish(ctx context.Context, evt any) (err error) {
	evtValue := reflect.ValueOf(evt)
	ctxValue := reflect.ValueOf(ctx)
	if handler, ok := mountsMap[evtValue.Type()]; ok {
		for i := range handler {
			callback := handler[i].callback
			result := callback.Call([]reflect.Value{ctxValue, evtValue})
			if !result[0].IsNil() {
				err = errors.Join(err, result[0].Interface().(error))
			}
		}
	}
	return err
}

func buildHandler(callback reflect.Value) Handler {
	return Handler{callback: callback}
}
