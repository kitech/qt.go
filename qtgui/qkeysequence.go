package qtgui

// /usr/include/qt/QtGui/qkeysequence.h
// #include <qkeysequence.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 30
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QKeySequence) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQKeySequenceFromPointer(cthis unsafe.Pointer) *QKeySequence {
	return &QKeySequence{&qtrt.CObject{cthis}}
}
func (*QKeySequence) NewFromPointer(cthis unsafe.Pointer) *QKeySequence {
	return NewQKeySequenceFromPointer(cthis)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequenceC2ERK7QStringNS_14SequenceFormatE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, format)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequenceC2Eiiii", ffiqt.FFI_TYPE_VOID, cthis, k1, k2, k3, k4)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequenceC2ENS_11StandardKeyE", ffiqt.FFI_TYPE_VOID, cthis, key)
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
func (this *QKeySequence) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QKeySequence5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qkeysequence.h:164
// index:0
// Public
// bool isEmpty()
func (this *QKeySequence) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QKeySequence7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qkeysequence.h:172
// index:0
// Public
// QString toString(enum QKeySequence::SequenceFormat)
func (this *QKeySequence) ToString(format int) *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QKeySequence8toStringENS_14SequenceFormatE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qkeysequence.h:173
// index:0
// Public static
// QKeySequence fromString(const class QString &, enum QKeySequence::SequenceFormat)
func (this *QKeySequence) FromString(str *qtcore.QString, format int) *QKeySequence /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequence10fromStringERK7QStringNS_14SequenceFormatE", ffiqt.FFI_TYPE_POINTER, str, format)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QKeySequence_FromString(str *qtcore.QString, format int) *QKeySequence /*123*/ {
	var nilthis *QKeySequence
	rv := nilthis.FromString(str, format)
	return rv
}

// /usr/include/qt/QtGui/qkeysequence.h:178
// index:0
// Public
// QKeySequence::SequenceMatch matches(const class QKeySequence &)
func (this *QKeySequence) Matches(seq *QKeySequence) int {
	var convArg0 = seq.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QKeySequence7matchesERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qkeysequence.h:179
// index:0
// Public static
// QKeySequence mnemonic(const class QString &)
func (this *QKeySequence) Mnemonic(text *qtcore.QString) *QKeySequence /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QKeySequence8mnemonicERK7QString", ffiqt.FFI_TYPE_POINTER, text)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QKeySequence_Mnemonic(text *qtcore.QString) *QKeySequence /*123*/ {
	var nilthis *QKeySequence
	rv := nilthis.Mnemonic(text)
	return rv
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
func (this *QKeySequence) IsDetached() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QKeySequence10isDetachedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QKeySequence__StandardKey = int

const QKeySequence__UnknownKey QKeySequence__StandardKey = 0
const QKeySequence__HelpContents QKeySequence__StandardKey = 1
const QKeySequence__WhatsThis QKeySequence__StandardKey = 2
const QKeySequence__Open QKeySequence__StandardKey = 3
const QKeySequence__Close QKeySequence__StandardKey = 4
const QKeySequence__Save QKeySequence__StandardKey = 5
const QKeySequence__New QKeySequence__StandardKey = 6
const QKeySequence__Delete QKeySequence__StandardKey = 7
const QKeySequence__Cut QKeySequence__StandardKey = 8
const QKeySequence__Copy QKeySequence__StandardKey = 9
const QKeySequence__Paste QKeySequence__StandardKey = 10
const QKeySequence__Undo QKeySequence__StandardKey = 11
const QKeySequence__Redo QKeySequence__StandardKey = 12
const QKeySequence__Back QKeySequence__StandardKey = 13
const QKeySequence__Forward QKeySequence__StandardKey = 14
const QKeySequence__Refresh QKeySequence__StandardKey = 15
const QKeySequence__ZoomIn QKeySequence__StandardKey = 16
const QKeySequence__ZoomOut QKeySequence__StandardKey = 17
const QKeySequence__Print QKeySequence__StandardKey = 18
const QKeySequence__AddTab QKeySequence__StandardKey = 19
const QKeySequence__NextChild QKeySequence__StandardKey = 20
const QKeySequence__PreviousChild QKeySequence__StandardKey = 21
const QKeySequence__Find QKeySequence__StandardKey = 22
const QKeySequence__FindNext QKeySequence__StandardKey = 23
const QKeySequence__FindPrevious QKeySequence__StandardKey = 24
const QKeySequence__Replace QKeySequence__StandardKey = 25
const QKeySequence__SelectAll QKeySequence__StandardKey = 26
const QKeySequence__Bold QKeySequence__StandardKey = 27
const QKeySequence__Italic QKeySequence__StandardKey = 28
const QKeySequence__Underline QKeySequence__StandardKey = 29
const QKeySequence__MoveToNextChar QKeySequence__StandardKey = 30
const QKeySequence__MoveToPreviousChar QKeySequence__StandardKey = 31
const QKeySequence__MoveToNextWord QKeySequence__StandardKey = 32
const QKeySequence__MoveToPreviousWord QKeySequence__StandardKey = 33
const QKeySequence__MoveToNextLine QKeySequence__StandardKey = 34
const QKeySequence__MoveToPreviousLine QKeySequence__StandardKey = 35
const QKeySequence__MoveToNextPage QKeySequence__StandardKey = 36
const QKeySequence__MoveToPreviousPage QKeySequence__StandardKey = 37
const QKeySequence__MoveToStartOfLine QKeySequence__StandardKey = 38
const QKeySequence__MoveToEndOfLine QKeySequence__StandardKey = 39
const QKeySequence__MoveToStartOfBlock QKeySequence__StandardKey = 40
const QKeySequence__MoveToEndOfBlock QKeySequence__StandardKey = 41
const QKeySequence__MoveToStartOfDocument QKeySequence__StandardKey = 42
const QKeySequence__MoveToEndOfDocument QKeySequence__StandardKey = 43
const QKeySequence__SelectNextChar QKeySequence__StandardKey = 44
const QKeySequence__SelectPreviousChar QKeySequence__StandardKey = 45
const QKeySequence__SelectNextWord QKeySequence__StandardKey = 46
const QKeySequence__SelectPreviousWord QKeySequence__StandardKey = 47
const QKeySequence__SelectNextLine QKeySequence__StandardKey = 48
const QKeySequence__SelectPreviousLine QKeySequence__StandardKey = 49
const QKeySequence__SelectNextPage QKeySequence__StandardKey = 50
const QKeySequence__SelectPreviousPage QKeySequence__StandardKey = 51
const QKeySequence__SelectStartOfLine QKeySequence__StandardKey = 52
const QKeySequence__SelectEndOfLine QKeySequence__StandardKey = 53
const QKeySequence__SelectStartOfBlock QKeySequence__StandardKey = 54
const QKeySequence__SelectEndOfBlock QKeySequence__StandardKey = 55
const QKeySequence__SelectStartOfDocument QKeySequence__StandardKey = 56
const QKeySequence__SelectEndOfDocument QKeySequence__StandardKey = 57
const QKeySequence__DeleteStartOfWord QKeySequence__StandardKey = 58
const QKeySequence__DeleteEndOfWord QKeySequence__StandardKey = 59
const QKeySequence__DeleteEndOfLine QKeySequence__StandardKey = 60
const QKeySequence__InsertParagraphSeparator QKeySequence__StandardKey = 61
const QKeySequence__InsertLineSeparator QKeySequence__StandardKey = 62
const QKeySequence__SaveAs QKeySequence__StandardKey = 63
const QKeySequence__Preferences QKeySequence__StandardKey = 64
const QKeySequence__Quit QKeySequence__StandardKey = 65
const QKeySequence__FullScreen QKeySequence__StandardKey = 66
const QKeySequence__Deselect QKeySequence__StandardKey = 67
const QKeySequence__DeleteCompleteLine QKeySequence__StandardKey = 68
const QKeySequence__Backspace QKeySequence__StandardKey = 69
const QKeySequence__Cancel QKeySequence__StandardKey = 70

type QKeySequence__SequenceFormat = int

const QKeySequence__NativeText QKeySequence__SequenceFormat = 0
const QKeySequence__PortableText QKeySequence__SequenceFormat = 1

type QKeySequence__SequenceMatch = int

const QKeySequence__NoMatch QKeySequence__SequenceMatch = 0
const QKeySequence__PartialMatch QKeySequence__SequenceMatch = 1
const QKeySequence__ExactMatch QKeySequence__SequenceMatch = 2

//  body block end
