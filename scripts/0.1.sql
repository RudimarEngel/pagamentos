
-- show databases;
-- SELECT CURRENT_USER();
-- ALTER USER 'root'@'localhost' IDENTIFIED BY '123';

drop database pagamentos;

create database pagamentos;

use pagamentos;

create table UsuarioTipo(
  UsuarioTipoId bigint not null AUTO_INCREMENT,
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
  FOREIGN KEY (UsuarioTipoId) REFERENCES UsuarioTipo(UsuarioTipoId)
)
AUTO_INCREMENT=1;


-- DADOS MOCK
insert into UsuarioTipo (Tipo) values ('lojistas'),('comuns');

insert into Usuario (UsuarioTipoId, Nome, CpfCnpj, Email, Senha, SenhaSal)
values (2,'Rudimar Engel', '12345678910', 'rudimar@gmail.com', 'senhateste1', 'salsenhateste1'),
       (1,"Moe's", '01234567891011', 'moe@moesbar.com', 'senhateste2', 'salsenhateste2');
