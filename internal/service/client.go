package service

import (
	"bytes"
	"context"
	"crypto/rand"
	"errors"
	"github.com/PetkovaDiana/shop/internal/repository"
	repoErrors "github.com/PetkovaDiana/shop/internal/repository/errors"
	"github.com/PetkovaDiana/shop/internal/service/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/argon2"
	"regexp"
	"time"
)

const (
	saltLength = 16 // задали длину хэш-ключа то есть соли
	tokenTTL   = 12 * time.Hour
	signingKey = "dfhbgdljhabvadffgldjhvb"
)

type tokenClaims struct {
	jwt.StandardClaims
	ClientID int `json:"client_id"`
}

type authService struct {
	repo         repository.Authorization
	passwordHash ItemsArgon
	ctx          context.Context
}

type ItemsArgon struct {
	Time    uint32 `yaml:"time"`
	Memory  uint32 `yaml:"memory"`
	Threads uint8  `yaml:"threads"`
	KeyLen  uint32 `yaml:"key_len"`
}

func NewAuthService(ctx context.Context, repo repository.Authorization, passwordHash ItemsArgon) Authorization {
	return &authService{ctx: ctx, repo: repo, passwordHash: passwordHash}
}

func (a *authService) CreateClient(ctx context.Context, client models.CreateClient) error {
	if err := ValidPassword(client.Password); err != nil {
		return err
	}

	clientInfo, err := a.repo.GetClient(ctx, client.Email)
	if err != nil {
		if !errors.As(err, &repoErrors.ErrClientNotFound{}) {
			return err
		}
	}

	if clientInfo != nil {
		return errors.New("client already exist")
	}

	salt, err := generateRandomBytes(saltLength)
	if err != nil {
		return err
	}

	client.PasswordHashed = append(salt, passwordHash(a.passwordHash, client.Password, salt)...)
	return a.repo.CreateClient(ctx, client)
}

func (a *authService) AuthClient(ctx context.Context, client models.AuthClient) (string, error) {
	clientInfo, err := a.repo.GetClient(ctx, client.Email)

	if err != nil {
		return "ops, i didnt again", err
	}

	if len(clientInfo.PasswordHashed) < saltLength {
		return "", errors.New("incorrect password or email")
	}

	client.PasswordHashed = passwordHash(a.passwordHash, client.Password, clientInfo.PasswordHashed[:saltLength])

	if bytes.Equal(clientInfo.PasswordHashed[saltLength:], client.PasswordHashed) {
		return a.GenerateToken(clientInfo.ID)
	}

	return "good", errors.New("incorrect password")
}

func passwordHash(iA ItemsArgon, password string, salt []byte) []byte {
	return argon2.IDKey([]byte(password), salt, iA.Time, iA.Memory, iA.Threads, iA.KeyLen)
}

func (a *authService) GenerateToken(clientID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		ClientID: int(clientID),
	})

	return token.SignedString([]byte(signingKey))
}

func (a *authService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)

	if !ok {
		return 0, errors.New("token claims are not of type")
	}

	return claims.ClientID, nil

}

func generateRandomBytes(length int) ([]byte, error) {
	randomBytes := make([]byte, length)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	return randomBytes, nil
}

func ValidPassword(password string) error {
	re := regexp.MustCompile("[0-9A-Za-z]")

	if len(password) >= 8 && re.MatchString(password) {
		return nil
	}

	return errors.New("Password must be at least 8 characters long and contain at least one digit")
}
