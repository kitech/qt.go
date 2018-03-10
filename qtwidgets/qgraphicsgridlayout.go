package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h
// #include <qgraphicsgridlayout.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QGraphicsGridLayout struct {
	*QGraphicsLayout
}
type QGraphicsGridLayout_ITF interface {
	QGraphicsLayout_ITF
	QGraphicsGridLayout_PTR() *QGraphicsGridLayout
}

func (ptr *QGraphicsGridLayout) QGraphicsGridLayout_PTR() *QGraphicsGridLayout { return ptr }

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
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsGridLayout(QGraphicsLayoutItem *)

/*
Constructs a QGraphicsGridLayout instance. parent is passed to QGraphicsLayout's constructor.
*/
func NewQGraphicsGridLayout(parent QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) *QGraphicsGridLayout {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = parent.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayoutC2EP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsGridLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsGridLayout)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsGridLayout(QGraphicsLayoutItem *)

/*
Constructs a QGraphicsGridLayout instance. parent is passed to QGraphicsLayout's constructor.
*/
func NewQGraphicsGridLayout__() *QGraphicsGridLayout {
	// arg: 0, QGraphicsLayoutItem *=Pointer, QGraphicsLayoutItem=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayoutC2EP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsGridLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsGridLayout)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsGridLayout()

/*

 */
func DeleteQGraphicsGridLayout(this *QGraphicsGridLayout) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayoutD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addItem(QGraphicsLayoutItem *, int, int, int, int, Qt::Alignment)

/*
Adds item to the grid on row and column. You can specify a rowSpan and columnSpan and an optional alignment.
*/
func (this *QGraphicsGridLayout) AddItem(item QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, row int, column int, rowSpan int, columnSpan int, alignment int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = item.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout7addItemEP19QGraphicsLayoutItemiiii6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, row, column, rowSpan, columnSpan, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addItem(QGraphicsLayoutItem *, int, int, int, int, Qt::Alignment)

/*
Adds item to the grid on row and column. You can specify a rowSpan and columnSpan and an optional alignment.
*/
func (this *QGraphicsGridLayout) AddItem__(item QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, row int, column int, rowSpan int, columnSpan int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = item.QGraphicsLayoutItem_PTR().GetCthis()
	}
	// arg: 5, Qt::Alignment=Elaborated, Qt::Alignment=Typedef, QFlags<Qt::AlignmentFlag>
	alignment := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout7addItemEP19QGraphicsLayoutItemiiii6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, row, column, rowSpan, columnSpan, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:61
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void addItem(QGraphicsLayoutItem *, int, int, Qt::Alignment)

/*
Adds item to the grid on row and column. You can specify a rowSpan and columnSpan and an optional alignment.
*/
func (this *QGraphicsGridLayout) AddItem_1(item QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, row int, column int, alignment int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = item.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout7addItemEP19QGraphicsLayoutItemii6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, row, column, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:61
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void addItem(QGraphicsLayoutItem *, int, int, Qt::Alignment)

/*
Adds item to the grid on row and column. You can specify a rowSpan and columnSpan and an optional alignment.
*/
func (this *QGraphicsGridLayout) AddItem_1_(item QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, row int, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = item.QGraphicsLayoutItem_PTR().GetCthis()
	}
	// arg: 3, Qt::Alignment=Elaborated, Qt::Alignment=Typedef, QFlags<Qt::AlignmentFlag>
	alignment := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout7addItemEP19QGraphicsLayoutItemii6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, row, column, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalSpacing(qreal)

/*
Sets the default horizontal spacing for the grid layout to spacing.

See also horizontalSpacing().
*/
func (this *QGraphicsGridLayout) SetHorizontalSpacing(spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout20setHorizontalSpacingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:64
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal horizontalSpacing() const

/*
Returns the default horizontal spacing for the grid layout.

See also setHorizontalSpacing().
*/
func (this *QGraphicsGridLayout) HorizontalSpacing() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout17horizontalSpacingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalSpacing(qreal)

/*
Sets the default vertical spacing for the grid layout to spacing.

See also verticalSpacing().
*/
func (this *QGraphicsGridLayout) SetVerticalSpacing(spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout18setVerticalSpacingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal verticalSpacing() const

/*
Returns the default vertical spacing for the grid layout.

See also setVerticalSpacing().
*/
func (this *QGraphicsGridLayout) VerticalSpacing() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout15verticalSpacingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpacing(qreal)

/*
Sets the grid layout's default spacing, both vertical and horizontal, to spacing.

See also rowSpacing() and columnSpacing().
*/
func (this *QGraphicsGridLayout) SetSpacing(spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout10setSpacingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowSpacing(int, qreal)

/*
Sets the spacing for row to spacing.

See also rowSpacing().
*/
func (this *QGraphicsGridLayout) SetRowSpacing(row int, spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout13setRowSpacingEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal rowSpacing(int) const

/*
Returns the row spacing for row.

See also setRowSpacing().
*/
func (this *QGraphicsGridLayout) RowSpacing(row int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout10rowSpacingEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnSpacing(int, qreal)

/*
Sets the spacing for column to spacing.

See also columnSpacing().
*/
func (this *QGraphicsGridLayout) SetColumnSpacing(column int, spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout16setColumnSpacingEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal columnSpacing(int) const

/*
Returns the column spacing for column.

See also setColumnSpacing().
*/
func (this *QGraphicsGridLayout) ColumnSpacing(column int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout13columnSpacingEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowStretchFactor(int, int)

/*
Sets the stretch factor for row to stretch.

See also rowStretchFactor().
*/
func (this *QGraphicsGridLayout) SetRowStretchFactor(row int, stretch int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout19setRowStretchFactorEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:75
// index:0
// Public Visibility=Default Availability=Available
// [4] int rowStretchFactor(int) const

/*
Returns the stretch factor for row.

See also setRowStretchFactor().
*/
func (this *QGraphicsGridLayout) RowStretchFactor(row int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout16rowStretchFactorEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnStretchFactor(int, int)

/*
Sets the stretch factor for column to stretch.

See also columnStretchFactor().
*/
func (this *QGraphicsGridLayout) SetColumnStretchFactor(column int, stretch int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout22setColumnStretchFactorEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnStretchFactor(int) const

/*
Returns the stretch factor for column.

See also setColumnStretchFactor().
*/
func (this *QGraphicsGridLayout) ColumnStretchFactor(column int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout19columnStretchFactorEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowMinimumHeight(int, qreal)

/*
Sets the minimum height for row, row, to height.

See also rowMinimumHeight().
*/
func (this *QGraphicsGridLayout) SetRowMinimumHeight(row int, height float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout19setRowMinimumHeightEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal rowMinimumHeight(int) const

/*
Returns the minimum height for row, row.

See also setRowMinimumHeight().
*/
func (this *QGraphicsGridLayout) RowMinimumHeight(row int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout16rowMinimumHeightEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowPreferredHeight(int, qreal)

/*
Sets the preferred height for row, row, to height.

See also rowPreferredHeight().
*/
func (this *QGraphicsGridLayout) SetRowPreferredHeight(row int, height float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout21setRowPreferredHeightEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal rowPreferredHeight(int) const

/*
Returns the preferred height for row, row.

See also setRowPreferredHeight().
*/
func (this *QGraphicsGridLayout) RowPreferredHeight(row int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout18rowPreferredHeightEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowMaximumHeight(int, qreal)

/*
Sets the maximum height for row, row, to height.

See also rowMaximumHeight().
*/
func (this *QGraphicsGridLayout) SetRowMaximumHeight(row int, height float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout19setRowMaximumHeightEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal rowMaximumHeight(int) const

/*
Returns the maximum height for row, row.

See also setRowMaximumHeight().
*/
func (this *QGraphicsGridLayout) RowMaximumHeight(row int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout16rowMaximumHeightEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowFixedHeight(int, qreal)

/*
Sets the fixed height for row, row, to height.
*/
func (this *QGraphicsGridLayout) SetRowFixedHeight(row int, height float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout17setRowFixedHeightEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnMinimumWidth(int, qreal)

/*
Sets the minimum width for column to width.

See also columnMinimumWidth().
*/
func (this *QGraphicsGridLayout) SetColumnMinimumWidth(column int, width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout21setColumnMinimumWidthEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal columnMinimumWidth(int) const

/*
Returns the minimum width for column.

See also setColumnMinimumWidth().
*/
func (this *QGraphicsGridLayout) ColumnMinimumWidth(column int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout18columnMinimumWidthEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnPreferredWidth(int, qreal)

/*
Sets the preferred width for column to width.

See also columnPreferredWidth().
*/
func (this *QGraphicsGridLayout) SetColumnPreferredWidth(column int, width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout23setColumnPreferredWidthEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal columnPreferredWidth(int) const

/*
Returns the preferred width for column.

See also setColumnPreferredWidth().
*/
func (this *QGraphicsGridLayout) ColumnPreferredWidth(column int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout20columnPreferredWidthEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnMaximumWidth(int, qreal)

/*
Sets the maximum width of column to width.

See also columnMaximumWidth().
*/
func (this *QGraphicsGridLayout) SetColumnMaximumWidth(column int, width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout21setColumnMaximumWidthEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal columnMaximumWidth(int) const

/*
Returns the maximum width for column.

See also setColumnMaximumWidth().
*/
func (this *QGraphicsGridLayout) ColumnMaximumWidth(column int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout18columnMaximumWidthEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnFixedWidth(int, qreal)

/*
Sets the fixed width of column to width.
*/
func (this *QGraphicsGridLayout) SetColumnFixedWidth(column int, width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout19setColumnFixedWidthEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowAlignment(int, Qt::Alignment)

/*
Sets the alignment of row to alignment.

See also rowAlignment().
*/
func (this *QGraphicsGridLayout) SetRowAlignment(row int, alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout15setRowAlignmentEi6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:96
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment rowAlignment(int) const

/*
Returns the alignment of row.

See also setRowAlignment().
*/
func (this *QGraphicsGridLayout) RowAlignment(row int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout12rowAlignmentEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnAlignment(int, Qt::Alignment)

/*
Sets the alignment for column to alignment.

See also columnAlignment().
*/
func (this *QGraphicsGridLayout) SetColumnAlignment(column int, alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout18setColumnAlignmentEi6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:98
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment columnAlignment(int) const

/*
Returns the alignment for column.

See also setColumnAlignment().
*/
func (this *QGraphicsGridLayout) ColumnAlignment(column int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout15columnAlignmentEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlignment(QGraphicsLayoutItem *, Qt::Alignment)

/*
Sets the alignment for item to alignment.

See also alignment().
*/
func (this *QGraphicsGridLayout) SetAlignment(item QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, alignment int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = item.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout12setAlignmentEP19QGraphicsLayoutItem6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment alignment(QGraphicsLayoutItem *) const

/*
Returns the alignment for item.

See also setAlignment().
*/
func (this *QGraphicsGridLayout) Alignment(item QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) int {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = item.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout9alignmentEP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:103
// index:0
// Public Visibility=Default Availability=Available
// [4] int rowCount() const

/*
Returns the number of rows in the grid layout. This is always one more than the index of the last row that is occupied by a layout item (empty rows are counted except for those at the end).
*/
func (this *QGraphicsGridLayout) RowCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout8rowCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:104
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnCount() const

/*
Returns the number of columns in the grid layout. This is always one more than the index of the last column that is occupied by a layout item (empty columns are counted except for those at the end).
*/
func (this *QGraphicsGridLayout) ColumnCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout11columnCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsLayoutItem * itemAt(int, int) const

/*
Returns a pointer to the layout item at (row, column).
*/
func (this *QGraphicsGridLayout) ItemAt(row int, column int) *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout6itemAtEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:110
// index:1
// Public virtual Visibility=Default Availability=Available
// [8] QGraphicsLayoutItem * itemAt(int) const

/*
Returns a pointer to the layout item at (row, column).
*/
func (this *QGraphicsGridLayout) ItemAt_1(index int) *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout6itemAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:109
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int count() const

/*
Reimplemented from QGraphicsLayout::count().

Returns the number of layout items in this grid layout.
*/
func (this *QGraphicsGridLayout) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:111
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void removeAt(int)

/*
Reimplemented from QGraphicsLayout::removeAt().

Removes the layout item at index without destroying it. Ownership of the item is transferred to the caller.

See also addItem().
*/
func (this *QGraphicsGridLayout) RemoveAt(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout8removeAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeItem(QGraphicsLayoutItem *)

/*
Removes the layout item item without destroying it. Ownership of the item is transferred to the caller.

See also addItem().
*/
func (this *QGraphicsGridLayout) RemoveItem(item QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = item.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout10removeItemEP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:114
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void invalidate()

/*
Reimplemented from QGraphicsLayout::invalidate().
*/
func (this *QGraphicsGridLayout) Invalidate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout10invalidateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:117
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRectF &)

/*
Reimplemented from QGraphicsLayoutItem::setGeometry().

Sets the bounding geometry of the grid layout to rect.
*/
func (this *QGraphicsGridLayout) SetGeometry(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout11setGeometryERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:118
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QSizeF sizeHint(Qt::SizeHint, const QSizeF &) const

/*
Reimplemented from QGraphicsLayoutItem::sizeHint().
*/
func (this *QGraphicsGridLayout) SizeHint(which int, constraint qtcore.QSizeF_ITF) *qtcore.QSizeF /*123*/ {
	var convArg1 unsafe.Pointer
	if constraint != nil && constraint.QSizeF_PTR() != nil {
		convArg1 = constraint.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout8sizeHintEN2Qt8SizeHintERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:118
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QSizeF sizeHint(Qt::SizeHint, const QSizeF &) const

/*
Reimplemented from QGraphicsLayoutItem::sizeHint().
*/
func (this *QGraphicsGridLayout) SizeHint__(which int) *qtcore.QSizeF /*123*/ {
	// arg: 1, const QSizeF &=LValueReference, QSizeF=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout8sizeHintEN2Qt8SizeHintERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
