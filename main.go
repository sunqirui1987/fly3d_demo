package main

import (
	"fmt"
	"runtime"

	log "github.com/sirupsen/logrus"
	"github.com/suiqirui1987/fly3d_demo/samples"

	"github.com/suiqirui1987/fly3d/core"
	"github.com/suiqirui1987/fly3d/engines"
	"github.com/suiqirui1987/fly3d/windows"
)

func main() {
	runtime.LockOSThread()
	fmt.Print("Engine Example \r\n")

	core.GlobalFly3D.ResRepository = "/Users/qiruisun/gitlab/golang-fly3d-engine/"
	if core.GlobalFly3D.IsDebug {
		log.SetLevel(log.DebugLevel)
	}

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
