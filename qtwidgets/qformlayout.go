// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qformlayout.h
// #include <qformlayout.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
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
type QFormLayout struct {
	*QLayout
}
type QFormLayout_ITF interface {
	QLayout_ITF
	QFormLayout_PTR() *QFormLayout
}

func (ptr *QFormLayout) QFormLayout_PTR() *QFormLayout { return ptr }

func (this *QFormLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QLayout.GetCthis()
	}
}
func (this *QFormLayout) SetCthis(cthis unsafe.Pointer) {
	this.QLayout = NewQLayoutFromPointer(cthis)
}
func NewQFormLayoutFromPointer(cthis unsafe.Pointer) *QFormLayout {
	bcthis0 := NewQLayoutFromPointer(cthis)
	return &QFormLayout{bcthis0}
}
func (*QFormLayout) NewFromPointer(cthis unsafe.Pointer) *QFormLayout {
	return NewQFormLayoutFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qformlayout.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QFormLayout) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qformlayout.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFormLayout(QWidget *)

/*
Constructs a new form layout with the given parent widget.

See also QWidget::setLayout().
*/
func (*QFormLayout) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QFormLayout {
	return NewQFormLayout(parent)
}
func NewQFormLayout(parent QWidget_ITF /*777 QWidget **/) *QFormLayout {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayoutC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFormLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFormLayout")
	return gothis
}

// /usr/include/qt/QtWidgets/qformlayout.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFormLayout(QWidget *)

/*
Constructs a new form layout with the given parent widget.

See also QWidget::setLayout().
*/
func (*QFormLayout) NewForInheritp() *QFormLayout {
	return NewQFormLayoutp()
}
func NewQFormLayoutp() *QFormLayout {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayoutC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFormLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFormLayout")
	return gothis
}

// /usr/include/qt/QtWidgets/qformlayout.h:92
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFormLayout()

/*

 */
func DeleteQFormLayout(this *QFormLayout) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayoutD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qformlayout.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFieldGrowthPolicy(QFormLayout::FieldGrowthPolicy)

/*

 */
func (this *QFormLayout) SetFieldGrowthPolicy(policy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout20setFieldGrowthPolicyENS_17FieldGrowthPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:95
// index:0
// Public Visibility=Default Availability=Available
// [4] QFormLayout::FieldGrowthPolicy fieldGrowthPolicy() const

/*

 */
func (this *QFormLayout) FieldGrowthPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout17fieldGrowthPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowWrapPolicy(QFormLayout::RowWrapPolicy)

/*

 */
func (this *QFormLayout) SetRowWrapPolicy(policy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout16setRowWrapPolicyENS_13RowWrapPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:97
// index:0
// Public Visibility=Default Availability=Available
// [4] QFormLayout::RowWrapPolicy rowWrapPolicy() const

/*

 */
func (this *QFormLayout) RowWrapPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout13rowWrapPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLabelAlignment(Qt::Alignment)

/*

 */
func (this *QFormLayout) SetLabelAlignment(alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout17setLabelAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment labelAlignment() const

/*

 */
func (this *QFormLayout) LabelAlignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout14labelAlignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFormAlignment(Qt::Alignment)

/*

 */
func (this *QFormLayout) SetFormAlignment(alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout16setFormAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment formAlignment() const

/*

 */
func (this *QFormLayout) FormAlignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout13formAlignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalSpacing(int)

/*

 */
func (this *QFormLayout) SetHorizontalSpacing(spacing int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout20setHorizontalSpacingEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:104
// index:0
// Public Visibility=Default Availability=Available
// [4] int horizontalSpacing() const

/*

 */
func (this *QFormLayout) HorizontalSpacing() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout17horizontalSpacingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qformlayout.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalSpacing(int)

/*

 */
func (this *QFormLayout) SetVerticalSpacing(spacing int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout18setVerticalSpacingEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:106
// index:0
// Public Visibility=Default Availability=Available
// [4] int verticalSpacing() const

/*

 */
func (this *QFormLayout) VerticalSpacing() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout15verticalSpacingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qformlayout.h:108
// index:0
// Public Visibility=Default Availability=Available
// [4] int spacing() const

/*
If the vertical spacing is equal to the horizontal spacing, this function returns that value; otherwise it returns -1.

See also setSpacing(), verticalSpacing(), and horizontalSpacing().
*/
func (this *QFormLayout) Spacing() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout7spacingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qformlayout.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpacing(int)

/*
This function sets both the vertical and horizontal spacing to spacing.

See also spacing(), setVerticalSpacing(), and setHorizontalSpacing().
*/
func (this *QFormLayout) SetSpacing(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout10setSpacingEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addRow(QWidget *, QWidget *)

/*
Adds a new row to the bottom of this form layout, with the given label and field.

See also insertRow().
*/
func (this *QFormLayout) AddRow(label QWidget_ITF /*777 QWidget **/, field QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if label != nil && label.QWidget_PTR() != nil {
		convArg0 = label.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if field != nil && field.QWidget_PTR() != nil {
		convArg1 = field.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout6addRowEP7QWidgetS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:112
// index:1
// Public Visibility=Default Availability=Available
// [-2] void addRow(QWidget *, QLayout *)

/*
Adds a new row to the bottom of this form layout, with the given label and field.

See also insertRow().
*/
func (this *QFormLayout) AddRow1(label QWidget_ITF /*777 QWidget **/, field QLayout_ITF /*777 QLayout **/) {
	var convArg0 unsafe.Pointer
	if label != nil && label.QWidget_PTR() != nil {
		convArg0 = label.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if field != nil && field.QLayout_PTR() != nil {
		convArg1 = field.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout6addRowEP7QWidgetP7QLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:113
// index:2
// Public Visibility=Default Availability=Available
// [-2] void addRow(const QString &, QWidget *)

/*
Adds a new row to the bottom of this form layout, with the given label and field.

See also insertRow().
*/
func (this *QFormLayout) AddRow2(labelText string, field QWidget_ITF /*777 QWidget **/) {
	var tmpArg0 = qtcore.NewQString5(labelText)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if field != nil && field.QWidget_PTR() != nil {
		convArg1 = field.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout6addRowERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:114
// index:3
// Public Visibility=Default Availability=Available
// [-2] void addRow(const QString &, QLayout *)

/*
Adds a new row to the bottom of this form layout, with the given label and field.

See also insertRow().
*/
func (this *QFormLayout) AddRow3(labelText string, field QLayout_ITF /*777 QLayout **/) {
	var tmpArg0 = qtcore.NewQString5(labelText)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if field != nil && field.QLayout_PTR() != nil {
		convArg1 = field.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout6addRowERK7QStringP7QLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:115
// index:4
// Public Visibility=Default Availability=Available
// [-2] void addRow(QWidget *)

/*
Adds a new row to the bottom of this form layout, with the given label and field.

See also insertRow().
*/
func (this *QFormLayout) AddRow4(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout6addRowEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:116
// index:5
// Public Visibility=Default Availability=Available
// [-2] void addRow(QLayout *)

/*
Adds a new row to the bottom of this form layout, with the given label and field.

See also insertRow().
*/
func (this *QFormLayout) AddRow5(layout QLayout_ITF /*777 QLayout **/) {
	var convArg0 unsafe.Pointer
	if layout != nil && layout.QLayout_PTR() != nil {
		convArg0 = layout.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout6addRowEP7QLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertRow(int, QWidget *, QWidget *)

/*
Inserts a new row at position row in this form layout, with the given label and field. If row is out of bounds, the new row is added at the end.

See also addRow().
*/
func (this *QFormLayout) InsertRow(row int, label QWidget_ITF /*777 QWidget **/, field QWidget_ITF /*777 QWidget **/) {
	var convArg1 unsafe.Pointer
	if label != nil && label.QWidget_PTR() != nil {
		convArg1 = label.QWidget_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if field != nil && field.QWidget_PTR() != nil {
		convArg2 = field.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiP7QWidgetS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:119
// index:1
// Public Visibility=Default Availability=Available
// [-2] void insertRow(int, QWidget *, QLayout *)

/*
Inserts a new row at position row in this form layout, with the given label and field. If row is out of bounds, the new row is added at the end.

See also addRow().
*/
func (this *QFormLayout) InsertRow1(row int, label QWidget_ITF /*777 QWidget **/, field QLayout_ITF /*777 QLayout **/) {
	var convArg1 unsafe.Pointer
	if label != nil && label.QWidget_PTR() != nil {
		convArg1 = label.QWidget_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if field != nil && field.QLayout_PTR() != nil {
		convArg2 = field.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiP7QWidgetP7QLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:120
// index:2
// Public Visibility=Default Availability=Available
// [-2] void insertRow(int, const QString &, QWidget *)

/*
Inserts a new row at position row in this form layout, with the given label and field. If row is out of bounds, the new row is added at the end.

See also addRow().
*/
func (this *QFormLayout) InsertRow2(row int, labelText string, field QWidget_ITF /*777 QWidget **/) {
	var tmpArg1 = qtcore.NewQString5(labelText)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if field != nil && field.QWidget_PTR() != nil {
		convArg2 = field.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiRK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:121
// index:3
// Public Visibility=Default Availability=Available
// [-2] void insertRow(int, const QString &, QLayout *)

/*
Inserts a new row at position row in this form layout, with the given label and field. If row is out of bounds, the new row is added at the end.

See also addRow().
*/
func (this *QFormLayout) InsertRow3(row int, labelText string, field QLayout_ITF /*777 QLayout **/) {
	var tmpArg1 = qtcore.NewQString5(labelText)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if field != nil && field.QLayout_PTR() != nil {
		convArg2 = field.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiRK7QStringP7QLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:122
// index:4
// Public Visibility=Default Availability=Available
// [-2] void insertRow(int, QWidget *)

/*
Inserts a new row at position row in this form layout, with the given label and field. If row is out of bounds, the new row is added at the end.

See also addRow().
*/
func (this *QFormLayout) InsertRow4(row int, widget QWidget_ITF /*777 QWidget **/) {
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:123
// index:5
// Public Visibility=Default Availability=Available
// [-2] void insertRow(int, QLayout *)

/*
Inserts a new row at position row in this form layout, with the given label and field. If row is out of bounds, the new row is added at the end.

See also addRow().
*/
func (this *QFormLayout) InsertRow5(row int, layout QLayout_ITF /*777 QLayout **/) {
	var convArg1 unsafe.Pointer
	if layout != nil && layout.QLayout_PTR() != nil {
		convArg1 = layout.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiP7QLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeRow(int)

/*
Deletes row row from this form layout.

row must be non-negative and less than rowCount().

After this call, rowCount() is decremented by one. All widgets and nested layouts that occupied this row are deleted. That includes both the field widget(s) and the label, if any. All following rows are shifted up one row and the freed vertical space is redistributed amongst the remaining rows.

You can use this function to undo a previous addRow() or insertRow():


  QFormLayout *flay = ...;
  QPointer<QLineEdit> le = new QLineEdit;
  flay->insertRow(2, "User:", le);
  // later:
  flay->removeRow(2); // le == nullptr at this point



If you want to remove the row from the layout without deleting the widgets, use takeRow() instead.

This function was introduced in  Qt 5.8.

See also takeRow().
*/
func (this *QFormLayout) RemoveRow(row int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout9removeRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:126
// index:1
// Public Visibility=Default Availability=Available
// [-2] void removeRow(QWidget *)

/*
Deletes row row from this form layout.

row must be non-negative and less than rowCount().

After this call, rowCount() is decremented by one. All widgets and nested layouts that occupied this row are deleted. That includes both the field widget(s) and the label, if any. All following rows are shifted up one row and the freed vertical space is redistributed amongst the remaining rows.

You can use this function to undo a previous addRow() or insertRow():


  QFormLayout *flay = ...;
  QPointer<QLineEdit> le = new QLineEdit;
  flay->insertRow(2, "User:", le);
  // later:
  flay->removeRow(2); // le == nullptr at this point



If you want to remove the row from the layout without deleting the widgets, use takeRow() instead.

This function was introduced in  Qt 5.8.

See also takeRow().
*/
func (this *QFormLayout) RemoveRow1(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout9removeRowEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:127
// index:2
// Public Visibility=Default Availability=Available
// [-2] void removeRow(QLayout *)

/*
Deletes row row from this form layout.

row must be non-negative and less than rowCount().

After this call, rowCount() is decremented by one. All widgets and nested layouts that occupied this row are deleted. That includes both the field widget(s) and the label, if any. All following rows are shifted up one row and the freed vertical space is redistributed amongst the remaining rows.

You can use this function to undo a previous addRow() or insertRow():


  QFormLayout *flay = ...;
  QPointer<QLineEdit> le = new QLineEdit;
  flay->insertRow(2, "User:", le);
  // later:
  flay->removeRow(2); // le == nullptr at this point



If you want to remove the row from the layout without deleting the widgets, use takeRow() instead.

This function was introduced in  Qt 5.8.

See also takeRow().
*/
func (this *QFormLayout) RemoveRow2(layout QLayout_ITF /*777 QLayout **/) {
	var convArg0 unsafe.Pointer
	if layout != nil && layout.QLayout_PTR() != nil {
		convArg0 = layout.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout9removeRowEP7QLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:129
// index:0
// Public Visibility=Default Availability=Available
// [16] QFormLayout::TakeRowResult takeRow(int)

/*
Removes the specified row from this form layout.

row must be non-negative and less than rowCount().

Note: This function doesn't delete anything.

After this call, rowCount() is decremented by one. All following rows are shifted up one row and the freed vertical space is redistributed amongst the remaining rows.

You can use this function to undo a previous addRow() or insertRow():


  QFormLayout *flay = ...;
  QPointer<QLineEdit> le = new QLineEdit;
  flay->insertRow(2, "User:", le);
  // later:
  QFormLayout::TakeRowResult result = flay->takeRow(2);



If you want to remove the row from the layout and delete the widgets, use removeRow() instead.

Returns A structure containing both the widget and corresponding label layout items

This function was introduced in  Qt 5.8.

See also removeRow().
*/
func (this *QFormLayout) TakeRow(row int) unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout7takeRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtWidgets/qformlayout.h:130
// index:1
// Public Visibility=Default Availability=Available
// [16] QFormLayout::TakeRowResult takeRow(QWidget *)

/*
Removes the specified row from this form layout.

row must be non-negative and less than rowCount().

Note: This function doesn't delete anything.

After this call, rowCount() is decremented by one. All following rows are shifted up one row and the freed vertical space is redistributed amongst the remaining rows.

You can use this function to undo a previous addRow() or insertRow():


  QFormLayout *flay = ...;
  QPointer<QLineEdit> le = new QLineEdit;
  flay->insertRow(2, "User:", le);
  // later:
  QFormLayout::TakeRowResult result = flay->takeRow(2);



If you want to remove the row from the layout and delete the widgets, use removeRow() instead.

Returns A structure containing both the widget and corresponding label layout items

This function was introduced in  Qt 5.8.

See also removeRow().
*/
func (this *QFormLayout) TakeRow1(widget QWidget_ITF /*777 QWidget **/) unsafe.Pointer /*444*/ {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout7takeRowEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtWidgets/qformlayout.h:131
// index:2
// Public Visibility=Default Availability=Available
// [16] QFormLayout::TakeRowResult takeRow(QLayout *)

/*
Removes the specified row from this form layout.

row must be non-negative and less than rowCount().

Note: This function doesn't delete anything.

After this call, rowCount() is decremented by one. All following rows are shifted up one row and the freed vertical space is redistributed amongst the remaining rows.

You can use this function to undo a previous addRow() or insertRow():


  QFormLayout *flay = ...;
  QPointer<QLineEdit> le = new QLineEdit;
  flay->insertRow(2, "User:", le);
  // later:
  QFormLayout::TakeRowResult result = flay->takeRow(2);



If you want to remove the row from the layout and delete the widgets, use removeRow() instead.

Returns A structure containing both the widget and corresponding label layout items

This function was introduced in  Qt 5.8.

See also removeRow().
*/
func (this *QFormLayout) TakeRow2(layout QLayout_ITF /*777 QLayout **/) unsafe.Pointer /*444*/ {
	var convArg0 unsafe.Pointer
	if layout != nil && layout.QLayout_PTR() != nil {
		convArg0 = layout.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout7takeRowEP7QLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtWidgets/qformlayout.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItem(int, QFormLayout::ItemRole, QLayoutItem *)

/*
Sets the item in the given row for the given role to item, extending the layout with empty rows if necessary.

If the cell is already occupied, the item is not inserted and an error message is sent to the console. The item spans both columns.

Warning: Do not use this function to add child layouts or child widget items. Use setLayout() or setWidget() instead.

See also setLayout().
*/
func (this *QFormLayout) SetItem(row int, role int, item QLayoutItem_ITF /*777 QLayoutItem **/) {
	var convArg2 unsafe.Pointer
	if item != nil && item.QLayoutItem_PTR() != nil {
		convArg2 = item.QLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout7setItemEiNS_8ItemRoleEP11QLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, role, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidget(int, QFormLayout::ItemRole, QWidget *)

/*
Sets the widget in the given row for the given role to widget, extending the layout with empty rows if necessary.

If the cell is already occupied, the widget is not inserted and an error message is sent to the console.

Note: For most applications, addRow() or insertRow() should be used instead of setWidget().

See also setLayout().
*/
func (this *QFormLayout) SetWidget(row int, role int, widget QWidget_ITF /*777 QWidget **/) {
	var convArg2 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg2 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout9setWidgetEiNS_8ItemRoleEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, role, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLayout(int, QFormLayout::ItemRole, QLayout *)

/*
Sets the sub-layout in the given row for the given role to layout, extending the form layout with empty rows if necessary.

If the cell is already occupied, the layout is not inserted and an error message is sent to the console.

Note: For most applications, addRow() or insertRow() should be used instead of setLayout().

See also setWidget().
*/
func (this *QFormLayout) SetLayout(row int, role int, layout QLayout_ITF /*777 QLayout **/) {
	var convArg2 unsafe.Pointer
	if layout != nil && layout.QLayout_PTR() != nil {
		convArg2 = layout.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout9setLayoutEiNS_8ItemRoleEP7QLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, role, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:137
// index:0
// Public Visibility=Default Availability=Available
// [8] QLayoutItem * itemAt(int, QFormLayout::ItemRole) const

/*
Returns the layout item in the given row with the specified role (column). Returns 0 if there is no such item.

See also QLayout::itemAt() and setItem().
*/
func (this *QFormLayout) ItemAt(row int, role int) *QLayoutItem /*777 QLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout6itemAtEiNS_8ItemRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, role)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qformlayout.h:146
// index:1
// Public virtual Visibility=Default Availability=Available
// [8] QLayoutItem * itemAt(int) const

/*
Returns the layout item in the given row with the specified role (column). Returns 0 if there is no such item.

See also QLayout::itemAt() and setItem().
*/
func (this *QFormLayout) ItemAt1(index int) *QLayoutItem /*777 QLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout6itemAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qformlayout.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getItemPosition(int, int *, QFormLayout::ItemRole *) const

/*
Retrieves the row and role (column) of the item at the specified index. If index is out of bounds, *rowPtr is set to -1; otherwise the row is stored in *rowPtr and the role is stored in *rolePtr.

See also itemAt(), count(), getLayoutPosition(), and getWidgetPosition().
*/
func (this *QFormLayout) GetItemPosition(index int, rowPtr unsafe.Pointer /*666*/, rolePtr unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout15getItemPositionEiPiPNS_8ItemRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, rowPtr, rolePtr)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getWidgetPosition(QWidget *, int *, QFormLayout::ItemRole *) const

/*
Retrieves the row and role (column) of the specified widget in the layout. If widget is not in the layout, *rowPtr is set to -1; otherwise the row is stored in *rowPtr and the role is stored in *rolePtr.

See also getItemPosition() and itemAt().
*/
func (this *QFormLayout) GetWidgetPosition(widget QWidget_ITF /*777 QWidget **/, rowPtr unsafe.Pointer /*666*/, rolePtr unsafe.Pointer /*666*/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout17getWidgetPositionEP7QWidgetPiPNS_8ItemRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, rowPtr, rolePtr)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getLayoutPosition(QLayout *, int *, QFormLayout::ItemRole *) const

/*
Retrieves the row and role (column) of the specified child layout. If layout is not in the form layout, *rowPtr is set to -1; otherwise the row is stored in *rowPtr and the role is stored in *rolePtr.
*/
func (this *QFormLayout) GetLayoutPosition(layout QLayout_ITF /*777 QLayout **/, rowPtr unsafe.Pointer /*666*/, rolePtr unsafe.Pointer /*666*/) {
	var convArg0 unsafe.Pointer
	if layout != nil && layout.QLayout_PTR() != nil {
		convArg0 = layout.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout17getLayoutPositionEP7QLayoutPiPNS_8ItemRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, rowPtr, rolePtr)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:141
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * labelForField(QWidget *) const

/*
Returns the label associated with the given field.

See also itemAt().
*/
func (this *QFormLayout) LabelForField(field QWidget_ITF /*777 QWidget **/) *QWidget /*777 QWidget **/ {
	var convArg0 unsafe.Pointer
	if field != nil && field.QWidget_PTR() != nil {
		convArg0 = field.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout13labelForFieldEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qformlayout.h:142
// index:1
// Public Visibility=Default Availability=Available
// [8] QWidget * labelForField(QLayout *) const

/*
Returns the label associated with the given field.

See also itemAt().
*/
func (this *QFormLayout) LabelForField1(field QLayout_ITF /*777 QLayout **/) *QWidget /*777 QWidget **/ {
	var convArg0 unsafe.Pointer
	if field != nil && field.QLayout_PTR() != nil {
		convArg0 = field.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout13labelForFieldEP7QLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qformlayout.h:145
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void addItem(QLayoutItem *)

/*
Reimplemented from QLayout::addItem().
*/
func (this *QFormLayout) AddItem(item QLayoutItem_ITF /*777 QLayoutItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QLayoutItem_PTR() != nil {
		convArg0 = item.QLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout7addItemEP11QLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:147
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QLayoutItem * takeAt(int)

/*
Reimplemented from QLayout::takeAt().
*/
func (this *QFormLayout) TakeAt(index int) *QLayoutItem /*777 QLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout6takeAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qformlayout.h:149
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRect &)

/*
Reimplemented from QLayoutItem::setGeometry().
*/
func (this *QFormLayout) SetGeometry(rect qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout11setGeometryERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:150
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSize() const

/*
Reimplemented from QLayoutItem::minimumSize().
*/
func (this *QFormLayout) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout11minimumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qformlayout.h:151
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QLayoutItem::sizeHint().
*/
func (this *QFormLayout) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qformlayout.h:152
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void invalidate()

/*
Reimplemented from QLayoutItem::invalidate().
*/
func (this *QFormLayout) Invalidate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFormLayout10invalidateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:154
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasHeightForWidth() const

/*
Reimplemented from QLayoutItem::hasHeightForWidth().
*/
func (this *QFormLayout) HasHeightForWidth() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout17hasHeightForWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qformlayout.h:155
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int heightForWidth(int) const

/*
Reimplemented from QLayoutItem::heightForWidth().
*/
func (this *QFormLayout) HeightForWidth(width int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout14heightForWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qformlayout.h:156
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::Orientations expandingDirections() const

/*
Reimplemented from QLayoutItem::expandingDirections().
*/
func (this *QFormLayout) ExpandingDirections() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout19expandingDirectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:157
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int count() const

/*
Reimplemented from QLayout::count().
*/
func (this *QFormLayout) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qformlayout.h:159
// index:0
// Public Visibility=Default Availability=Available
// [4] int rowCount() const

/*
Returns the number of rows in the form.

See also QLayout::count().
*/
func (this *QFormLayout) RowCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFormLayout8rowCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

/*
This enum specifies the different policies that can be used to control the way in which the form's fields grow.



See also fieldGrowthPolicy.

*/
type QFormLayout__FieldGrowthPolicy = int

// The fields never grow beyond their effective size hint. This is the default for QMacStyle.
const QFormLayout__FieldsStayAtSizeHint QFormLayout__FieldGrowthPolicy = 0

// Fields with an horizontal size policy of Expanding or MinimumExpanding will grow to fill the available space. The other fields will not grow beyond their effective size hint. This is the default policy for Plastique.
const QFormLayout__ExpandingFieldsGrow QFormLayout__FieldGrowthPolicy = 1

// All fields with a size policy that allows them to grow will grow to fill the available space. This is the default policy for most styles.
const QFormLayout__AllNonFixedFieldsGrow QFormLayout__FieldGrowthPolicy = 2

func (this *QFormLayout) FieldGrowthPolicyItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QFormLayout_FieldGrowthPolicyItemName(val int) string {
	var nilthis *QFormLayout
	return nilthis.FieldGrowthPolicyItemName(val)
}

/*
This enum specifies the different policies that can be used to control the way in which the form's rows wrap.



See also rowWrapPolicy.

*/
type QFormLayout__RowWrapPolicy = int

// Fields are always laid out next to their label. This is the default policy for all styles except Qt Extended styles.
const QFormLayout__DontWrapRows QFormLayout__RowWrapPolicy = 0

// Labels are given enough horizontal space to fit the widest label, and the rest of the space is given to the fields. If the minimum size of a field pair is wider than the available space, the field is wrapped to the next line. This is the default policy for Qt Extended styles.
const QFormLayout__WrapLongRows QFormLayout__RowWrapPolicy = 1

// Fields are always laid out below their label.
const QFormLayout__WrapAllRows QFormLayout__RowWrapPolicy = 2

func (this *QFormLayout) RowWrapPolicyItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QFormLayout_RowWrapPolicyItemName(val int) string {
	var nilthis *QFormLayout
	return nilthis.RowWrapPolicyItemName(val)
}

/*
This enum specifies the types of widgets (or other layout items) that may appear in a row.



See also itemAt() and getItemPosition().

*/
type QFormLayout__ItemRole = int

// A label widget.
const QFormLayout__LabelRole QFormLayout__ItemRole = 0

// A field widget.
const QFormLayout__FieldRole QFormLayout__ItemRole = 1

// A widget that spans label and field columns.
const QFormLayout__SpanningRole QFormLayout__ItemRole = 2

func (this *QFormLayout) ItemRoleItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QFormLayout_ItemRoleItemName(val int) string {
	var nilthis *QFormLayout
	return nilthis.ItemRoleItemName(val)
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
