package qtwidgets

// /usr/include/qt/QtWidgets/qsizepolicy.h
// #include <qsizepolicy.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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
type QSizePolicy struct {
	*qtrt.CObject
}

func (this *QSizePolicy) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSizePolicy) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQSizePolicyFromPointer(cthis unsafe.Pointer) *QSizePolicy {
	return &QSizePolicy{&qtrt.CObject{cthis}}
}
func (*QSizePolicy) NewFromPointer(cthis unsafe.Pointer) *QSizePolicy {
	return NewQSizePolicyFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:113
// index:0
// Public inline
// void QSizePolicy()
func NewQSizePolicy() *QSizePolicy {
	cthis := qtrt.Calloc(1, 256) // 4
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSizePolicyC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQSizePolicyFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:116
// index:1
// Public inline
// void QSizePolicy(enum QSizePolicy::Policy, enum QSizePolicy::Policy, enum QSizePolicy::ControlType)
func NewQSizePolicy_1(horizontal int, vertical int, type_ int) *QSizePolicy {
	cthis := qtrt.Calloc(1, 256) // 4
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSizePolicyC2ENS_6PolicyES0_NS_11ControlTypeE", ffiqt.FFI_TYPE_VOID, cthis, horizontal, vertical, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQSizePolicyFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:128
// index:0
// Public inline
// QSizePolicy::Policy horizontalPolicy()
func (this *QSizePolicy) HorizontalPolicy() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSizePolicy16horizontalPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:129
// index:0
// Public inline
// QSizePolicy::Policy verticalPolicy()
func (this *QSizePolicy) VerticalPolicy() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSizePolicy14verticalPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:130
// index:0
// Public
// QSizePolicy::ControlType controlType()
func (this *QSizePolicy) ControlType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSizePolicy11controlTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:132
// index:0
// Public inline
// void setHorizontalPolicy(enum QSizePolicy::Policy)
func (this *QSizePolicy) SetHorizontalPolicy(d int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSizePolicy19setHorizontalPolicyENS_6PolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), d)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:133
// index:0
// Public inline
// void setVerticalPolicy(enum QSizePolicy::Policy)
func (this *QSizePolicy) SetVerticalPolicy(d int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSizePolicy17setVerticalPolicyENS_6PolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), d)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:134
// index:0
// Public
// void setControlType(enum QSizePolicy::ControlType)
func (this *QSizePolicy) SetControlType(type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSizePolicy14setControlTypeENS_11ControlTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:136
// index:0
// Public inline
// Qt::Orientations expandingDirections()
func (this *QSizePolicy) ExpandingDirections() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSizePolicy19expandingDirectionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:141
// index:0
// Public inline
// void setHeightForWidth(_Bool)
func (this *QSizePolicy) SetHeightForWidth(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSizePolicy17setHeightForWidthEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:142
// index:0
// Public inline
// bool hasHeightForWidth()
func (this *QSizePolicy) HasHeightForWidth() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSizePolicy17hasHeightForWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:143
// index:0
// Public inline
// void setWidthForHeight(_Bool)
func (this *QSizePolicy) SetWidthForHeight(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSizePolicy17setWidthForHeightEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:144
// index:0
// Public inline
// bool hasWidthForHeight()
func (this *QSizePolicy) HasWidthForHeight() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSizePolicy17hasWidthForHeightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:153
// index:0
// Public inline
// int horizontalStretch()
func (this *QSizePolicy) HorizontalStretch() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSizePolicy17horizontalStretchEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:154
// index:0
// Public inline
// int verticalStretch()
func (this *QSizePolicy) VerticalStretch() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSizePolicy15verticalStretchEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:155
// index:0
// Public inline
// void setHorizontalStretch(int)
func (this *QSizePolicy) SetHorizontalStretch(stretchFactor int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSizePolicy20setHorizontalStretchEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), stretchFactor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:156
// index:0
// Public inline
// void setVerticalStretch(int)
func (this *QSizePolicy) SetVerticalStretch(stretchFactor int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSizePolicy18setVerticalStretchEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), stretchFactor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:158
// index:0
// Public inline
// bool retainSizeWhenHidden()
func (this *QSizePolicy) RetainSizeWhenHidden() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSizePolicy20retainSizeWhenHiddenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:159
// index:0
// Public inline
// void setRetainSizeWhenHidden(_Bool)
func (this *QSizePolicy) SetRetainSizeWhenHidden(retainSize bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSizePolicy23setRetainSizeWhenHiddenEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), retainSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:161
// index:0
// Public inline
// void transpose()
func (this *QSizePolicy) Transpose() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSizePolicy9transposeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:166
// index:0
// Public inline
// QSizePolicy transposed()
func (this *QSizePolicy) Transposed() *QSizePolicy /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSizePolicy10transposedEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSizePolicyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

type QSizePolicy__PolicyFlag = int

const QSizePolicy__GrowFlag QSizePolicy__PolicyFlag = 1
const QSizePolicy__ExpandFlag QSizePolicy__PolicyFlag = 2
const QSizePolicy__ShrinkFlag QSizePolicy__PolicyFlag = 4
const QSizePolicy__IgnoreFlag QSizePolicy__PolicyFlag = 8

type QSizePolicy__Policy = int

const QSizePolicy__Fixed QSizePolicy__Policy = 0
const QSizePolicy__Minimum QSizePolicy__Policy = 1
const QSizePolicy__Maximum QSizePolicy__Policy = 4
const QSizePolicy__Preferred QSizePolicy__Policy = 5
const QSizePolicy__MinimumExpanding QSizePolicy__Policy = 3
const QSizePolicy__Expanding QSizePolicy__Policy = 7
const QSizePolicy__Ignored QSizePolicy__Policy = 13

type QSizePolicy__ControlType = int

const QSizePolicy__DefaultType QSizePolicy__ControlType = 1
const QSizePolicy__ButtonBox QSizePolicy__ControlType = 2
const QSizePolicy__CheckBox QSizePolicy__ControlType = 4
const QSizePolicy__ComboBox QSizePolicy__ControlType = 8
const QSizePolicy__Frame QSizePolicy__ControlType = 16
const QSizePolicy__GroupBox QSizePolicy__ControlType = 32
const QSizePolicy__Label QSizePolicy__ControlType = 64
const QSizePolicy__Line QSizePolicy__ControlType = 128
const QSizePolicy__LineEdit QSizePolicy__ControlType = 256
const QSizePolicy__PushButton QSizePolicy__ControlType = 512
const QSizePolicy__RadioButton QSizePolicy__ControlType = 1024
const QSizePolicy__Slider QSizePolicy__ControlType = 2048
const QSizePolicy__SpinBox QSizePolicy__ControlType = 4096
const QSizePolicy__TabWidget QSizePolicy__ControlType = 8192
const QSizePolicy__ToolButton QSizePolicy__ControlType = 16384

//  body block end
