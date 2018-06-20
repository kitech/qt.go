package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediacontent.h
// #include <qmediacontent.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 30
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
type QMediaContent struct {
	*qtrt.CObject
}
type QMediaContent_ITF interface {
	QMediaContent_PTR() *QMediaContent
}

func (ptr *QMediaContent) QMediaContent_PTR() *QMediaContent { return ptr }

func (this *QMediaContent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMediaContent) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMediaContentFromPointer(cthis unsafe.Pointer) *QMediaContent {
	return &QMediaContent{&qtrt.CObject{cthis}}
}
func (*QMediaContent) NewFromPointer(cthis unsafe.Pointer) *QMediaContent {
	return NewQMediaContentFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediacontent.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMediaContent()

/*
Constructs a null QMediaContent.
*/
func NewQMediaContent() *QMediaContent {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMediaContentC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaContentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMediaContent)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediacontent.h:59
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QMediaContent(const QUrl &)

/*
Constructs a null QMediaContent.
*/
func NewQMediaContent_1(contentUrl qtcore.QUrl_ITF) *QMediaContent {
	var convArg0 unsafe.Pointer
	if contentUrl != nil && contentUrl.QUrl_PTR() != nil {
		convArg0 = contentUrl.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMediaContentC2ERK4QUrl", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaContentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMediaContent)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediacontent.h:60
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QMediaContent(const QNetworkRequest &)

/*
Constructs a null QMediaContent.
*/
func NewQMediaContent_2(contentRequest qtnetwork.QNetworkRequest_ITF) *QMediaContent {
	var convArg0 unsafe.Pointer
	if contentRequest != nil && contentRequest.QNetworkRequest_PTR() != nil {
		convArg0 = contentRequest.QNetworkRequest_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMediaContentC2ERK15QNetworkRequest", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaContentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMediaContent)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediacontent.h:61
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QMediaContent(const QMediaResource &)

/*
Constructs a null QMediaContent.
*/
func NewQMediaContent_3(contentResource QMediaResource_ITF) *QMediaContent {
	var convArg0 unsafe.Pointer
	if contentResource != nil && contentResource.QMediaResource_PTR() != nil {
		convArg0 = contentResource.QMediaResource_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMediaContentC2ERK14QMediaResource", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaContentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMediaContent)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediacontent.h:64
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QMediaContent(QMediaPlaylist *, const QUrl &, bool)

/*
Constructs a null QMediaContent.
*/
func NewQMediaContent_4(playlist QMediaPlaylist_ITF /*777 QMediaPlaylist **/, contentUrl qtcore.QUrl_ITF, takeOwnership bool) *QMediaContent {
	var convArg0 unsafe.Pointer
	if playlist != nil && playlist.QMediaPlaylist_PTR() != nil {
		convArg0 = playlist.QMediaPlaylist_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if contentUrl != nil && contentUrl.QUrl_PTR() != nil {
		convArg1 = contentUrl.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMediaContentC2EP14QMediaPlaylistRK4QUrlb", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, takeOwnership)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaContentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMediaContent)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediacontent.h:64
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QMediaContent(QMediaPlaylist *, const QUrl &, bool)

/*
Constructs a null QMediaContent.
*/
func NewQMediaContent_4_(playlist QMediaPlaylist_ITF /*777 QMediaPlaylist **/) *QMediaContent {
	var convArg0 unsafe.Pointer
	if playlist != nil && playlist.QMediaPlaylist_PTR() != nil {
		convArg0 = playlist.QMediaPlaylist_PTR().GetCthis()
	}
	// arg: 1, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg1 = qtcore.NewQUrl()
	// arg: 2, bool=Bool, =Invalid, , Invalid
	takeOwnership := false
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMediaContentC2EP14QMediaPlaylistRK4QUrlb", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, takeOwnership)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaContentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMediaContent)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediacontent.h:64
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QMediaContent(QMediaPlaylist *, const QUrl &, bool)

/*
Constructs a null QMediaContent.
*/
func NewQMediaContent_4_1(playlist QMediaPlaylist_ITF /*777 QMediaPlaylist **/, contentUrl qtcore.QUrl_ITF) *QMediaContent {
	var convArg0 unsafe.Pointer
	if playlist != nil && playlist.QMediaPlaylist_PTR() != nil {
		convArg0 = playlist.QMediaPlaylist_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if contentUrl != nil && contentUrl.QUrl_PTR() != nil {
		convArg1 = contentUrl.QUrl_PTR().GetCthis()
	}
	// arg: 2, bool=Bool, =Invalid, , Invalid
	takeOwnership := false
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMediaContentC2EP14QMediaPlaylistRK4QUrlb", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, takeOwnership)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaContentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMediaContent)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediacontent.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QMediaContent()

/*

 */
func DeleteQMediaContent(this *QMediaContent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMediaContentD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediacontent.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QMediaContent & operator=(const QMediaContent &)

/*

 */
func (this *QMediaContent) Operator_equal(other QMediaContent_ITF) *QMediaContent {
	var convArg0 unsafe.Pointer
	if other != nil && other.QMediaContent_PTR() != nil {
		convArg0 = other.QMediaContent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMediaContentaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMediaContentFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMediaContent)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediacontent.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QMediaContent &) const

/*

 */
func (this *QMediaContent) Operator_equal_equal(other QMediaContent_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QMediaContent_PTR() != nil {
		convArg0 = other.QMediaContent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMediaContenteqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediacontent.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QMediaContent &) const

/*

 */
func (this *QMediaContent) Operator_not_equal(other QMediaContent_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QMediaContent_PTR() != nil {
		convArg0 = other.QMediaContent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMediaContentneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediacontent.h:72
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if this media content is null (uninitialized); false otherwise.
*/
func (this *QMediaContent) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMediaContent6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediacontent.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl canonicalUrl() const

/*
Returns a QUrl that represents that canonical resource for this media content.
*/
func (this *QMediaContent) CanonicalUrl() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMediaContent12canonicalUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediacontent.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkRequest canonicalRequest() const

/*
Returns a QNetworkRequest that represents that canonical resource for this media content.
*/
func (this *QMediaContent) CanonicalRequest() *qtnetwork.QNetworkRequest /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMediaContent16canonicalRequestEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtnetwork.NewQNetworkRequestFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtnetwork.DeleteQNetworkRequest)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediacontent.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QMediaResource canonicalResource() const

/*
Returns a QMediaResource that represents that canonical resource for this media content.
*/
func (this *QMediaContent) CanonicalResource() *QMediaResource /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMediaContent17canonicalResourceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMediaResourceFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMediaResource)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediacontent.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QMediaResourceList resources() const

/*
Returns a list of alternative resources for this media content. The first item in this list is always the canonical resource.
*/
func (this *QMediaContent) Resources() *QMediaResourceList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMediaContent9resourcesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMediaResourceListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediacontent.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QMediaPlaylist * playlist() const

/*
Returns a playlist for this media content or 0 if this QMediaContent is not a playlist.
*/
func (this *QMediaContent) Playlist() *QMediaPlaylist /*777 QMediaPlaylist **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMediaContent8playlistEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMediaPlaylistFromPointer(unsafe.Pointer(uintptr(rv))) // 444
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
