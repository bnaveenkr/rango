package rango

type Scene struct {
	NObjects uint64   /* Number of objects in the scene */
	NMax     uint64   /* Max number of objects scene can hold */
	Objects  []Object /* Object list of the Scene */
}

func InitScene(scene *Scene, maxObjects int8) *Scene {
	scene.NObjects = 0
	scene.NMax = uint64(maxObjects)
	scene.Objects = make([]Object, 0)
	return scene
}

func AddObjectsToScene(scene *Scene, object Object) *Scene {

	if scene.NObjects < scene.NMax {
		scene.Objects = append(scene.Objects, object)
		scene.NObjects = scene.NObjects + 1
	}
	return scene
}

func SetScene() {

}
