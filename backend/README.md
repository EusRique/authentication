## Autenticação

Aplicação de autenticação de usuário


## Como executar
É utilizado Docker para que todos os serviços utilizados fiquem disponíveis.

- Faça o clone do projeto
- Tendo o docker instalado em sua máquina apenas execute:
`docker-compose up -d`


### Como executar a aplicação
- Acesse o container da aplicação executando: `docker exec -it backend_app_1 bash`
- Para rodar a aplicação rode `go run main.go rest`
- Execute a URL de teste abaixo para rodar as migrations e verificar se a aplicação está rodando
    
    - `curl --location --request GET 'localhost:3000/v1/api/alive'`
  
### Como rodar os teste unitários
- Dentro da pasta backend: `ginkgo cover cmd domain/model/`
      
### Como executar a interface gráfica do PgAdmin
- Serviço disponível em `localhost:9000'`
- Login e senha: `admin@user.com | 123456`

No dashbord botão direito em **Serves** > **Register** > **Server**

Dados de Conexão aba connection:
- **Host:** `172.17.0.1` 
- **username:** `postgres` 
- **senha:** `root` 
- **database:** authentication
  
### Serviços utilizados ao executar o docker-compose

- Aplicação principal
- Postgres
- PgAdmin
- Go


## **Abaixo os links para a documentação da API**

1. [Registrar novo usuário](./requirements/new-user.md)
2. [Listar todos os registros](./requirements/login.md)