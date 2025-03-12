# Documentación del Proyecto: Task Manager

## Introducción
Task Manager es una API RESTful desarrollada en **Go (Golang)** que permite a los usuarios gestionar sus tareas de manera eficiente. Los usuarios pueden registrarse, autenticarse mediante **JWT**, y realizar operaciones CRUD sobre sus tareas. El sistema almacena la información en una base de datos **MySQL**.

## Tecnologías Utilizadas
- **Lenguaje**: Go (Golang)
- **Framework Web**: Gin
- **Base de Datos**: MySQL
- **ORM**: GORM
- **Autenticación**: JSON Web Tokens (JWT)
- **Pruebas**: Testify, sqlmock

## Arquitectura
La API sigue una arquitectura modular basada en controladores para manejar las distintas operaciones del sistema:

1. **Autenticación** (`auth_controller.go`)
2. **Gestión de Tareas** (`task_controller.go`)
3. **Generación de Reportes** (`report_controller.go`)

# Controladores

## 1. Controlador de Autenticación (`auth_controller.go`)
Este controlador gestiona el registro e inicio de sesión de los usuarios.

### **Rutas**
```
POST   /register   Permite registrar un nuevo usuario.
POST   /login      Autentica al usuario y genera un token JWT.
```

### **Detalles de Implementación**
#### **Registro de Usuario** (`Register`)
- Recibe un JSON con `email` y `password`.
- Hashea la contraseña antes de guardarla en la base de datos.
- Responde con un mensaje de éxito o error.

#### **Inicio de Sesión** (`Login`)
- Recibe un JSON con `email` y `password`.
- Verifica la existencia del usuario y compara la contraseña.
- Genera y retorna un **JWT** con la información del usuario.

## 2. Controlador de Tareas (`task_controller.go`)
Este controlador permite la gestión de tareas de los usuarios autenticados.

### **Rutas**
```
POST   /tasks        Crea una nueva tarea.
GET    /tasks        Obtiene todas las tareas del usuario autenticado.
GET    /tasks/:id    Obtiene una tarea específica.
PUT    /tasks/:id    Actualiza una tarea.
DELETE /tasks/:id    Elimina una tarea.
```

### **Detalles de Implementación**
#### **Crear una Tarea** (`CreateTask`)
- Recibe un JSON con `title`, `description` y `dueDate`.
- Obtiene el `user_id` del token JWT y almacena la tarea en la base de datos.

#### **Obtener Tareas** (`GetTasks`, `GetTaskByID`)
- Filtra las tareas por el `user_id` del usuario autenticado.

#### **Actualizar Tarea** (`UpdateTask`)
- Verifica que la tarea pertenezca al usuario antes de actualizarla.

#### **Eliminar Tarea** (`DeleteTask`)
- Elimina una tarea si pertenece al usuario autenticado.

## 3. Controlador de Reportes (`report_controller.go`)
Este controlador permite generar reportes de las tareas en formato **PDF** y **CSV**.

### **Rutas**
```
GET    /reports/pdf   Genera un reporte en formato PDF con todas las tareas del usuario.
GET    /reports/csv   Genera un reporte en formato CSV con todas las tareas del usuario.
```

### **Detalles de Implementación**
#### **Generación de PDF** (`GeneratePDFReport`)
- Obtiene todas las tareas del usuario autenticado.
- Usa la librería `gofpdf` para crear un documento PDF.
- Guarda el archivo y lo devuelve en la respuesta.

#### **Generación de CSV** (`GenerateCSVReport`)
- Obtiene todas las tareas del usuario autenticado.
- Crea un archivo CSV con encabezados y las tareas registradas.
- Guarda el archivo y lo devuelve en la respuesta.

# Pruebas

Se han implementado pruebas unitarias para los controladores de **autenticación** y **reportes**.

- Se usa `sqlmock` para simular la base de datos en los tests.
- Se implementa un middleware falso (`FakeAuthMiddleware`) para evitar la autenticación real.
- Se usan `httptest` y `testify` para validar las respuestas de la API.

## 1. Pruebas de Autenticación (`auth_controller_test.go`)
- **TestRegister**: Verifica que un usuario se pueda registrar exitosamente.
- **TestLogin**: Valida la autenticación de un usuario y la generación del JWT.

## 2. Pruebas de Reportes (`report_controller_test.go`)
- **TestGenerateCSVReport**: Prueba la generación del CSV.
- **TestGeneratePDFReport**: Prueba la generación del PDF.

# Conclusión
Task Manager es una API funcional que permite la gestión de tareas con autenticación segura mediante JWT. Se han implementado endpoints CRUD y la posibilidad de exportar las tareas en diferentes formatos. La arquitectura modular facilita la escalabilidad y mantenimiento del sistema.
