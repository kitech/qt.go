package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h
// #include <qgraphicsgridlayout.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QGraphicsGridLayout struct {
	*QGraphicsLayout
}

func (this *QGraphicsGridLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsLayout.GetCthis()
	}
}
func (this *QGraphicsGridLayout) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsLayout = NewQGraphicsLayoutFromPointer(cthis)
}
func NewQGraphicsGridLayoutFromPointer(cthis unsafe.Pointer) *QGraphicsGridLayout {
	bcthis0 := NewQGraphicsLayoutFromPointer(cthis)
	return &QGraphicsGridLayout{bcthis0}
}
func (*QGraphicsGridLayout) NewFromPointer(cthis unsafe.Pointer) *QGraphicsGridLayout {
	return NewQGraphicsGridLayoutFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:56
// index:0
// Public
// void QGraphicsGridLayout(class QGraphicsLayoutItem *)
func NewQGraphicsGridLayout(parent *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/) *QGraphicsGridLayout {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayoutC2EP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsGridLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:57
// index:0
// Public virtual
// void ~QGraphicsGridLayout()
func DeleteQGraphicsGridLayout(*QGraphicsGridLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:63
// index:0
// Public
// void setHorizontalSpacing(qreal)
func (this *QGraphicsGridLayout) SetHorizontalSpacing(spacing float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout20setHorizontalSpacingEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:64
// index:0
// Public
// qreal horizontalSpacing()
func (this *QGraphicsGridLayout) HorizontalSpacing() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout17horizontalSpacingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:65
// index:0
// Public
// void setVerticalSpacing(qreal)
func (this *QGraphicsGridLayout) SetVerticalSpacing(spacing float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout18setVerticalSpacingEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:66
// index:0
// Public
// qreal verticalSpacing()
func (this *QGraphicsGridLayout) VerticalSpacing() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout15verticalSpacingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:67
// index:0
// Public
// void setSpacing(qreal)
func (this *QGraphicsGridLayout) SetSpacing(spacing float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout10setSpacingEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:69
// index:0
// Public
// void setRowSpacing(int, qreal)
func (this *QGraphicsGridLayout) SetRowSpacing(row int, spacing float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout13setRowSpacingEid", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:70
// index:0
// Public
// qreal rowSpacing(int)
func (this *QGraphicsGridLayout) RowSpacing(row int) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout10rowSpacingEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:71
// index:0
// Public
// void setColumnSpacing(int, qreal)
func (this *QGraphicsGridLayout) SetColumnSpacing(column int, spacing float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout16setColumnSpacingEid", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:72
// index:0
// Public
// qreal columnSpacing(int)
func (this *QGraphicsGridLayout) ColumnSpacing(column int) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout13columnSpacingEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:74
// index:0
// Public
// void setRowStretchFactor(int, int)
func (this *QGraphicsGridLayout) SetRowStretchFactor(row int, stretch int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout19setRowStretchFactorEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:75
// index:0
// Public
// int rowStretchFactor(int)
func (this *QGraphicsGridLayout) RowStretchFactor(row int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout16rowStretchFactorEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:76
// index:0
// Public
// void setColumnStretchFactor(int, int)
func (this *QGraphicsGridLayout) SetColumnStretchFactor(column int, stretch int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout22setColumnStretchFactorEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:77
// index:0
// Public
// int columnStretchFactor(int)
func (this *QGraphicsGridLayout) ColumnStretchFactor(column int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout19columnStretchFactorEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:79
// index:0
// Public
// void setRowMinimumHeight(int, qreal)
func (this *QGraphicsGridLayout) SetRowMinimumHeight(row int, height float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout19setRowMinimumHeightEid", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:80
// index:0
// Public
// qreal rowMinimumHeight(int)
func (this *QGraphicsGridLayout) RowMinimumHeight(row int) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout16rowMinimumHeightEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:81
// index:0
// Public
// void setRowPreferredHeight(int, qreal)
func (this *QGraphicsGridLayout) SetRowPreferredHeight(row int, height float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout21setRowPreferredHeightEid", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:82
// index:0
// Public
// qreal rowPreferredHeight(int)
func (this *QGraphicsGridLayout) RowPreferredHeight(row int) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout18rowPreferredHeightEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:83
// index:0
// Public
// void setRowMaximumHeight(int, qreal)
func (this *QGraphicsGridLayout) SetRowMaximumHeight(row int, height float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout19setRowMaximumHeightEid", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:84
// index:0
// Public
// qreal rowMaximumHeight(int)
func (this *QGraphicsGridLayout) RowMaximumHeight(row int) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout16rowMaximumHeightEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:85
// index:0
// Public
// void setRowFixedHeight(int, qreal)
func (this *QGraphicsGridLayout) SetRowFixedHeight(row int, height float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout17setRowFixedHeightEid", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:87
// index:0
// Public
// void setColumnMinimumWidth(int, qreal)
func (this *QGraphicsGridLayout) SetColumnMinimumWidth(column int, width float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout21setColumnMinimumWidthEid", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:88
// index:0
// Public
// qreal columnMinimumWidth(int)
func (this *QGraphicsGridLayout) ColumnMinimumWidth(column int) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout18columnMinimumWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:89
// index:0
// Public
// void setColumnPreferredWidth(int, qreal)
func (this *QGraphicsGridLayout) SetColumnPreferredWidth(column int, width float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout23setColumnPreferredWidthEid", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:90
// index:0
// Public
// qreal columnPreferredWidth(int)
func (this *QGraphicsGridLayout) ColumnPreferredWidth(column int) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout20columnPreferredWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:91
// index:0
// Public
// void setColumnMaximumWidth(int, qreal)
func (this *QGraphicsGridLayout) SetColumnMaximumWidth(column int, width float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout21setColumnMaximumWidthEid", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:92
// index:0
// Public
// qreal columnMaximumWidth(int)
func (this *QGraphicsGridLayout) ColumnMaximumWidth(column int) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout18columnMaximumWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:93
// index:0
// Public
// void setColumnFixedWidth(int, qreal)
func (this *QGraphicsGridLayout) SetColumnFixedWidth(column int, width float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout19setColumnFixedWidthEid", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:96
// index:0
// Public
// Qt::Alignment rowAlignment(int)
func (this *QGraphicsGridLayout) RowAlignment(row int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout12rowAlignmentEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:98
// index:0
// Public
// Qt::Alignment columnAlignment(int)
func (this *QGraphicsGridLayout) ColumnAlignment(column int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout15columnAlignmentEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:101
// index:0
// Public
// Qt::Alignment alignment(class QGraphicsLayoutItem *)
func (this *QGraphicsGridLayout) Alignment(item *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/) int {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout9alignmentEP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:103
// index:0
// Public
// int rowCount()
func (this *QGraphicsGridLayout) RowCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout8rowCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:104
// index:0
// Public
// int columnCount()
func (this *QGraphicsGridLayout) ColumnCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout11columnCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:106
// index:0
// Public
// QGraphicsLayoutItem * itemAt(int, int)
func (this *QGraphicsGridLayout) ItemAt(row int, column int) *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout6itemAtEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:110
// index:1
// Public virtual
// QGraphicsLayoutItem * itemAt(int)
func (this *QGraphicsGridLayout) ItemAt_1(index int) *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout6itemAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:109
// index:0
// Public virtual
// int count()
func (this *QGraphicsGridLayout) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:111
// index:0
// Public virtual
// void removeAt(int)
func (this *QGraphicsGridLayout) RemoveAt(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout8removeAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:112
// index:0
// Public
// void removeItem(class QGraphicsLayoutItem *)
func (this *QGraphicsGridLayout) RemoveItem(item *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout10removeItemEP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:114
// index:0
// Public virtual
// void invalidate()
func (this *QGraphicsGridLayout) Invalidate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout10invalidateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:117
// index:0
// Public virtual
// void setGeometry(const class QRectF &)
func (this *QGraphicsGridLayout) SetGeometry(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout11setGeometryERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:118
// index:0
// Public virtual
// QSizeF sizeHint(Qt::SizeHint, const class QSizeF &)
func (this *QGraphicsGridLayout) SizeHint(which int, constraint *qtcore.QSizeF) *qtcore.QSizeF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = constraint.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout8sizeHintEN2Qt8SizeHintERK6QSizeF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), which, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
