package main

import (
	"fmt"
	"sync"
	"time"

	//"github.com/shibukawa/glfw"
	"github.com/suiqirui1987/fly3d/glfw"
	"github.com/suiqirui1987/fly3d/gui/canvas"
	"github.com/suiqirui1987/fly3d/gui/nanogui"
)

type Application struct {
	screen *nanogui.Screen
}

func DemoButton(screen *nanogui.Screen) {

	window := nanogui.NewWindow(screen, "Button demo")
	window.SetPosition(15, 15)
	window.SetLayout(nanogui.NewGroupLayout())

	nanogui.NewLabel(window, "Push buttons").SetFont("sans-bold")

	b1 := nanogui.NewButton(window, "Plain button")
	b1.SetCallback(func() {
		fmt.Println("pushed!")
	})

	b2 := nanogui.NewButton(window, "Styled")
	b2.SetBackgroundColor(canvas.RGBA(0, 0, 255, 25))
	b2.SetIcon(nanogui.IconRocket)
	b2.SetCallback(func() {
		fmt.Println("pushed!")
	})

	nanogui.NewLabel(window, "Toggle button").SetFont("sans-bold")
	b3 := nanogui.NewButton(window, "Toggle me")
	b3.SetFlags(nanogui.ToggleButtonType)
	b3.SetChangeCallback(func(state bool) {
		fmt.Println("Toggle button state:", state)
	})

	nanogui.NewLabel(window, "Radio buttons").SetFont("sans-bold")
	b4 := nanogui.NewButton(window, "Radio button 1")
	b4.SetFlags(nanogui.RadioButtonType)
	b5 := nanogui.NewButton(window, "Radio button 2")
	b5.SetFlags(nanogui.RadioButtonType)

	nanogui.NewLabel(window, "A tool palette").SetFont("sans-bold")
	tools := nanogui.NewWidget(window)
	tools.SetLayout(nanogui.NewBoxLayout(nanogui.Horizontal, nanogui.Middle, 0, 6))

	nanogui.NewToolButton(tools, nanogui.IconCloud)
	nanogui.NewToolButton(tools, nanogui.IconFastForward)
	nanogui.NewToolButton(tools, nanogui.IconCompass)
	nanogui.NewToolButton(tools, nanogui.IconInstall)

	nanogui.NewLabel(window, "Popup buttons").SetFont("sans-bold")
	b6 := nanogui.NewPopupButton(window, "Popup")
	b6.SetIcon(nanogui.IconExport)
	popup := b6.Popup()
	popup.SetLayout(nanogui.NewGroupLayout())

	nanogui.NewLabel(popup, "Arbitrary widgets can be placed here").SetFont("sans-bold")
	nanogui.NewCheckBox(popup, "A check box")
	b7 := nanogui.NewPopupButton(popup, "Recursive popup")
	b7.SetIcon(nanogui.IconFlash)
	popup2 := b7.Popup()

	popup2.SetLayout(nanogui.NewGroupLayout())
	nanogui.NewCheckBox(popup2, "Another check box")
}

func (a *Application) init() {
	//glfw.WindowHint(glfw.Samples, 4)
	a.screen = nanogui.NewScreen(600, 600, "NanoGUI.Go Test", false, false)

	//a.screen.NVGContext().CreateFont("japanese", "font/GenShinGothic-P-Regular.ttf")

	DemoButton(a.screen)

	//a.screen.SetDrawContentsCallback(func() {

	//})

	a.screen.PerformLayout()
	//a.screen.DebugPrint()

}

var mainloopActive bool = false

func (a *Application) Run() {

	defer glfw.Terminate()

	mainloopActive = true

	var wg sync.WaitGroup

	/* If there are no mouse/keyboard events, try to refresh the
	view roughly every 50 ms; this is to support animations
	such as progress bars while keeping the system load
	reasonably low */
	wg.Add(1)
	go func() {
		for mainloopActive {
			time.Sleep(50 * time.Millisecond)
			glfw.PostEmptyEvent()
		}
		wg.Done()
	}()
	for mainloopActive {

		haveActiveScreen := true
		if !a.screen.Visible() {
			haveActiveScreen = false
		} else if a.screen.GLFWWindow().ShouldClose() {
			haveActiveScreen = false
			a.screen.SetVisible(false)

		} else {
			haveActiveScreen = true

			a.screen.DrawAll()

		}
		if !haveActiveScreen {
			mainloopActive = false
			break
		}

		glfw.WaitEvents()
	}

	wg.Wait()

}

func main() {
	nanogui.Init()
	//nanogui.SetDebug(true)
	app := Application{}
	app.init()
	//app.screen.DrawAll()
	//app.screen.SetVisible(true)
	nanogui.MainLoop()
}
