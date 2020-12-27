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
// size 40
type QFileOpenEvent struct {
	*qtcore.QEvent
}
type QFileOpenEvent_ITF interface {
	qtcore.QEvent_ITF
	QFileOpenEvent_PTR() *QFileOpenEvent
}

func (ptr *QFileOpenEvent) QFileOpenEvent_PTR() *QFileOpenEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QFileOpenEventFromptr(cthis Voidptr) *QFileOpenEvent {
	bcthis0 := qtcore.QEventFromptr(cthis)
	return &QFileOpenEvent{bcthis0}
}
func (*QFileOpenEvent) Fromptr(cthis Voidptr) *QFileOpenEvent {
	return QFileOpenEventFromptr(cthis)
}

// /usr/include/qt/QtGui/qevent.h:784
// index:0
// Public inline Indirect Visibility=Default Availability=Available
// [8] QString file() const

/*
 */
func (this *QFileOpenEvent) File() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(1192897722, "_ZNK14QFileOpenEvent4fileEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint3(err, rv)
	rv.High = uint64(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(rv.Ptr())
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

func DeleteQFileOpenEvent(this *QFileOpenEvent) {
	rv, err := qtrt.Qtcc3(3021811242, "_ZN14QFileOpenEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10125() {
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
