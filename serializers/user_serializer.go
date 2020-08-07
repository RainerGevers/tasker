package serializers

type UserSerializer struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type LoginSerializer struct {
	Uuid         string `json:"uuid"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	AuthToken    string `json:"authToken"`
	RefreshToken string `json:"refreshToken"`
}
