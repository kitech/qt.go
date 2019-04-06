package qtandroidextras

// /usr/include/qt/QtAndroidExtras/qandroidjnienvironment.h
// #include <qandroidjnienvironment.h>
// #include <QtAndroidExtras>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QAndroidJniExceptionCleaner struct {
	*qtrt.CObject
}
type QAndroidJniExceptionCleaner_ITF interface {
	QAndroidJniExceptionCleaner_PTR() *QAndroidJniExceptionCleaner
}

func (ptr *QAndroidJniExceptionCleaner) QAndroidJniExceptionCleaner_PTR() *QAndroidJniExceptionCleaner {
	return ptr
}

func (this *QAndroidJniExceptionCleaner) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAndroidJniExceptionCleaner) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAndroidJniExceptionCleanerFromPointer(cthis unsafe.Pointer) *QAndroidJniExceptionCleaner {
	return &QAndroidJniExceptionCleaner{&qtrt.CObject{cthis}}
}
func (*QAndroidJniExceptionCleaner) NewFromPointer(cthis unsafe.Pointer) *QAndroidJniExceptionCleaner {
	return NewQAndroidJniExceptionCleanerFromPointer(cthis)
}

// /usr/include/qt/QtAndroidExtras/qandroidjnienvironment.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAndroidJniExceptionCleaner(QAndroidJniExceptionCleaner::OutputMode)

/*

 */
func (*QAndroidJniExceptionCleaner) NewForInherit(outputMode int) *QAndroidJniExceptionCleaner {
	return NewQAndroidJniExceptionCleaner(outputMode)
}
func NewQAndroidJniExceptionCleaner(outputMode int) *QAndroidJniExceptionCleaner {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAndroidJniExceptionCleanerC2ENS_10OutputModeE", qtrt.FFI_TYPE_POINTER, outputMode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAndroidJniExceptionCleanerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidJniExceptionCleaner)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/qandroidjnienvironment.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAndroidJniExceptionCleaner(QAndroidJniExceptionCleaner::OutputMode)

/*

 */
func (*QAndroidJniExceptionCleaner) NewForInheritp() *QAndroidJniExceptionCleaner {
	return NewQAndroidJniExceptionCleanerp()
}
func NewQAndroidJniExceptionCleanerp() *QAndroidJniExceptionCleaner {
	// arg: 0, QAndroidJniExceptionCleaner::OutputMode=Enum, QAndroidJniExceptionCleaner::OutputMode=Enum, , Invalid
	outputMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAndroidJniExceptionCleanerC2ENS_10OutputModeE", qtrt.FFI_TYPE_POINTER, outputMode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAndroidJniExceptionCleanerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidJniExceptionCleaner)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/qandroidjnienvironment.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QAndroidJniExceptionCleaner()

/*

 */
func DeleteQAndroidJniExceptionCleaner(this *QAndroidJniExceptionCleaner) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAndroidJniExceptionCleanerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 4)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtAndroidExtras/qandroidjnienvironment.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clean()

/*

 */
func (this *QAndroidJniExceptionCleaner) Clean() {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAndroidJniExceptionCleaner5cleanEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QAndroidJniExceptionCleaner__OutputMode = int

//
const QAndroidJniExceptionCleaner__Silent QAndroidJniExceptionCleaner__OutputMode = 0

//
const QAndroidJniExceptionCleaner__Verbose QAndroidJniExceptionCleaner__OutputMode = 1

func (this *QAndroidJniExceptionCleaner) OutputModeItemName(val int) string {
	switch val {
	case QAndroidJniExceptionCleaner__Silent: // 0
		return "Silent"
	case QAndroidJniExceptionCleaner__Verbose: // 1
		return "Verbose"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QAndroidJniExceptionCleaner_OutputModeItemName(val int) string {
	var nilthis *QAndroidJniExceptionCleaner
	return nilthis.OutputModeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_14913() {
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
