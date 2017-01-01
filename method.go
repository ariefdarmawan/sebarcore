package sebarcore

import (
	"reflect"
)

type Method struct {
	Name    string
	ParmIn  []*MethodParm
	ParmOut []*MethodParm

	mvalue reflect.Value
	mtype  reflect.Type
}

type MethodParm struct {
	Type string

	ReflectType reflect.Type
}

func NewMethodFromRM(rm reflect.Method) *Method {
	m := new(Method)
	m.Name = rm.Name

	mtype := rm.Type
	numin := mtype.NumIn()
	numout := mtype.NumOut()

	m.mvalue = rm.Func

	for idx := 1; idx < numin; idx++ {
		op := mtype.In(idx)
		parm := new(MethodParm)
		parm.Type = op.String()
		parm.ReflectType = op
		m.ParmIn = append(m.ParmIn, parm)
	}

	for idx := 0; idx < numout; idx++ {
		parm := new(MethodParm)
		m.ParmIn = append(m.ParmIn, parm)
	}

	return m
}

func (m *Method) Do(in []interface{}) []reflect.Value {
	invs := []reflect.Value{}
	for _, v := range in {
		invs = append(invs, reflect.ValueOf(v))
	}
	outs := m.mvalue.Call(invs)

	var ifaces []reflect.Value
	for _, o := range outs {
		ifaces = append(ifaces, o)
	}
	return ifaces
}

func (m *Method) Exec(in []interface{}, outs []interface{}) error {
	invs := []reflect.Value{}
	for _, v := range in {
		invs = append(invs, reflect.ValueOf(v))
	}
	outvs := m.mvalue.Call(invs)

	for idx, outv := range outvs {
		rf := reflect.ValueOf(outs[idx])
		rf.Elem().Set(outv)
	}

	return nil
}
