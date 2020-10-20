package models

//RespuestaLogin tiene el token que es devuelve con el login
type RespuestaLogin struct {
	Token string `json: "token,omitempty"`
}
