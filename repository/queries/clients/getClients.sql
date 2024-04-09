SELECT JSON_ARRAYAGG(
    JSON_OBJECT(
        'cedula', cedula,
        'nombres', nombres,
        'apellidos', apellidos,
        'direccion', direccion,
        'telefono', telefono,
        'correo', correo
    )
) AS resultado
FROM Cliente;
