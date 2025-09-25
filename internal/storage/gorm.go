package storage

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GORMStore struct {
	db        *gorm.DB
	idCounter int
}

// NewGORMStore crée une nouvelle instance de GORMStore avec une base de données SQLite
func NewGORMStore(dbPath string) *GORMStore {
	// Configuration GORM pour supprimer les logs verbeux
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Impossible de se connecter à la base de données:", err)
	}

	// Auto-migration : crée la table si elle n'existe pas
	err = db.AutoMigrate(&Contact{})
	if err != nil {
		log.Fatal("Impossible de migrer la base de données:", err)
	}

	store := &GORMStore{
		db: db,
	}

	// Initialiser le compteur d'ID
	store.initIDCounter()

	return store
}

// initIDCounter initialise le compteur d'ID basé sur le dernier ID dans la base
func (gs *GORMStore) initIDCounter() {
	var lastContact Contact
	result := gs.db.Last(&lastContact)
	if result.Error != nil {
		// Aucun contact trouvé, commencer à 1
		gs.idCounter = 1
	} else {
		gs.idCounter = lastContact.ID + 1
	}
}

// Add ajoute un nouveau contact à la base de données
func (gs *GORMStore) Add(contact *Contact) error {
	result := gs.db.Create(contact)
	if result.Error != nil {
		return fmt.Errorf("erreur lors de l'ajout du contact : %w", result.Error)
	}
	return nil
}

// GetAll récupère tous les contacts de la base de données
func (gs *GORMStore) GetAll() map[int]*Contact {
	var contacts []Contact
	gs.db.Find(&contacts)

	contactMap := make(map[int]*Contact)
	for i := range contacts {
		contactMap[contacts[i].ID] = &contacts[i]
	}

	return contactMap
}

// GetByID récupère un contact par son ID
func (gs *GORMStore) GetByID(id int) (*Contact, bool) {
	var contact Contact
	result := gs.db.First(&contact, id)
	if result.Error != nil {
		return nil, false
	}
	return &contact, true
}

// Remove supprime un contact par son ID
func (gs *GORMStore) Remove(id int) error {
	result := gs.db.Delete(&Contact{}, id)
	if result.Error != nil {
		return fmt.Errorf("erreur lors de la suppression : %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("contact avec l'ID %d n'existe pas", id)
	}
	return nil
}

// Update met à jour un contact existant
func (gs *GORMStore) Update(id int, name, email string) error {
	contact, exists := gs.GetByID(id)
	if !exists {
		return fmt.Errorf("contact avec l'ID %d n'existe pas", id)
	}

	// Mettre à jour seulement les champs non-vides
	updates := make(map[string]interface{})
	if name != "" {
		updates["name"] = name
	}
	if email != "" {
		updates["email"] = email
	}

	if len(updates) == 0 {
		return nil // Aucune mise à jour nécessaire
	}

	result := gs.db.Model(contact).Updates(updates)
	if result.Error != nil {
		return fmt.Errorf("erreur lors de la mise à jour : %w", result.Error)
	}

	return nil
}

// GetNextID retourne le prochain ID disponible
func (gs *GORMStore) GetNextID() int {
	id := gs.idCounter
	gs.idCounter++
	return id
}
