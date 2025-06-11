-- Usar la base de datos gimnasio
USE gimnasio;

-- Borrar tablas si ya existen
DROP TABLE IF EXISTS inscripciones;
DROP TABLE IF EXISTS actividades;
DROP TABLE IF EXISTS usuarios;

-- Crear tabla usuarios
CREATE TABLE usuarios (
  id_usuarios INT NOT NULL AUTO_INCREMENT,
  nombre VARCHAR(45) NOT NULL,
  email VARCHAR(45) NOT NULL,
  password VARCHAR(100) NOT NULL,
  tipo_usuarios_id TINYINT NOT NULL,
  PRIMARY KEY (id_usuarios)
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
  usuarios_id INT,
  fecha_inscripcion DATE,
  actividades_id INT,
  PRIMARY KEY (id_inscripciones),
  FOREIGN KEY (usuarios_id) REFERENCES usuarios(id_usuarios),
  FOREIGN KEY (actividades_id) REFERENCES actividades(id_actividades)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Insertar usuario de prueba (admin)
INSERT INTO usuarios (nombre, email, password, tipo_usuarios_id)
VALUES ('admin', 'admin@gmail.com', '14535', 1);
