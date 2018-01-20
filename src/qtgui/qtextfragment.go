//  header block begin
// /usr/include/qt/QtGui/qtextobject.h
// #include <qtextobject.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 35
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
type QTextFragment struct {
	*qtrt.CObject
}

func (this *QTextFragment) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQTextFragmentFromPointer(cthis unsafe.Pointer) *QTextFragment {
	return &QTextFragment{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qtextobject.h:306
// index:0
// Public inline
// void QTextFragment(const class QTextDocumentPrivate *, int, int)
func NewQTextFragment(priv unsafe.Pointer, f int, fe int) *QTextFragment {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextFragmentC2EPK20QTextDocumentPrivateii", ffiqt.FFI_TYPE_VOID, cthis, priv, &f, &fe)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextFragmentFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextobject.h:307
// index:1
// Public inline
// void QTextFragment()
func NewQTextFragment_1() *QTextFragment {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextFragmentC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextFragmentFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextobject.h:311
// index:0
// Public inline
// bool isValid()
func (this *QTextFragment) IsValid() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextFragment7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextobject.h:317
// index:0
// Public
// int position()
func (this *QTextFragment) Position() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextFragment8positionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextobject.h:318
// index:0
// Public
// int length()
func (this *QTextFragment) Length() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextFragment6lengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextobject.h:319
// index:0
// Public
// bool contains(int)
func (this *QTextFragment) Contains(position int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextFragment8containsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &position)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextobject.h:321
// index:0
// Public
// QTextCharFormat charFormat()
func (this *QTextFragment) CharFormat() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextFragment10charFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextobject.h:322
// index:0
// Public
// int charFormatIndex()
func (this *QTextFragment) CharFormatIndex() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextFragment15charFormatIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextobject.h:323
// index:0
// Public
// QString text()
func (this *QTextFragment) Text() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextFragment4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextobject.h:326
// index:0
// Public
// QList<QGlyphRun> glyphRuns(int, int)
func (this *QTextFragment) GlyphRuns(from int, length int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextFragment9glyphRunsEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &from, &length)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
