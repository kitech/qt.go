package qtwidgets

// /usr/include/qt/QtWidgets/qlayoutitem.h
// #include <qlayoutitem.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 56
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QLayoutItem struct {
	*qtrt.CObject
}

func (this *QLayoutItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QLayoutItem) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQLayoutItemFromPointer(cthis unsafe.Pointer) *QLayoutItem {
	return &QLayoutItem{&qtrt.CObject{cthis}}
}
func (*QLayoutItem) NewFromPointer(cthis unsafe.Pointer) *QLayoutItem {
	return NewQLayoutItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:64
// index:0
// Public virtual
// void ~QLayoutItem()
func DeleteQLayoutItem(*QLayoutItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QLayoutItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:65
// index:0
// Public pure virtual
// QSize sizeHint()
func (this *QLayoutItem) SizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLayoutItem8sizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:66
// index:0
// Public pure virtual
// QSize minimumSize()
func (this *QLayoutItem) MinimumSize() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLayoutItem11minimumSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:67
// index:0
// Public pure virtual
// QSize maximumSize()
func (this *QLayoutItem) MaximumSize() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLayoutItem11maximumSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:68
// index:0
// Public pure virtual
// Qt::Orientations expandingDirections()
func (this *QLayoutItem) ExpandingDirections() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLayoutItem19expandingDirectionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:69
// index:0
// Public pure virtual
// void setGeometry(const QRect &)
func (this *QLayoutItem) SetGeometry(arg0 *qtcore.QRect) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QLayoutItem11setGeometryERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:70
// index:0
// Public pure virtual
// QRect geometry()
func (this *QLayoutItem) Geometry() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLayoutItem8geometryEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:71
// index:0
// Public pure virtual
// bool isEmpty()
func (this *QLayoutItem) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLayoutItem7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:72
// index:0
// Public virtual
// bool hasHeightForWidth()
func (this *QLayoutItem) HasHeightForWidth() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLayoutItem17hasHeightForWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:73
// index:0
// Public virtual
// int heightForWidth(int)
func (this *QLayoutItem) HeightForWidth(arg0 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLayoutItem14heightForWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:74
// index:0
// Public virtual
// int minimumHeightForWidth(int)
func (this *QLayoutItem) MinimumHeightForWidth(arg0 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLayoutItem21minimumHeightForWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:75
// index:0
// Public virtual
// void invalidate()
func (this *QLayoutItem) Invalidate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QLayoutItem10invalidateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:77
// index:0
// Public virtual
// QWidget * widget()
func (this *QLayoutItem) Widget() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QLayoutItem6widgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:78
// index:0
// Public virtual
// QLayout * layout()
func (this *QLayoutItem) Layout() *QLayout /*777 QLayout **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QLayoutItem6layoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:79
// index:0
// Public virtual
// QSpacerItem * spacerItem()
func (this *QLayoutItem) SpacerItem() *QSpacerItem /*777 QSpacerItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QLayoutItem10spacerItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSpacerItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:81
// index:0
// Public inline
// Qt::Alignment alignment()
func (this *QLayoutItem) Alignment() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLayoutItem9alignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:83
// index:0
// Public virtual
// QSizePolicy::ControlTypes controlTypes()
func (this *QLayoutItem) ControlTypes() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLayoutItem12controlTypesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

//  body block end
