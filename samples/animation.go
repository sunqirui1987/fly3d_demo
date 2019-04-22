package samples

import (
	"math"

	"github.com/suiqirui1987/fly3d/engines"
	"github.com/suiqirui1987/fly3d/math32"
	"github.com/suiqirui1987/fly3d/module/animations"
	"github.com/suiqirui1987/fly3d/module/cameras"
	"github.com/suiqirui1987/fly3d/module/lights"
	"github.com/suiqirui1987/fly3d/module/materials"
	"github.com/suiqirui1987/fly3d/module/meshs"
	"github.com/suiqirui1987/fly3d/windows"
)

func Animations_Sample(scene *engines.Scene, app *windows.App) {

	lights.NewPointLight("Omni", math32.NewVector3(0, 2, 8), scene)
	var camera = cameras.NewArcRotateCamera("ArcRotateCamera", 1, 0.8, 20, math32.NewVector3(0, 0, 0), scene)
	camera.AttachControl(app)

	// Fountain object
	var mes = meshs.CreateBox("foutain", 1.0, scene, false)

	// Ground
	var m = materials.NewStandardMaterial("groundMat", scene)
	m.BackFaceCulling = false
	m.DiffuseColor = math32.NewColor3(0.3, 0.3, 1)

	var ground = meshs.CreatePlane("ground", 50.0, scene, false)
	ground.Position = math32.NewVector3(0, -10, 0)
	ground.Rotation = math32.NewVector3(math.Pi/2, 0, 0)
	ground.Material = m

	var animation = animations.NewAnimation("animation", "Rotation.X", 30, animations.ANIMATIONTYPE_FLOAT,
		animations.ANIMATIONLOOPMODE_CYCLE)

	var keys = make([]*animations.AnimationKeyFrame, 0)
	keys = append(keys, &animations.AnimationKeyFrame{
		Frame: 0,
		Value: 0,
	})
	keys = append(keys, &animations.AnimationKeyFrame{
		Frame: 50,
		Value: math.Pi,
	})
	keys = append(keys, &animations.AnimationKeyFrame{
		Frame: 100,
		Value: 0,
	})

	animation.SetKeys(keys)

	mes.AddAnimation(animation)

	animations.BeginAnimation(scene, mes, 0, 100, true, 1.0)
}
