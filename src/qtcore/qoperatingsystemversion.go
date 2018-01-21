package qtcore

// /usr/include/qt/QtCore/qoperatingsystemversion.h
// #include <qoperatingsystemversion.h>
// #include <QtCore>

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
	*qtrt.CObject
}

func (this *QOperatingSystemVersion) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQOperatingSystemVersionFromPointer(cthis unsafe.Pointer) *QOperatingSystemVersion {
	return &QOperatingSystemVersion{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:85
// index:0
// Public inline
// void QOperatingSystemVersion(enum QOperatingSystemVersion::OSType, int, int, int)
func NewQOperatingSystemVersion(osType int, vmajor int, vminor int, vmicro int) *QOperatingSystemVersion {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QOperatingSystemVersionC2ENS_6OSTypeEiii", ffiqt.FFI_TYPE_VOID, cthis, &osType, &vmajor, &vminor, &vmicro)
	gopp.ErrPrint(err, rv)
	gothis := NewQOperatingSystemVersionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:93
// index:0
// Public static
// QOperatingSystemVersion current()
func (this *QOperatingSystemVersion) Current() *QOperatingSystemVersion /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QOperatingSystemVersion7currentEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQOperatingSystemVersionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QOperatingSystemVersion_Current() *QOperatingSystemVersion /*123*/ {
	var nilthis *QOperatingSystemVersion
	rv := nilthis.Current()
	return rv
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:95
// index:0
// Public static inline
// QOperatingSystemVersion::OSType currentType()
func (this *QOperatingSystemVersion) CurrentType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QOperatingSystemVersion11currentTypeEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QOperatingSystemVersion_CurrentType() int {
	var nilthis *QOperatingSystemVersion
	rv := nilthis.CurrentType()
	return rv
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:114
// index:0
// Public inline
// int majorVersion()
func (this *QOperatingSystemVersion) MajorVersion() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QOperatingSystemVersion12majorVersionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:115
// index:0
// Public inline
// int minorVersion()
func (this *QOperatingSystemVersion) MinorVersion() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QOperatingSystemVersion12minorVersionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:116
// index:0
// Public inline
// int microVersion()
func (this *QOperatingSystemVersion) MicroVersion() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QOperatingSystemVersion12microVersionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:118
// index:0
// Public inline
// int segmentCount()
func (this *QOperatingSystemVersion) SegmentCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QOperatingSystemVersion12segmentCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:124
// index:0
// Public inline
// QOperatingSystemVersion::OSType type()
func (this *QOperatingSystemVersion) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QOperatingSystemVersion4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:125
// index:0
// Public
// QString name()
func (this *QOperatingSystemVersion) Name() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QOperatingSystemVersion4nameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
