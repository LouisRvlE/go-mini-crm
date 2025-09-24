package storage

import "fmt"

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (c *Contact) updateContact(name, email string) {
	if name != "" {
		c.Name = name
	}
	if email != "" {
		c.Email = email
	}
}

func (c *Contact) Display() {
	fmt.Printf("Id: %d, Nom: %s, Email: %s\n", c.ID, c.Name, c.Email)
}

func NewContact(id int, name, email string) *Contact {
	return &Contact{
		ID:    id,
		Name:  name,
		Email: email,
	}
}

func (c *Contact) GetID() int {
	return c.ID
}

func (c *Contact) GetName() string {
	return c.Name
}

func (c *Contact) GetEmail() string {
	return c.Email
}