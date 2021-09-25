package rosengo

type Gameobject interface {
	Update() error
	Draw() error
}
