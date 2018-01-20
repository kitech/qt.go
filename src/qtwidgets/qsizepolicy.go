//  header block begin
// /usr/include/qt/QtWidgets/qsizepolicy.h
// #include <qsizepolicy.h>
// #include <QtWidgets>
package qtwidgets

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
type QSizePolicy struct {
	*qtrt.CObject
}

func (this *QSizePolicy) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQSizePolicyFromPointer(cthis unsafe.Pointer) *QSizePolicy {
	return &QSizePolicy{&qtrt.CObject{cthis}}
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSizePolicyC2ENS_6PolicyES0_NS_11ControlTypeE", ffiqt.FFI_TYPE_VOID, cthis, &horizontal, &vertical, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQSizePolicyFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:128
// index:0
// Public inline
// QSizePolicy::Policy horizontalPolicy()
func (this *QSizePolicy) HorizontalPolicy() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSizePolicy16horizontalPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:129
// index:0
// Public inline
// QSizePolicy::Policy verticalPolicy()
func (this *QSizePolicy) VerticalPolicy() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSizePolicy14verticalPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:130
// index:0
// Public
// QSizePolicy::ControlType controlType()
func (this *QSizePolicy) ControlType() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSizePolicy11controlTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:132
// index:0
// Public inline
// void setHorizontalPolicy(enum QSizePolicy::Policy)
func (this *QSizePolicy) SetHorizontalPolicy(d int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSizePolicy19setHorizontalPolicyENS_6PolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &d)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:133
// index:0
// Public inline
// void setVerticalPolicy(enum QSizePolicy::Policy)
func (this *QSizePolicy) SetVerticalPolicy(d int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSizePolicy17setVerticalPolicyENS_6PolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &d)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:134
// index:0
// Public
// void setControlType(enum QSizePolicy::ControlType)
func (this *QSizePolicy) SetControlType(type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSizePolicy14setControlTypeENS_11ControlTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:136
// index:0
// Public inline
// Qt::Orientations expandingDirections()
func (this *QSizePolicy) ExpandingDirections() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSizePolicy19expandingDirectionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:141
// index:0
// Public inline
// void setHeightForWidth(_Bool)
func (this *QSizePolicy) SetHeightForWidth(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSizePolicy17setHeightForWidthEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:142
// index:0
// Public inline
// bool hasHeightForWidth()
func (this *QSizePolicy) HasHeightForWidth() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSizePolicy17hasHeightForWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:143
// index:0
// Public inline
// void setWidthForHeight(_Bool)
func (this *QSizePolicy) SetWidthForHeight(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSizePolicy17setWidthForHeightEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:144
// index:0
// Public inline
// bool hasWidthForHeight()
func (this *QSizePolicy) HasWidthForHeight() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSizePolicy17hasWidthForHeightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:153
// index:0
// Public inline
// int horizontalStretch()
func (this *QSizePolicy) HorizontalStretch() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSizePolicy17horizontalStretchEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:154
// index:0
// Public inline
// int verticalStretch()
func (this *QSizePolicy) VerticalStretch() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSizePolicy15verticalStretchEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:155
// index:0
// Public inline
// void setHorizontalStretch(int)
func (this *QSizePolicy) SetHorizontalStretch(stretchFactor int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSizePolicy20setHorizontalStretchEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &stretchFactor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:156
// index:0
// Public inline
// void setVerticalStretch(int)
func (this *QSizePolicy) SetVerticalStretch(stretchFactor int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSizePolicy18setVerticalStretchEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &stretchFactor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:158
// index:0
// Public inline
// bool retainSizeWhenHidden()
func (this *QSizePolicy) RetainSizeWhenHidden() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSizePolicy20retainSizeWhenHiddenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:159
// index:0
// Public inline
// void setRetainSizeWhenHidden(_Bool)
func (this *QSizePolicy) SetRetainSizeWhenHidden(retainSize bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSizePolicy23setRetainSizeWhenHiddenEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &retainSize)
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
func (this *QSizePolicy) Transposed() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSizePolicy10transposedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
