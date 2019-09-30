use EnZonaRD
INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'David', N'Mendez', N'1234', N'David Bujosa', N'8296441368', N'F', 'Calle 10 Alma Rosa')

INSERT [dbo].[Locaciones] ([IdUsuario], [Longitud], [Latitud]) 
VALUES (2, CAST(18.474541000 AS Decimal(9, 6)), CAST(- -69.888562000 AS Decimal(9, 6)))

INSERT [dbo].[Delitos] ([IdCategoria], [Id_Locacion], [Fecha], [Hora], [Descripcion], [Modo]) 
VALUES (2,2, CAST(N'2019-08-13' AS Date), CAST(N'16:13:47.8170000' AS Time), N'Me robaron mi billetera y el celcular', N'Motor')



INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'Maria', N'Rodriguez', N'mimi22', N'MariRo@gmail.com', N'8096451357', N'F', 'Calle 36 San Isidro')

INSERT [dbo].[Locaciones] ([IdUsuario], [Longitud], [Latitud]) 
VALUES (3, CAST(18.487341000 AS Decimal(9, 6)), CAST(-69.878562000 AS Decimal(9, 6)))

INSERT [dbo].[Delitos] ([IdCategoria], [Id_Locacion], [Fecha], [Hora], [Descripcion], [Modo]) 
VALUES (3,3, CAST(N'2019-09-18' AS Date), CAST(N'22:13:52.8170000' AS Time), N'Me llevaron el carro mientras conducia', N'Carro')




INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'Pedro', N'Montez', N'pipo234', N'Eldestructor24@gmail.com', N'8295816742', N'M', 'Calle A 27 Gualey')

INSERT [dbo].[Locaciones] ([IdUsuario], [Longitud], [Latitud]) 
VALUES (4, CAST(18.491703000 AS Decimal(9, 6)), CAST(-69.921513000 AS Decimal(9, 6)))

INSERT [dbo].[Delitos] ([IdCategoria], [Id_Locacion], [Fecha], [Hora], [Descripcion], [Modo]) 
VALUES (1,4, CAST(N'2019-07-19' AS Date), CAST(N'23:22:52.8170000' AS Time), N'Robaron mi celular mientras salia del club', N'Motor')


INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'Ocasio', N'Pichardo', N'pami264', N'Oca99tru@gmail.com', N'8294735871', N'M', 'Calle Na 27 El millon')

INSERT [dbo].[Locaciones] ([IdUsuario], [Longitud], [Latitud]) 
VALUES (5, CAST(18.497525000 AS Decimal(9, 6)), CAST(-69.936806000 AS Decimal(9, 6)))

INSERT [dbo].[Delitos] ([IdCategoria], [Id_Locacion], [Fecha], [Hora], [Descripcion], [Modo]) 
VALUES (1,5, CAST(N'2019-07-30' AS Date), CAST(N'20:31:48.8170000' AS Time), N'Una mujer se llevo mi billetera', N'Persona')

INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'Irina', N'Rojas', N'r4uid', N'RoIri89@gmail.com', N'8097864895', N'F', 'Calle Mencion Piantini')

INSERT [dbo].[Locaciones] ([IdUsuario], [Longitud], [Latitud]) 
VALUES (6, CAST(18.497971000 AS Decimal(9, 6)), CAST( -69.946883000 AS Decimal(9, 6)))

INSERT [dbo].[Delitos] ([IdCategoria], [Id_Locacion], [Fecha], [Hora], [Descripcion], [Modo]) 
VALUES (1,6, CAST(N'2019-08-26' AS Date), CAST(N'18:47:35.8170000' AS Time), N'Se llevaron mi cartera mientras caminaba', N'Motor')


INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'Joshua', N'Berrocal', N'uctUsCUe4', N'jhu24@gmail.com', N'8297668903', N'M', 'Calle Manzana 34 los mina')

INSERT [dbo].[Locaciones] ([IdUsuario], [Longitud], [Latitud]) 
VALUES (7, CAST(18.491743000 AS Decimal(9, 6)), CAST(-69.921513000 AS Decimal(9, 6)))

INSERT [dbo].[Delitos] ([IdCategoria], [Id_Locacion], [Fecha], [Hora], [Descripcion], [Modo]) 
VALUES (2,7, CAST(N'2019-07-06' AS Date), CAST(N'15:56:52.8170000' AS Time), N'Me quitaron un collar que tenia', N'Carro Publico')



INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'Alan', N'Elvira', N'8zXwZZqj', N'Ou8h99tru@gmail.com', N'8295667890', N'M', 'Calle 23 z')

INSERT [dbo].[Locaciones] ([IdUsuario], [Longitud], [Latitud]) 
VALUES (8, CAST(18.497625000 AS Decimal(9, 6)), CAST(-69.946806000 AS Decimal(9, 6)))

INSERT [dbo].[Delitos] ([IdCategoria], [Id_Locacion], [Fecha], [Hora], [Descripcion], [Modo]) 
VALUES (2,8, CAST(N'2019-07-13' AS Date), CAST(N'21:34:48.8170000' AS Time), N'Me robaron el celular', N'Motor')


INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'Nicole', N'Espinar', N'gSzax6GG', N'niIri89@gmail.com', N'8091490344', N'F', 'Calle Cruzada La puya')

INSERT [dbo].[Locaciones] ([IdUsuario], [Longitud], [Latitud]) 
VALUES (9, CAST(18.486336000 AS Decimal(9, 6)), CAST( -69.958527000 AS Decimal(9, 6)))

INSERT [dbo].[Delitos] ([IdCategoria], [Id_Locacion], [Fecha], [Hora], [Descripcion], [Modo]) 
VALUES (1,9, CAST(N'2019-09-11' AS Date), CAST(N'19:35:15.8170000' AS Time), N'Se llevaron mi computadora de la mesa', N'Persona')


INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'Octavio', N'Arnal', N'zVHDXZjE', N'fsdf45@gmail.com', N'8295889032', N'F', 'Calle churchill 23 ')

INSERT [dbo].[Locaciones] ([IdUsuario], [Longitud], [Latitud]) 
VALUES (10, CAST(18.483541000 AS Decimal(9, 6)), CAST(- -69.878562000 AS Decimal(9, 6)))

INSERT [dbo].[Delitos] ([IdCategoria], [Id_Locacion], [Fecha], [Hora], [Descripcion], [Modo]) 
VALUES (1,10, CAST(N'2019-08-30' AS Date), CAST(N'08:13:47.8170000' AS Time), N'Me robaron mi cartera con todo adentro', N'Motor')



INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'Lilian', N'Varela', N'mAXzb98U', N'rfdsfdsk@gmail.com', N'8096778431', N'F', 'Calle Lope de vega 45')

INSERT [dbo].[Locaciones] ([IdUsuario], [Longitud], [Latitud]) 
VALUES (11, CAST(18.4885341000 AS Decimal(9, 6)), CAST(-69.884562000 AS Decimal(9, 6)))

INSERT [dbo].[Delitos] ([IdCategoria], [Id_Locacion], [Fecha], [Hora], [Descripcion], [Modo]) 
VALUES (2,11, CAST(N'2019-09-25' AS Date), CAST(N'20:34:52.8170000' AS Time), N'El hombre se llevo mi celular', N'Pasola')




INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'Francisco', N'Alguacil', N'uctUsCUe4', N'cdskk44@gmail.com', N'8292559038', N'M', 'Calle 45 los mina')

INSERT [dbo].[Locaciones] ([IdUsuario], [Longitud], [Latitud]) 
VALUES (12, CAST(18.492743000 AS Decimal(9, 6)), CAST(-69.941513000 AS Decimal(9, 6)))

INSERT [dbo].[Delitos] ([IdCategoria], [Id_Locacion], [Fecha], [Hora], [Descripcion], [Modo]) 
VALUES (2,12, CAST(N'2019-07-23' AS Date), CAST(N'15:56:52.8170000' AS Time), N'Me quitaron un collar que tenia', N'Carro Publico')




INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'Israel', N'Escribano', N'8zXwZZqj', N'Ofddsf4u8h99tru@gmail.com', N'8294778294', N'M', 'Calle 23 girasoles')

INSERT [dbo].[Locaciones] ([IdUsuario], [Longitud], [Latitud]) 
VALUES (13, CAST(18.496625000 AS Decimal(9, 6)), CAST(-69.976806000 AS Decimal(9, 6)))

INSERT [dbo].[Delitos] ([IdCategoria], [Id_Locacion], [Fecha], [Hora], [Descripcion], [Modo]) 
VALUES (1,13, CAST(N'2019-07-15' AS Date), CAST(N'21:34:48.8170000' AS Time), N'Me robaron el celular', N'Motor')



INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'Elena', N'Laguna', N'gSzax6GG', N'sdf4fnen@gmail.com', N'8093486731', N'F', 'Calle Martin lutero')

INSERT [dbo].[Locaciones] ([IdUsuario], [Longitud], [Latitud]) 
VALUES (14, CAST(18.496336000 AS Decimal(9, 6)), CAST( -69.968527000 AS Decimal(9, 6)))

INSERT [dbo].[Delitos] ([IdCategoria], [Id_Locacion], [Fecha], [Hora], [Descripcion], [Modo]) 
VALUES (1,14, CAST(N'2019-08-13' AS Date), CAST(N'19:35:15.8170000' AS Time), N'Un hombre se llevo mi cartera', N'Persona')


INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'Manuel', N'Arnaud', N'zVHDfXZjE', N'fdfds@gmail.com', N'8295889032', N'F', 'Calle luperon 45 ')

INSERT [dbo].[Locaciones] ([IdUsuario], [Longitud], [Latitud]) 
VALUES (15, CAST(18.493541000 AS Decimal(9, 6)), CAST(- -69.879562000 AS Decimal(9, 6)))

INSERT [dbo].[Delitos] ([IdCategoria], [Id_Locacion], [Fecha], [Hora], [Descripcion], [Modo]) 
VALUES (1,15, CAST(N'2019-08-25' AS Date), CAST(N'13:26:47.8170000' AS Time), N'Me robaron mi cartera con todo adentro', N'Motor')




INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'Manuela', N'Nomare', N'mAXzb98U', N'afjkdjfe4@gmail.com', N'8096778431', N'F', 'Calle Murlon 54')

INSERT [dbo].[Locaciones] ([IdUsuario], [Longitud], [Latitud]) 
VALUES (16, CAST(18.4895241000 AS Decimal(9, 6)), CAST(-69.884562000 AS Decimal(9, 6)))

INSERT [dbo].[Delitos] ([IdCategoria], [Id_Locacion], [Fecha], [Hora], [Descripcion], [Modo]) 
VALUES (2,16, CAST(N'2019-07-25' AS Date), CAST(N'22:45:52.8170000' AS Time), N'El hombre se llevo mi celular', N'Pasola')




INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'Francis', N'Algustin', N'uctUsCUe4', N'ek;awrk65@gmail.com', N'8292559038', N'M', 'Calle rosal 45')

INSERT [dbo].[Locaciones] ([IdUsuario], [Longitud], [Latitud]) 
VALUES (17, CAST(18.493743000 AS Decimal(9, 6)), CAST(-69.941513000 AS Decimal(9, 6)))

INSERT [dbo].[Delitos] ([IdCategoria], [Id_Locacion], [Fecha], [Hora], [Descripcion], [Modo]) 
VALUES (2,17, CAST(N'2019-07-12' AS Date), CAST(N'18:56:52.8170000' AS Time), N'Me quitaron un collar que tenia', N'Carro Publico')



INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'Emmanuel', N'Moreno', N'8zXwZZqj', N'fksdflj54@gmail.com', N'8294778294', N'M', 'Calle Mejia 22')

INSERT [dbo].[Locaciones] ([IdUsuario], [Longitud], [Latitud]) 
VALUES (18, CAST(18.496625000 AS Decimal(9, 6)), CAST(-69.977806000 AS Decimal(9, 6)))

INSERT [dbo].[Delitos] ([IdCategoria], [Id_Locacion], [Fecha], [Hora], [Descripcion], [Modo]) 
VALUES (1,18, CAST(N'2019-07-15' AS Date), CAST(N'23:54:48.8170000' AS Time), N'Me robaron el celular', N'Motor')


INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'Francia', N'Ricart', N'gSzax6GG', N'frkwekr43@gmail.com', N'809345876', N'F', 'Calle Luo ricart')

INSERT [dbo].[Locaciones] ([IdUsuario], [Longitud], [Latitud]) 
VALUES (19, CAST(18.496536000 AS Decimal(9, 6)), CAST( -69.967527000 AS Decimal(9, 6)))

INSERT [dbo].[Delitos] ([IdCategoria], [Id_Locacion], [Fecha], [Hora], [Descripcion], [Modo]) 
VALUES (2,19, CAST(N'2019-08-13' AS Date), CAST(N'19:56:15.8170000' AS Time), N'Un hombre se llevo mi cartera', N'Persona')

INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'Paola', N'Quesada', N'zVHDXZjE', N'lilovvcf@gmail.com', N'8296847934', N'F', 'Calle 12 Respaldo ')

INSERT [dbo].[Locaciones] ([IdUsuario], [Longitud], [Latitud]) 
VALUES (20, CAST(18.473541000 AS Decimal(9, 6)), CAST(- -69.878562000 AS Decimal(9, 6)))

INSERT [dbo].[Delitos] ([IdCategoria], [Id_Locacion], [Fecha], [Hora], [Descripcion], [Modo]) 
VALUES (2,20, CAST(N'2019-08-28' AS Date), CAST(N'08:13:47.8170000' AS Time), N'Me robaron mi cartera con todo adentro', N'Motor')


INSERT [dbo].[Usuarios] ([Nombre], [Apellido], [Contrasena], [Email], [Telefono], [Genero], [Direccion]) 
VALUES (N'Clara', N'Cortes', N'mAXzb98U', N'rtg65vxg@gmail.com', N'8091865932', N'F', 'Calle 39 Capotillo')

INSERT [dbo].[Locaciones] ([IdUsuario], [Longitud], [Latitud]) 
VALUES (21, CAST(18.4865341000 AS Decimal(9, 6)), CAST(-69.874562000 AS Decimal(9, 6)))

INSERT [dbo].[Delitos] ([IdCategoria], [Id_Locacion], [Fecha], [Hora], [Descripcion], [Modo]) 
VALUES (2,21, CAST(N'2019-09-23' AS Date), CAST(N'20:34:52.8170000' AS Time), N'El hombre se llevo mi celular', N'Pasola')







