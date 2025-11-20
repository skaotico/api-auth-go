# 🔐 api-auth: Servicio de Autenticación Centralizada

![Go Version](https://img.shields.io/badge/Go-1.25.4-blue?style=for-the-badge&logo=go)
![Gin Framework](https://img.shields.io/badge/Gin-v1.11.0-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

## 📋 Descripción

**api-auth** es un servicio de micro-autenticación robusto y eficiente desarrollado en **Go** utilizando el framework **Gin**. Su misión es centralizar la gestión de usuarios y la emisión segura de tokens de acceso, sirviendo como la puerta de entrada confiable para tu ecosistema de microservicios.

Actualmente, el servicio expone un endpoint principal (`/login`) que autentica usuarios vía credenciales (email/password) y genera un **JSON Web Token (JWT)** firmado, habilitando el acceso seguro a otros recursos protegidos.

### ✨ Características Principales

- **🔒 Autenticación Robusta**: Implementación de **bcrypt** para el hashing y salting seguro de contraseñas.
- **🔑 Gestión de Sesiones JWT**: Generación y firma de tokens estándar para autenticación stateless.
- **🏗️ Arquitectura Limpia**: Diseño modular basado en Clean Architecture para máxima mantenibilidad y testabilidad.
- **📄 Documentación Viva**: Integración con Swagger para documentación automática de la API.
- **🐳 Docker Ready**: Contenerización lista para despliegue con Docker.
- **✅ Calidad de Código**: Linter integrado (golangci-lint) y hooks de pre-commit para asegurar estándares.

---

## 🛠️ Tecnologías

| Tecnología | Versión | Descripción |
| :--- | :--- | :--- |
| **Go** | `1.25.4` | Lenguaje principal, concurrente y tipado. |
| **Gin** | `v1.11.0` | Framework HTTP de alto rendimiento. |
| **JWT** | `v5` | Estándar para transmisión segura de información. |
| **PostgreSQL** | `lib/pq` | Motor de base de datos relacional. |
| **Godotenv** | `v1` | Gestión de configuración via `.env`. |

---

## 📂 Estructura del Proyecto

El proyecto sigue una estructura idiomática de Go, separando responsabilidades claramente:

```bash
├── cmd/
│   └── server/          # 🚀 Punto de entrada (main.go)
├── internal/            # 🧠 Lógica de negocio privada
│   ├── domain/          # Modelos y contratos (Interfaces)
│   ├── service/         # Casos de uso y lógica de aplicación
│   ├── repository/      # Acceso a datos (SQL implementation)
│   └── handler/         # Controladores HTTP (Gin handlers)
└── pkg/                 # 📦 Paquetes reutilizables (Logger, DB, JWT)
```

---

## 🚀 Instalación y Ejecución

### Requisitos Previos
- **Go** 1.25.4+
- **Git**
- **PostgreSQL** (Local o Docker)

### 1. Clonar el Repositorio
```bash
git clone https://github.com/skaotico/api-auth-go
cd api-auth
```

### 2. Configuración de Entorno
Crea un archivo `.env` en la raíz basado en el siguiente template:

```env
# --- APP ---
ENV=development
APP_PORT=8080
VERSION=1.0.0

# --- DATABASE ---
DB_HOST=localhost
DB_USER=postgres
DB_PASS=tu_password
DB_NAME=auth_db

# --- SECURITY (JWT) --- 
# ⚠️ Usa una clave secreta fuerte en producción
JWT_SECRET=super_secret_key_change_me
JWT_EXPIRATION=24h
JWT_REFRESH_TTL=72h
```

### 3. Instalar Dependencias
```bash
go mod tidy
```

### 4. Ejecutar el Servidor
```bash
go run ./cmd/server
```
> El servicio estará disponible en: `http://localhost:8080`

---

## 🐳 Dockerización

Para facilitar la ejecución del servicio y su despliegue, `api-auth` puede ejecutarse dentro de un contenedor Docker.

### Dockerfile

El archivo `Dockerfile` ya se encuentra en la raíz del proyecto.  

```bash
# Construye una imagen Docker a partir del Dockerfile en el directorio actual (.)
# -t asigna un nombre y etiqueta a la imagen: "api-auth-go:latest"
docker build -t api-auth-go:latest .

# =========================
# Ejecutar el contenedor Docker
# =========================

# Ejecuta un contenedor a partir de la imagen "api-auth-go:latest"
# -d ejecuta el contenedor en segundo plano (detached)
# -p 8022:8022 mapea el puerto 8022 del contenedor al puerto 8022 de la máquina local
# --name api-auth-go asigna un nombre al contenedor para poder gestionarlo fácilmente
docker run -d -p 8022:8022 --name api-auth-go api-auth-go:latest
```

ya con esto tendrias levantado el proyecto en docker 
 

## 📡 API Endpoints

### `POST /v1/auth/login`

Autentica un usuario y devuelve sus credenciales de acceso.

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "securePassword123"
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "username": "usuario_ejemplo",
    "email": "correo@dominio.com",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  },
  "message": "Operación exitosa",
  "timestamp": "2025-11-20T10:00:00-03:00"
}
```

**Response (Error):**
```json
{
  "success": false,
  "message": "Credenciales inválidas",
  "error_code": "UNAUTHORIZED"
}
```

---

## 🛡️ Calidad de Código

Este proyecto utiliza **golangci-lint** para mantener un código limpio.

### Ejecutar Linter Localmente
```bash
golangci-lint run
```

### Configurar Pre-commit Hook
Evita commits con errores configurando el hook de git:

```bash
# Crear el hook
echo '#!/bin/sh
echo "🔍 Ejecutando linter..."
golangci-lint run
if [ $? -ne 0 ]; then
  echo " Error de Lint! Corrige los errores antes de commitear."
  exit 1
fi
echo " Lint pasado."
exit 0' > .git/hooks/pre-commit

# Dar permisos de ejecución
chmod +x .git/hooks/pre-commit
```

---

## 🤝 Contribución

¡Las contribuciones son bienvenidas!

1.  Haz un **Fork** del proyecto.
2.  Crea tu rama de funcionalidad (`git checkout -b feature/AmazingFeature`).
3.  Haz tus cambios y **Commit** (`git commit -m 'Add some AmazingFeature'`).
4.  **Push** a la rama (`git push origin feature/AmazingFeature`).
5.  Abre un **Pull Request**.

 
