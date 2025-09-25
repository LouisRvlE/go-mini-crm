#!/bin/bash

# Script de test pour les cas d'erreur du Mini-CRM
# Teste la robustesse de l'application face aux entrées invalides

# Couleurs pour l'affichage
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

BINARY_PATH="./mini-crm"
TESTS_PASSED=0
TESTS_TOTAL=0

log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
    ((TESTS_PASSED++))
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

log_test() {
    echo -e "${YELLOW}[TEST]${NC} $1"
    ((TESTS_TOTAL++))
}

# Test d'une commande qui doit échouer
test_should_fail() {
    local test_name="$1"
    local command="$2"
    
    log_test "$test_name"
    
    if eval "$command" > /dev/null 2>&1; then
        log_error "❌ La commande a réussi mais devrait échouer: $command"
        return 1
    else
        log_success "✅ La commande a échoué comme attendu"
        return 0
    fi
}

# Test d'une commande qui doit réussir
test_should_succeed() {
    local test_name="$1"
    local command="$2"
    
    log_test "$test_name"
    
    if eval "$command" > /dev/null 2>&1; then
        log_success "✅ La commande a réussi comme attendu"
        return 0
    else
        log_error "❌ La commande a échoué mais devrait réussir: $command"
        return 1
    fi
}

# Construire le binaire si nécessaire
if [ ! -f "$BINARY_PATH" ]; then
    log_info "Construction du binaire..."
    go build -o "$BINARY_PATH" .
fi

echo -e "${BLUE}================================${NC}"
echo -e "${BLUE}  MINI-CRM - Tests d'erreur     ${NC}"
echo -e "${BLUE}================================${NC}"

log_info "Test de la robustesse de l'application\n"

# Tests des arguments manquants
echo -e "${YELLOW}=== Tests d'arguments manquants ===${NC}"
test_should_fail "Ajout sans nom" "$BINARY_PATH add --email test@test.com"
test_should_fail "Ajout sans email" "$BINARY_PATH add --name 'Test User'"
test_should_fail "Ajout sans aucun argument" "$BINARY_PATH add"
test_should_fail "Mise à jour sans ID" "$BINARY_PATH update --name 'Test'"
test_should_fail "Suppression sans ID" "$BINARY_PATH delete"

echo

# Tests avec IDs invalides
echo -e "${YELLOW}=== Tests avec IDs invalides ===${NC}"
test_should_fail "Mise à jour avec ID inexistant" "$BINARY_PATH update --id 99999 --name 'Test'"
test_should_fail "Suppression avec ID inexistant" "$BINARY_PATH delete --id 99999"
test_should_fail "Mise à jour avec ID négatif" "$BINARY_PATH update --id -1 --name 'Test'"
test_should_fail "Suppression avec ID négatif" "$BINARY_PATH delete --id -1"
test_should_fail "Mise à jour avec ID non numérique" "$BINARY_PATH update --id abc --name 'Test'"

echo

# Tests avec des emails invalides (si validation implémentée)
echo -e "${YELLOW}=== Tests d'emails invalides ===${NC}"
test_should_fail "Email sans @" "$BINARY_PATH add --name 'Test' --email 'invalidemail'"
test_should_fail "Email sans domaine" "$BINARY_PATH add --name 'Test' --email 'test@'"
test_should_fail "Email vide" "$BINARY_PATH add --name 'Test' --email ''"

echo

# Tests avec des noms vides ou invalides
echo -e "${YELLOW}=== Tests de noms invalides ===${NC}"
test_should_fail "Nom vide" "$BINARY_PATH add --name '' --email 'test@test.com'"
test_should_fail "Nom avec seulement des espaces" "$BINARY_PATH add --name '   ' --email 'test@test.com'"

echo

# Tests de commandes inexistantes
echo -e "${YELLOW}=== Tests de commandes inexistantes ===${NC}"
test_should_fail "Commande inexistante" "$BINARY_PATH inexistante"
test_should_fail "Sous-commande inexistante" "$BINARY_PATH add inexistante"

echo

# Tests de flags invalides
echo -e "${YELLOW}=== Tests de flags invalides ===${NC}"
test_should_fail "Flag inexistant pour add" "$BINARY_PATH add --invalid-flag value"
test_should_fail "Flag inexistant pour list" "$BINARY_PATH list --invalid-flag"
test_should_fail "Flag inexistant pour update" "$BINARY_PATH update --invalid-flag value"
test_should_fail "Flag inexistant pour delete" "$BINARY_PATH delete --invalid-flag"

echo

# Tests positifs pour s'assurer que les commandes valides fonctionnent
echo -e "${YELLOW}=== Tests de validation positive ===${NC}"
test_should_succeed "Ajout valide" "$BINARY_PATH add --name 'Test Valid' --email 'valid@test.com'"
test_should_succeed "Liste" "$BINARY_PATH list"
test_should_succeed "Mise à jour valide" "$BINARY_PATH update --id 1 --name 'Test Updated'"
test_should_succeed "Aide générale" "$BINARY_PATH --help"
test_should_succeed "Aide pour add" "$BINARY_PATH add --help"

# Nettoyer le contact de test
$BINARY_PATH delete --id 1 > /dev/null 2>&1 || true

echo
echo -e "${BLUE}================================${NC}"
echo -e "${BLUE}         RÉSULTATS              ${NC}"
echo -e "${BLUE}================================${NC}"

echo -e "Tests réussis: ${GREEN}$TESTS_PASSED${NC}/$TESTS_TOTAL"

if [ $TESTS_PASSED -eq $TESTS_TOTAL ]; then
    echo -e "${GREEN}✅ Tous les tests d'erreur sont passés !${NC}"
    echo -e "${GREEN}L'application gère correctement les cas d'erreur.${NC}"
    exit 0
else
    local failed=$((TESTS_TOTAL - TESTS_PASSED))
    echo -e "${RED}❌ $failed test(s) d'erreur échoué(s)${NC}"
    echo -e "${RED}L'application ne gère pas correctement certains cas d'erreur.${NC}"
    exit 1
fi