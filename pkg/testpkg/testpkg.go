package testpkg

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Test struct {
	value string
	ctx   context.Context
}
type TestObj struct {
	name string `json:"name"`
	job  string `json:"job"`
}

func NewTest(value string) *Test {
	return &Test{value: value}
}

func (t *Test) Startup(ctx context.Context) {
	t.ctx = ctx
	runtime.EventsOn(t.ctx, "change_state", func(optionalData ...interface{}) {
		switch optionalData[0].(type) {
		default:
			runtime.LogError(t.ctx, "Incorrect data type")
		case string:
			t.SetValue(optionalData[0].(string))
		}
	})
	runtime.EventsEmit(t.ctx, "state_updated", t.value)
}

func (t *Test) SetValue(str string) {
	t.value = str
	runtime.EventsEmit(t.ctx, "state_updated", t.value)
}

func (t *Test) GetValue() string {
	return t.value
}

func (t *Test) Greet(name string) string {
	return t.value + " welcome " + name
}

func (t *Test) GetJson() *TestObj {
	return &TestObj{name: "Test", job: "Dev"}
}
