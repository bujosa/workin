USE EnZonaRD

INSERT [dbo].[Categorias] ([Categoria]) VALUES (N'Atraco')
INSERT [dbo].[Categorias] ([Categoria]) VALUES (N'Robo')
INSERT [dbo].[Categorias] ([Categoria]) VALUES (N'Homicidio')

INSERT [dbo].[Delitos] ([IdCategoria], [Id_Locacion], [Fecha], [Hora], [Descripcion], [Modo]) 
VALUES (1, 1, CAST(N'2019-09-27' AS Date), CAST(N'19:13:47.8170000' AS Time), N'Me atracaron con una pistola en la noche y me quitaron Tokyo', N'Pasola')

INSERT [dbo].[Locaciones] ([IdUsuario], [Longitud], [Latitud]) 
VALUES (1, CAST(18.483984000 AS Decimal(9, 6)), CAST(-69.934888000 AS Decimal(9, 6)))

INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Codigo], [Activar]) 
VALUES (N'David', N'Bujosa', N'OpenMind', N'davidbujosa@gmail.com', N'8292666009', N'M', 945123, 1)
