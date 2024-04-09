SELECT
    JSON_ARRAYAGG(
        JSON_OBJECT(
            'cedula_cliente',
            c.cedula,
            'nombres_cliente',
            c.nombres,
            'apellidos_cliente',
            c.apellidos,
            'mascotas',
            IFNULL(mascotas_json, JSON_ARRAY())
        )
    ) AS resultado
FROM
    Cliente c
    LEFT JOIN (
        SELECT
            m.cliente_cedula,
            JSON_ARRAYAGG(
                JSON_OBJECT(
                    'nombre_mascota',
                    m.nombre,
                    'descripcion_receta',
                    IFNULL(m2.descripcion, ''),
                    'dosis_receta',
                    IFNULL(m2.dosis, '')
                )
            ) AS mascotas_json
        FROM
            Mascota m
            INNER JOIN Receta r ON r.mascota_id = m.id
            INNER JOIN Medicamento m2 ON m2.id = r.medicamento_id
        GROUP BY
            m.cliente_cedula
    ) AS mascotas_agrupadas ON mascotas_agrupadas.cliente_cedula = c.cedula
WHERE
    c.cedula = 'cedula_replace';