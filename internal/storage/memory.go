package storage

import "fmt"

type MemoryStore struct {
	contacts  map[int]*Contact
	idCounter int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts:  make(map[int]*Contact),
		idCounter: 1,
	}
}

func (ms *MemoryStore) Add(contact *Contact) error {
	ms.contacts[contact.ID] = contact
	return nil
}

func (ms *MemoryStore) GetAll() map[int]*Contact {
	return ms.contacts
}

func (ms *MemoryStore) GetByID(id int) (*Contact, bool) {
	contact, exists := ms.contacts[id]
	return contact, exists
}

func (ms *MemoryStore) Remove(id int) error {
	if _, exists := ms.contacts[id]; !exists {
		return fmt.Errorf("contact avec l'ID %d n'existe pas", id)
	}
	delete(ms.contacts, id)
	return nil
}

func (ms *MemoryStore) Update(id int, name, email string) error {
	contact, exists := ms.contacts[id]
	if !exists {
		return fmt.Errorf("contact avec l'ID %d n'existe pas", id)
	}
	contact.updateContact(name, email)
	return nil
}

func (ms *MemoryStore) GetNextID() int {
	id := ms.idCounter
	ms.idCounter++
	return id
}