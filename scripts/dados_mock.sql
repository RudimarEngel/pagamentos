truncate table Acao;
insert into Acao (Acao) values('transferencia'), ('erro'), ('rollback');

truncate table UsuarioTipo;
insert into UsuarioTipo (Tipo) values ('lojistas'),('comuns');

truncate table Usuario;
insert into Usuario (UsuarioTipoId, Nome, CpfCnpj, Email, Senha, SenhaSal)
values (2,'Rudimar Engel', '12345678910', 'rudimar@gmail.com', 'senhateste1', 'salsenhateste1'),
       (1,"Moe's", '01234567891011', 'moe@moesbar.com', 'senhateste2', 'salsenhateste2');

truncate table Conta;
insert into Conta (UsuarioId, Saldo) values (1,300.12),(2, 25012.01);

truncate table Maquina;
insert into Maquina(UsuarioId, Marca, Modelo) VALUES (2, 'Q2', 'Queridona Smart');
