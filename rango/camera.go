package rango

type Camera struct {
	Position Vector
	Look     Vector
	Hori     Vector
	Vert     Vector
	Up       Vector
	Fov2     float64
	Ar       float64
	Width    float64
	Height   float64
}

func SetCamera(camera *Camera, position Vector, lookat Vector, fov float64, width float64, height float64) *Camera {

	/* Camera assumes up Vector is (0,1,0) */
	up := V(0, 1, 0)
	look := Normalize(Subtract(lookat, position))
	hori := Cross(look, up)
	vert := Cross(hori, look)

	camera.Position = position
	camera.Look = look
	camera.Hori = hori
	camera.Vert = vert
	camera.Up = up

	camera.Fov2 = deg2rad(fov * 0.5)
	camera.Ar = width / height
	camera.Width = width
	camera.Height = height

	return camera
}
