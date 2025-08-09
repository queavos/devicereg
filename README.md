🌀  Local Token Server

Un servidor HTTP minimalista escrito en Go que expone un token de autenticación junto con el hostname del sistema. Diseñado para integrarse con frontends locales y facilitar pruebas de comunicación entre capas, respetando políticas CORS.


🚀 Características

Expone un endpoint GET /token en http://localhost:8085

Devuelve un JSON con:

token: valor personalizado

hostname: nombre del sistema

status: estado de respuesta

Soporte completo para CORS (incluye manejo de preflight OPTIONS)

Ideal para pruebas locales con frontends como React, Vue, etc.


Asegurate de tener Go instalado (go version) y configurado correctamente.

🧪 Ejemplo de respuesta

{
  "token": "abc123",
  "hostname": "ozba-machine",
  "status": "ok"
}

🌐 CORS

Este servidor incluye los encabezados necesarios para permitir peticiones desde cualquier origen:

Access-Control-Allow-Origin: *
Access-Control-Allow-Methods: GET, OPTIONS
Access-Control-Allow-Headers: Content-Type

Para mayor seguridad, podés restringir el origen modificando el valor de Access-Control-Allow-Origin.

🧠 Filosofía

Este proyecto forma parte de una infraestructura simbólica que combina precisión técnica con narrativa digital. Cada componente refleja una intención: conectar, proteger y representar.

🛠️ Futuras mejoras

Validación de tokens

Configuración dinámica del puerto y origen CORS
