package testpkg

import "context"

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
