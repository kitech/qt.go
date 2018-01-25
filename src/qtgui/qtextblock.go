package qtgui

// /usr/include/qt/QtGui/qtextobject.h
// #include <qtextobject.h>
// #include <QtGui>

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
	*qtrt.CObject
}

func (this *QTextBlock) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextBlock) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQTextBlockFromPointer(cthis unsafe.Pointer) *QTextBlock {
	return &QTextBlock{&qtrt.CObject{cthis}}
}
func (*QTextBlock) NewFromPointer(cthis unsafe.Pointer) *QTextBlock {
	return NewQTextBlockFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextobject.h:206
// index:0
// Public inline
// void QTextBlock()
func NewQTextBlock() *QTextBlock {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextBlockC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextBlockFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextobject.h:210
// index:0
// Public
// bool isValid()
func (this *QTextBlock) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextobject.h:216
// index:0
// Public
// int position()
func (this *QTextBlock) Position() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock8positionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextobject.h:217
// index:0
// Public
// int length()
func (this *QTextBlock) Length() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock6lengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextobject.h:218
// index:0
// Public
// bool contains(int)
func (this *QTextBlock) Contains(position int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock8containsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), position)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextobject.h:220
// index:0
// Public
// QTextLayout * layout()
func (this *QTextBlock) Layout() *QTextLayout /*444 QTextLayout **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock6layoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:221
// index:0
// Public
// void clearLayout()
func (this *QTextBlock) ClearLayout() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextBlock11clearLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:222
// index:0
// Public
// QTextBlockFormat blockFormat()
func (this *QTextBlock) BlockFormat() *QTextBlockFormat /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock11blockFormatEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextBlockFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:223
// index:0
// Public
// int blockFormatIndex()
func (this *QTextBlock) BlockFormatIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock16blockFormatIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextobject.h:224
// index:0
// Public
// QTextCharFormat charFormat()
func (this *QTextBlock) CharFormat() *QTextCharFormat /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock10charFormatEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:225
// index:0
// Public
// int charFormatIndex()
func (this *QTextBlock) CharFormatIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock15charFormatIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextobject.h:227
// index:0
// Public
// Qt::LayoutDirection textDirection()
func (this *QTextBlock) TextDirection() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock13textDirectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtextobject.h:229
// index:0
// Public
// QString text()
func (this *QTextBlock) Text() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock4textEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:233
// index:0
// Public
// const QTextDocument * document()
func (this *QTextBlock) Document() *QTextDocument /*444 const QTextDocument **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock8documentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:235
// index:0
// Public
// QTextList * textList()
func (this *QTextBlock) TextList() *QTextList /*444 QTextList **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock8textListEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextListFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:237
// index:0
// Public
// QTextBlockUserData * userData()
func (this *QTextBlock) UserData() *QTextBlockUserData /*444 QTextBlockUserData **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock8userDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextBlockUserDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:238
// index:0
// Public
// void setUserData(class QTextBlockUserData *)
func (this *QTextBlock) SetUserData(data *QTextBlockUserData /*444 QTextBlockUserData **/) {
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextBlock11setUserDataEP18QTextBlockUserData", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:240
// index:0
// Public
// int userState()
func (this *QTextBlock) UserState() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock9userStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextobject.h:241
// index:0
// Public
// void setUserState(int)
func (this *QTextBlock) SetUserState(state int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextBlock12setUserStateEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:243
// index:0
// Public
// int revision()
func (this *QTextBlock) Revision() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock8revisionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextobject.h:244
// index:0
// Public
// void setRevision(int)
func (this *QTextBlock) SetRevision(rev int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextBlock11setRevisionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), rev)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:246
// index:0
// Public
// bool isVisible()
func (this *QTextBlock) IsVisible() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock9isVisibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextobject.h:247
// index:0
// Public
// void setVisible(_Bool)
func (this *QTextBlock) SetVisible(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextBlock10setVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:249
// index:0
// Public
// int blockNumber()
func (this *QTextBlock) BlockNumber() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock11blockNumberEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextobject.h:250
// index:0
// Public
// int firstLineNumber()
func (this *QTextBlock) FirstLineNumber() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock15firstLineNumberEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextobject.h:252
// index:0
// Public
// void setLineCount(int)
func (this *QTextBlock) SetLineCount(count int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextBlock12setLineCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:253
// index:0
// Public
// int lineCount()
func (this *QTextBlock) LineCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock9lineCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextobject.h:283
// index:0
// Public
// QTextBlock::iterator begin()
func (this *QTextBlock) Begin() unsafe.Pointer /*444*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock5beginEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qtextobject.h:284
// index:0
// Public
// QTextBlock::iterator end()
func (this *QTextBlock) End() unsafe.Pointer /*444*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock3endEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qtextobject.h:286
// index:0
// Public
// QTextBlock next()
func (this *QTextBlock) Next() *QTextBlock /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock4nextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:287
// index:0
// Public
// QTextBlock previous()
func (this *QTextBlock) Previous() *QTextBlock /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock8previousEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:290
// index:0
// Public inline
// int fragmentIndex()
func (this *QTextBlock) FragmentIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextBlock13fragmentIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

//  body block end
