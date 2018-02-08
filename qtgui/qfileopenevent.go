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
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
func (this *QFileOpenEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQFileOpenEventFromPointer(cthis unsafe.Pointer) *QFileOpenEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QFileOpenEvent{bcthis0}
}
func (*QFileOpenEvent) NewFromPointer(cthis unsafe.Pointer) *QFileOpenEvent {
	return NewQFileOpenEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:738
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFileOpenEvent(const QString &)
func NewQFileOpenEvent(file string) *QFileOpenEvent {
	var tmpArg0 = qtcore.NewQString_5(file)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QFileOpenEventC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileOpenEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFileOpenEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:739
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFileOpenEvent(const QUrl &)
func NewQFileOpenEvent_1(url *qtcore.QUrl) *QFileOpenEvent {
	var convArg0 = url.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QFileOpenEventC2ERK4QUrl", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileOpenEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFileOpenEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:740
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFileOpenEvent()
func DeleteQFileOpenEvent(this *QFileOpenEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QFileOpenEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:742
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString file()
func (this *QFileOpenEvent) File() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QFileOpenEvent4fileEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qevent.h:743
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QUrl url()
func (this *QFileOpenEvent) Url() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QFileOpenEvent3urlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:744
// index:0
// Public Visibility=Default Availability=Available
// [1] bool openFile(QFile &, QIODevice::OpenMode)
func (this *QFileOpenEvent) OpenFile(file *qtcore.QFile, flags int) bool {
	var convArg0 = file.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QFileOpenEvent8openFileER5QFile6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

//  body block end
