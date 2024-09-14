package interfaces

type HashingServiceInterface interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(hashedPassword string, password string) error
}
