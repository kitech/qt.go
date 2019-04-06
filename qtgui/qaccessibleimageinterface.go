package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

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
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QAccessibleImageInterface struct {
	*qtrt.CObject
}
type QAccessibleImageInterface_ITF interface {
	QAccessibleImageInterface_PTR() *QAccessibleImageInterface
}

func (ptr *QAccessibleImageInterface) QAccessibleImageInterface_PTR() *QAccessibleImageInterface {
	return ptr
}

func (this *QAccessibleImageInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAccessibleImageInterface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAccessibleImageInterfaceFromPointer(cthis unsafe.Pointer) *QAccessibleImageInterface {
	return &QAccessibleImageInterface{&qtrt.CObject{cthis}}
}
func (*QAccessibleImageInterface) NewFromPointer(cthis unsafe.Pointer) *QAccessibleImageInterface {
	return NewQAccessibleImageInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessible.h:655
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleImageInterface()

/*

 */
func DeleteQAccessibleImageInterface(this *QAccessibleImageInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QAccessibleImageInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessible.h:657
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString imageDescription() const

/*

 */
func (this *QAccessibleImageInterface) ImageDescription() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QAccessibleImageInterface16imageDescriptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qaccessible.h:658
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSize imageSize() const

/*

 */
func (this *QAccessibleImageInterface) ImageSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QAccessibleImageInterface9imageSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qaccessible.h:659
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QPoint imagePosition() const

/*

 */
func (this *QAccessibleImageInterface) ImagePosition() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QAccessibleImageInterface13imagePositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

//  body block end

//  keep block begin

func init_unused_10815() {
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
