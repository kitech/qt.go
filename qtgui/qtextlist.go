package qtgui

// /usr/include/qt/QtGui/qtextlist.h
// #include <qtextlist.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
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
type QTextList struct {
	*QTextBlockGroup
}

func (this *QTextList) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QTextBlockGroup.GetCthis()
	}
}
func (this *QTextList) SetCthis(cthis unsafe.Pointer) {
	this.QTextBlockGroup = NewQTextBlockGroupFromPointer(cthis)
}
func NewQTextListFromPointer(cthis unsafe.Pointer) *QTextList {
	bcthis0 := NewQTextBlockGroupFromPointer(cthis)
	return &QTextList{bcthis0}
}
func (*QTextList) NewFromPointer(cthis unsafe.Pointer) *QTextList {
	return NewQTextListFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextlist.h:55
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QTextList) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextList10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextlist.h:57
// index:0
// Public
// void QTextList(QTextDocument *)
func NewQTextList(doc *QTextDocument /*777 QTextDocument **/) *QTextList {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = doc.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextListC2EP13QTextDocument", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
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
func (this *QTextList) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextList5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlist.h:62
// index:0
// Public inline
// bool isEmpty()
func (this *QTextList) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextList7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextlist.h:65
// index:0
// Public
// QTextBlock item(int)
func (this *QTextList) Item(i int) *QTextBlock /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextList4itemEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextlist.h:67
// index:0
// Public
// int itemNumber(const QTextBlock &)
func (this *QTextList) ItemNumber(arg0 *QTextBlock) int {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextList10itemNumberERK10QTextBlock", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlist.h:68
// index:0
// Public
// QString itemText(const QTextBlock &)
func (this *QTextList) ItemText(arg0 *QTextBlock) *qtcore.QString /*123*/ {
	var convArg0 = arg0.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextList8itemTextERK10QTextBlock", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextlist.h:70
// index:0
// Public
// void removeItem(int)
func (this *QTextList) RemoveItem(i int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextList10removeItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:71
// index:0
// Public
// void remove(const QTextBlock &)
func (this *QTextList) Remove(arg0 *QTextBlock) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextList6removeERK10QTextBlock", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:73
// index:0
// Public
// void add(const QTextBlock &)
func (this *QTextList) Add(block *QTextBlock) {
	var convArg0 = block.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextList3addERK10QTextBlock", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:75
// index:0
// Public inline
// void setFormat(const QTextListFormat &)
func (this *QTextList) SetFormat(format *QTextListFormat) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextList9setFormatERK15QTextListFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:76
// index:0
// Public inline
// QTextListFormat format()
func (this *QTextList) Format() *QTextListFormat /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextList6formatEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextListFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
