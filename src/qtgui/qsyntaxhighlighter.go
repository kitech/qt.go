//  header block begin
// /usr/include/qt/QtGui/qsyntaxhighlighter.h
// #include <qsyntaxhighlighter.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QSyntaxHighlighter struct {
	*qtcore.QObject
}

func (this *QSyntaxHighlighter) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:62
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QSyntaxHighlighter) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:65
// index:0
// void QSyntaxHighlighter(class QObject *)
func NewQSyntaxHighlighter(parent unsafe.Pointer) *QSyntaxHighlighter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighterC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSyntaxHighlighterFromPointer(cthis)
	return gothis
}
func NewQSyntaxHighlighterFromPointer(cthis unsafe.Pointer) *QSyntaxHighlighter {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QSyntaxHighlighter{bcthis0}
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:66
// index:1
// void QSyntaxHighlighter(class QTextDocument *)
func NewQSyntaxHighlighter_1(parent unsafe.Pointer) *QSyntaxHighlighter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighterC1EP13QTextDocument", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSyntaxHighlighterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:67
// index:0
// virtual
// void ~QSyntaxHighlighter()
func DeleteQSyntaxHighlighter(*QSyntaxHighlighter) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighterD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:69
// index:0
// void setDocument(class QTextDocument *)
func (this *QSyntaxHighlighter) SetDocument(doc unsafe.Pointer) {
	// 0: (, doc QTextDocument *), (doc)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter11setDocumentEP13QTextDocument", ffiqt.FFI_TYPE_VOID, this.GetCthis(), doc)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:70
// index:0
// QTextDocument * document()
func (this *QSyntaxHighlighter) Document() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter8documentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:73
// index:0
// void rehighlight()
func (this *QSyntaxHighlighter) Rehighlight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter11rehighlightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:74
// index:0
// void rehighlightBlock(const class QTextBlock &)
func (this *QSyntaxHighlighter) RehighlightBlock(block unsafe.Pointer) {
	// 0: (, block const QTextBlock &), (block)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter16rehighlightBlockERK10QTextBlock", ffiqt.FFI_TYPE_VOID, this.GetCthis(), block)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:77
// index:0
// pure virtual
// void highlightBlock(const class QString &)
func (this *QSyntaxHighlighter) HighlightBlock(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter14highlightBlockERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:79
// index:0
// void setFormat(int, int, const class QTextCharFormat &)
func (this *QSyntaxHighlighter) SetFormat(start int, count int, format unsafe.Pointer) {
	// 0: (, start int, count int, format const QTextCharFormat &), (&start, &count, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter9setFormatEiiRK15QTextCharFormat", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &start, &count, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:80
// index:1
// void setFormat(int, int, const class QColor &)
func (this *QSyntaxHighlighter) SetFormat_1(start int, count int, color unsafe.Pointer) {
	// 1: (, start int, count int, color const QColor &), (&start, &count, color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter9setFormatEiiRK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &start, &count, color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:81
// index:2
// void setFormat(int, int, const class QFont &)
func (this *QSyntaxHighlighter) SetFormat_2(start int, count int, font unsafe.Pointer) {
	// 2: (, start int, count int, font const QFont &), (&start, &count, font)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter9setFormatEiiRK5QFont", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &start, &count, font)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:82
// index:0
// QTextCharFormat format(int)
func (this *QSyntaxHighlighter) Format(pos int) {
	// 0: (, pos int), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter6formatEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:84
// index:0
// int previousBlockState()
func (this *QSyntaxHighlighter) PreviousBlockState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter18previousBlockStateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:85
// index:0
// int currentBlockState()
func (this *QSyntaxHighlighter) CurrentBlockState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter17currentBlockStateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:86
// index:0
// void setCurrentBlockState(int)
func (this *QSyntaxHighlighter) SetCurrentBlockState(newState int) {
	// 0: (, newState int), (&newState)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter20setCurrentBlockStateEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &newState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:88
// index:0
// void setCurrentBlockUserData(class QTextBlockUserData *)
func (this *QSyntaxHighlighter) SetCurrentBlockUserData(data unsafe.Pointer) {
	// 0: (, data QTextBlockUserData *), (data)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSyntaxHighlighter23setCurrentBlockUserDataEP18QTextBlockUserData", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:89
// index:0
// QTextBlockUserData * currentBlockUserData()
func (this *QSyntaxHighlighter) CurrentBlockUserData() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter20currentBlockUserDataEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsyntaxhighlighter.h:91
// index:0
// QTextBlock currentBlock()
func (this *QSyntaxHighlighter) CurrentBlock() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSyntaxHighlighter12currentBlockEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
