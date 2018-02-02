package qtwidgets

// /usr/include/qt/QtWidgets/qformlayout.h
// #include <qformlayout.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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

type QFormLayout struct {
	*QLayout
}

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
// [8] const QMetaObject * metaObject()
func (this *QFormLayout) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qformlayout.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFormLayout(QWidget *)
func NewQFormLayout(parent *QWidget /*777 QWidget **/) *QFormLayout {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayoutC2EP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFormLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qformlayout.h:92
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFormLayout()
func DeleteQFormLayout(this *QFormLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayoutD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qformlayout.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFieldGrowthPolicy(enum QFormLayout::FieldGrowthPolicy)
func (this *QFormLayout) SetFieldGrowthPolicy(policy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout20setFieldGrowthPolicyENS_17FieldGrowthPolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:95
// index:0
// Public Visibility=Default Availability=Available
// [4] QFormLayout::FieldGrowthPolicy fieldGrowthPolicy()
func (this *QFormLayout) FieldGrowthPolicy() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout17fieldGrowthPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowWrapPolicy(enum QFormLayout::RowWrapPolicy)
func (this *QFormLayout) SetRowWrapPolicy(policy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout16setRowWrapPolicyENS_13RowWrapPolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:97
// index:0
// Public Visibility=Default Availability=Available
// [4] QFormLayout::RowWrapPolicy rowWrapPolicy()
func (this *QFormLayout) RowWrapPolicy() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout13rowWrapPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment labelAlignment()
func (this *QFormLayout) LabelAlignment() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout14labelAlignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment formAlignment()
func (this *QFormLayout) FormAlignment() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout13formAlignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalSpacing(int)
func (this *QFormLayout) SetHorizontalSpacing(spacing int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout20setHorizontalSpacingEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:104
// index:0
// Public Visibility=Default Availability=Available
// [4] int horizontalSpacing()
func (this *QFormLayout) HorizontalSpacing() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout17horizontalSpacingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qformlayout.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalSpacing(int)
func (this *QFormLayout) SetVerticalSpacing(spacing int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout18setVerticalSpacingEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:106
// index:0
// Public Visibility=Default Availability=Available
// [4] int verticalSpacing()
func (this *QFormLayout) VerticalSpacing() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout15verticalSpacingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qformlayout.h:108
// index:0
// Public Visibility=Default Availability=Available
// [4] int spacing()
func (this *QFormLayout) Spacing() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout7spacingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qformlayout.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpacing(int)
func (this *QFormLayout) SetSpacing(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout10setSpacingEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addRow(QWidget *, QWidget *)
func (this *QFormLayout) AddRow(label *QWidget /*777 QWidget **/, field *QWidget /*777 QWidget **/) {
	var convArg0 = label.GetCthis()
	var convArg1 = field.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout6addRowEP7QWidgetS1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:112
// index:1
// Public Visibility=Default Availability=Available
// [-2] void addRow(QWidget *, QLayout *)
func (this *QFormLayout) AddRow_1(label *QWidget /*777 QWidget **/, field *QLayout /*777 QLayout **/) {
	var convArg0 = label.GetCthis()
	var convArg1 = field.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout6addRowEP7QWidgetP7QLayout", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:113
// index:2
// Public Visibility=Default Availability=Available
// [-2] void addRow(const QString &, QWidget *)
func (this *QFormLayout) AddRow_2(labelText *qtcore.QString, field *QWidget /*777 QWidget **/) {
	var convArg0 = labelText.GetCthis()
	var convArg1 = field.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout6addRowERK7QStringP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:114
// index:3
// Public Visibility=Default Availability=Available
// [-2] void addRow(const QString &, QLayout *)
func (this *QFormLayout) AddRow_3(labelText *qtcore.QString, field *QLayout /*777 QLayout **/) {
	var convArg0 = labelText.GetCthis()
	var convArg1 = field.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout6addRowERK7QStringP7QLayout", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:115
// index:4
// Public Visibility=Default Availability=Available
// [-2] void addRow(QWidget *)
func (this *QFormLayout) AddRow_4(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout6addRowEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:116
// index:5
// Public Visibility=Default Availability=Available
// [-2] void addRow(QLayout *)
func (this *QFormLayout) AddRow_5(layout *QLayout /*777 QLayout **/) {
	var convArg0 = layout.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout6addRowEP7QLayout", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertRow(int, QWidget *, QWidget *)
func (this *QFormLayout) InsertRow(row int, label *QWidget /*777 QWidget **/, field *QWidget /*777 QWidget **/) {
	var convArg1 = label.GetCthis()
	var convArg2 = field.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiP7QWidgetS1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:119
// index:1
// Public Visibility=Default Availability=Available
// [-2] void insertRow(int, QWidget *, QLayout *)
func (this *QFormLayout) InsertRow_1(row int, label *QWidget /*777 QWidget **/, field *QLayout /*777 QLayout **/) {
	var convArg1 = label.GetCthis()
	var convArg2 = field.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiP7QWidgetP7QLayout", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:120
// index:2
// Public Visibility=Default Availability=Available
// [-2] void insertRow(int, const QString &, QWidget *)
func (this *QFormLayout) InsertRow_2(row int, labelText *qtcore.QString, field *QWidget /*777 QWidget **/) {
	var convArg1 = labelText.GetCthis()
	var convArg2 = field.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiRK7QStringP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:121
// index:3
// Public Visibility=Default Availability=Available
// [-2] void insertRow(int, const QString &, QLayout *)
func (this *QFormLayout) InsertRow_3(row int, labelText *qtcore.QString, field *QLayout /*777 QLayout **/) {
	var convArg1 = labelText.GetCthis()
	var convArg2 = field.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiRK7QStringP7QLayout", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:122
// index:4
// Public Visibility=Default Availability=Available
// [-2] void insertRow(int, QWidget *)
func (this *QFormLayout) InsertRow_4(row int, widget *QWidget /*777 QWidget **/) {
	var convArg1 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:123
// index:5
// Public Visibility=Default Availability=Available
// [-2] void insertRow(int, QLayout *)
func (this *QFormLayout) InsertRow_5(row int, layout *QLayout /*777 QLayout **/) {
	var convArg1 = layout.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiP7QLayout", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeRow(int)
func (this *QFormLayout) RemoveRow(row int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout9removeRowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:126
// index:1
// Public Visibility=Default Availability=Available
// [-2] void removeRow(QWidget *)
func (this *QFormLayout) RemoveRow_1(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout9removeRowEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:127
// index:2
// Public Visibility=Default Availability=Available
// [-2] void removeRow(QLayout *)
func (this *QFormLayout) RemoveRow_2(layout *QLayout /*777 QLayout **/) {
	var convArg0 = layout.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout9removeRowEP7QLayout", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:129
// index:0
// Public Visibility=Default Availability=Available
// [16] QFormLayout::TakeRowResult takeRow(int)
func (this *QFormLayout) TakeRow(row int) unsafe.Pointer /*444*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout7takeRowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtWidgets/qformlayout.h:130
// index:1
// Public Visibility=Default Availability=Available
// [16] QFormLayout::TakeRowResult takeRow(QWidget *)
func (this *QFormLayout) TakeRow_1(widget *QWidget /*777 QWidget **/) unsafe.Pointer /*444*/ {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout7takeRowEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtWidgets/qformlayout.h:131
// index:2
// Public Visibility=Default Availability=Available
// [16] QFormLayout::TakeRowResult takeRow(QLayout *)
func (this *QFormLayout) TakeRow_2(layout *QLayout /*777 QLayout **/) unsafe.Pointer /*444*/ {
	var convArg0 = layout.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout7takeRowEP7QLayout", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtWidgets/qformlayout.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItem(int, enum QFormLayout::ItemRole, QLayoutItem *)
func (this *QFormLayout) SetItem(row int, role int, item *QLayoutItem /*777 QLayoutItem **/) {
	var convArg2 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout7setItemEiNS_8ItemRoleEP11QLayoutItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, role, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidget(int, enum QFormLayout::ItemRole, QWidget *)
func (this *QFormLayout) SetWidget(row int, role int, widget *QWidget /*777 QWidget **/) {
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout9setWidgetEiNS_8ItemRoleEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, role, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLayout(int, enum QFormLayout::ItemRole, QLayout *)
func (this *QFormLayout) SetLayout(row int, role int, layout *QLayout /*777 QLayout **/) {
	var convArg2 = layout.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout9setLayoutEiNS_8ItemRoleEP7QLayout", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, role, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:137
// index:0
// Public Visibility=Default Availability=Available
// [8] QLayoutItem * itemAt(int, enum QFormLayout::ItemRole)
func (this *QFormLayout) ItemAt(row int, role int) *QLayoutItem /*777 QLayoutItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout6itemAtEiNS_8ItemRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qformlayout.h:146
// index:1
// Public virtual Visibility=Default Availability=Available
// [8] QLayoutItem * itemAt(int)
func (this *QFormLayout) ItemAt_1(index int) *QLayoutItem /*777 QLayoutItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout6itemAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qformlayout.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getItemPosition(int, int *, enum QFormLayout::ItemRole *)
func (this *QFormLayout) GetItemPosition(index int, rowPtr unsafe.Pointer /*666*/, rolePtr unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout15getItemPositionEiPiPNS_8ItemRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, &rowPtr, &rolePtr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getWidgetPosition(QWidget *, int *, enum QFormLayout::ItemRole *)
func (this *QFormLayout) GetWidgetPosition(widget *QWidget /*777 QWidget **/, rowPtr unsafe.Pointer /*666*/, rolePtr unsafe.Pointer /*666*/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout17getWidgetPositionEP7QWidgetPiPNS_8ItemRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &rowPtr, &rolePtr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getLayoutPosition(QLayout *, int *, enum QFormLayout::ItemRole *)
func (this *QFormLayout) GetLayoutPosition(layout *QLayout /*777 QLayout **/, rowPtr unsafe.Pointer /*666*/, rolePtr unsafe.Pointer /*666*/) {
	var convArg0 = layout.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout17getLayoutPositionEP7QLayoutPiPNS_8ItemRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &rowPtr, &rolePtr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:141
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * labelForField(QWidget *)
func (this *QFormLayout) LabelForField(field *QWidget /*777 QWidget **/) *QWidget /*777 QWidget **/ {
	var convArg0 = field.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout13labelForFieldEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qformlayout.h:142
// index:1
// Public Visibility=Default Availability=Available
// [8] QWidget * labelForField(QLayout *)
func (this *QFormLayout) LabelForField_1(field *QLayout /*777 QLayout **/) *QWidget /*777 QWidget **/ {
	var convArg0 = field.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout13labelForFieldEP7QLayout", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qformlayout.h:145
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void addItem(QLayoutItem *)
func (this *QFormLayout) AddItem(item *QLayoutItem /*777 QLayoutItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout7addItemEP11QLayoutItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:147
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QLayoutItem * takeAt(int)
func (this *QFormLayout) TakeAt(index int) *QLayoutItem /*777 QLayoutItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout6takeAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qformlayout.h:149
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRect &)
func (this *QFormLayout) SetGeometry(rect *qtcore.QRect) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout11setGeometryERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:150
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSize()
func (this *QFormLayout) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout11minimumSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qformlayout.h:151
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QFormLayout) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qformlayout.h:152
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void invalidate()
func (this *QFormLayout) Invalidate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout10invalidateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:154
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasHeightForWidth()
func (this *QFormLayout) HasHeightForWidth() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout17hasHeightForWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qformlayout.h:155
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int heightForWidth(int)
func (this *QFormLayout) HeightForWidth(width int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout14heightForWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), width)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qformlayout.h:156
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::Orientations expandingDirections()
func (this *QFormLayout) ExpandingDirections() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout19expandingDirectionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:157
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int count()
func (this *QFormLayout) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qformlayout.h:159
// index:0
// Public Visibility=Default Availability=Available
// [4] int rowCount()
func (this *QFormLayout) RowCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout8rowCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

type QFormLayout__FieldGrowthPolicy = int

const QFormLayout__FieldsStayAtSizeHint QFormLayout__FieldGrowthPolicy = 0
const QFormLayout__ExpandingFieldsGrow QFormLayout__FieldGrowthPolicy = 1
const QFormLayout__AllNonFixedFieldsGrow QFormLayout__FieldGrowthPolicy = 2

type QFormLayout__RowWrapPolicy = int

const QFormLayout__DontWrapRows QFormLayout__RowWrapPolicy = 0
const QFormLayout__WrapLongRows QFormLayout__RowWrapPolicy = 1
const QFormLayout__WrapAllRows QFormLayout__RowWrapPolicy = 2

type QFormLayout__ItemRole = int

const QFormLayout__LabelRole QFormLayout__ItemRole = 0
const QFormLayout__FieldRole QFormLayout__ItemRole = 1
const QFormLayout__SpanningRole QFormLayout__ItemRole = 2

//  body block end
