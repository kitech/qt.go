//  header block begin
// /usr/include/qt/QtGui/qpicture.h
// #include <qpicture.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
type QPicture struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qpicture.h:59
// index:0
// void QPicture(int)
func NewQPicture(formatVersion int) *QPicture {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPictureC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &formatVersion)
	gopp.ErrPrint(err, rv)
	return &QPicture{cthis}
}

// /usr/include/qt/QtGui/qpicture.h:61
// index:0
// virtual
// void ~QPicture()
func DeleteQPicture(*QPicture) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPictureD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:63
// index:0
// bool isNull()
func (this *QPicture) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPicture6isNullEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:65
// index:0
// virtual
// int devType()
func (this *QPicture) DevType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPicture7devTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:66
// index:0
// uint size()
func (this *QPicture) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPicture4sizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:67
// index:0
// const char * data()
func (this *QPicture) Data() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPicture4dataEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:68
// index:0
// virtual
// void setData(const char *, uint)
func (this *QPicture) SetData(data unsafe.Pointer, size uint) {
	// 0: (, const char * data, uint size), (data, &size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPicture7setDataEPKcj", ffiqt.FFI_TYPE_VOID, this.cthis, data, &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:70
// index:0
// bool play(class QPainter *)
func (this *QPicture) Play(p unsafe.Pointer) {
	// 0: (, QPainter * p), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPicture4playEP8QPainter", ffiqt.FFI_TYPE_VOID, this.cthis, p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:72
// index:0
// bool load(class QIODevice *, const char *)
func (this *QPicture) Load(dev unsafe.Pointer, format unsafe.Pointer) {
	// 0: (, QIODevice * dev, const char * format), (dev, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPicture4loadEP9QIODevicePKc", ffiqt.FFI_TYPE_VOID, this.cthis, dev, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:73
// index:1
// bool load(const class QString &, const char *)
func (this *QPicture) Load_1(fileName unsafe.Pointer, format unsafe.Pointer) {
	// 1: (, const QString & fileName, const char * format), (fileName, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPicture4loadERK7QStringPKc", ffiqt.FFI_TYPE_VOID, this.cthis, fileName, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:74
// index:0
// bool save(class QIODevice *, const char *)
func (this *QPicture) Save(dev unsafe.Pointer, format unsafe.Pointer) {
	// 0: (, QIODevice * dev, const char * format), (dev, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPicture4saveEP9QIODevicePKc", ffiqt.FFI_TYPE_VOID, this.cthis, dev, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:75
// index:1
// bool save(const class QString &, const char *)
func (this *QPicture) Save_1(fileName unsafe.Pointer, format unsafe.Pointer) {
	// 1: (, const QString & fileName, const char * format), (fileName, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPicture4saveERK7QStringPKc", ffiqt.FFI_TYPE_VOID, this.cthis, fileName, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:77
// index:0
// QRect boundingRect()
func (this *QPicture) BoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPicture12boundingRectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:78
// index:0
// void setBoundingRect(const class QRect &)
func (this *QPicture) SetBoundingRect(r unsafe.Pointer) {
	// 0: (, const QRect & r), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPicture15setBoundingRectERK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:85
// index:0
// inline
// void swap(class QPicture &)
func (this *QPicture) Swap(other unsafe.Pointer) {
	// 0: (, QPicture & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPicture4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:87
// index:0
// void detach()
func (this *QPicture) Detach() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPicture6detachEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:88
// index:0
// bool isDetached()
func (this *QPicture) IsDetached() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPicture10isDetachedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:94
// index:0
// static
// const char * pictureFormat(const class QString &)
func (this *QPicture) PictureFormat(fileName unsafe.Pointer) {
	// 0: (const QString & fileName), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPicture13pictureFormatERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPicture_PictureFormat(fileName unsafe.Pointer) {
	// 0: (const QString & fileName), (fileName)
	var nilthis *QPicture
	nilthis.PictureFormat(fileName)
}

// /usr/include/qt/QtGui/qpicture.h:95
// index:0
// static
// QList<QByteArray> inputFormats()
func (this *QPicture) InputFormats() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPicture12inputFormatsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPicture_InputFormats() {
	// 0: (), ()
	var nilthis *QPicture
	nilthis.InputFormats()
}

// /usr/include/qt/QtGui/qpicture.h:96
// index:0
// static
// QList<QByteArray> outputFormats()
func (this *QPicture) OutputFormats() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPicture13outputFormatsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPicture_OutputFormats() {
	// 0: (), ()
	var nilthis *QPicture
	nilthis.OutputFormats()
}

// /usr/include/qt/QtGui/qpicture.h:97
// index:0
// static
// QStringList inputFormatList()
func (this *QPicture) InputFormatList() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPicture15inputFormatListEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPicture_InputFormatList() {
	// 0: (), ()
	var nilthis *QPicture
	nilthis.InputFormatList()
}

// /usr/include/qt/QtGui/qpicture.h:98
// index:0
// static
// QStringList outputFormatList()
func (this *QPicture) OutputFormatList() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPicture16outputFormatListEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPicture_OutputFormatList() {
	// 0: (), ()
	var nilthis *QPicture
	nilthis.OutputFormatList()
}

// /usr/include/qt/QtGui/qpicture.h:101
// index:0
// virtual
// QPaintEngine * paintEngine()
func (this *QPicture) PaintEngine() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPicture11paintEngineEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
