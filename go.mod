module github.com/suiqirui1987/fly3d_demo

require (
	github.com/go-gl/gl v0.0.0-20190320180904-bf2b1f2f34d7 // indirect
	github.com/go-gl/glfw v0.0.0-20190409004039-e6da0acd62b1 // indirect
	github.com/suiqirui1987/fly3d v1.0.0
	golang.org/x/mobile v0.0.0-20190415191353-3e0bab5405d6 // indirect
)

replace github.com/suiqirui1987/fly3d => ../fly3d

replace github.com/suiqirui1987/fly3d_demo => ./

go 1.12
