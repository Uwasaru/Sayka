package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/oklog/ulid"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"crypto/sha256"
	"math/rand"
	"os"
	"strconv"
	"time"
	"encoding/hex"
)

func Ulid() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}

func Uuid() string {
	uuidObj, _ := uuid.NewRandom()
  return uuidObj.String()
}

func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func ComparePassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

func GenerateToken(user_id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_id,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateVerificationToken() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(999999-100000) + 100000
	return strconv.Itoa(code)
}

func GetCurrentTimeStamps() time.Time {
	return time.Now().UTC()
}

func GetExpireTimeStamps() time.Time {
	return time.Now().AddDate(0, 1, 0).UTC()
}

func EncodeToken(token string) string {
	hasher := sha256.New()
	hasher.Write([]byte(token))
	encryptedBytes := hasher.Sum(nil)
	return hex.EncodeToString(encryptedBytes)
}