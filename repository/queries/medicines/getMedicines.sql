SELECT JSON_ARRAYAGG(
    JSON_OBJECT(
        'id', id,
        'nombre', nombre,
        'descripcion', descripcion,
        'dosis', dosis
    )
) AS resultado
FROM Medicamento;
