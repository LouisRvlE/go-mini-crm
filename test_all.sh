#!/bin/bash

set -e  

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' 

BINARY_PATH="./mini-crm"
CONFIG_FILE="config.yaml"
CONFIG_BACKUP="config.yaml.backup"
TEST_DB="contacts.db"
TEST_JSON="contacts.json"
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
log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}
increment_test_counter() {
    ((TESTS_TOTAL++))
}

backup_config() {
    if [ -f "$CONFIG_FILE" ]; then
        cp "$CONFIG_FILE" "$CONFIG_BACKUP"
        log_info "Configuration sauvegardée"
    fi
}

restore_config() {
    if [ -f "$CONFIG_BACKUP" ]; then
        cp "$CONFIG_BACKUP" "$CONFIG_FILE"
        rm "$CONFIG_BACKUP"
        log_info "Configuration restaurée"
    fi
}

set_storage_type() {
    local storage_type=$1
    log_info "Configuration du stockage : $storage_type"
    
    
    cat > "$CONFIG_FILE" << EOF
    
storage:
  type: "$storage_type"
  jsonFile: "contacts_test.json"
  dbPath: "contacts_test.db"
EOF
}

cleanup_test_files() {
    local files=("contacts_test.db" "contacts_test.json" "contacts.db" "contacts.json")
    for file in "\${files[@]}"; do
        if [ -f "$file" ]; then
            rm "$file"
        fi
    done
}

build_binary() {
    log_info "Construction du binaire..."
    if go build -o "$BINARY_PATH" .; then
        log_success "Binaire construit avec succès"
        return 0
    else
        log_error "Échec de la construction du binaire"
        return 1
    fi
}

test_add_contact() {
    local name="$1"
    local email="$2"
    local expected_result="$3"
    
    increment_test_counter
    log_info "Test ajout contact: $name <$email>"
    
    if $BINARY_PATH add --name "$name" --email "$email" > /dev/null 2>&1; then
        if [ "$expected_result" = "success" ]; then
            log_success "Ajout réussi comme attendu"
        else
            log_error "Ajout réussi mais échec attendu"
            return 1
        fi
    else
        if [ "$expected_result" = "fail" ]; then
            log_success "Échec d'ajout comme attendu"
        else
            log_error "Échec d'ajout inattendu"
            return 1
        fi
    fi
    return 0
}

test_list_contacts() {
    increment_test_counter
    log_info "Test liste des contacts"
    
    if $BINARY_PATH list > /dev/null 2>&1; then
        log_success "Liste affichée avec succès"
        return 0
    else
        log_error "Échec d'affichage de la liste"
        return 1
    fi
}

test_update_contact() {
    local id="$1"
    local name="$2"
    local email="$3"
    local expected_result="$4"
    
    increment_test_counter
    log_info "Test mise à jour contact ID=$id"
    
    local cmd_args="--id $id"
    if [ -n "$name" ]; then
        cmd_args="$cmd_args --name \"$name\""
    fi
    if [ -n "$email" ]; then
        cmd_args="$cmd_args --email \"$email\""
    fi
    
    if eval "$BINARY_PATH update $cmd_args" > /dev/null 2>&1; then
        if [ "$expected_result" = "success" ]; then
            log_success "Mise à jour réussie comme attendu"
        else
            log_error "Mise à jour réussie mais échec attendu"
            return 1
        fi
    else
        if [ "$expected_result" = "fail" ]; then
            log_success "Échec de mise à jour comme attendu"
        else
            log_error "Échec de mise à jour inattendu"
            return 1
        fi
    fi
    return 0
}

test_delete_contact() {
    local id="$1"
    local expected_result="$2"
    
    increment_test_counter
    log_info "Test suppression contact ID=$id"
    
    if $BINARY_PATH delete --id "$id" > /dev/null 2>&1; then
        if [ "$expected_result" = "success" ]; then
            log_success "Suppression réussie comme attendu"
        else
            log_error "Suppression réussie mais échec attendu"
            return 1
        fi
    else
        if [ "$expected_result" = "fail" ]; then
            log_success "Échec de suppression comme attendu"
        else
            log_error "Échec de suppression inattendu"
            return 1
        fi
    fi
    return 0
}

test_invalid_commands() {
    log_info "=== Test des commandes invalides ==="
    
    increment_test_counter
    if $BINARY_PATH add --name "Test" > /dev/null 2>&1; then
        log_error "Ajout sans email devrait échouer"
    else
        log_success "Ajout sans email échoue comme attendu"
    fi
    
    increment_test_counter
    if $BINARY_PATH add --email "test@test.com" > /dev/null 2>&1; then
        log_error "Ajout sans nom devrait échouer"
    else
        log_success "Ajout sans nom échoue comme attendu"
    fi
    
    increment_test_counter
    if $BINARY_PATH update --name "Test" > /dev/null 2>&1; then
        log_error "Mise à jour sans ID devrait échouer"
    else
        log_success "Mise à jour sans ID échoue comme attendu"
    fi
    
    increment_test_counter
    if $BINARY_PATH delete > /dev/null 2>&1; then
        log_error "Suppression sans ID devrait échouer"
    else
        log_success "Suppression sans ID échoue comme attendu"
    fi
}

run_crud_tests() {
    local storage_type="$1"
    
    log_info "=== Tests CRUD pour le stockage: $storage_type ==="
    
    cleanup_test_files
    
    set_storage_type "$storage_type"

    test_add_contact "Alice Dupont" "alice@example.com" "success"
    test_add_contact "Bob Martin" "bob@example.com" "success"
    test_add_contact "Charlie Durand" "charlie@example.com" "success"
    
    test_list_contacts
    
    test_update_contact "1" "Alice DUPONT" "" "success"
    test_update_contact "2" "" "bob.martin@example.com" "success"
    test_update_contact "3" "Charles Durand" "charles@example.com" "success"
    
    test_delete_contact "2" "success"
    
    test_list_contacts
    
    test_update_contact "999" "Test" "test@test.com" "fail"
    test_delete_contact "999" "fail"
    
    log_info "Tests $storage_type terminés\n"
}

main() {
    echo -e "${BLUE}================================${NC}"
    echo -e "${BLUE}  MINI-CRM - Tests automatiques  ${NC}"
    echo -e "${BLUE}================================${NC}\\n"
    
    backup_config
    
    if ! build_binary; then
        log_error "Impossible de construire le binaire. Arrêt des tests."
        restore_config
        exit 1
    fi
    
    increment_test_counter
    if $BINARY_PATH --help > /dev/null 2>&1; then
        log_success "Binaire fonctionnel"
    else
        log_error "Le binaire ne fonctionne pas correctement"
        restore_config
        exit 1
    fi
    
    test_invalid_commands
    
    local storage_types=("memory" "json" "gorm")
    
    for storage_type in "${storage_types[@]}"; do
        run_crud_tests "$storage_type"
    done
    
    log_info "=== Tests de performance basiques ==="
    set_storage_type "gorm"
    cleanup_test_files
    
    increment_test_counter
    log_info "Ajout de 100 contacts en masse..."
    local start_time=$(date +%s)
    for i in $(seq 1 100); do
        $BINARY_PATH add --name "Contact$i" --email "contact$i@test.com" > /dev/null 2>&1
    done
    local end_time=$(date +%s)
    local duration=$((end_time - start_time))
    log_success "100 contacts ajoutés en ${duration}s"
    
    increment_test_counter
    if $BINARY_PATH list > /dev/null 2>&1; then
        log_success "Liste de 100 contacts affichée"
    else
        log_error "Échec d'affichage de la liste de 100 contacts"
    fi
    
    cleanup_test_files
    
    restore_config
    
    echo
    echo -e "${BLUE}================================${NC}"
    echo -e "${BLUE}       RÉSULTATS DES TESTS       ${NC}"
    echo -e "${BLUE}================================${NC}"
    echo -e "Tests réussis: ${GREEN}$TESTS_PASSED${NC}/$TESTS_TOTAL"
    
    if [ $TESTS_PASSED -eq $TESTS_TOTAL ]; then
        echo -e "${GREEN}✅ Tous les tests sont passés !${NC}"
        exit 0
    else
        local failed=$((TESTS_TOTAL - TESTS_PASSED))
        echo -e "${RED}❌ $failed test(s) échoué(s)${NC}"
        exit 1
    fi
}

trap 'log_warning "Tests interrompus"; cleanup_test_files; restore_config; exit 1' INT TERM

main "\$@"