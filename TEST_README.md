# Scripts de Test - Mini-CRM

Ce dossier contient plusieurs scripts pour tester toutes les fonctionnalités du Mini-CRM de manière automatisée.

## Scripts disponibles

### 🧪 `test_all.sh` - Tests complets
Script principal qui teste toutes les fonctionnalités avec tous les types de stockage.

**Fonctionnalités testées :**
- ✅ Construction du binaire
- ✅ Tests CRUD complets (Create, Read, Update, Delete)
- ✅ Tous les types de stockage (memory, json, gorm)
- ✅ Gestion des erreurs et cas limites
- ✅ Tests de performance basiques
- ✅ Validation des arguments
- ✅ Restauration automatique de la configuration

**Usage :**
```bash
./test_all.sh
```

**Résultats :**
- Code de sortie 0 : Tous les tests passés ✅
- Code de sortie 1 : Au moins un test échoué ❌
- Rapport détaillé avec compteur de réussite/échec

### ⚡ `test_quick.sh` - Tests rapides
Script de test rapide qui utilise la configuration actuelle sans la modifier.

**Fonctionnalités testées :**
- ✅ Tests CRUD basiques
- ✅ Utilise la configuration existante
- ✅ Pas de modification de fichiers

**Usage :**
```bash
./test_quick.sh
```

### 🚨 `test_errors.sh` - Tests de robustesse
Script spécialisé dans les tests de gestion d'erreur et de validation.

**Fonctionnalités testées :**
- ✅ Arguments manquants
- ✅ IDs invalides (inexistants, négatifs, non-numériques)
- ✅ Emails invalides
- ✅ Noms invalides
- ✅ Commandes inexistantes
- ✅ Flags invalides
- ✅ Tests de validation positive

**Usage :**
```bash
./test_errors.sh
```

### 🧹 `cleanup.sh` - Nettoyage
Script pour nettoyer tous les fichiers temporaires créés pendant les tests.

**Fichiers nettoyés :**
- Fichiers de base de données de test (`contacts_test.db`)
- Fichiers JSON de test (`contacts_test.json`)
- Sauvegardes de configuration (`config.yaml.backup`)
- Binaires compilés (`mini-crm`, `go-mini-crm`)
- Fichiers de test Go (`*.test`)

**Usage :**
```bash
./cleanup.sh           # Nettoyage standard
./cleanup.sh --full    # Nettoyage complet avec cache Go
```

## Exemples d'utilisation

### Test complet avant livraison
```bash
# Lancer tous les tests
./test_all.sh

# Si tous les tests passent, nettoyer
./cleanup.sh
```

### Développement rapide
```bash
# Tests rapides pendant le développement
./test_quick.sh

# Tests d'erreur spécifiques
./test_errors.sh
```

### Intégration continue
```bash
#!/bin/bash
# Script CI/CD
set -e

echo "Lancement des tests complets..."
./test_all.sh

echo "Tests de robustesse..."
./test_errors.sh

echo "Nettoyage..."
./cleanup.sh

echo "✅ Tous les tests sont passés !"
```

## Configuration des tests

Les scripts `test_all.sh` créent temporairement une configuration de test qui :
- Utilise `contacts_test.db` au lieu de `contacts.db`
- Utilise `contacts_test.json` au lieu de `contacts.json`
- Sauvegarde et restaure automatiquement la configuration originale

## Types de stockage testés

### Memory Storage
- ✅ Données en mémoire uniquement
- ✅ Rapide pour les tests
- ✅ Aucun fichier créé

### JSON Storage
- ✅ Persistence dans un fichier JSON
- ✅ Format lisible
- ✅ Bon pour le debug

### GORM Storage
- ✅ Base de données SQLite
- ✅ Performance optimale
- ✅ Robustesse en production

## Sorties des tests

### Format des logs
```
[INFO] Message informatif
[SUCCESS] ✅ Test réussi
[ERROR] ❌ Test échoué  
[WARNING] ⚠️  Avertissement
```

### Rapport final
```
================================
       RÉSULTATS DES TESTS       
================================
Tests réussis: 25/25
✅ Tous les tests sont passés !
```

## Gestion des erreurs

Tous les scripts incluent :
- ✅ Gestion des interruptions (Ctrl+C)
- ✅ Nettoyage automatique en cas d'erreur
- ✅ Restauration de la configuration originale
- ✅ Codes de sortie appropriés pour l'intégration CI/CD

## Dépendances

Les scripts nécessitent :
- Bash (compatible fish shell)
- Go compilateur
- Permissions d'écriture dans le répertoire de travail

## Personnalisation

### Modifier les contacts de test
Éditez les variables dans `test_all.sh` :
```bash
# Exemple dans la fonction test_add_contact
test_add_contact "Votre Nom" "votre@email.com" "success"
```

### Changer les types de stockage testés
Modifiez le tableau dans `test_all.sh` :
```bash
local storage_types=("memory" "json" "gorm")
# Ou seulement certains types :
local storage_types=("gorm")
```

### Ajuster les tests de performance
Modifiez le nombre de contacts dans `test_all.sh` :
```bash
for i in $(seq 1 100); do  # Changer 100 pour un autre nombre
```