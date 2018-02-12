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
// extern C begin: 18
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QTextBoundaryFinder struct {
	*qtrt.CObject
}

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
// [-2] void QTextBoundaryFinder(enum QTextBoundaryFinder::BoundaryType, const QString &)
func NewQTextBoundaryFinder_1(type_ int, string string) *QTextBoundaryFinder {
	var tmpArg1 = NewQString_5(string)
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
// [-2] void QTextBoundaryFinder(enum QTextBoundaryFinder::BoundaryType, const QChar *, int, unsigned char *, int)
func NewQTextBoundaryFinder_2(type_ int, chars *QChar /*777 const QChar **/, length int, buffer unsafe.Pointer /*666*/, bufferSize int) *QTextBoundaryFinder {
	var convArg1 = chars.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextBoundaryFinderC2ENS_12BoundaryTypeEPK5QChariPhi", qtrt.FFI_TYPE_POINTER, type_, convArg1, length, buffer, bufferSize)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextBoundaryFinderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextBoundaryFinder)
	return gothis
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTextBoundaryFinder()
func DeleteQTextBoundaryFinder(this *QTextBoundaryFinder) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextBoundaryFinderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QTextBoundaryFinder) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QTextBoundaryFinder::BoundaryType type()
func (this *QTextBoundaryFinder) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QString string()
func (this *QTextBoundaryFinder) String() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder6stringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toStart()
func (this *QTextBoundaryFinder) ToStart() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextBoundaryFinder7toStartEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toEnd()
func (this *QTextBoundaryFinder) ToEnd() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextBoundaryFinder5toEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:86
// index:0
// Public Visibility=Default Availability=Available
// [4] int position()
func (this *QTextBoundaryFinder) Position() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPosition(int)
func (this *QTextBoundaryFinder) SetPosition(position int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextBoundaryFinder11setPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:89
// index:0
// Public Visibility=Default Availability=Available
// [4] int toNextBoundary()
func (this *QTextBoundaryFinder) ToNextBoundary() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextBoundaryFinder14toNextBoundaryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] int toPreviousBoundary()
func (this *QTextBoundaryFinder) ToPreviousBoundary() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextBoundaryFinder18toPreviousBoundaryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAtBoundary()
func (this *QTextBoundaryFinder) IsAtBoundary() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder12isAtBoundaryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] QTextBoundaryFinder::BoundaryReasons boundaryReasons()
func (this *QTextBoundaryFinder) BoundaryReasons() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder15boundaryReasonsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

type QTextBoundaryFinder__BoundaryType = int

const QTextBoundaryFinder__Grapheme QTextBoundaryFinder__BoundaryType = 0
const QTextBoundaryFinder__Word QTextBoundaryFinder__BoundaryType = 1
const QTextBoundaryFinder__Sentence QTextBoundaryFinder__BoundaryType = 2
const QTextBoundaryFinder__Line QTextBoundaryFinder__BoundaryType = 3

type QTextBoundaryFinder__BoundaryReason = int

const QTextBoundaryFinder__NotAtBoundary QTextBoundaryFinder__BoundaryReason = 0
const QTextBoundaryFinder__BreakOpportunity QTextBoundaryFinder__BoundaryReason = 31
const QTextBoundaryFinder__StartOfItem QTextBoundaryFinder__BoundaryReason = 32
const QTextBoundaryFinder__EndOfItem QTextBoundaryFinder__BoundaryReason = 64
const QTextBoundaryFinder__MandatoryBreak QTextBoundaryFinder__BoundaryReason = 128
const QTextBoundaryFinder__SoftHyphen QTextBoundaryFinder__BoundaryReason = 256

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
		qtrt.KeepMe()
	}
}

//  keep block end
