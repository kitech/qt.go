//  header block begin

// +build !minimal

package qtqml

// /usr/include/qt/QtQml/qqmlfile.h
// #include <qqmlfile.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

// /usr/include/qt/QtQml/qqmlfile.h:84
// index:0
// Public Visibility=Default Availability=Available
// [1] bool connectFinished(QObject *, const char *)

/*

 */
func (this *QQmlFile) ConnectFinished(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 string) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(arg1)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlFile15connectFinishedEP7QObjectPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlfile.h:85
// index:1
// Public Visibility=Default Availability=Available
// [1] bool connectFinished(QObject *, int)

/*

 */
func (this *QQmlFile) ConnectFinished1(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 int) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlFile15connectFinishedEP7QObjecti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, arg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlfile.h:86
// index:0
// Public Visibility=Default Availability=Available
// [1] bool connectDownloadProgress(QObject *, const char *)

/*

 */
func (this *QQmlFile) ConnectDownloadProgress(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 string) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(arg1)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlFile23connectDownloadProgressEP7QObjectPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlfile.h:87
// index:1
// Public Visibility=Default Availability=Available
// [1] bool connectDownloadProgress(QObject *, int)

/*

 */
func (this *QQmlFile) ConnectDownloadProgress1(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 int) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlFile23connectDownloadProgressEP7QObjecti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, arg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

//  body block end

//  keep block begin

func init_unused_11530() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
