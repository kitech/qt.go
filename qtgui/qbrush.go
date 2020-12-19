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
func QBrushFromptr(cthis unsafe.Pointer) *QBrush {
	return &QBrush{&qtrt.CObject{cthis}}
}
func (*QBrush) Fromptr(cthis unsafe.Pointer) *QBrush {
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
	rv, err := qtrt.Qtcc1(3258578779, "_ZN6QBrushC2Ev", qtrt.FFITY_POINTER, cthis)
	qtrt.ErrPrint(err, rv)
	gothis := QBrushFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQBrush)
	return gothis
}

func DeleteQBrush(this *QBrush) {
	rv, err := qtrt.Qtcc1(0, "_ZN6QBrushD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10057() {
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
