package models

import (
	"api/src/security"
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

// User representa um usuário do sistema
type User struct {
	ID       uint64 `json:"id,omitempty"` // ID no banco de dados - esse dado não vai ser visível para o usuário
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Active   bool   `json:"active,omitempty"` // indica se o usuário está ativo ou não - esse dado não vai ser visível para o usuário
}

// Prepare vai chamar os métodos para formatar e validar os dados do usuário. Argumentos: 0-cadastro, 1-edição
func (user *User) Prepare(prepareType int16) error {

	// retirando os espaços em branco
	user.format()

	// validando os dados
	if err := user.validate(prepareType); err != nil {
		return err
	}

	// colocando hash na senha
	if prepareType == 0 { // se está cadastrando
		if err := user.putHashInPassword(); err != nil {
			return err
		}
	}

	return nil
}

// Argumentos: 0-cadastro, 1-edição
func (user *User) validate(prepareType int16) error {

	// verificando se o nome está em branco
	if user.Name == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}

	// verificando se o e-mail está em branco
	if user.Email == "" {
		return errors.New("o E-mail é obrigatório e não pode estar em branco")
	}

	// verificando se o e-mail é valid
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("o E-mail inserido é invalido")
	}

	// verificando se a senha está em branco
	if prepareType == 0 && user.Password == "" { // prepareType = 0: cadastro, 1:edição
		return errors.New("a senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
}

func (user *User) putHashInPassword() error {
	// criando o hash
	passwordWithHash, err := security.Hash(user.Password)
	if err != nil {
		return err
	}

	user.Password = string(passwordWithHash)

	return nil
}
