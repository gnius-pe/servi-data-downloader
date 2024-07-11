# Usa una imagen base oficial de Go para construir la aplicación
FROM golang:1.20-alpine AS builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos de tu proyecto al directorio de trabajo en el contenedor
COPY . .

# Descarga las dependencias y construye la aplicación
RUN go mod tidy
RUN go build -o main cmd/main.go

# Usa una imagen base de Alpine para ejecutar la aplicación y añadir las dependencias necesarias
FROM alpine:latest

# Instala las dependencias necesarias para wkhtmltopdf
RUN apk add --no-cache \
    libxrender \
    libxext \
    openssl3 \
    ca-certificates \
    fontconfig \
    && apk add --no-cache --virtual .build-deps \
    build-base \
    xz \
    && wget https://github.com/wkhtmltopdf/wkhtmltopdf/releases/download/0.12.4/wkhtmltox-0.12.4_linux-generic-amd64.tar.xz \
    && tar -xJf wkhtmltox-0.12.4_linux-generic-amd64.tar.xz \
    && cp wkhtmltox/bin/wkhtmltopdf /usr/local/bin/ \
    && cp -r wkhtmltox/share /usr/local/ \
    && chmod +x /usr/local/bin/wkhtmltopdf \
    && apk del .build-deps \
    && rm -rf wkhtmltox-0.12.4_linux-generic-amd64.tar.xz wkhtmltox

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
