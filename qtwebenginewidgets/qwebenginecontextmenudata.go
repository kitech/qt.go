package qtwebenginewidgets

// /usr/include/qt/QtWebEngineWidgets/qwebenginecontextmenudata.h
// #include <qwebenginecontextmenudata.h>
// #include <QtWebEngineWidgets>

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
type QWebEngineContextMenuData struct {
	*qtrt.CObject
}
type QWebEngineContextMenuData_ITF interface {
	QWebEngineContextMenuData_PTR() *QWebEngineContextMenuData
}

func (ptr *QWebEngineContextMenuData) QWebEngineContextMenuData_PTR() *QWebEngineContextMenuData {
	return ptr
}

func (this *QWebEngineContextMenuData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWebEngineContextMenuData) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWebEngineContextMenuDataFromPointer(cthis unsafe.Pointer) *QWebEngineContextMenuData {
	return &QWebEngineContextMenuData{&qtrt.CObject{cthis}}
}
func (*QWebEngineContextMenuData) NewFromPointer(cthis unsafe.Pointer) *QWebEngineContextMenuData {
	return NewQWebEngineContextMenuDataFromPointer(cthis)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginecontextmenudata.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWebEngineContextMenuData()

/*

 */
func (*QWebEngineContextMenuData) NewForInherit() *QWebEngineContextMenuData {
	return NewQWebEngineContextMenuData()
}
func NewQWebEngineContextMenuData() *QWebEngineContextMenuData {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QWebEngineContextMenuDataC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEngineContextMenuDataFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWebEngineContextMenuData)
	return gothis
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginecontextmenudata.h:60
// index:0
// Public Visibility=Default Availability=Available
// [8] QWebEngineContextMenuData & operator=(const QWebEngineContextMenuData &)

/*

 */
func (this *QWebEngineContextMenuData) Operator_equal(other QWebEngineContextMenuData_ITF) *QWebEngineContextMenuData {
	var convArg0 unsafe.Pointer
	if other != nil && other.QWebEngineContextMenuData_PTR() != nil {
		convArg0 = other.QWebEngineContextMenuData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN25QWebEngineContextMenuDataaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQWebEngineContextMenuDataFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQWebEngineContextMenuData)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginecontextmenudata.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QWebEngineContextMenuData()

/*

 */
func DeleteQWebEngineContextMenuData(this *QWebEngineContextMenuData) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QWebEngineContextMenuDataD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginecontextmenudata.h:104
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*

 */
func (this *QWebEngineContextMenuData) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QWebEngineContextMenuData7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginecontextmenudata.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint position() const

/*

 */
func (this *QWebEngineContextMenuData) Position() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QWebEngineContextMenuData8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginecontextmenudata.h:107
// index:0
// Public Visibility=Default Availability=Available
// [8] QString selectedText() const

/*

 */
func (this *QWebEngineContextMenuData) SelectedText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QWebEngineContextMenuData12selectedTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginecontextmenudata.h:108
// index:0
// Public Visibility=Default Availability=Available
// [8] QString linkText() const

/*

 */
func (this *QWebEngineContextMenuData) LinkText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QWebEngineContextMenuData8linkTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginecontextmenudata.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl linkUrl() const

/*

 */
func (this *QWebEngineContextMenuData) LinkUrl() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QWebEngineContextMenuData7linkUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginecontextmenudata.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl mediaUrl() const

/*

 */
func (this *QWebEngineContextMenuData) MediaUrl() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QWebEngineContextMenuData8mediaUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginecontextmenudata.h:111
// index:0
// Public Visibility=Default Availability=Available
// [4] QWebEngineContextMenuData::MediaType mediaType() const

/*

 */
func (this *QWebEngineContextMenuData) MediaType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QWebEngineContextMenuData9mediaTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginecontextmenudata.h:112
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isContentEditable() const

/*

 */
func (this *QWebEngineContextMenuData) IsContentEditable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QWebEngineContextMenuData17isContentEditableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginecontextmenudata.h:113
// index:0
// Public Visibility=Default Availability=Available
// [8] QString misspelledWord() const

/*

 */
func (this *QWebEngineContextMenuData) MisspelledWord() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QWebEngineContextMenuData14misspelledWordEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginecontextmenudata.h:114
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList spellCheckerSuggestions() const

/*

 */
func (this *QWebEngineContextMenuData) SpellCheckerSuggestions() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QWebEngineContextMenuData23spellCheckerSuggestionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginecontextmenudata.h:115
// index:0
// Public Visibility=Default Availability=Available
// [4] QWebEngineContextMenuData::MediaFlags mediaFlags() const

/*

 */
func (this *QWebEngineContextMenuData) MediaFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QWebEngineContextMenuData10mediaFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginecontextmenudata.h:116
// index:0
// Public Visibility=Default Availability=Available
// [4] QWebEngineContextMenuData::EditFlags editFlags() const

/*

 */
func (this *QWebEngineContextMenuData) EditFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QWebEngineContextMenuData9editFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

/*


 */
type QWebEngineContextMenuData__MediaType = int

//
const QWebEngineContextMenuData__MediaTypeNone QWebEngineContextMenuData__MediaType = 0

//
const QWebEngineContextMenuData__MediaTypeImage QWebEngineContextMenuData__MediaType = 1

//
const QWebEngineContextMenuData__MediaTypeVideo QWebEngineContextMenuData__MediaType = 2

//
const QWebEngineContextMenuData__MediaTypeAudio QWebEngineContextMenuData__MediaType = 3

//
const QWebEngineContextMenuData__MediaTypeCanvas QWebEngineContextMenuData__MediaType = 4

//
const QWebEngineContextMenuData__MediaTypeFile QWebEngineContextMenuData__MediaType = 5

//
const QWebEngineContextMenuData__MediaTypePlugin QWebEngineContextMenuData__MediaType = 6

func (this *QWebEngineContextMenuData) MediaTypeItemName(val int) string {
	switch val {
	case QWebEngineContextMenuData__MediaTypeNone: // 0
		return "MediaTypeNone"
	case QWebEngineContextMenuData__MediaTypeImage: // 1
		return "MediaTypeImage"
	case QWebEngineContextMenuData__MediaTypeVideo: // 2
		return "MediaTypeVideo"
	case QWebEngineContextMenuData__MediaTypeAudio: // 3
		return "MediaTypeAudio"
	case QWebEngineContextMenuData__MediaTypeCanvas: // 4
		return "MediaTypeCanvas"
	case QWebEngineContextMenuData__MediaTypeFile: // 5
		return "MediaTypeFile"
	case QWebEngineContextMenuData__MediaTypePlugin: // 6
		return "MediaTypePlugin"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QWebEngineContextMenuData_MediaTypeItemName(val int) string {
	var nilthis *QWebEngineContextMenuData
	return nilthis.MediaTypeItemName(val)
}

/*


 */
type QWebEngineContextMenuData__MediaFlag = int

//
const QWebEngineContextMenuData__MediaInError QWebEngineContextMenuData__MediaFlag = 1

//
const QWebEngineContextMenuData__MediaPaused QWebEngineContextMenuData__MediaFlag = 2

//
const QWebEngineContextMenuData__MediaMuted QWebEngineContextMenuData__MediaFlag = 4

//
const QWebEngineContextMenuData__MediaLoop QWebEngineContextMenuData__MediaFlag = 8

//
const QWebEngineContextMenuData__MediaCanSave QWebEngineContextMenuData__MediaFlag = 16

//
const QWebEngineContextMenuData__MediaHasAudio QWebEngineContextMenuData__MediaFlag = 32

//
const QWebEngineContextMenuData__MediaCanToggleControls QWebEngineContextMenuData__MediaFlag = 64

//
const QWebEngineContextMenuData__MediaControls QWebEngineContextMenuData__MediaFlag = 128

//
const QWebEngineContextMenuData__MediaCanPrint QWebEngineContextMenuData__MediaFlag = 256

//
const QWebEngineContextMenuData__MediaCanRotate QWebEngineContextMenuData__MediaFlag = 512

func (this *QWebEngineContextMenuData) MediaFlagItemName(val int) string {
	switch val {
	case QWebEngineContextMenuData__MediaInError: // 1
		return "MediaInError"
	case QWebEngineContextMenuData__MediaPaused: // 2
		return "MediaPaused"
	case QWebEngineContextMenuData__MediaMuted: // 4
		return "MediaMuted"
	case QWebEngineContextMenuData__MediaLoop: // 8
		return "MediaLoop"
	case QWebEngineContextMenuData__MediaCanSave: // 16
		return "MediaCanSave"
	case QWebEngineContextMenuData__MediaHasAudio: // 32
		return "MediaHasAudio"
	case QWebEngineContextMenuData__MediaCanToggleControls: // 64
		return "MediaCanToggleControls"
	case QWebEngineContextMenuData__MediaControls: // 128
		return "MediaControls"
	case QWebEngineContextMenuData__MediaCanPrint: // 256
		return "MediaCanPrint"
	case QWebEngineContextMenuData__MediaCanRotate: // 512
		return "MediaCanRotate"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QWebEngineContextMenuData_MediaFlagItemName(val int) string {
	var nilthis *QWebEngineContextMenuData
	return nilthis.MediaFlagItemName(val)
}

/*


 */
type QWebEngineContextMenuData__EditFlag = int

//
const QWebEngineContextMenuData__CanUndo QWebEngineContextMenuData__EditFlag = 1

//
const QWebEngineContextMenuData__CanRedo QWebEngineContextMenuData__EditFlag = 2

//
const QWebEngineContextMenuData__CanCut QWebEngineContextMenuData__EditFlag = 4

//
const QWebEngineContextMenuData__CanCopy QWebEngineContextMenuData__EditFlag = 8

//
const QWebEngineContextMenuData__CanPaste QWebEngineContextMenuData__EditFlag = 16

//
const QWebEngineContextMenuData__CanDelete QWebEngineContextMenuData__EditFlag = 32

//
const QWebEngineContextMenuData__CanSelectAll QWebEngineContextMenuData__EditFlag = 64

//
const QWebEngineContextMenuData__CanTranslate QWebEngineContextMenuData__EditFlag = 128

//
const QWebEngineContextMenuData__CanEditRichly QWebEngineContextMenuData__EditFlag = 256

func (this *QWebEngineContextMenuData) EditFlagItemName(val int) string {
	switch val {
	case QWebEngineContextMenuData__CanUndo: // 1
		return "CanUndo"
	case QWebEngineContextMenuData__CanRedo: // 2
		return "CanRedo"
	case QWebEngineContextMenuData__CanCut: // 4
		return "CanCut"
	case QWebEngineContextMenuData__CanCopy: // 8
		return "CanCopy"
	case QWebEngineContextMenuData__CanPaste: // 16
		return "CanPaste"
	case QWebEngineContextMenuData__CanDelete: // 32
		return "CanDelete"
	case QWebEngineContextMenuData__CanSelectAll: // 64
		return "CanSelectAll"
	case QWebEngineContextMenuData__CanTranslate: // 128
		return "CanTranslate"
	case QWebEngineContextMenuData__CanEditRichly: // 256
		return "CanEditRichly"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QWebEngineContextMenuData_EditFlagItemName(val int) string {
	var nilthis *QWebEngineContextMenuData
	return nilthis.EditFlagItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11711() {
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
