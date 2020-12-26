package qtgui

// /usr/include/qt/QtGui/qbrush.h
// #include <qbrush.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
// size 8
type QBrush struct {
	*qtrt.CObject
}
type QBrush_ITF interface {
	QBrush_PTR() *QBrush
}

func (ptr *QBrush) QBrush_PTR() *QBrush { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QBrushFromptr(cthis Voidptr) *QBrush {
	return &QBrush{&qtrt.CObject{cthis}}
}
func (*QBrush) Fromptr(cthis Voidptr) *QBrush {
	return QBrushFromptr(cthis)
}

// /usr/include/qt/QtGui/qbrush.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QBrush()

/*
 */
func (*QBrush) NewForInherit() *QBrush {
	return NewQBrush()
}
func NewQBrush() *QBrush {
	cthis := qtrt.Malloc(8)
	rv, err := qtrt.Qtcc3(3258578779, "_ZN6QBrushC2Ev", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, Voidptr(&cthis))
	qtrt.ErrPrint2(err, rv)
	gothis := QBrushFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQBrush)
	return gothis
}

func DeleteQBrush(this *QBrush) {
	rv, err := qtrt.Qtcc3(1609486306, "_ZN6QBrushD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10143() {
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
