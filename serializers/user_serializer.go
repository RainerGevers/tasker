package serializers

type UserSerializer struct {
	Id uint `json:"id"`
	Email string `json:"email"`
	Username string `json:"username"`
}