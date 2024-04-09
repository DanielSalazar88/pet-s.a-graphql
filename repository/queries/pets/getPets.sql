SELECT JSON_ARRAYAGG(
    JSON_OBJECT(
        'id', id,
        'nombre', nombre,
        'raza', raza,
        'edad', edad,
        'peso', peso,
        'cedula_cliente', cliente_cedula,
        'nombre_cliente', c.nombres,
        'apellidos_cliente', c.apellidos,
        'telefono_cliente', c.telefono,
        'direccion_cliente', c.direccion,
        'correo_cliente', c.correo
        )
) AS resultado
FROM Mascota
INNER JOIN Cliente c on c.cedula = cliente_cedula;
