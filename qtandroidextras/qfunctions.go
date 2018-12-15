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
// /usr/include/qt/QtAndroidExtras/qandroidjniobject.h:194
// index:0
// Invalid inline Visibility=Default Availability=Available
// [1] bool operator!=(const QAndroidJniObject &, const QAndroidJniObject &)

/*

 */
func Operator_not_equal(obj1 QAndroidJniObject_ITF, obj2 QAndroidJniObject_ITF) bool {
	var convArg0 unsafe.Pointer
	if obj1 != nil && obj1.QAndroidJniObject_PTR() != nil {
		convArg0 = obj1.QAndroidJniObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if obj2 != nil && obj2.QAndroidJniObject_PTR() != nil {
		convArg1 = obj2.QAndroidJniObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZneRK17QAndroidJniObjectS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtAndroidExtras/qandroidjniobject.h:189
// index:0
// Invalid inline Visibility=Default Availability=Available
// [1] bool operator==(const QAndroidJniObject &, const QAndroidJniObject &)

/*

 */
func Operator_equal_equal(obj1 QAndroidJniObject_ITF, obj2 QAndroidJniObject_ITF) bool {
	var convArg0 unsafe.Pointer
	if obj1 != nil && obj1.QAndroidJniObject_PTR() != nil {
		convArg0 = obj1.QAndroidJniObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if obj2 != nil && obj2.QAndroidJniObject_PTR() != nil {
		convArg1 = obj2.QAndroidJniObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZeqRK17QAndroidJniObjectS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtAndroidExtras/qandroidfunctions.h:71
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void runOnAndroidThreadSync(const QtAndroid::Runnable &, int)

/*

 */
func RunOnAndroidThreadSync(runnable unsafe.Pointer /*555*/, timeoutMs int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QtAndroid22runOnAndroidThreadSyncERKSt8functionIFvvEEi", qtrt.FFI_TYPE_POINTER, runnable, timeoutMs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtAndroidExtras/qandroidfunctions.h:70
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void runOnAndroidThread(const QtAndroid::Runnable &)

/*

 */
func RunOnAndroidThread(runnable unsafe.Pointer /*555*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QtAndroid18runOnAndroidThreadERKSt8functionIFvvEE", qtrt.FFI_TYPE_POINTER, runnable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtAndroidExtras/qandroidfunctions.h:62
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void startIntentSender(const QAndroidJniObject &, int, QAndroidActivityResultReceiver *)

/*

 */
func StartIntentSender(intentSender QAndroidJniObject_ITF, receiverRequestCode int, resultReceiver QAndroidActivityResultReceiver_ITF /*777 QAndroidActivityResultReceiver **/) {
	var convArg0 unsafe.Pointer
	if intentSender != nil && intentSender.QAndroidJniObject_PTR() != nil {
		convArg0 = intentSender.QAndroidJniObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if resultReceiver != nil && resultReceiver.QAndroidActivityResultReceiver_PTR() != nil {
		convArg2 = resultReceiver.QAndroidActivityResultReceiver_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QtAndroid17startIntentSenderERK17QAndroidJniObjectiP30QAndroidActivityResultReceiver", qtrt.FFI_TYPE_POINTER, convArg0, receiverRequestCode, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtAndroidExtras/qandroidfunctions.h:60
// index:0
// Invalid Visibility=Default Availability=Available
// [4] int androidSdkVersion()

/*

 */
func AndroidSdkVersion() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QtAndroid17androidSdkVersionEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtAndroidExtras/qandroidfunctions.h:73
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void hideSplashScreen()

/*

 */
func HideSplashScreen() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QtAndroid16hideSplashScreenEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtAndroidExtras/qandroidfunctions.h:57
// index:0
// Invalid Visibility=Default Availability=Available
// [16] QAndroidJniObject androidActivity()

/*

 */
func AndroidActivity() *QAndroidJniObject /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QtAndroid15androidActivityEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAndroidJniObject)
	return rv2
}

// /usr/include/qt/QtAndroidExtras/qandroidfunctions.h:58
// index:0
// Invalid Visibility=Default Availability=Available
// [16] QAndroidJniObject androidService()

/*

 */
func AndroidService() *QAndroidJniObject /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QtAndroid14androidServiceEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAndroidJniObject)
	return rv2
}

// /usr/include/qt/QtAndroidExtras/qandroidfunctions.h:59
// index:0
// Invalid Visibility=Default Availability=Available
// [16] QAndroidJniObject androidContext()

/*

 */
func AndroidContext() *QAndroidJniObject /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QtAndroid14androidContextEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAndroidJniObject)
	return rv2
}

// /usr/include/qt/QtAndroidExtras/qandroidfunctions.h:65
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void startActivity(const QAndroidJniObject &, int, QAndroidActivityResultReceiver *)

/*

 */
func StartActivity(intent QAndroidJniObject_ITF, receiverRequestCode int, resultReceiver QAndroidActivityResultReceiver_ITF /*777 QAndroidActivityResultReceiver **/) {
	var convArg0 unsafe.Pointer
	if intent != nil && intent.QAndroidJniObject_PTR() != nil {
		convArg0 = intent.QAndroidJniObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if resultReceiver != nil && resultReceiver.QAndroidActivityResultReceiver_PTR() != nil {
		convArg2 = resultReceiver.QAndroidActivityResultReceiver_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QtAndroid13startActivityERK17QAndroidJniObjectiP30QAndroidActivityResultReceiver", qtrt.FFI_TYPE_POINTER, convArg0, receiverRequestCode, convArg2)
	qtrt.ErrPrint(err, rv)
}

//  body block end
