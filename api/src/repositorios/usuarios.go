package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

// Usuarios representa um repositorio de usuários
type Usuarios struct {
	db *sql.DB
}

// Função NovoRepositorioDeUsuarios cria um repositorio de usuarios
func NovoRespositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuario no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into usuarios (nome, nick, email, senha) values(?,?,?,?)")
	if erro != nil {
		return 0, nil
	}

	defer statement.Close()
	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, nil
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, nil
	}

	fmt.Println(ultimoIDInserido)

	return uint64(ultimoIDInserido), nil

}
