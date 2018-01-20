//  header block begin
// /usr/include/qt/QtGui/qtextlist.h
// #include <qtextlist.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
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
type QTextList struct {
	*QTextBlockGroup
}

func (this *QTextList) GetCthis() unsafe.Pointer {
	return this.QTextBlockGroup.GetCthis()
}

// /usr/include/qt/QtGui/qtextlist.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QTextList) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextList10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:57
// index:0
// void QTextList(class QTextDocument *)
func NewQTextList(doc unsafe.Pointer) *QTextList {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextListC2EP13QTextDocument", ffiqt.FFI_TYPE_VOID, cthis, doc)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextListFromPointer(cthis)
	return gothis
}
func NewQTextListFromPointer(cthis unsafe.Pointer) *QTextList {
	bcthis0 := NewQTextBlockGroupFromPointer(cthis)
	return &QTextList{bcthis0}
}

// /usr/include/qt/QtGui/qtextlist.h:58
// index:0
// virtual
// void ~QTextList()
func DeleteQTextList(*QTextList) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextListD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:60
// index:0
// int count()
func (this *QTextList) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextList5countEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:62
// index:0
// inline
// bool isEmpty()
func (this *QTextList) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextList7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:65
// index:0
// QTextBlock item(int)
func (this *QTextList) Item(i int) {
	// 0: (, i int), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextList4itemEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:67
// index:0
// int itemNumber(const class QTextBlock &)
func (this *QTextList) ItemNumber(arg0 unsafe.Pointer) {
	// 0: (, const QTextBlock & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextList10itemNumberERK10QTextBlock", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:68
// index:0
// QString itemText(const class QTextBlock &)
func (this *QTextList) ItemText(arg0 unsafe.Pointer) {
	// 0: (, const QTextBlock & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextList8itemTextERK10QTextBlock", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:70
// index:0
// void removeItem(int)
func (this *QTextList) RemoveItem(i int) {
	// 0: (, i int), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextList10removeItemEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:71
// index:0
// void remove(const class QTextBlock &)
func (this *QTextList) Remove(arg0 unsafe.Pointer) {
	// 0: (, const QTextBlock & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextList6removeERK10QTextBlock", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:73
// index:0
// void add(const class QTextBlock &)
func (this *QTextList) Add(block unsafe.Pointer) {
	// 0: (, block const QTextBlock &), (block)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextList3addERK10QTextBlock", ffiqt.FFI_TYPE_VOID, this.GetCthis(), block)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:75
// index:0
// inline
// void setFormat(const class QTextListFormat &)
func (this *QTextList) SetFormat(format unsafe.Pointer) {
	// 0: (, format const QTextListFormat &), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextList9setFormatERK15QTextListFormat", ffiqt.FFI_TYPE_VOID, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:76
// index:0
// inline
// QTextListFormat format()
func (this *QTextList) Format() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextList6formatEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
