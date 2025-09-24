package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type JSONStore struct {
	contacts  map[int]*Contact
	idCounter int
	filePath  string
	mu        sync.RWMutex
}

func NewJSONStore() *JSONStore {
	store := &JSONStore{
		contacts:  make(map[int]*Contact),
		idCounter: 1,
		filePath:  "contacts.json",
	}
	store.loadFromFile()
	return store
}

func (jsonStore *JSONStore) Add(contact *Contact) error {
	jsonStore.mu.Lock()
	defer jsonStore.mu.Unlock()
	jsonStore.contacts[contact.ID] = contact
	return jsonStore.saveToFile()
}

func (jsonStore *JSONStore) GetAll() map[int]*Contact {
	jsonStore.mu.RLock()
	defer jsonStore.mu.RUnlock()
	return jsonStore.contacts
}

func (jsonStore *JSONStore) GetByID(id int) (*Contact, bool) {
	jsonStore.mu.RLock()
	defer jsonStore.mu.RUnlock()
	contact, exists := jsonStore.contacts[id]
	return contact, exists
}

func (jsonStore *JSONStore) Remove(id int) error {
	jsonStore.mu.Lock()
	defer jsonStore.mu.Unlock()
	if _, exists := jsonStore.contacts[id]; !exists {
		return fmt.Errorf("contact avec l'ID %d n'existe pas", id)
	}
	delete(jsonStore.contacts, id)
	return jsonStore.saveToFile()
}

func (jsonStore *JSONStore) Update(id int, name, email string) error {
	jsonStore.mu.Lock()
	defer jsonStore.mu.Unlock()
	contact, exists := jsonStore.contacts[id]
	if !exists {
		return fmt.Errorf("contact avec l'ID %d n'existe pas", id)
	}
	contact.updateContact(name, email)
	return jsonStore.saveToFile()
}

func (jsonStore *JSONStore) GetNextID() int {
	jsonStore.mu.Lock()
	defer jsonStore.mu.Unlock()
	id := jsonStore.idCounter
	jsonStore.idCounter++
	return id
}

func (jsonStore *JSONStore) loadFromFile() error {
	data, err := os.ReadFile(jsonStore.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	var contacts []*Contact
	if err := json.Unmarshal(data, &contacts); err != nil {
		return err
	}

	jsonStore.contacts = make(map[int]*Contact)
	maxID := 0
	for _, contact := range contacts {
		jsonStore.contacts[contact.ID] = contact
		if contact.ID > maxID {
			maxID = contact.ID
		}
	}
	jsonStore.idCounter = maxID + 1

	return nil
}

func (jsonStore *JSONStore) saveToFile() error {
	contacts := make([]*Contact, 0, len(jsonStore.contacts))
	for _, contact := range jsonStore.contacts {
		contacts = append(contacts, contact)
	}

	data, err := json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(jsonStore.filePath, data, 0644)
}