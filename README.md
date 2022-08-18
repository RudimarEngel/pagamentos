# pagamentos

# REFRENCIAS:
- https://go.dev/doc/tutorial/web-service-gin
- https://go.dev/src/net/http/status.go
- https://medium.com/wesionary-team/a-clean-architecture-for-web-application-in-go-lang-4b802dd130bb
- https://go.dev/doc/code
- https://setapp.com/how-to/use-go-with-mysql#:~:text=Now%2C%20let's%20connect%20to%20MySQL,you've%20installed%20Go%20in.


# LEITURAS:
- https://www.vivaolinux.com.br/dica/Access-denied-for-user-rootlocalhost-no-MySQL-Server-Community-57-Resolvido-CentOS7-x86-64


# BANCO:
-- mysql -u root -p pagamentos
  -- password: 
-- show databases;
-- SELECT CURRENT_USER();
-- ALTER USER 'root'@'localhost' IDENTIFIED BY 'senha_fake';

# REQUISIÇÃO:
POST, http://localhost:8080/transferencia
Body:
{
    "valor": 0.02,
    "IdPagante": 1,
    "IdRecebedor": 2,
    "MaquinaId": 1
}
