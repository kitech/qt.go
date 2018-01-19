//  header block begin
// /usr/include/qt/QtCore/qtextboundaryfinder.h
// #include <qtextboundaryfinder.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QTextBoundaryFinder struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:54
// index:0
// void QTextBoundaryFinder()
func NewQTextBoundaryFinder() *QTextBoundaryFinder {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextBoundaryFinderC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QTextBoundaryFinder{cthis}
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:76
// index:1
// void QTextBoundaryFinder(enum QTextBoundaryFinder::BoundaryType, const class QString &)
func NewQTextBoundaryFinder_1(type_ int, string unsafe.Pointer) *QTextBoundaryFinder {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextBoundaryFinderC2ENS_12BoundaryTypeERK7QString", ffiqt.FFI_TYPE_VOID, cthis, &type_, string)
	gopp.ErrPrint(err, rv)
	return &QTextBoundaryFinder{cthis}
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:77
// index:2
// void QTextBoundaryFinder(enum QTextBoundaryFinder::BoundaryType, const class QChar *, int, unsigned char *, int)
func NewQTextBoundaryFinder_2(type_ int, chars unsafe.Pointer, length int, buffer unsafe.Pointer, bufferSize int) *QTextBoundaryFinder {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextBoundaryFinderC2ENS_12BoundaryTypeEPK5QChariPhi", ffiqt.FFI_TYPE_VOID, cthis, &type_, chars, &length, buffer, &bufferSize)
	gopp.ErrPrint(err, rv)
	return &QTextBoundaryFinder{cthis}
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:57
// index:0
// void ~QTextBoundaryFinder()
func DeleteQTextBoundaryFinder(*QTextBoundaryFinder) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextBoundaryFinderD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:79
// index:0
// inline
// bool isValid()
func (this *QTextBoundaryFinder) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:81
// index:0
// inline
// QTextBoundaryFinder::BoundaryType type()
func (this *QTextBoundaryFinder) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder4typeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:82
// index:0
// QString string()
func (this *QTextBoundaryFinder) String() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder6stringEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:84
// index:0
// void toStart()
func (this *QTextBoundaryFinder) ToStart() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextBoundaryFinder7toStartEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:85
// index:0
// void toEnd()
func (this *QTextBoundaryFinder) ToEnd() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextBoundaryFinder5toEndEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:86
// index:0
// int position()
func (this *QTextBoundaryFinder) Position() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder8positionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:87
// index:0
// void setPosition(int)
func (this *QTextBoundaryFinder) SetPosition(position int) {
	// 0: (, int position), (&position)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextBoundaryFinder11setPositionEi", ffiqt.FFI_TYPE_VOID, this.cthis, &position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:89
// index:0
// int toNextBoundary()
func (this *QTextBoundaryFinder) ToNextBoundary() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextBoundaryFinder14toNextBoundaryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:90
// index:0
// int toPreviousBoundary()
func (this *QTextBoundaryFinder) ToPreviousBoundary() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextBoundaryFinder18toPreviousBoundaryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:92
// index:0
// bool isAtBoundary()
func (this *QTextBoundaryFinder) IsAtBoundary() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder12isAtBoundaryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextboundaryfinder.h:93
// index:0
// QTextBoundaryFinder::BoundaryReasons boundaryReasons()
func (this *QTextBoundaryFinder) BoundaryReasons() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextBoundaryFinder15boundaryReasonsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
