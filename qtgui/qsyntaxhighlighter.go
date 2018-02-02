package qtgui

// /usr/include/qt/QtGui/qsyntaxhighlighter.h
// #include <qsyntaxhighlighter.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
// void highlightBlock(const class QString &)
func (this *QSyntaxHighlighter) InheritHighlightBlock(f func(text *qtcore.QString)) {
	ffiqt.SetAllInheritCallback(this, "highlightBlock", f)
}

// void setFormat(int, int, const class QTextCharFormat &)
func (this *QSyntaxHighlighter) InheritSetFormat(f func(start int, count int, format *QTextCharFormat)) {
	ffiqt.SetAllInheritCallback(this, "setFormat", f)
}

// QTextCharFormat format(int)
func (this *QSyntaxHighlighter) InheritFormat(f func(pos int) unsafe.Pointer) {
	ffiqt.SetAllInheritCallback(this, "format", f)
}

// int previousBlockState()
func (this *QSyntaxHighlighter) InheritPreviousBlockState(f func() int) {
	ffiqt.SetAllInheritCallback(this, "previousBlockState", f)
}

// int currentBlockState()
func (this *QSyntaxHighlighter) InheritCurrentBlockState(f func() int) {
	ffiqt.SetAllInheritCallback(this, "currentBlockState", f)
}

// void setCurrentBlockState(int)
func (this *QSyntaxHighlighter) InheritSetCurrentBlockState(f func(newState int)) {
	ffiqt.SetAllInheritCallback(this, "setCurrentBlockState", f)
}

// void setCurrentBlockUserData(class QTextBlockUserData *)
func (this *QSyntaxHighlighter) InheritSetCurrentBlockUserData(f func(data *QTextBlockUserData /*777 QTextBlockUserData **/)) {
	ffiqt.SetAllInheritCallback(this, "setCurrentBlockUserData", f)
}

// QTextBlockUserData * currentBlockUserData()
func (this *QSyntaxHighlighter) InheritCurrentBlockUserData(f func() unsafe.Pointer /*666*/) {
	ffiqt.SetAllInheritCallback(this, "currentBlockUserData", f)
}

// QTextBlock currentBlock()
func (this *QSyntaxHighlighter) InheritCurrentBlock(f func() unsafe.Pointer) {
	ffiqt.SetAllInheritCallback(this, "currentBlock", f)
}

type QSyntaxHighlighter struct {
	*qtcore.QObject
}

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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSyntaxHighlighter(QObject *)
func NewQSyntaxHighlighter(parent *qtcore.QObject /*777 QObject **/) *QSyntaxHighlighter {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighterC1EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSyntaxHighlighterFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:66
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSyntaxHighlighter(QTextDocument *)
func NewQSyntaxHighlighter_1(parent *QTextDocument /*777 QTextDocument **/) *QSyntaxHighlighter {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighterC1EP13QTextDocument", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSyntaxHighlighterFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSyntaxHighlighter()
func DeleteQSyntaxHighlighter(this *QSyntaxHighlighter) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighterD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDocument(QTextDocument *)
func (this *QSyntaxHighlighter) SetDocument(doc *QTextDocument /*777 QTextDocument **/) {
	var convArg0 = doc.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter11setDocumentEP13QTextDocument", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextDocument * document()
func (this *QSyntaxHighlighter) Document() *QTextDocument /*777 QTextDocument **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter8documentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rehighlight()
func (this *QSyntaxHighlighter) Rehighlight() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter11rehighlightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rehighlightBlock(const QTextBlock &)
func (this *QSyntaxHighlighter) RehighlightBlock(block *QTextBlock) {
	var convArg0 = block.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter16rehighlightBlockERK10QTextBlock", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:77
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [-2] void highlightBlock(const QString &)
func (this *QSyntaxHighlighter) HighlightBlock(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter14highlightBlockERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:79
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setFormat(int, int, const QTextCharFormat &)
func (this *QSyntaxHighlighter) SetFormat(start int, count int, format *QTextCharFormat) {
	var convArg2 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter9setFormatEiiRK15QTextCharFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), start, count, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:80
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void setFormat(int, int, const QColor &)
func (this *QSyntaxHighlighter) SetFormat_1(start int, count int, color *QColor) {
	var convArg2 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter9setFormatEiiRK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), start, count, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:81
// index:2
// Protected Visibility=Default Availability=Available
// [-2] void setFormat(int, int, const QFont &)
func (this *QSyntaxHighlighter) SetFormat_2(start int, count int, font *QFont) {
	var convArg2 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter9setFormatEiiRK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), start, count, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:82
// index:0
// Protected Visibility=Default Availability=Available
// [16] QTextCharFormat format(int)
func (this *QSyntaxHighlighter) Format(pos int) *QTextCharFormat /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter6formatEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCharFormat)
	return rv2
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:84
// index:0
// Protected Visibility=Default Availability=Available
// [4] int previousBlockState()
func (this *QSyntaxHighlighter) PreviousBlockState() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter18previousBlockStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:85
// index:0
// Protected Visibility=Default Availability=Available
// [4] int currentBlockState()
func (this *QSyntaxHighlighter) CurrentBlockState() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter17currentBlockStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:86
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setCurrentBlockState(int)
func (this *QSyntaxHighlighter) SetCurrentBlockState(newState int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter20setCurrentBlockStateEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), newState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:88
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setCurrentBlockUserData(QTextBlockUserData *)
func (this *QSyntaxHighlighter) SetCurrentBlockUserData(data *QTextBlockUserData /*777 QTextBlockUserData **/) {
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter23setCurrentBlockUserDataEP18QTextBlockUserData", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:89
// index:0
// Protected Visibility=Default Availability=Available
// [8] QTextBlockUserData * currentBlockUserData()
func (this *QSyntaxHighlighter) CurrentBlockUserData() *QTextBlockUserData /*777 QTextBlockUserData **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter20currentBlockUserDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextBlockUserDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:91
// index:0
// Protected Visibility=Default Availability=Available
// [16] QTextBlock currentBlock()
func (this *QSyntaxHighlighter) CurrentBlock() *QTextBlock /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter12currentBlockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

//  body block end
