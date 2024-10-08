# Dar servicio de descarga de archivos

### Descripción

Servicio que da recursos de pdf y exel

### Requisitos

- Go 1.16 o superior

## Instalación

#### Paso 1: Clonar el repositorio

```sh
git clone https://github.com/tu-usuario/tu-repositorio.git
```

#### Paso 2: Instalar dependencias

```sh
go mod tidy
```

#### Paso 2: Ejecutar la aplicación

```sh
go run cmd/main.go
```

### Ejeccion en Docker

#### Paso 1: Crear la imagen
```sh
docker build -t dar-service-download .
```

#### Paso 1: Se crea el container y correr:

```sh
docker run -d -p 3010:3010 --name dar-service-download-container dar-service-download
```

## Local Host
http://localhost:3010/

### API Endpoints

| Método | Endpoint                         | Descripción                                | File        |
|--------|----------------------------------|--------------------------------------------|--------------|
| GET    | `http://localhost:3010/api/patient/downloader/:id`    | Obtiene los detalles de un paciente por ID. | ![PDF](https://img.icons8.com/material-outlined/24/000000/pdf.png) |
| GET    | `http://localhost:3010/api/patient/download/csv`      | Exporta la lista de pacientes en formato CSV. | ![CSV](https://img.icons8.com/material-outlined/24/000000/csv.png) |


## Derechos de Autor

© 2024 gnius-pe. Todos los derechos reservados.

Este proyecto está alojado en el repositorio [dar-servicio-download](https://github.com/gnius-pe/servi-data-downloader). El uso del código y los recursos de este proyecto están sujetos a los términos de la licencia [Licencia MIT](./LICENSE).
