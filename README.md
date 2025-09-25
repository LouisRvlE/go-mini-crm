# Mini-CRM CLI# Mini-CRM



Un gestionnaire de contacts simple et efficace en ligne de commande, écrit en Go. Ce projet illustre les bonnes pratiques de développement Go avec une architecture modulaire, une CLI professionnelle et plusieurs options de persistance.Un système de gestion de contacts simple écrit en Go.



## 🚀 Fonctionnalités## Description



- ✅ **Gestion complète des contacts (CRUD)** : Ajouter, Lister, Mettre à jour et Supprimer des contactsCe projet est un mini-CRM (Customer Relationship Management) en ligne de commande qui permet de gérer une liste de contacts. Il offre des fonctionnalités de base pour ajouter, lister, modifier et supprimer des contacts.

- ✅ **Interface en ligne de commande professionnelle** : Commandes et sous-commandes avec Cobra

- ✅ **Configuration externe** : Gestion via fichier YAML avec Viper## Fonctionnalités

- ✅ **Persistance multiple** : Support de 3 backends de stockage :

  - **GORM/SQLite** : Base de données SQL robuste dans un fichier-   ✅ Ajouter un contact (nom et email)

  - **JSON** : Sauvegarde simple et lisible-   ✅ Lister tous les contacts

  - **Mémoire** : Stockage éphémère pour tests-   ✅ Mettre à jour un contact existant

- ✅ **Architecture découplée** : Packages séparés avec interfaces-   ✅ Supprimer un contact

- ✅ **Injection de dépendances** : Via les interfaces Go-   ✅ Interface en ligne de commande interactive

-   ✅ Ajout de contact via flags de commande

## 📋 Prérequis

## Prérequis

- Go 1.25.1 ou supérieur

- Git (pour le clonage)-   Go 1.25.1 ou supérieur



## 🔧 Installation## Installation



1. **Clonez le repository** :1. Clonez le repository :

```bash

git clone https://github.com/LouisRvlE/go-mini-crm```bash

cd go-mini-crmgit clone https://github.com/LouisRvlE/go-mini-crm

```cd go-mini-crm

```

2. **Installez les dépendances** :

```bash2. Compilez le programme :

go mod tidy

``````bash

go build main.go

3. **Compilez l'application** :```

```bash

go build -o mini-crm .## Utilisation

```

Lancez le programme sans arguments pour utiliser le menu interactif :

## ⚙️ Configuration

```bash

L'application utilise un fichier `config.yaml` pour sa configuration. Si le fichier n'existe pas, il sera créé automatiquement au premier lancement.go run main.go

```

### Fichier config.yaml

Le programme affichera un menu avec les options suivantes :

```yaml

# Configuration du Mini-CRM1. Ajouter un contact

storage:2. Lister tous les contacts

  # Types de stockage : "memory", "json", "gorm"3. Supprimer un contact

  type: "gorm"4. Mettre à jour un contact

  5. Quitter

  # Fichier JSON (utilisé si type="json")

  jsonFile: "contacts.json"Vous pouvez également ajouter un contact directement via les flags :

  

  # Base de données SQLite (utilisée si type="gorm")```bash

  dbPath: "contacts.db"go run main.go -add -name="Michel" -email="michel@g.com"

``````



### Types de stockage disponibles#### Flags disponibles



1. **`gorm`** *(recommandé)* : Base de données SQLite via GORM-   `-add` : Active le mode d'ajout direct

   - Persistance robuste-   `-name` : Nom du contact (obligatoire avec -add)

   - Transactions ACID-   `-email` : Email du contact (obligatoire avec -add)

   - Contraintes d'intégrité (emails uniques)

## Structure du projet

2. **`json`** : Fichier JSON simple

   - Données lisibles```

   - Facile à éditer manuellement.

   - Bonne pour le développement├── main.go        # Point d'entrée principal et gestion des flags

├── contact.go     # Définition de la structure Contact et ses méthodes

3. **`memory`** : Stockage en mémoire├── store.go       # Interface Storer et implémentation MemoryStore

   - Très rapide├── handlers.go    # Fonctions de gestion des interactions utilisateur

   - Données perdues à l'arrêt├── go.mod         # Module Go

   - Parfait pour les tests└── README.md      # Documentation

```

## 💻 Utilisation

## Architecture

### Commandes disponibles

Le projet est maintenant organisé en modules séparés pour une meilleure maintenabilité :

```bash

# Afficher l'aide générale-   **`main.go`** : Point d'entrée de l'application, gestion des arguments de ligne de commande

./mini-crm help-   **`contact.go`** : Définit la structure `Contact` avec ses méthodes pour la gestion des données de contact

-   **`store.go`** : Contient l'interface `Storer` et l'implémentation `MemoryStore` pour la persistance des données

# Voir la configuration actuelle-   **`handlers.go`** : Regroupe toutes les fonctions de gestion des interactions utilisateur (menu, ajout, suppression, etc.)

./mini-crm config

### Nouvelles fonctionnalités

# Mode interactif (recommandé pour débuter)

./mini-crm interactive-   ✅ **Architecture modulaire** : Code organisé en fichiers séparés par responsabilité

-   ✅ **Méthodes getter** : Nouvelles méthodes `GetID()`, `GetName()`, `GetEmail()` pour accéder aux propriétés des contacts

# Ajouter un contact-   ✅ **Fonction de création** : `NewContact()` pour créer des contacts de manière uniforme

./mini-crm add --name="John Doe" --email="john@example.com"-   ✅ **Séparation des préoccupations** : Interface utilisateur séparée de la logique métier


# Lister tous les contacts
./mini-crm list

# Mettre à jour un contact
./mini-crm update --id=1 --name="Jane Doe" --email="jane@example.com"

# Supprimer un contact
./mini-crm delete --id=1
```

### Mode interactif

Le mode interactif offre une interface simple avec menu :

```bash
./mini-crm interactive
```

```
--- [mini-CRM] ---
1. Ajouter un contact
2. Lister tous les contacts
3. Supprimer un contact
4. Mettre à jour un contact
5. Quitter
```

### Exemples d'utilisation

```bash
# Configuration initiale (vérifier le type de stockage)
./mini-crm config

# Ajouter quelques contacts
./mini-crm add --name="Alice Martin" --email="alice@example.com"
./mini-crm add --name="Bob Dupont" --email="bob@example.com"
./mini-crm add --name="Claire Durand" --email="claire@example.com"

# Lister tous les contacts
./mini-crm list

# Mettre à jour un contact
./mini-crm update --id=2 --name="Robert Dupont"

# Supprimer un contact
./mini-crm delete --id=1
```

### Changer de type de stockage

1. Modifiez le fichier `config.yaml`
2. Changez la valeur `storage.type`
3. Relancez l'application

**Exemple** : Passer du stockage GORM au stockage JSON
```yaml
storage:
  type: "json"  # Changé de "gorm" à "json"
  jsonFile: "contacts.json"
  dbPath: "contacts.db"
```

## 🏗️ Architecture du projet

```
go-mini-crm/
├── main.go                    # Point d'entrée de l'application
├── config.yaml               # Fichier de configuration
├── go.mod                    # Dépendances Go
├── cmd/
│   └── root.go               # Commandes CLI avec Cobra
├── internal/
│   ├── app/
│   │   └── app.go            # Logique applicative
│   ├── config/
│   │   └── config.go         # Gestion de la configuration avec Viper
│   └── storage/
│       ├── storage.go        # Interface Storer
│       ├── contact.go        # Modèle Contact
│       ├── memory.go         # Stockage en mémoire
│       ├── json.go           # Stockage JSON
│       └── gorm.go           # Stockage GORM/SQLite
└── README.md                 # Cette documentation
```

### Composants clés

- **Interface `Storer`** : Contrat pour tous les backends de stockage
- **Modèle `Contact`** : Structure de données avec tags JSON et GORM
- **Configuration dynamique** : Sélection du backend via config.yaml
- **CLI professionnelle** : Commandes structurées avec Cobra
- **Gestion d'erreurs** : Retours d'erreurs idiomatiques Go

## 🧪 Tests et développement

### Test rapide de toutes les fonctionnalités

```bash
# 1. Tester avec le stockage GORM (par défaut)
./mini-crm config
./mini-crm add --name="Test User" --email="test@example.com"
./mini-crm list

# 2. Passer au stockage JSON
# Modifiez config.yaml : type: "json"
./mini-crm config
./mini-crm list  # Aucun contact (nouveau backend)

# 3. Passer au stockage mémoire
# Modifiez config.yaml : type: "memory"
./mini-crm config
./mini-crm list  # Aucun contact (stockage éphémère)
```

### Développement et extension

Pour ajouter un nouveau type de stockage :

1. Implémentez l'interface `storage.Storer`
2. Ajoutez le nouveau type dans `config/config.go`
3. Modifiez `cmd/root.go` pour instancier votre store
4. Mettez à jour la documentation

## 🔍 Détails techniques

### Dépendances principales

- **[Cobra](https://github.com/spf13/cobra)** : Framework CLI
- **[Viper](https://github.com/spf13/viper)** : Gestion de configuration
- **[GORM](https://gorm.io/)** : ORM pour Go
- **[SQLite Driver](https://github.com/mattn/go-sqlite3)** : Driver SQLite pour GORM

### Fonctionnalités avancées

- **Concurrence sûre** : Mutex dans JSONStore
- **Auto-migration** : Tables créées automatiquement avec GORM
- **Validation** : Emails uniques dans la base GORM
- **Gestion d'erreurs** : Messages d'erreur clairs
- **Configuration par défaut** : Valeurs sensées si config manquante

## 📝 Notes sur l'implémentation

### Partie 1 : GORM/SQLite (45%)
- ✅ Ajout des dépendances GORM et SQLite
- ✅ Mise à jour de la struct Contact avec tags GORM
- ✅ Implémentation complète de GORMStore
- ✅ Intégration dans cmd/root.go avec configuration

### Partie 2 : CLI Cobra & Viper (55%)
- ✅ Structure projet orientée Cobra
- ✅ Fichier de configuration YAML avec Viper
- ✅ Commande racine avec sélection dynamique du storage
- ✅ Toutes les sous-commandes (add, list, update, delete, config)
- ✅ Mode interactif préservé
- ✅ Documentation complète

## 🤝 Contribution

Ce projet est conçu comme un exemple pédagogique. Les améliorations suggérées :

- Tests unitaires
- Validation des emails
- Import/export de données
- Interface web (optionnelle)
- Docker support

## 📄 Licence

Projet éducatif - EFREI Paris

---

**Auteur** : Louis Revelle  
**Cours** : Développement Go - EFREI Paris  
**Date** : 2025