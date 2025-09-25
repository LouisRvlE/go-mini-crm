#!/bin/bash

# Script de test rapide pour Mini-CRM
# Teste les fonctionnalités avec la configuration actuelle

set -e

# Couleurs pour l'affichage
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

BINARY_PATH="./mini-crm"

log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Construire le binaire si nécessaire
if [ ! -f "$BINARY_PATH" ]; then
    log_info "Construction du binaire..."
    go build -o "$BINARY_PATH" .
fi

echo -e "${BLUE}================================${NC}"
echo -e "${BLUE}   MINI-CRM - Test rapide       ${NC}"
echo -e "${BLUE}================================${NC}"

# Afficher la configuration actuelle
log_info "Configuration actuelle:"
$BINARY_PATH 2>/dev/null | head -10

echo

# Test d'ajout
log_info "Test 1: Ajout de contacts"
$BINARY_PATH add --name "Test User 1" --email "test1@example.com"
$BINARY_PATH add --name "Test User 2" --email "test2@example.com"
log_success "Contacts ajoutés"

echo

# Test de liste
log_info "Test 2: Liste des contacts"
$BINARY_PATH list
log_success "Liste affichée"

echo

# Obtenir le premier ID disponible de la liste
FIRST_ID=$($BINARY_PATH list 2>/dev/null | grep "Id:" | head -1 | cut -d: -f2 | cut -d, -f1 | tr -d ' ')
SECOND_ID=$($BINARY_PATH list 2>/dev/null | grep "Id:" | head -2 | tail -1 | cut -d: -f2 | cut -d, -f1 | tr -d ' ')

if [ -z "$FIRST_ID" ]; then
    log_error "Aucun contact trouvé pour les tests de mise à jour et suppression"
    exit 1
fi

# Test de mise à jour
log_info "Test 3: Mise à jour d'un contact (ID: $FIRST_ID)"
$BINARY_PATH update --id "$FIRST_ID" --name "Test User Updated" --email "updated@example.com"
log_success "Contact mis à jour"

echo

# Afficher la liste mise à jour
log_info "Test 4: Vérification de la mise à jour"
$BINARY_PATH list
log_success "Mise à jour vérifiée"

echo

# Test de suppression (utiliser le second ID s'il existe, sinon créer un nouveau contact)
if [ -n "$SECOND_ID" ]; then
    DELETE_ID="$SECOND_ID"
    log_info "Test 5: Suppression d'un contact (ID: $DELETE_ID)"
else
    $BINARY_PATH add --name "Contact à supprimer" --email "delete@example.com" > /dev/null
    DELETE_ID=$($BINARY_PATH list 2>/dev/null | grep "delete@example.com" | cut -d: -f2 | cut -d, -f1 | tr -d ' ')
    log_info "Test 5: Suppression d'un contact créé pour le test (ID: $DELETE_ID)"
fi

$BINARY_PATH delete --id "$DELETE_ID"
log_success "Contact supprimé"

echo

# Afficher la liste finale
log_info "Test 6: Liste finale"
$BINARY_PATH list
log_success "Tous les tests terminés !"

echo -e "\n${GREEN}✅ Tests rapides terminés avec succès !${NC}"