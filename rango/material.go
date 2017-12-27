package rango

type Material struct {
	Color        Color   /* Color of the object */
	Kambient     float64 /* Ambient constant */
	Kdiffuse     float64 /* Diffuse constant */
	Kspecular    float64 /* Specular constant */
	Shininess    float64 /* Exponent of specular component */
	Reflectivity float64 /* Amount of reflectivity of the surface, 0.0f means a matt object, 1.0f means a mirror */
	Traslucency  float64 /* Amount of translucency of the surface, 0.0f means opaque object, 1.0f means a Saint Gobain's glass */
	Ir           float64 /* Index of refraction of a translucent object */
}

func SetMaterial(material *Material, color Color, ambient float64, diffuse float64, specular float64, shininess float64, reflectivity float64, translucency float64, ir float64) *Material {

	material.Color = color
	material.Kambient = ambient
	material.Kdiffuse = diffuse
	material.Kspecular = specular
	material.Shininess = shininess
	material.Reflectivity = reflectivity
	material.Traslucency = translucency
	material.Ir = ir

	return material
}
