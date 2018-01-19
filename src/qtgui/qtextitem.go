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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qpaintengine.h:75
// index:0
// qreal descent()
func (this *QTextItem) Descent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextItem7descentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:76
// index:0
// qreal ascent()
func (this *QTextItem) Ascent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextItem6ascentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:77
// index:0
// qreal width()
func (this *QTextItem) Width() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextItem5widthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:79
// index:0
// QTextItem::RenderFlags renderFlags()
func (this *QTextItem) RenderFlags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextItem11renderFlagsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:80
// index:0
// QString text()
func (this *QTextItem) Text() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextItem4textEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:81
// index:0
// QFont font()
func (this *QTextItem) Font() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextItem4fontEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
