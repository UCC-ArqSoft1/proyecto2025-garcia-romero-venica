DROP TABLE IF EXISTS `actividades`;
CREATE TABLE `actividades` (
  `id_actividades` int NOT NULL AUTO_INCREMENT,
  `nombre` varchar(45) DEFAULT NULL,
  `descripcion` varchar(100) DEFAULT NULL,
  `categoria` varchar(100) DEFAULT NULL,
  `estado` tinyint DEFAULT NULL,
  `cupo` int DEFAULT NULL,
  `horarios` varchar(45) DEFAULT NULL,
  `profesor` varchar(45) DEFAULT NULL,
  `disponibles` int DEFAULT NULL,
  PRIMARY KEY (`id_actividades`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `inscripciones`;
CREATE TABLE `inscripciones` (
  `id_inscripciones` int NOT NULL AUTO_INCREMENT,
  `usuarios_id` int DEFAULT NULL,
  `fecha_inscripcion` time DEFAULT NULL,
  `actividades_id` int DEFAULT NULL,
  PRIMARY KEY (`id_inscripciones`),
  KEY `fk_usuario_id_idx` (`usuarios_id`),
  KEY `fk_actividad_id_idx` (`actividades_id`),
  CONSTRAINT `fk_actividad_id` FOREIGN KEY (`actividades_id`) REFERENCES `actividades` (`id_actividades`),
  CONSTRAINT `fk_usuario_id` FOREIGN KEY (`usuarios_id`) REFERENCES `usuarios` (`id_usuarios`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `usuarios`;
CREATE TABLE `usuarios` (
  `id_usuarios` int NOT NULL AUTO_INCREMENT,
  `nombre` varchar(45) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `contraseña_hash` varchar(100) DEFAULT NULL,
  `tipo_usuarios_id` tinyint DEFAULT NULL,
  PRIMARY KEY (`id_usuarios`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `usuarios` WRITE;
INSERT INTO `usuarios` (nombre, email, contraseña_hash, tipo_usuarios_id)
VALUES 
('Andrew', 'andrew@gmail.com', '$2a$10$YFzqg4YYYmAYIW5s7SC88ejRpD4cPhXvkGnF15cw5sKBmHIcBBEna', 1),
('Pablo Socio', 'pablo.socio@example.com', '$2a$10$xHv8X4OBqTj2tvFJ/Gryv.wv3kgnSYTOSqZg48RQk6HiNYU8ITqzu', 1),
('Pablo Admin', 'admin@example.com', '$2a$10$Wb6ID44Eaz9iVxhKYMkRkOaOYr7asVabY/Bz7vnELIRtFbV5eIVTO', 2);
UNLOCK TABLES;
