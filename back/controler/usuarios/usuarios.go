package usuarios

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

var (
	loginVal  login
	loginsVal []login
)

func init() {

	loginVal.Username = "teste"
	loginVal.Password = "teste"

	loginsVal = append(loginsVal, loginVal)

	loginVal.Username = "borodin"
	loginVal.Password = "teste"

	loginsVal = append(loginsVal, loginVal)
}

// AutorizaUsuario : zz
func AutorizaUsuario(usuario string, senha string) bool {
	for _, login := range loginsVal {
		if login.Username == usuario && login.Password == senha {
			return true
		}
	}

	return false
}
