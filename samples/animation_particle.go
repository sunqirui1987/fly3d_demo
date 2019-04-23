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
	"github.com/suiqirui1987/fly3d/module/particles"
	"github.com/suiqirui1987/fly3d/module/textures"
	"github.com/suiqirui1987/fly3d/windows"
)

func Animations_Particle_Sample(scene *engines.Scene, app windows.IWindow) {

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

	// particle system demo
	// Create a particle system
	var particleSystem = particles.NewParticleSystem("particles", 2000, scene)

	//Texture of each particle
	particleSystem.ParticleTexture = textures.NewTexture(ResRepository+"Assets/flare.png", scene, false, 0)

	// Where the particles come from
	particleSystem.Emitter = mes
	particleSystem.EmitRate = 1500

	particleSystem.MinEmitBox = math32.NewVector3(-1, 0, 0)
	particleSystem.MaxEmitBox = math32.NewVector3(1, 0, 0)

	// Colors of all particles
	particleSystem.Color1 = math32.NewColor4(0.7, 0.8, 1.0, 1.0)
	particleSystem.Color2 = math32.NewColor4(0.2, 0.5, 1.0, 1.0)
	particleSystem.ColorDead = math32.NewColor4(0, 0, 0.2, 0.0)

	// Size of each particle (random between...
	particleSystem.MinSize = 0.1
	particleSystem.MaxSize = 0.5

	// Life time of each particle (random between...
	particleSystem.MinLifeTime = 0.3
	particleSystem.MaxLifeTime = 1.5

	// Blend mode : BLENDMODE_ONEONE, or BLENDMODE_STANDARD
	particleSystem.BlendMode = particles.BLENDMODE_ONEONE

	// Set the gravity of all particles
	particleSystem.Gravity = math32.NewVector3(0, -9.81, 0)

	// Direction of each particle after it has been emitted
	particleSystem.Direction1 = math32.NewVector3(-7, 8, 3)
	particleSystem.Direction2 = math32.NewVector3(7, 8, -3)

	// Angular speed, in radians
	particleSystem.MinAngularSpeed = 0
	particleSystem.MaxAngularSpeed = math.Pi

	// Speed
	particleSystem.MinEmitPower = 1
	particleSystem.MaxEmitPower = 3
	particleSystem.UpdateSpeed = 0.005

	// Start the particle system
	particleSystem.Start()

}
