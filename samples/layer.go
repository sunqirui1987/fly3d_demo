package samples

import (
	"github.com/suiqirui1987/fly3d/engines"
	"github.com/suiqirui1987/fly3d/math32"
	"github.com/suiqirui1987/fly3d/module/cameras"
	"github.com/suiqirui1987/fly3d/module/layers"
	"github.com/suiqirui1987/fly3d/windows"
)

func Layer_Sample(scene *engines.Scene, app windows.IWindow) {

	camera := cameras.NewArcRotateCamera("Camera", 0, 0, 10, math32.NewVector3(0, 0, 0), scene)

	// Positions the camera overwriting alpha, beta, radius
	camera.SetPosition(math32.NewVector3(0, 0, 20))

	// This attaches the camera to the canvas
	camera.AttachControl(app)

	layers.NewLayer("1", ResRepository+"Assets/Layer0_0.png", scene, true, nil)

}
