package rango

import "math"

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
func Specular(hit Hit, scene Scene, light Light) Vector {
	objId := hit.ObjectId

	reflect := ReflectRay(hit)
	lightDir := Normalize(Subtract(light.Positon, hit.Position))
	t := math.Pow(Dot(lightDir, reflect.Dir), scene.Objects[objId].Material.Shininess) * scene.Objects[objId].Material.Kspecular

	return FloatVecMult(t, MultiplyColorVector(scene.Objects[objId].Material.Color, light.Color))
}
