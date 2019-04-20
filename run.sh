#! /bin/sh


js(){
     echo "js"

    gopherjs  build
}

android(){
    echo "android"

    gomobile build -target=android -bundleid demo
}

ios(){
    echo "IOS" 

    gomobile build -target=ios -bundleid demo
}

os(){
    OS=`uname -s`
    echo "config go env"
    export GOPATH=$GOPATH:$currPWD
    export GO111MODULE=on
    export GOPROXY=https://goproxy.io


    echo "output demo os "
    gox -os="linux" -arch="amd64" -output="./bin/demo_linux64" ./
    gox -os="darwin" -arch="amd64" -output="./bin/demo_darwin64" ./
    gox -os="windows" -arch="amd64" -output="./bin/demo_windows" ./



}

case "$1" in
  js)
	js
        ;;
  android)
	android
        ;;
  ios)
	ios
        ;;
  os)
	os
        ;;
  *)
	FULLPATH=/etc/init.d/$PIDNAME
	echo "Usage: $FULLPATH {start|stop|restart|force-reload|status|configtest|terminate}"
	exit 1
	;;
esac