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

func Fog_Sample(scene *engines.Scene, app windows.IWindow) {

	var camera = cameras.NewFreeCamera("Camera", math32.NewVector3(0, 0, -20), scene)
	camera.SetTarget(math32.NewVector3(0, 0, 0))
	camera.AttachControl(app)

	lights.NewPointLight("Omni", math32.NewVector3(20, 100, 2), scene)

	var sphere0 = meshs.CreateSphere("Sphere0", 16, 3, scene, false)
	var sphere1 = meshs.CreateSphere("Sphere1", 16, 3, scene, false)
	var sphere2 = meshs.CreateSphere("Sphere2", 16, 3, scene, false)

	var material0 = materials.NewStandardMaterial("mat0", scene)
	material0.DiffuseColor = math32.NewColor3(1, 0, 0)
	sphere0.Material = material0
	sphere0.Position = math32.NewVector3(-10, 0, 0)

	var material1 = materials.NewStandardMaterial("mat1", scene)
	material1.DiffuseColor = math32.NewColor3(1, 1, 0)
	sphere1.Material = material1

	var material2 = materials.NewStandardMaterial("mat2", scene)
	material2.DiffuseColor = math32.NewColor3(1, 0, 1)
	sphere2.Material = material2
	sphere2.Position = math32.NewVector3(10, 0, 0)

	// Fog
	scene.FogMode = core.FOGMODE_EXP
	scene.FogDensity = 0.1

	// Animations
	var alpha float32 = 0.0
	scene.RegisterBeforeRender(func() {
		sphere0.Position.Z = 1.3 * math32.Cos(alpha)
		sphere1.Position.Z = 1.3 * math32.Sin(alpha)
		sphere2.Position.Z = 1.3 * math32.Cos(alpha)

		alpha += 0.1
	})

}
