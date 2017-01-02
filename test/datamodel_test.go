package coretest

import (
	"eaciit/sebarcore"
	"fmt"
	"reflect"
	"testing"

	"github.com/eaciit/toolkit"
)

type DummyModel struct {
	ID           string
	Title        string
	RandomNumber int
}

func (dm *DummyModel) Hello(variable string) (string, error) {
	/*
		if dm.ID == "" || dm.Title == "" {
			fmt.Println("RESULT -------> ", "ERROR")
			return "", errors.New("Either ID or Title is empty")
		}
	*/

	if dm.Title == "" {
		dm.Title = "Mr"
	}

	ret := fmt.Sprintf("Hello %s, your message is %s", dm.Title, variable)
	//fmt.Println("RESULT -------> ", ret)
	return ret, nil
}

func TestInit(t *testing.T) {
	if err := sebarcore.SetConfig("../config.json"); err != nil {
		fmt.Println("Unable to load config:", err.Error())
	}
}

func TestLoadModel(t *testing.T) {
	dm := sebarcore.Obj2DataModel(new(DummyModel))
	if len(dm.Fields) != 3 || len(dm.Methods) != 1 {
		t.Error("Unable to receive model")
	} else {
		t.Logf("Model received: %s\nFields: %s\nMethods: %s",
			toolkit.JsonString(dm),
			toolkit.JsonString(dm.Fields),
			toolkit.JsonString(dm.Methods))
	}
}

var outs []reflect.Value

func TestDo(t *testing.T) {
	model := new(DummyModel)
	dm := sebarcore.Obj2DataModel(model)
	outs = dm.Methods[0].Do([]interface{}{model, "Arief Darmawan"})
	if len(outs) != 2 {
		t.Errorf("Unable to call")
	} else {
		t.Logf("Result: %s", toolkit.JsonString(outs))
	}
}

func TestParseOut(t *testing.T) {
	ret := ""
	var err error

	if len(outs) != 2 {
		t.Errorf("Unable to call")
	} else {
		v1 := reflect.ValueOf(&ret)
		v2 := outs[0]
		if errset := setvalue(v1, v2); errset != nil {
			t.Errorf(errset.Error())
		}
		if ret == "" {
			t.Errorf("Unable to receive result")
		}
		t.Logf("Result: %s", toolkit.JsonString(map[string]interface{}{"message": ret, "error": err}))
	}
}

func TestExec(t *testing.T) {
	ret := ""
	var err error

	model := new(DummyModel)
	dm := sebarcore.Obj2DataModel(model)
	errExec := dm.Methods[0].Exec([]interface{}{model, "Arief Darmawan"},
		[]interface{}{&ret, &err})
	if errExec != nil {
		t.Errorf("Unable to call: %s", errExec.Error())
	} else {
		t.Logf("Result: \nMsg: %s\nErr: %s", ret, toolkit.JsonString(err))
	}
}

func setvalue(v1, v2 reflect.Value) error {
	//v1t, v2t := v1.Type(), v2.Type()
	//_, v2ts := v1t.String(), v2t.String()
	/*
		if v1ts != v2ts {
			return fmt.Errorf("Type is different: %s and %s", v1ts, v2ts)
		}
	*/

	v1.Elem().Set(v2)
	return nil
}
