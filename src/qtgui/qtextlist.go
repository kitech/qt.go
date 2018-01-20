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
func NewQTextListFromPointer(cthis unsafe.Pointer) *QTextList {
	bcthis0 := NewQTextBlockGroupFromPointer(cthis)
	return &QTextList{bcthis0}
}

// /usr/include/qt/QtGui/qtextlist.h:55
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QTextList) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextList10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlist.h:57
// index:0
// Public
// void QTextList(class QTextDocument *)
func NewQTextList(doc unsafe.Pointer) *QTextList {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextListC2EP13QTextDocument", ffiqt.FFI_TYPE_VOID, cthis, doc)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextListFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextlist.h:58
// index:0
// Public virtual
// void ~QTextList()
func DeleteQTextList(*QTextList) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextListD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:60
// index:0
// Public
// int count()
func (this *QTextList) Count() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextList5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlist.h:62
// index:0
// Public inline
// bool isEmpty()
func (this *QTextList) IsEmpty() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextList7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlist.h:65
// index:0
// Public
// QTextBlock item(int)
func (this *QTextList) Item(i int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextList4itemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlist.h:67
// index:0
// Public
// int itemNumber(const class QTextBlock &)
func (this *QTextList) ItemNumber(arg0 *QTextBlock) interface{} {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextList10itemNumberERK10QTextBlock", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlist.h:68
// index:0
// Public
// QString itemText(const class QTextBlock &)
func (this *QTextList) ItemText(arg0 *QTextBlock) interface{} {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextList8itemTextERK10QTextBlock", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlist.h:70
// index:0
// Public
// void removeItem(int)
func (this *QTextList) RemoveItem(i int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextList10removeItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:71
// index:0
// Public
// void remove(const class QTextBlock &)
func (this *QTextList) Remove(arg0 *QTextBlock) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextList6removeERK10QTextBlock", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:73
// index:0
// Public
// void add(const class QTextBlock &)
func (this *QTextList) Add(block *QTextBlock) {
	var convArg0 = block.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextList3addERK10QTextBlock", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:75
// index:0
// Public inline
// void setFormat(const class QTextListFormat &)
func (this *QTextList) SetFormat(format *QTextListFormat) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextList9setFormatERK15QTextListFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:76
// index:0
// Public inline
// QTextListFormat format()
func (this *QTextList) Format() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextList6formatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
