# Auditoría de Código y Mejoras Propuestas

Este documento detalla los hallazgos de la revisión de código, enfocándose en seguridad, deuda técnica y oportunidades de mejora para el proyecto `api-auth-go`.

## 1. Seguridad 🛡️

### 1.1. Invalidación Estricta por Detección de Robo (Token Reuse)
- **Ubicación:** `internal/service/auth/impl/authService.go` (Método `RefreshToken`)
- **Problema:** Actualmente, cuando se detecta que un Refresh Token está siendo reutilizado (señal clara de robo o condición de carrera), el sistema solo loguea una advertencia y rechaza la petición.
- **Riesgo:** Si un atacante tiene el token, puede intentar usarlo en otro momento o ya haber generado un token válido antes.
- **Acción Recomendada:** **Descomentar y activar** la línea `s.cacheService.DeleteAll(...)`. Es preferible cerrar la sesión del usuario legítimo (forzándolo a loguearse de nuevo) para expulsar inmediatamente al atacante.

### 1.2. Rate Limiting en Refresh Token
- **Ubicación:** `internal/app/app.go`
- **Problema:** El endpoint `/auth/refresh` no parece tener un Rate Limit específico visible en las rutas (aunque Login sí lo tiene).
- **Riesgo:** Un atacante podría intentar saturar el servicio de validación de tokens o intentar fuerza bruta sobre tokens.
- **Acción Recomendada:** Aplicar el middleware de Rate Limit también a la ruta `/auth/refresh`.

### 1.3. Sanitización y Validación de Inputs
- **Ubicación:** DTOs (`internal/handler/auth/dto/request`)
- **Problema:** Se confía en `ShouldBindJSON`.
- **Acción Recomendada:** Asegurar el uso extensivo de tags de validación (`binding:"required,email,min=8,alphanum"`) para rechazar datos mal formados antes de que toquen la capa de servicio.

## 2. Deuda Técnica 🛠️

### 2.1. Código Muerto / Comentarios de Debug
- **Ubicación:** `internal/service/auth/impl/authService.go` (Método `Login`)
- **Problema:** Existe un bloque grande de código comentado (`// Verificar datos guardados en Redis...`).
- **Impacto:** Ensucia el código y dificulta la lectura.
- **Acción Recomendada:** Eliminar este código. La verificación debe hacerse mediante **Tests de Integración** automatizados, no mediante trazas manuales en el código productivo.

### 2.2. "Magic Strings" en Redis Keys
- **Ubicación:** `internal/service/cache/impl/cacheServiceImpl.go`
- **Problema:** Las claves se construyen manualmente: `"auth:jwt:" + jwt`.
- **Impacto:** Si se decide cambiar el prefijo o la estructura de las claves, hay que buscar y reemplazar en múltiples lugares, aumentando el riesgo de errores.
- **Acción Recomendada:** Crear constantes o métodos privados para generar las claves (ej: `func getJwtKey(token string) string`).

### 2.3. Cobertura de Tests (Testing)
- **Estado General:** Faltan tests unitarios y de integración visibles.
- **Acción Recomendada:**
    - Implementar tests unitarios para `AuthService` (mockeando Redis y UserRepo).
    - Implementar tests para `UserService`.
    - Usar librerías como `testify` para aserciones y mocks.

## 3. Arquitectura y Mantenibilidad 🚀

### 3.1. Gestión de Errores de Dominio
- **Problema:** El manejo de errores a veces retorna errores directos de librerías o strings simples.
- **Acción Recomendada:** Definir un paquete `internal/domain/errors` con errores tipados (ej: `ErrUserNotFound`, `ErrInvalidToken`, `ErrDatabaseDown`). Esto permite que el Handler decida el código HTTP (404, 401, 500) de forma determinista basándose en el tipo de error.

### 3.2. SQL Hardcodeado
- **Ubicación:** `internal/repository/user/postgres_repository.go`
- **Observación:** Las queries SQL están escritas como strings dentro de los métodos.
- **Acción Recomendada:** Para un proyecto pequeño está bien. Si crece, considerar mover las queries a constantes o usar un Query Builder / ORM ligero para facilitar el mantenimiento y evitar errores de sintaxis SQL.

## 4. Resumen de Prioridades

1.  🔴 **Alta:** Activar invalidación por Token Reuse (Seguridad).
2.  🔴 **Alta:** Eliminar código muerto en `AuthService` (Limpieza).
3.  🟡 **Media:** Implementar Tests Unitarios básicos (Calidad).
4.  🟢 **Baja:** Refactorizar Magic Strings de Redis (Mantenibilidad).
