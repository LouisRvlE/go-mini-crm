package app

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	. "go-mini-crm/internal/storage"
)

func Run(store Storer) {
	addFlag := flag.Bool("add", false, "Ajouter un contact directement")
	nameFlag := flag.String("name", "", "Nom du contact à ajouter")
	emailFlag := flag.String("email", "", "Email du contact à ajouter")
	flag.Parse()

	if *addFlag {
		handleDirectAdd(store, *nameFlag, *emailFlag)
	}

	runInteractiveMode(store)
}

func HandleAddContact(store Storer) {
	fmt.Print("\nNom : ")
	var name string
	fmt.Scanln(&name)

	fmt.Print("Email : ")
	var email string
	fmt.Scanln(&email)

	contact := NewContact(store.GetNextID(), name, email)
	err := store.Add(contact)
	if err != nil {
		fmt.Printf("Erreur lors de l'ajout du contact : %v\n", err)
		return
	}
	fmt.Println("Contact ajouté avec succès")
}

func HandleListContacts(store Storer) {
	contacts := store.GetAll()
	if len(contacts) == 0 {
		fmt.Println("\nAucun contact enregistré")
		return
	}

	fmt.Println("\n--- Liste des contacts ---")
	for _, contact := range contacts {
		contact.Display()
	}
}

func HandleRemoveContact(store Storer) {
	fmt.Print("Id du contact à supprimer : ")
	var idStr string
	fmt.Scanln(&idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Id invalide")
		return
	}

	err = store.Remove(id)
	if err != nil {
		fmt.Printf("Erreur : %v\n", err)
		return
	}

	fmt.Println("Contact supprimé avec succès")
}

func HandleUpdateContact(store Storer) {
	fmt.Print("Id du contact à mettre à jour : ")
	var idStr string
	fmt.Scanln(&idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Id invalide")
		return
	}

	contact, exists := store.GetByID(id)
	if !exists {
		fmt.Println("Contact non trouvé")
		return
	}

	fmt.Printf("Nom actuel : %s\nNouveau nom (laisser vide pour ne pas changer) : ", contact.GetName())
	var newName string
	fmt.Scanln(&newName)

	fmt.Printf("Email actuel : %s\nNouvel email (laisser vide pour ne pas changer) : ", contact.GetEmail())
	var newEmail string
	fmt.Scanln(&newEmail)

	err = store.Update(id, newName, newEmail)
	if err != nil {
		fmt.Printf("Erreur lors de la mise à jour : %v\n", err)
		return
	}
	fmt.Println("Contact mis à jour avec succès")
}

func handleDirectAdd(store Storer, name, email string) {
	if name == "" || email == "" {
		fmt.Println("Erreur : Les flags -name et -email sont requis avec -add")
		fmt.Println("Usage: go run main.go -add -name=\"John Doe\" -email=\"john@example.com\"")
		os.Exit(1)
	}

	contact := NewContact(store.GetNextID(), name, email)
	err := store.Add(contact)
	if err != nil {
		fmt.Printf("Erreur lors de l'ajout du contact : %v\n", err)
		os.Exit(1)
	}
	fmt.Print("Contact ajouté avec succès : ")
	contact.Display()
}

func showMenu() {
	fmt.Println("\n--- [mini-CRM] ---")
	fmt.Println("1. Ajouter un contact")
	fmt.Println("2. Lister tous les contacts")
	fmt.Println("3. Supprimer un contact")
	fmt.Println("4. Mettre à jour un contact")
	fmt.Println("5. Quitter")
}

func runInteractiveMode(store Storer) {
	for {
		showMenu()
		fmt.Print("Choisissez une option : ")
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			HandleAddContact(store)
		case "2":
			HandleListContacts(store)
		case "3":
			HandleRemoveContact(store)
		case "4":
			HandleUpdateContact(store)
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
