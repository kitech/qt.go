package qtwidgets

// /usr/include/qt/QtWidgets/qlistwidget.h
// #include <qlistwidget.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 68
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
type QListWidgetItem struct {
	*qtrt.CObject
}
type QListWidgetItem_ITF interface {
	QListWidgetItem_PTR() *QListWidgetItem
}

func (ptr *QListWidgetItem) QListWidgetItem_PTR() *QListWidgetItem { return ptr }

func (this *QListWidgetItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QListWidgetItem) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQListWidgetItemFromPointer(cthis unsafe.Pointer) *QListWidgetItem {
	return &QListWidgetItem{&qtrt.CObject{cthis}}
}
func (*QListWidgetItem) NewFromPointer(cthis unsafe.Pointer) *QListWidgetItem {
	return NewQListWidgetItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QListWidgetItem(QListWidget *, int)

/*

 */
func (*QListWidgetItem) NewForInherit(view QListWidget_ITF /*777 QListWidget **/, type_ int) *QListWidgetItem {
	return NewQListWidgetItem(view, type_)
}
func NewQListWidgetItem(view QListWidget_ITF /*777 QListWidget **/, type_ int) *QListWidgetItem {
	var convArg0 unsafe.Pointer
	if view != nil && view.QListWidget_PTR() != nil {
		convArg0 = view.QListWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItemC2EP11QListWidgeti", qtrt.FFI_TYPE_POINTER, convArg0, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQListWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qlistwidget.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QListWidgetItem(QListWidget *, int)

/*

 */
func (*QListWidgetItem) NewForInheritp() *QListWidgetItem {
	return NewQListWidgetItemp()
}
func NewQListWidgetItemp() *QListWidgetItem {
	// arg: 0, QListWidget *=Pointer, QListWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	type_ := 0 /*Type*/
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItemC2EP11QListWidgeti", qtrt.FFI_TYPE_POINTER, convArg0, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQListWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qlistwidget.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QListWidgetItem(QListWidget *, int)

/*

 */
func (*QListWidgetItem) NewForInheritp1(view QListWidget_ITF /*777 QListWidget **/) *QListWidgetItem {
	return NewQListWidgetItemp1(view)
}
func NewQListWidgetItemp1(view QListWidget_ITF /*777 QListWidget **/) *QListWidgetItem {
	var convArg0 unsafe.Pointer
	if view != nil && view.QListWidget_PTR() != nil {
		convArg0 = view.QListWidget_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	type_ := 0 /*Type*/
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItemC2EP11QListWidgeti", qtrt.FFI_TYPE_POINTER, convArg0, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQListWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qlistwidget.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QListWidgetItem(const QString &, QListWidget *, int)

/*

 */
func (*QListWidgetItem) NewForInherit1(text string, view QListWidget_ITF /*777 QListWidget **/, type_ int) *QListWidgetItem {
	return NewQListWidgetItem1(text, view, type_)
}
func NewQListWidgetItem1(text string, view QListWidget_ITF /*777 QListWidget **/, type_ int) *QListWidgetItem {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if view != nil && view.QListWidget_PTR() != nil {
		convArg1 = view.QListWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItemC2ERK7QStringP11QListWidgeti", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQListWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qlistwidget.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QListWidgetItem(const QString &, QListWidget *, int)

/*

 */
func (*QListWidgetItem) NewForInherit1p(text string) *QListWidgetItem {
	return NewQListWidgetItem1p(text)
}
func NewQListWidgetItem1p(text string) *QListWidgetItem {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QListWidget *=Pointer, QListWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, int=Int, =Invalid, , Invalid
	type_ := 0 /*Type*/
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItemC2ERK7QStringP11QListWidgeti", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQListWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qlistwidget.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QListWidgetItem(const QString &, QListWidget *, int)

/*

 */
func (*QListWidgetItem) NewForInherit1p1(text string, view QListWidget_ITF /*777 QListWidget **/) *QListWidgetItem {
	return NewQListWidgetItem1p1(text, view)
}
func NewQListWidgetItem1p1(text string, view QListWidget_ITF /*777 QListWidget **/) *QListWidgetItem {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if view != nil && view.QListWidget_PTR() != nil {
		convArg1 = view.QListWidget_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid, , Invalid
	type_ := 0 /*Type*/
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItemC2ERK7QStringP11QListWidgeti", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQListWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qlistwidget.h:66
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QListWidgetItem(const QIcon &, const QString &, QListWidget *, int)

/*

 */
func (*QListWidgetItem) NewForInherit2(icon qtgui.QIcon_ITF, text string, view QListWidget_ITF /*777 QListWidget **/, type_ int) *QListWidgetItem {
	return NewQListWidgetItem2(icon, text, view, type_)
}
func NewQListWidgetItem2(icon qtgui.QIcon_ITF, text string, view QListWidget_ITF /*777 QListWidget **/, type_ int) *QListWidgetItem {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if view != nil && view.QListWidget_PTR() != nil {
		convArg2 = view.QListWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItemC2ERK5QIconRK7QStringP11QListWidgeti", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQListWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qlistwidget.h:66
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QListWidgetItem(const QIcon &, const QString &, QListWidget *, int)

/*

 */
func (*QListWidgetItem) NewForInherit2p(icon qtgui.QIcon_ITF, text string) *QListWidgetItem {
	return NewQListWidgetItem2p(icon, text)
}
func NewQListWidgetItem2p(icon qtgui.QIcon_ITF, text string) *QListWidgetItem {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, QListWidget *=Pointer, QListWidget=Record, , Invalid
	var convArg2 unsafe.Pointer
	// arg: 3, int=Int, =Invalid, , Invalid
	type_ := 0 /*Type*/
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItemC2ERK5QIconRK7QStringP11QListWidgeti", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQListWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qlistwidget.h:66
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QListWidgetItem(const QIcon &, const QString &, QListWidget *, int)

/*

 */
func (*QListWidgetItem) NewForInherit2p1(icon qtgui.QIcon_ITF, text string, view QListWidget_ITF /*777 QListWidget **/) *QListWidgetItem {
	return NewQListWidgetItem2p1(icon, text, view)
}
func NewQListWidgetItem2p1(icon qtgui.QIcon_ITF, text string, view QListWidget_ITF /*777 QListWidget **/) *QListWidgetItem {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if view != nil && view.QListWidget_PTR() != nil {
		convArg2 = view.QListWidget_PTR().GetCthis()
	}
	// arg: 3, int=Int, =Invalid, , Invalid
	type_ := 0 /*Type*/
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItemC2ERK5QIconRK7QStringP11QListWidgeti", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQListWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qlistwidget.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QListWidgetItem()

/*

 */
func DeleteQListWidgetItem(this *QListWidgetItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QListWidgetItem * clone() const

/*

 */
func (this *QListWidgetItem) Clone() *QListWidgetItem /*777 QListWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QListWidgetItem5cloneEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlistwidget.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QListWidget * listWidget() const

/*

 */
func (this *QListWidgetItem) ListWidget() *QListWidget /*777 QListWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QListWidgetItem10listWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQListWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlistwidget.h:75
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setSelected(bool)

/*

 */
func (this *QListWidgetItem) SetSelected(select_ bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItem11setSelectedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), select_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:76
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isSelected() const

/*

 */
func (this *QListWidgetItem) IsSelected() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QListWidgetItem10isSelectedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistwidget.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHidden(bool)

/*

 */
func (this *QListWidgetItem) SetHidden(hide bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItem9setHiddenEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hide)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isHidden() const

/*

 */
func (this *QListWidgetItem) IsHidden() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QListWidgetItem8isHiddenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistwidget.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags() const

/*

 */
func (this *QListWidgetItem) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QListWidgetItem5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(Qt::ItemFlags)

/*

 */
func (this *QListWidgetItem) SetFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItem8setFlagsE6QFlagsIN2Qt8ItemFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:84
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString text() const

/*

 */
func (this *QListWidgetItem) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QListWidgetItem4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qlistwidget.h:86
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setText(const QString &)

/*

 */
func (this *QListWidgetItem) SetText(text string) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItem7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:88
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QIcon icon() const

/*

 */
func (this *QListWidgetItem) Icon() *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QListWidgetItem4iconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:90
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setIcon(const QIcon &)

/*

 */
func (this *QListWidgetItem) SetIcon(icon qtgui.QIcon_ITF) {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItem7setIconERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:92
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString statusTip() const

/*

 */
func (this *QListWidgetItem) StatusTip() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QListWidgetItem9statusTipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qlistwidget.h:94
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setStatusTip(const QString &)

/*

 */
func (this *QListWidgetItem) SetStatusTip(statusTip string) {
	var tmpArg0 = qtcore.NewQString5(statusTip)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItem12setStatusTipERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:97
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString toolTip() const

/*

 */
func (this *QListWidgetItem) ToolTip() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QListWidgetItem7toolTipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qlistwidget.h:99
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setToolTip(const QString &)

/*

 */
func (this *QListWidgetItem) SetToolTip(toolTip string) {
	var tmpArg0 = qtcore.NewQString5(toolTip)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItem10setToolTipERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:103
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString whatsThis() const

/*

 */
func (this *QListWidgetItem) WhatsThis() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QListWidgetItem9whatsThisEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qlistwidget.h:105
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWhatsThis(const QString &)

/*

 */
func (this *QListWidgetItem) SetWhatsThis(whatsThis string) {
	var tmpArg0 = qtcore.NewQString5(whatsThis)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItem12setWhatsThisERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:108
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QFont font() const

/*

 */
func (this *QListWidgetItem) Font() *qtgui.QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QListWidgetItem4fontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:110
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)

/*

 */
func (this *QListWidgetItem) SetFont(font qtgui.QFont_ITF) {
	var convArg0 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg0 = font.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItem7setFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:112
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int textAlignment() const

/*

 */
func (this *QListWidgetItem) TextAlignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QListWidgetItem13textAlignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlistwidget.h:114
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTextAlignment(int)

/*

 */
func (this *QListWidgetItem) SetTextAlignment(alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItem16setTextAlignmentEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:117
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QColor backgroundColor() const

/*

 */
func (this *QListWidgetItem) BackgroundColor() *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QListWidgetItem15backgroundColorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:119
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [-2] void setBackgroundColor(const QColor &)

/*

 */
func (this *QListWidgetItem) SetBackgroundColor(color qtgui.QColor_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItem18setBackgroundColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:122
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QBrush background() const

/*

 */
func (this *QListWidgetItem) Background() *qtgui.QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QListWidgetItem10backgroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:124
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBackground(const QBrush &)

/*

 */
func (this *QListWidgetItem) SetBackground(brush qtgui.QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg0 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItem13setBackgroundERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:127
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QColor textColor() const

/*

 */
func (this *QListWidgetItem) TextColor() *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QListWidgetItem9textColorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:129
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTextColor(const QColor &)

/*

 */
func (this *QListWidgetItem) SetTextColor(color qtgui.QColor_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItem12setTextColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:132
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QBrush foreground() const

/*

 */
func (this *QListWidgetItem) Foreground() *qtgui.QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QListWidgetItem10foregroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:134
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setForeground(const QBrush &)

/*

 */
func (this *QListWidgetItem) SetForeground(brush qtgui.QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg0 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItem13setForegroundERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:137
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::CheckState checkState() const

/*

 */
func (this *QListWidgetItem) CheckState() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QListWidgetItem10checkStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:139
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setCheckState(Qt::CheckState)

/*

 */
func (this *QListWidgetItem) SetCheckState(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItem13setCheckStateEN2Qt10CheckStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:142
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*

 */
func (this *QListWidgetItem) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QListWidgetItem8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:144
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setSizeHint(const QSize &)

/*

 */
func (this *QListWidgetItem) SetSizeHint(size qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItem11setSizeHintERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:147
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant data(int) const

/*

 */
func (this *QListWidgetItem) Data(role int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QListWidgetItem4dataEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:148
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setData(int, const QVariant &)

/*

 */
func (this *QListWidgetItem) SetData(role int, value qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItem7setDataEiRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:150
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool operator<(const QListWidgetItem &) const

/*

 */
func (this *QListWidgetItem) Operator_less_than(other QListWidgetItem_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QListWidgetItem_PTR() != nil {
		convArg0 = other.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QListWidgetItemltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistwidget.h:153
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void read(QDataStream &)

/*

 */
func (this *QListWidgetItem) Read(in qtcore.QDataStream_ITF) {
	var convArg0 unsafe.Pointer
	if in != nil && in.QDataStream_PTR() != nil {
		convArg0 = in.QDataStream_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItem4readER11QDataStream", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:154
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void write(QDataStream &) const

/*

 */
func (this *QListWidgetItem) Write(out_ qtcore.QDataStream_ITF) {
	var convArg0 unsafe.Pointer
	if out_ != nil && out_.QDataStream_PTR() != nil {
		convArg0 = out_.QDataStream_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QListWidgetItem5writeER11QDataStream", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:156
// index:0
// Public Visibility=Default Availability=Available
// [48] QListWidgetItem & operator=(const QListWidgetItem &)

/*

 */
func (this *QListWidgetItem) Operator_equal(other QListWidgetItem_ITF) *QListWidgetItem {
	var convArg0 unsafe.Pointer
	if other != nil && other.QListWidgetItem_PTR() != nil {
		convArg0 = other.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QListWidgetItemaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQListWidgetItem)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:158
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int type() const

/*

 */
func (this *QListWidgetItem) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QListWidgetItem4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

/*


 */
type QListWidgetItem__ItemType = int

//
const QListWidgetItem__Type QListWidgetItem__ItemType = 0

//
const QListWidgetItem__UserType QListWidgetItem__ItemType = 1000

func (this *QListWidgetItem) ItemTypeItemName(val int) string {
	switch val {
	case QListWidgetItem__Type: // 0
		return "Type"
	case QListWidgetItem__UserType: // 1000
		return "UserType"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QListWidgetItem_ItemTypeItemName(val int) string {
	var nilthis *QListWidgetItem
	return nilthis.ItemTypeItemName(val)
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
