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
type QTextBlock struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qtextobject.h:205
// index:0
// inline
// void QTextBlock(class QTextDocumentPrivate *, int)
func NewQTextBlock(priv unsafe.Pointer, b int) *QTextBlock {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextBlockC2EP20QTextDocumentPrivatei", ffiqt.FFI_TYPE_VOID, cthis, priv, &b)
	gopp.ErrPrint(err, rv)
	return &QTextBlock{cthis}
}

// /usr/include/qt/QtGui/qtextobject.h:206
// index:1
// inline
// void QTextBlock()
func NewQTextBlock_1() *QTextBlock {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextBlockC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QTextBlock{cthis}
}

// /usr/include/qt/QtGui/qtextobject.h:210
// index:0
// bool isValid()
func (this *QTextBlock) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:216
// index:0
// int position()
func (this *QTextBlock) Position() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock8positionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:217
// index:0
// int length()
func (this *QTextBlock) Length() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock6lengthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:218
// index:0
// bool contains(int)
func (this *QTextBlock) Contains(position int) {
	// 0: (, int position), (&position)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock8containsEi", ffiqt.FFI_TYPE_VOID, this.cthis, &position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:220
// index:0
// QTextLayout * layout()
func (this *QTextBlock) Layout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock6layoutEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:221
// index:0
// void clearLayout()
func (this *QTextBlock) ClearLayout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextBlock11clearLayoutEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:222
// index:0
// QTextBlockFormat blockFormat()
func (this *QTextBlock) BlockFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock11blockFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:223
// index:0
// int blockFormatIndex()
func (this *QTextBlock) BlockFormatIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock16blockFormatIndexEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:224
// index:0
// QTextCharFormat charFormat()
func (this *QTextBlock) CharFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock10charFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:225
// index:0
// int charFormatIndex()
func (this *QTextBlock) CharFormatIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock15charFormatIndexEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:227
// index:0
// Qt::LayoutDirection textDirection()
func (this *QTextBlock) TextDirection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock13textDirectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:229
// index:0
// QString text()
func (this *QTextBlock) Text() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock4textEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:231
// index:0
// QVector<QTextLayout::FormatRange> textFormats()
func (this *QTextBlock) TextFormats() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock11textFormatsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:233
// index:0
// const QTextDocument * document()
func (this *QTextBlock) Document() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock8documentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:235
// index:0
// QTextList * textList()
func (this *QTextBlock) TextList() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock8textListEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:237
// index:0
// QTextBlockUserData * userData()
func (this *QTextBlock) UserData() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock8userDataEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:238
// index:0
// void setUserData(class QTextBlockUserData *)
func (this *QTextBlock) SetUserData(data unsafe.Pointer) {
	// 0: (, QTextBlockUserData * data), (data)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextBlock11setUserDataEP18QTextBlockUserData", ffiqt.FFI_TYPE_VOID, this.cthis, data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:240
// index:0
// int userState()
func (this *QTextBlock) UserState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock9userStateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:241
// index:0
// void setUserState(int)
func (this *QTextBlock) SetUserState(state int) {
	// 0: (, int state), (&state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextBlock12setUserStateEi", ffiqt.FFI_TYPE_VOID, this.cthis, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:243
// index:0
// int revision()
func (this *QTextBlock) Revision() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock8revisionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:244
// index:0
// void setRevision(int)
func (this *QTextBlock) SetRevision(rev int) {
	// 0: (, int rev), (&rev)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextBlock11setRevisionEi", ffiqt.FFI_TYPE_VOID, this.cthis, &rev)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:246
// index:0
// bool isVisible()
func (this *QTextBlock) IsVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock9isVisibleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:247
// index:0
// void setVisible(_Bool)
func (this *QTextBlock) SetVisible(visible bool) {
	// 0: (, bool visible), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextBlock10setVisibleEb", ffiqt.FFI_TYPE_VOID, this.cthis, &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:249
// index:0
// int blockNumber()
func (this *QTextBlock) BlockNumber() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock11blockNumberEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:250
// index:0
// int firstLineNumber()
func (this *QTextBlock) FirstLineNumber() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock15firstLineNumberEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:252
// index:0
// void setLineCount(int)
func (this *QTextBlock) SetLineCount(count int) {
	// 0: (, int count), (&count)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextBlock12setLineCountEi", ffiqt.FFI_TYPE_VOID, this.cthis, &count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:253
// index:0
// int lineCount()
func (this *QTextBlock) LineCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock9lineCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:283
// index:0
// QTextBlock::iterator begin()
func (this *QTextBlock) Begin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock5beginEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:284
// index:0
// QTextBlock::iterator end()
func (this *QTextBlock) End() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock3endEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:286
// index:0
// QTextBlock next()
func (this *QTextBlock) Next() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock4nextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:287
// index:0
// QTextBlock previous()
func (this *QTextBlock) Previous() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock8previousEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:289
// index:0
// inline
// QTextDocumentPrivate * docHandle()
func (this *QTextBlock) DocHandle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock9docHandleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:290
// index:0
// inline
// int fragmentIndex()
func (this *QTextBlock) FragmentIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock13fragmentIndexEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
