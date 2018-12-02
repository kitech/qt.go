package qtgui

// /usr/include/qt/QtGui/qtextlist.h
// #include <qtextlist.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
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

/*

 */
type QTextList struct {
	*QTextBlockGroup
}
type QTextList_ITF interface {
	QTextBlockGroup_ITF
	QTextList_PTR() *QTextList
}

func (ptr *QTextList) QTextList_PTR() *QTextList { return ptr }

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QTextList) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextList10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextlist.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextList(QTextDocument *)

/*

 */
func (*QTextList) NewForInherit(doc QTextDocument_ITF /*777 QTextDocument **/) *QTextList {
	return NewQTextList(doc)
}
func NewQTextList(doc QTextDocument_ITF /*777 QTextDocument **/) *QTextList {
	var convArg0 unsafe.Pointer
	if doc != nil && doc.QTextDocument_PTR() != nil {
		convArg0 = doc.QTextDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextListC2EP13QTextDocument", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextListFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTextList")
	return gothis
}

// /usr/include/qt/QtGui/qtextlist.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTextList()

/*

 */
func DeleteQTextList(this *QTextList) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextListD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtextlist.h:60
// index:0
// Public Visibility=Default Availability=Available
// [4] int count() const

/*
Returns the number of items in the list.
*/
func (this *QTextList) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextList5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlist.h:62
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*

 */
func (this *QTextList) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextList7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextlist.h:65
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlock item(int) const

/*
Returns the i-th text block in the list.

See also count() and itemText().
*/
func (this *QTextList) Item(i int) *QTextBlock /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextList4itemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

// /usr/include/qt/QtGui/qtextlist.h:67
// index:0
// Public Visibility=Default Availability=Available
// [4] int itemNumber(const QTextBlock &) const

/*
Returns the index of the list item that corresponds to the given block. Returns -1 if the block was not present in the list.
*/
func (this *QTextList) ItemNumber(arg0 QTextBlock_ITF) int {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTextBlock_PTR() != nil {
		convArg0 = arg0.QTextBlock_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextList10itemNumberERK10QTextBlock", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlist.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QString itemText(const QTextBlock &) const

/*
Returns the text of the list item that corresponds to the given block.
*/
func (this *QTextList) ItemText(arg0 QTextBlock_ITF) string {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTextBlock_PTR() != nil {
		convArg0 = arg0.QTextBlock_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextList8itemTextERK10QTextBlock", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextlist.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeItem(int)

/*
Removes the item at item position i from the list. When the last item in the list is removed, the list is automatically deleted by the QTextDocument that owns it.

See also add() and remove().
*/
func (this *QTextList) RemoveItem(i int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextList10removeItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void remove(const QTextBlock &)

/*
Removes the given block from the list.

See also add() and removeItem().
*/
func (this *QTextList) Remove(arg0 QTextBlock_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTextBlock_PTR() != nil {
		convArg0 = arg0.QTextBlock_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextList6removeERK10QTextBlock", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void add(const QTextBlock &)

/*
Makes the given block part of the list.

See also remove() and removeItem().
*/
func (this *QTextList) Add(block QTextBlock_ITF) {
	var convArg0 unsafe.Pointer
	if block != nil && block.QTextBlock_PTR() != nil {
		convArg0 = block.QTextBlock_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextList3addERK10QTextBlock", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:75
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFormat(const QTextListFormat &)

/*
Sets the list's format to format.

See also format().
*/
func (this *QTextList) SetFormat(format QTextListFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QTextListFormat_PTR() != nil {
		convArg0 = format.QTextListFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextList9setFormatERK15QTextListFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlist.h:76
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QTextListFormat format() const

/*
Returns the list's format.

See also setFormat().
*/
func (this *QTextList) Format() *QTextListFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextList6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextListFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextListFormat)
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
