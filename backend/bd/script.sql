-- Eliminar base existente (opcional)
DROP DATABASE IF EXISTS gimnasio;

-- Crear base de datos
CREATE DATABASE gimnasio;
USE gimnasio;

-- Crear tabla usuarios
CREATE TABLE usuarios (
  id_usuarios INT NOT NULL AUTO_INCREMENT,
  nombre VARCHAR(45) NOT NULL,
  email VARCHAR(45) NOT NULL,
  password VARCHAR(100) NOT NULL,
  tipo_usuarios_id VARCHAR(45) NOT NULL,
  PRIMARY KEY (id_usuarios),
  UNIQUE (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Crear tabla actividades
CREATE TABLE actividades (
  id_actividades INT NOT NULL AUTO_INCREMENT,
  nombre VARCHAR(45) NOT NULL,
  descripcion VARCHAR(100),
  categoria VARCHAR(100),
  estado TINYINT DEFAULT 1,
  cupo INT,
  horarios VARCHAR(45),
  profesor VARCHAR(45),
  disponibles INT,
  PRIMARY KEY (id_actividades)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Crear tabla inscripciones
CREATE TABLE inscripciones (
  id_inscripciones INT NOT NULL AUTO_INCREMENT,
  usuarios_id INT NOT NULL,
  fecha_inscripcion DATE,
  actividades_id INT NOT NULL,
  PRIMARY KEY (id_inscripciones),
  FOREIGN KEY (usuarios_id) REFERENCES usuarios(id_usuarios),
  FOREIGN KEY (actividades_id) REFERENCES actividades(id_actividades)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Insertar usuarios con contraseñas hasheadas (SHA256)
-- admin -> password
-- socio -> password

INSERT INTO usuarios (nombre, email, password, tipo_usuarios_id)
VALUES 
('admin', 'admin@gmail.com', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', 'admin'),
('socio', 'socio@gmail.com', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', 'socio');
