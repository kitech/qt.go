package qtwidgets

// /usr/include/qt/QtWidgets/qgridlayout.h
// #include <qgridlayout.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 53
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

// void addItem(class QLayoutItem *)
func (this *QGridLayout) InheritAddItem(f func(arg0 *QLayoutItem /*777 QLayoutItem **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "addItem", f)
}

type QGridLayout struct {
	*QLayout
}
type QGridLayout_ITF interface {
	QLayout_ITF
	QGridLayout_PTR() *QGridLayout
}

func (ptr *QGridLayout) QGridLayout_PTR() *QGridLayout { return ptr }

func (this *QGridLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QLayout.GetCthis()
	}
}
func (this *QGridLayout) SetCthis(cthis unsafe.Pointer) {
	this.QLayout = NewQLayoutFromPointer(cthis)
}
func NewQGridLayoutFromPointer(cthis unsafe.Pointer) *QGridLayout {
	bcthis0 := NewQLayoutFromPointer(cthis)
	return &QGridLayout{bcthis0}
}
func (*QGridLayout) NewFromPointer(cthis unsafe.Pointer) *QGridLayout {
	return NewQGridLayoutFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QGridLayout) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgridlayout.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGridLayout(QWidget *)
func NewQGridLayout(parent QWidget_ITF /*777 QWidget **/) *QGridLayout {
	var convArg0 = parent.QWidget_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGridLayoutC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGridLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qgridlayout.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGridLayout()
func NewQGridLayout_1() *QGridLayout {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGridLayoutC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGridLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qgridlayout.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGridLayout()
func DeleteQGridLayout(this *QGridLayout) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGridLayoutD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QGridLayout) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qgridlayout.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSize()
func (this *QGridLayout) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout11minimumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qgridlayout.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize maximumSize()
func (this *QGridLayout) MaximumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout11maximumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qgridlayout.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalSpacing(int)
func (this *QGridLayout) SetHorizontalSpacing(spacing int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGridLayout20setHorizontalSpacingEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:74
// index:0
// Public Visibility=Default Availability=Available
// [4] int horizontalSpacing()
func (this *QGridLayout) HorizontalSpacing() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout17horizontalSpacingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgridlayout.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalSpacing(int)
func (this *QGridLayout) SetVerticalSpacing(spacing int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGridLayout18setVerticalSpacingEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int verticalSpacing()
func (this *QGridLayout) VerticalSpacing() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout15verticalSpacingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgridlayout.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpacing(int)
func (this *QGridLayout) SetSpacing(spacing int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGridLayout10setSpacingEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] int spacing()
func (this *QGridLayout) Spacing() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout7spacingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgridlayout.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowStretch(int, int)
func (this *QGridLayout) SetRowStretch(row int, stretch int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGridLayout13setRowStretchEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnStretch(int, int)
func (this *QGridLayout) SetColumnStretch(column int, stretch int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGridLayout16setColumnStretchEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] int rowStretch(int)
func (this *QGridLayout) RowStretch(row int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout10rowStretchEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgridlayout.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnStretch(int)
func (this *QGridLayout) ColumnStretch(column int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout13columnStretchEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgridlayout.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowMinimumHeight(int, int)
func (this *QGridLayout) SetRowMinimumHeight(row int, minSize int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGridLayout19setRowMinimumHeightEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, minSize)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnMinimumWidth(int, int)
func (this *QGridLayout) SetColumnMinimumWidth(column int, minSize int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGridLayout21setColumnMinimumWidthEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, minSize)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] int rowMinimumHeight(int)
func (this *QGridLayout) RowMinimumHeight(row int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout16rowMinimumHeightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgridlayout.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnMinimumWidth(int)
func (this *QGridLayout) ColumnMinimumWidth(column int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout18columnMinimumWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgridlayout.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnCount()
func (this *QGridLayout) ColumnCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout11columnCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgridlayout.h:91
// index:0
// Public Visibility=Default Availability=Available
// [4] int rowCount()
func (this *QGridLayout) RowCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout8rowCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgridlayout.h:93
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect cellRect(int, int)
func (this *QGridLayout) CellRect(row int, column int) *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout8cellRectEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qgridlayout.h:95
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasHeightForWidth()
func (this *QGridLayout) HasHeightForWidth() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout17hasHeightForWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgridlayout.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int heightForWidth(int)
func (this *QGridLayout) HeightForWidth(arg0 int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout14heightForWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgridlayout.h:97
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int minimumHeightForWidth(int)
func (this *QGridLayout) MinimumHeightForWidth(arg0 int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout21minimumHeightForWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgridlayout.h:99
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::Orientations expandingDirections()
func (this *QGridLayout) ExpandingDirections() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout19expandingDirectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:100
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void invalidate()
func (this *QGridLayout) Invalidate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGridLayout10invalidateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:102
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void addWidget(QWidget *)
func (this *QGridLayout) AddWidget(w QWidget_ITF /*777 QWidget **/) {
	var convArg0 = w.QWidget_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGridLayout9addWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:103
// index:1
// Public Visibility=Default Availability=Available
// [-2] void addWidget(QWidget *, int, int, Qt::Alignment)
func (this *QGridLayout) AddWidget_1(arg0 QWidget_ITF /*777 QWidget **/, row int, column int, arg3 int) {
	var convArg0 = arg0.QWidget_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGridLayout9addWidgetEP7QWidgetii6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, row, column, arg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:104
// index:2
// Public Visibility=Default Availability=Available
// [-2] void addWidget(QWidget *, int, int, int, int, Qt::Alignment)
func (this *QGridLayout) AddWidget_2(arg0 QWidget_ITF /*777 QWidget **/, row int, column int, rowSpan int, columnSpan int, arg5 int) {
	var convArg0 = arg0.QWidget_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGridLayout9addWidgetEP7QWidgetiiii6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, row, column, rowSpan, columnSpan, arg5)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addLayout(QLayout *, int, int, Qt::Alignment)
func (this *QGridLayout) AddLayout(arg0 QLayout_ITF /*777 QLayout **/, row int, column int, arg3 int) {
	var convArg0 = arg0.QLayout_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGridLayout9addLayoutEP7QLayoutii6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, row, column, arg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:106
// index:1
// Public Visibility=Default Availability=Available
// [-2] void addLayout(QLayout *, int, int, int, int, Qt::Alignment)
func (this *QGridLayout) AddLayout_1(arg0 QLayout_ITF /*777 QLayout **/, row int, column int, rowSpan int, columnSpan int, arg5 int) {
	var convArg0 = arg0.QLayout_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGridLayout9addLayoutEP7QLayoutiiii6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, row, column, rowSpan, columnSpan, arg5)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOriginCorner(Qt::Corner)
func (this *QGridLayout) SetOriginCorner(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGridLayout15setOriginCornerEN2Qt6CornerE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:109
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Corner originCorner()
func (this *QGridLayout) OriginCorner() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout12originCornerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:111
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QLayoutItem * itemAt(int)
func (this *QGridLayout) ItemAt(index int) *QLayoutItem /*777 QLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout6itemAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgridlayout.h:112
// index:0
// Public Visibility=Default Availability=Available
// [8] QLayoutItem * itemAtPosition(int, int)
func (this *QGridLayout) ItemAtPosition(row int, column int) *QLayoutItem /*777 QLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout14itemAtPositionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgridlayout.h:113
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QLayoutItem * takeAt(int)
func (this *QGridLayout) TakeAt(index int) *QLayoutItem /*777 QLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGridLayout6takeAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgridlayout.h:114
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int count()
func (this *QGridLayout) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgridlayout.h:115
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRect &)
func (this *QGridLayout) SetGeometry(arg0 qtcore.QRect_ITF) {
	var convArg0 = arg0.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGridLayout11setGeometryERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addItem(QLayoutItem *, int, int, int, int, Qt::Alignment)
func (this *QGridLayout) AddItem(item QLayoutItem_ITF /*777 QLayoutItem **/, row int, column int, rowSpan int, columnSpan int, arg5 int) {
	var convArg0 = item.QLayoutItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGridLayout7addItemEP11QLayoutItemiiii6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, row, column, rowSpan, columnSpan, arg5)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:123
// index:1
// Protected virtual Visibility=Default Availability=Available
// [-2] void addItem(QLayoutItem *)
func (this *QGridLayout) AddItem_1(arg0 QLayoutItem_ITF /*777 QLayoutItem **/) {
	var convArg0 = arg0.QLayoutItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGridLayout7addItemEP11QLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultPositioning(int, Qt::Orientation)
func (this *QGridLayout) SetDefaultPositioning(n int, orient int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGridLayout21setDefaultPositioningEiN2Qt11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n, orient)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getItemPosition(int, int *, int *, int *, int *)
func (this *QGridLayout) GetItemPosition(idx int, row unsafe.Pointer /*666*/, column unsafe.Pointer /*666*/, rowSpan unsafe.Pointer /*666*/, columnSpan unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGridLayout15getItemPositionEiPiS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), idx, &row, &column, &rowSpan, &columnSpan)
	qtrt.ErrPrint(err, rv)
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
