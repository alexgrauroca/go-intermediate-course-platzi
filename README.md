# Nemon Power Optimizer
## Primeros pasos
En primer lugar hay que instalar los paquetes del servicio con el comando:

`go mod vendor`

Lógicamente, es imprescindible tener Go instalado en el SO, ya sea mediante un contenedor o directamente en el propio host.

## Iniciar el servicio
Para levantar el servicio se puede hacer compilando el código y ejecutando el archivo compilado, o bien utilizando el comando de go que genera y ejecuta un archivo compilado temporal.

Para compilar:

`go build src/main.go`

Generará un archivo compilado en la raíz, llamado main. Para ejecutarlo:

`./main`

Para utilizar el compilador temporal:

`go run src/main.go`

El puerto sobre el cual corre el servicio actualmente es el 1323.

