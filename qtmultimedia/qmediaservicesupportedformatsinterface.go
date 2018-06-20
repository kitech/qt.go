package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h
// #include <qmediaserviceproviderplugin.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
type QMediaServiceSupportedFormatsInterface struct {
	*qtrt.CObject
}
type QMediaServiceSupportedFormatsInterface_ITF interface {
	QMediaServiceSupportedFormatsInterface_PTR() *QMediaServiceSupportedFormatsInterface
}

func (ptr *QMediaServiceSupportedFormatsInterface) QMediaServiceSupportedFormatsInterface_PTR() *QMediaServiceSupportedFormatsInterface {
	return ptr
}

func (this *QMediaServiceSupportedFormatsInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMediaServiceSupportedFormatsInterface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMediaServiceSupportedFormatsInterfaceFromPointer(cthis unsafe.Pointer) *QMediaServiceSupportedFormatsInterface {
	return &QMediaServiceSupportedFormatsInterface{&qtrt.CObject{cthis}}
}
func (*QMediaServiceSupportedFormatsInterface) NewFromPointer(cthis unsafe.Pointer) *QMediaServiceSupportedFormatsInterface {
	return NewQMediaServiceSupportedFormatsInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:126
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [-2] void ~QMediaServiceSupportedFormatsInterface()

/*

 */
func DeleteQMediaServiceSupportedFormatsInterface(this *QMediaServiceSupportedFormatsInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN38QMediaServiceSupportedFormatsInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:127
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QMultimedia::SupportEstimate hasSupport(const QString &, const QStringList &) const

/*

 */
func (this *QMediaServiceSupportedFormatsInterface) HasSupport(mimeType string, codecs qtcore.QStringList_ITF) int {
	var tmpArg0 = qtcore.NewQString_5(mimeType)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if codecs != nil && codecs.QStringList_PTR() != nil {
		convArg1 = codecs.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK38QMediaServiceSupportedFormatsInterface10hasSupportERK7QStringRK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:128
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QStringList supportedMimeTypes() const

/*

 */
func (this *QMediaServiceSupportedFormatsInterface) SupportedMimeTypes() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK38QMediaServiceSupportedFormatsInterface18supportedMimeTypesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
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
