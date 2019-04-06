package qtmacextras

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h
// #include <qmactoolbaritem.h>
// #include <Qtmacextras>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
type QMacToolBarItem struct {
	*qtcore.QObject
}
type QMacToolBarItem_ITF interface {
	qtcore.QObject_ITF
	QMacToolBarItem_PTR() *QMacToolBarItem
}

func (ptr *QMacToolBarItem) QMacToolBarItem_PTR() *QMacToolBarItem { return ptr }

func (this *QMacToolBarItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QMacToolBarItem) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQMacToolBarItemFromPointer(cthis unsafe.Pointer) *QMacToolBarItem {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QMacToolBarItem{bcthis0}
}
func (*QMacToolBarItem) NewFromPointer(cthis unsafe.Pointer) *QMacToolBarItem {
	return NewQMacToolBarItemFromPointer(cthis)
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMacToolBarItem) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QMacToolBarItem10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMacToolBarItem(QObject *)

/*
Constructs a QMacToolBarItem with parent.
*/
func (*QMacToolBarItem) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QMacToolBarItem {
	return NewQMacToolBarItem(parent)
}
func NewQMacToolBarItem(parent qtcore.QObject_ITF /*777 QObject **/) *QMacToolBarItem {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMacToolBarItemC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMacToolBarItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMacToolBarItem")
	return gothis
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMacToolBarItem(QObject *)

/*
Constructs a QMacToolBarItem with parent.
*/
func (*QMacToolBarItem) NewForInheritp() *QMacToolBarItem {
	return NewQMacToolBarItemp()
}
func NewQMacToolBarItemp() *QMacToolBarItem {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMacToolBarItemC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMacToolBarItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMacToolBarItem")
	return gothis
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMacToolBarItem()

/*

 */
func DeleteQMacToolBarItem(this *QMacToolBarItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMacToolBarItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool selectable() const

/*

 */
func (this *QMacToolBarItem) Selectable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QMacToolBarItem10selectableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelectable(bool)

/*

 */
func (this *QMacToolBarItem) SetSelectable(selectable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMacToolBarItem13setSelectableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), selectable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] QMacToolBarItem::StandardItem standardItem() const

/*

 */
func (this *QMacToolBarItem) StandardItem() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QMacToolBarItem12standardItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStandardItem(QMacToolBarItem::StandardItem)

/*

 */
func (this *QMacToolBarItem) SetStandardItem(standardItem int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMacToolBarItem15setStandardItemENS_12StandardItemE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), standardItem)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text() const

/*

 */
func (this *QMacToolBarItem) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QMacToolBarItem4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)

/*

 */
func (this *QMacToolBarItem) SetText(text string) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMacToolBarItem7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:83
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon icon() const

/*

 */
func (this *QMacToolBarItem) Icon() *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QMacToolBarItem4iconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIcon(const QIcon &)

/*

 */
func (this *QMacToolBarItem) SetIcon(icon qtgui.QIcon_ITF) {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMacToolBarItem7setIconERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] NSToolbarItem * nativeToolBarItem() const

/*
Returns the native NSToolbarItem.
*/
func (this *QMacToolBarItem) NativeToolBarItem() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QMacToolBarItem17nativeToolBarItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activated()

/*
This signal is emitted when the toolbar item is clicked or otherwise activated.
*/
func (this *QMacToolBarItem) Activated() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMacToolBarItem9activatedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

/*

 */
type QMacToolBarItem__StandardItem = int

// Don't use a standard item
const QMacToolBarItem__NoStandardItem QMacToolBarItem__StandardItem = 0

// A spacing item
const QMacToolBarItem__Space QMacToolBarItem__StandardItem = 1

// A spacing item which grows to fill available space
const QMacToolBarItem__FlexibleSpace QMacToolBarItem__StandardItem = 2

func (this *QMacToolBarItem) StandardItemItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QMacToolBarItem_StandardItemItemName(val int) string {
	var nilthis *QMacToolBarItem
	return nilthis.StandardItemItemName(val)
}

//  body block end

//  keep block begin

func init_unused_14883() {
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
