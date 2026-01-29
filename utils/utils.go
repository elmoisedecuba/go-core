package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// Generate Username
func GetUsername(email string) string {
	// Dividir el string en partes usando el símbolo '@'
	username := strings.Split(email, "@")
	return username[0] // Si no hay '@', retorna el email completo
}

// Generate RandomHash
func RandomHash() string {
	newUUID := uuid.New()
	min := 111111111111111111
	max := 999999999999999999
	random := rand.Intn(max-min) + min
	randomString := strconv.Itoa(random)
	randomUUID := newUUID.String() + randomString
	randomhash := sha256.Sum256([]byte(randomUUID))
	hash := hex.EncodeToString(randomhash[:])
	return hash
}

// Encode String
func EncodeString(text string) string {
	textHash := sha256.Sum256([]byte(text))
	hash := hex.EncodeToString(textHash[:])
	return hash
}

// Encode UUID
func EncodeUUID(text string) string {
	textHash := md5.Sum([]byte(text))
	hash := hex.EncodeToString(textHash[:])
	return hash
}
