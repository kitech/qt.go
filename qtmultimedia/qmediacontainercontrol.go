package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediacontainercontrol.h
// #include <qmediacontainercontrol.h>
// #include <QtMultimedia>

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
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type QMediaContainerControl struct {
	*QMediaControl
}
type QMediaContainerControl_ITF interface {
	QMediaControl_ITF
	QMediaContainerControl_PTR() *QMediaContainerControl
}

func (ptr *QMediaContainerControl) QMediaContainerControl_PTR() *QMediaContainerControl { return ptr }

func (this *QMediaContainerControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QMediaContainerControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQMediaContainerControlFromPointer(cthis unsafe.Pointer) *QMediaContainerControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QMediaContainerControl{bcthis0}
}
func (*QMediaContainerControl) NewFromPointer(cthis unsafe.Pointer) *QMediaContainerControl {
	return NewQMediaContainerControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediacontainercontrol.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMediaContainerControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QMediaContainerControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediacontainercontrol.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMediaContainerControl()

/*

 */
func DeleteQMediaContainerControl(this *QMediaContainerControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QMediaContainerControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediacontainercontrol.h:58
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QStringList supportedContainers() const

/*
Returns a list of MIME types of supported container formats.
*/
func (this *QMediaContainerControl) SupportedContainers() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QMediaContainerControl19supportedContainersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediacontainercontrol.h:59
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString containerFormat() const

/*
Returns the selected container format.

See also setContainerFormat().
*/
func (this *QMediaContainerControl) ContainerFormat() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QMediaContainerControl15containerFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qmediacontainercontrol.h:60
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setContainerFormat(const QString &)

/*
Sets the current container format.

See also containerFormat().
*/
func (this *QMediaContainerControl) SetContainerFormat(format string) {
	var tmpArg0 = qtcore.NewQString5(format)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN22QMediaContainerControl18setContainerFormatERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediacontainercontrol.h:62
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString containerDescription(const QString &) const

/*
Returns a description of the container formatMimeType.
*/
func (this *QMediaContainerControl) ContainerDescription(formatMimeType string) string {
	var tmpArg0 = qtcore.NewQString5(formatMimeType)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QMediaContainerControl20containerDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qmediacontainercontrol.h:65
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMediaContainerControl(QObject *)

/*
Constructs a new media container control with the given parent.
*/
func (*QMediaContainerControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QMediaContainerControl {
	return NewQMediaContainerControl(parent)
}
func NewQMediaContainerControl(parent qtcore.QObject_ITF /*777 QObject **/) *QMediaContainerControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QMediaContainerControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaContainerControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaContainerControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediacontainercontrol.h:65
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMediaContainerControl(QObject *)

/*
Constructs a new media container control with the given parent.
*/
func (*QMediaContainerControl) NewForInheritp() *QMediaContainerControl {
	return NewQMediaContainerControlp()
}
func NewQMediaContainerControlp() *QMediaContainerControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN22QMediaContainerControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaContainerControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaContainerControl")
	return gothis
}

//  body block end

//  keep block begin

func init_unused_11851() {
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
