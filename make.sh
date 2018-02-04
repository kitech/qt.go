source ~/triline/shell/android-ndk-env.sh


export CGO_ENABLED=1
export GOOS=android
export GOARCH=arm
export GOARM=7
export CGO_LDFLAGS="$CGO_LDFLAGS -L/androidsys/lib"

set -x
# rm -rf ~/oss/pkg/android_arm
# go install -p 1 -v -x  -pkgdir ~/oss/pkg/android_arm github.com/gonuts/ffi
go install -p 1 -v  -pkgdir ~/oss/pkg/android_arm ./qtrt ./qtcore ./qtgui ./qtwidgets
