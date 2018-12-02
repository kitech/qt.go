package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediatimerange.h
// #include <qmediatimerange.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 54
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type QMediaTimeInterval struct {
	*qtrt.CObject
}
type QMediaTimeInterval_ITF interface {
	QMediaTimeInterval_PTR() *QMediaTimeInterval
}

func (ptr *QMediaTimeInterval) QMediaTimeInterval_PTR() *QMediaTimeInterval { return ptr }

func (this *QMediaTimeInterval) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMediaTimeInterval) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMediaTimeIntervalFromPointer(cthis unsafe.Pointer) *QMediaTimeInterval {
	return &QMediaTimeInterval{&qtrt.CObject{cthis}}
}
func (*QMediaTimeInterval) NewFromPointer(cthis unsafe.Pointer) *QMediaTimeInterval {
	return NewQMediaTimeIntervalFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMediaTimeInterval()

/*

 */
func (*QMediaTimeInterval) NewForInherit() *QMediaTimeInterval {
	return NewQMediaTimeInterval()
}
func NewQMediaTimeInterval() *QMediaTimeInterval {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QMediaTimeIntervalC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaTimeIntervalFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMediaTimeInterval)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:56
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QMediaTimeInterval(qint64, qint64)

/*

 */
func (*QMediaTimeInterval) NewForInherit_1(start int64, end_ int64) *QMediaTimeInterval {
	return NewQMediaTimeInterval_1(start, end_)
}
func NewQMediaTimeInterval_1(start int64, end_ int64) *QMediaTimeInterval {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QMediaTimeIntervalC2Exx", qtrt.FFI_TYPE_POINTER, start, end_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaTimeIntervalFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMediaTimeInterval)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:59
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 start() const

/*

 */
func (this *QMediaTimeInterval) Start() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QMediaTimeInterval5startEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:60
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 end() const

/*

 */
func (this *QMediaTimeInterval) End() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QMediaTimeInterval3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:62
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(qint64) const

/*
Returns true if the specified time lies within the time range.
*/
func (this *QMediaTimeInterval) Contains(time int64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QMediaTimeInterval8containsEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), time)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:64
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNormal() const

/*

 */
func (this *QMediaTimeInterval) IsNormal() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QMediaTimeInterval8isNormalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:65
// index:0
// Public Visibility=Default Availability=Available
// [16] QMediaTimeInterval normalized() const

/*

 */
func (this *QMediaTimeInterval) Normalized() *QMediaTimeInterval /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QMediaTimeInterval10normalizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMediaTimeIntervalFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMediaTimeInterval)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:66
// index:0
// Public Visibility=Default Availability=Available
// [16] QMediaTimeInterval translated(qint64) const

/*

 */
func (this *QMediaTimeInterval) Translated(offset int64) *QMediaTimeInterval /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QMediaTimeInterval10translatedEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), offset)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMediaTimeIntervalFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMediaTimeInterval)
	return rv2
}

func DeleteQMediaTimeInterval(this *QMediaTimeInterval) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QMediaTimeIntervalD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
