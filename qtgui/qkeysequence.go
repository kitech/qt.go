package qtgui

// /usr/include/qt/QtGui/qkeysequence.h
// #include <qkeysequence.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 48
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QKeySequence struct {
	*qtrt.CObject
}
type QKeySequence_ITF interface {
	QKeySequence_PTR() *QKeySequence
}

func (ptr *QKeySequence) QKeySequence_PTR() *QKeySequence { return ptr }

func (this *QKeySequence) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QKeySequence) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQKeySequenceFromPointer(cthis unsafe.Pointer) *QKeySequence {
	return &QKeySequence{&qtrt.CObject{cthis}}
}
func (*QKeySequence) NewFromPointer(cthis unsafe.Pointer) *QKeySequence {
	return NewQKeySequenceFromPointer(cthis)
}

// /usr/include/qt/QtGui/qkeysequence.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QKeySequence()

/*
Constructs an empty key sequence.
*/
func NewQKeySequence() *QKeySequence {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QKeySequenceC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQKeySequence)
	return gothis
}

// /usr/include/qt/QtGui/qkeysequence.h:157
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QKeySequence(const QString &, QKeySequence::SequenceFormat)

/*
Constructs an empty key sequence.
*/
func NewQKeySequence_1(key string, format int) *QKeySequence {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QKeySequenceC2ERK7QStringNS_14SequenceFormatE", qtrt.FFI_TYPE_POINTER, convArg0, format)
	qtrt.ErrPrint(err, rv)
	gothis := NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQKeySequence)
	return gothis
}

// /usr/include/qt/QtGui/qkeysequence.h:157
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QKeySequence(const QString &, QKeySequence::SequenceFormat)

/*
Constructs an empty key sequence.
*/
func NewQKeySequence_1_(key string) *QKeySequence {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QKeySequence::SequenceFormat=Enum, QKeySequence::SequenceFormat=Enum,
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QKeySequenceC2ERK7QStringNS_14SequenceFormatE", qtrt.FFI_TYPE_POINTER, convArg0, format)
	qtrt.ErrPrint(err, rv)
	gothis := NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQKeySequence)
	return gothis
}

// /usr/include/qt/QtGui/qkeysequence.h:158
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QKeySequence(int, int, int, int)

/*
Constructs an empty key sequence.
*/
func NewQKeySequence_2(k1 int, k2 int, k3 int, k4 int) *QKeySequence {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QKeySequenceC2Eiiii", qtrt.FFI_TYPE_POINTER, k1, k2, k3, k4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQKeySequence)
	return gothis
}

// /usr/include/qt/QtGui/qkeysequence.h:158
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QKeySequence(int, int, int, int)

/*
Constructs an empty key sequence.
*/
func NewQKeySequence_2_(k1 int) *QKeySequence {
	// arg: 1, int=Int, =Invalid,
	k2 := int(0)
	// arg: 2, int=Int, =Invalid,
	k3 := int(0)
	// arg: 3, int=Int, =Invalid,
	k4 := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QKeySequenceC2Eiiii", qtrt.FFI_TYPE_POINTER, k1, k2, k3, k4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQKeySequence)
	return gothis
}

// /usr/include/qt/QtGui/qkeysequence.h:158
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QKeySequence(int, int, int, int)

/*
Constructs an empty key sequence.
*/
func NewQKeySequence_2_1(k1 int, k2 int) *QKeySequence {
	// arg: 2, int=Int, =Invalid,
	k3 := int(0)
	// arg: 3, int=Int, =Invalid,
	k4 := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QKeySequenceC2Eiiii", qtrt.FFI_TYPE_POINTER, k1, k2, k3, k4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQKeySequence)
	return gothis
}

// /usr/include/qt/QtGui/qkeysequence.h:158
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QKeySequence(int, int, int, int)

/*
Constructs an empty key sequence.
*/
func NewQKeySequence_2_2(k1 int, k2 int, k3 int) *QKeySequence {
	// arg: 3, int=Int, =Invalid,
	k4 := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QKeySequenceC2Eiiii", qtrt.FFI_TYPE_POINTER, k1, k2, k3, k4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQKeySequence)
	return gothis
}

// /usr/include/qt/QtGui/qkeysequence.h:160
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QKeySequence(QKeySequence::StandardKey)

/*
Constructs an empty key sequence.
*/
func NewQKeySequence_3(key int) *QKeySequence {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QKeySequenceC2ENS_11StandardKeyE", qtrt.FFI_TYPE_POINTER, key)
	qtrt.ErrPrint(err, rv)
	gothis := NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQKeySequence)
	return gothis
}

// /usr/include/qt/QtGui/qkeysequence.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QKeySequence()

/*

 */
func DeleteQKeySequence(this *QKeySequence) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QKeySequenceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qkeysequence.h:163
// index:0
// Public Visibility=Default Availability=Available
// [4] int count() const

/*
Returns the number of keys in the key sequence. The maximum is 4.
*/
func (this *QKeySequence) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QKeySequence5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qkeysequence.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if the key sequence is empty; otherwise returns false.
*/
func (this *QKeySequence) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QKeySequence7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qkeysequence.h:172
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString(QKeySequence::SequenceFormat) const

/*
Return a string representation of the key sequence, based on format.

For example, the value Qt::CTRL+Qt::Key_O results in "Ctrl+O". If the key sequence has multiple key codes, each is separated by commas in the string returned, such as "Alt+X, Ctrl+Y, Z". The strings, "Ctrl", "Shift", etc. are translated using QObject::tr() in the "QShortcut" context.

If the key sequence has no keys, an empty string is returned.

On macOS, the string returned resembles the sequence that is shown in the menu bar if format is QKeySequence::NativeText; otherwise, the string uses the "portable" format, suitable for writing to a file.

This function was introduced in  Qt 4.1.

See also fromString().
*/
func (this *QKeySequence) ToString(format int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QKeySequence8toStringENS_14SequenceFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qkeysequence.h:172
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString(QKeySequence::SequenceFormat) const

/*
Return a string representation of the key sequence, based on format.

For example, the value Qt::CTRL+Qt::Key_O results in "Ctrl+O". If the key sequence has multiple key codes, each is separated by commas in the string returned, such as "Alt+X, Ctrl+Y, Z". The strings, "Ctrl", "Shift", etc. are translated using QObject::tr() in the "QShortcut" context.

If the key sequence has no keys, an empty string is returned.

On macOS, the string returned resembles the sequence that is shown in the menu bar if format is QKeySequence::NativeText; otherwise, the string uses the "portable" format, suitable for writing to a file.

This function was introduced in  Qt 4.1.

See also fromString().
*/
func (this *QKeySequence) ToString__() string {
	// arg: 0, QKeySequence::SequenceFormat=Enum, QKeySequence::SequenceFormat=Enum,
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QKeySequence8toStringENS_14SequenceFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qkeysequence.h:173
// index:0
// Public static Visibility=Default Availability=Available
// [8] QKeySequence fromString(const QString &, QKeySequence::SequenceFormat)

/*
Return a QKeySequence from the string str based on format.

This function was introduced in  Qt 4.1.

See also toString().
*/
func (this *QKeySequence) FromString(str string, format int) *QKeySequence /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QKeySequence10fromStringERK7QStringNS_14SequenceFormatE", qtrt.FFI_TYPE_POINTER, convArg0, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQKeySequence)
	return rv2
}
func QKeySequence_FromString(str string, format int) *QKeySequence /*123*/ {
	var nilthis *QKeySequence
	rv := nilthis.FromString(str, format)
	return rv
}

// /usr/include/qt/QtGui/qkeysequence.h:173
// index:0
// Public static Visibility=Default Availability=Available
// [8] QKeySequence fromString(const QString &, QKeySequence::SequenceFormat)

/*
Return a QKeySequence from the string str based on format.

This function was introduced in  Qt 4.1.

See also toString().
*/
func (this *QKeySequence) FromString__(str string) *QKeySequence /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QKeySequence::SequenceFormat=Enum, QKeySequence::SequenceFormat=Enum,
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QKeySequence10fromStringERK7QStringNS_14SequenceFormatE", qtrt.FFI_TYPE_POINTER, convArg0, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQKeySequence)
	return rv2
}

// /usr/include/qt/QtGui/qkeysequence.h:178
// index:0
// Public Visibility=Default Availability=Available
// [4] QKeySequence::SequenceMatch matches(const QKeySequence &) const

/*
Matches the sequence with seq. Returns ExactMatch if successful, PartialMatch if seq matches incompletely, and NoMatch if the sequences have nothing in common. Returns NoMatch if seq is shorter.
*/
func (this *QKeySequence) Matches(seq QKeySequence_ITF) int {
	var convArg0 unsafe.Pointer
	if seq != nil && seq.QKeySequence_PTR() != nil {
		convArg0 = seq.QKeySequence_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QKeySequence7matchesERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qkeysequence.h:179
// index:0
// Public static Visibility=Default Availability=Available
// [8] QKeySequence mnemonic(const QString &)

/*
Returns the shortcut key sequence for the mnemonic in text, or an empty key sequence if no mnemonics are found.

For example, mnemonic("E&xit") returns Qt::ALT+Qt::Key_X, mnemonic("&Quit") returns ALT+Key_Q, and mnemonic("Quit") returns an empty QKeySequence.

We provide a list of common mnemonics in English. At the time of writing, Microsoft and Open Group do not appear to have issued equivalent recommendations for other languages.
*/
func (this *QKeySequence) Mnemonic(text string) *QKeySequence /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QKeySequence8mnemonicERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQKeySequence)
	return rv2
}
func QKeySequence_Mnemonic(text string) *QKeySequence /*123*/ {
	var nilthis *QKeySequence
	rv := nilthis.Mnemonic(text)
	return rv
}

// /usr/include/qt/QtGui/qkeysequence.h:187
// index:0
// Public Visibility=Default Availability=Available
// [4] int operator[](uint) const

/*

 */
func (this *QKeySequence) Operator_get_index(i uint) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QKeySequenceixEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qkeysequence.h:188
// index:0
// Public Visibility=Default Availability=Available
// [8] QKeySequence & operator=(const QKeySequence &)

/*

 */
func (this *QKeySequence) Operator_equal(other QKeySequence_ITF) *QKeySequence {
	var convArg0 unsafe.Pointer
	if other != nil && other.QKeySequence_PTR() != nil {
		convArg0 = other.QKeySequence_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QKeySequenceaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQKeySequence)
	return rv2
}

// /usr/include/qt/QtGui/qkeysequence.h:190
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QKeySequence & operator=(QKeySequence &&)

/*

 */
func (this *QKeySequence) Operator_equal_1(other unsafe.Pointer /*333*/) *QKeySequence {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QKeySequenceaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQKeySequence)
	return rv2
}

// /usr/include/qt/QtGui/qkeysequence.h:192
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QKeySequence &)

/*
Swaps key sequence other with this key sequence. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
func (this *QKeySequence) Swap(other QKeySequence_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QKeySequence_PTR() != nil {
		convArg0 = other.QKeySequence_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QKeySequence4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qkeysequence.h:194
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QKeySequence &) const

/*

 */
func (this *QKeySequence) Operator_equal_equal(other QKeySequence_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QKeySequence_PTR() != nil {
		convArg0 = other.QKeySequence_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QKeySequenceeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qkeysequence.h:195
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QKeySequence &) const

/*

 */
func (this *QKeySequence) Operator_not_equal(other QKeySequence_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QKeySequence_PTR() != nil {
		convArg0 = other.QKeySequence_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QKeySequenceneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qkeysequence.h:197
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator<(const QKeySequence &) const

/*

 */
func (this *QKeySequence) Operator_less_than(ks QKeySequence_ITF) bool {
	var convArg0 unsafe.Pointer
	if ks != nil && ks.QKeySequence_PTR() != nil {
		convArg0 = ks.QKeySequence_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QKeySequenceltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qkeysequence.h:198
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>(const QKeySequence &) const

/*

 */
func (this *QKeySequence) Operator_greater_than(other QKeySequence_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QKeySequence_PTR() != nil {
		convArg0 = other.QKeySequence_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QKeySequencegtERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qkeysequence.h:200
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<=(const QKeySequence &) const

/*

 */
func (this *QKeySequence) Operator_less_than_equal(other QKeySequence_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QKeySequence_PTR() != nil {
		convArg0 = other.QKeySequence_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QKeySequenceleERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qkeysequence.h:202
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>=(const QKeySequence &) const

/*

 */
func (this *QKeySequence) Operator_greater_than_equal(other QKeySequence_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QKeySequence_PTR() != nil {
		convArg0 = other.QKeySequence_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QKeySequencegeERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qkeysequence.h:205
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDetached() const

/*

 */
func (this *QKeySequence) IsDetached() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QKeySequence10isDetachedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*
This enum represent standard key bindings. They can be used to assign platform dependent keyboard shortcuts to a QAction.

Note that the key bindings are platform dependent. The currently bound shortcuts can be queried using keyBindings().



This enum was introduced or modified in  Qt 4.2.

*/
type QKeySequence__StandardKey = int

// Unbound key.
const QKeySequence__UnknownKey QKeySequence__StandardKey = 0

// Open help contents.
const QKeySequence__HelpContents QKeySequence__StandardKey = 1

// Activate "what's this".
const QKeySequence__WhatsThis QKeySequence__StandardKey = 2

// Open document.
const QKeySequence__Open QKeySequence__StandardKey = 3

// Close document/tab.
const QKeySequence__Close QKeySequence__StandardKey = 4

// Save document.
const QKeySequence__Save QKeySequence__StandardKey = 5

// Create new document.
const QKeySequence__New QKeySequence__StandardKey = 6

// Delete.
const QKeySequence__Delete QKeySequence__StandardKey = 7

// Cut.
const QKeySequence__Cut QKeySequence__StandardKey = 8

// Copy.
const QKeySequence__Copy QKeySequence__StandardKey = 9

//
const QKeySequence__Paste QKeySequence__StandardKey = 10

//
const QKeySequence__Undo QKeySequence__StandardKey = 11

//
const QKeySequence__Redo QKeySequence__StandardKey = 12

//
const QKeySequence__Back QKeySequence__StandardKey = 13

//
const QKeySequence__Forward QKeySequence__StandardKey = 14

//
const QKeySequence__Refresh QKeySequence__StandardKey = 15

//
const QKeySequence__ZoomIn QKeySequence__StandardKey = 16

//
const QKeySequence__ZoomOut QKeySequence__StandardKey = 17

//
const QKeySequence__Print QKeySequence__StandardKey = 18

//
const QKeySequence__AddTab QKeySequence__StandardKey = 19

//
const QKeySequence__NextChild QKeySequence__StandardKey = 20

//
const QKeySequence__PreviousChild QKeySequence__StandardKey = 21

//
const QKeySequence__Find QKeySequence__StandardKey = 22

//
const QKeySequence__FindNext QKeySequence__StandardKey = 23

//
const QKeySequence__FindPrevious QKeySequence__StandardKey = 24

//
const QKeySequence__Replace QKeySequence__StandardKey = 25

//
const QKeySequence__SelectAll QKeySequence__StandardKey = 26

//
const QKeySequence__Bold QKeySequence__StandardKey = 27

//
const QKeySequence__Italic QKeySequence__StandardKey = 28

//
const QKeySequence__Underline QKeySequence__StandardKey = 29

//
const QKeySequence__MoveToNextChar QKeySequence__StandardKey = 30

//
const QKeySequence__MoveToPreviousChar QKeySequence__StandardKey = 31

//
const QKeySequence__MoveToNextWord QKeySequence__StandardKey = 32

//
const QKeySequence__MoveToPreviousWord QKeySequence__StandardKey = 33

//
const QKeySequence__MoveToNextLine QKeySequence__StandardKey = 34

//
const QKeySequence__MoveToPreviousLine QKeySequence__StandardKey = 35

//
const QKeySequence__MoveToNextPage QKeySequence__StandardKey = 36

//
const QKeySequence__MoveToPreviousPage QKeySequence__StandardKey = 37

//
const QKeySequence__MoveToStartOfLine QKeySequence__StandardKey = 38

//
const QKeySequence__MoveToEndOfLine QKeySequence__StandardKey = 39

//
const QKeySequence__MoveToStartOfBlock QKeySequence__StandardKey = 40

//
const QKeySequence__MoveToEndOfBlock QKeySequence__StandardKey = 41

//
const QKeySequence__MoveToStartOfDocument QKeySequence__StandardKey = 42

//
const QKeySequence__MoveToEndOfDocument QKeySequence__StandardKey = 43

//
const QKeySequence__SelectNextChar QKeySequence__StandardKey = 44

//
const QKeySequence__SelectPreviousChar QKeySequence__StandardKey = 45

//
const QKeySequence__SelectNextWord QKeySequence__StandardKey = 46

//
const QKeySequence__SelectPreviousWord QKeySequence__StandardKey = 47

//
const QKeySequence__SelectNextLine QKeySequence__StandardKey = 48

//
const QKeySequence__SelectPreviousLine QKeySequence__StandardKey = 49

//
const QKeySequence__SelectNextPage QKeySequence__StandardKey = 50

//
const QKeySequence__SelectPreviousPage QKeySequence__StandardKey = 51

//
const QKeySequence__SelectStartOfLine QKeySequence__StandardKey = 52

//
const QKeySequence__SelectEndOfLine QKeySequence__StandardKey = 53

//
const QKeySequence__SelectStartOfBlock QKeySequence__StandardKey = 54

//
const QKeySequence__SelectEndOfBlock QKeySequence__StandardKey = 55

//
const QKeySequence__SelectStartOfDocument QKeySequence__StandardKey = 56

//
const QKeySequence__SelectEndOfDocument QKeySequence__StandardKey = 57

//
const QKeySequence__DeleteStartOfWord QKeySequence__StandardKey = 58

//
const QKeySequence__DeleteEndOfWord QKeySequence__StandardKey = 59

//
const QKeySequence__DeleteEndOfLine QKeySequence__StandardKey = 60

//
const QKeySequence__InsertParagraphSeparator QKeySequence__StandardKey = 61

//
const QKeySequence__InsertLineSeparator QKeySequence__StandardKey = 62

//
const QKeySequence__SaveAs QKeySequence__StandardKey = 63

//
const QKeySequence__Preferences QKeySequence__StandardKey = 64

//
const QKeySequence__Quit QKeySequence__StandardKey = 65

//
const QKeySequence__FullScreen QKeySequence__StandardKey = 66

//
const QKeySequence__Deselect QKeySequence__StandardKey = 67

//
const QKeySequence__DeleteCompleteLine QKeySequence__StandardKey = 68

//
const QKeySequence__Backspace QKeySequence__StandardKey = 69

//
const QKeySequence__Cancel QKeySequence__StandardKey = 70

/*

 */
type QKeySequence__SequenceFormat = int

// The key sequence as a platform specific string. This means that it will be shown translated and on the Mac it will resemble a key sequence from the menu bar. This enum is best used when you want to display the string to the user.
const QKeySequence__NativeText QKeySequence__SequenceFormat = 0

//
const QKeySequence__PortableText QKeySequence__SequenceFormat = 1

/*

 */
type QKeySequence__SequenceMatch = int

// The key sequences are different; not even partially matching.
const QKeySequence__NoMatch QKeySequence__SequenceMatch = 0

// The key sequences match partially, but are not the same.
const QKeySequence__PartialMatch QKeySequence__SequenceMatch = 1

// The key sequences are the same.
const QKeySequence__ExactMatch QKeySequence__SequenceMatch = 2

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
