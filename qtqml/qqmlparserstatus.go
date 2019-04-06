package qtqml

// /usr/include/qt/QtQml/qqmlparserstatus.h
// #include <qqmlparserstatus.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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

/*

 */
type QQmlParserStatus struct {
	*qtrt.CObject
}
type QQmlParserStatus_ITF interface {
	QQmlParserStatus_PTR() *QQmlParserStatus
}

func (ptr *QQmlParserStatus) QQmlParserStatus_PTR() *QQmlParserStatus { return ptr }

func (this *QQmlParserStatus) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlParserStatus) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQQmlParserStatusFromPointer(cthis unsafe.Pointer) *QQmlParserStatus {
	return &QQmlParserStatus{&qtrt.CObject{cthis}}
}
func (*QQmlParserStatus) NewFromPointer(cthis unsafe.Pointer) *QQmlParserStatus {
	return NewQQmlParserStatusFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlparserstatus.h:52
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlParserStatus()

/*

 */
func (*QQmlParserStatus) NewForInherit() *QQmlParserStatus {
	return NewQQmlParserStatus()
}
func NewQQmlParserStatus() *QQmlParserStatus {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QQmlParserStatusC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlParserStatusFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQmlParserStatus)
	return gothis
}

// /usr/include/qt/QtQml/qqmlparserstatus.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQmlParserStatus()

/*

 */
func DeleteQQmlParserStatus(this *QQmlParserStatus) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QQmlParserStatusD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlparserstatus.h:55
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void classBegin()

/*
Invoked after class creation, but before any properties have been set.
*/
func (this *QQmlParserStatus) ClassBegin() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QQmlParserStatus10classBeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlparserstatus.h:56
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void componentComplete()

/*
Invoked after the root component that caused this instantiation has completed construction. At this point all static values and binding values have been assigned to the class.
*/
func (this *QQmlParserStatus) ComponentComplete() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QQmlParserStatus17componentCompleteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11499() {
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
