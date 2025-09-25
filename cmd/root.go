package cmd

import (
	"fmt"
	"log"
	"os"

	"go-mini-crm/internal/app"
	"go-mini-crm/internal/config"
	"go-mini-crm/internal/storage"

	"github.com/spf13/cobra"
)

var store storage.Storer
var cfg *config.Config

var rootCmd = &cobra.Command{
	Use:   "mini-crm",
	Short: "Un mini-CRM en ligne de commande",
	Long:  `Un mini-CRM simple pour gérer vos contacts avec les opérations CRUD de base.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("=== Mini-CRM ===")
		fmt.Println("Un gestionnaire de contacts simple et efficace")
		fmt.Println()
		fmt.Println("Configuration actuelle:")
		fmt.Println(config.GetConfigInfo(cfg))
		fmt.Println()
		fmt.Println("Commandes disponibles:")
		fmt.Println("  mini-crm interactive    - Lance le mode interactif")
		fmt.Println("  mini-crm add           - Ajoute un contact")
		fmt.Println("  mini-crm list          - Liste tous les contacts")
		fmt.Println("  mini-crm update        - Met à jour un contact")
		fmt.Println("  mini-crm delete        - Supprime un contact")
		fmt.Println("  mini-crm help          - Affiche cette aide")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Charger la configuration
	var err error
	cfg, err = config.LoadConfig()
	if err != nil {
		log.Fatalf("Erreur lors du chargement de la configuration : %v", err)
	}

	// Initialiser le store selon la configuration
	store = initializeStore(cfg)

	// Ajouter toutes les commandes
	rootCmd.AddCommand(interactiveCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(configCmd)
}

// initializeStore initialise le store selon la configuration
func initializeStore(cfg *config.Config) storage.Storer {
	switch cfg.Storage.Type {
	case "memory":
		log.Println("Initialisation du stockage en mémoire")
		return storage.NewMemoryStore()
	case "json":
		log.Printf("Initialisation du stockage JSON : %s", cfg.Storage.JSONFile)
		return storage.NewJSONStore()
	case "gorm":
		log.Printf("Initialisation du stockage GORM : %s", cfg.Storage.DBPath)
		return storage.NewGORMStore(cfg.Storage.DBPath)
	default:
		log.Fatalf("Type de stockage inconnu : %s", cfg.Storage.Type)
		return nil
	}
}

// Commande interactive
var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Lance le mode interactif",
	Long:  `Lance le mode interactif du CRM avec un menu pour naviguer entre les options.`,
	Run: func(cmd *cobra.Command, args []string) {
		runInteractiveMode(store)
	},
}

// Commande add
var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Ajoute un nouveau contact",
	Long:    `Ajoute un nouveau contact au CRM. Vous devez spécifier le nom et l'email.`,
	Example: "mini-crm add --name=\"John Doe\" --email=\"john@example.com\"",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		email, _ := cmd.Flags().GetString("email")

		if name == "" || email == "" {
			fmt.Println("Erreur : Les flags --name et --email sont requis")
			fmt.Println("Usage: mini-crm add --name=\"John Doe\" --email=\"john@example.com\"")
			os.Exit(1)
		}

		contact := storage.NewContact(store.GetNextID(), name, email)
		err := store.Add(contact)
		if err != nil {
			fmt.Printf("Erreur lors de l'ajout du contact : %v\n", err)
			os.Exit(1)
		}
		fmt.Print("Contact ajouté avec succès : ")
		contact.Display()
	},
}

// Commande list
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Liste tous les contacts",
	Long:  `Affiche la liste de tous les contacts enregistrés dans le CRM.`,
	Run: func(cmd *cobra.Command, args []string) {
		contacts := store.GetAll()
		if len(contacts) == 0 {
			fmt.Println("Aucun contact enregistré")
			return
		}

		fmt.Println("--- Liste des contacts ---")
		for _, contact := range contacts {
			contact.Display()
		}
	},
}

// Commande update
var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Met à jour un contact existant",
	Long:    `Met à jour un contact existant en spécifiant son ID et les nouvelles valeurs.`,
	Example: "mini-crm update --id=1 --name=\"Jane Doe\" --email=\"jane@example.com\"",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")
		email, _ := cmd.Flags().GetString("email")

		if id == 0 {
			fmt.Println("Erreur : Le flag --id est requis")
			fmt.Println("Usage: mini-crm update --id=1 --name=\"Jane Doe\" --email=\"jane@example.com\"")
			os.Exit(1)
		}

		_, exists := store.GetByID(id)
		if !exists {
			fmt.Println("Contact non trouvé")
			os.Exit(1)
		}

		err := store.Update(id, name, email)
		if err != nil {
			fmt.Printf("Erreur lors de la mise à jour : %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Contact mis à jour avec succès")
	},
}

// Commande delete
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Supprime un contact",
	Long:    `Supprime un contact du CRM en spécifiant son ID.`,
	Example: "mini-crm delete --id=1",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")

		if id == 0 {
			fmt.Println("Erreur : Le flag --id est requis")
			fmt.Println("Usage: mini-crm delete --id=1")
			os.Exit(1)
		}

		err := store.Remove(id)
		if err != nil {
			fmt.Printf("Erreur : %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Contact supprimé avec succès")
	},
}

func init() {
	// Flags pour la commande add
	addCmd.Flags().StringP("name", "n", "", "Nom du contact (requis)")
	addCmd.Flags().StringP("email", "e", "", "Email du contact (requis)")
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("email")

	// Flags pour la commande update
	updateCmd.Flags().IntP("id", "i", 0, "ID du contact à mettre à jour (requis)")
	updateCmd.Flags().StringP("name", "n", "", "Nouveau nom du contact")
	updateCmd.Flags().StringP("email", "e", "", "Nouvel email du contact")
	updateCmd.MarkFlagRequired("id")

	// Flags pour la commande delete
	deleteCmd.Flags().IntP("id", "i", 0, "ID du contact à supprimer (requis)")
	deleteCmd.MarkFlagRequired("id")
}

// Fonction pour le mode interactif (copiée et adaptée depuis app.go)
func runInteractiveMode(store storage.Storer) {
	for {
		showMenu()
		fmt.Print("Choisissez une option : ")
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			app.HandleAddContact(store)
		case "2":
			app.HandleListContacts(store)
		case "3":
			app.HandleRemoveContact(store)
		case "4":
			app.HandleUpdateContact(store)
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

// Commande config pour afficher les informations de configuration
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Affiche les informations de configuration",
	Long:  `Affiche les informations sur la configuration actuelle de l'application, notamment le type de stockage utilisé.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("=== Configuration du Mini-CRM ===")
		fmt.Println(config.GetConfigInfo(cfg))
		fmt.Println()
		fmt.Println("Pour changer la configuration, modifiez le fichier config.yaml")
		fmt.Println("Types de stockage disponibles : memory, json, gorm")
	},
}
