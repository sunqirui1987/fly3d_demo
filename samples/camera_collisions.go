package samples

import (
	"github.com/suiqirui1987/fly3d/engines"
	"github.com/suiqirui1987/fly3d/interfaces"
	"github.com/suiqirui1987/fly3d/math32"
	"github.com/suiqirui1987/fly3d/module/cameras"
	"github.com/suiqirui1987/fly3d/module/lights"
	"github.com/suiqirui1987/fly3d/module/materials"
	"github.com/suiqirui1987/fly3d/module/meshs"
	"github.com/suiqirui1987/fly3d/module/textures"
	"github.com/suiqirui1987/fly3d/windows"
)

func Camera_Collisions_Sample(scene *engines.Scene, app windows.IWindow) {

	// Lights
	lights.NewDirectionalLight("Omni", math32.NewVector3(-2, -5, 2), scene)
	lights.NewPointLight("Omni", math32.NewVector3(2, -5, -2), scene)

	// Need a free camera for collisions
	var camera = cameras.NewFreeCamera("FreeCamera", math32.NewVector3(0, -8, -20), scene)
	camera.AttachControl(app)

	var tex = textures.NewTexture(ResRepository+"Assets/t.Png", scene, false, 0)
	tex.EnableHasAlpha(true)

	var m = materials.NewStandardMaterial("Mat", scene)
	m.DiffuseTexture = tex
	//Simple crate
	var box = meshs.CreateBox("box", 2, scene, false)
	box.Material = m
	box.Position = math32.NewVector3(5, -9, -10)

	// Enable Collisions
	scene.CollisionsEnabled = true

	//Then apply collisions and gravity to the active camera
	camera.CheckCollisions = true
	//Set the ellipsoid around the camera
	camera.Ellipsoid = math32.NewVector3(1, 1, 1)

	//finally,
	box.Checkcollisions = true

	camera.OnCollide = func(colmesh interfaces.IMesh) {
		if colmesh.GetName() == "box" {
			//camera.Position = math32.NewVector3(0, -8, -20)
		}
	}

}
