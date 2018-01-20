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

// /usr/include/qt/QtGui/qkeysequence.h:156
// index:0
// void QKeySequence()
func NewQKeySequence() *QKeySequence {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequenceC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQKeySequenceFromPointer(cthis)
	return gothis
}
func NewQKeySequenceFromPointer(cthis unsafe.Pointer) *QKeySequence {
	return &QKeySequence{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qkeysequence.h:157
// index:1
// void QKeySequence(const class QString &, enum QKeySequence::SequenceFormat)
func NewQKeySequence_1(key unsafe.Pointer, format int) *QKeySequence {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequenceC2ERK7QStringNS_14SequenceFormatE", ffiqt.FFI_TYPE_VOID, cthis, key, &format)
	gopp.ErrPrint(err, rv)
	gothis := NewQKeySequenceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qkeysequence.h:158
// index:2
// void QKeySequence(int, int, int, int)
func NewQKeySequence_2(k1 int, k2 int, k3 int, k4 int) *QKeySequence {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequenceC2Eiiii", ffiqt.FFI_TYPE_VOID, cthis, &k1, &k2, &k3, &k4)
	gopp.ErrPrint(err, rv)
	gothis := NewQKeySequenceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qkeysequence.h:160
// index:3
// void QKeySequence(enum QKeySequence::StandardKey)
func NewQKeySequence_3(key int) *QKeySequence {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequenceC2ENS_11StandardKeyE", ffiqt.FFI_TYPE_VOID, cthis, &key)
	gopp.ErrPrint(err, rv)
	gothis := NewQKeySequenceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qkeysequence.h:161
// index:0
// void ~QKeySequence()
func DeleteQKeySequence(*QKeySequence) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequenceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qkeysequence.h:163
// index:0
// int count()
func (this *QKeySequence) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QKeySequence5countEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qkeysequence.h:164
// index:0
// bool isEmpty()
func (this *QKeySequence) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QKeySequence7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qkeysequence.h:172
// index:0
// QString toString(enum QKeySequence::SequenceFormat)
func (this *QKeySequence) ToString(format int) {
	// 0: (, format QKeySequence::SequenceFormat), (&format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QKeySequence8toStringENS_14SequenceFormatE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qkeysequence.h:173
// index:0
// static
// QKeySequence fromString(const class QString &, enum QKeySequence::SequenceFormat)
func (this *QKeySequence) FromString(str unsafe.Pointer, format int) {
	// 0: (str const QString &, format QKeySequence::SequenceFormat), (str, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequence10fromStringERK7QStringNS_14SequenceFormatE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QKeySequence_FromString(str unsafe.Pointer, format int) {
	// 0: (str const QString &, format QKeySequence::SequenceFormat), (str, format)
	var nilthis *QKeySequence
	nilthis.FromString(str, format)
}

// /usr/include/qt/QtGui/qkeysequence.h:175
// index:0
// static
// QList<QKeySequence> listFromString(const class QString &, enum QKeySequence::SequenceFormat)
func (this *QKeySequence) ListFromString(str unsafe.Pointer, format int) {
	// 0: (str const QString &, format QKeySequence::SequenceFormat), (str, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequence14listFromStringERK7QStringNS_14SequenceFormatE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QKeySequence_ListFromString(str unsafe.Pointer, format int) {
	// 0: (str const QString &, format QKeySequence::SequenceFormat), (str, format)
	var nilthis *QKeySequence
	nilthis.ListFromString(str, format)
}

// /usr/include/qt/QtGui/qkeysequence.h:178
// index:0
// QKeySequence::SequenceMatch matches(const class QKeySequence &)
func (this *QKeySequence) Matches(seq unsafe.Pointer) {
	// 0: (, seq const QKeySequence &), (seq)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QKeySequence7matchesERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), seq)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qkeysequence.h:179
// index:0
// static
// QKeySequence mnemonic(const class QString &)
func (this *QKeySequence) Mnemonic(text unsafe.Pointer) {
	// 0: (text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequence8mnemonicERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QKeySequence_Mnemonic(text unsafe.Pointer) {
	// 0: (text const QString &), (text)
	var nilthis *QKeySequence
	nilthis.Mnemonic(text)
}

// /usr/include/qt/QtGui/qkeysequence.h:180
// index:0
// static
// QList<QKeySequence> keyBindings(enum QKeySequence::StandardKey)
func (this *QKeySequence) KeyBindings(key int) {
	// 0: (key QKeySequence::StandardKey), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequence11keyBindingsENS_11StandardKeyE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QKeySequence_KeyBindings(key int) {
	// 0: (key QKeySequence::StandardKey), (key)
	var nilthis *QKeySequence
	nilthis.KeyBindings(key)
}

// /usr/include/qt/QtGui/qkeysequence.h:192
// index:0
// inline
// void swap(class QKeySequence &)
func (this *QKeySequence) Swap(other unsafe.Pointer) {
	// 0: (, other QKeySequence &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequence4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qkeysequence.h:205
// index:0
// bool isDetached()
func (this *QKeySequence) IsDetached() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QKeySequence10isDetachedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
