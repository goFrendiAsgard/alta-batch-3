package model

type PersonMemModel struct {
	data []Person
}

func NewPersonMemModel() *PersonMemModel {
	return &PersonMemModel{
		data: []Person{},
	}
}

func (pm *PersonMemModel) GetAll() ([]Person, error) {
	return pm.data, nil
}

func (pm *PersonMemModel) Add(p Person) (Person, error) {
	pm.data = append(pm.data, p)
	return p, nil
}
