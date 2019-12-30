package di

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

type container struct {
	funStore map[reflect.Type]interface{}
	store    map[reflect.Type]interface{}
}

func New() *container {
	return &container{
		funStore: map[reflect.Type]interface{}{},
		store:    map[reflect.Type]interface{}{},
	}
}

func (c *container) registerFun(f interface{}) error {
	t := reflect.TypeOf(f)

	if t.Kind() != reflect.Func {
		return errors.New("f must be a function")
	}

	if t.NumOut() != 1 {
		return errors.New("f must have 1 return value")
	}

	c.funStore[t.Out(0)] = f

	return nil
}

func (c *container) execute(f interface{}) error {
	typeF := reflect.TypeOf(f)
	if typeF.Kind() != reflect.Func {
		return errors.New("not a func")
	}

	if typeF.NumOut() > 0 {
		return errors.New("f must not have any return values")
	}

	// Build parameter values
	params := make([]reflect.Value, typeF.NumIn())
	for i := range params {
		expected := typeF.In(i)

		// Check for evaluated value
		if res, ok := c.store[expected]; ok {
			params[i] = reflect.ValueOf(res)
			continue
		}

		// Check for function that evaluate to desired result
		if res, ok := c.funStore[expected]; ok {
			out := reflect.ValueOf(res).Call(nil)
			c.store[expected] = out[0]
			delete(c.funStore, expected)
			params[i] = out[0]
			continue
		}

		return errors.Errorf("no object with type %s found in store", expected.Name())
	}

	for _, param := range params {
		fmt.Println(param)
	}

	ffff := reflect.ValueOf(f)

	ffff.Call(params)

	return nil
}
