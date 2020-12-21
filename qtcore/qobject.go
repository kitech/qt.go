package qtcore

// /usr/include/qt/QtCore/qobject.h
// #include <qobject.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*
 */
// size 16
type QObject struct {
	*qtrt.CObject
}
type QObject_ITF interface {
	QObject_PTR() *QObject
}

func (ptr *QObject) QObject_PTR() *QObject { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QObjectFromptr(cthis Voidptr) *QObject {
	return &QObject{&qtrt.CObject{cthis}}
}
func (*QObject) Fromptr(cthis Voidptr) *QObject {
	return QObjectFromptr(cthis)
}

// /usr/include/qt/QtCore/qobject.h:150
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString objectName() const

/*
 */
func (this *QObject) ObjectName() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(799704896, "_ZNK7QObject10objectNameEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := /*==*/ QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qobject.h:151
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setObjectName(const QString &)

/*
 */
func (this *QObject) SetObjectName(name string) {
	var tmpArg0 = NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(2393916980, "_ZN7QObject13setObjectNameERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:153
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool isWidgetType() const

/*
 */
func (this *QObject) IsWidgetType() bool {
	rv, err := qtrt.Qtcc3(4140218935, "_ZNK7QObject12isWidgetTypeEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:154
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool isWindowType() const

/*
 */
func (this *QObject) IsWindowType() bool {
	rv, err := qtrt.Qtcc3(2017529966, "_ZNK7QObject12isWindowTypeEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:156
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool signalsBlocked() const

/*
 */
func (this *QObject) SignalsBlocked() bool {
	rv, err := qtrt.Qtcc3(2511842610, "_ZNK7QObject14signalsBlockedEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:157
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool blockSignals(bool)

/*
 */
func (this *QObject) BlockSignals(b bool) bool {
	rv, err := qtrt.Qtcc3(2150786866, "_ZN7QObject12blockSignalsEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&b))
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:159
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QThread * thread() const

/*
 */
func (this *QObject) Thread() *QThread /*777 QThread **/ {
	rv, err := qtrt.Qtcc3(2989794423, "_ZNK7QObject6threadEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QThreadFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobject.h:160
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void moveToThread(QThread *)

/*
 */
func (this *QObject) MoveToThread(thread QThread_ITF /*777 QThread **/) {
	var convArg0 Voidptr
	if thread != nil && thread.QThread_PTR() != nil {
		convArg0 = thread.QThread_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(2881015986, "_ZN7QObject12moveToThreadEP7QThread", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:218
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setParent(QObject *)

/*
 */
func (this *QObject) SetParent(parent QObject_ITF /*777 QObject **/) {
	var convArg0 Voidptr
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(600858033, "_ZN7QObject9setParentEPS_", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

func DeleteQObject(this *QObject) {
	rv, err := qtrt.Qtcc1(0, "_ZN7QObjectD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10011() {
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
}

//  keep block end
