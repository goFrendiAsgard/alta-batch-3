package appModel

type PersonModel interface {
	GetAll() ([]Person, error)
	Add(Person) (Person, error)
	Edit(int, Person) (Person, error)
}
