package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

func InstanceService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("bWASt4rTUP_s3CREt_K3y")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	// generate payload jwt
	claim := jwt.MapClaims{}
	claim["user_id"] = userID // insert payload data

	// generate jwt algoritm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// generate secret key
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}