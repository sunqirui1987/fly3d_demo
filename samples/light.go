package samples

import (
	"github.com/suiqirui1987/fly3d/engines"
	"github.com/suiqirui1987/fly3d/math32"
	"github.com/suiqirui1987/fly3d/module/cameras"
	"github.com/suiqirui1987/fly3d/module/lights"
	"github.com/suiqirui1987/fly3d/module/materials"
	"github.com/suiqirui1987/fly3d/module/meshs"
	"github.com/suiqirui1987/fly3d/module/textures"
	"github.com/suiqirui1987/fly3d/windows"
)

func Light_Sample(scene *engines.Scene, app windows.IWindow) {

	var camera = cameras.NewArcRotateCamera("Camera", 0, 0, 10, math32.NewVector3Zero(), scene)
	camera.SetPosition(math32.NewVector3(-10, 10, 0))
	camera.AttachControl(app)

	var light0 = lights.NewPointLight("Omni0", math32.NewVector3(0, 10, 0), scene)
	// Lights colors
	light0.Diffuse = math32.NewColor3(1.0, 0.0, 0.0)
	light0.Specular = math32.NewColor3(1.0, 0.0, 0.0)

	var light1 = lights.NewPointLight("Omni1", math32.NewVector3(0, -10, 0), scene)
	var light2 = lights.NewPointLight("Omni2", math32.NewVector3(10, 0, 0), scene)
	var light3 = lights.NewDirectionalLight("Dir0", math32.NewVector3(1, -1, 0), scene)

	light1.Diffuse = math32.NewColor3(0, 1, 0)
	light1.Specular = math32.NewColor3(0, 1, 0)

	light2.Diffuse = math32.NewColor3(0, 0, 1)
	light2.Specular = math32.NewColor3(0, 0, 1)

	light3.Diffuse = math32.NewColor3(1, 1, 1)
	light3.Specular = math32.NewColor3(1, 1, 1)

	// Creating light sphere
	var lightSphere0 = meshs.CreateSphere("Sphere0", 16, 0.5, scene, false)
	var lightSphere1 = meshs.CreateSphere("Sphere1", 16, 0.5, scene, false)
	var lightSphere2 = meshs.CreateSphere("Sphere2", 16, 0.5, scene, false)

	m0 := materials.NewStandardMaterial("red", scene)
	m0.DiffuseColor = math32.NewColor3(0, 0, 0)
	m0.SpecularColor = math32.NewColor3(0, 0, 0)
	m0.EmissiveColor = math32.NewColor3(1, 0, 0)

	lightSphere0.Material = m0

	m1 := materials.NewStandardMaterial("green", scene)
	m1.DiffuseColor = math32.NewColor3(0, 0, 0)
	m1.SpecularColor = math32.NewColor3(0, 0, 0)
	m1.EmissiveColor = math32.NewColor3(0, 1, 0)

	lightSphere1.Material = m1

	m2 := materials.NewStandardMaterial("blue", scene)
	m2.DiffuseColor = math32.NewColor3(0, 0, 0)
	m2.SpecularColor = math32.NewColor3(0, 0, 0)
	m2.EmissiveColor = math32.NewColor3(0, 0, 1)

	lightSphere2.Material = m2

	// Sphere material

	var sphere = meshs.CreateSphere("Sphere", 16, 3, scene, false)
	var material = materials.NewStandardMaterial("kosh", scene)
	material.DiffuseColor = math32.NewColor3(1, 1, 1)
	sphere.Material = material

	tex := textures.NewCubeTexture(ResRepository+"Assets/skybox/skybox", scene, nil)
	tex.CoordinatesMode = textures.SKYBOX_MODE
	// Skybox
	var skybox = meshs.CreateBox("skyBox", 100.0, scene, false)
	var skyboxMaterial = materials.NewStandardMaterial("skyBox", scene)
	skyboxMaterial.BackFaceCulling = false
	skyboxMaterial.ReflectionTexture = tex
	skyboxMaterial.DiffuseColor = math32.NewColor3(0, 0, 0)
	skyboxMaterial.SpecularColor = math32.NewColor3(0, 0, 0)
	skybox.Material = skyboxMaterial

	// Animations
	var alpha float32 = 0.0
	scene.RegisterBeforeRender(func() {

		light0.Position = math32.NewVector3(10*math32.Sin(alpha), 0, 10*math32.Cos(alpha))
		light1.Position = math32.NewVector3(10*math32.Sin(alpha), 0, -10*math32.Cos(alpha))
		light2.Position = math32.NewVector3(10*math32.Cos(alpha), 0, 10*math32.Sin(alpha))

		lightSphere0.Position = light0.Position
		lightSphere1.Position = light1.Position
		lightSphere2.Position = light2.Position

		alpha += 0.01
	})

}
