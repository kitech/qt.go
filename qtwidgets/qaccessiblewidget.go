package qtwidgets

// /usr/include/qt/QtWidgets/qaccessiblewidget.h
// #include <qaccessiblewidget.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 134
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
func (this *QAccessibleWidget) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQAccessibleWidgetFromPointer(cthis unsafe.Pointer) *QAccessibleWidget {
	return &QAccessibleWidget{&qtrt.CObject{cthis}}
}
func (*QAccessibleWidget) NewFromPointer(cthis unsafe.Pointer) *QAccessibleWidget {
	return NewQAccessibleWidgetFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAccessibleWidget(QWidget *, QAccessible::Role, const QString &)
func NewQAccessibleWidget(o *QWidget /*777 QWidget **/, r int, name *qtcore.QString) *QAccessibleWidget {
	var convArg0 = o.GetCthis()
	var convArg2 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleWidgetC2EP7QWidgetN11QAccessible4RoleERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0, r, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QAccessibleWidget) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QWindow * window()
func (this *QAccessibleWidget) Window() *qtgui.QWindow /*777 QWindow **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget6windowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int childCount()
func (this *QAccessibleWidget) ChildCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget10childCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int indexOfChild(const QAccessibleInterface *)
func (this *QAccessibleWidget) IndexOfChild(child *qtgui.QAccessibleInterface /*777 const QAccessibleInterface **/) int {
	var convArg0 = child.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget12indexOfChildEPK20QAccessibleInterface", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * focusChild()
func (this *QAccessibleWidget) FocusChild() *qtgui.QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget10focusChildEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect rect()
func (this *QAccessibleWidget) Rect() *qtcore.QRect /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget4rectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * parent()
func (this *QAccessibleWidget) Parent() *qtgui.QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget6parentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:68
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * child(int)
func (this *QAccessibleWidget) Child(index int) *qtgui.QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget5childEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString text(QAccessible::Text)
func (this *QAccessibleWidget) Text(t int) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget4textEN11QAccessible4TextE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), t)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QAccessible::Role role()
func (this *QAccessibleWidget) Role() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget4roleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessible::State state()
func (this *QAccessibleWidget) State() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget5stateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QColor foregroundColor()
func (this *QAccessibleWidget) ForegroundColor() *qtgui.QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget15foregroundColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QColor backgroundColor()
func (this *QAccessibleWidget) BackgroundColor() *qtgui.QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget15backgroundColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] void * interface_cast(QAccessible::InterfaceType)
func (this *QAccessibleWidget) Interface_cast(t int) unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleWidget14interface_castEN11QAccessible13InterfaceTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), t)
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void doAction(const QString &)
func (this *QAccessibleWidget) DoAction(actionName *qtcore.QString) {
	var convArg0 = actionName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleWidget8doActionERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:84
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleWidget()
func DeleteQAccessibleWidget(*QAccessibleWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:85
// index:0
// Protected Visibility=Default Availability=Available
// [8] QWidget * widget()
func (this *QAccessibleWidget) Widget() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget6widgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:86
// index:0
// Protected Visibility=Default Availability=Available
// [8] QObject * parentObject()
func (this *QAccessibleWidget) ParentObject() *qtcore.QObject /*777 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget12parentObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:88
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void addControllingSignal(const QString &)
func (this *QAccessibleWidget) AddControllingSignal(signal *qtcore.QString) {
	var convArg0 = signal.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleWidget20addControllingSignalERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
