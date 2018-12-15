package qtcore

// /usr/include/qt/QtCore/qtextboundaryfinder.h
// #include <qtextboundaryfinder.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QTextBoundaryFinder struct {
	*qtrt.CObject
}
type QTextBoundaryFinder_ITF interface {
	QTextBoundaryFinder_PTR() *QTextBoundaryFinder
}

func (ptr *QTextBoundaryFinder) QTextBoundaryFinder_PTR() *QTextBoundaryFinder { return ptr }

func (this *QTextBoundaryFinder) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextBoundaryFinder) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextBoundaryFinderFromPointer(cthis unsafe.Pointer) *QTextBoundaryFinder {
	return &QTextBoundaryFinder{&qtrt.CObject{cthis}}
}
func (*QTextBoundaryFinder) NewFromPointer(cthis unsafe.Pointer) *QTextBoundaryFinder {
	return NewQTextBoundaryFinderFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:54
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextBoundaryFinder()

/*
Constructs an invalid QTextBoundaryFinder object.
*/
func (*QTextBoundaryFinder) NewForInherit() *QTextBoundaryFinder {
	return NewQTextBoundaryFinder()
}
func NewQTextBoundaryFinder() *QTextBoundaryFinder {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextBoundaryFinderC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextBoundaryFinderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextBoundaryFinder)
	return gothis
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:76
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextBoundaryFinder(QTextBoundaryFinder::BoundaryType, const QString &)

/*
Constructs an invalid QTextBoundaryFinder object.
*/
func (*QTextBoundaryFinder) NewForInherit1(type_ int, string string) *QTextBoundaryFinder {
	return NewQTextBoundaryFinder1(type_, string)
}
func NewQTextBoundaryFinder1(type_ int, string string) *QTextBoundaryFinder {
	var tmpArg1 = NewQString5(string)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextBoundaryFinderC2ENS_12BoundaryTypeERK7QString", qtrt.FFI_TYPE_POINTER, type_, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextBoundaryFinderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextBoundaryFinder)
	return gothis
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:77
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTextBoundaryFinder(QTextBoundaryFinder::BoundaryType, const QChar *, int, unsigned char *, int)

/*
Constructs an invalid QTextBoundaryFinder object.
*/
func (*QTextBoundaryFinder) NewForInherit2(type_ int, chars QChar_ITF /*777 const QChar **/, length int, buffer unsafe.Pointer /*666*/, bufferSize int) *QTextBoundaryFinder {
	return NewQTextBoundaryFinder2(type_, chars, length, buffer, bufferSize)
}
func NewQTextBoundaryFinder2(type_ int, chars QChar_ITF /*777 const QChar **/, length int, buffer unsafe.Pointer /*666*/, bufferSize int) *QTextBoundaryFinder {
	var convArg1 unsafe.Pointer
	if chars != nil && chars.QChar_PTR() != nil {
		convArg1 = chars.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextBoundaryFinderC2ENS_12BoundaryTypeEPK5QChariPhi", qtrt.FFI_TYPE_POINTER, type_, convArg1, length, buffer, bufferSize)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextBoundaryFinderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextBoundaryFinder)
	return gothis
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:77
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTextBoundaryFinder(QTextBoundaryFinder::BoundaryType, const QChar *, int, unsigned char *, int)

/*
Constructs an invalid QTextBoundaryFinder object.
*/
func (*QTextBoundaryFinder) NewForInherit2p(type_ int, chars QChar_ITF /*777 const QChar **/, length int) *QTextBoundaryFinder {
	return NewQTextBoundaryFinder2p(type_, chars, length)
}
func NewQTextBoundaryFinder2p(type_ int, chars QChar_ITF /*777 const QChar **/, length int) *QTextBoundaryFinder {
	var convArg1 unsafe.Pointer
	if chars != nil && chars.QChar_PTR() != nil {
		convArg1 = chars.QChar_PTR().GetCthis()
	}
	// arg: 3, unsigned char *=Pointer, =Invalid, , Invalid
	var buffer unsafe.Pointer
	// arg: 4, int=Int, =Invalid, , Invalid
	bufferSize := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextBoundaryFinderC2ENS_12BoundaryTypeEPK5QChariPhi", qtrt.FFI_TYPE_POINTER, type_, convArg1, length, buffer, bufferSize)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextBoundaryFinderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextBoundaryFinder)
	return gothis
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:77
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTextBoundaryFinder(QTextBoundaryFinder::BoundaryType, const QChar *, int, unsigned char *, int)

/*
Constructs an invalid QTextBoundaryFinder object.
*/
func (*QTextBoundaryFinder) NewForInherit2p1(type_ int, chars QChar_ITF /*777 const QChar **/, length int, buffer unsafe.Pointer /*666*/) *QTextBoundaryFinder {
	return NewQTextBoundaryFinder2p1(type_, chars, length, buffer)
}
func NewQTextBoundaryFinder2p1(type_ int, chars QChar_ITF /*777 const QChar **/, length int, buffer unsafe.Pointer /*666*/) *QTextBoundaryFinder {
	var convArg1 unsafe.Pointer
	if chars != nil && chars.QChar_PTR() != nil {
		convArg1 = chars.QChar_PTR().GetCthis()
	}
	// arg: 4, int=Int, =Invalid, , Invalid
	bufferSize := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextBoundaryFinderC2ENS_12BoundaryTypeEPK5QChariPhi", qtrt.FFI_TYPE_POINTER, type_, convArg1, length, buffer, bufferSize)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextBoundaryFinderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextBoundaryFinder)
	return gothis
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:56
// index:0
// Public Visibility=Default Availability=Available
// [48] QTextBoundaryFinder & operator=(const QTextBoundaryFinder &)

/*

 */
func (this *QTextBoundaryFinder) Operator_equal(other QTextBoundaryFinder_ITF) *QTextBoundaryFinder {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTextBoundaryFinder_PTR() != nil {
		convArg0 = other.QTextBoundaryFinder_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextBoundaryFinderaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBoundaryFinderFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBoundaryFinder)
	return rv2
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTextBoundaryFinder()

/*

 */
func DeleteQTextBoundaryFinder(this *QTextBoundaryFinder) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextBoundaryFinderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the text boundary finder is valid; otherwise returns false. A default QTextBoundaryFinder is invalid.
*/
func (this *QTextBoundaryFinder) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QTextBoundaryFinder::BoundaryType type() const

/*
Returns the type of the QTextBoundaryFinder.
*/
func (this *QTextBoundaryFinder) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QString string() const

/*
Returns the string the QTextBoundaryFinder object operates on.
*/
func (this *QTextBoundaryFinder) String() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder6stringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toStart()

/*
Moves the finder to the start of the string. This is equivalent to setPosition(0).

See also setPosition() and position().
*/
func (this *QTextBoundaryFinder) ToStart() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextBoundaryFinder7toStartEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toEnd()

/*
Moves the finder to the end of the string. This is equivalent to setPosition(string.length()).

See also setPosition() and position().
*/
func (this *QTextBoundaryFinder) ToEnd() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextBoundaryFinder5toEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:86
// index:0
// Public Visibility=Default Availability=Available
// [4] int position() const

/*
Returns the current position of the QTextBoundaryFinder.

The range is from 0 (the beginning of the string) to the length of the string inclusive.

See also setPosition().
*/
func (this *QTextBoundaryFinder) Position() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPosition(int)

/*
Sets the current position of the QTextBoundaryFinder to position.

If position is out of bounds, it will be bound to only valid positions. In this case, valid positions are from 0 to the length of the string inclusive.

See also position().
*/
func (this *QTextBoundaryFinder) SetPosition(position int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextBoundaryFinder11setPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:89
// index:0
// Public Visibility=Default Availability=Available
// [4] int toNextBoundary()

/*
Moves the QTextBoundaryFinder to the next boundary position and returns that position.

Returns -1 if there is no next boundary.
*/
func (this *QTextBoundaryFinder) ToNextBoundary() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextBoundaryFinder14toNextBoundaryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] int toPreviousBoundary()

/*
Moves the QTextBoundaryFinder to the previous boundary position and returns that position.

Returns -1 if there is no previous boundary.
*/
func (this *QTextBoundaryFinder) ToPreviousBoundary() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextBoundaryFinder18toPreviousBoundaryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAtBoundary() const

/*
Returns true if the object's position() is currently at a valid text boundary.
*/
func (this *QTextBoundaryFinder) IsAtBoundary() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder12isAtBoundaryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] QTextBoundaryFinder::BoundaryReasons boundaryReasons() const

/*
Returns the reasons for the boundary finder to have chosen the current position as a boundary.
*/
func (this *QTextBoundaryFinder) BoundaryReasons() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder15boundaryReasonsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

/*

 */
type QTextBoundaryFinder__BoundaryType = int

// Finds a grapheme which is the smallest boundary. It including letters, punctuation marks, numerals and more.
const QTextBoundaryFinder__Grapheme QTextBoundaryFinder__BoundaryType = 0

// Finds a word.
const QTextBoundaryFinder__Word QTextBoundaryFinder__BoundaryType = 1

// Finds sentence boundaries. These include periods, question marks etc.
const QTextBoundaryFinder__Sentence QTextBoundaryFinder__BoundaryType = 2

// Finds possible positions for breaking the text into multiple lines.
const QTextBoundaryFinder__Line QTextBoundaryFinder__BoundaryType = 3

func (this *QTextBoundaryFinder) BoundaryTypeItemName(val int) string {
	switch val {
	case QTextBoundaryFinder__Grapheme: // 0
		return "Grapheme"
	case QTextBoundaryFinder__Word: // 1
		return "Word"
	case QTextBoundaryFinder__Sentence: // 2
		return "Sentence"
	case QTextBoundaryFinder__Line: // 3
		return "Line"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextBoundaryFinder_BoundaryTypeItemName(val int) string {
	var nilthis *QTextBoundaryFinder
	return nilthis.BoundaryTypeItemName(val)
}

/*


 */
type QTextBoundaryFinder__BoundaryReason = int

//
const QTextBoundaryFinder__NotAtBoundary QTextBoundaryFinder__BoundaryReason = 0

//
const QTextBoundaryFinder__BreakOpportunity QTextBoundaryFinder__BoundaryReason = 31

//
const QTextBoundaryFinder__StartOfItem QTextBoundaryFinder__BoundaryReason = 32

//
const QTextBoundaryFinder__EndOfItem QTextBoundaryFinder__BoundaryReason = 64

//
const QTextBoundaryFinder__MandatoryBreak QTextBoundaryFinder__BoundaryReason = 128

//
const QTextBoundaryFinder__SoftHyphen QTextBoundaryFinder__BoundaryReason = 256

func (this *QTextBoundaryFinder) BoundaryReasonItemName(val int) string {
	switch val {
	case QTextBoundaryFinder__NotAtBoundary: // 0
		return "NotAtBoundary"
	case QTextBoundaryFinder__BreakOpportunity: // 31
		return "BreakOpportunity"
	case QTextBoundaryFinder__StartOfItem: // 32
		return "StartOfItem"
	case QTextBoundaryFinder__EndOfItem: // 64
		return "EndOfItem"
	case QTextBoundaryFinder__MandatoryBreak: // 128
		return "MandatoryBreak"
	case QTextBoundaryFinder__SoftHyphen: // 256
		return "SoftHyphen"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextBoundaryFinder_BoundaryReasonItemName(val int) string {
	var nilthis *QTextBoundaryFinder
	return nilthis.BoundaryReasonItemName(val)
}

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
}

//  keep block end
