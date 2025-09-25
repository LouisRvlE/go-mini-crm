#!/bin/bash

# Script de démonstration pour Mini-CRM
# Montre toutes les fonctionnalités de manière interactive

# Couleurs pour l'affichage
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m'

BINARY_PATH="./mini-crm"

log_demo() {
    echo -e "${CYAN}[DEMO]${NC} $1"
}

log_command() {
    echo -e "${YELLOW}$ $1${NC}"
}

pause_demo() {
    echo -e "${BLUE}Press Enter to continue...${NC}"
    read -r
}

# Construire le binaire si nécessaire
if [ ! -f "$BINARY_PATH" ]; then
    echo -e "${BLUE}Construction du binaire...${NC}"
    go build -o "$BINARY_PATH" .
fi

echo -e "${GREEN}================================${NC}"
echo -e "${GREEN}  MINI-CRM - Démonstration      ${NC}"
echo -e "${GREEN}================================${NC}"
echo

log_demo "Bienvenue dans la démonstration du Mini-CRM !"
log_demo "Ce script va vous montrer toutes les fonctionnalités disponibles."
echo
pause_demo

# 1. Afficher l'aide
log_demo "1. Affichage de l'aide générale"
log_command "mini-crm --help"
$BINARY_PATH --help
echo
pause_demo

# 2. Afficher les informations générales
log_demo "2. Informations générales et configuration"
log_command "mini-crm"
$BINARY_PATH
echo
pause_demo

# 3. Ajouter des contacts
log_demo "3. Ajout de nouveaux contacts"

log_command "mini-crm add --name \"Alice Dupont\" --email \"alice@example.com\""
$BINARY_PATH add --name "Alice Dupont" --email "alice@example.com"
echo

log_command "mini-crm add --name \"Bob Martin\" --email \"bob@example.com\""
$BINARY_PATH add --name "Bob Martin" --email "bob@example.com"
echo

log_command "mini-crm add --name \"Charlie Durand\" --email \"charlie@example.com\""
$BINARY_PATH add --name "Charlie Durand" --email "charlie@example.com"
echo
pause_demo

# 4. Lister tous les contacts
log_demo "4. Liste de tous les contacts"
log_command "mini-crm list"
$BINARY_PATH list
echo
pause_demo

# 5. Mettre à jour un contact
log_demo "5. Mise à jour d'un contact"

# Obtenir le premier ID disponible
FIRST_ID=$($BINARY_PATH list 2>/dev/null | grep "Id:" | head -1 | cut -d: -f2 | cut -d, -f1 | tr -d ' ')

if [ -n "$FIRST_ID" ]; then
    log_command "mini-crm update --id $FIRST_ID --name \"Alice DUPONT\" --email \"alice.dupont@example.com\""
    $BINARY_PATH update --id "$FIRST_ID" --name "Alice DUPONT" --email "alice.dupont@example.com"
    echo
    
    log_demo "Vérification de la mise à jour :"
    log_command "mini-crm list"
    $BINARY_PATH list
else
    log_demo "Aucun contact disponible pour la mise à jour"
fi
echo
pause_demo

# 6. Supprimer un contact
log_demo "6. Suppression d'un contact"

# Obtenir le dernier ID disponible
LAST_ID=$($BINARY_PATH list 2>/dev/null | grep "Id:" | tail -1 | cut -d: -f2 | cut -d, -f1 | tr -d ' ')

if [ -n "$LAST_ID" ]; then
    log_command "mini-crm delete --id $LAST_ID"
    $BINARY_PATH delete --id "$LAST_ID"
    echo
    
    log_demo "Vérification de la suppression :"
    log_command "mini-crm list"
    $BINARY_PATH list
else
    log_demo "Aucun contact disponible pour la suppression"
fi
echo
pause_demo

# 7. Démonstration des erreurs
log_demo "7. Gestion des erreurs"

log_demo "Tentative d'ajout sans email (doit échouer) :"
log_command "mini-crm add --name \"Test User\""
$BINARY_PATH add --name "Test User" 2>/dev/null && echo "Erreur : Devrait échouer" || echo -e "${GREEN}✅ Échec attendu - Validation OK${NC}"
echo

log_demo "Tentative de suppression avec ID inexistant (doit échouer) :"
log_command "mini-crm delete --id 99999"
$BINARY_PATH delete --id 99999 2>/dev/null && echo "Erreur : Devrait échouer" || echo -e "${GREEN}✅ Échec attendu - Validation OK${NC}"
echo
pause_demo

# 8. Test des différents types de stockage
log_demo "8. Démonstration des types de stockage"
log_demo "Le Mini-CRM supporte plusieurs types de stockage :"
echo -e "  ${GREEN}memory${NC} : Données en mémoire (perdues à la fermeture)"
echo -e "  ${GREEN}json${NC}   : Stockage dans un fichier JSON"
echo -e "  ${GREEN}gorm${NC}   : Base de données SQLite via GORM"
echo

log_demo "Configuration actuelle dans config.yaml :"
grep -A 10 "storage:" config.yaml | head -8
echo
pause_demo

# 9. Aide pour les commandes spécifiques
log_demo "9. Aide pour les commandes spécifiques"

log_command "mini-crm add --help"
$BINARY_PATH add --help
echo
pause_demo

log_command "mini-crm update --help"
$BINARY_PATH update --help
echo
pause_demo

# 10. État final
log_demo "10. État final du système"
log_command "mini-crm list"
$BINARY_PATH list
echo

# Afficher les fichiers créés
log_demo "Fichiers créés par l'application :"
ls -la *.db *.json 2>/dev/null | grep -E "\.(db|json)$" || echo "Aucun fichier de données visible (stockage en mémoire ou pas encore créé)"
echo

echo -e "${GREEN}================================${NC}"
echo -e "${GREEN}   FIN DE LA DÉMONSTRATION       ${NC}"
echo -e "${GREEN}================================${NC}"
echo
echo -e "${CYAN}Fonctionnalités démontrées :${NC}"
echo -e "  ✅ Ajout de contacts"
echo -e "  ✅ Liste des contacts"
echo -e "  ✅ Mise à jour de contacts"
echo -e "  ✅ Suppression de contacts"
echo -e "  ✅ Gestion des erreurs"
echo -e "  ✅ Types de stockage multiples"
echo -e "  ✅ Aide contextuelle"
echo
echo -e "${YELLOW}Pour nettoyer les données de test :${NC}"
echo -e "  ./cleanup.sh"
echo
echo -e "${YELLOW}Pour lancer les tests automatiques :${NC}"
echo -e "  ./test_all.sh    # Tests complets"
echo -e "  ./test_quick.sh  # Tests rapides"
echo -e "  ./test_errors.sh # Tests d'erreur"