package rango

type Light struct {
	Positon Vector
	Color   Color   /* Light color, (1.0f, 1.0f, 1.0f) in most cases */
	Shadow  float64 /* Shadow factor, 0.0f means no shadow, 1.0f means solid black shadow */
}

func setLight(light *Light, position Vector, color Color, shadow float64) *Light {

	light.Positon = position
	light.Color = color
	light.Shadow = shadow

	return light
}
