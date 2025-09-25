package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Config représente la configuration de l'application
type Config struct {
	Storage StorageConfig `yaml:"storage"`
}

// StorageConfig contient la configuration du stockage
type StorageConfig struct {
	Type     string `yaml:"type"`     // "memory", "json", ou "gorm"
	JSONFile string `yaml:"jsonFile"` // Chemin du fichier JSON (pour le type "json")
	DBPath   string `yaml:"dbPath"`   // Chemin de la base de données SQLite (pour le type "gorm")
}

// LoadConfig charge la configuration depuis le fichier config.yaml
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// Valeurs par défaut
	viper.SetDefault("storage.type", "json")
	viper.SetDefault("storage.jsonFile", "contacts.json")
	viper.SetDefault("storage.dbPath", "contacts.db")

	// Essayer de lire le fichier de configuration
	if err := viper.ReadInConfig(); err != nil {
		// Si le fichier n'existe pas, créer un fichier de configuration par défaut
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("Fichier de configuration non trouvé, création du fichier par défaut...")
			if err := createDefaultConfig(); err != nil {
				return nil, fmt.Errorf("impossible de créer le fichier de configuration par défaut : %w", err)
			}
		} else {
			return nil, fmt.Errorf("erreur lors de la lecture du fichier de configuration : %w", err)
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("impossible de désérialiser la configuration : %w", err)
	}

	// Valider la configuration
	if err := validateConfig(&config); err != nil {
		return nil, fmt.Errorf("configuration invalide : %w", err)
	}

	log.Printf("Configuration chargée : type de stockage = %s", config.Storage.Type)
	return &config, nil
}

// createDefaultConfig crée un fichier de configuration par défaut
func createDefaultConfig() error {
	viper.Set("storage.type", "json")
	viper.Set("storage.jsonFile", "contacts.json")
	viper.Set("storage.dbPath", "contacts.db")

	return viper.WriteConfigAs("config.yaml")
}

// validateConfig valide la configuration
func validateConfig(config *Config) error {
	validTypes := []string{"memory", "json", "gorm"}
	isValid := false
	for _, validType := range validTypes {
		if config.Storage.Type == validType {
			isValid = true
			break
		}
	}

	if !isValid {
		return fmt.Errorf("type de stockage invalide : %s. Types valides : %v", config.Storage.Type, validTypes)
	}

	return nil
}

// GetConfigInfo retourne des informations sur la configuration actuelle
func GetConfigInfo(config *Config) string {
	info := fmt.Sprintf("Type de stockage : %s\n", config.Storage.Type)

	switch config.Storage.Type {
	case "json":
		info += fmt.Sprintf("Fichier JSON : %s", config.Storage.JSONFile)
	case "gorm":
		info += fmt.Sprintf("Base de données SQLite : %s", config.Storage.DBPath)
	case "memory":
		info += "Stockage en mémoire (données perdues à la fermeture)"
	}

	return info
}
