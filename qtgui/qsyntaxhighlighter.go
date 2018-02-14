package qtgui

// /usr/include/qt/QtGui/qsyntaxhighlighter.h
// #include <qsyntaxhighlighter.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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

// void highlightBlock(const class QString &)
func (this *QSyntaxHighlighter) InheritHighlightBlock(f func(text string) /*void*/) {
	qtrt.SetAllInheritCallback(this, "highlightBlock", f)
}

// void setFormat(int, int, const class QTextCharFormat &)
func (this *QSyntaxHighlighter) InheritSetFormat(f func(start int, count int, format *QTextCharFormat) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setFormat", f)
}

// QTextCharFormat format(int)
func (this *QSyntaxHighlighter) InheritFormat(f func(pos int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "format", f)
}

// int previousBlockState()
func (this *QSyntaxHighlighter) InheritPreviousBlockState(f func() int) {
	qtrt.SetAllInheritCallback(this, "previousBlockState", f)
}

// int currentBlockState()
func (this *QSyntaxHighlighter) InheritCurrentBlockState(f func() int) {
	qtrt.SetAllInheritCallback(this, "currentBlockState", f)
}

// void setCurrentBlockState(int)
func (this *QSyntaxHighlighter) InheritSetCurrentBlockState(f func(newState int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setCurrentBlockState", f)
}

// void setCurrentBlockUserData(class QTextBlockUserData *)
func (this *QSyntaxHighlighter) InheritSetCurrentBlockUserData(f func(data *QTextBlockUserData /*777 QTextBlockUserData **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setCurrentBlockUserData", f)
}

// QTextBlockUserData * currentBlockUserData()
func (this *QSyntaxHighlighter) InheritCurrentBlockUserData(f func() unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "currentBlockUserData", f)
}

// QTextBlock currentBlock()
func (this *QSyntaxHighlighter) InheritCurrentBlock(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "currentBlock", f)
}

type QSyntaxHighlighter struct {
	*qtcore.QObject
}
type QSyntaxHighlighter_ITF interface {
	qtcore.QObject_ITF
	QSyntaxHighlighter_PTR() *QSyntaxHighlighter
}

func (ptr *QSyntaxHighlighter) QSyntaxHighlighter_PTR() *QSyntaxHighlighter { return ptr }

func (this *QSyntaxHighlighter) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QSyntaxHighlighter) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQSyntaxHighlighterFromPointer(cthis unsafe.Pointer) *QSyntaxHighlighter {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QSyntaxHighlighter{bcthis0}
}
func (*QSyntaxHighlighter) NewFromPointer(cthis unsafe.Pointer) *QSyntaxHighlighter {
	return NewQSyntaxHighlighterFromPointer(cthis)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QSyntaxHighlighter) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSyntaxHighlighter(QObject *)
func NewQSyntaxHighlighter(parent qtcore.QObject_ITF /*777 QObject **/) *QSyntaxHighlighter {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSyntaxHighlighterC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSyntaxHighlighterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSyntaxHighlighter")
	return gothis
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:66
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSyntaxHighlighter(QTextDocument *)
func NewQSyntaxHighlighter_1(parent QTextDocument_ITF /*777 QTextDocument **/) *QSyntaxHighlighter {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QTextDocument_PTR() != nil {
		convArg0 = parent.QTextDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSyntaxHighlighterC2EP13QTextDocument", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSyntaxHighlighterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSyntaxHighlighter")
	return gothis
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSyntaxHighlighter()
func DeleteQSyntaxHighlighter(this *QSyntaxHighlighter) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSyntaxHighlighterD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDocument(QTextDocument *)
func (this *QSyntaxHighlighter) SetDocument(doc QTextDocument_ITF /*777 QTextDocument **/) {
	var convArg0 unsafe.Pointer
	if doc != nil && doc.QTextDocument_PTR() != nil {
		convArg0 = doc.QTextDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSyntaxHighlighter11setDocumentEP13QTextDocument", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextDocument * document()
func (this *QSyntaxHighlighter) Document() *QTextDocument /*777 QTextDocument **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter8documentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rehighlight()
func (this *QSyntaxHighlighter) Rehighlight() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSyntaxHighlighter11rehighlightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rehighlightBlock(const QTextBlock &)
func (this *QSyntaxHighlighter) RehighlightBlock(block QTextBlock_ITF) {
	var convArg0 unsafe.Pointer
	if block != nil && block.QTextBlock_PTR() != nil {
		convArg0 = block.QTextBlock_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSyntaxHighlighter16rehighlightBlockERK10QTextBlock", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:77
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [-2] void highlightBlock(const QString &)
func (this *QSyntaxHighlighter) HighlightBlock(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSyntaxHighlighter14highlightBlockERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:79
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setFormat(int, int, const QTextCharFormat &)
func (this *QSyntaxHighlighter) SetFormat(start int, count int, format QTextCharFormat_ITF) {
	var convArg2 unsafe.Pointer
	if format != nil && format.QTextCharFormat_PTR() != nil {
		convArg2 = format.QTextCharFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSyntaxHighlighter9setFormatEiiRK15QTextCharFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), start, count, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:80
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void setFormat(int, int, const QColor &)
func (this *QSyntaxHighlighter) SetFormat_1(start int, count int, color QColor_ITF) {
	var convArg2 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg2 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSyntaxHighlighter9setFormatEiiRK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), start, count, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:81
// index:2
// Protected Visibility=Default Availability=Available
// [-2] void setFormat(int, int, const QFont &)
func (this *QSyntaxHighlighter) SetFormat_2(start int, count int, font QFont_ITF) {
	var convArg2 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg2 = font.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSyntaxHighlighter9setFormatEiiRK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), start, count, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:82
// index:0
// Protected Visibility=Default Availability=Available
// [16] QTextCharFormat format(int)
func (this *QSyntaxHighlighter) Format(pos int) *QTextCharFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter6formatEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCharFormat)
	return rv2
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:84
// index:0
// Protected Visibility=Default Availability=Available
// [4] int previousBlockState()
func (this *QSyntaxHighlighter) PreviousBlockState() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter18previousBlockStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:85
// index:0
// Protected Visibility=Default Availability=Available
// [4] int currentBlockState()
func (this *QSyntaxHighlighter) CurrentBlockState() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter17currentBlockStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:86
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setCurrentBlockState(int)
func (this *QSyntaxHighlighter) SetCurrentBlockState(newState int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSyntaxHighlighter20setCurrentBlockStateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newState)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:88
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setCurrentBlockUserData(QTextBlockUserData *)
func (this *QSyntaxHighlighter) SetCurrentBlockUserData(data QTextBlockUserData_ITF /*777 QTextBlockUserData **/) {
	var convArg0 unsafe.Pointer
	if data != nil && data.QTextBlockUserData_PTR() != nil {
		convArg0 = data.QTextBlockUserData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSyntaxHighlighter23setCurrentBlockUserDataEP18QTextBlockUserData", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:89
// index:0
// Protected Visibility=Default Availability=Available
// [8] QTextBlockUserData * currentBlockUserData()
func (this *QSyntaxHighlighter) CurrentBlockUserData() *QTextBlockUserData /*777 QTextBlockUserData **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter20currentBlockUserDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextBlockUserDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:91
// index:0
// Protected Visibility=Default Availability=Available
// [16] QTextBlock currentBlock()
func (this *QSyntaxHighlighter) CurrentBlock() *QTextBlock /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter12currentBlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

//  body block end

//  keep block begin

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
