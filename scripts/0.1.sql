
-- mysql -u root -p pagamentos
  -- password: 
-- show databases;
-- SELECT CURRENT_USER();
-- ALTER USER 'root'@'localhost' IDENTIFIED BY 'senha_fake';

drop database pagamentos;

create database pagamentos;

use pagamentos;

create table UsuarioTipo(
  UsuarioTipoId igint not null AUTO_INCREMENT,
  Tipo varchar(10),

  primary key (UsuarioTipoId)
);

create table Usuario (
  UsuarioId bigint not null AUTO_INCREMENT,
  UsuarioTipoId bigint not null,
  Nome varchar(20),
  CpfCnpj varchar(14) not null unique,
  Email varchar(25) not null unique,
  Senha varchar(50) not null,
  SenhaSal varchar(50) not null,

  primary key (UsuarioId),
  foreign key (UsuarioTipoId) references UsuarioTipo(UsuarioTipoId)
)
AUTO_INCREMENT=1;

create table Conta (
  ContaId bigint not null AUTO_INCREMENT,
  UsuarioId bigint not null,
  Saldo decimal(9,2) not null default '0.00',

  primary key (ContaId),
  foreign key (UsuarioId) references Usuario(UsuarioId)
)
AUTO_INCREMENT=1;

create table Acao (
  AcaoId  bigint not null AUTO_INCREMENT,
)



-- DADOS MOCK
truncate table UsuarioTipo;
insert into UsuarioTipo (Tipo) values ('lojistas'),('comuns');

truncate table Usuario;
insert into Usuario (UsuarioTipoId, Nome, CpfCnpj, Email, Senha, SenhaSal)
values (2,'Rudimar Engel', '12345678910', 'rudimar@gmail.com', 'senhateste1', 'salsenhateste1'),
       (1,"Moe's", '01234567891011', 'moe@moesbar.com', 'senhateste2', 'salsenhateste2');

truncate table Conta;
insert into Conta (UsuarioId, Saldo) values (1,300.12),(2, 25012.01);
