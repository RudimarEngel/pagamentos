insert into Acao (Acao) values('transferencia'), ('erro'), ('rollback');

insert into UsuarioTipo (Tipo) values ('lojista'),('comum');

insert into Usuario (UsuarioTipoId, Nome, CpfCnpj, Email, Senha, SenhaSal)
values (2,'Rudimar Engel', '12345678910', 'rudimar@gmail.com', 'senhateste1', 'salsenhateste1'),
       (1,"Moe's", '01234567891011', 'moe@moesbar.com', 'senhateste2', 'salsenhateste2'),
       (2, 'Homer Simpson', '01987654321', 'homer@moesbar.com', 'senhateste3', 'salsenhateste3' );

insert into Conta (UsuarioId, Saldo) values (1,300.12),(2, 25012.01), (3,256.55);

insert into Maquina(UsuarioId, Marca, Modelo) VALUES (2, 'Q2', 'Queridona Smart');


select * from Usuario; select * from UsuarioTipo;





