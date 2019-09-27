CREATE DATABASE EnZonaRD

USE EnZonaRD

CREATE TABLE Usuarios
(
    IdUsuario INT PRIMARY KEY IDENTITY,
   	Nombre VARCHAR(50) NOT NULL,
    Apellido VARCHAR(50) NOT NULL,
	Contrasena VARCHAR(50) NOT NULL,
	Email VARCHAR(50) NOT NULL UNIQUE,
	Telefono CHAR(10) NOT NULL,
    Genero char(1) NOT NULL,
	Codigo int NOT NULL,
	Activar bit default 0
);

CREATE TABLE Categoria
(
    IdCategoria INT PRIMARY KEY IDENTITY,
   	Categoria VARCHAR(50) NOT NULL,
);

CREATE TABLE Locaciones
(
    IdLocacion INT PRIMARY KEY IDENTITY,
    IdUsuario INT FOREIGN KEY REFERENCES Usuarios(IdUsuario),
    Longitud DECIMAL(9,6) NOT NULL,
    Latitud DECIMAL(9,6) NOT NULL,
);

CREATE TABLE Delitos
(
    IdDelitos INT PRIMARY KEY IDENTITY,
    IdCategoria INT FOREIGN KEY REFERENCES Categoria(IdCategoria),
    Id_Locacion INT FOREIGN KEY REFERENCES Locaciones(IdLocacion),
    FECHA DATE Default(CONVERT (date,GETDATE())), 
	HORA TIME  Default(CONVERT (time, GETDATE())),  
    Descripcion VARCHAR(300),
	ModeDeTransporte VARCHAR(50),
);



