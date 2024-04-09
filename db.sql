-- Crear la tabla Cliente
CREATE TABLE Cliente (
    cedula VARCHAR(20) PRIMARY KEY,
    nombres VARCHAR(100),
    apellidos VARCHAR(100),
    direccion VARCHAR(255),
    telefono VARCHAR(20),
    correo varchar(50)
);


-- Crear la tabla Medicamento
CREATE TABLE Medicamento (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100),
    descripcion TEXT,
    dosis VARCHAR(50)
);

-- Crear la tabla Mascota
CREATE TABLE Mascota (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100),
    raza VARCHAR(100),
    edad INT,
    peso FLOAT,
    cliente_cedula VARCHAR(20),
    FOREIGN KEY (cliente_cedula) REFERENCES Cliente(cedula)
);

-- Crear la tabla Receta
CREATE TABLE Receta (
    id INT AUTO_INCREMENT PRIMARY KEY,
    mascota_id INT,
    medicamento_id INT,
    FOREIGN KEY (mascota_id) REFERENCES Mascota(id),
    FOREIGN KEY (medicamento_id) REFERENCES Medicamento(id)
);