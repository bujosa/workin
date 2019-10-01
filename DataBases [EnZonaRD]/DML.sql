USE EnZonaRD

INSERT [dbo].[Categorias] ([Categoria]) VALUES (N'Atraco')
INSERT [dbo].[Categorias] ([Categoria]) VALUES (N'Robo')
INSERT [dbo].[Categorias] ([Categoria]) VALUES (N'Homicidio')

INSERT [dbo].[Delitos] ([IdCategoria], [Id_Locacion], [Fecha], [Hora], [Descripcion], [Modo]) 
VALUES (1, 1, CAST(N'2019-09-27' AS Date), CAST(N'19:13:47.8170000' AS Time), N'Me atracaron con una pistola en la noche y me quitaron Tokyo', N'Pasola')

INSERT Delitos (IdCategoria, Id_Locacion, Descripcion, Modo)
values (2,1, 'Me robaron todas mis pertenencias mientras iba caminando en la calle' , 'Carro')

INSERT [dbo].[Locaciones] ([IdUsuario], [Longitud], [Latitud]) 
VALUES (1, CAST(18.483984000 AS Decimal(9, 6)), CAST(-69.934888000 AS Decimal(9, 6)))

INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'David', N'Bujosa', N'OpenMind', N'davidbujosa@gmail.com', N'8292666009', N'M', 'Calle 14 E #10 Lucerna')

INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'Jose', N'Bujosa', N'OpenMind', N'josebujosa@gmail.com', N'8494032525', N'M', 'Calle 14 E #10 Vista Hermosa')

INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'Tiron', N'Pelaez', N'OpenMind', N'TironPelaez@gmail.com', N'8290001111', N'M', 'Calle 14 E #10 Cancino Viejo')

INSERT Usuarios (Nombre,Apellido,Contrasena, Email,Telefono, Genero,Direccion)

select Top(1) IdLocacion from Locaciones order by IdLocacion desc 
select * from Delitos
select * from Categorias
select * from Usuarios

USE EnZonaRD
select D.Fecha, D.Hora, D.Descripcion, D.Modo, L.Longitud, L.Latitud, U.IdUsuario, C.Categoria from Delitos as D
INNER JOIN Locaciones as L  on D.Id_Locacion = L.IdLocacion
INNER JOIN Usuarios as U on L.IdUsuario = U.IdUsuario
INNER JOIN Categorias as C on D.IdCategoria = C.IdCategoria