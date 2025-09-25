#!/bin/bash

# Script de nettoyage pour Mini-CRM
# Supprime tous les fichiers de test et sauvegarde créés pendant les tests

# Couleurs pour l'affichage
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

# Fichiers à nettoyer
FILES_TO_CLEAN=(
    "contacts_test.db"
    "contacts_test.json"
    "config.yaml.backup"
    "mini-crm"
    "go-mini-crm"
    ".test_results.log"
)

# Dossiers à nettoyer (si vides)
DIRS_TO_CLEAN=(
    "test_output"
    "logs"
)

echo -e "${BLUE}================================${NC}"
echo -e "${BLUE}    MINI-CRM - Nettoyage        ${NC}"
echo -e "${BLUE}================================${NC}"

log_info "Début du nettoyage..."

# Nettoyer les fichiers
for file in "${FILES_TO_CLEAN[@]}"; do
    if [ -f "$file" ]; then
        rm "$file"
        log_success "Supprimé: $file"
    fi
done

# Nettoyer les dossiers vides
for dir in "${DIRS_TO_CLEAN[@]}"; do
    if [ -d "$dir" ] && [ -z "$(ls -A "$dir")" ]; then
        rmdir "$dir"
        log_success "Supprimé le dossier vide: $dir"
    fi
done

# Nettoyer les fichiers temporaires Go
if ls *.test 2>/dev/null; then
    rm *.test
    log_success "Fichiers de test Go supprimés"
fi

# Nettoyer le cache Go (optionnel)
if [ "$1" = "--full" ]; then
    log_info "Nettoyage complet du cache Go..."
    go clean -cache
    go clean -modcache 2>/dev/null || log_warning "Impossible de nettoyer le cache des modules"
    log_success "Cache Go nettoyé"
fi

log_success "Nettoyage terminé !"
echo
echo "Usage:"
echo "  ./cleanup.sh       - Nettoyage standard"
echo "  ./cleanup.sh --full - Nettoyage complet avec cache Go"