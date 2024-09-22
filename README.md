# Proyecto de Servidor de Datos

## Descripción

Servicio que da recursos de pdf y exel

## Requisitos

- Go 1.16 o superior

## Instalación

### Paso 1: Clonar el repositorio

```sh
    git clone https://github.com/tu-usuario/tu-repositorio.git
```

### Paso 2: Instalar dependencias

```sh
    go mod tidy
```

### Paso 2: Ejecutar la aplicación

```sh
    go run cmd/main.go
```

Este proyecto está licenciado bajo la Licencia MIT.

## Ejeccion en Docker

### Paso 1: Crear la imagen
```sh
docker build -t dar-service-download .
```

### Paso 1: Se crea el container y correr:

```sh
docker run -d -p 3010:3010 --name dar-service-donwload-container dar-service-download
```


### Local Host
http://localhost:3010/