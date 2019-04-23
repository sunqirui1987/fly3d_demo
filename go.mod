module github.com/suiqirui1987/fly3d_demo

require (
	github.com/goxjs/gl v0.0.0-20171128034433-dc8f4a9a3c9c // indirect
	github.com/suiqirui1987/fly3d v1.0.0
)

replace github.com/suiqirui1987/fly3d => ../fly3d

replace github.com/suiqirui1987/fly3d_demo => ./

go 1.12
