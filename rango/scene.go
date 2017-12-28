package rango

type Scene struct {
	NObjects uint64   /* Number of objects in the scene */
	NMax     uint64   /* Max number of objects scene can hold */
	Objects  []Object /* Object list of the Scene */
}

func SetScene() {

}
