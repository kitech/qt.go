package qtandroidextras

// this file is write by hand, don't delete
import (
	"unsafe"

	"github.com/kitech/qt.go/qtandroidrt"
	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
)

func AndroidActivity_() *QAndroidJniObject {
	rv, err := qtrt.InvokeQtFunc6("QtAndroid_androidActivity", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv)))
	// qtrt.SetFinalizer(gothis, DeleteQAndroidJniObject)
	return gothis
}

func AndroidContext_() *QAndroidJniObject {
	rv, err := qtrt.InvokeQtFunc6("QtAndroid_androidContext", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv)))
	// qtrt.SetFinalizer(gothis, DeleteQAndroidJniObject)
	return gothis
}

func AndroidSdkVersion_() int {
	rv, err := qtrt.InvokeQtFunc6("QtAndroid_androidAdkVersion", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

func AndroidService_() *QAndroidJniObject {
	rv, err := qtrt.InvokeQtFunc6("QtAndroid_androidService", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv)))
	// qtrt.SetFinalizer(gothis, DeleteQAndroidJniObject)
	return gothis
}

func BindService_(serviceIntent *QAndroidIntent, serviceConnection *QAndroidServiceConnection, flags int) bool {
	var convArg0 = serviceIntent.GetCthis()
	var convArg1 = serviceConnection.GetCthis()
	var convArg2 = flags
	rv, err := qtrt.InvokeQtFunc6("QtAndroid_androidService", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	// qtrt.Cretval2go("bool", rv)
	return qtrt.IfElse(rv == 0, false, true).(bool)
}

func CheckPermission_(permission string) int {
	var tmpArg0 = qtcore.NewQString5(permission)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("QtAndroid_checkPermission", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

func HideSplashScreen_() {
	rv, err := qtrt.InvokeQtFunc6("QtAndroid_hideSplashScreen", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}

func HideSplashScreen_2(duration int) {
	var convArg0 = duration
	rv, err := qtrt.InvokeQtFunc6("QtAndroid_hideSplashScreen_2", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}

// TODO QtAndroid::Runnable
/*
func RunOnAndroidThread(runnable unsafe.Pointer) {
	var convArg0 = runnable
	rv, err := qtrt.InvokeQtFunc6("QtAndroid_runOnAndroidThread", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
*/
// Note: can not callback to go scope, or it's will crash or blocked
func RunOnAndroidThread_(runnable func()) {
	if runnable == nil {
		return
	}
	cbno, cbfn := cmemory.AddOnceRunner(runnable)
	var convArg0 = cbfn
	rv, err := qtrt.InvokeQtFunc6("QtAndroid_runOnAndroidThread", qtrt.FFI_TYPE_POINTER, cbno, convArg0)
	qtrt.ErrPrint(err, rv)
}

func RunOnAndroidThread2(symname string, a0 uint64, a1 uint64) {
	symptr := qtrt.GetQtSymAddr(symname)
	rv, err := qtrt.InvokeQtFunc6("QtAndroid_runOnAndroidThread2", qtrt.FFI_TYPE_POINTER, symptr, a0, a1)
	qtrt.ErrPrint(err, rv)
}

// TODO QtAndroid::Runnable
/*
func RunOnAndroidThreadSync(runnable unsafe.Pointer, timeoutMS int) {
	var convArg0 = runnable
	var convArg1 = timeoutMS
	rv, err := qtrt.InvokeQtFunc6("QtAndroid_runOnAndroidThreadSync", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}
*/
func RunOnAndroidThreadSync_(runnable func(), timeoutMS int) {
	if runnable == nil {
		return
	}
	cbno, cbfn := cmemory.AddOnceRunner(runnable)
	var convArg0 = cbfn
	var convArg1 = timeoutMS
	rv, err := qtrt.InvokeQtFunc6("QtAndroid_runOnAndroidThreadSync", qtrt.FFI_TYPE_POINTER, cbno, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

func ShouldShowRequestPermissionRationale_handby(permission string) int {
	var tmpArg0 = qtcore.NewQString5(permission)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("QtAndroid_shouldShowRequestPermissionRationale", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

func StartActivity_(intent *QAndroidJniObject, receiverRequestCode int, resultReceiver *QAndroidActivityResultReceiver) {
	var convArg0 = intent.GetCthis()
	var convArg1 = receiverRequestCode
	var convArg2 = resultReceiver.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("QtAndroid_startActivity", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

func StartIntentSender_(intentSender *QAndroidJniObject, receiverRequestCode int, resultReceiver *QAndroidActivityResultReceiver) {
	var convArg0 = intentSender.GetCthis()
	var convArg1 = receiverRequestCode
	var convArg2 = resultReceiver.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("QtAndroid_startIntentSender", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

///////// heavy wrapper, not just binding
func KeepScreenOn(on bool) {
	ion := qtrt.IfElseInt(on, 1, 0)
	RunOnAndroidThread2("C_AndroidKeepScreenOnRaw", uint64(ion), 0)
}

func ShowToast(message string, duration uint64) {
	msgp := qtrt.CString(message) // memory free by callee
	RunOnAndroidThread2("C_AndroidShowToast", uint64(uintptr(msgp)), duration)
}
