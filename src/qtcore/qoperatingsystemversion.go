//  header block begin
// /usr/include/qt/QtCore/qoperatingsystemversion.h
// #include <qoperatingsystemversion.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
type QOperatingSystemVersion struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:85
// index:0
// inline
// void QOperatingSystemVersion(enum QOperatingSystemVersion::OSType, int, int, int)
func NewQOperatingSystemVersion(osType int, vmajor int, vminor int, vmicro int) *QOperatingSystemVersion {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QOperatingSystemVersionC2ENS_6OSTypeEiii", ffiqt.FFI_TYPE_VOID, cthis, &osType, &vmajor, &vminor, &vmicro)
	gopp.ErrPrint(err, rv)
	return &QOperatingSystemVersion{cthis}
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:93
// index:0
// static
// QOperatingSystemVersion current()
func (this *QOperatingSystemVersion) Current() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QOperatingSystemVersion7currentEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QOperatingSystemVersion_Current() {
	// 0: (), ()
	var nilthis *QOperatingSystemVersion
	nilthis.Current()
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:95
// index:0
// static inline
// QOperatingSystemVersion::OSType currentType()
func (this *QOperatingSystemVersion) CurrentType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QOperatingSystemVersion11currentTypeEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QOperatingSystemVersion_CurrentType() {
	// 0: (), ()
	var nilthis *QOperatingSystemVersion
	nilthis.CurrentType()
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:114
// index:0
// inline
// int majorVersion()
func (this *QOperatingSystemVersion) MajorVersion() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QOperatingSystemVersion12majorVersionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:115
// index:0
// inline
// int minorVersion()
func (this *QOperatingSystemVersion) MinorVersion() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QOperatingSystemVersion12minorVersionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:116
// index:0
// inline
// int microVersion()
func (this *QOperatingSystemVersion) MicroVersion() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QOperatingSystemVersion12microVersionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:118
// index:0
// inline
// int segmentCount()
func (this *QOperatingSystemVersion) SegmentCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QOperatingSystemVersion12segmentCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:124
// index:0
// inline
// QOperatingSystemVersion::OSType type()
func (this *QOperatingSystemVersion) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QOperatingSystemVersion4typeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:125
// index:0
// QString name()
func (this *QOperatingSystemVersion) Name() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QOperatingSystemVersion4nameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
