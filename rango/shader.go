package rango

func Ambinet(hit Hit, scene Scene, light Light) Vector {
	objId := hit.ObjectId

	return FloatVecMult(scene.Objects[objId].Material.Kambient, MultiplyColorVector(scene.Objects[objId].Material.Color, light.Color))
}

func Diffuse(hit Hit, scene Scene, light Light) Vector {
	objId := hit.ObjectId

	lightDirection := Normalize(Subtract(light.Positon, hit.Position))
	t := Dot(hit.Normal, lightDirection) * scene.Objects[objId].Material.Kdiffuse

	return FloatVecMult(t, MultiplyColorVector(scene.Objects[objId].Material.Color, light.Color))

}
