package main

import (
	"fmt"
	"runtime"

	"github.com/suiqirui1987/fly3d_demo/samples"

	"github.com/suiqirui1987/fly3d/engines"
	"github.com/suiqirui1987/fly3d/windows"
)

func main() {
	runtime.LockOSThread()
	fmt.Print("Engine Example \r\n")

	//core.GlobalFly3D.SetIsDebug(true)

	appoption := &windows.AppOption{
		Width:  600,
		Height: 600,
		HInt:   8,
	}
	app, err := windows.NewApp(appoption)
	if err != nil {
		panic(err)
	}

	engine := engines.NewEngine(app, true)
	scene := engines.NewScene(engine)

	//layer
	//samples.Layer_Sample(scene, app)
	//samples.BoxMesh_Sample(scene, app)
	samples.Shadow_Sample(scene, app)

	engine.RunRenderLoop(func() {
		scene.Render()
	})

	app.Run()

}
