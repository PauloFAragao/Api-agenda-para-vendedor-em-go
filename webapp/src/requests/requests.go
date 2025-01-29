package requests

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

func MakeRequestWithAuthentication(r *http.Request, method, url string, data io.Reader) (*http.Response, error) {
	// criando a requisição
	request, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	// lendo o cookie
	cookie, _ := cookies.Read(r)

	// adicionando o token ao request
	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	// criando o client
	client := &http.Client{}

	// fazendo a requisição
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
