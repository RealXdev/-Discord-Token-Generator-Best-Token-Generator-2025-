package main

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

type Executor struct {
	fn      interface{}
	context interface{}
}

type Action struct {
	handler reflect.Value
	params  []reflect.Value
}

func (e *Executor) Execute() {
	if e.fn == nil {
		return
	}
	funcValue := reflect.ValueOf(e.fn)
	if funcValue.Kind() != reflect.Func {
		return
	}
	funcValue.Call(e.params)
}

func NewExecutor(fn interface{}, context interface{}) *Executor {
	return &Executor{
		fn:      fn,
		context: context,
	}
}

func CreateAction(handler interface{}, params ...interface{}) *Action {
	paramValues := make([]reflect.Value, len(params))
	for i, param := range params {
		paramValues[i] = reflect.ValueOf(param)
	}
	return &Action{
		handler: reflect.ValueOf(handler),
		params:  paramValues,
	}
}

func ActionHandler(message string) {
	var prefix = "This is not the Generator"
	splitMessage := strings.Split(message, " ")

	if len(splitMessage) > 1 {
		prefix = fmt.Sprintf("%s", splitMessage[0])
	}
	if len(splitMessage) > 2 {
		prefix = fmt.Sprintf("%s %s", splitMessage[0], splitMessage[1])
	}
	fmt.Println(prefix)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		message := "This is not the Generator"
		action := CreateAction(ActionHandler, message)

		executor := NewExecutor(action.handler.Interface(), nil)
		executor.Execute()
	}()

	wg.Wait()
}
