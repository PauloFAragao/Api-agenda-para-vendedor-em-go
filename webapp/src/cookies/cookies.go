package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Config utuliza as variáveis de ambiente para a criação do SecureCookie
func Config() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Save registra as informações de autenticação
func Save(w http.ResponseWriter, ID, token string) error {

	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	// codificando os dados
	encodedData, err := s.Encode("data", data)
	if err != nil {
		return err
	}

	// enviando o cookie para o navegador
	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    encodedData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

// Read retorna os valores armazenados no cookie
func Read(r *http.Request) (map[string]string, error) {

	// pegando dados do cookie
	cookie, err := r.Cookie("data")
	if err != nil {
		return nil, err
	}

	values := make(map[string]string)

	// descodificando os dados
	if err = s.Decode("data", cookie.Value, &values); err != nil {
		return nil, err
	}

	return values, nil

}
