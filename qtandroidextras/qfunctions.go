package qtandroidextras

import "unsafe"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

func init() {
	if false {
		_ = unsafe.Pointer(uintptr(0))
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  header block end

//  body block begin
// /usr/include/qt/QtAndroidExtras/qandroidfunctions.h:107
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool shouldShowRequestPermissionRationale(const QString &)

/*

 */
func ShouldShowRequestPermissionRationale(permission string) bool {
	var tmpArg0 = qtcore.NewQString_5(permission)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QtAndroid36shouldShowRequestPermissionRationaleERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtAndroidExtras/qandroidfunctions.h:104
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void requestPermissions(const QStringList &, const QtAndroid::PermissionResultCallback &)

/*

 */
func RequestPermissions(permissions qtcore.QStringList_ITF, callbackFunc unsafe.Pointer /*555*/) {
	var convArg0 unsafe.Pointer
	if permissions != nil && permissions.QStringList_PTR() != nil {
		convArg0 = permissions.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QtAndroid18requestPermissionsERK11QStringListRKSt8functionIFvRK5QHashI7QStringNS_16PermissionResultEEEE", qtrt.FFI_TYPE_POINTER, convArg0, callbackFunc)
	qtrt.ErrPrint(err, rv)
}

//  body block end
