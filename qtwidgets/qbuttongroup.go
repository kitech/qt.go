package qtwidgets

// /usr/include/qt/QtWidgets/qbuttongroup.h
// #include <qbuttongroup.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
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

type QButtonGroup struct {
	*qtcore.QObject
}

func (this *QButtonGroup) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QButtonGroup) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQButtonGroupFromPointer(cthis unsafe.Pointer) *QButtonGroup {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QButtonGroup{bcthis0}
}
func (*QButtonGroup) NewFromPointer(cthis unsafe.Pointer) *QButtonGroup {
	return NewQButtonGroupFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QButtonGroup) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QButtonGroup10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QButtonGroup(QObject *)
func NewQButtonGroup(parent *qtcore.QObject /*777 QObject **/) *QButtonGroup {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QButtonGroupC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQButtonGroupFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QButtonGroup()
func DeleteQButtonGroup(this *QButtonGroup) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QButtonGroupD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExclusive(_Bool)
func (this *QButtonGroup) SetExclusive(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QButtonGroup12setExclusiveEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:64
// index:0
// Public Visibility=Default Availability=Available
// [1] bool exclusive()
func (this *QButtonGroup) Exclusive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QButtonGroup9exclusiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addButton(QAbstractButton *, int)
func (this *QButtonGroup) AddButton(arg0 *QAbstractButton /*777 QAbstractButton **/, id int) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QButtonGroup9addButtonEP15QAbstractButtoni", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeButton(QAbstractButton *)
func (this *QButtonGroup) RemoveButton(arg0 *QAbstractButton /*777 QAbstractButton **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QButtonGroup12removeButtonEP15QAbstractButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractButton * checkedButton()
func (this *QButtonGroup) CheckedButton() *QAbstractButton /*777 QAbstractButton **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QButtonGroup13checkedButtonEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractButton * button(int)
func (this *QButtonGroup) Button(id int) *QAbstractButton /*777 QAbstractButton **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QButtonGroup6buttonEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setId(QAbstractButton *, int)
func (this *QButtonGroup) SetId(button *QAbstractButton /*777 QAbstractButton **/, id int) {
	var convArg0 = button.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QButtonGroup5setIdEP15QAbstractButtoni", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int id(QAbstractButton *)
func (this *QButtonGroup) Id(button *QAbstractButton /*777 QAbstractButton **/) int {
	var convArg0 = button.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QButtonGroup2idEP15QAbstractButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] int checkedId()
func (this *QButtonGroup) CheckedId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QButtonGroup9checkedIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void buttonClicked(QAbstractButton *)
func (this *QButtonGroup) ButtonClicked(arg0 *QAbstractButton /*777 QAbstractButton **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QButtonGroup13buttonClickedEP15QAbstractButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:81
// index:1
// Public Visibility=Default Availability=Available
// [-2] void buttonClicked(int)
func (this *QButtonGroup) ButtonClicked_1(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QButtonGroup13buttonClickedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void buttonPressed(QAbstractButton *)
func (this *QButtonGroup) ButtonPressed(arg0 *QAbstractButton /*777 QAbstractButton **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QButtonGroup13buttonPressedEP15QAbstractButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:83
// index:1
// Public Visibility=Default Availability=Available
// [-2] void buttonPressed(int)
func (this *QButtonGroup) ButtonPressed_1(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QButtonGroup13buttonPressedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void buttonReleased(QAbstractButton *)
func (this *QButtonGroup) ButtonReleased(arg0 *QAbstractButton /*777 QAbstractButton **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QButtonGroup14buttonReleasedEP15QAbstractButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:85
// index:1
// Public Visibility=Default Availability=Available
// [-2] void buttonReleased(int)
func (this *QButtonGroup) ButtonReleased_1(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QButtonGroup14buttonReleasedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void buttonToggled(QAbstractButton *, _Bool)
func (this *QButtonGroup) ButtonToggled(arg0 *QAbstractButton /*777 QAbstractButton **/, arg1 bool) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QButtonGroup13buttonToggledEP15QAbstractButtonb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:87
// index:1
// Public Visibility=Default Availability=Available
// [-2] void buttonToggled(int, _Bool)
func (this *QButtonGroup) ButtonToggled_1(arg0 int, arg1 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QButtonGroup13buttonToggledEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
}

//  body block end
