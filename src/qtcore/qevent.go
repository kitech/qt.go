package qtcore

// /usr/include/qt/QtCore/qcoreevent.h
// #include <qcoreevent.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QEvent struct {
	*qtrt.CObject
}

func (this *QEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQEventFromPointer(cthis unsafe.Pointer) *QEvent {
	return &QEvent{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qcoreevent.h:297
// index:0
// Public
// void QEvent(enum QEvent::Type)
func NewQEvent(type_ int) *QEvent {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QEventC2ENS_4TypeE", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qcoreevent.h:299
// index:0
// Public virtual
// void ~QEvent()
func DeleteQEvent(*QEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreevent.h:301
// index:0
// Public inline
// QEvent::Type type()
func (this *QEvent) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QEvent4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qcoreevent.h:302
// index:0
// Public inline
// bool spontaneous()
func (this *QEvent) Spontaneous() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QEvent11spontaneousEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qcoreevent.h:304
// index:0
// Public inline
// void setAccepted(_Bool)
func (this *QEvent) SetAccepted(accepted bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QEvent11setAcceptedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &accepted)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreevent.h:305
// index:0
// Public inline
// bool isAccepted()
func (this *QEvent) IsAccepted() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QEvent10isAcceptedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qcoreevent.h:307
// index:0
// Public inline
// void accept()
func (this *QEvent) Accept() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QEvent6acceptEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreevent.h:308
// index:0
// Public inline
// void ignore()
func (this *QEvent) Ignore() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QEvent6ignoreEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreevent.h:310
// index:0
// Public static
// int registerEventType(int)
func (this *QEvent) RegisterEventType(hint int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QEvent17registerEventTypeEi", ffiqt.FFI_TYPE_POINTER, hint)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QEvent_RegisterEventType(hint int) int {
	var nilthis *QEvent
	rv := nilthis.RegisterEventType(hint)
	return rv
}

//  body block end
