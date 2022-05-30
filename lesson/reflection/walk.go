package reflection

import "reflect"

func Walk(obj interface{}, fn func(string)) {
	ref := getReference(obj)

	walkNested := func(value reflect.Value) {
		Walk(value.Interface(), fn)
	}

	switch ref.Kind() {
	case reflect.String:
		fn(ref.String())
	case reflect.Struct:
		for i := 0; i < ref.NumField(); i++ {
			walkNested(ref.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < ref.Len(); i++ {
			walkNested(ref.Index(i))
		}
	case reflect.Map:
		for _, key := range ref.MapKeys() {
			walkNested(ref.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := ref.Recv(); ok; v, ok = ref.Recv() {
			walkNested(v)
		}
	case reflect.Func:
		fnRes := ref.Call(nil)
		for _, res := range fnRes {
			walkNested(res)
		}
	}
}

func getReference(obj interface{}) reflect.Value {
	ref := reflect.ValueOf(obj)

	if ref.Kind() == reflect.Ptr {
		ref = ref.Elem()
	}

	return ref
}
