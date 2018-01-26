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
import "mkuse/cffiqt"
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
// Public virtual
// const QMetaObject * metaObject()
func (this *QSyntaxHighlighter) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:65
// index:0
// Public
// void QSyntaxHighlighter(class QObject *)
func NewQSyntaxHighlighter(parent *qtcore.QObject /*777 QObject **/) *QSyntaxHighlighter {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighterC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSyntaxHighlighterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:66
// index:1
// Public
// void QSyntaxHighlighter(class QTextDocument *)
func NewQSyntaxHighlighter_1(parent *QTextDocument /*777 QTextDocument **/) *QSyntaxHighlighter {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighterC1EP13QTextDocument", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSyntaxHighlighterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:67
// index:0
// Public virtual
// void ~QSyntaxHighlighter()
func DeleteQSyntaxHighlighter(*QSyntaxHighlighter) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighterD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:69
// index:0
// Public
// void setDocument(class QTextDocument *)
func (this *QSyntaxHighlighter) SetDocument(doc *QTextDocument /*777 QTextDocument **/) {
	var convArg0 = doc.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter11setDocumentEP13QTextDocument", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:70
// index:0
// Public
// QTextDocument * document()
func (this *QSyntaxHighlighter) Document() *QTextDocument /*777 QTextDocument **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter8documentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:73
// index:0
// Public
// void rehighlight()
func (this *QSyntaxHighlighter) Rehighlight() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter11rehighlightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:74
// index:0
// Public
// void rehighlightBlock(const class QTextBlock &)
func (this *QSyntaxHighlighter) RehighlightBlock(block *QTextBlock) {
	var convArg0 = block.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter16rehighlightBlockERK10QTextBlock", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:77
// index:0
// Protected pure virtual
// void highlightBlock(const class QString &)
func (this *QSyntaxHighlighter) HighlightBlock(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter14highlightBlockERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:79
// index:0
// Protected
// void setFormat(int, int, const class QTextCharFormat &)
func (this *QSyntaxHighlighter) SetFormat(start int, count int, format *QTextCharFormat) {
	var convArg2 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter9setFormatEiiRK15QTextCharFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), start, count, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:80
// index:1
// Protected
// void setFormat(int, int, const class QColor &)
func (this *QSyntaxHighlighter) SetFormat_1(start int, count int, color *QColor) {
	var convArg2 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter9setFormatEiiRK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), start, count, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:81
// index:2
// Protected
// void setFormat(int, int, const class QFont &)
func (this *QSyntaxHighlighter) SetFormat_2(start int, count int, font *QFont) {
	var convArg2 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter9setFormatEiiRK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), start, count, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:82
// index:0
// Protected
// QTextCharFormat format(int)
func (this *QSyntaxHighlighter) Format(pos int) *QTextCharFormat /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter6formatEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:84
// index:0
// Protected
// int previousBlockState()
func (this *QSyntaxHighlighter) PreviousBlockState() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter18previousBlockStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:85
// index:0
// Protected
// int currentBlockState()
func (this *QSyntaxHighlighter) CurrentBlockState() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter17currentBlockStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:86
// index:0
// Protected
// void setCurrentBlockState(int)
func (this *QSyntaxHighlighter) SetCurrentBlockState(newState int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter20setCurrentBlockStateEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), newState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:88
// index:0
// Protected
// void setCurrentBlockUserData(class QTextBlockUserData *)
func (this *QSyntaxHighlighter) SetCurrentBlockUserData(data *QTextBlockUserData /*777 QTextBlockUserData **/) {
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter23setCurrentBlockUserDataEP18QTextBlockUserData", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:89
// index:0
// Protected
// QTextBlockUserData * currentBlockUserData()
func (this *QSyntaxHighlighter) CurrentBlockUserData() *QTextBlockUserData /*777 QTextBlockUserData **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter20currentBlockUserDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextBlockUserDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:91
// index:0
// Protected
// QTextBlock currentBlock()
func (this *QSyntaxHighlighter) CurrentBlock() *QTextBlock /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter12currentBlockEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
