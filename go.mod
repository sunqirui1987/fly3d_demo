module github.com/suiqirui1987/fly3d_demo

require (
	github.com/sirupsen/logrus v1.4.1
	github.com/suiqirui1987/fly3d v1.0.0
)

replace github.com/suiqirui1987/fly3d => ../fly3d

replace github.com/suiqirui1987/fly3d_demo => ./

go 1.12
