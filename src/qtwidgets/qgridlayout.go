//  header block begin
// /usr/include/qt/QtWidgets/qgridlayout.h
// #include <qgridlayout.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 54
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
type QGridLayout struct {
	*QLayout
}

func (this *QGridLayout) GetCthis() unsafe.Pointer {
	return this.QLayout.GetCthis()
}

// /usr/include/qt/QtWidgets/qgridlayout.h:58
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGridLayout) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:64
// index:0
// void QGridLayout(class QWidget *)
func NewQGridLayout(parent unsafe.Pointer) *QGridLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QGridLayoutC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGridLayoutFromPointer(cthis)
	return gothis
}
func NewQGridLayoutFromPointer(cthis unsafe.Pointer) *QGridLayout {
	bcthis0 := NewQLayoutFromPointer(cthis)
	return &QGridLayout{bcthis0}
}

// /usr/include/qt/QtWidgets/qgridlayout.h:65
// index:1
// void QGridLayout()
func NewQGridLayout_1() *QGridLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QGridLayoutC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQGridLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgridlayout.h:67
// index:0
// virtual
// void ~QGridLayout()
func DeleteQGridLayout(*QGridLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QGridLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:69
// index:0
// virtual
// QSize sizeHint()
func (this *QGridLayout) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:70
// index:0
// virtual
// QSize minimumSize()
func (this *QGridLayout) MinimumSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout11minimumSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:71
// index:0
// virtual
// QSize maximumSize()
func (this *QGridLayout) MaximumSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout11maximumSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:73
// index:0
// void setHorizontalSpacing(int)
func (this *QGridLayout) SetHorizontalSpacing(spacing int) {
	// 0: (, spacing int), (&spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QGridLayout20setHorizontalSpacingEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:74
// index:0
// int horizontalSpacing()
func (this *QGridLayout) HorizontalSpacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout17horizontalSpacingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:75
// index:0
// void setVerticalSpacing(int)
func (this *QGridLayout) SetVerticalSpacing(spacing int) {
	// 0: (, spacing int), (&spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QGridLayout18setVerticalSpacingEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:76
// index:0
// int verticalSpacing()
func (this *QGridLayout) VerticalSpacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout15verticalSpacingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:77
// index:0
// void setSpacing(int)
func (this *QGridLayout) SetSpacing(spacing int) {
	// 0: (, spacing int), (&spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QGridLayout10setSpacingEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:78
// index:0
// int spacing()
func (this *QGridLayout) Spacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout7spacingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:80
// index:0
// void setRowStretch(int, int)
func (this *QGridLayout) SetRowStretch(row int, stretch int) {
	// 0: (, row int, stretch int), (&row, &stretch)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QGridLayout13setRowStretchEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:81
// index:0
// void setColumnStretch(int, int)
func (this *QGridLayout) SetColumnStretch(column int, stretch int) {
	// 0: (, column int, stretch int), (&column, &stretch)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QGridLayout16setColumnStretchEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:82
// index:0
// int rowStretch(int)
func (this *QGridLayout) RowStretch(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout10rowStretchEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:83
// index:0
// int columnStretch(int)
func (this *QGridLayout) ColumnStretch(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout13columnStretchEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:85
// index:0
// void setRowMinimumHeight(int, int)
func (this *QGridLayout) SetRowMinimumHeight(row int, minSize int) {
	// 0: (, row int, minSize int), (&row, &minSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QGridLayout19setRowMinimumHeightEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &minSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:86
// index:0
// void setColumnMinimumWidth(int, int)
func (this *QGridLayout) SetColumnMinimumWidth(column int, minSize int) {
	// 0: (, column int, minSize int), (&column, &minSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QGridLayout21setColumnMinimumWidthEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &minSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:87
// index:0
// int rowMinimumHeight(int)
func (this *QGridLayout) RowMinimumHeight(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout16rowMinimumHeightEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:88
// index:0
// int columnMinimumWidth(int)
func (this *QGridLayout) ColumnMinimumWidth(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout18columnMinimumWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:90
// index:0
// int columnCount()
func (this *QGridLayout) ColumnCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout11columnCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:91
// index:0
// int rowCount()
func (this *QGridLayout) RowCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout8rowCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:93
// index:0
// QRect cellRect(int, int)
func (this *QGridLayout) CellRect(row int, column int) {
	// 0: (, row int, column int), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout8cellRectEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:95
// index:0
// virtual
// bool hasHeightForWidth()
func (this *QGridLayout) HasHeightForWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout17hasHeightForWidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:96
// index:0
// virtual
// int heightForWidth(int)
func (this *QGridLayout) HeightForWidth(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout14heightForWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:97
// index:0
// virtual
// int minimumHeightForWidth(int)
func (this *QGridLayout) MinimumHeightForWidth(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout21minimumHeightForWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:99
// index:0
// virtual
// Qt::Orientations expandingDirections()
func (this *QGridLayout) ExpandingDirections() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout19expandingDirectionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:100
// index:0
// virtual
// void invalidate()
func (this *QGridLayout) Invalidate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QGridLayout10invalidateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:102
// index:0
// inline
// void addWidget(class QWidget *)
func (this *QGridLayout) AddWidget(w unsafe.Pointer) {
	// 0: (, w QWidget *), (w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QGridLayout9addWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:103
// index:1
// void addWidget(class QWidget *, int, int, Qt::Alignment)
func (this *QGridLayout) AddWidget_1(arg0 unsafe.Pointer, row int, column int, arg3 int) {
	// 1: (, QWidget * arg0, row int, column int, Qt::Alignment arg3), (arg0, &row, &column, &arg3)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QGridLayout9addWidgetEP7QWidgetii6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, &row, &column, &arg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:104
// index:2
// void addWidget(class QWidget *, int, int, int, int, Qt::Alignment)
func (this *QGridLayout) AddWidget_2(arg0 unsafe.Pointer, row int, column int, rowSpan int, columnSpan int, arg5 int) {
	// 2: (, QWidget * arg0, row int, column int, rowSpan int, columnSpan int, Qt::Alignment arg5), (arg0, &row, &column, &rowSpan, &columnSpan, &arg5)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QGridLayout9addWidgetEP7QWidgetiiii6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, &row, &column, &rowSpan, &columnSpan, &arg5)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:105
// index:0
// void addLayout(class QLayout *, int, int, Qt::Alignment)
func (this *QGridLayout) AddLayout(arg0 unsafe.Pointer, row int, column int, arg3 int) {
	// 0: (, QLayout * arg0, row int, column int, Qt::Alignment arg3), (arg0, &row, &column, &arg3)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QGridLayout9addLayoutEP7QLayoutii6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, &row, &column, &arg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:106
// index:1
// void addLayout(class QLayout *, int, int, int, int, Qt::Alignment)
func (this *QGridLayout) AddLayout_1(arg0 unsafe.Pointer, row int, column int, rowSpan int, columnSpan int, arg5 int) {
	// 1: (, QLayout * arg0, row int, column int, rowSpan int, columnSpan int, Qt::Alignment arg5), (arg0, &row, &column, &rowSpan, &columnSpan, &arg5)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QGridLayout9addLayoutEP7QLayoutiiii6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, &row, &column, &rowSpan, &columnSpan, &arg5)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:108
// index:0
// void setOriginCorner(Qt::Corner)
func (this *QGridLayout) SetOriginCorner(arg0 int) {
	// 0: (, Qt::Corner arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QGridLayout15setOriginCornerEN2Qt6CornerE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:109
// index:0
// Qt::Corner originCorner()
func (this *QGridLayout) OriginCorner() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout12originCornerEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:111
// index:0
// virtual
// QLayoutItem * itemAt(int)
func (this *QGridLayout) ItemAt(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout6itemAtEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:112
// index:0
// QLayoutItem * itemAtPosition(int, int)
func (this *QGridLayout) ItemAtPosition(row int, column int) {
	// 0: (, row int, column int), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout14itemAtPositionEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:113
// index:0
// virtual
// QLayoutItem * takeAt(int)
func (this *QGridLayout) TakeAt(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QGridLayout6takeAtEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:114
// index:0
// virtual
// int count()
func (this *QGridLayout) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout5countEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:115
// index:0
// virtual
// void setGeometry(const class QRect &)
func (this *QGridLayout) SetGeometry(arg0 unsafe.Pointer) {
	// 0: (, const QRect & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QGridLayout11setGeometryERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:117
// index:0
// void addItem(class QLayoutItem *, int, int, int, int, Qt::Alignment)
func (this *QGridLayout) AddItem(item unsafe.Pointer, row int, column int, rowSpan int, columnSpan int, arg5 int) {
	// 0: (, item QLayoutItem *, row int, column int, rowSpan int, columnSpan int, Qt::Alignment arg5), (item, &row, &column, &rowSpan, &columnSpan, &arg5)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QGridLayout7addItemEP11QLayoutItemiiii6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &row, &column, &rowSpan, &columnSpan, &arg5)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:123
// index:1
// virtual
// void addItem(class QLayoutItem *)
func (this *QGridLayout) AddItem_1(arg0 unsafe.Pointer) {
	// 1: (, QLayoutItem * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QGridLayout7addItemEP11QLayoutItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:119
// index:0
// void setDefaultPositioning(int, Qt::Orientation)
func (this *QGridLayout) SetDefaultPositioning(n int, orient int) {
	// 0: (, n int, orient Qt::Orientation), (&n, &orient)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QGridLayout21setDefaultPositioningEiN2Qt11OrientationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &n, &orient)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgridlayout.h:120
// index:0
// void getItemPosition(int, int *, int *, int *, int *)
func (this *QGridLayout) GetItemPosition(idx int, row unsafe.Pointer, column unsafe.Pointer, rowSpan unsafe.Pointer, columnSpan unsafe.Pointer) {
	// 0: (, idx int, row int *, column int *, rowSpan int *, columnSpan int *), (&idx, row, column, rowSpan, columnSpan)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QGridLayout15getItemPositionEiPiS0_S0_S0_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &idx, row, column, rowSpan, columnSpan)
	gopp.ErrPrint(err, rv)
}

//  body block end
