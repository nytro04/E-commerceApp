package users

// user properties
type User struct {
	ID      uint64
	Name    string
	Hash    []byte
	Address string
}



func CreateUser(name, password, address string) (*User, error) {

}