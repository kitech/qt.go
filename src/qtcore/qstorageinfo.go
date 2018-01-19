//  header block begin
// /usr/include/qt/QtCore/qstorageinfo.h
// #include <qstorageinfo.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 26
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
type QStorageInfo struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qstorageinfo.h:58
// index:0
// void QStorageInfo()
func NewQStorageInfo() *QStorageInfo {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStorageInfoC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QStorageInfo{cthis}
}

// /usr/include/qt/QtCore/qstorageinfo.h:59
// index:1
// void QStorageInfo(const class QString &)
func NewQStorageInfo_1(path unsafe.Pointer) *QStorageInfo {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStorageInfoC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, path)
	gopp.ErrPrint(err, rv)
	return &QStorageInfo{cthis}
}

// /usr/include/qt/QtCore/qstorageinfo.h:60
// index:2
// void QStorageInfo(const class QDir &)
func NewQStorageInfo_2(dir unsafe.Pointer) *QStorageInfo {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStorageInfoC2ERK4QDir", ffiqt.FFI_TYPE_VOID, cthis, dir)
	gopp.ErrPrint(err, rv)
	return &QStorageInfo{cthis}
}

// /usr/include/qt/QtCore/qstorageinfo.h:62
// index:0
// void ~QStorageInfo()
func DeleteQStorageInfo(*QStorageInfo) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStorageInfoD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:69
// index:0
// inline
// void swap(class QStorageInfo &)
func (this *QStorageInfo) Swap(other unsafe.Pointer) {
	// 0: (, QStorageInfo & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStorageInfo4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:72
// index:0
// void setPath(const class QString &)
func (this *QStorageInfo) SetPath(path unsafe.Pointer) {
	// 0: (, const QString & path), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStorageInfo7setPathERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, path)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:74
// index:0
// QString rootPath()
func (this *QStorageInfo) RootPath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo8rootPathEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:75
// index:0
// QByteArray device()
func (this *QStorageInfo) Device() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo6deviceEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:76
// index:0
// QByteArray subvolume()
func (this *QStorageInfo) Subvolume() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo9subvolumeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:77
// index:0
// QByteArray fileSystemType()
func (this *QStorageInfo) FileSystemType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo14fileSystemTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:78
// index:0
// QString name()
func (this *QStorageInfo) Name() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo4nameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:79
// index:0
// QString displayName()
func (this *QStorageInfo) DisplayName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo11displayNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:81
// index:0
// qint64 bytesTotal()
func (this *QStorageInfo) BytesTotal() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo10bytesTotalEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:82
// index:0
// qint64 bytesFree()
func (this *QStorageInfo) BytesFree() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo9bytesFreeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:83
// index:0
// qint64 bytesAvailable()
func (this *QStorageInfo) BytesAvailable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo14bytesAvailableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:84
// index:0
// int blockSize()
func (this *QStorageInfo) BlockSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo9blockSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:86
// index:0
// inline
// bool isRoot()
func (this *QStorageInfo) IsRoot() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo6isRootEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:87
// index:0
// bool isReadOnly()
func (this *QStorageInfo) IsReadOnly() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo10isReadOnlyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:88
// index:0
// bool isReady()
func (this *QStorageInfo) IsReady() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo7isReadyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:89
// index:0
// bool isValid()
func (this *QStorageInfo) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:91
// index:0
// void refresh()
func (this *QStorageInfo) Refresh() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStorageInfo7refreshEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:93
// index:0
// static
// QList<QStorageInfo> mountedVolumes()
func (this *QStorageInfo) MountedVolumes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStorageInfo14mountedVolumesEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStorageInfo_MountedVolumes() {
	// 0: (), ()
	var nilthis *QStorageInfo
	nilthis.MountedVolumes()
}

// /usr/include/qt/QtCore/qstorageinfo.h:94
// index:0
// static
// QStorageInfo root()
func (this *QStorageInfo) Root() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStorageInfo4rootEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStorageInfo_Root() {
	// 0: (), ()
	var nilthis *QStorageInfo
	nilthis.Root()
}

//  body block end
