package rosengo

type GameObject interface {
	Update() error
	Draw() error
}

type gameObject struct {
}
