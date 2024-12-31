package evtbus

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

type Mount struct {
}

type Evt1 struct {
	Id int
}
type Evt2 struct {
	Id int
}

var err1 = errors.New("h1 error")
var err2 = errors.New("h2 error")

func (m Mount) Handler1(ctx context.Context, evt Evt1) error {
	fmt.Printf("mount handler called: %v\n", evt)
	return err1
}
func (m Mount) Handler2(ctx context.Context, evt Evt2) error {
	fmt.Printf("mount handler2 called: %v\n", evt)
	return err2
}

func TestPublish(t *testing.T) {
	// The mount struct is like a namespace, with no other actual functionality or significance
	// And it must be a struct;
	//When a pointer is passed, it can map to both value receiver methods
	//and pointer receiver methods. When a value is passed,
	//it can only map to value receiver methods.
	//Generally, passing a pointer is sufficient.
	InitMounts(&Mount{})
	publish(context.Background(), Evt2{Id: 2})
	publish(context.Background(), Evt1{Id: 1})
	publish(context.Background(), Evt1{Id: 1})
	publish(context.Background(), Evt2{Id: 2})
}
