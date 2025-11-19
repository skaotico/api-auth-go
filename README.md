# api-auth: Servicio de Autenticación Centralizada

## Descripción

`api-auth` es un servicio de micro-autenticación desarrollado en **Go** (1.25.4) utilizando el framework **Gin**. Su propósito principal es centralizar la gestión de usuarios y la generación de tokens de acceso seguros.

Actualmente, expone un endpoint (`/login`) para autenticar usuarios mediante correo electrónico y contraseña. Genera un **JSON Web Token (JWT)** válido, permitiendo el acceso a otros microservicios.

### Características

- Autenticación segura mediante **bcrypt** para el almacenamiento de contraseñas.
- Generación de **JWT** para la gestión de sesiones.
- Diseño modular y de capas (Clean Architecture) para facilitar el mantenimiento y la escalabilidad.
- Documentación automática con Swagger.

## Tecnologías Principales

| Tecnología    | Versión              | Descripción                                |
|:------------- |:-------------------- |:------------------------------------------ |
| Go            | 1.25.4               | Lenguaje de programación principal         |
| Gin           | v1.11.0              | Framework web rápido y minimalista para Go |
| JWT           | golang-jwt/jwt/v5    | Implementación de JWT para Go              |
| PostgreSQL    | lib/pq               | Driver para conexión a PostgreSQL          |
| Configuración | Variables de entorno | Carga de variables con `joho/godotenv`     |

## Estructura del Proyecto

El proyecto sigue una estructura modular de Go, alineada con la Arquitectura Limpia/Hexagonal:

```
├── cmd/server           # Punto de entrada para iniciar el servidor
├── internal/            # Lógica de aplicación privada (core del negocio)
│   ├── domain           # Entidades y reglas de negocio
│   ├── service          # Lógica de aplicación
│   ├── repository       # Capa de persistencia (acceso a DB)
│   └── handler          # Capa de presentación (controladores HTTP Gin)
└── pkg/                 # Código reutilizable (configuración, logging, JWT, DB)
```

## Instalación y Ejecución

### Requisitos Previos

1. Go 1.25.4 o superior  
2. Git  
3. Base de Datos PostgreSQL  

### 1. Clonar el Repositorio

```bash
git clone [URL_DE_TU_REPOSITORIO]
cd api-auth
```

### 2. Configuración de Entorno

Cree un archivo `.env` en la raíz del proyecto con la configuración necesaria:

```env
# ===========================
# Ambiente de la Aplicación
# ===========================
ENV=
APP_PORT=
VERSION=

# ===========================
# Configuración de Base de Datos
# ===========================
DB_HOST=
DB_USER=
DB_PASS=
DB_NAME=

# ===========================
# Configuración JWT
# ===========================
# La clave debe ser larga, compleja y única para proteger tus tokens.
JWT_SECRET=
JWT_EXPIRATION=
JWT_REFRESH_TTL=
```

### 3. Instalar Dependencias

```bash
go mod tidy
```

### 4. Ejecutar el Servidor

```bash
go run ./cmd/server
```

El servicio se iniciará y estará disponible en `http://localhost:<env.APP_PORT>`.

## Dockerización

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

## API Endpoints

### Autenticación (Login)

- **Método:** POST  
- **Endpoint:** `/v1/auth/login`  
- **Descripción:** Autentica a un usuario y devuelve un token JWT.

#### Ejemplo de Solicitud (Request Body)

```json
{
  "email": "user@example.com",
  "password": "securePassword123"
}
```

#### Ejemplo de Respuesta Exitosa (200 OK)

```json
{
  "success": true,
  "data": {
    "id": 0,
    "username": "usuario_ejemplo",
    "email": "correo_ejemplo@dominio.com",
    "first_name": "Nombre",
    "last_name": "Apellido",
    "phone": "+56900000000",
    "country_id": 0,
    "address_line": "Dirección de ejemplo 123",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.ejemploTokenDeJWT"
  },
  "message": "Operación exitosa",
  "timestamp": "YYYY-MM-DDTHH:MM:SS-03:00",
  "path": "/v1/auth/login"
}
```

#### Ejemplo de error (500 NOT-OK)

```json
{
  "success": false,
  "message": "usuario no encontrado",
  "timestamp": "2025-11-19T17:23:12-03:00",
  "path": "/v1/auth/login",
  "error_code": "UNAUTHORIZED"
}
```

#### Ejemplo de error (502 NOT-OK)

```json
{
  "success": false,
  "message": "Key: 'LoginRequestDto.Password' Error:Field validation for 'Password' failed on the 'required' tag",
  "timestamp": "2025-11-19T17:32:29-03:00",
  "path": "/v1/auth/login",
  "error_code": "502"
}
```

## Contribución

1. Hacer un fork del repositorio.  
2. Crear una nueva rama: `git checkout -b feature/nueva-funcionalidad`.  
3. Realizar los cambios necesarios.  
4. Asegurarse de que las pruebas pasen (si existen).  
5. Hacer commit de los cambios: `git commit -am 'feat: Añadir nueva funcionalidad X'`.  
6. Subir la rama: `git push origin feature/nueva-funcionalidad`.  
7. Abrir un Pull Request (PR).
