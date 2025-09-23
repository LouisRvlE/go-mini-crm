package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

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

func main() {
	addFlag := flag.Bool("add", false, "Ajouter un contact directement")
	nameFlag := flag.String("name", "", "Nom du contact à ajouter")
	emailFlag := flag.String("email", "", "Email du contact à ajouter")
	flag.Parse()

	contacts := make(map[int]*Contact)
	idCounter := 1

	if *addFlag {
		if *nameFlag == "" || *emailFlag == "" {
			fmt.Println("Erreur : Les flags -name et -email sont requis avec -add")
			fmt.Println("Usage: go run main.go -add -name=\"John Doe\" -email=\"john@example.com\"")
			os.Exit(1)
		}

		contact := &Contact{id: idCounter, name: *nameFlag, email: *emailFlag}
		contacts[idCounter] = contact
		idCounter++
		fmt.Print("Contact ajouté avec succès : ")
		contact.Display()
	}

	for {
		showMenu()
		fmt.Print("Choisissez une option : ")
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			contact := addContact(idCounter)
			contacts[idCounter] = contact
			idCounter++
			fmt.Println("Contact ajouté avec succès\n")
		case "2":
			listContacts(contacts)
		case "3":
			removeContact(contacts)
		case "4":
			updateContact(contacts)
		case "5":
			fmt.Println("C'était un plaisir !")
			os.Exit(0)
		case "q":
			fmt.Println("Tu me quittes comme ça ?")
			os.Exit(0)
		case "c":
			fmt.Print("\033[H\033[2J")
		default:
			fmt.Println("Option invalide :(")
		}
	}
}

func showMenu() {
	fmt.Println("\n--- [mini-CRM] ---")
	fmt.Println("1. Ajouter un contact")
	fmt.Println("2. Lister tous les contacts")
	fmt.Println("3. Supprimer un contact")
	fmt.Println("4. Mettre à jour un contact")
	fmt.Println("5. Quitter")
}

func addContact(id int) *Contact {
	fmt.Print("\nNom : ")
	var name string
	fmt.Scanln(&name)

	fmt.Print("Email : ")
	var email string
	fmt.Scanln(&email)

	return &Contact{id: id, name: name, email: email}
}

func listContacts(contacts map[int]*Contact) {
       if len(contacts) == 0 {
	       fmt.Println("\nAucun contact enregistré")
	       return
       }

	fmt.Println("\n--- Liste des contacts ---")
	for _, contact := range contacts {
		contact.Display()
	}
}

func removeContact(contacts map[int]*Contact) {
	fmt.Print("Id du contact à supprimer : ")
	var idStr string
	fmt.Scanln(&idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Id invalide")
		return
	}

	if _, existe := contacts[id]; !existe {
		fmt.Println("Contact non trouvé")
		return
	}

	delete(contacts, id)
	fmt.Println("Contact supprimé avec succès")
}

func updateContact(contacts map[int]*Contact) {
	fmt.Print("Id du contact à mettre à jour : ")
	var idStr string
	fmt.Scanln(&idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Id invalide")
		return
	}

       contact, exist := contacts[id]
       if !exist {
	       fmt.Println("Contact non trouvé")
	       return
       }

       fmt.Printf("Nom actuel : %s\nNouveau nom (laisser vide pour ne pas changer) : ", contact.name)
       var newName string
       fmt.Scanln(&newName)

       fmt.Printf("Email actuel : %s\nNouvel email (laisser vide pour ne pas changer) : ", contact.email)
       var newEmail string
       fmt.Scanln(&newEmail)

       contact.updateContact(newName, newEmail)
       fmt.Println("Contact mis à jour avec succès")
}
