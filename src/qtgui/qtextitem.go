//  header block begin
// /usr/include/qt/QtGui/qpaintengine.h
// #include <qpaintengine.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 197
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
type QTextItem struct {
	*qtrt.CObject
}

func (this *QTextItem) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQTextItemFromPointer(cthis unsafe.Pointer) *QTextItem {
	return &QTextItem{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qpaintengine.h:75
// index:0
// Public
// qreal descent()
func (this *QTextItem) Descent() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextItem7descentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpaintengine.h:76
// index:0
// Public
// qreal ascent()
func (this *QTextItem) Ascent() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextItem6ascentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpaintengine.h:77
// index:0
// Public
// qreal width()
func (this *QTextItem) Width() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextItem5widthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpaintengine.h:79
// index:0
// Public
// QTextItem::RenderFlags renderFlags()
func (this *QTextItem) RenderFlags() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextItem11renderFlagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpaintengine.h:80
// index:0
// Public
// QString text()
func (this *QTextItem) Text() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextItem4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpaintengine.h:81
// index:0
// Public
// QFont font()
func (this *QTextItem) Font() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextItem4fontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
