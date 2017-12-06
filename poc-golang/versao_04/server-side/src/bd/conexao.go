package bd

import (
	"database/sql"
	"fmt"
)

var (
	servidor = "pwbt.database.windows.net"
	usuario  = "admin-jose"
	senha    = "123abc!@#"
	banco    = "PWBT"
	porta    = 1433
)

var stringDeConexao = fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;port=%d", servidor, usuario, senha, banco, porta)

var bancoDeDados, _ = sql.Open("tipoBanco", "stringDeConexao")
var erro error
