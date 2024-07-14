# Usa una imagen base oficial de Go para construir la aplicación
FROM golang:1.20-alpine AS builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos de tu proyecto al directorio de trabajo en el contenedor
COPY . .

# Descarga las dependencias y construye la aplicación
RUN go mod tidy
RUN go build -o main cmd/main.go

# Usa una imagen base de Ubuntu Focal para ejecutar la aplicación y añadir las dependencias necesarias
FROM ubuntu:focal

# Evita las interacciones con la instalación de paquetes
ENV DEBIAN_FRONTEND=noninteractive

# Actualiza el sistema, instala los certificados CA y las dependencias necesarias
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    wget \
    libxrender1 \
    libfontconfig1 \
    libx11-dev \
    libjpeg62 \
    libxtst6 \
    fontconfig \
    xfonts-75dpi \
    xfonts-base \
    libjpeg-turbo8 \
    && update-ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Descarga e instala wkhtmltopdf
RUN wget https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6-1/wkhtmltox_0.12.6-1.focal_amd64.deb \
    && apt-get update \
    && apt-get install -y --no-install-recommends ./wkhtmltox_0.12.6-1.focal_amd64.deb \
    && rm -rf /var/lib/apt/lists/* \
    && rm wkhtmltox_0.12.6-1.focal_amd64.deb

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el binario de la fase de construcción a la fase de ejecución
COPY --from=builder /app/main .

# Copia el archivo .env
COPY .env .env

# Copia el directorio de plantillas
COPY templates ./templates

# Expone el puerto en el que tu aplicación escucha
EXPOSE 3010

# Define el comando de inicio de tu aplicación
CMD ["./main"]