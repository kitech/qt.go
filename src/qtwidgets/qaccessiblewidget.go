package qtwidgets

// /usr/include/qt/QtWidgets/qaccessiblewidget.h
// #include <qaccessiblewidget.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 133
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
type QAccessibleWidget struct {
	*qtrt.CObject
}

func (this *QAccessibleWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQAccessibleWidgetFromPointer(cthis unsafe.Pointer) *QAccessibleWidget {
	return &QAccessibleWidget{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:56
// index:0
// Public
// void QAccessibleWidget(class QWidget *, class QAccessible::Role, const class QString &)
func NewQAccessibleWidget(o *QWidget /*444 QWidget **/, r int, name *qtcore.QString) *QAccessibleWidget {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = o.GetCthis()
	var convArg2 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleWidgetC2EP7QWidgetN11QAccessible4RoleERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &r, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleWidgetFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:57
// index:0
// Public virtual
// bool isValid()
func (this *QAccessibleWidget) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:59
// index:0
// Public virtual
// QWindow * window()
func (this *QAccessibleWidget) Window() *qtgui.QWindow /*444 QWindow **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget6windowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:60
// index:0
// Public virtual
// int childCount()
func (this *QAccessibleWidget) ChildCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget10childCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:61
// index:0
// Public virtual
// int indexOfChild(const class QAccessibleInterface *)
func (this *QAccessibleWidget) IndexOfChild(child *qtgui.QAccessibleInterface /*444 const QAccessibleInterface **/) int {
	var convArg0 = child.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget12indexOfChildEPK20QAccessibleInterface", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:63
// index:0
// Public virtual
// QAccessibleInterface * focusChild()
func (this *QAccessibleWidget) FocusChild() *qtgui.QAccessibleInterface /*444 QAccessibleInterface **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget10focusChildEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:65
// index:0
// Public virtual
// QRect rect()
func (this *QAccessibleWidget) Rect() *qtcore.QRect /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget4rectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:67
// index:0
// Public virtual
// QAccessibleInterface * parent()
func (this *QAccessibleWidget) Parent() *qtgui.QAccessibleInterface /*444 QAccessibleInterface **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget6parentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:68
// index:0
// Public virtual
// QAccessibleInterface * child(int)
func (this *QAccessibleWidget) Child(index int) *qtgui.QAccessibleInterface /*444 QAccessibleInterface **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget5childEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:70
// index:0
// Public virtual
// QString text(class QAccessible::Text)
func (this *QAccessibleWidget) Text(t int) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget4textEN11QAccessible4TextE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &t)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:71
// index:0
// Public virtual
// QAccessible::Role role()
func (this *QAccessibleWidget) Role() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget4roleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:72
// index:0
// Public virtual
// QAccessible::State state()
func (this *QAccessibleWidget) State() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget5stateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:74
// index:0
// Public virtual
// QColor foregroundColor()
func (this *QAccessibleWidget) ForegroundColor() *qtgui.QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget15foregroundColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:75
// index:0
// Public virtual
// QColor backgroundColor()
func (this *QAccessibleWidget) BackgroundColor() *qtgui.QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget15backgroundColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:77
// index:0
// Public virtual
// void * interface_cast(class QAccessible::InterfaceType)
func (this *QAccessibleWidget) Interface_cast(t int) unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleWidget14interface_castEN11QAccessible13InterfaceTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &t)
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:81
// index:0
// Public virtual
// void doAction(const class QString &)
func (this *QAccessibleWidget) DoAction(actionName *qtcore.QString) {
	var convArg0 = actionName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleWidget8doActionERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:84
// index:0
// Protected virtual
// void ~QAccessibleWidget()
func DeleteQAccessibleWidget(*QAccessibleWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:85
// index:0
// Protected
// QWidget * widget()
func (this *QAccessibleWidget) Widget() *QWidget /*444 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget6widgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:86
// index:0
// Protected
// QObject * parentObject()
func (this *QAccessibleWidget) ParentObject() *qtcore.QObject /*444 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget12parentObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:88
// index:0
// Protected
// void addControllingSignal(const class QString &)
func (this *QAccessibleWidget) AddControllingSignal(signal *qtcore.QString) {
	var convArg0 = signal.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleWidget20addControllingSignalERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
