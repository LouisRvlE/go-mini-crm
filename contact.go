package main

import "fmt"

type Contact struct {
	id    int
	name  string
	email string
}

func (c *Contact) updateContact(name, email string) {
	if name != "" {
		c.name = name
	}
	if email != "" {
		c.email = email
	}
}

func (c *Contact) Display() {
	fmt.Printf("Id: %d, Nom: %s, Email: %s\n", c.id, c.name, c.email)
}

func NewContact(id int, name, email string) *Contact {
	return &Contact{
		id:    id,
		name:  name,
		email: email,
	}
}

func (c *Contact) GetID() int {
	return c.id
}

func (c *Contact) GetName() string {
	return c.name
}

func (c *Contact) GetEmail() string {
	return c.email
}