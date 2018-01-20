//  header block begin
// /usr/include/qt/QtWidgets/qbuttongroup.h
// #include <qbuttongroup.h>
// #include <QtWidgets>
package qtwidgets

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
type QButtonGroup struct {
	*qtcore.QObject
}

func (this *QButtonGroup) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQButtonGroupFromPointer(cthis unsafe.Pointer) *QButtonGroup {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QButtonGroup{bcthis0}
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:56
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QButtonGroup) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QButtonGroup10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:60
// index:0
// Public
// void QButtonGroup(class QObject *)
func NewQButtonGroup(parent unsafe.Pointer) *QButtonGroup {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QButtonGroupC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQButtonGroupFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:61
// index:0
// Public virtual
// void ~QButtonGroup()
func DeleteQButtonGroup(*QButtonGroup) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QButtonGroupD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:63
// index:0
// Public
// void setExclusive(_Bool)
func (this *QButtonGroup) SetExclusive(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QButtonGroup12setExclusiveEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:64
// index:0
// Public
// bool exclusive()
func (this *QButtonGroup) Exclusive() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QButtonGroup9exclusiveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:66
// index:0
// Public
// void addButton(class QAbstractButton *, int)
func (this *QButtonGroup) AddButton(arg0 unsafe.Pointer, id int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QButtonGroup9addButtonEP15QAbstractButtoni", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, &id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:67
// index:0
// Public
// void removeButton(class QAbstractButton *)
func (this *QButtonGroup) RemoveButton(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QButtonGroup12removeButtonEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:69
// index:0
// Public
// QList<QAbstractButton *> buttons()
func (this *QButtonGroup) Buttons() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QButtonGroup7buttonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:71
// index:0
// Public
// QAbstractButton * checkedButton()
func (this *QButtonGroup) CheckedButton() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QButtonGroup13checkedButtonEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:74
// index:0
// Public
// QAbstractButton * button(int)
func (this *QButtonGroup) Button(id int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QButtonGroup6buttonEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &id)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:75
// index:0
// Public
// void setId(class QAbstractButton *, int)
func (this *QButtonGroup) SetId(button unsafe.Pointer, id int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QButtonGroup5setIdEP15QAbstractButtoni", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button, &id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:76
// index:0
// Public
// int id(class QAbstractButton *)
func (this *QButtonGroup) Id(button unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QButtonGroup2idEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:77
// index:0
// Public
// int checkedId()
func (this *QButtonGroup) CheckedId() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QButtonGroup9checkedIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:80
// index:0
// Public
// void buttonClicked(class QAbstractButton *)
func (this *QButtonGroup) ButtonClicked(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QButtonGroup13buttonClickedEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:81
// index:1
// Public
// void buttonClicked(int)
func (this *QButtonGroup) ButtonClicked_1(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QButtonGroup13buttonClickedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:82
// index:0
// Public
// void buttonPressed(class QAbstractButton *)
func (this *QButtonGroup) ButtonPressed(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QButtonGroup13buttonPressedEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:83
// index:1
// Public
// void buttonPressed(int)
func (this *QButtonGroup) ButtonPressed_1(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QButtonGroup13buttonPressedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:84
// index:0
// Public
// void buttonReleased(class QAbstractButton *)
func (this *QButtonGroup) ButtonReleased(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QButtonGroup14buttonReleasedEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:85
// index:1
// Public
// void buttonReleased(int)
func (this *QButtonGroup) ButtonReleased_1(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QButtonGroup14buttonReleasedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:86
// index:0
// Public
// void buttonToggled(class QAbstractButton *, _Bool)
func (this *QButtonGroup) ButtonToggled(arg0 unsafe.Pointer, arg1 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QButtonGroup13buttonToggledEP15QAbstractButtonb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, &arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qbuttongroup.h:87
// index:1
// Public
// void buttonToggled(int, _Bool)
func (this *QButtonGroup) ButtonToggled_1(arg0 int, arg1 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QButtonGroup13buttonToggledEib", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &arg1)
	gopp.ErrPrint(err, rv)
}

//  body block end
