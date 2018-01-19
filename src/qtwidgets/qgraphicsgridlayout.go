//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h
// #include <qgraphicsgridlayout.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:56
// index:0
// void QGraphicsGridLayout(class QGraphicsLayoutItem *)
func NewQGraphicsGridLayout(parent unsafe.Pointer) *QGraphicsGridLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayoutC2EP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsGridLayout{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:57
// index:0
// virtual
// void ~QGraphicsGridLayout()
func DeleteQGraphicsGridLayout(*QGraphicsGridLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:63
// index:0
// void setHorizontalSpacing(qreal)
func (this *QGraphicsGridLayout) SetHorizontalSpacing(spacing float64) {
	// 0: (, qreal spacing), (&spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout20setHorizontalSpacingEd", ffiqt.FFI_TYPE_VOID, this.cthis, &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:64
// index:0
// qreal horizontalSpacing()
func (this *QGraphicsGridLayout) HorizontalSpacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout17horizontalSpacingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:65
// index:0
// void setVerticalSpacing(qreal)
func (this *QGraphicsGridLayout) SetVerticalSpacing(spacing float64) {
	// 0: (, qreal spacing), (&spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout18setVerticalSpacingEd", ffiqt.FFI_TYPE_VOID, this.cthis, &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:66
// index:0
// qreal verticalSpacing()
func (this *QGraphicsGridLayout) VerticalSpacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout15verticalSpacingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:67
// index:0
// void setSpacing(qreal)
func (this *QGraphicsGridLayout) SetSpacing(spacing float64) {
	// 0: (, qreal spacing), (&spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout10setSpacingEd", ffiqt.FFI_TYPE_VOID, this.cthis, &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:69
// index:0
// void setRowSpacing(int, qreal)
func (this *QGraphicsGridLayout) SetRowSpacing(row int, spacing float64) {
	// 0: (, int row, qreal spacing), (&row, &spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout13setRowSpacingEid", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:70
// index:0
// qreal rowSpacing(int)
func (this *QGraphicsGridLayout) RowSpacing(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout10rowSpacingEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:71
// index:0
// void setColumnSpacing(int, qreal)
func (this *QGraphicsGridLayout) SetColumnSpacing(column int, spacing float64) {
	// 0: (, int column, qreal spacing), (&column, &spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout16setColumnSpacingEid", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:72
// index:0
// qreal columnSpacing(int)
func (this *QGraphicsGridLayout) ColumnSpacing(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout13columnSpacingEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:74
// index:0
// void setRowStretchFactor(int, int)
func (this *QGraphicsGridLayout) SetRowStretchFactor(row int, stretch int) {
	// 0: (, int row, int stretch), (&row, &stretch)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout19setRowStretchFactorEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:75
// index:0
// int rowStretchFactor(int)
func (this *QGraphicsGridLayout) RowStretchFactor(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout16rowStretchFactorEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:76
// index:0
// void setColumnStretchFactor(int, int)
func (this *QGraphicsGridLayout) SetColumnStretchFactor(column int, stretch int) {
	// 0: (, int column, int stretch), (&column, &stretch)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout22setColumnStretchFactorEii", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:77
// index:0
// int columnStretchFactor(int)
func (this *QGraphicsGridLayout) ColumnStretchFactor(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout19columnStretchFactorEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:79
// index:0
// void setRowMinimumHeight(int, qreal)
func (this *QGraphicsGridLayout) SetRowMinimumHeight(row int, height float64) {
	// 0: (, int row, qreal height), (&row, &height)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout19setRowMinimumHeightEid", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:80
// index:0
// qreal rowMinimumHeight(int)
func (this *QGraphicsGridLayout) RowMinimumHeight(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout16rowMinimumHeightEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:81
// index:0
// void setRowPreferredHeight(int, qreal)
func (this *QGraphicsGridLayout) SetRowPreferredHeight(row int, height float64) {
	// 0: (, int row, qreal height), (&row, &height)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout21setRowPreferredHeightEid", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:82
// index:0
// qreal rowPreferredHeight(int)
func (this *QGraphicsGridLayout) RowPreferredHeight(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout18rowPreferredHeightEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:83
// index:0
// void setRowMaximumHeight(int, qreal)
func (this *QGraphicsGridLayout) SetRowMaximumHeight(row int, height float64) {
	// 0: (, int row, qreal height), (&row, &height)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout19setRowMaximumHeightEid", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:84
// index:0
// qreal rowMaximumHeight(int)
func (this *QGraphicsGridLayout) RowMaximumHeight(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout16rowMaximumHeightEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:85
// index:0
// void setRowFixedHeight(int, qreal)
func (this *QGraphicsGridLayout) SetRowFixedHeight(row int, height float64) {
	// 0: (, int row, qreal height), (&row, &height)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout17setRowFixedHeightEid", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:87
// index:0
// void setColumnMinimumWidth(int, qreal)
func (this *QGraphicsGridLayout) SetColumnMinimumWidth(column int, width float64) {
	// 0: (, int column, qreal width), (&column, &width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout21setColumnMinimumWidthEid", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:88
// index:0
// qreal columnMinimumWidth(int)
func (this *QGraphicsGridLayout) ColumnMinimumWidth(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout18columnMinimumWidthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:89
// index:0
// void setColumnPreferredWidth(int, qreal)
func (this *QGraphicsGridLayout) SetColumnPreferredWidth(column int, width float64) {
	// 0: (, int column, qreal width), (&column, &width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout23setColumnPreferredWidthEid", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:90
// index:0
// qreal columnPreferredWidth(int)
func (this *QGraphicsGridLayout) ColumnPreferredWidth(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout20columnPreferredWidthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:91
// index:0
// void setColumnMaximumWidth(int, qreal)
func (this *QGraphicsGridLayout) SetColumnMaximumWidth(column int, width float64) {
	// 0: (, int column, qreal width), (&column, &width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout21setColumnMaximumWidthEid", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:92
// index:0
// qreal columnMaximumWidth(int)
func (this *QGraphicsGridLayout) ColumnMaximumWidth(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout18columnMaximumWidthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:93
// index:0
// void setColumnFixedWidth(int, qreal)
func (this *QGraphicsGridLayout) SetColumnFixedWidth(column int, width float64) {
	// 0: (, int column, qreal width), (&column, &width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout19setColumnFixedWidthEid", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:96
// index:0
// Qt::Alignment rowAlignment(int)
func (this *QGraphicsGridLayout) RowAlignment(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout12rowAlignmentEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:98
// index:0
// Qt::Alignment columnAlignment(int)
func (this *QGraphicsGridLayout) ColumnAlignment(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout15columnAlignmentEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:101
// index:0
// Qt::Alignment alignment(class QGraphicsLayoutItem *)
func (this *QGraphicsGridLayout) Alignment(item unsafe.Pointer) {
	// 0: (, QGraphicsLayoutItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout9alignmentEP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:103
// index:0
// int rowCount()
func (this *QGraphicsGridLayout) RowCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout8rowCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:104
// index:0
// int columnCount()
func (this *QGraphicsGridLayout) ColumnCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout11columnCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:106
// index:0
// QGraphicsLayoutItem * itemAt(int, int)
func (this *QGraphicsGridLayout) ItemAt(row int, column int) {
	// 0: (, int row, int column), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout6itemAtEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:110
// index:1
// virtual
// QGraphicsLayoutItem * itemAt(int)
func (this *QGraphicsGridLayout) ItemAt_1(index int) {
	// 1: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout6itemAtEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:109
// index:0
// virtual
// int count()
func (this *QGraphicsGridLayout) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout5countEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:111
// index:0
// virtual
// void removeAt(int)
func (this *QGraphicsGridLayout) RemoveAt(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout8removeAtEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:112
// index:0
// void removeItem(class QGraphicsLayoutItem *)
func (this *QGraphicsGridLayout) RemoveItem(item unsafe.Pointer) {
	// 0: (, QGraphicsLayoutItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout10removeItemEP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:114
// index:0
// virtual
// void invalidate()
func (this *QGraphicsGridLayout) Invalidate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout10invalidateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:117
// index:0
// virtual
// void setGeometry(const class QRectF &)
func (this *QGraphicsGridLayout) SetGeometry(rect unsafe.Pointer) {
	// 0: (, const QRectF & rect), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsGridLayout11setGeometryERK6QRectF", ffiqt.FFI_TYPE_VOID, this.cthis, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:118
// index:0
// virtual
// QSizeF sizeHint(Qt::SizeHint, const class QSizeF &)
func (this *QGraphicsGridLayout) SizeHint(which int, constraint unsafe.Pointer) {
	// 0: (, Qt::SizeHint which, const QSizeF & constraint), (&which, constraint)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout8sizeHintEN2Qt8SizeHintERK6QSizeF", ffiqt.FFI_TYPE_VOID, this.cthis, &which, constraint)
	gopp.ErrPrint(err, rv)
}

//  body block end
