package models

import "fmt"

type User struct {
	ID int         `json:"id"`
	Name string    `json:"name"`
	Email string   `json:"email"`
	Password string `json:"password"`
	//privado encapsulado  
}	

//Setter de contraseña con validación
func (u *User) SetPassword(pwd string) error {
	if len(pwd) < 8 {
		return fmt.Errorf("la contraseña debe tener al menos 8 caracteres")
	}
	u.Password = pwd
	return nil
}

//Getter validando la contraseña
func (u *User) CheckPassword(pwd string) bool {
	return u.Password == pwd
}
