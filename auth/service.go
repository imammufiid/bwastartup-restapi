package auth

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	
}