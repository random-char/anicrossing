package scene

type Scene interface {
	Init()
    Update()
    Render()
	Teardown()
	GetChildren() []Scene
}
