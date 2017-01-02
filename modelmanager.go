package sebarcore

type ModelManager struct {
	models map[string]*DataModel
}

func NewModelManager() *ModelManager {
	mm := new(ModelManager)
	mm.models = map[string]*DataModel{}
	return mm
}

func (mm *ModelManager) Register(o interface{}, name string) *DataModel {
	model := Obj2DataModel(o)
	if name == "" {
		name = model.ID
	}
	mm.models[name] = model
	return model
}

func (mm *ModelManager) Unregister(name string) {
	delete(mm.models, name)
}

func (mm *ModelManager) New(name string) (interface{}, error) {
	m, b := mm.models[name]
	if !b {
		return nil, ThrowErr("sebarcore", "ModelManager.New", "Model "+name+" is not exist")
	}
	model := m.New()
	return model, nil
}
