SELECT JSON_ARRAYAGG(
    JSON_OBJECT(
        'id', r.id,
        'medicamento', m2.nombre,
        'dosis', m2.dosis,
        'descripcion', m2.descripcion,
        'nombre_mascota', m.nombre,
        'cedula_cliente', m.cliente_cedula
    )
) AS resultado
FROM Receta r 
inner join Mascota m on r.mascota_id = m.id 
inner join Medicamento m2 on r.medicamento_id = m2.id 