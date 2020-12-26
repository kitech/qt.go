package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

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

//  ext block end

//  body block begin

/*
 */
// size 64
type QKeyEvent struct {
	*QInputEvent
}
type QKeyEvent_ITF interface {
	QInputEvent_ITF
	QKeyEvent_PTR() *QKeyEvent
}

func (ptr *QKeyEvent) QKeyEvent_PTR() *QKeyEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QKeyEventFromptr(cthis Voidptr) *QKeyEvent {
	bcthis0 := QInputEventFromptr(cthis)
	return &QKeyEvent{bcthis0}
}
func (*QKeyEvent) Fromptr(cthis Voidptr) *QKeyEvent {
	return QKeyEventFromptr(cthis)
}

// /usr/include/qt/QtGui/qevent.h:387
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int key() const

/*
 */
func (this *QKeyEvent) Key() int {
	rv, err := qtrt.Qtcc3(3921584077, "_ZNK9QKeyEvent3keyEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:392
// index:0
// Public inline Indirect Visibility=Default Availability=Available
// [8] QString text() const

/*
 */
func (this *QKeyEvent) Text() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(3489740836, "_ZNK9QKeyEvent4textEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qevent.h:393
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool isAutoRepeat() const

/*
 */
func (this *QKeyEvent) IsAutoRepeat() bool {
	rv, err := qtrt.Qtcc3(3618180515, "_ZNK9QKeyEvent12isAutoRepeatEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qevent.h:394
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int count() const

/*
 */
func (this *QKeyEvent) Count() int {
	rv, err := qtrt.Qtcc3(2383005192, "_ZNK9QKeyEvent5countEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:396
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] quint32 nativeScanCode() const

/*
 */
func (this *QKeyEvent) NativeScanCode() uint {
	rv, err := qtrt.Qtcc3(1110117359, "_ZNK9QKeyEvent14nativeScanCodeEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qevent.h:397
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] quint32 nativeVirtualKey() const

/*
 */
func (this *QKeyEvent) NativeVirtualKey() uint {
	rv, err := qtrt.Qtcc3(3927640315, "_ZNK9QKeyEvent16nativeVirtualKeyEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qevent.h:398
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] quint32 nativeModifiers() const

/*
 */
func (this *QKeyEvent) NativeModifiers() uint {
	rv, err := qtrt.Qtcc3(2110805296, "_ZNK9QKeyEvent15nativeModifiersEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return uint(rv) // 222
}

func DeleteQKeyEvent(this *QKeyEvent) {
	rv, err := qtrt.Qtcc3(4215536018, "_ZN9QKeyEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10073() {
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
