package samples

import (
	"github.com/suiqirui1987/fly3d/engines"
	"github.com/suiqirui1987/fly3d/math32"
	"github.com/suiqirui1987/fly3d/module/cameras"
	"github.com/suiqirui1987/fly3d/module/lights"
	"github.com/suiqirui1987/fly3d/module/meshs"
	"github.com/suiqirui1987/fly3d/windows"
)

func Shadow_Sample(scene *engines.Scene, app *windows.App) {

	camera := cameras.NewArcRotateCamera("Camera", 0, 0, 10, math32.NewVector3Zero(), scene)
	camera.AttachControl(app)

	light := lights.NewDirectionalLight("dir01", math32.NewVector3(0, -1, -0.2), scene)
	light.Intensity = 0.6
	/*
		light2 := lights.NewDirectionalLight("dir02", math32.NewVector3(-1, -2, -1), scene)
		light2.Position = math32.NewVector3(10, 20, 10)
		light2.Intensity = 0.6
	*/

	camera.SetPosition(math32.NewVector3(-20, 20, 0))
	/*
		tex := textures.NewCubeTexture(core.GlobalFly3D.ResRepository+"example/Assets/skybox/night", scene, nil)
		tex.CoordinatesMode = textures.SKYBOX_MODE

		var skybox *meshs.Mesh
		// Skybox
		skybox = meshs.CreateBox("skyBox", 1000.0, scene, false)

		skyboxMaterial := materials.NewStandardMaterial("skyBox", scene)
		skyboxMaterial.BackFaceCulling = false
		skyboxMaterial.ReflectionTexture = tex
		skyboxMaterial.DiffuseColor = math32.NewColor3(0, 0, 0)
		skyboxMaterial.SpecularColor = math32.NewColor3(0, 0, 0)
		skybox.Material = skyboxMaterial

		// Ground
		tex2 := textures.NewTexture(core.GlobalFly3D.ResRepository+"example/Assets/grass.jpg", scene, false, 0)
		tex2.UScale = 60
		tex2.VScale = 60

		ground := meshs.CreateGround("ground", 1000, 1000, 1, scene, false)
		groundMaterial := materials.NewStandardMaterial("ground", scene)
		groundMaterial.BackFaceCulling = false
		groundMaterial.DiffuseTexture = tex2
		groundMaterial.DiffuseColor = math32.NewColor3(0, 0, 0)
		groundMaterial.SpecularColor = math32.NewColor3(0, 0, 0)
		//groundMaterial.Wireframe = true
		ground.Position.Y = -2.05
		ground.Material = groundMaterial
	*/
	// Torus
	torus := meshs.CreateTorus("torus", 8, 2, 32, scene, false)
	torus.Position.Y = 6.0
	//torus2 := meshs.CreateTorus("torus2", 4, 1, 32, scene, false)
	//torus2.Position.Y = 6.0

}
