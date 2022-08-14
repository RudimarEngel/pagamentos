
-- show databases;

drop database pagamentos;

create database pagamentos;

use pagamentos;

create table UsuarioTipo(
  UsuarioTipoId bigint not null AUTO_INCREMENT,
  Tipo varchar(10)

  primary key (UsuarioTipoId)
)

create table Usuario (
  UsuarioId bigint not null AUTO_INCREMENT,
  UsuarioTipoId bigint not null,
  Nome varchar(20),
  CpfCnpj varchar(14),
  Email varchar(25),
  Senha varchar(50),
  SenhaSal varchar(50),

  primary key (UsuarioId)
)
AUTO_INCREMENT=1;


-- DADOS MOCK
UsuarioTipo
insert into Usuario (UsuarioTipoId, Nome, CpfCnpj, Email, SenhaSal)
values (1,'Rudimar Engel', '')