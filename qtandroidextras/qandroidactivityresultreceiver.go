package qtandroidextras

// /usr/include/qt/QtAndroidExtras/qandroidactivityresultreceiver.h
// #include <qandroidactivityresultreceiver.h>
// #include <QtAndroidExtras>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QAndroidActivityResultReceiver struct {
	*qtrt.CObject
}
type QAndroidActivityResultReceiver_ITF interface {
	QAndroidActivityResultReceiver_PTR() *QAndroidActivityResultReceiver
}

func (ptr *QAndroidActivityResultReceiver) QAndroidActivityResultReceiver_PTR() *QAndroidActivityResultReceiver {
	return ptr
}

func (this *QAndroidActivityResultReceiver) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAndroidActivityResultReceiver) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAndroidActivityResultReceiverFromPointer(cthis unsafe.Pointer) *QAndroidActivityResultReceiver {
	return &QAndroidActivityResultReceiver{&qtrt.CObject{cthis}}
}
func (*QAndroidActivityResultReceiver) NewFromPointer(cthis unsafe.Pointer) *QAndroidActivityResultReceiver {
	return NewQAndroidActivityResultReceiverFromPointer(cthis)
}

// /usr/include/qt/QtAndroidExtras/qandroidactivityresultreceiver.h:52
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAndroidActivityResultReceiver()

/*

 */
func (*QAndroidActivityResultReceiver) NewForInherit() *QAndroidActivityResultReceiver {
	return NewQAndroidActivityResultReceiver()
}
func NewQAndroidActivityResultReceiver() *QAndroidActivityResultReceiver {
	rv, err := qtrt.InvokeQtFunc6("_ZN30QAndroidActivityResultReceiverC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAndroidActivityResultReceiverFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidActivityResultReceiver)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/qandroidactivityresultreceiver.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAndroidActivityResultReceiver()

/*

 */
func DeleteQAndroidActivityResultReceiver(this *QAndroidActivityResultReceiver) {
	rv, err := qtrt.InvokeQtFunc6("_ZN30QAndroidActivityResultReceiverD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtAndroidExtras/qandroidactivityresultreceiver.h:54
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void handleActivityResult(int, int, const QAndroidJniObject &)

/*
Reimplement this function to get activity results after starting an activity using either QtAndroid::startActivity() or QtAndroid::startIntentSender(). The receiverRequestCode is the request code unique to this receiver which was originally passed to the startActivity() or startIntentSender() functions. The resultCode is the result returned by the activity, and data is either null or a Java object of the class android.content.Intent. Both the last to arguments are identical to the arguments passed to onActivityResult().
*/
func (this *QAndroidActivityResultReceiver) HandleActivityResult(receiverRequestCode int, resultCode int, data QAndroidJniObject_ITF) {
	var convArg2 unsafe.Pointer
	if data != nil && data.QAndroidJniObject_PTR() != nil {
		convArg2 = data.QAndroidJniObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN30QAndroidActivityResultReceiver20handleActivityResultEiiRK17QAndroidJniObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), receiverRequestCode, resultCode, convArg2)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init() {
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
}

//  keep block end
