package qtwidgets

// /usr/include/qt/QtWidgets/qplaintextedit.h
// #include <qplaintextedit.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 114
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
// void documentChanged(int, int, int)
func (this *QPlainTextDocumentLayout) InheritDocumentChanged(f func(from int, arg1 int, charsAdded int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "documentChanged", f)
}

type QPlainTextDocumentLayout struct {
	*qtgui.QAbstractTextDocumentLayout
}

func (this *QPlainTextDocumentLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractTextDocumentLayout.GetCthis()
	}
}
func (this *QPlainTextDocumentLayout) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractTextDocumentLayout = qtgui.NewQAbstractTextDocumentLayoutFromPointer(cthis)
}
func NewQPlainTextDocumentLayoutFromPointer(cthis unsafe.Pointer) *QPlainTextDocumentLayout {
	bcthis0 := qtgui.NewQAbstractTextDocumentLayoutFromPointer(cthis)
	return &QPlainTextDocumentLayout{bcthis0}
}
func (*QPlainTextDocumentLayout) NewFromPointer(cthis unsafe.Pointer) *QPlainTextDocumentLayout {
	return NewQPlainTextDocumentLayoutFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:297
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QPlainTextDocumentLayout) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:302
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPlainTextDocumentLayout(QTextDocument *)
func NewQPlainTextDocumentLayout(document *qtgui.QTextDocument /*777 QTextDocument **/) *QPlainTextDocumentLayout {
	var convArg0 = document.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN24QPlainTextDocumentLayoutC2EP13QTextDocument", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPlainTextDocumentLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:303
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPlainTextDocumentLayout()
func DeleteQPlainTextDocumentLayout(this *QPlainTextDocumentLayout) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QPlainTextDocumentLayoutD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:306
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int hitTest(const QPointF &, Qt::HitTestAccuracy)
func (this *QPlainTextDocumentLayout) HitTest(arg0 *qtcore.QPointF, arg1 int) int {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout7hitTestERK7QPointFN2Qt15HitTestAccuracyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, arg1)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:308
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int pageCount()
func (this *QPlainTextDocumentLayout) PageCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout9pageCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:309
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QSizeF documentSize()
func (this *QPlainTextDocumentLayout) DocumentSize() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout12documentSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:311
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF frameBoundingRect(QTextFrame *)
func (this *QPlainTextDocumentLayout) FrameBoundingRect(arg0 *qtgui.QTextFrame /*777 QTextFrame **/) *qtcore.QRectF /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout17frameBoundingRectEP10QTextFrame", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:312
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF blockBoundingRect(const QTextBlock &)
func (this *QPlainTextDocumentLayout) BlockBoundingRect(block *qtgui.QTextBlock) *qtcore.QRectF /*123*/ {
	var convArg0 = block.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout17blockBoundingRectERK10QTextBlock", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:314
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensureBlockLayout(const QTextBlock &)
func (this *QPlainTextDocumentLayout) EnsureBlockLayout(block *qtgui.QTextBlock) {
	var convArg0 = block.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout17ensureBlockLayoutERK10QTextBlock", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:316
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCursorWidth(int)
func (this *QPlainTextDocumentLayout) SetCursorWidth(width int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QPlainTextDocumentLayout14setCursorWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:317
// index:0
// Public Visibility=Default Availability=Available
// [4] int cursorWidth()
func (this *QPlainTextDocumentLayout) CursorWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout11cursorWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:319
// index:0
// Public Visibility=Default Availability=Available
// [-2] void requestUpdate()
func (this *QPlainTextDocumentLayout) RequestUpdate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QPlainTextDocumentLayout13requestUpdateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:322
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void documentChanged(int, int, int)
func (this *QPlainTextDocumentLayout) DocumentChanged(from int, arg1 int, charsAdded int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QPlainTextDocumentLayout15documentChangedEiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), from, arg1, charsAdded)
	gopp.ErrPrint(err, rv)
}

//  body block end
