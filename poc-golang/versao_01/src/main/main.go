// Todo codigo em Go precisa iniciar com este package
package main
 
 // Importa todas as bibliotecas que serão utilizadas no código
import (

    // bibliotecas Nativas do Golang
    "encoding/json"
    "strconv" 
    "log"
    "net/http"
	"fmt"
    "database/sql"
    "io/ioutil"
    "bytes"

    // Pode ser usado bibliotecas não nativas 
    // /gorilla/mux para o servidor HTTP
    // /denisenkom/go-mssqldb para a comunicação com Banco SQLServer
    "github.com/gorilla/mux"
    _ "github.com/denisenkom/go-mssqldb"
)
 
// Struct Principal criada para armazenar os dados retornados do Banco de dados
// Structs são mais comuns para converter para JSON em Golang
type Usuario struct {

    // "omitempty", quando usado, omite este atributo caso seja 'null'
    ID      string  `json:"id,omitempty"`
    Nome    string  `json:"nome,omitempty"`
    Email   string  `json:"email,omitempty"`
    Senha   string  `json:"senha,omitempty"`
}

type CategoriaFull struct{
    ID      string  `json:"id,omitempty"`
    Nome    string  `json:"nome,omitempty"`
    Total   string  `json:"total,omitempty"`
}

type CategoriaEach struct{
    ID      string `json:"id,omitempty"`
    Nome    string `json:"nome,omitempty"`
    Regiao  string `json:"regiao,omitempty"`
    Total   string `json:"total,omitempty"`
}

type NovaDenuncia struct{
    Categoria string    `json:"categoria,omitempty"`
    Localidade string   `json:"localidade,omitempty"`
}

/* VARIAVEIS GLOBAIS PODEM SER USADAS EM QUALQUER PARTE DO CÓDiGO*/
// array para salvar os Usuarios
var cadastros []Usuario 
// array usado para enviar o total de cada categoria
var categorias []CategoriaFull
// array usado para enviar o total de denuncias por regiao
var categoriasRegiao []CategoriaEach
// Usado para armazenar o ultimo 'id' do banco de dados
var countBanco int 
var countCategoria int 
/*
    @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
    BD/BD.GO e BD/CONEXAO.GO
    @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
*/
func AtualizaCategorias() {
    
    // Inicia conexão com Banco de dados no Azure 
    db, err := sql.Open("mssql", "server=pwbt.database.windows.net;user id=admin-jose;password=12ab!@;database=PWBT;port=1433")
    // Caso ocorrer algum erro na conexão, o ERRO será salvo em 'err'
    // Caso não houver ERRO o retorno será vazio 'nil'
    if err != nil {
        log.Println("Erro ao conectar com o Banco de dados:", err.Error())
    }

    // Select traz todos os Nomes das categorias e as quantidades de denuncias de cada um
    rowsFull, err := db.Query("SELECT d.id_categoria, c.categoria, COUNT(d.id_categoria) FROM tab_denuncia d JOIN tab_categoria c ON d.id_categoria = c.id GROUP BY d.id_categoria, c.categoria") 
    if err != nil {
        log.Println("Erro no SELECT das CategoriasFull:", err.Error())
    }

    // select usado para trazer a Categoria e a quantidade de denuncias por regia
    rowsEach, err := db.Query("SELECT d.id_categoria, c.categoria, l.regiao, COUNT(d.id_localidade) FROM tab_denuncia d JOIN tab_categoria c ON d.id_categoria = c.id JOIN tab_localidade l ON d.id_localidade = l.id GROUP BY d.id_localidade, d.id_categoria, c.categoria, l.regiao") 
    if err != nil {
        log.Println("Erro no SELECT das CategoriasCada:", err.Error())
    }
    // select apenas para trazer o valor do ultimo ID
    rowsCount, err := db.Query("select MAX(id) from tab_denuncia") 
    if err != nil {
        log.Println("Erro no SELECT count categoria:", err.Error())
    }

    // finaliza o comando para Query
    defer rowsFull.Close() 
    defer rowsEach.Close()
    defer rowsCount.Close()
    // fecha conexão com o Banco
    defer db.Close()   
    
    // Zera o Array antes de buscar novos dados no banco.
    // Dessa forma evita dados repetidos
    categorias = categorias[:0]
    categoriasRegiao = categoriasRegiao[:0]

    // rowsFull.Next usado para varrer o objeto 'rowsFull' e pegar os valores retornados da Query
    for rowsFull.Next() {
        // Struct criada para receber os dados do banco de dados
        // Uma Struct permite receber dados de diferentes tipos: String, Int ...
        addCategoria := CategoriaFull{} 

        // rowsFull.Scan varre o objeto rowsFull e salva os valores na variaveis citadas abaixo.
        // As variaveis são salvas na mesma ordem que são coletados do Banco de dados
        if err := rowsFull.Scan(&addCategoria.ID, &addCategoria.Nome, &addCategoria.Total); err != nil { 
            log.Println("Erro ao salvar as categoriasFull retornados do Banco:", err.Error())
        }
        
        // adiciona na struct principal os dados do banco
        // Sera está struct 'cadastros' que será convertida para o formato JSON
        categorias = append(categorias, addCategoria)

        // Pega o valor do ID do ultimo dado buscado no banco e converte de String para Int
    }

    for rowsEach.Next(){
        addCategoria := CategoriaEach{}
        if err := rowsEach.Scan(&addCategoria.ID, &addCategoria.Nome, &addCategoria.Regiao, &addCategoria.Total); err != nil{
            log.Println("Erro ao salvar as categoriasEach retornados do Banco:", err.Error())
        } 

        categoriasRegiao = append(categoriasRegiao, addCategoria)
    }

    for rowsCount.Next(){
        
        if err := rowsCount.Scan(&countCategoria); err != nil{
            log.Println("Erro ao salvar categoriasCount retornados do Banco:", err.Error())
        } 
    }

    log.Printf("Categorias atualizadas!")
}
/*
	@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
	REST/REST.GO
	@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
*/
func GetCadastros(w http.ResponseWriter, req *http.Request) {
    log.Printf("Get cadastros") 
    json.NewEncoder(w).Encode(cadastros) //correto cadastros
}

// Adicona mais uma denuncia
func PostDenuncia(w http.ResponseWriter, req *http.Request){
    // modelo que deve enviado
    // {"categoria":"4","localidade":"2"}
    log.Printf("Post Nova Denuncia") 
    var novaD NovaDenuncia
    
    // grava em 'novaD' os dados enviados
    _ = json.NewDecoder(req.Body).Decode(&novaD)
    // imprime no terminal os valores recebidos
    fmt.Println(novaD)
    // usado para sempre o numero do 'id' ser id+1
    countCategoria++ 

    db, err := sql.Open("mssql", "server=pwbt.database.windows.net;user id=admin-jose;password=12ab!@;database=PWBT;port=1433")
    if err != nil {
        log.Println("Erro ao conectar com o Banco de dados:", err.Error())
    }

    rows, err := db.Query("INSERT into tab_denuncia (id, id_categoria, id_localidade) VALUES (?1, ?2, ?3)", countCategoria, novaD.Categoria, novaD.Localidade)
    if err != nil {
        log.Println("Erro no INSERT:", err.Error())
    }
    defer rows.Close() // fecha o comando Query
    defer db.Close()   // fecha conexão com o Banco
    // atualiza struct no banco
    AtualizaCategorias() 
    
    //json.NewEncoder(w).Encode(categorias)
}

// função para enviar apenas uma categoria com o total por regiao
func GetUMACategoria(w http.ResponseWriter, req *http.Request) {
    // OBSERVAÇÂO: comentarios de como funciona esta na 'func GetUsuario'
    log.Printf("Get uma Categoria") 
    params := mux.Vars(req) 
    var categoriaFound []CategoriaEach
    for _, item := range categoriasRegiao { 
        if item.ID == params["id"] { 
            categoriaFound = append(categoriaFound,item)
        }
    }
    json.NewEncoder(w).Encode(categoriasRegiao) 
    //json.NewEncoder(w).Encode(categoriaFound) 
}

// envia os dados das categorias via GET
func GetCategorias(w http.ResponseWriter, req *http.Request){
    log.Printf("Get categorias") 
    json.NewEncoder(w).Encode(categorias)
}
/*
	@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
	REST/SERVIDOR.GO
	@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
*/
func main() {

    AtualizaUsuarios()   
    AtualizaCategorias() 
    router := mux.NewRouter()

    // Caso queira trabalhar com cadastros
    router.HandleFunc("/cadastros/", GetCadastros).Methods("GET") // JSON com todos os cadastros
    router.HandleFunc("/cadastros/{id}", GetUsuario).Methods("GET") // JSON com apenas 1 cadastro
    router.HandleFunc("/cadastros/", PostUsuario).Methods("POST") // JSON para receber o um novo usuario
    router.HandleFunc("/cadastros/{id}", DeleteUsuario).Methods("DELETE") // JSON para deletar um usuario

    router.HandleFunc("/categorias/", GetCategorias).Methods("GET") // JSON com todas as categorias
    router.HandleFunc("/categorias/{id}", GetUMACategoria).Methods("GET") // devolve apenas uma categoria
    router.HandleFunc("/categorias/", PostDenuncia).Methods("POST") // adiciona nova categoria

    // Não retorna nem envia nada, apenas atualiza o arquivo tabela.json aqui no server [util para teste]
    router.HandleFunc("/categorias/api/{id}", CriaArquivoJSON).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", router)) // Server na porta 8080 [ localhost:8080 ]
}
