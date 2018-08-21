package qtwebenginewidgets

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h
// #include <qwebenginedownloaditem.h>
// #include <QtWebEngineWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtqml"
import "github.com/kitech/qt.go/qtpositioning"
import "github.com/kitech/qt.go/qtwebchannel"
import "github.com/kitech/qt.go/qtquickwidgets"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtwidgets"
import "github.com/kitech/qt.go/qtprintsupport"
import "github.com/kitech/qt.go/qtquick"
import "github.com/kitech/qt.go/qtwebenginecore"

//  ext block end

//  body block begin

/*

 */
type QWebEngineDownloadItem struct {
	*qtcore.QObject
}
type QWebEngineDownloadItem_ITF interface {
	qtcore.QObject_ITF
	QWebEngineDownloadItem_PTR() *QWebEngineDownloadItem
}

func (ptr *QWebEngineDownloadItem) QWebEngineDownloadItem_PTR() *QWebEngineDownloadItem { return ptr }

func (this *QWebEngineDownloadItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QWebEngineDownloadItem) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQWebEngineDownloadItemFromPointer(cthis unsafe.Pointer) *QWebEngineDownloadItem {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QWebEngineDownloadItem{bcthis0}
}
func (*QWebEngineDownloadItem) NewFromPointer(cthis unsafe.Pointer) *QWebEngineDownloadItem {
	return NewQWebEngineDownloadItemFromPointer(cthis)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QWebEngineDownloadItem) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QWebEngineDownloadItem10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWebEngineDownloadItem()

/*

 */
func DeleteQWebEngineDownloadItem(this *QWebEngineDownloadItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QWebEngineDownloadItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:114
// index:0
// Public Visibility=Default Availability=Available
// [4] quint32 id() const

/*

 */
func (this *QWebEngineDownloadItem) Id() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QWebEngineDownloadItem2idEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:115
// index:0
// Public Visibility=Default Availability=Available
// [4] QWebEngineDownloadItem::DownloadState state() const

/*

 */
func (this *QWebEngineDownloadItem) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QWebEngineDownloadItem5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 totalBytes() const

/*

 */
func (this *QWebEngineDownloadItem) TotalBytes() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QWebEngineDownloadItem10totalBytesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:117
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 receivedBytes() const

/*

 */
func (this *QWebEngineDownloadItem) ReceivedBytes() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QWebEngineDownloadItem13receivedBytesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl url() const

/*

 */
func (this *QWebEngineDownloadItem) Url() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QWebEngineDownloadItem3urlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QString mimeType() const

/*

 */
func (this *QWebEngineDownloadItem) MimeType() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QWebEngineDownloadItem8mimeTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:120
// index:0
// Public Visibility=Default Availability=Available
// [8] QString path() const

/*

 */
func (this *QWebEngineDownloadItem) Path() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QWebEngineDownloadItem4pathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPath(QString)

/*

 */
func (this *QWebEngineDownloadItem) SetPath(path string) {
	var tmpArg0 = qtcore.NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN22QWebEngineDownloadItem7setPathE7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:122
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFinished() const

/*

 */
func (this *QWebEngineDownloadItem) IsFinished() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QWebEngineDownloadItem10isFinishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:123
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPaused() const

/*

 */
func (this *QWebEngineDownloadItem) IsPaused() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QWebEngineDownloadItem8isPausedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:124
// index:0
// Public Visibility=Default Availability=Available
// [4] QWebEngineDownloadItem::SavePageFormat savePageFormat() const

/*

 */
func (this *QWebEngineDownloadItem) SavePageFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QWebEngineDownloadItem14savePageFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSavePageFormat(QWebEngineDownloadItem::SavePageFormat)

/*

 */
func (this *QWebEngineDownloadItem) SetSavePageFormat(format int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QWebEngineDownloadItem17setSavePageFormatENS_14SavePageFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:126
// index:0
// Public Visibility=Default Availability=Available
// [4] QWebEngineDownloadItem::DownloadType type() const

/*

 */
func (this *QWebEngineDownloadItem) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QWebEngineDownloadItem4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:127
// index:0
// Public Visibility=Default Availability=Available
// [4] QWebEngineDownloadItem::DownloadInterruptReason interruptReason() const

/*

 */
func (this *QWebEngineDownloadItem) InterruptReason() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QWebEngineDownloadItem15interruptReasonEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] QString interruptReasonString() const

/*

 */
func (this *QWebEngineDownloadItem) InterruptReasonString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QWebEngineDownloadItem21interruptReasonStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void accept()

/*

 */
func (this *QWebEngineDownloadItem) Accept() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QWebEngineDownloadItem6acceptEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cancel()

/*

 */
func (this *QWebEngineDownloadItem) Cancel() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QWebEngineDownloadItem6cancelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void pause()

/*

 */
func (this *QWebEngineDownloadItem) Pause() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QWebEngineDownloadItem5pauseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resume()

/*

 */
func (this *QWebEngineDownloadItem) Resume() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QWebEngineDownloadItem6resumeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void finished()

/*

 */
func (this *QWebEngineDownloadItem) Finished() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QWebEngineDownloadItem8finishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QWebEngineDownloadItem::DownloadState)

/*

 */
func (this *QWebEngineDownloadItem) StateChanged(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QWebEngineDownloadItem12stateChangedENS_13DownloadStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void downloadProgress(qint64, qint64)

/*

 */
func (this *QWebEngineDownloadItem) DownloadProgress(bytesReceived int64, bytesTotal int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QWebEngineDownloadItem16downloadProgressExx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bytesReceived, bytesTotal)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginedownloaditem.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void isPausedChanged(bool)

/*

 */
func (this *QWebEngineDownloadItem) IsPausedChanged(isPaused bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QWebEngineDownloadItem15isPausedChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), isPaused)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QWebEngineDownloadItem__DownloadState = int

//
const QWebEngineDownloadItem__DownloadRequested QWebEngineDownloadItem__DownloadState = 0

//
const QWebEngineDownloadItem__DownloadInProgress QWebEngineDownloadItem__DownloadState = 1

//
const QWebEngineDownloadItem__DownloadCompleted QWebEngineDownloadItem__DownloadState = 2

//
const QWebEngineDownloadItem__DownloadCancelled QWebEngineDownloadItem__DownloadState = 3

//
const QWebEngineDownloadItem__DownloadInterrupted QWebEngineDownloadItem__DownloadState = 4

func (this *QWebEngineDownloadItem) DownloadStateItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QWebEngineDownloadItem_DownloadStateItemName(val int) string {
	var nilthis *QWebEngineDownloadItem
	return nilthis.DownloadStateItemName(val)
}

/*


 */
type QWebEngineDownloadItem__SavePageFormat = int

//
const QWebEngineDownloadItem__UnknownSaveFormat QWebEngineDownloadItem__SavePageFormat = -1

//
const QWebEngineDownloadItem__SingleHtmlSaveFormat QWebEngineDownloadItem__SavePageFormat = 0

//
const QWebEngineDownloadItem__CompleteHtmlSaveFormat QWebEngineDownloadItem__SavePageFormat = 1

//
const QWebEngineDownloadItem__MimeHtmlSaveFormat QWebEngineDownloadItem__SavePageFormat = 2

func (this *QWebEngineDownloadItem) SavePageFormatItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QWebEngineDownloadItem_SavePageFormatItemName(val int) string {
	var nilthis *QWebEngineDownloadItem
	return nilthis.SavePageFormatItemName(val)
}

/*


 */
type QWebEngineDownloadItem__DownloadInterruptReason = int

//
const QWebEngineDownloadItem__NoReason QWebEngineDownloadItem__DownloadInterruptReason = 0

//
const QWebEngineDownloadItem__FileFailed QWebEngineDownloadItem__DownloadInterruptReason = 1

//
const QWebEngineDownloadItem__FileAccessDenied QWebEngineDownloadItem__DownloadInterruptReason = 2

//
const QWebEngineDownloadItem__FileNoSpace QWebEngineDownloadItem__DownloadInterruptReason = 3

//
const QWebEngineDownloadItem__FileNameTooLong QWebEngineDownloadItem__DownloadInterruptReason = 5

//
const QWebEngineDownloadItem__FileTooLarge QWebEngineDownloadItem__DownloadInterruptReason = 6

//
const QWebEngineDownloadItem__FileVirusInfected QWebEngineDownloadItem__DownloadInterruptReason = 7

//
const QWebEngineDownloadItem__FileTransientError QWebEngineDownloadItem__DownloadInterruptReason = 10

//
const QWebEngineDownloadItem__FileBlocked QWebEngineDownloadItem__DownloadInterruptReason = 11

//
const QWebEngineDownloadItem__FileSecurityCheckFailed QWebEngineDownloadItem__DownloadInterruptReason = 12

//
const QWebEngineDownloadItem__FileTooShort QWebEngineDownloadItem__DownloadInterruptReason = 13

//
const QWebEngineDownloadItem__FileHashMismatch QWebEngineDownloadItem__DownloadInterruptReason = 14

//
const QWebEngineDownloadItem__NetworkFailed QWebEngineDownloadItem__DownloadInterruptReason = 20

//
const QWebEngineDownloadItem__NetworkTimeout QWebEngineDownloadItem__DownloadInterruptReason = 21

//
const QWebEngineDownloadItem__NetworkDisconnected QWebEngineDownloadItem__DownloadInterruptReason = 22

//
const QWebEngineDownloadItem__NetworkServerDown QWebEngineDownloadItem__DownloadInterruptReason = 23

//
const QWebEngineDownloadItem__NetworkInvalidRequest QWebEngineDownloadItem__DownloadInterruptReason = 24

//
const QWebEngineDownloadItem__ServerFailed QWebEngineDownloadItem__DownloadInterruptReason = 30

//
const QWebEngineDownloadItem__ServerBadContent QWebEngineDownloadItem__DownloadInterruptReason = 33

//
const QWebEngineDownloadItem__ServerUnauthorized QWebEngineDownloadItem__DownloadInterruptReason = 34

//
const QWebEngineDownloadItem__ServerCertProblem QWebEngineDownloadItem__DownloadInterruptReason = 35

//
const QWebEngineDownloadItem__ServerForbidden QWebEngineDownloadItem__DownloadInterruptReason = 36

//
const QWebEngineDownloadItem__ServerUnreachable QWebEngineDownloadItem__DownloadInterruptReason = 37

//
const QWebEngineDownloadItem__UserCanceled QWebEngineDownloadItem__DownloadInterruptReason = 40

func (this *QWebEngineDownloadItem) DownloadInterruptReasonItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QWebEngineDownloadItem_DownloadInterruptReasonItemName(val int) string {
	var nilthis *QWebEngineDownloadItem
	return nilthis.DownloadInterruptReasonItemName(val)
}

/*


 */
type QWebEngineDownloadItem__DownloadType = int

//
const QWebEngineDownloadItem__Attachment QWebEngineDownloadItem__DownloadType = 0

//
const QWebEngineDownloadItem__DownloadAttribute QWebEngineDownloadItem__DownloadType = 1

//
const QWebEngineDownloadItem__UserRequested QWebEngineDownloadItem__DownloadType = 2

//
const QWebEngineDownloadItem__SavePage QWebEngineDownloadItem__DownloadType = 3

func (this *QWebEngineDownloadItem) DownloadTypeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QWebEngineDownloadItem_DownloadTypeItemName(val int) string {
	var nilthis *QWebEngineDownloadItem
	return nilthis.DownloadTypeItemName(val)
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
		qtqml.KeepMe()
	}
	if false {
		qtpositioning.KeepMe()
	}
	if false {
		qtwebchannel.KeepMe()
	}
	if false {
		qtquickwidgets.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtwidgets.KeepMe()
	}
	if false {
		qtprintsupport.KeepMe()
	}
	if false {
		qtquick.KeepMe()
	}
	if false {
		qtwebenginecore.KeepMe()
	}
}

//  keep block end
