//  header block begin
// /usr/include/qt/QtGui/qtextobject.h
// #include <qtextobject.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
type QTextFrame struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qtextobject.h:120
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QTextFrame) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextFrame10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:123
// index:0
// void QTextFrame(class QTextDocument *)
func NewQTextFrame(doc unsafe.Pointer) *QTextFrame {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextFrameC2EP13QTextDocument", ffiqt.FFI_TYPE_VOID, cthis, doc)
	gopp.ErrPrint(err, rv)
	return &QTextFrame{cthis}
}

// /usr/include/qt/QtGui/qtextobject.h:124
// index:0
// virtual
// void ~QTextFrame()
func DeleteQTextFrame(*QTextFrame) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextFrameD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:126
// index:0
// inline
// void setFrameFormat(const class QTextFrameFormat &)
func (this *QTextFrame) SetFrameFormat(format unsafe.Pointer) {
	// 0: (, const QTextFrameFormat & format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextFrame14setFrameFormatERK16QTextFrameFormat", ffiqt.FFI_TYPE_VOID, this.cthis, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:127
// index:0
// inline
// QTextFrameFormat frameFormat()
func (this *QTextFrame) FrameFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextFrame11frameFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:129
// index:0
// QTextCursor firstCursorPosition()
func (this *QTextFrame) FirstCursorPosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextFrame19firstCursorPositionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:130
// index:0
// QTextCursor lastCursorPosition()
func (this *QTextFrame) LastCursorPosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextFrame18lastCursorPositionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:131
// index:0
// int firstPosition()
func (this *QTextFrame) FirstPosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextFrame13firstPositionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:132
// index:0
// int lastPosition()
func (this *QTextFrame) LastPosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextFrame12lastPositionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:134
// index:0
// QTextFrameLayoutData * layoutData()
func (this *QTextFrame) LayoutData() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextFrame10layoutDataEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:135
// index:0
// void setLayoutData(class QTextFrameLayoutData *)
func (this *QTextFrame) SetLayoutData(data unsafe.Pointer) {
	// 0: (, QTextFrameLayoutData * data), (data)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextFrame13setLayoutDataEP20QTextFrameLayoutData", ffiqt.FFI_TYPE_VOID, this.cthis, data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:137
// index:0
// QList<QTextFrame *> childFrames()
func (this *QTextFrame) ChildFrames() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextFrame11childFramesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:138
// index:0
// QTextFrame * parentFrame()
func (this *QTextFrame) ParentFrame() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextFrame11parentFrameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:181
// index:0
// QTextFrame::iterator begin()
func (this *QTextFrame) Begin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextFrame5beginEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:182
// index:0
// QTextFrame::iterator end()
func (this *QTextFrame) End() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextFrame3endEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
