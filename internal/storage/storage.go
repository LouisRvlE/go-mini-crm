package storage

type Storer interface {
	Add(contact *Contact) error
	GetAll() map[int]*Contact
	GetByID(id int) (*Contact, bool)
	Remove(id int) error
	Update(id int, name, email string) error
	GetNextID() int
}
