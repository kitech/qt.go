package qtgui

// /usr/include/qt/QtGui/qbackingstore.h
// #include <qbackingstore.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 145
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QBackingStore struct {
	*qtrt.CObject
}

func (this *QBackingStore) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QBackingStore) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQBackingStoreFromPointer(cthis unsafe.Pointer) *QBackingStore {
	return &QBackingStore{&qtrt.CObject{cthis}}
}
func (*QBackingStore) NewFromPointer(cthis unsafe.Pointer) *QBackingStore {
	return NewQBackingStoreFromPointer(cthis)
}

// /usr/include/qt/QtGui/qbackingstore.h:62
// index:0
// Public
// void QBackingStore(class QWindow *)
func NewQBackingStore(window *QWindow /*777 QWindow **/) *QBackingStore {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = window.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStoreC2EP7QWindow", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQBackingStoreFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbackingstore.h:63
// index:0
// Public
// void ~QBackingStore()
func DeleteQBackingStore(*QBackingStore) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStoreD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:65
// index:0
// Public
// QWindow * window()
func (this *QBackingStore) Window() *QWindow /*777 QWindow **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QBackingStore6windowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qbackingstore.h:67
// index:0
// Public
// QPaintDevice * paintDevice()
func (this *QBackingStore) PaintDevice() *QPaintDevice /*777 QPaintDevice **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStore11paintDeviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qbackingstore.h:69
// index:0
// Public
// void flush(const class QRegion &, class QWindow *, const class QPoint &)
func (this *QBackingStore) Flush(region *QRegion, window *QWindow /*777 QWindow **/, offset *qtcore.QPoint) {
	var convArg0 = region.GetCthis()
	var convArg1 = window.GetCthis()
	var convArg2 = offset.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStore5flushERK7QRegionP7QWindowRK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:71
// index:0
// Public
// void resize(const class QSize &)
func (this *QBackingStore) Resize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStore6resizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:72
// index:0
// Public
// QSize size()
func (this *QBackingStore) Size() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QBackingStore4sizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qbackingstore.h:74
// index:0
// Public
// bool scroll(const class QRegion &, int, int)
func (this *QBackingStore) Scroll(area *QRegion, dx int, dy int) bool {
	var convArg0 = area.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStore6scrollERK7QRegionii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, dx, dy)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qbackingstore.h:76
// index:0
// Public
// void beginPaint(const class QRegion &)
func (this *QBackingStore) BeginPaint(arg0 *QRegion) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStore10beginPaintERK7QRegion", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:77
// index:0
// Public
// void endPaint()
func (this *QBackingStore) EndPaint() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStore8endPaintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:79
// index:0
// Public
// void setStaticContents(const class QRegion &)
func (this *QBackingStore) SetStaticContents(region *QRegion) {
	var convArg0 = region.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStore17setStaticContentsERK7QRegion", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:80
// index:0
// Public
// QRegion staticContents()
func (this *QBackingStore) StaticContents() *QRegion /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QBackingStore14staticContentsEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qbackingstore.h:81
// index:0
// Public
// bool hasStaticContents()
func (this *QBackingStore) HasStaticContents() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QBackingStore17hasStaticContentsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
