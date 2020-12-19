package qtcore

// /usr/include/qt/QtCore/qsize.h
// #include <qsize.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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
// size 8
type QSize struct {
	*qtrt.CObject
}
type QSize_ITF interface {
	QSize_PTR() *QSize
}

func (ptr *QSize) QSize_PTR() *QSize { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
func QSizeFromptr(cthis unsafe.Pointer) *QSize {
	return &QSize{&qtrt.CObject{cthis}}
}
func (*QSize) Fromptr(cthis unsafe.Pointer) *QSize {
	return QSizeFromptr(cthis)
}

// /usr/include/qt/QtCore/qsize.h:57
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QSize(int, int)

/*
 */
func (*QSize) NewForInherit(w int, h int) *QSize {
	return NewQSize(w, h)
}
func NewQSize(w int, h int) *QSize {
	cthis := qtrt.Malloc(8)
	rv, err := qtrt.Qtcc1(1650702872, "_ZN5QSizeC2Eii", qtrt.FFITY_POINTER, cthis, w, h)
	qtrt.ErrPrint(err, rv)
	gothis := QSizeFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQSize)
	return gothis
}

// /usr/include/qt/QtCore/qsize.h:63
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int width() const

/*
 */
func (this *QSize) Width() int {
	rv, err := qtrt.Qtcc1(978481712, "_ZNK5QSize5widthEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qsize.h:64
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int height() const

/*
 */
func (this *QSize) Height() int {
	rv, err := qtrt.Qtcc1(1546467447, "_ZNK5QSize6heightEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qsize.h:83
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int & rwidth()

/*
 */
func (this *QSize) Rwidth() int {
	rv, err := qtrt.Qtcc1(1400378653, "_ZN5QSize6rwidthEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cpretval2go("int", rv).(int) // 3331
}

// /usr/include/qt/QtCore/qsize.h:84
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int & rheight()

/*
 */
func (this *QSize) Rheight() int {
	rv, err := qtrt.Qtcc1(3486783902, "_ZN5QSize7rheightEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cpretval2go("int", rv).(int) // 3331
}

func DeleteQSize(this *QSize) {
	rv, err := qtrt.Qtcc1(0, "_ZN5QSizeD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10039() {
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
