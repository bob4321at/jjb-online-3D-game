package scenes

type Scene interface {
	Setup()
	Update()
	Draw()
	GetSetup() bool
}

var Scenes = []Scene{&MainMenuScene{}, &GameScene{}}
var Current_Scene_Id = 0
