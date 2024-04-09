insert into Receta (
	mascota_id,
	medicamento_id 
)
select m.id, me.id 
from Mascota m, Medicamento me
where m.id = id_mascota_replace and me.id = id_medicamento_replace; 

