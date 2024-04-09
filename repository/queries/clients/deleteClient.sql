-- Primero, elimina las filas relacionadas en la tabla Receta que hacen referencia a la Mascota que deseas eliminar
DELETE FROM Receta 
WHERE mascota_id IN (
    SELECT id FROM Mascota 
    WHERE cliente_cedula = 'cedula_replace'
    );

-- Luego, elimina la fila de la Mascota
DELETE FROM Mascota 
WHERE cliente_cedula = 'cedula_replace';

-- Finalmente, elimina la fila del Cliente
DELETE FROM Cliente 
WHERE cedula = 'cedula_replace';
