package sebarcore

import "reflect"

type DataModel struct {
	ID   string
	Name string

	Fields  []*Field
	Methods []*Method

	rtype  reflect.Type
	rvalue reflect.Value
}

func (d *DataModel) AddField(fs ...*Field) {
	d.Fields = append(d.Fields, fs...)
}

func (d *DataModel) RemoveField(id string) {
	for i, v := range d.Fields {
		if v.ID == id {
			d.Fields = append(d.Fields[0:i], d.Fields[i+1:]...)
		}
	}
}

func Obj2DataModel(o interface{}) *DataModel {
	//sebarCheck()
	dm := new(DataModel)

	v := reflect.Indirect(reflect.ValueOf(o))
	t := v.Type()

	numField := t.NumField()

	dm.ID = t.Kind().String()
	dm.Name = dm.ID

	dm.rtype = t
	dm.rvalue = v

	for fidx := 0; fidx < numField; fidx++ {
		f := t.Field(fidx)
		df := new(Field)
		df.ID = f.Name
		df.Type = f.Type.String()
		df.Title = f.Name
		dm.AddField(df)
	}

	tm := reflect.TypeOf(o)
	numMethod := tm.NumMethod()
	for fidx := 0; fidx < numMethod; fidx++ {
		tmtd := tm.Method(fidx)
		dm.AddMethod(tmtd)
	}

	return dm
}

func (dm *DataModel) MethodByName(name string) *Method {
	for _, m := range dm.Methods {
		if m.Name == name {
			return m
		}
	}
	return nil
}

func (dm *DataModel) AddMethod(mtd reflect.Method) error {
	m := NewMethodFromRM(mtd)
	dm.Methods = append(dm.Methods, m)
	return nil
}
