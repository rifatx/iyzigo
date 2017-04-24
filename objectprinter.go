package iyzigo

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

type ObjectPrinter struct {
	buf               *bytes.Buffer
	keyValueSeparator string
	fieldSeparator    string
	arraySeparator    string
	nameOpener        string
	nameCloser        string
	numericOpener     string
	numericCloser     string
	stringOpener      string
	stringCloser      string
	boolOpener        string
	boolCloser        string
	objectOpener      string
	objectCloser      string
	arrayOpener       string
	arrayCloser       string
}

const (
	tagName = "opt"
)

func InitObjectPrinter() *ObjectPrinter {
	return &ObjectPrinter{buf: new(bytes.Buffer)}
}

func (op *ObjectPrinter) SetKeyValueSeparator(kvps string) *ObjectPrinter {
	op.keyValueSeparator = kvps
	return op
}

func (op *ObjectPrinter) SetFieldSeparator(fs string) *ObjectPrinter {
	op.fieldSeparator = fs
	return op
}

func (op *ObjectPrinter) SetArraySeparator(as string) *ObjectPrinter {
	op.arraySeparator = as
	return op
}

func (op *ObjectPrinter) SetNameEncapsulation(opener string, closer string) *ObjectPrinter {
	op.nameOpener = opener
	op.nameCloser = closer
	return op
}

func (op *ObjectPrinter) SetNumericEncapsulation(opener string, closer string) *ObjectPrinter {
	op.numericOpener = opener
	op.numericCloser = closer
	return op
}

func (op *ObjectPrinter) SetStringEncapsulation(opener string, closer string) *ObjectPrinter {
	op.stringOpener = opener
	op.stringCloser = closer
	return op
}

func (op *ObjectPrinter) SetBoolEncapsulation(opener string, closer string) *ObjectPrinter {
	op.boolOpener = opener
	op.boolCloser = closer
	return op
}

func (op *ObjectPrinter) SetObjectEncapsulation(opener string, closer string) *ObjectPrinter {
	op.objectOpener = opener
	op.objectCloser = closer
	return op
}

func (op *ObjectPrinter) SetArrayEncapsulation(opener string, closer string) *ObjectPrinter {
	op.arrayOpener = opener
	op.arrayCloser = closer
	return op
}

func (op *ObjectPrinter) writeKvp(sf reflect.StructField, valueOpener string, value interface{}, valueCloser string) {
	op.writeName(sf)
	op.buf.WriteString(fmt.Sprintf("%s%v%s", valueOpener, value, valueCloser))
}

func (op *ObjectPrinter) writeName(sf reflect.StructField) {
	if len(sf.Name) == 0 {
		return
	}
	name := sf.Name
	if n, ok := sf.Tag.Lookup(tagName); ok {
		ts := strings.Split(n, ",")
		tn := strings.Split(ts[0], ":")
		if len(tn) == 2 && tn[0] == "name" {
			name = tn[1]
		}
	}
	if len(name) > 0 {
		op.buf.WriteString(fmt.Sprintf("%s%s%s%s", op.nameOpener, name, op.nameCloser, op.keyValueSeparator))
	}
}

func (op *ObjectPrinter) writeSingle(valueOpener string, value interface{}, valueCloser string) {
	op.buf.WriteString(fmt.Sprintf("%s%v%s", valueOpener, value, valueCloser))
}

func (op *ObjectPrinter) openObject() {
	op.buf.WriteString(op.objectOpener)
}

func (op *ObjectPrinter) closeObject() {
	op.buf.WriteString(op.objectCloser)
}

func (op *ObjectPrinter) openArray() {
	op.buf.WriteString(op.arrayOpener)
}

func (op *ObjectPrinter) closeArray() {
	op.buf.WriteString(op.arrayCloser)
}

func (op *ObjectPrinter) writeArraySeparator() {
	op.buf.WriteString(op.arraySeparator)
}

func (op *ObjectPrinter) writeFieldSeparator() {
	op.buf.WriteString(op.fieldSeparator)
}

func isOmitted(sf reflect.StructField, v reflect.Value) bool {
	omitempty := false
	if n, ok := sf.Tag.Lookup(tagName); ok {
		ts := strings.Split(n, ",")
		if len(ts) > 1 && ts[1] == "omitempty" {
			omitempty = true
		}
	}
	if !omitempty {
		return false
	}
	switch sf.Type.Kind() {
	case reflect.String, reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		return v.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Println(sf.Name, " ", v.Int())
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0.0
	case reflect.Struct:
		return len(InitObjectPrinter().GetObjectString(v.Interface())) == 0
	case reflect.Ptr:
		return v.IsNil() == true
	}
	return false
}

func (op *ObjectPrinter) readField(t reflect.StructField, v reflect.Value, written bool) bool {
	if omitted := isOmitted(t, v); omitted {
		return false
	}
	if written {
		op.writeFieldSeparator()
	}
	switch t.Type.Kind() {
	case reflect.Ptr:
		op.writeName(t)
		op.crawlInterface(v.Elem().Interface())
	case reflect.String:
		op.writeKvp(t, op.stringOpener, v.String(), op.stringCloser)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		op.writeKvp(t, op.numericOpener, v.Int(), op.numericCloser)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		op.writeKvp(t, op.numericOpener, v.Uint(), op.numericCloser)
	case reflect.Float32, reflect.Float64:
		op.writeKvp(t, op.numericOpener, v.Float(), op.numericCloser)
	case reflect.Bool:
		op.writeKvp(t, op.boolOpener, v.Bool(), op.boolCloser)
	case reflect.Slice:
		op.writeName(t)
		op.openArray()
		for j := 0; j < v.Len(); j++ {
			ev := reflect.ValueOf(v.Slice(j, j+1))
			switch ev.Kind() {
			case reflect.Interface, reflect.Ptr:
				ev = ev.Elem()
			}
			switch ev.Kind() {
			case reflect.Struct:
				op.crawlInterface(v.Index(j).Interface())
			}
			if j < v.Len()-1 {
				op.writeArraySeparator()
			}
		}
		op.closeArray()
	case reflect.Struct:
		op.writeName(t)
		op.crawlInterface(v.Interface())
	}
	return true
}

func (op *ObjectPrinter) crawlInterface(o interface{}) {
	if o == nil {
		return
	}
	t := reflect.TypeOf(o)
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		t = t.Elem()
	}
	v := reflect.ValueOf(o)
	switch v.Kind() {
	case reflect.Interface, reflect.Ptr:
		v = v.Elem()
	}
	switch t.Kind() {
	case reflect.Struct:
		op.openObject()
		written := false
		for i := 0; i < t.NumField(); i++ {
			written = op.readField(t.Field(i), v.Field(i), written) || written
		}
		op.closeObject()
	case reflect.String:
		op.writeSingle(op.stringOpener, v.String(), op.stringCloser)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		op.writeSingle(op.numericOpener, v.Int(), op.numericCloser)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		op.writeSingle(op.numericOpener, v.Uint(), op.numericCloser)
	case reflect.Float32, reflect.Float64:
		op.writeSingle(op.numericOpener, v.Float(), op.numericCloser)
	case reflect.Bool:
		op.writeSingle(op.boolOpener, v.Bool(), op.boolCloser)
	}
}

func (op *ObjectPrinter) GetObjectString(o interface{}) string {
	op.crawlInterface(o)
	return op.buf.String()
}
