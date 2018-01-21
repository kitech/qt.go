package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QFileOpenEvent struct {
	*qtcore.QEvent
}

func (this *QFileOpenEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func NewQFileOpenEventFromPointer(cthis unsafe.Pointer) *QFileOpenEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QFileOpenEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:738
// index:0
// Public
// void QFileOpenEvent(const class QString &)
func NewQFileOpenEvent(file *qtcore.QString) *QFileOpenEvent {
	cthis := qtrt.Calloc(1, 256) // 40
	var convArg0 = file.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QFileOpenEventC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileOpenEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:739
// index:1
// Public
// void QFileOpenEvent(const class QUrl &)
func NewQFileOpenEvent_1(url *qtcore.QUrl) *QFileOpenEvent {
	cthis := qtrt.Calloc(1, 256) // 40
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QFileOpenEventC2ERK4QUrl", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileOpenEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:740
// index:0
// Public virtual
// void ~QFileOpenEvent()
func DeleteQFileOpenEvent(*QFileOpenEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QFileOpenEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:742
// index:0
// Public inline
// QString file()
func (this *QFileOpenEvent) File() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QFileOpenEvent4fileEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:743
// index:0
// Public inline
// QUrl url()
func (this *QFileOpenEvent) Url() *qtcore.QUrl /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QFileOpenEvent3urlEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
