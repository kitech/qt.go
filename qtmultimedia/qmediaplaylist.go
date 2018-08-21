package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediaplaylist.h
// #include <qmediaplaylist.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 39
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

// bool setMediaObject(QMediaObject *)
func (this *QMediaPlaylist) InheritSetMediaObject(f func(object *QMediaObject /*777 QMediaObject **/) bool) {
	qtrt.SetAllInheritCallback(this, "setMediaObject", f)
}

/*

 */
type QMediaPlaylist struct {
	*qtcore.QObject
	*QMediaBindableInterface
}
type QMediaPlaylist_ITF interface {
	qtcore.QObject_ITF
	QMediaBindableInterface_ITF
	QMediaPlaylist_PTR() *QMediaPlaylist
}

func (ptr *QMediaPlaylist) QMediaPlaylist_PTR() *QMediaPlaylist { return ptr }

func (this *QMediaPlaylist) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QMediaPlaylist) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
	this.QMediaBindableInterface = NewQMediaBindableInterfaceFromPointer(cthis)
}
func NewQMediaPlaylistFromPointer(cthis unsafe.Pointer) *QMediaPlaylist {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := NewQMediaBindableInterfaceFromPointer(cthis)
	return &QMediaPlaylist{bcthis0, bcthis1}
}
func (*QMediaPlaylist) NewFromPointer(cthis unsafe.Pointer) *QMediaPlaylist {
	return NewQMediaPlaylistFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMediaPlaylist) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaPlaylist10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMediaPlaylist(QObject *)

/*
Create a new playlist object with the given parent.
*/
func NewQMediaPlaylist(parent qtcore.QObject_ITF /*777 QObject **/) *QMediaPlaylist {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylistC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaPlaylistFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaPlaylist")
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMediaPlaylist(QObject *)

/*
Create a new playlist object with the given parent.
*/
func NewQMediaPlaylist__() *QMediaPlaylist {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylistC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaPlaylistFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaPlaylist")
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMediaPlaylist()

/*

 */
func DeleteQMediaPlaylist(this *QMediaPlaylist) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylistD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QMediaObject * mediaObject() const

/*
Reimplemented from QMediaBindableInterface::mediaObject().

Returns the QMediaObject instance that this QMediaPlaylist is bound too, or 0 otherwise.
*/
func (this *QMediaPlaylist) MediaObject() *QMediaObject /*777 QMediaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaPlaylist11mediaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMediaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:75
// index:0
// Public Visibility=Default Availability=Available
// [4] QMediaPlaylist::PlaybackMode playbackMode() const

/*

 */
func (this *QMediaPlaylist) PlaybackMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaPlaylist12playbackModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPlaybackMode(QMediaPlaylist::PlaybackMode)

/*

 */
func (this *QMediaPlaylist) SetPlaybackMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist15setPlaybackModeENS_12PlaybackModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentIndex() const

/*
Returns position of the current media content in the playlist.

Note: Getter function for property currentIndex.

See also setCurrentIndex().
*/
func (this *QMediaPlaylist) CurrentIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaPlaylist12currentIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QMediaContent currentMedia() const

/*
Returns the current media content.

Note: Getter function for property currentMedia.
*/
func (this *QMediaPlaylist) CurrentMedia() *QMediaContent /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaPlaylist12currentMediaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMediaContentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMediaContent)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] int nextIndex(int) const

/*
Returns the index of the item, which would be current after calling next() steps times.

Returned value depends on the size of playlist, current position and playback mode.

See also QMediaPlaylist::playbackMode() and previousIndex().
*/
func (this *QMediaPlaylist) NextIndex(steps int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaPlaylist9nextIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), steps)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] int nextIndex(int) const

/*
Returns the index of the item, which would be current after calling next() steps times.

Returned value depends on the size of playlist, current position and playback mode.

See also QMediaPlaylist::playbackMode() and previousIndex().
*/
func (this *QMediaPlaylist) NextIndex__() int {
	// arg: 0, int=Int, =Invalid, , Invalid
	steps := int(1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaPlaylist9nextIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), steps)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] int previousIndex(int) const

/*
Returns the index of the item, which would be current after calling previous() steps times.

See also QMediaPlaylist::playbackMode() and nextIndex().
*/
func (this *QMediaPlaylist) PreviousIndex(steps int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaPlaylist13previousIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), steps)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] int previousIndex(int) const

/*
Returns the index of the item, which would be current after calling previous() steps times.

See also QMediaPlaylist::playbackMode() and nextIndex().
*/
func (this *QMediaPlaylist) PreviousIndex__() int {
	// arg: 0, int=Int, =Invalid, , Invalid
	steps := int(1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaPlaylist13previousIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), steps)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] QMediaContent media(int) const

/*
Returns the media content at index in the playlist.
*/
func (this *QMediaPlaylist) Media(index int) *QMediaContent /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaPlaylist5mediaEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMediaContentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMediaContent)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:86
// index:0
// Public Visibility=Default Availability=Available
// [4] int mediaCount() const

/*
Returns the number of items in the playlist.

See also isEmpty().
*/
func (this *QMediaPlaylist) MediaCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaPlaylist10mediaCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:87
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if the playlist contains no items, otherwise returns false.

See also mediaCount().
*/
func (this *QMediaPlaylist) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaPlaylist7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:88
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadOnly() const

/*
Returns true if the playlist can be modified, otherwise returns false.

See also mediaCount().
*/
func (this *QMediaPlaylist) IsReadOnly() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaPlaylist10isReadOnlyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:90
// index:0
// Public Visibility=Default Availability=Available
// [1] bool addMedia(const QMediaContent &)

/*
Append the media content to the playlist.

Returns true if the operation is successful, otherwise returns false.
*/
func (this *QMediaPlaylist) AddMedia(content QMediaContent_ITF) bool {
	var convArg0 unsafe.Pointer
	if content != nil && content.QMediaContent_PTR() != nil {
		convArg0 = content.QMediaContent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist8addMediaERK13QMediaContent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool insertMedia(int, const QMediaContent &)

/*
Insert the media content to the playlist at position pos.

Returns true if the operation is successful, otherwise returns false.
*/
func (this *QMediaPlaylist) InsertMedia(index int, content QMediaContent_ITF) bool {
	var convArg1 unsafe.Pointer
	if content != nil && content.QMediaContent_PTR() != nil {
		convArg1 = content.QMediaContent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist11insertMediaEiRK13QMediaContent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool moveMedia(int, int)

/*
Move the item from position from to position to.

Returns true if the operation is successful, otherwise false.

This function was introduced in  Qt 5.7.
*/
func (this *QMediaPlaylist) MoveMedia(from int, to int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist9moveMediaEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), from, to)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:95
// index:0
// Public Visibility=Default Availability=Available
// [1] bool removeMedia(int)

/*
Remove the item from the playlist at position pos.

Returns true if the operation is successful, otherwise return false.
*/
func (this *QMediaPlaylist) RemoveMedia(pos int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist11removeMediaEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:96
// index:1
// Public Visibility=Default Availability=Available
// [1] bool removeMedia(int, int)

/*
Remove the item from the playlist at position pos.

Returns true if the operation is successful, otherwise return false.
*/
func (this *QMediaPlaylist) RemoveMedia_1(start int, end int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist11removeMediaEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), start, end)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:97
// index:0
// Public Visibility=Default Availability=Available
// [1] bool clear()

/*
Remove all the items from the playlist.

Returns true if the operation is successful, otherwise return false.
*/
func (this *QMediaPlaylist) Clear() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void load(const QNetworkRequest &, const char *)

/*
Load playlist using network request. If format is specified, it is used, otherwise format is guessed from playlist name and data.

New items are appended to playlist.

QMediaPlaylist::loaded() signal is emitted if playlist was loaded successfully, otherwise the playlist emits loadFailed().
*/
func (this *QMediaPlaylist) Load(request qtnetwork.QNetworkRequest_ITF, format string) {
	var convArg0 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg0 = request.QNetworkRequest_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist4loadERK15QNetworkRequestPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void load(const QNetworkRequest &, const char *)

/*
Load playlist using network request. If format is specified, it is used, otherwise format is guessed from playlist name and data.

New items are appended to playlist.

QMediaPlaylist::loaded() signal is emitted if playlist was loaded successfully, otherwise the playlist emits loadFailed().
*/
func (this *QMediaPlaylist) Load__(request qtnetwork.QNetworkRequest_ITF) {
	var convArg0 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg0 = request.QNetworkRequest_PTR().GetCthis()
	}
	// arg: 1, const char *=Pointer, =Invalid, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist4loadERK15QNetworkRequestPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:100
// index:1
// Public Visibility=Default Availability=Available
// [-2] void load(const QUrl &, const char *)

/*
Load playlist using network request. If format is specified, it is used, otherwise format is guessed from playlist name and data.

New items are appended to playlist.

QMediaPlaylist::loaded() signal is emitted if playlist was loaded successfully, otherwise the playlist emits loadFailed().
*/
func (this *QMediaPlaylist) Load_1(location qtcore.QUrl_ITF, format string) {
	var convArg0 unsafe.Pointer
	if location != nil && location.QUrl_PTR() != nil {
		convArg0 = location.QUrl_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist4loadERK4QUrlPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:100
// index:1
// Public Visibility=Default Availability=Available
// [-2] void load(const QUrl &, const char *)

/*
Load playlist using network request. If format is specified, it is used, otherwise format is guessed from playlist name and data.

New items are appended to playlist.

QMediaPlaylist::loaded() signal is emitted if playlist was loaded successfully, otherwise the playlist emits loadFailed().
*/
func (this *QMediaPlaylist) Load_1_(location qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if location != nil && location.QUrl_PTR() != nil {
		convArg0 = location.QUrl_PTR().GetCthis()
	}
	// arg: 1, const char *=Pointer, =Invalid, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist4loadERK4QUrlPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:101
// index:2
// Public Visibility=Default Availability=Available
// [-2] void load(QIODevice *, const char *)

/*
Load playlist using network request. If format is specified, it is used, otherwise format is guessed from playlist name and data.

New items are appended to playlist.

QMediaPlaylist::loaded() signal is emitted if playlist was loaded successfully, otherwise the playlist emits loadFailed().
*/
func (this *QMediaPlaylist) Load_2(device qtcore.QIODevice_ITF /*777 QIODevice **/, format string) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist4loadEP9QIODevicePKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:101
// index:2
// Public Visibility=Default Availability=Available
// [-2] void load(QIODevice *, const char *)

/*
Load playlist using network request. If format is specified, it is used, otherwise format is guessed from playlist name and data.

New items are appended to playlist.

QMediaPlaylist::loaded() signal is emitted if playlist was loaded successfully, otherwise the playlist emits loadFailed().
*/
func (this *QMediaPlaylist) Load_2_(device qtcore.QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	// arg: 1, const char *=Pointer, =Invalid, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist4loadEP9QIODevicePKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:103
// index:0
// Public Visibility=Default Availability=Available
// [1] bool save(const QUrl &, const char *)

/*
Save playlist to location. If format is specified, it is used, otherwise format is guessed from location name.

Returns true if playlist was saved successfully, otherwise returns false.
*/
func (this *QMediaPlaylist) Save(location qtcore.QUrl_ITF, format string) bool {
	var convArg0 unsafe.Pointer
	if location != nil && location.QUrl_PTR() != nil {
		convArg0 = location.QUrl_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist4saveERK4QUrlPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:103
// index:0
// Public Visibility=Default Availability=Available
// [1] bool save(const QUrl &, const char *)

/*
Save playlist to location. If format is specified, it is used, otherwise format is guessed from location name.

Returns true if playlist was saved successfully, otherwise returns false.
*/
func (this *QMediaPlaylist) Save__(location qtcore.QUrl_ITF) bool {
	var convArg0 unsafe.Pointer
	if location != nil && location.QUrl_PTR() != nil {
		convArg0 = location.QUrl_PTR().GetCthis()
	}
	// arg: 1, const char *=Pointer, =Invalid, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist4saveERK4QUrlPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:104
// index:1
// Public Visibility=Default Availability=Available
// [1] bool save(QIODevice *, const char *)

/*
Save playlist to location. If format is specified, it is used, otherwise format is guessed from location name.

Returns true if playlist was saved successfully, otherwise returns false.
*/
func (this *QMediaPlaylist) Save_1(device qtcore.QIODevice_ITF /*777 QIODevice **/, format string) bool {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist4saveEP9QIODevicePKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:106
// index:0
// Public Visibility=Default Availability=Available
// [4] QMediaPlaylist::Error error() const

/*
Returns the last error condition.
*/
func (this *QMediaPlaylist) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaPlaylist5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:107
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns the string describing the last error condition.
*/
func (this *QMediaPlaylist) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMediaPlaylist11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void shuffle()

/*
Shuffle items in the playlist.
*/
func (this *QMediaPlaylist) Shuffle() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist7shuffleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void next()

/*
Advance to the next media content in playlist.
*/
func (this *QMediaPlaylist) Next() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist4nextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void previous()

/*
Return to the previous media content in playlist.
*/
func (this *QMediaPlaylist) Previous() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist8previousEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentIndex(int)

/*
Activate media content from playlist at position playlistPosition.

Note: Setter function for property currentIndex.

See also currentIndex().
*/
func (this *QMediaPlaylist) SetCurrentIndex(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist15setCurrentIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentIndexChanged(int)

/*
Signal emitted when playlist position changed to position.

Note: Notifier signal for property currentIndex.
*/
func (this *QMediaPlaylist) CurrentIndexChanged(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist19currentIndexChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void playbackModeChanged(QMediaPlaylist::PlaybackMode)

/*
Signal emitted when playback mode changed to mode.

Note: Notifier signal for property playbackMode.
*/
func (this *QMediaPlaylist) PlaybackModeChanged(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist19playbackModeChangedENS_12PlaybackModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentMediaChanged(const QMediaContent &)

/*
Signal emitted when current media changes to content.

Note: Notifier signal for property currentMedia.
*/
func (this *QMediaPlaylist) CurrentMediaChanged(arg0 QMediaContent_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMediaContent_PTR() != nil {
		convArg0 = arg0.QMediaContent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist19currentMediaChangedERK13QMediaContent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mediaAboutToBeInserted(int, int)

/*
Signal emitted when items are to be inserted at start and ending at end.
*/
func (this *QMediaPlaylist) MediaAboutToBeInserted(start int, end int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist22mediaAboutToBeInsertedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), start, end)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mediaInserted(int, int)

/*
This signal is emitted after media has been inserted into the playlist. The new items are those between start and end inclusive.
*/
func (this *QMediaPlaylist) MediaInserted(start int, end int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist13mediaInsertedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), start, end)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mediaAboutToBeRemoved(int, int)

/*
Signal emitted when item are to be deleted at start and ending at end.
*/
func (this *QMediaPlaylist) MediaAboutToBeRemoved(start int, end int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist21mediaAboutToBeRemovedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), start, end)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mediaRemoved(int, int)

/*
This signal is emitted after media has been removed from the playlist. The removed items are those between start and end inclusive.
*/
func (this *QMediaPlaylist) MediaRemoved(start int, end int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist12mediaRemovedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), start, end)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mediaChanged(int, int)

/*
This signal is emitted after media has been changed in the playlist between start and end positions inclusive.
*/
func (this *QMediaPlaylist) MediaChanged(start int, end int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist12mediaChangedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), start, end)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void loaded()

/*
Signal emitted when playlist finished loading.
*/
func (this *QMediaPlaylist) Loaded() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist6loadedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void loadFailed()

/*
Signal emitted if failed to load playlist.
*/
func (this *QMediaPlaylist) LoadFailed() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist10loadFailedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaplaylist.h:132
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool setMediaObject(QMediaObject *)

/*

 */
func (this *QMediaPlaylist) SetMediaObject(object QMediaObject_ITF /*777 QMediaObject **/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QMediaObject_PTR() != nil {
		convArg0 = object.QMediaObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMediaPlaylist14setMediaObjectEP12QMediaObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*
The QMediaPlaylist::PlaybackMode describes the order items in playlist are played.


*/
type QMediaPlaylist__PlaybackMode = int

// The current item is played only once.
const QMediaPlaylist__CurrentItemOnce QMediaPlaylist__PlaybackMode = 0

// The current item is played repeatedly in a loop.
const QMediaPlaylist__CurrentItemInLoop QMediaPlaylist__PlaybackMode = 1

// Playback starts from the current and moves through each successive item until the last is reached and then stops. The next item is a null item when the last one is currently playing.
const QMediaPlaylist__Sequential QMediaPlaylist__PlaybackMode = 2

// Playback restarts at the first item after the last has finished playing.
const QMediaPlaylist__Loop QMediaPlaylist__PlaybackMode = 3

// Play items in random order.
const QMediaPlaylist__Random QMediaPlaylist__PlaybackMode = 4

func (this *QMediaPlaylist) PlaybackModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QMediaPlaylist_PlaybackModeItemName(val int) string {
	var nilthis *QMediaPlaylist
	return nilthis.PlaybackModeItemName(val)
}

/*
This enum describes the QMediaPlaylist error codes.


*/
type QMediaPlaylist__Error = int

// No errors.
const QMediaPlaylist__NoError QMediaPlaylist__Error = 0

// Format error.
const QMediaPlaylist__FormatError QMediaPlaylist__Error = 1

// Format not supported.
const QMediaPlaylist__FormatNotSupportedError QMediaPlaylist__Error = 2

// Network error.
const QMediaPlaylist__NetworkError QMediaPlaylist__Error = 3

// Access denied error.
const QMediaPlaylist__AccessDeniedError QMediaPlaylist__Error = 4

func (this *QMediaPlaylist) ErrorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QMediaPlaylist_ErrorItemName(val int) string {
	var nilthis *QMediaPlaylist
	return nilthis.ErrorItemName(val)
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
