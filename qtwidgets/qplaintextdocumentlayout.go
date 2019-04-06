// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qplaintextedit.h
// #include <qplaintextedit.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 114
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// void documentChanged(int, int, int)
func (this *QPlainTextDocumentLayout) InheritDocumentChanged(f func(from int, arg1 int, charsAdded int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "documentChanged", f)
}

/*

 */
type QPlainTextDocumentLayout struct {
	*qtgui.QAbstractTextDocumentLayout
}
type QPlainTextDocumentLayout_ITF interface {
	qtgui.QAbstractTextDocumentLayout_ITF
	QPlainTextDocumentLayout_PTR() *QPlainTextDocumentLayout
}

func (ptr *QPlainTextDocumentLayout) QPlainTextDocumentLayout_PTR() *QPlainTextDocumentLayout {
	return ptr
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

// /usr/include/qt/QtWidgets/qplaintextedit.h:298
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QPlainTextDocumentLayout) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:303
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPlainTextDocumentLayout(QTextDocument *)

/*

 */
func (*QPlainTextDocumentLayout) NewForInherit(document qtgui.QTextDocument_ITF /*777 QTextDocument **/) *QPlainTextDocumentLayout {
	return NewQPlainTextDocumentLayout(document)
}
func NewQPlainTextDocumentLayout(document qtgui.QTextDocument_ITF /*777 QTextDocument **/) *QPlainTextDocumentLayout {
	var convArg0 unsafe.Pointer
	if document != nil && document.QTextDocument_PTR() != nil {
		convArg0 = document.QTextDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QPlainTextDocumentLayoutC2EP13QTextDocument", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPlainTextDocumentLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPlainTextDocumentLayout")
	return gothis
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:304
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPlainTextDocumentLayout()

/*

 */
func DeleteQPlainTextDocumentLayout(this *QPlainTextDocumentLayout) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QPlainTextDocumentLayoutD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:307
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int hitTest(const QPointF &, Qt::HitTestAccuracy) const

/*

 */
func (this *QPlainTextDocumentLayout) HitTest(arg0 qtcore.QPointF_ITF, arg1 int) int {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPointF_PTR() != nil {
		convArg0 = arg0.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout7hitTestERK7QPointFN2Qt15HitTestAccuracyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, arg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:309
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int pageCount() const

/*

 */
func (this *QPlainTextDocumentLayout) PageCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout9pageCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:310
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QSizeF documentSize() const

/*

 */
func (this *QPlainTextDocumentLayout) DocumentSize() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout12documentSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:312
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF frameBoundingRect(QTextFrame *) const

/*

 */
func (this *QPlainTextDocumentLayout) FrameBoundingRect(arg0 qtgui.QTextFrame_ITF /*777 QTextFrame **/) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTextFrame_PTR() != nil {
		convArg0 = arg0.QTextFrame_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout17frameBoundingRectEP10QTextFrame", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:313
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF blockBoundingRect(const QTextBlock &) const

/*
Returns the bounding rectangle of the text block in the block's own coordinates.

See also blockBoundingGeometry().
*/
func (this *QPlainTextDocumentLayout) BlockBoundingRect(block qtgui.QTextBlock_ITF) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if block != nil && block.QTextBlock_PTR() != nil {
		convArg0 = block.QTextBlock_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout17blockBoundingRectERK10QTextBlock", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:315
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensureBlockLayout(const QTextBlock &) const

/*

 */
func (this *QPlainTextDocumentLayout) EnsureBlockLayout(block qtgui.QTextBlock_ITF) {
	var convArg0 unsafe.Pointer
	if block != nil && block.QTextBlock_PTR() != nil {
		convArg0 = block.QTextBlock_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout17ensureBlockLayoutERK10QTextBlock", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:317
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCursorWidth(int)

/*

 */
func (this *QPlainTextDocumentLayout) SetCursorWidth(width int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QPlainTextDocumentLayout14setCursorWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:318
// index:0
// Public Visibility=Default Availability=Available
// [4] int cursorWidth() const

/*

 */
func (this *QPlainTextDocumentLayout) CursorWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout11cursorWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:320
// index:0
// Public Visibility=Default Availability=Available
// [-2] void requestUpdate()

/*

 */
func (this *QPlainTextDocumentLayout) RequestUpdate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QPlainTextDocumentLayout13requestUpdateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:323
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void documentChanged(int, int, int)

/*

 */
func (this *QPlainTextDocumentLayout) DocumentChanged(from int, arg1 int, charsAdded int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QPlainTextDocumentLayout15documentChangedEiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), from, arg1, charsAdded)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11291() {
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
