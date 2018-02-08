package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

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
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QAccessibleInterface struct {
	*qtrt.CObject
}

func (this *QAccessibleInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAccessibleInterface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAccessibleInterfaceFromPointer(cthis unsafe.Pointer) *QAccessibleInterface {
	return &QAccessibleInterface{&qtrt.CObject{cthis}}
}
func (*QAccessibleInterface) NewFromPointer(cthis unsafe.Pointer) *QAccessibleInterface {
	return NewQAccessibleInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessible.h:460
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleInterface()
func DeleteQAccessibleInterface(this *QAccessibleInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAccessibleInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessible.h:464
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QAccessibleInterface) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAccessibleInterface7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qaccessible.h:465
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QObject * object()
func (this *QAccessibleInterface) Object() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAccessibleInterface6objectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qaccessible.h:466
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QWindow * window()
func (this *QAccessibleInterface) Window() *QWindow /*777 QWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAccessibleInterface6windowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qaccessible.h:470
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * focusChild()
func (this *QAccessibleInterface) FocusChild() *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAccessibleInterface10focusChildEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qaccessible.h:472
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * childAt(int, int)
func (this *QAccessibleInterface) ChildAt(x int, y int) *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAccessibleInterface7childAtEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qaccessible.h:475
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * parent()
func (this *QAccessibleInterface) Parent() *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAccessibleInterface6parentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qaccessible.h:476
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * child(int)
func (this *QAccessibleInterface) Child(index int) *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAccessibleInterface5childEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qaccessible.h:477
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int childCount()
func (this *QAccessibleInterface) ChildCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAccessibleInterface10childCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:478
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int indexOfChild(const QAccessibleInterface *)
func (this *QAccessibleInterface) IndexOfChild(arg0 *QAccessibleInterface /*777 const QAccessibleInterface **/) int {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAccessibleInterface12indexOfChildEPKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:481
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString text(QAccessible::Text)
func (this *QAccessibleInterface) Text(t int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAccessibleInterface4textEN11QAccessible4TextE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), t)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qaccessible.h:482
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setText(QAccessible::Text, const QString &)
func (this *QAccessibleInterface) SetText(t int, text string) {
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAccessibleInterface7setTextEN11QAccessible4TextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), t, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:483
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QRect rect()
func (this *QAccessibleInterface) Rect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAccessibleInterface4rectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qaccessible.h:484
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QAccessible::Role role()
func (this *QAccessibleInterface) Role() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAccessibleInterface4roleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qaccessible.h:485
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAccessible::State state()
func (this *QAccessibleInterface) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAccessibleInterface5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qaccessible.h:487
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QColor foregroundColor()
func (this *QAccessibleInterface) ForegroundColor() *QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAccessibleInterface15foregroundColorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qaccessible.h:488
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QColor backgroundColor()
func (this *QAccessibleInterface) BackgroundColor() *QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAccessibleInterface15backgroundColorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qaccessible.h:490
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QAccessibleTextInterface * textInterface()
func (this *QAccessibleInterface) TextInterface() *QAccessibleTextInterface /*777 QAccessibleTextInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAccessibleInterface13textInterfaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleTextInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qaccessible.h:493
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QAccessibleEditableTextInterface * editableTextInterface()
func (this *QAccessibleInterface) EditableTextInterface() *QAccessibleEditableTextInterface /*777 QAccessibleEditableTextInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAccessibleInterface21editableTextInterfaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleEditableTextInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qaccessible.h:496
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QAccessibleValueInterface * valueInterface()
func (this *QAccessibleInterface) ValueInterface() *QAccessibleValueInterface /*777 QAccessibleValueInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAccessibleInterface14valueInterfaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleValueInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qaccessible.h:502
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QAccessibleImageInterface * imageInterface()
func (this *QAccessibleInterface) ImageInterface() *QAccessibleImageInterface /*777 QAccessibleImageInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAccessibleInterface14imageInterfaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleImageInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qaccessible.h:505
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QAccessibleTableInterface * tableInterface()
func (this *QAccessibleInterface) TableInterface() *QAccessibleTableInterface /*777 QAccessibleTableInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAccessibleInterface14tableInterfaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleTableInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qaccessible.h:508
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QAccessibleTableCellInterface * tableCellInterface()
func (this *QAccessibleInterface) TableCellInterface() *QAccessibleTableCellInterface /*777 QAccessibleTableCellInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAccessibleInterface18tableCellInterfaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleTableCellInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qaccessible.h:511
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void virtual_hook(int, void *)
func (this *QAccessibleInterface) Virtual_hook(id int, data unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAccessibleInterface12virtual_hookEiPv", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:513
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [8] void * interface_cast(QAccessible::InterfaceType)
func (this *QAccessibleInterface) Interface_cast(arg0 int) unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAccessibleInterface14interface_castEN11QAccessible13InterfaceTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

//  body block end
