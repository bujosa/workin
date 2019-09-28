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
	Direccion varchar(100) NOT NULL,
);

CREATE TABLE Categorias
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
    IdCategoria INT FOREIGN KEY REFERENCES Categorias(IdCategoria),
    Id_Locacion INT FOREIGN KEY REFERENCES Locaciones(IdLocacion),
    Fecha DATE Default(CONVERT (date,GETDATE())), 
	Hora TIME  Default(CONVERT (time, GETDATE())),  
    Descripcion VARCHAR(300),
	Modo VARCHAR(50),
);



