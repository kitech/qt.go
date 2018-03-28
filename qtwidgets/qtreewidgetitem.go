package qtwidgets

// /usr/include/qt/QtWidgets/qtreewidget.h
// #include <qtreewidget.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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

// void emitDataChanged()
func (this *QTreeWidgetItem) InheritEmitDataChanged(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "emitDataChanged", f)
}

/*

 */
type QTreeWidgetItem struct {
	*qtrt.CObject
}
type QTreeWidgetItem_ITF interface {
	QTreeWidgetItem_PTR() *QTreeWidgetItem
}

func (ptr *QTreeWidgetItem) QTreeWidgetItem_PTR() *QTreeWidgetItem { return ptr }

func (this *QTreeWidgetItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTreeWidgetItem) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTreeWidgetItemFromPointer(cthis unsafe.Pointer) *QTreeWidgetItem {
	return &QTreeWidgetItem{&qtrt.CObject{cthis}}
}
func (*QTreeWidgetItem) NewFromPointer(cthis unsafe.Pointer) *QTreeWidgetItem {
	return NewQTreeWidgetItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidgetItem(int)

/*

 */
func NewQTreeWidgetItem(type_ int) *QTreeWidgetItem {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2Ei", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTreeWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidgetItem(int)

/*

 */
func NewQTreeWidgetItem__() *QTreeWidgetItem {
	// arg: 0, int=Int, =Invalid,
	type_ := 0 /*Type*/
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2Ei", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTreeWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:68
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidgetItem(const QStringList &, int)

/*

 */
func NewQTreeWidgetItem_1(strings qtcore.QStringList_ITF, type_ int) *QTreeWidgetItem {
	var convArg0 unsafe.Pointer
	if strings != nil && strings.QStringList_PTR() != nil {
		convArg0 = strings.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2ERK11QStringListi", qtrt.FFI_TYPE_POINTER, convArg0, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTreeWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:68
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidgetItem(const QStringList &, int)

/*

 */
func NewQTreeWidgetItem_1_(strings qtcore.QStringList_ITF) *QTreeWidgetItem {
	var convArg0 unsafe.Pointer
	if strings != nil && strings.QStringList_PTR() != nil {
		convArg0 = strings.QStringList_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	type_ := 0 /*Type*/
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2ERK11QStringListi", qtrt.FFI_TYPE_POINTER, convArg0, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTreeWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:69
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidgetItem(QTreeWidget *, int)

/*

 */
func NewQTreeWidgetItem_2(view QTreeWidget_ITF /*777 QTreeWidget **/, type_ int) *QTreeWidgetItem {
	var convArg0 unsafe.Pointer
	if view != nil && view.QTreeWidget_PTR() != nil {
		convArg0 = view.QTreeWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EP11QTreeWidgeti", qtrt.FFI_TYPE_POINTER, convArg0, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTreeWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:69
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidgetItem(QTreeWidget *, int)

/*

 */
func NewQTreeWidgetItem_2_(view QTreeWidget_ITF /*777 QTreeWidget **/) *QTreeWidgetItem {
	var convArg0 unsafe.Pointer
	if view != nil && view.QTreeWidget_PTR() != nil {
		convArg0 = view.QTreeWidget_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	type_ := 0 /*Type*/
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EP11QTreeWidgeti", qtrt.FFI_TYPE_POINTER, convArg0, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTreeWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:70
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidgetItem(QTreeWidget *, const QStringList &, int)

/*

 */
func NewQTreeWidgetItem_3(view QTreeWidget_ITF /*777 QTreeWidget **/, strings qtcore.QStringList_ITF, type_ int) *QTreeWidgetItem {
	var convArg0 unsafe.Pointer
	if view != nil && view.QTreeWidget_PTR() != nil {
		convArg0 = view.QTreeWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if strings != nil && strings.QStringList_PTR() != nil {
		convArg1 = strings.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EP11QTreeWidgetRK11QStringListi", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTreeWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:70
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidgetItem(QTreeWidget *, const QStringList &, int)

/*

 */
func NewQTreeWidgetItem_3_(view QTreeWidget_ITF /*777 QTreeWidget **/, strings qtcore.QStringList_ITF) *QTreeWidgetItem {
	var convArg0 unsafe.Pointer
	if view != nil && view.QTreeWidget_PTR() != nil {
		convArg0 = view.QTreeWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if strings != nil && strings.QStringList_PTR() != nil {
		convArg1 = strings.QStringList_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid,
	type_ := 0 /*Type*/
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EP11QTreeWidgetRK11QStringListi", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTreeWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:71
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidgetItem(QTreeWidget *, QTreeWidgetItem *, int)

/*

 */
func NewQTreeWidgetItem_4(view QTreeWidget_ITF /*777 QTreeWidget **/, after QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, type_ int) *QTreeWidgetItem {
	var convArg0 unsafe.Pointer
	if view != nil && view.QTreeWidget_PTR() != nil {
		convArg0 = view.QTreeWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if after != nil && after.QTreeWidgetItem_PTR() != nil {
		convArg1 = after.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EP11QTreeWidgetPS_i", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTreeWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:71
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidgetItem(QTreeWidget *, QTreeWidgetItem *, int)

/*

 */
func NewQTreeWidgetItem_4_(view QTreeWidget_ITF /*777 QTreeWidget **/, after QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) *QTreeWidgetItem {
	var convArg0 unsafe.Pointer
	if view != nil && view.QTreeWidget_PTR() != nil {
		convArg0 = view.QTreeWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if after != nil && after.QTreeWidgetItem_PTR() != nil {
		convArg1 = after.QTreeWidgetItem_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid,
	type_ := 0 /*Type*/
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EP11QTreeWidgetPS_i", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTreeWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:72
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidgetItem(QTreeWidgetItem *, int)

/*

 */
func NewQTreeWidgetItem_5(parent QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, type_ int) *QTreeWidgetItem {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QTreeWidgetItem_PTR() != nil {
		convArg0 = parent.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EPS_i", qtrt.FFI_TYPE_POINTER, convArg0, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTreeWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:72
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidgetItem(QTreeWidgetItem *, int)

/*

 */
func NewQTreeWidgetItem_5_(parent QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) *QTreeWidgetItem {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QTreeWidgetItem_PTR() != nil {
		convArg0 = parent.QTreeWidgetItem_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	type_ := 0 /*Type*/
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EPS_i", qtrt.FFI_TYPE_POINTER, convArg0, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTreeWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:73
// index:6
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidgetItem(QTreeWidgetItem *, const QStringList &, int)

/*

 */
func NewQTreeWidgetItem_6(parent QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, strings qtcore.QStringList_ITF, type_ int) *QTreeWidgetItem {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QTreeWidgetItem_PTR() != nil {
		convArg0 = parent.QTreeWidgetItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if strings != nil && strings.QStringList_PTR() != nil {
		convArg1 = strings.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EPS_RK11QStringListi", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTreeWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:73
// index:6
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidgetItem(QTreeWidgetItem *, const QStringList &, int)

/*

 */
func NewQTreeWidgetItem_6_(parent QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, strings qtcore.QStringList_ITF) *QTreeWidgetItem {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QTreeWidgetItem_PTR() != nil {
		convArg0 = parent.QTreeWidgetItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if strings != nil && strings.QStringList_PTR() != nil {
		convArg1 = strings.QStringList_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid,
	type_ := 0 /*Type*/
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EPS_RK11QStringListi", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTreeWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:74
// index:7
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidgetItem(QTreeWidgetItem *, QTreeWidgetItem *, int)

/*

 */
func NewQTreeWidgetItem_7(parent QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, after QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, type_ int) *QTreeWidgetItem {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QTreeWidgetItem_PTR() != nil {
		convArg0 = parent.QTreeWidgetItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if after != nil && after.QTreeWidgetItem_PTR() != nil {
		convArg1 = after.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EPS_S0_i", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTreeWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:74
// index:7
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidgetItem(QTreeWidgetItem *, QTreeWidgetItem *, int)

/*

 */
func NewQTreeWidgetItem_7_(parent QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, after QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) *QTreeWidgetItem {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QTreeWidgetItem_PTR() != nil {
		convArg0 = parent.QTreeWidgetItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if after != nil && after.QTreeWidgetItem_PTR() != nil {
		convArg1 = after.QTreeWidgetItem_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid,
	type_ := 0 /*Type*/
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EPS_S0_i", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTreeWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTreeWidgetItem()

/*

 */
func DeleteQTreeWidgetItem(this *QTreeWidgetItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 64)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QTreeWidgetItem * clone() const

/*

 */
func (this *QTreeWidgetItem) Clone() *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem5cloneEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:80
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QTreeWidget * treeWidget() const

/*

 */
func (this *QTreeWidgetItem) TreeWidget() *QTreeWidget /*777 QTreeWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10treeWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:82
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setSelected(bool)

/*

 */
func (this *QTreeWidgetItem) SetSelected(select_ bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem11setSelectedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), select_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:83
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isSelected() const

/*

 */
func (this *QTreeWidgetItem) IsSelected() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10isSelectedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:85
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHidden(bool)

/*

 */
func (this *QTreeWidgetItem) SetHidden(hide bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem9setHiddenEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hide)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:86
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isHidden() const

/*

 */
func (this *QTreeWidgetItem) IsHidden() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem8isHiddenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:88
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setExpanded(bool)

/*

 */
func (this *QTreeWidgetItem) SetExpanded(expand bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem11setExpandedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), expand)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isExpanded() const

/*

 */
func (this *QTreeWidgetItem) IsExpanded() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10isExpandedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:91
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFirstColumnSpanned(bool)

/*

 */
func (this *QTreeWidgetItem) SetFirstColumnSpanned(span bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem21setFirstColumnSpannedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), span)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:92
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isFirstColumnSpanned() const

/*

 */
func (this *QTreeWidgetItem) IsFirstColumnSpanned() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem20isFirstColumnSpannedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:94
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setDisabled(bool)

/*

 */
func (this *QTreeWidgetItem) SetDisabled(disabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem11setDisabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), disabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:95
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDisabled() const

/*

 */
func (this *QTreeWidgetItem) IsDisabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10isDisabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setChildIndicatorPolicy(QTreeWidgetItem::ChildIndicatorPolicy)

/*

 */
func (this *QTreeWidgetItem) SetChildIndicatorPolicy(policy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem23setChildIndicatorPolicyENS_20ChildIndicatorPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] QTreeWidgetItem::ChildIndicatorPolicy childIndicatorPolicy() const

/*

 */
func (this *QTreeWidgetItem) ChildIndicatorPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem20childIndicatorPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags() const

/*

 */
func (this *QTreeWidgetItem) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(Qt::ItemFlags)

/*

 */
func (this *QTreeWidgetItem) SetFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem8setFlagsE6QFlagsIN2Qt8ItemFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:104
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString text(int) const

/*

 */
func (this *QTreeWidgetItem) Text(column int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem4textEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtreewidget.h:106
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setText(int, const QString &)

/*

 */
func (this *QTreeWidgetItem) SetText(column int, text string) {
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem7setTextEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:108
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QIcon icon(int) const

/*

 */
func (this *QTreeWidgetItem) Icon(column int) *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem4iconEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:110
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setIcon(int, const QIcon &)

/*

 */
func (this *QTreeWidgetItem) SetIcon(column int, icon qtgui.QIcon_ITF) {
	var convArg1 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg1 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem7setIconEiRK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:112
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString statusTip(int) const

/*

 */
func (this *QTreeWidgetItem) StatusTip(column int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem9statusTipEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtreewidget.h:114
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setStatusTip(int, const QString &)

/*

 */
func (this *QTreeWidgetItem) SetStatusTip(column int, statusTip string) {
	var tmpArg1 = qtcore.NewQString_5(statusTip)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem12setStatusTipEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:117
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString toolTip(int) const

/*

 */
func (this *QTreeWidgetItem) ToolTip(column int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem7toolTipEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtreewidget.h:119
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setToolTip(int, const QString &)

/*

 */
func (this *QTreeWidgetItem) SetToolTip(column int, toolTip string) {
	var tmpArg1 = qtcore.NewQString_5(toolTip)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem10setToolTipEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:123
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString whatsThis(int) const

/*

 */
func (this *QTreeWidgetItem) WhatsThis(column int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem9whatsThisEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtreewidget.h:125
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWhatsThis(int, const QString &)

/*

 */
func (this *QTreeWidgetItem) SetWhatsThis(column int, whatsThis string) {
	var tmpArg1 = qtcore.NewQString_5(whatsThis)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem12setWhatsThisEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:128
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QFont font(int) const

/*

 */
func (this *QTreeWidgetItem) Font(column int) *qtgui.QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem4fontEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:130
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFont(int, const QFont &)

/*

 */
func (this *QTreeWidgetItem) SetFont(column int, font qtgui.QFont_ITF) {
	var convArg1 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg1 = font.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem7setFontEiRK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:132
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int textAlignment(int) const

/*

 */
func (this *QTreeWidgetItem) TextAlignment(column int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem13textAlignmentEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreewidget.h:134
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTextAlignment(int, int)

/*

 */
func (this *QTreeWidgetItem) SetTextAlignment(column int, alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem16setTextAlignmentEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:137
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QColor backgroundColor(int) const

/*

 */
func (this *QTreeWidgetItem) BackgroundColor(column int) *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem15backgroundColorEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:139
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBackgroundColor(int, const QColor &)

/*

 */
func (this *QTreeWidgetItem) SetBackgroundColor(column int, color qtgui.QColor_ITF) {
	var convArg1 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg1 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem18setBackgroundColorEiRK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:142
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QBrush background(int) const

/*

 */
func (this *QTreeWidgetItem) Background(column int) *qtgui.QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10backgroundEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:144
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBackground(int, const QBrush &)

/*

 */
func (this *QTreeWidgetItem) SetBackground(column int, brush qtgui.QBrush_ITF) {
	var convArg1 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg1 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem13setBackgroundEiRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:147
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QColor textColor(int) const

/*

 */
func (this *QTreeWidgetItem) TextColor(column int) *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem9textColorEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:149
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTextColor(int, const QColor &)

/*

 */
func (this *QTreeWidgetItem) SetTextColor(column int, color qtgui.QColor_ITF) {
	var convArg1 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg1 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem12setTextColorEiRK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:152
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QBrush foreground(int) const

/*

 */
func (this *QTreeWidgetItem) Foreground(column int) *qtgui.QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10foregroundEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:154
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setForeground(int, const QBrush &)

/*

 */
func (this *QTreeWidgetItem) SetForeground(column int, brush qtgui.QBrush_ITF) {
	var convArg1 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg1 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem13setForegroundEiRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:157
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::CheckState checkState(int) const

/*

 */
func (this *QTreeWidgetItem) CheckState(column int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10checkStateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:159
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setCheckState(int, Qt::CheckState)

/*

 */
func (this *QTreeWidgetItem) SetCheckState(column int, state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem13setCheckStateEiN2Qt10CheckStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:162
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSize sizeHint(int) const

/*

 */
func (this *QTreeWidgetItem) SizeHint(column int) *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem8sizeHintEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:164
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setSizeHint(int, const QSize &)

/*

 */
func (this *QTreeWidgetItem) SetSizeHint(column int, size qtcore.QSize_ITF) {
	var convArg1 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg1 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem11setSizeHintEiRK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:167
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant data(int, int) const

/*

 */
func (this *QTreeWidgetItem) Data(column int, role int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem4dataEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, role)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:168
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setData(int, int, const QVariant &)

/*

 */
func (this *QTreeWidgetItem) SetData(column int, role int, value qtcore.QVariant_ITF) {
	var convArg2 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg2 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem7setDataEiiRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, role, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:170
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool operator<(const QTreeWidgetItem &) const

/*

 */
func (this *QTreeWidgetItem) Operator_less_than(other QTreeWidgetItem_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTreeWidgetItem_PTR() != nil {
		convArg0 = other.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItemltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:173
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void read(QDataStream &)

/*

 */
func (this *QTreeWidgetItem) Read(in qtcore.QDataStream_ITF) {
	var convArg0 unsafe.Pointer
	if in != nil && in.QDataStream_PTR() != nil {
		convArg0 = in.QDataStream_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem4readER11QDataStream", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:174
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void write(QDataStream &) const

/*

 */
func (this *QTreeWidgetItem) Write(out qtcore.QDataStream_ITF) {
	var convArg0 unsafe.Pointer
	if out != nil && out.QDataStream_PTR() != nil {
		convArg0 = out.QDataStream_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem5writeER11QDataStream", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:176
// index:0
// Public Visibility=Default Availability=Available
// [64] QTreeWidgetItem & operator=(const QTreeWidgetItem &)

/*

 */
func (this *QTreeWidgetItem) Operator_equal(other QTreeWidgetItem_ITF) *QTreeWidgetItem {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTreeWidgetItem_PTR() != nil {
		convArg0 = other.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItemaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTreeWidgetItem)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:178
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QTreeWidgetItem * parent() const

/*

 */
func (this *QTreeWidgetItem) Parent() *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem6parentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:179
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QTreeWidgetItem * child(int) const

/*

 */
func (this *QTreeWidgetItem) Child(index int) *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem5childEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:185
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int childCount() const

/*

 */
func (this *QTreeWidgetItem) ChildCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10childCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreewidget.h:186
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int columnCount() const

/*

 */
func (this *QTreeWidgetItem) ColumnCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem11columnCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreewidget.h:187
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int indexOfChild(QTreeWidgetItem *) const

/*

 */
func (this *QTreeWidgetItem) IndexOfChild(child QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) int {
	var convArg0 unsafe.Pointer
	if child != nil && child.QTreeWidgetItem_PTR() != nil {
		convArg0 = child.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem12indexOfChildEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreewidget.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addChild(QTreeWidgetItem *)

/*

 */
func (this *QTreeWidgetItem) AddChild(child QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if child != nil && child.QTreeWidgetItem_PTR() != nil {
		convArg0 = child.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem8addChildEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertChild(int, QTreeWidgetItem *)

/*

 */
func (this *QTreeWidgetItem) InsertChild(index int, child QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg1 unsafe.Pointer
	if child != nil && child.QTreeWidgetItem_PTR() != nil {
		convArg1 = child.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem11insertChildEiPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeChild(QTreeWidgetItem *)

/*

 */
func (this *QTreeWidgetItem) RemoveChild(child QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if child != nil && child.QTreeWidgetItem_PTR() != nil {
		convArg0 = child.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem11removeChildEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:192
// index:0
// Public Visibility=Default Availability=Available
// [8] QTreeWidgetItem * takeChild(int)

/*

 */
func (this *QTreeWidgetItem) TakeChild(index int) *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem9takeChildEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:198
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int type() const

/*

 */
func (this *QTreeWidgetItem) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreewidget.h:199
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void sortChildren(int, Qt::SortOrder)

/*

 */
func (this *QTreeWidgetItem) SortChildren(column int, order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem12sortChildrenEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:203
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void emitDataChanged()

/*

 */
func (this *QTreeWidgetItem) EmitDataChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem15emitDataChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QTreeWidgetItem__ItemType = int

//
const QTreeWidgetItem__Type QTreeWidgetItem__ItemType = 0

//
const QTreeWidgetItem__UserType QTreeWidgetItem__ItemType = 1000

/*


 */
type QTreeWidgetItem__ChildIndicatorPolicy = int

//
const QTreeWidgetItem__ShowIndicator QTreeWidgetItem__ChildIndicatorPolicy = 0

//
const QTreeWidgetItem__DontShowIndicator QTreeWidgetItem__ChildIndicatorPolicy = 1

//
const QTreeWidgetItem__DontShowIndicatorWhenChildless QTreeWidgetItem__ChildIndicatorPolicy = 2

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
