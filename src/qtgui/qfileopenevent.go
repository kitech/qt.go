//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:738
// index:0
// void QFileOpenEvent(const class QString &)
func NewQFileOpenEvent(file unsafe.Pointer) *QFileOpenEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QFileOpenEventC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, file)
	gopp.ErrPrint(err, rv)
	return &QFileOpenEvent{cthis}
}

// /usr/include/qt/QtGui/qevent.h:739
// index:1
// void QFileOpenEvent(const class QUrl &)
func NewQFileOpenEvent_1(url unsafe.Pointer) *QFileOpenEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QFileOpenEventC2ERK4QUrl", ffiqt.FFI_TYPE_VOID, cthis, url)
	gopp.ErrPrint(err, rv)
	return &QFileOpenEvent{cthis}
}

// /usr/include/qt/QtGui/qevent.h:740
// index:0
// virtual
// void ~QFileOpenEvent()
func DeleteQFileOpenEvent(*QFileOpenEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QFileOpenEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:742
// index:0
// inline
// QString file()
func (this *QFileOpenEvent) File() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QFileOpenEvent4fileEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:743
// index:0
// inline
// QUrl url()
func (this *QFileOpenEvent) Url() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QFileOpenEvent3urlEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
