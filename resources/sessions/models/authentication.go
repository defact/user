package sessions

type Authentication struct {
	Authentication string `json:"authentication"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Token          string `json:"token"`
}
