# Proyecto: Comprendiendo el Patrón Pub/Sub

Este proyecto, construido en Golang, pretende servir como ejemplo y guía para el uso del patrón de diseño de Publicador/Suscriptor (Pub/Sub) utilizando Google Cloud Platform (GCP).

## Descripción

El proyecto consta de una API Gateway y dos microservicios: uno para manejar salas y otro para usuarios. Cada microservicio expone una API para atender las solicitudes de manera concurrente, al mismo tiempo que está pendiente de los mensajes que recibe a través de suscripciones.

## Tecnologías Utilizadas

- **Golang:** Lenguaje de programación principal del proyecto.
- **Google Cloud Platform (GCP):** Plataforma en la nube utilizada para implementar el patrón Pub/Sub.
- **API Gateway:** Actúa como intermediario entre los clientes y los microservicios.
- **Microservicios:** Dos microservicios separados, uno para el manejo de salas y otro para usuarios.

## Funcionalidades Principales

1. **API Gateway:**
   - Recibe las solicitudes de los clientes y las enruta a los microservicios correspondientes.

2. **Microservicio de Salas:**
   - Maneja las operaciones relacionadas con las salas, como crear, actualizar y eliminar salas.
   - Está pendiente de los mensajes recibidos a través de suscripciones Pub/Sub relacionadas con las salas.

3. **Microservicio de Usuarios:**
   - Maneja las operaciones relacionadas con los usuarios, como registrar, autenticar y actualizar usuarios.
   - Está pendiente de los mensajes recibidos a través de suscripciones Pub/Sub relacionadas con los usuarios.

## Instalación y Uso

1. **Clonar el Repositorio:**

2. **Configuración de GCP:**
- Configurar una cuenta en Google Cloud Platform.
- Crear un proyecto en GCP y habilitar los servicios necesarios, como Pub/Sub.
- Configurar las credenciales en el proyecto local.

3. **Ejecutar el Proyecto:**
- Iniciar la API Gateway y los microservicios.
- Realizar solicitudes a la API Gateway para interactuar con los microservicios.

## Contribución

¡Las contribuciones son bienvenidas! Si deseas contribuir a este proyecto, sigue estos pasos:

1. Haz un fork del proyecto.
2. Crea una nueva rama (`git checkout -b feature/nueva-funcionalidad`).
3. Realiza tus cambios y haz commit (`git commit -am 'Agrega una nueva funcionalidad'`).
4. Haz push a la rama (`git push origin feature/nueva-funcionalidad`).
5. Crea un nuevo Pull Request.

## Licencia

Este proyecto está bajo la Licencia [MIT](https://opensource.org/licenses/MIT).
