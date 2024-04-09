SELECT
    IFNULL(
        JSON_ARRAYAGG(
            JSON_OBJECT(
                'nombre_mascota',
                m3.nombre,
                'raza_mascota',
                COALESCE(m3.raza, ''),
                'peso_mascota',
                COALESCE(m3.peso, ''),
                'edad_mascota',
                COALESCE(m3.edad, ''),
                'recetas',
                COALESCE(recetas_json, JSON_ARRAY())
            )
        ),
        JSON_ARRAY()
    ) AS resultado
FROM
    Mascota m3
    LEFT JOIN (
        SELECT
            r.mascota_id,
            JSON_ARRAYAGG(
                JSON_OBJECT(
                    'medicamento',
                    m4.nombre,
                    'descripcion_receta',
                    COALESCE(m4.descripcion, ''),
                    'dosis_receta',
                    COALESCE(m4.dosis, '')
                )
            ) AS recetas_json
        FROM
            Receta r
            INNER JOIN Medicamento m4 ON r.medicamento_id = m4.id
        GROUP BY
            r.mascota_id
    ) AS recetas_agrupadas ON recetas_agrupadas.mascota_id = m3.id
WHERE
    m3.id = id_mascota_replace;