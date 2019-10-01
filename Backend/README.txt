los paquetes que hacen faltan 

tienen que instarlarlo en la ruta de src/github.com

usar los siguientes comandos:


go get -u github.com/gorilla/mux

go get -u github.com/denisenkom/go-mssqldb

go get -u github.com/rs/cors
_______________________

el archivo config es la conexion a la base de datos

luego de que tenga todos los paquetes instalado

OJO : las carpetas que estan en el bakend tienen que estar en el SRC lo recomendable es crear una carpeta llamada 'Backend' y 
de ahi ponerla en la ruta SRC donde tienes instalado go adentro de la carpeta SRC

para correr usar comando : go run main.go