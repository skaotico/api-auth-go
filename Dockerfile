# Etapa 1: Build
FROM golang:1.25.4-alpine AS builder

# Instalar dependencias necesarias
RUN apk add --no-cache git ca-certificates tzdata

# Establecer directorio de trabajo
WORKDIR /app

# Copiar archivos de módulos y descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código fuente
COPY . .

# Construir la aplicación en modo release (binario)
RUN go build -o server ./cmd/server

# Etapa 2: Imagen final ligera
FROM alpine:latest

# Instalar certificados y zona horaria
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /root/

# Copiar el binario compilado desde la etapa build
COPY --from=builder /app/server .

# Copiar archivos de configuración si existen
# COPY .env .

# Exponer puerto de la aplicación
EXPOSE 8022

# Comando por defecto para ejecutar la aplicación
CMD ["./server"]
