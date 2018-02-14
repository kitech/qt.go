package qtwidgets

// /usr/include/qt/QtWidgets/qlayoutitem.h
// #include <qlayoutitem.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 58
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

type QLayoutItem struct {
	*qtrt.CObject
}
type QLayoutItem_ITF interface {
	QLayoutItem_PTR() *QLayoutItem
}

func (ptr *QLayoutItem) QLayoutItem_PTR() *QLayoutItem { return ptr }

func (this *QLayoutItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QLayoutItem) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQLayoutItemFromPointer(cthis unsafe.Pointer) *QLayoutItem {
	return &QLayoutItem{&qtrt.CObject{cthis}}
}
func (*QLayoutItem) NewFromPointer(cthis unsafe.Pointer) *QLayoutItem {
	return NewQLayoutItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:63
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QLayoutItem(Qt::Alignment)
func NewQLayoutItem(alignment int) *QLayoutItem {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QLayoutItemC2E6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, alignment)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLayoutItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QLayoutItem()
func DeleteQLayoutItem(this *QLayoutItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QLayoutItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:65
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QLayoutItem) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLayoutItem8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:66
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSize minimumSize()
func (this *QLayoutItem) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLayoutItem11minimumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:67
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSize maximumSize()
func (this *QLayoutItem) MaximumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLayoutItem11maximumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:68
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] Qt::Orientations expandingDirections()
func (this *QLayoutItem) ExpandingDirections() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLayoutItem19expandingDirectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:69
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRect &)
func (this *QLayoutItem) SetGeometry(arg0 qtcore.QRect_ITF) {
	var convArg0 = arg0.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QLayoutItem11setGeometryERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:70
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QRect geometry()
func (this *QLayoutItem) Geometry() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLayoutItem8geometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:71
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QLayoutItem) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLayoutItem7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasHeightForWidth()
func (this *QLayoutItem) HasHeightForWidth() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLayoutItem17hasHeightForWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int heightForWidth(int)
func (this *QLayoutItem) HeightForWidth(arg0 int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLayoutItem14heightForWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int minimumHeightForWidth(int)
func (this *QLayoutItem) MinimumHeightForWidth(arg0 int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLayoutItem21minimumHeightForWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void invalidate()
func (this *QLayoutItem) Invalidate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QLayoutItem10invalidateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QWidget * widget()
func (this *QLayoutItem) Widget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QLayoutItem6widgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QLayout * layout()
func (this *QLayoutItem) Layout() *QLayout /*777 QLayout **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QLayoutItem6layoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSpacerItem * spacerItem()
func (this *QLayoutItem) SpacerItem() *QSpacerItem /*777 QSpacerItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QLayoutItem10spacerItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSpacerItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::Alignment alignment()
func (this *QLayoutItem) Alignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLayoutItem9alignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlignment(Qt::Alignment)
func (this *QLayoutItem) SetAlignment(a int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QLayoutItem12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QSizePolicy::ControlTypes controlTypes()
func (this *QLayoutItem) ControlTypes() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLayoutItem12controlTypesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
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
