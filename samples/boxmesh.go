package samples

import (
	"github.com/suiqirui1987/fly3d/core"
	"github.com/suiqirui1987/fly3d/engines"
	"github.com/suiqirui1987/fly3d/math32"
	"github.com/suiqirui1987/fly3d/module/cameras"
	"github.com/suiqirui1987/fly3d/module/lights"
	"github.com/suiqirui1987/fly3d/module/materials"
	"github.com/suiqirui1987/fly3d/module/meshs"
	"github.com/suiqirui1987/fly3d/windows"
)

func BoxMesh_Sample(scene *engines.Scene, app *windows.App) {

	core.GlobalFly3D.ResRepository = "/Users/qiruisun/gitlab/golang-fly3d-engine/"

	camera := cameras.NewArcRotateCamera("Camera", 0, 0, 10, math32.NewVector3(0, 0, 0), scene)

	lights.NewPointLight("Omni", math32.NewVector3(20, 100, 2), scene)

	box := meshs.CreateBox("box", 2.0, scene, false)
	material := materials.NewStandardMaterial("kosh", scene)
	//material.Wireframe = true
	//material.OpacityTexture = textures.NewTexture(core.GlobalFly3D.ResRepository+"example/Assets/normalMap.jpg", scene, false, 0)
	material.DiffuseColor = math32.NewColor3(1, 0, 0)

	box.Material = material

	camera.SetPosition(math32.NewVector3(-3, 3, 0))
	// This attaches the camera to the canvas
	camera.AttachControl(app)
}
