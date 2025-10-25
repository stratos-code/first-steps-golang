¡Perfecto 🙌!
Aquí tienes un **plan de estudios de Go de 6 semanas**, organizado **por días (1 h diaria / 5 días a la semana)**.
➡️ Ideal para alguien con experiencia en TypeScript que ya entiende conceptos de programación moderna.
📅 **Duración total:** 30 días de estudio efectivo.

---

## 🧭 Semana 1: Fundamentos y sintaxis básica 🧠

**Objetivo:** Entender la base del lenguaje y su filosofía minimalista.

| Día | Tema                | Tareas prácticas                                                                               |
| --- | ------------------- | ---------------------------------------------------------------------------------------------- |
| 1   | Filosofía de Go     | Leer sobre la filosofía de Go, instalar Go y configurar entorno (`go version`, `go mod init`). |
| 2   | Estructura básica   | Crear tu primer `hello.go`. Probar `go run`, `go build`.                                       |
| 3   | Tipos y variables   | Declarar variables, constantes y usar `:=`. Practicar con slices, maps y arrays.               |
| 4   | Control de flujo    | Ejercicios con `if`, `for`, `switch`.                                                          |
| 5   | Funciones y errores | Crear funciones con múltiples retornos. Manejar errores sin excepciones.                       |

📚 Recursos: [Tour oficial de Go](https://go.dev/tour) · [Go by Example](https://gobyexample.com)

---

## 🧱 Semana 2: Tipos compuestos, punteros y organización 🧭

**Objetivo:** Entender cómo Go estructura programas más grandes.

| Día | Tema                  | Tareas prácticas                                              |
| --- | --------------------- | ------------------------------------------------------------- |
| 6   | Structs y métodos     | Definir structs y métodos asociados.                          |
| 7   | Interfaces            | Crear interfaces y structs que las implementen (duck typing). |
| 8   | Punteros              | Crear funciones que modifiquen structs por referencia.        |
| 9   | Paquetes y módulos    | Dividir un proyecto en múltiples paquetes.                    |
| 10  | Estilo y convenciones | Usar `go fmt`, `go vet`, y aprender convenciones idiomáticas. |

📚 Recursos: [Effective Go](https://golang.org/doc/effective_go)

---

## ⚡ Semana 3: Concurrencia (goroutines & channels) 🕹️

**Objetivo:** Aprender una de las características más potentes de Go.

| Día | Tema          | Tareas prácticas                                                  |
| --- | ------------- | ----------------------------------------------------------------- |
| 11  | Goroutines    | Crear funciones concurrentes y comparar rendimiento.              |
| 12  | Channels      | Comunicar goroutines con channels simples.                        |
| 13  | Select        | Implementar un worker pool.                                       |
| 14  | Context       | Usar `context.WithTimeout` para cancelar tareas.                  |
| 15  | Mini proyecto | Construir un pequeño scraper concurrente (p. ej. URLs de prueba). |

📚 Recurso recomendado: charla “Go Concurrency Patterns” de Rob Pike

---

## 🌐 Semana 4: HTTP, APIs y JSON 🚀

**Objetivo:** Crear un backend RESTful idiomático.

| Día | Tema          | Tareas prácticas                                |
| --- | ------------- | ----------------------------------------------- |
| 16  | net/http      | Crear un servidor básico con `http.HandleFunc`. |
| 17  | Ruteo         | Implementar un router con chi o gin.            |
| 18  | JSON          | Parsear body de requests y responder en JSON.   |
| 19  | Middlewares   | Autenticación sencilla (header token fake).     |
| 20  | Mini proyecto | API CRUD básica (users o tasks).                |

📚 Recurso: [https://go.dev/doc](https://go.dev/doc)

---

## 🧪 Semana 5: Bases de datos y testing 🧰

**Objetivo:** Persistencia de datos y asegurar calidad de código.

| Día | Tema                 | Tareas prácticas                        |
| --- | -------------------- | --------------------------------------- |
| 21  | database/sql         | Conectar a PostgreSQL y hacer queries.  |
| 22  | ORM o query builder  | Probar GORM o sqlc.                     |
| 23  | Migraciones          | Crear tablas y seeders simples.         |
| 24  | Testing              | Test unitarios con `testing` y testify. |
| 25  | Tests de integración | Probar endpoints reales con DB.         |

📚 Recurso: [https://pkg.go.dev/testing](https://pkg.go.dev/testing)

---

## 🚢 Semana 6: Buenas prácticas, despliegue y arquitectura 🏗️

**Objetivo:** Preparar proyectos reales para producción.

| Día | Tema               | Tareas prácticas                                                    |
| --- | ------------------ | ------------------------------------------------------------------- |
| 26  | Clean Architecture | Estructurar proyecto (domain, usecase, infra).                      |
| 27  | Logs y Config      | Usar `zerolog` o `zap`. Leer `.env` con `os.Getenv`.                |
| 28  | Build & binarios   | `go build`, versiones, flags.                                       |
| 29  | Docker y deploy    | Dockerizar tu API y subirla a Railway o Fly.io.                     |
| 30  | CI/CD & cierre     | Pipeline con GitHub Actions. Documentar y presentar proyecto final. |

📚 Recurso: [https://12factor.net](https://12factor.net) · [https://golangci-lint.run](https://golangci-lint.run)

---

## 🧰 Proyecto final sugerido

👉 **“Task Manager API”**

* CRUD de usuarios y tareas.
* Concurrency para notificaciones simuladas.
* JWT fake + middlewares.
* PostgreSQL como base de datos.
* Tests unitarios e integración.
* Dockerfile + deploy gratuito.

---

## 📝 Consejos para aprovechar el plan

* ✅ No te saltes la semana 3: entender goroutines y channels te separa de un dev junior.
* 🧪 Testea siempre desde semana 4.
* 🐳 Aprende a compilar y dockerizar temprano para evitar bloqueos al final.
* 🧭 Usa el comando `go doc` para explorar funciones sin salir de la terminal.
* 🌱 Lee código open source: por ejemplo, el repo de gin es un excelente ejemplo idiomático.
