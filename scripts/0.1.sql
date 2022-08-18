

drop database pagamentos;

create database pagamentos;

use pagamentos;

create table UsuarioTipo(
  UsuarioTipoId bigint not null AUTO_INCREMENT,
  Tipo varchar(10),
  Ativo boolean default true,
  CreatedAt timestamp not null default current_timestamp(),
  DeletedAt  timestamp not null default '0000-00-00 00:00:00',

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
  Ativo boolean default true,
  CreatedAt timestamp not null default current_timestamp(),
  DeletedAt  timestamp not null default '0000-00-00 00:00:00',

  primary key (UsuarioId),
  foreign key (UsuarioTipoId) references UsuarioTipo(UsuarioTipoId)
)
AUTO_INCREMENT=1;

create table Conta (
  ContaId bigint not null AUTO_INCREMENT,
  UsuarioId bigint not null,
  Saldo decimal(9,2) not null default '0.00',
  Ativo boolean default true,
  CreatedAt timestamp not null default current_timestamp(),
  DeletedAt  timestamp not null default '0000-00-00 00:00:00',

  primary key (ContaId),
  foreign key (UsuarioId) references Usuario(UsuarioId)
)
AUTO_INCREMENT=1;

create table Acao (
  AcaoId  bigint not null AUTO_INCREMENT,
  Acao varchar(14) not null unique,
  Ativo boolean default true,
  CreatedAt timestamp not null default current_timestamp(),
  DeletedAt  timestamp not null default '0000-00-00 00:00:00',

  primary key (AcaoId)
)
AUTO_INCREMENT=1;

create table Maquina (
  MaquinaId bigint not null AUTO_INCREMENT,
  UsuarioId bigint not null,
  Marca varchar(15),
  Modelo varchar(15),
  Ativo boolean default true,
  CreatedAt timestamp not null default current_timestamp(),
  DeletedAt  timestamp not null default '0000-00-00 00:00:00',

  primary key (MaquinaId)
) AUTO_INCREMENT=1;


/*
-- Tabela criada automaticamente
CREATE TABLE IF NOT EXISTS Bilhetes_20220817  ( 
	Bilhetes_20220817Id bigint not null AUTO_INCREMENT,
	IdUsuarioPag bigint,
	IdUsuarioRec bigint,
	AcaoId bigint not null,
	Valor decimal(9,2),
	MaquinaId bigint,
	CreatedAt timestamp default current_timestamp(),
	PRIMARY KEY (Bilhetes_20220817Id)
) AUTO_INCREMENT=1;
*/


show tables;