package main

import (
	"fmt"
	"strings"
)

type Contact struct {
	Name  string
	Email string
	Phone string
}

type ContactManager struct {
	names map[string]Contact
}

func (cm *ContactManager) Add(c Contact) error {
	key := strings.ToLower(c.Name)
	if _, ok := cm.names[key]; ok {
		return fmt.Errorf("Contact with name %s already exists", c.Name)
	}

	cm.names[key] = c

	return nil
}

func (cm *ContactManager) Find(name string) (Contact, error) {
	key := strings.ToLower(name)

	if c, ok := cm.names[key]; ok {
		return c, nil
	}

	return Contact{}, fmt.Errorf("Contact with name %s not found", name)
}

func (cm *ContactManager) Delete(name string) error {
	key := strings.ToLower(name)

	if _, ok := cm.names[key]; ok {
		delete(cm.names, key)
		return nil
	}

	return fmt.Errorf("Contact with name %s not found", name)
}

func (cm *ContactManager) List() []Contact {
	contacts := make([]Contact, 0)

	for _, v := range cm.names {
		contacts = append(contacts, v)
	}

	return contacts
}

func main() {
	cm := ContactManager{
		names: make(map[string]Contact),
	}

	err := cm.Add(Contact{Name: "Alice Wonderland", Email: "alice@example.com", Phone: "111-2222"})

	if err != nil {
		fmt.Println(err.Error())
	}

	err = cm.Add(Contact{Name: "Bob The Builder", Email: "bob@example.com", Phone: "333-4444"})

	if err != nil {
		fmt.Println(err.Error())
	}

	err = cm.Add(Contact{Name: "Charlie Brown", Email: "charlie@example.com", Phone: "555-6666"})

	if err != nil {
		fmt.Println(err.Error())
	}

	err = cm.Add(Contact{Name: "Alice Wonderland", Email: "alice@example.com", Phone: "111-2222"})

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Contact List", cm.List())

	contactFound, err := cm.Find("bob the builder")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Contact found:", contactFound)
	}

	contactFound, err = cm.Find("bobby")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Contact found:", contactFound)
	}

	err = cm.Delete("Alice Wonderland")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Contact deleted:", cm.List())
	}
}
