package evtbus

import (
	"context"
)

func InitMounts(mounts ...any) {
	initMount(mounts...)
}

//func SetHandlerOrder(evt any, handlers []any) {
//	evtTyp := reflect.TypeOf(evt)
//	handlerVals := make([]Handler, len(handlers))
//	for _, h := range handlers {
//		mountsMap[evtTyp] = append(mountsMap[evtTyp], buildHandler(reflect.ValueOf(h)))
//	}
//	mountsMap[evtTyp] = handlerVals
//}

func Publish(ctx context.Context, evt any) error {
	return publish(ctx, evt)
}
