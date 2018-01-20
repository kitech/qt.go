//  header block begin
// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
}

//  ext block end

//  body block begin
type QAccessibleInterface struct {
	*qtrt.CObject
}

func (this *QAccessibleInterface) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQAccessibleInterfaceFromPointer(cthis unsafe.Pointer) *QAccessibleInterface {
	return &QAccessibleInterface{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qaccessible.h:460
// index:0
// Protected virtual
// void ~QAccessibleInterface()
func DeleteQAccessibleInterface(*QAccessibleInterface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QAccessibleInterfaceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:464
// index:0
// Public pure virtual
// bool isValid()
func (this *QAccessibleInterface) IsValid() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAccessibleInterface7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:465
// index:0
// Public pure virtual
// QObject * object()
func (this *QAccessibleInterface) Object() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAccessibleInterface6objectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:466
// index:0
// Public virtual
// QWindow * window()
func (this *QAccessibleInterface) Window() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAccessibleInterface6windowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:470
// index:0
// Public virtual
// QAccessibleInterface * focusChild()
func (this *QAccessibleInterface) FocusChild() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAccessibleInterface10focusChildEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:472
// index:0
// Public pure virtual
// QAccessibleInterface * childAt(int, int)
func (this *QAccessibleInterface) ChildAt(x int, y int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAccessibleInterface7childAtEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:475
// index:0
// Public pure virtual
// QAccessibleInterface * parent()
func (this *QAccessibleInterface) Parent() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAccessibleInterface6parentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:476
// index:0
// Public pure virtual
// QAccessibleInterface * child(int)
func (this *QAccessibleInterface) Child(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAccessibleInterface5childEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:477
// index:0
// Public pure virtual
// int childCount()
func (this *QAccessibleInterface) ChildCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAccessibleInterface10childCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:478
// index:0
// Public pure virtual
// int indexOfChild(const class QAccessibleInterface *)
func (this *QAccessibleInterface) IndexOfChild(arg0 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAccessibleInterface12indexOfChildEPKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:481
// index:0
// Public pure virtual
// QString text(class QAccessible::Text)
func (this *QAccessibleInterface) Text(t int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAccessibleInterface4textEN11QAccessible4TextE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &t)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:482
// index:0
// Public pure virtual
// void setText(class QAccessible::Text, const class QString &)
func (this *QAccessibleInterface) SetText(t int, text *qtcore.QString) {
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QAccessibleInterface7setTextEN11QAccessible4TextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &t, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:483
// index:0
// Public pure virtual
// QRect rect()
func (this *QAccessibleInterface) Rect() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAccessibleInterface4rectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:484
// index:0
// Public pure virtual
// QAccessible::Role role()
func (this *QAccessibleInterface) Role() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAccessibleInterface4roleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:485
// index:0
// Public pure virtual
// QAccessible::State state()
func (this *QAccessibleInterface) State() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAccessibleInterface5stateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:487
// index:0
// Public virtual
// QColor foregroundColor()
func (this *QAccessibleInterface) ForegroundColor() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAccessibleInterface15foregroundColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:488
// index:0
// Public virtual
// QColor backgroundColor()
func (this *QAccessibleInterface) BackgroundColor() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAccessibleInterface15backgroundColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:490
// index:0
// Public inline
// QAccessibleTextInterface * textInterface()
func (this *QAccessibleInterface) TextInterface() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QAccessibleInterface13textInterfaceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:493
// index:0
// Public inline
// QAccessibleEditableTextInterface * editableTextInterface()
func (this *QAccessibleInterface) EditableTextInterface() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QAccessibleInterface21editableTextInterfaceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:496
// index:0
// Public inline
// QAccessibleValueInterface * valueInterface()
func (this *QAccessibleInterface) ValueInterface() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QAccessibleInterface14valueInterfaceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:499
// index:0
// Public inline
// QAccessibleActionInterface * actionInterface()
func (this *QAccessibleInterface) ActionInterface() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QAccessibleInterface15actionInterfaceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:502
// index:0
// Public inline
// QAccessibleImageInterface * imageInterface()
func (this *QAccessibleInterface) ImageInterface() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QAccessibleInterface14imageInterfaceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:505
// index:0
// Public inline
// QAccessibleTableInterface * tableInterface()
func (this *QAccessibleInterface) TableInterface() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QAccessibleInterface14tableInterfaceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:508
// index:0
// Public inline
// QAccessibleTableCellInterface * tableCellInterface()
func (this *QAccessibleInterface) TableCellInterface() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QAccessibleInterface18tableCellInterfaceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:511
// index:0
// Public virtual
// void virtual_hook(int, void *)
func (this *QAccessibleInterface) Virtual_hook(id int, data unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QAccessibleInterface12virtual_hookEiPv", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &id, data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:513
// index:0
// Public inline virtual
// void * interface_cast(class QAccessible::InterfaceType)
func (this *QAccessibleInterface) Interface_cast(arg0 int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QAccessibleInterface14interface_castEN11QAccessible13InterfaceTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
