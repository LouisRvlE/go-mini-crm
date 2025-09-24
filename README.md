# Mini-CRM

Un système de gestion de contacts simple écrit en Go.

## Description

Ce projet est un mini-CRM (Customer Relationship Management) en ligne de commande qui permet de gérer une liste de contacts. Il offre des fonctionnalités de base pour ajouter, lister, modifier et supprimer des contacts.

## Fonctionnalités

-   ✅ Ajouter un contact (nom et email)
-   ✅ Lister tous les contacts
-   ✅ Mettre à jour un contact existant
-   ✅ Supprimer un contact
-   ✅ Interface en ligne de commande interactive
-   ✅ Ajout de contact via flags de commande

## Prérequis

-   Go 1.25.1 ou supérieur

## Installation

1. Clonez le repository :

```bash
git clone https://github.com/LouisRvlE/go-mini-crm
cd go-mini-crm
```

2. Compilez le programme :

```bash
go build main.go
```

## Utilisation

Lancez le programme sans arguments pour utiliser le menu interactif :

```bash
go run main.go
```

Le programme affichera un menu avec les options suivantes :

1. Ajouter un contact
2. Lister tous les contacts
3. Supprimer un contact
4. Mettre à jour un contact
5. Quitter

Vous pouvez également ajouter un contact directement via les flags :

```bash
go run main.go -add -name="Michel" -email="michel@g.com"
```

#### Flags disponibles

-   `-add` : Active le mode d'ajout direct
-   `-name` : Nom du contact (obligatoire avec -add)
-   `-email` : Email du contact (obligatoire avec -add)

## Structure du projet

```
.
├── main.go        # Point d'entrée principal et gestion des flags
├── contact.go     # Définition de la structure Contact et ses méthodes
├── store.go       # Interface Storer et implémentation MemoryStore
├── handlers.go    # Fonctions de gestion des interactions utilisateur
├── go.mod         # Module Go
└── README.md      # Documentation
```

## Architecture

Le projet est maintenant organisé en modules séparés pour une meilleure maintenabilité :

-   **`main.go`** : Point d'entrée de l'application, gestion des arguments de ligne de commande
-   **`contact.go`** : Définit la structure `Contact` avec ses méthodes pour la gestion des données de contact
-   **`store.go`** : Contient l'interface `Storer` et l'implémentation `MemoryStore` pour la persistance des données
-   **`handlers.go`** : Regroupe toutes les fonctions de gestion des interactions utilisateur (menu, ajout, suppression, etc.)

### Nouvelles fonctionnalités

-   ✅ **Architecture modulaire** : Code organisé en fichiers séparés par responsabilité
-   ✅ **Méthodes getter** : Nouvelles méthodes `GetID()`, `GetName()`, `GetEmail()` pour accéder aux propriétés des contacts
-   ✅ **Fonction de création** : `NewContact()` pour créer des contacts de manière uniforme
-   ✅ **Séparation des préoccupations** : Interface utilisateur séparée de la logique métier
