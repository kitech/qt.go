//  header block begin
// /usr/include/qt/QtGui/qkeysequence.h
// #include <qkeysequence.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
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
type QKeySequence struct {
	*qtrt.CObject
}

func (this *QKeySequence) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQKeySequenceFromPointer(cthis unsafe.Pointer) *QKeySequence {
	return &QKeySequence{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qkeysequence.h:156
// index:0
// Public
// void QKeySequence()
func NewQKeySequence() *QKeySequence {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequenceC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQKeySequenceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qkeysequence.h:157
// index:1
// Public
// void QKeySequence(const class QString &, enum QKeySequence::SequenceFormat)
func NewQKeySequence_1(key *qtcore.QString, format int) *QKeySequence {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequenceC2ERK7QStringNS_14SequenceFormatE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &format)
	gopp.ErrPrint(err, rv)
	gothis := NewQKeySequenceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qkeysequence.h:158
// index:2
// Public
// void QKeySequence(int, int, int, int)
func NewQKeySequence_2(k1 int, k2 int, k3 int, k4 int) *QKeySequence {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequenceC2Eiiii", ffiqt.FFI_TYPE_VOID, cthis, &k1, &k2, &k3, &k4)
	gopp.ErrPrint(err, rv)
	gothis := NewQKeySequenceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qkeysequence.h:160
// index:3
// Public
// void QKeySequence(enum QKeySequence::StandardKey)
func NewQKeySequence_3(key int) *QKeySequence {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequenceC2ENS_11StandardKeyE", ffiqt.FFI_TYPE_VOID, cthis, &key)
	gopp.ErrPrint(err, rv)
	gothis := NewQKeySequenceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qkeysequence.h:161
// index:0
// Public
// void ~QKeySequence()
func DeleteQKeySequence(*QKeySequence) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequenceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qkeysequence.h:163
// index:0
// Public
// int count()
func (this *QKeySequence) Count() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QKeySequence5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qkeysequence.h:164
// index:0
// Public
// bool isEmpty()
func (this *QKeySequence) IsEmpty() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QKeySequence7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qkeysequence.h:172
// index:0
// Public
// QString toString(enum QKeySequence::SequenceFormat)
func (this *QKeySequence) ToString(format int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QKeySequence8toStringENS_14SequenceFormatE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &format)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qkeysequence.h:173
// index:0
// Public static
// QKeySequence fromString(const class QString &, enum QKeySequence::SequenceFormat)
func (this *QKeySequence) FromString(str *qtcore.QString, format int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequence10fromStringERK7QStringNS_14SequenceFormatE", ffiqt.FFI_TYPE_POINTER, str, format)
	gopp.ErrPrint(err, rv)
	return rv
}
func QKeySequence_FromString(str *qtcore.QString, format int) {
	var nilthis *QKeySequence
	nilthis.FromString(str, format)
}

// /usr/include/qt/QtGui/qkeysequence.h:175
// index:0
// Public static
// QList<QKeySequence> listFromString(const class QString &, enum QKeySequence::SequenceFormat)
func (this *QKeySequence) ListFromString(str *qtcore.QString, format int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequence14listFromStringERK7QStringNS_14SequenceFormatE", ffiqt.FFI_TYPE_POINTER, str, format)
	gopp.ErrPrint(err, rv)
	return rv
}
func QKeySequence_ListFromString(str *qtcore.QString, format int) {
	var nilthis *QKeySequence
	nilthis.ListFromString(str, format)
}

// /usr/include/qt/QtGui/qkeysequence.h:178
// index:0
// Public
// QKeySequence::SequenceMatch matches(const class QKeySequence &)
func (this *QKeySequence) Matches(seq *QKeySequence) interface{} {
	var convArg0 = seq.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QKeySequence7matchesERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qkeysequence.h:179
// index:0
// Public static
// QKeySequence mnemonic(const class QString &)
func (this *QKeySequence) Mnemonic(text *qtcore.QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequence8mnemonicERK7QString", ffiqt.FFI_TYPE_POINTER, text)
	gopp.ErrPrint(err, rv)
	return rv
}
func QKeySequence_Mnemonic(text *qtcore.QString) {
	var nilthis *QKeySequence
	nilthis.Mnemonic(text)
}

// /usr/include/qt/QtGui/qkeysequence.h:180
// index:0
// Public static
// QList<QKeySequence> keyBindings(enum QKeySequence::StandardKey)
func (this *QKeySequence) KeyBindings(key int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequence11keyBindingsENS_11StandardKeyE", ffiqt.FFI_TYPE_POINTER, key)
	gopp.ErrPrint(err, rv)
	return rv
}
func QKeySequence_KeyBindings(key int) {
	var nilthis *QKeySequence
	nilthis.KeyBindings(key)
}

// /usr/include/qt/QtGui/qkeysequence.h:192
// index:0
// Public inline
// void swap(class QKeySequence &)
func (this *QKeySequence) Swap(other *QKeySequence) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequence4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qkeysequence.h:205
// index:0
// Public
// bool isDetached()
func (this *QKeySequence) IsDetached() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QKeySequence10isDetachedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
