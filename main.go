package main

import (
	"fmt"
	"runtime"

	"github.com/suiqirui1987/fly3d/engines"
	"github.com/suiqirui1987/fly3d/gui/nanogui"
	"github.com/suiqirui1987/fly3d_demo/samples"

	"github.com/suiqirui1987/fly3d/windows"
)

var engine *engines.Engine
var scene *engines.Scene
var app *windows.App

func DemoButton(screen *nanogui.Screen) {

	window := nanogui.NewWindow(screen, "Demo")
	window.SetPosition(15, 15)
	window.SetLayout(nanogui.NewGroupLayout())

	b1 := nanogui.NewButton(window, "Animations_Particle_Sample")
	b1.SetCallback(func() {
		if scene != nil {
			scene.Dispose()
		}

		scene = engines.NewScene(engine)
		samples.Animations_Particle_Sample(scene, app)
	})

	b2 := nanogui.NewButton(window, "Shadow_Sample")
	b2.SetCallback(func() {
		if scene != nil {
			scene.Dispose()
		}

		scene = engines.NewScene(engine)
		samples.Shadow_Sample(scene, app)
	})

	b3 := nanogui.NewButton(window, "Light_Sample")
	b3.SetCallback(func() {
		if scene != nil {
			scene.Dispose()
		}

		scene = engines.NewScene(engine)
		samples.Light_Sample(scene, app)
	})

	b4 := nanogui.NewButton(window, "Fog_Sample")
	b4.SetCallback(func() {
		if scene != nil {
			scene.Dispose()
		}

		scene = engines.NewScene(engine)
		samples.Fog_Sample(scene, app)
	})

	b5 := nanogui.NewButton(window, "Camera_Collisions_Sample")
	b5.SetCallback(func() {
		if scene != nil {
			scene.Dispose()
		}

		scene = engines.NewScene(engine)
		samples.Camera_Collisions_Sample(scene, app)
	})

}

func main() {
	runtime.LockOSThread()
	fmt.Print("Engine Example \r\n")

	//core.GlobalFly3D.SetIsDebug(true)

	appoption := &windows.AppOption{
		Width:  800,
		Height: 600,
	}
	var err error
	app, err = windows.NewApp(appoption)
	if err != nil {
		panic(err)
	}

	DemoButton(app.Screen)
	app.Screen.PerformLayout()

	engine = engines.NewEngine(app, true)
	scene = engines.NewScene(engine)

	//layer
	//samples.Layer_Sample(scene, app)
	//samples.BoxMesh_Sample(scene, app)
	//samples.Shadow_Sample(scene, app)
	//samples.Light_Sample(scene, app)
	//samples.Fog_Sample(scene, app)
	//samples.Camera_Collisions_Sample(scene, app)
	//samples.Animations_Sample(scene, app)
	samples.Animations_Particle_Sample(scene, app)

	engine.RunRenderLoop(func() {

		if scene != nil {
			scene.Render()
		}
		engine.WipeCaches() // for draw ui
	})

	app.Run()

}
