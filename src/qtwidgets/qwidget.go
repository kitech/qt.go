//  header block begin
// /usr/include/qt/QtWidgets/qwidget.h
// #include <qwidget.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QWidget struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qwidget.h:130
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QWidget) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:215
// index:0
// virtual
// void ~QWidget()
func DeleteQWidget(*QWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:217
// index:0
// virtual
// int devType()
func (this *QWidget) DevType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget7devTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:219
// index:0
// WId winId()
func (this *QWidget) WinId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget5winIdEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:220
// index:0
// void createWinId()
func (this *QWidget) CreateWinId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11createWinIdEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:221
// index:0
// inline
// WId internalWinId()
func (this *QWidget) InternalWinId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13internalWinIdEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:222
// index:0
// WId effectiveWinId()
func (this *QWidget) EffectiveWinId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14effectiveWinIdEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:225
// index:0
// QStyle * style()
func (this *QWidget) Style() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget5styleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:226
// index:0
// void setStyle(class QStyle *)
func (this *QWidget) SetStyle(arg0 unsafe.Pointer) {
	// 0: (, QStyle * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget8setStyleEP6QStyle", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:229
// index:0
// bool isTopLevel()
func (this *QWidget) IsTopLevel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget10isTopLevelEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:230
// index:0
// bool isWindow()
func (this *QWidget) IsWindow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget8isWindowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:232
// index:0
// bool isModal()
func (this *QWidget) IsModal() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget7isModalEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:233
// index:0
// Qt::WindowModality windowModality()
func (this *QWidget) WindowModality() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14windowModalityEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:234
// index:0
// void setWindowModality(Qt::WindowModality)
func (this *QWidget) SetWindowModality(windowModality int) {
	// 0: (, Qt::WindowModality windowModality), (&windowModality)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17setWindowModalityEN2Qt14WindowModalityE", ffiqt.FFI_TYPE_VOID, this.cthis, &windowModality)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:236
// index:0
// bool isEnabled()
func (this *QWidget) IsEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget9isEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:237
// index:0
// bool isEnabledTo(const class QWidget *)
func (this *QWidget) IsEnabledTo(arg0 unsafe.Pointer) {
	// 0: (, const QWidget * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11isEnabledToEPKS_", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:238
// index:0
// bool isEnabledToTLW()
func (this *QWidget) IsEnabledToTLW() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14isEnabledToTLWEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:241
// index:0
// void setEnabled(_Bool)
func (this *QWidget) SetEnabled(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget10setEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:242
// index:0
// void setDisabled(_Bool)
func (this *QWidget) SetDisabled(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11setDisabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:243
// index:0
// void setWindowModified(_Bool)
func (this *QWidget) SetWindowModified(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17setWindowModifiedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:248
// index:0
// QRect frameGeometry()
func (this *QWidget) FrameGeometry() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13frameGeometryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:249
// index:0
// const QRect & geometry()
func (this *QWidget) Geometry() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget8geometryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:250
// index:0
// QRect normalGeometry()
func (this *QWidget) NormalGeometry() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14normalGeometryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:252
// index:0
// int x()
func (this *QWidget) X() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget1xEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:253
// index:0
// int y()
func (this *QWidget) Y() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget1yEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:254
// index:0
// QPoint pos()
func (this *QWidget) Pos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget3posEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:255
// index:0
// QSize frameSize()
func (this *QWidget) FrameSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget9frameSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:256
// index:0
// QSize size()
func (this *QWidget) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget4sizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:257
// index:0
// inline
// int width()
func (this *QWidget) Width() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget5widthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:258
// index:0
// inline
// int height()
func (this *QWidget) Height() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget6heightEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:259
// index:0
// inline
// QRect rect()
func (this *QWidget) Rect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget4rectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:260
// index:0
// QRect childrenRect()
func (this *QWidget) ChildrenRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget12childrenRectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:261
// index:0
// QRegion childrenRegion()
func (this *QWidget) ChildrenRegion() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14childrenRegionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:263
// index:0
// QSize minimumSize()
func (this *QWidget) MinimumSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11minimumSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:264
// index:0
// QSize maximumSize()
func (this *QWidget) MaximumSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11maximumSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:265
// index:0
// int minimumWidth()
func (this *QWidget) MinimumWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget12minimumWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:266
// index:0
// int minimumHeight()
func (this *QWidget) MinimumHeight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13minimumHeightEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:267
// index:0
// int maximumWidth()
func (this *QWidget) MaximumWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget12maximumWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:268
// index:0
// int maximumHeight()
func (this *QWidget) MaximumHeight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13maximumHeightEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:269
// index:0
// void setMinimumSize(const class QSize &)
func (this *QWidget) SetMinimumSize(arg0 unsafe.Pointer) {
	// 0: (, const QSize & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14setMinimumSizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:270
// index:1
// void setMinimumSize(int, int)
func (this *QWidget) SetMinimumSize_1(minw int, minh int) {
	// 1: (, int minw, int minh), (&minw, &minh)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14setMinimumSizeEii", ffiqt.FFI_TYPE_VOID, this.cthis, &minw, &minh)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:271
// index:0
// void setMaximumSize(const class QSize &)
func (this *QWidget) SetMaximumSize(arg0 unsafe.Pointer) {
	// 0: (, const QSize & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14setMaximumSizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:272
// index:1
// void setMaximumSize(int, int)
func (this *QWidget) SetMaximumSize_1(maxw int, maxh int) {
	// 1: (, int maxw, int maxh), (&maxw, &maxh)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14setMaximumSizeEii", ffiqt.FFI_TYPE_VOID, this.cthis, &maxw, &maxh)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:273
// index:0
// void setMinimumWidth(int)
func (this *QWidget) SetMinimumWidth(minw int) {
	// 0: (, int minw), (&minw)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget15setMinimumWidthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &minw)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:274
// index:0
// void setMinimumHeight(int)
func (this *QWidget) SetMinimumHeight(minh int) {
	// 0: (, int minh), (&minh)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget16setMinimumHeightEi", ffiqt.FFI_TYPE_VOID, this.cthis, &minh)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:275
// index:0
// void setMaximumWidth(int)
func (this *QWidget) SetMaximumWidth(maxw int) {
	// 0: (, int maxw), (&maxw)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget15setMaximumWidthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &maxw)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:276
// index:0
// void setMaximumHeight(int)
func (this *QWidget) SetMaximumHeight(maxh int) {
	// 0: (, int maxh), (&maxh)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget16setMaximumHeightEi", ffiqt.FFI_TYPE_VOID, this.cthis, &maxh)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:282
// index:0
// QSize sizeIncrement()
func (this *QWidget) SizeIncrement() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13sizeIncrementEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:283
// index:0
// void setSizeIncrement(const class QSize &)
func (this *QWidget) SetSizeIncrement(arg0 unsafe.Pointer) {
	// 0: (, const QSize & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget16setSizeIncrementERK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:284
// index:1
// void setSizeIncrement(int, int)
func (this *QWidget) SetSizeIncrement_1(w int, h int) {
	// 1: (, int w, int h), (&w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget16setSizeIncrementEii", ffiqt.FFI_TYPE_VOID, this.cthis, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:285
// index:0
// QSize baseSize()
func (this *QWidget) BaseSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget8baseSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:286
// index:0
// void setBaseSize(const class QSize &)
func (this *QWidget) SetBaseSize(arg0 unsafe.Pointer) {
	// 0: (, const QSize & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11setBaseSizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:287
// index:1
// void setBaseSize(int, int)
func (this *QWidget) SetBaseSize_1(basew int, baseh int) {
	// 1: (, int basew, int baseh), (&basew, &baseh)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11setBaseSizeEii", ffiqt.FFI_TYPE_VOID, this.cthis, &basew, &baseh)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:289
// index:0
// void setFixedSize(const class QSize &)
func (this *QWidget) SetFixedSize(arg0 unsafe.Pointer) {
	// 0: (, const QSize & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12setFixedSizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:290
// index:1
// void setFixedSize(int, int)
func (this *QWidget) SetFixedSize_1(w int, h int) {
	// 1: (, int w, int h), (&w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12setFixedSizeEii", ffiqt.FFI_TYPE_VOID, this.cthis, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:291
// index:0
// void setFixedWidth(int)
func (this *QWidget) SetFixedWidth(w int) {
	// 0: (, int w), (&w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13setFixedWidthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:292
// index:0
// void setFixedHeight(int)
func (this *QWidget) SetFixedHeight(h int) {
	// 0: (, int h), (&h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14setFixedHeightEi", ffiqt.FFI_TYPE_VOID, this.cthis, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:296
// index:0
// QPoint mapToGlobal(const class QPoint &)
func (this *QWidget) MapToGlobal(arg0 unsafe.Pointer) {
	// 0: (, const QPoint & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11mapToGlobalERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:297
// index:0
// QPoint mapFromGlobal(const class QPoint &)
func (this *QWidget) MapFromGlobal(arg0 unsafe.Pointer) {
	// 0: (, const QPoint & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13mapFromGlobalERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:298
// index:0
// QPoint mapToParent(const class QPoint &)
func (this *QWidget) MapToParent(arg0 unsafe.Pointer) {
	// 0: (, const QPoint & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11mapToParentERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:299
// index:0
// QPoint mapFromParent(const class QPoint &)
func (this *QWidget) MapFromParent(arg0 unsafe.Pointer) {
	// 0: (, const QPoint & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13mapFromParentERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:300
// index:0
// QPoint mapTo(const class QWidget *, const class QPoint &)
func (this *QWidget) MapTo(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	// 0: (, const QWidget * arg0, const QPoint & arg1), (arg0, arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget5mapToEPKS_RK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, arg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:301
// index:0
// QPoint mapFrom(const class QWidget *, const class QPoint &)
func (this *QWidget) MapFrom(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	// 0: (, const QWidget * arg0, const QPoint & arg1), (arg0, arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget7mapFromEPKS_RK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, arg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:303
// index:0
// QWidget * window()
func (this *QWidget) Window() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget6windowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:304
// index:0
// QWidget * nativeParentWidget()
func (this *QWidget) NativeParentWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget18nativeParentWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:305
// index:0
// inline
// QWidget * topLevelWidget()
func (this *QWidget) TopLevelWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14topLevelWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:308
// index:0
// const QPalette & palette()
func (this *QWidget) Palette() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget7paletteEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:309
// index:0
// void setPalette(const class QPalette &)
func (this *QWidget) SetPalette(arg0 unsafe.Pointer) {
	// 0: (, const QPalette & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget10setPaletteERK8QPalette", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:311
// index:0
// void setBackgroundRole(class QPalette::ColorRole)
func (this *QWidget) SetBackgroundRole(arg0 int) {
	// 0: (, QPalette::ColorRole arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17setBackgroundRoleEN8QPalette9ColorRoleE", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:312
// index:0
// QPalette::ColorRole backgroundRole()
func (this *QWidget) BackgroundRole() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14backgroundRoleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:314
// index:0
// void setForegroundRole(class QPalette::ColorRole)
func (this *QWidget) SetForegroundRole(arg0 int) {
	// 0: (, QPalette::ColorRole arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17setForegroundRoleEN8QPalette9ColorRoleE", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:315
// index:0
// QPalette::ColorRole foregroundRole()
func (this *QWidget) ForegroundRole() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14foregroundRoleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:317
// index:0
// const QFont & font()
func (this *QWidget) Font() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget4fontEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:318
// index:0
// void setFont(const class QFont &)
func (this *QWidget) SetFont(arg0 unsafe.Pointer) {
	// 0: (, const QFont & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget7setFontERK5QFont", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:319
// index:0
// QFontMetrics fontMetrics()
func (this *QWidget) FontMetrics() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11fontMetricsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:320
// index:0
// QFontInfo fontInfo()
func (this *QWidget) FontInfo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget8fontInfoEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:323
// index:0
// QCursor cursor()
func (this *QWidget) Cursor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget6cursorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:324
// index:0
// void setCursor(const class QCursor &)
func (this *QWidget) SetCursor(arg0 unsafe.Pointer) {
	// 0: (, const QCursor & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9setCursorERK7QCursor", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:325
// index:0
// void unsetCursor()
func (this *QWidget) UnsetCursor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11unsetCursorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:328
// index:0
// void setMouseTracking(_Bool)
func (this *QWidget) SetMouseTracking(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget16setMouseTrackingEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:329
// index:0
// bool hasMouseTracking()
func (this *QWidget) HasMouseTracking() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget16hasMouseTrackingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:330
// index:0
// bool underMouse()
func (this *QWidget) UnderMouse() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget10underMouseEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:332
// index:0
// void setTabletTracking(_Bool)
func (this *QWidget) SetTabletTracking(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17setTabletTrackingEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:333
// index:0
// bool hasTabletTracking()
func (this *QWidget) HasTabletTracking() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget17hasTabletTrackingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:335
// index:0
// void setMask(const class QBitmap &)
func (this *QWidget) SetMask(arg0 unsafe.Pointer) {
	// 0: (, const QBitmap & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget7setMaskERK7QBitmap", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:336
// index:1
// void setMask(const class QRegion &)
func (this *QWidget) SetMask_1(arg0 unsafe.Pointer) {
	// 1: (, const QRegion & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget7setMaskERK7QRegion", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:337
// index:0
// QRegion mask()
func (this *QWidget) Mask() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget4maskEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:338
// index:0
// void clearMask()
func (this *QWidget) ClearMask() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9clearMaskEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:348
// index:0
// QPixmap grab(const class QRect &)
func (this *QWidget) Grab(rectangle unsafe.Pointer) {
	// 0: (, const QRect & rectangle), (rectangle)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget4grabERK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, rectangle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:351
// index:0
// QGraphicsEffect * graphicsEffect()
func (this *QWidget) GraphicsEffect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14graphicsEffectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:352
// index:0
// void setGraphicsEffect(class QGraphicsEffect *)
func (this *QWidget) SetGraphicsEffect(effect unsafe.Pointer) {
	// 0: (, QGraphicsEffect * effect), (effect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17setGraphicsEffectEP15QGraphicsEffect", ffiqt.FFI_TYPE_VOID, this.cthis, effect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:357
// index:0
// void ungrabGesture(Qt::GestureType)
func (this *QWidget) UngrabGesture(type_ int) {
	// 0: (, Qt::GestureType type), (&type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13ungrabGestureEN2Qt11GestureTypeE", ffiqt.FFI_TYPE_VOID, this.cthis, &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:361
// index:0
// void setWindowTitle(const class QString &)
func (this *QWidget) SetWindowTitle(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14setWindowTitleERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:363
// index:0
// void setStyleSheet(const class QString &)
func (this *QWidget) SetStyleSheet(styleSheet unsafe.Pointer) {
	// 0: (, const QString & styleSheet), (styleSheet)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13setStyleSheetERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, styleSheet)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:367
// index:0
// QString styleSheet()
func (this *QWidget) StyleSheet() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget10styleSheetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:369
// index:0
// QString windowTitle()
func (this *QWidget) WindowTitle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11windowTitleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:370
// index:0
// void setWindowIcon(const class QIcon &)
func (this *QWidget) SetWindowIcon(icon unsafe.Pointer) {
	// 0: (, const QIcon & icon), (icon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13setWindowIconERK5QIcon", ffiqt.FFI_TYPE_VOID, this.cthis, icon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:371
// index:0
// QIcon windowIcon()
func (this *QWidget) WindowIcon() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget10windowIconEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:372
// index:0
// void setWindowIconText(const class QString &)
func (this *QWidget) SetWindowIconText(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17setWindowIconTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:373
// index:0
// QString windowIconText()
func (this *QWidget) WindowIconText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14windowIconTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:374
// index:0
// void setWindowRole(const class QString &)
func (this *QWidget) SetWindowRole(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13setWindowRoleERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:375
// index:0
// QString windowRole()
func (this *QWidget) WindowRole() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget10windowRoleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:376
// index:0
// void setWindowFilePath(const class QString &)
func (this *QWidget) SetWindowFilePath(filePath unsafe.Pointer) {
	// 0: (, const QString & filePath), (filePath)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17setWindowFilePathERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, filePath)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:377
// index:0
// QString windowFilePath()
func (this *QWidget) WindowFilePath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14windowFilePathEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:379
// index:0
// void setWindowOpacity(qreal)
func (this *QWidget) SetWindowOpacity(level float64) {
	// 0: (, qreal level), (&level)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget16setWindowOpacityEd", ffiqt.FFI_TYPE_VOID, this.cthis, &level)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:380
// index:0
// qreal windowOpacity()
func (this *QWidget) WindowOpacity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13windowOpacityEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:382
// index:0
// bool isWindowModified()
func (this *QWidget) IsWindowModified() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget16isWindowModifiedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:384
// index:0
// void setToolTip(const class QString &)
func (this *QWidget) SetToolTip(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget10setToolTipERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:385
// index:0
// QString toolTip()
func (this *QWidget) ToolTip() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget7toolTipEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:386
// index:0
// void setToolTipDuration(int)
func (this *QWidget) SetToolTipDuration(msec int) {
	// 0: (, int msec), (&msec)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget18setToolTipDurationEi", ffiqt.FFI_TYPE_VOID, this.cthis, &msec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:387
// index:0
// int toolTipDuration()
func (this *QWidget) ToolTipDuration() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget15toolTipDurationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:390
// index:0
// void setStatusTip(const class QString &)
func (this *QWidget) SetStatusTip(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12setStatusTipERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:391
// index:0
// QString statusTip()
func (this *QWidget) StatusTip() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget9statusTipEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:394
// index:0
// void setWhatsThis(const class QString &)
func (this *QWidget) SetWhatsThis(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12setWhatsThisERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:395
// index:0
// QString whatsThis()
func (this *QWidget) WhatsThis() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget9whatsThisEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:398
// index:0
// QString accessibleName()
func (this *QWidget) AccessibleName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14accessibleNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:399
// index:0
// void setAccessibleName(const class QString &)
func (this *QWidget) SetAccessibleName(name unsafe.Pointer) {
	// 0: (, const QString & name), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17setAccessibleNameERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:400
// index:0
// QString accessibleDescription()
func (this *QWidget) AccessibleDescription() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget21accessibleDescriptionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:401
// index:0
// void setAccessibleDescription(const class QString &)
func (this *QWidget) SetAccessibleDescription(description unsafe.Pointer) {
	// 0: (, const QString & description), (description)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget24setAccessibleDescriptionERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, description)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:404
// index:0
// void setLayoutDirection(Qt::LayoutDirection)
func (this *QWidget) SetLayoutDirection(direction int) {
	// 0: (, Qt::LayoutDirection direction), (&direction)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget18setLayoutDirectionEN2Qt15LayoutDirectionE", ffiqt.FFI_TYPE_VOID, this.cthis, &direction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:405
// index:0
// Qt::LayoutDirection layoutDirection()
func (this *QWidget) LayoutDirection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget15layoutDirectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:406
// index:0
// void unsetLayoutDirection()
func (this *QWidget) UnsetLayoutDirection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget20unsetLayoutDirectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:408
// index:0
// void setLocale(const class QLocale &)
func (this *QWidget) SetLocale(locale unsafe.Pointer) {
	// 0: (, const QLocale & locale), (locale)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9setLocaleERK7QLocale", ffiqt.FFI_TYPE_VOID, this.cthis, locale)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:409
// index:0
// QLocale locale()
func (this *QWidget) Locale() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget6localeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:410
// index:0
// void unsetLocale()
func (this *QWidget) UnsetLocale() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11unsetLocaleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:412
// index:0
// inline
// bool isRightToLeft()
func (this *QWidget) IsRightToLeft() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13isRightToLeftEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:413
// index:0
// inline
// bool isLeftToRight()
func (this *QWidget) IsLeftToRight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13isLeftToRightEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:416
// index:0
// inline
// void setFocus()
func (this *QWidget) SetFocus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget8setFocusEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:423
// index:1
// void setFocus(Qt::FocusReason)
func (this *QWidget) SetFocus_1(reason int) {
	// 1: (, Qt::FocusReason reason), (&reason)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget8setFocusEN2Qt11FocusReasonE", ffiqt.FFI_TYPE_VOID, this.cthis, &reason)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:419
// index:0
// bool isActiveWindow()
func (this *QWidget) IsActiveWindow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14isActiveWindowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:420
// index:0
// void activateWindow()
func (this *QWidget) ActivateWindow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14activateWindowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:421
// index:0
// void clearFocus()
func (this *QWidget) ClearFocus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget10clearFocusEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:424
// index:0
// Qt::FocusPolicy focusPolicy()
func (this *QWidget) FocusPolicy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11focusPolicyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:425
// index:0
// void setFocusPolicy(Qt::FocusPolicy)
func (this *QWidget) SetFocusPolicy(policy int) {
	// 0: (, Qt::FocusPolicy policy), (&policy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14setFocusPolicyEN2Qt11FocusPolicyE", ffiqt.FFI_TYPE_VOID, this.cthis, &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:426
// index:0
// bool hasFocus()
func (this *QWidget) HasFocus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget8hasFocusEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:427
// index:0
// static
// void setTabOrder(class QWidget *, class QWidget *)
func (this *QWidget) SetTabOrder(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	// 0: (QWidget * arg0, QWidget * arg1), (arg0, arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11setTabOrderEPS_S0_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QWidget_SetTabOrder(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	// 0: (QWidget * arg0, QWidget * arg1), (arg0, arg1)
	var nilthis *QWidget
	nilthis.SetTabOrder(arg0, arg1)
}

// /usr/include/qt/QtWidgets/qwidget.h:428
// index:0
// void setFocusProxy(class QWidget *)
func (this *QWidget) SetFocusProxy(arg0 unsafe.Pointer) {
	// 0: (, QWidget * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13setFocusProxyEPS_", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:429
// index:0
// QWidget * focusProxy()
func (this *QWidget) FocusProxy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget10focusProxyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:430
// index:0
// Qt::ContextMenuPolicy contextMenuPolicy()
func (this *QWidget) ContextMenuPolicy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget17contextMenuPolicyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:431
// index:0
// void setContextMenuPolicy(Qt::ContextMenuPolicy)
func (this *QWidget) SetContextMenuPolicy(policy int) {
	// 0: (, Qt::ContextMenuPolicy policy), (&policy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget20setContextMenuPolicyEN2Qt17ContextMenuPolicyE", ffiqt.FFI_TYPE_VOID, this.cthis, &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:434
// index:0
// void grabMouse()
func (this *QWidget) GrabMouse() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9grabMouseEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:436
// index:1
// void grabMouse(const class QCursor &)
func (this *QWidget) GrabMouse_1(arg0 unsafe.Pointer) {
	// 1: (, const QCursor & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9grabMouseERK7QCursor", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:438
// index:0
// void releaseMouse()
func (this *QWidget) ReleaseMouse() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12releaseMouseEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:439
// index:0
// void grabKeyboard()
func (this *QWidget) GrabKeyboard() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12grabKeyboardEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:440
// index:0
// void releaseKeyboard()
func (this *QWidget) ReleaseKeyboard() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget15releaseKeyboardEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:442
// index:0
// int grabShortcut(const class QKeySequence &, Qt::ShortcutContext)
func (this *QWidget) GrabShortcut(key unsafe.Pointer, context int) {
	// 0: (, const QKeySequence & key, Qt::ShortcutContext context), (key, &context)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12grabShortcutERK12QKeySequenceN2Qt15ShortcutContextE", ffiqt.FFI_TYPE_VOID, this.cthis, key, &context)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:443
// index:0
// void releaseShortcut(int)
func (this *QWidget) ReleaseShortcut(id int) {
	// 0: (, int id), (&id)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget15releaseShortcutEi", ffiqt.FFI_TYPE_VOID, this.cthis, &id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:444
// index:0
// void setShortcutEnabled(int, _Bool)
func (this *QWidget) SetShortcutEnabled(id int, enable bool) {
	// 0: (, int id, bool enable), (&id, &enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget18setShortcutEnabledEib", ffiqt.FFI_TYPE_VOID, this.cthis, &id, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:445
// index:0
// void setShortcutAutoRepeat(int, _Bool)
func (this *QWidget) SetShortcutAutoRepeat(id int, enable bool) {
	// 0: (, int id, bool enable), (&id, &enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget21setShortcutAutoRepeatEib", ffiqt.FFI_TYPE_VOID, this.cthis, &id, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:447
// index:0
// static
// QWidget * mouseGrabber()
func (this *QWidget) MouseGrabber() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12mouseGrabberEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QWidget_MouseGrabber() {
	// 0: (), ()
	var nilthis *QWidget
	nilthis.MouseGrabber()
}

// /usr/include/qt/QtWidgets/qwidget.h:448
// index:0
// static
// QWidget * keyboardGrabber()
func (this *QWidget) KeyboardGrabber() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget15keyboardGrabberEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QWidget_KeyboardGrabber() {
	// 0: (), ()
	var nilthis *QWidget
	nilthis.KeyboardGrabber()
}

// /usr/include/qt/QtWidgets/qwidget.h:451
// index:0
// inline
// bool updatesEnabled()
func (this *QWidget) UpdatesEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14updatesEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:452
// index:0
// void setUpdatesEnabled(_Bool)
func (this *QWidget) SetUpdatesEnabled(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17setUpdatesEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:455
// index:0
// QGraphicsProxyWidget * graphicsProxyWidget()
func (this *QWidget) GraphicsProxyWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget19graphicsProxyWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:459
// index:0
// void update()
func (this *QWidget) Update() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget6updateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:463
// index:1
// inline
// void update(int, int, int, int)
func (this *QWidget) Update_1(x int, y int, w int, h int) {
	// 1: (, int x, int y, int w, int h), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget6updateEiiii", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:464
// index:2
// void update(const class QRect &)
func (this *QWidget) Update_2(arg0 unsafe.Pointer) {
	// 2: (, const QRect & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget6updateERK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:465
// index:3
// void update(const class QRegion &)
func (this *QWidget) Update_3(arg0 unsafe.Pointer) {
	// 3: (, const QRegion & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget6updateERK7QRegion", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:460
// index:0
// void repaint()
func (this *QWidget) Repaint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget7repaintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:467
// index:1
// void repaint(int, int, int, int)
func (this *QWidget) Repaint_1(x int, y int, w int, h int) {
	// 1: (, int x, int y, int w, int h), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget7repaintEiiii", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:468
// index:2
// void repaint(const class QRect &)
func (this *QWidget) Repaint_2(arg0 unsafe.Pointer) {
	// 2: (, const QRect & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget7repaintERK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:469
// index:3
// void repaint(const class QRegion &)
func (this *QWidget) Repaint_3(arg0 unsafe.Pointer) {
	// 3: (, const QRegion & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget7repaintERK7QRegion", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:474
// index:0
// virtual
// void setVisible(_Bool)
func (this *QWidget) SetVisible(visible bool) {
	// 0: (, bool visible), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget10setVisibleEb", ffiqt.FFI_TYPE_VOID, this.cthis, &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:475
// index:0
// void setHidden(_Bool)
func (this *QWidget) SetHidden(hidden bool) {
	// 0: (, bool hidden), (&hidden)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9setHiddenEb", ffiqt.FFI_TYPE_VOID, this.cthis, &hidden)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:476
// index:0
// void show()
func (this *QWidget) Show() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget4showEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:477
// index:0
// void hide()
func (this *QWidget) Hide() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget4hideEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:479
// index:0
// void showMinimized()
func (this *QWidget) ShowMinimized() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13showMinimizedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:480
// index:0
// void showMaximized()
func (this *QWidget) ShowMaximized() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13showMaximizedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:481
// index:0
// void showFullScreen()
func (this *QWidget) ShowFullScreen() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14showFullScreenEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:482
// index:0
// void showNormal()
func (this *QWidget) ShowNormal() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget10showNormalEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:484
// index:0
// bool close()
func (this *QWidget) Close() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget5closeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:485
// index:0
// void raise()
func (this *QWidget) Raise() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget5raiseEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:486
// index:0
// void lower()
func (this *QWidget) Lower() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget5lowerEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:489
// index:0
// void stackUnder(class QWidget *)
func (this *QWidget) StackUnder(arg0 unsafe.Pointer) {
	// 0: (, QWidget * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget10stackUnderEPS_", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:490
// index:0
// void move(int, int)
func (this *QWidget) Move(x int, y int) {
	// 0: (, int x, int y), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget4moveEii", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:491
// index:1
// void move(const class QPoint &)
func (this *QWidget) Move_1(arg0 unsafe.Pointer) {
	// 1: (, const QPoint & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget4moveERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:492
// index:0
// void resize(int, int)
func (this *QWidget) Resize(w int, h int) {
	// 0: (, int w, int h), (&w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget6resizeEii", ffiqt.FFI_TYPE_VOID, this.cthis, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:493
// index:1
// void resize(const class QSize &)
func (this *QWidget) Resize_1(arg0 unsafe.Pointer) {
	// 1: (, const QSize & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget6resizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:494
// index:0
// inline
// void setGeometry(int, int, int, int)
func (this *QWidget) SetGeometry(x int, y int, w int, h int) {
	// 0: (, int x, int y, int w, int h), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11setGeometryEiiii", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:495
// index:1
// void setGeometry(const class QRect &)
func (this *QWidget) SetGeometry_1(arg0 unsafe.Pointer) {
	// 1: (, const QRect & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11setGeometryERK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:496
// index:0
// QByteArray saveGeometry()
func (this *QWidget) SaveGeometry() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget12saveGeometryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:497
// index:0
// bool restoreGeometry(const class QByteArray &)
func (this *QWidget) RestoreGeometry(geometry unsafe.Pointer) {
	// 0: (, const QByteArray & geometry), (geometry)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget15restoreGeometryERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.cthis, geometry)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:498
// index:0
// void adjustSize()
func (this *QWidget) AdjustSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget10adjustSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:499
// index:0
// bool isVisible()
func (this *QWidget) IsVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget9isVisibleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:500
// index:0
// bool isVisibleTo(const class QWidget *)
func (this *QWidget) IsVisibleTo(arg0 unsafe.Pointer) {
	// 0: (, const QWidget * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11isVisibleToEPKS_", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:501
// index:0
// inline
// bool isHidden()
func (this *QWidget) IsHidden() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget8isHiddenEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:503
// index:0
// bool isMinimized()
func (this *QWidget) IsMinimized() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11isMinimizedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:504
// index:0
// bool isMaximized()
func (this *QWidget) IsMaximized() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11isMaximizedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:505
// index:0
// bool isFullScreen()
func (this *QWidget) IsFullScreen() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget12isFullScreenEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:507
// index:0
// Qt::WindowStates windowState()
func (this *QWidget) WindowState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11windowStateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:511
// index:0
// virtual
// QSize sizeHint()
func (this *QWidget) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:512
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QWidget) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:514
// index:0
// QSizePolicy sizePolicy()
func (this *QWidget) SizePolicy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget10sizePolicyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:515
// index:0
// void setSizePolicy(class QSizePolicy)
func (this *QWidget) SetSizePolicy(arg0 unsafe.Pointer) {
	// 0: (, QSizePolicy arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13setSizePolicyE11QSizePolicy", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:516
// index:1
// inline
// void setSizePolicy(class QSizePolicy::Policy, class QSizePolicy::Policy)
func (this *QWidget) SetSizePolicy_1(horizontal int, vertical int) {
	// 1: (, QSizePolicy::Policy horizontal, QSizePolicy::Policy vertical), (&horizontal, &vertical)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13setSizePolicyEN11QSizePolicy6PolicyES1_", ffiqt.FFI_TYPE_VOID, this.cthis, &horizontal, &vertical)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:517
// index:0
// virtual
// int heightForWidth(int)
func (this *QWidget) HeightForWidth(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14heightForWidthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:518
// index:0
// virtual
// bool hasHeightForWidth()
func (this *QWidget) HasHeightForWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget17hasHeightForWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:520
// index:0
// QRegion visibleRegion()
func (this *QWidget) VisibleRegion() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13visibleRegionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:522
// index:0
// void setContentsMargins(int, int, int, int)
func (this *QWidget) SetContentsMargins(left int, top int, right int, bottom int) {
	// 0: (, int left, int top, int right, int bottom), (&left, &top, &right, &bottom)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget18setContentsMarginsEiiii", ffiqt.FFI_TYPE_VOID, this.cthis, &left, &top, &right, &bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:523
// index:1
// void setContentsMargins(const class QMargins &)
func (this *QWidget) SetContentsMargins_1(margins unsafe.Pointer) {
	// 1: (, const QMargins & margins), (margins)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget18setContentsMarginsERK8QMargins", ffiqt.FFI_TYPE_VOID, this.cthis, margins)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:524
// index:0
// void getContentsMargins(int *, int *, int *, int *)
func (this *QWidget) GetContentsMargins(left unsafe.Pointer, top unsafe.Pointer, right unsafe.Pointer, bottom unsafe.Pointer) {
	// 0: (, int * left, int * top, int * right, int * bottom), (left, top, right, bottom)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget18getContentsMarginsEPiS0_S0_S0_", ffiqt.FFI_TYPE_VOID, this.cthis, left, top, right, bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:525
// index:0
// QMargins contentsMargins()
func (this *QWidget) ContentsMargins() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget15contentsMarginsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:527
// index:0
// QRect contentsRect()
func (this *QWidget) ContentsRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget12contentsRectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:530
// index:0
// QLayout * layout()
func (this *QWidget) Layout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget6layoutEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:531
// index:0
// void setLayout(class QLayout *)
func (this *QWidget) SetLayout(arg0 unsafe.Pointer) {
	// 0: (, QLayout * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9setLayoutEP7QLayout", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:532
// index:0
// void updateGeometry()
func (this *QWidget) UpdateGeometry() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14updateGeometryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:534
// index:0
// void setParent(class QWidget *)
func (this *QWidget) SetParent(parent unsafe.Pointer) {
	// 0: (, QWidget * parent), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9setParentEPS_", ffiqt.FFI_TYPE_VOID, this.cthis, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:537
// index:0
// void scroll(int, int)
func (this *QWidget) Scroll(dx int, dy int) {
	// 0: (, int dx, int dy), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget6scrollEii", ffiqt.FFI_TYPE_VOID, this.cthis, &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:538
// index:1
// void scroll(int, int, const class QRect &)
func (this *QWidget) Scroll_1(dx int, dy int, arg2 unsafe.Pointer) {
	// 1: (, int dx, int dy, const QRect & arg2), (&dx, &dy, arg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget6scrollEiiRK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, &dx, &dy, arg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:542
// index:0
// QWidget * focusWidget()
func (this *QWidget) FocusWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11focusWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:543
// index:0
// QWidget * nextInFocusChain()
func (this *QWidget) NextInFocusChain() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget16nextInFocusChainEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:544
// index:0
// QWidget * previousInFocusChain()
func (this *QWidget) PreviousInFocusChain() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget20previousInFocusChainEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:547
// index:0
// bool acceptDrops()
func (this *QWidget) AcceptDrops() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11acceptDropsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:548
// index:0
// void setAcceptDrops(_Bool)
func (this *QWidget) SetAcceptDrops(on bool) {
	// 0: (, bool on), (&on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14setAcceptDropsEb", ffiqt.FFI_TYPE_VOID, this.cthis, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:552
// index:0
// void addAction(class QAction *)
func (this *QWidget) AddAction(action unsafe.Pointer) {
	// 0: (, QAction * action), (action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9addActionEP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:560
// index:0
// void insertAction(class QAction *, class QAction *)
func (this *QWidget) InsertAction(before unsafe.Pointer, action unsafe.Pointer) {
	// 0: (, QAction * before, QAction * action), (before, action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12insertActionEP7QActionS1_", ffiqt.FFI_TYPE_VOID, this.cthis, before, action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:561
// index:0
// void removeAction(class QAction *)
func (this *QWidget) RemoveAction(action unsafe.Pointer) {
	// 0: (, QAction * action), (action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12removeActionEP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:562
// index:0
// QList<QAction *> actions()
func (this *QWidget) Actions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget7actionsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:565
// index:0
// QWidget * parentWidget()
func (this *QWidget) ParentWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget12parentWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:568
// index:0
// inline
// Qt::WindowFlags windowFlags()
func (this *QWidget) WindowFlags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11windowFlagsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:569
// index:0
// void setWindowFlag(Qt::WindowType, _Bool)
func (this *QWidget) SetWindowFlag(arg0 int, on bool) {
	// 0: (, Qt::WindowType arg0, bool on), (&arg0, &on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13setWindowFlagEN2Qt10WindowTypeEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:572
// index:0
// inline
// Qt::WindowType windowType()
func (this *QWidget) WindowType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget10windowTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:574
// index:0
// static
// QWidget * find(WId)
func (this *QWidget) Find(arg0 uint64) {
	// 0: (WId arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget4findEy", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QWidget_Find(arg0 uint64) {
	// 0: (WId arg0), (arg0)
	var nilthis *QWidget
	nilthis.Find(arg0)
}

// /usr/include/qt/QtWidgets/qwidget.h:575
// index:0
// inline
// QWidget * childAt(int, int)
func (this *QWidget) ChildAt(x int, y int) {
	// 0: (, int x, int y), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget7childAtEii", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:576
// index:1
// QWidget * childAt(const class QPoint &)
func (this *QWidget) ChildAt_1(p unsafe.Pointer) {
	// 1: (, const QPoint & p), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget7childAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:578
// index:0
// void setAttribute(Qt::WidgetAttribute, _Bool)
func (this *QWidget) SetAttribute(arg0 int, on bool) {
	// 0: (, Qt::WidgetAttribute arg0, bool on), (&arg0, &on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12setAttributeEN2Qt15WidgetAttributeEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:579
// index:0
// inline
// bool testAttribute(Qt::WidgetAttribute)
func (this *QWidget) TestAttribute(arg0 int) {
	// 0: (, Qt::WidgetAttribute arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13testAttributeEN2Qt15WidgetAttributeE", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:581
// index:0
// virtual
// QPaintEngine * paintEngine()
func (this *QWidget) PaintEngine() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11paintEngineEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:583
// index:0
// void ensurePolished()
func (this *QWidget) EnsurePolished() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14ensurePolishedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:585
// index:0
// bool isAncestorOf(const class QWidget *)
func (this *QWidget) IsAncestorOf(child unsafe.Pointer) {
	// 0: (, const QWidget * child), (child)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget12isAncestorOfEPKS_", ffiqt.FFI_TYPE_VOID, this.cthis, child)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:592
// index:0
// bool autoFillBackground()
func (this *QWidget) AutoFillBackground() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget18autoFillBackgroundEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:593
// index:0
// void setAutoFillBackground(_Bool)
func (this *QWidget) SetAutoFillBackground(enabled bool) {
	// 0: (, bool enabled), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget21setAutoFillBackgroundEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:595
// index:0
// QBackingStore * backingStore()
func (this *QWidget) BackingStore() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget12backingStoreEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:597
// index:0
// QWindow * windowHandle()
func (this *QWidget) WindowHandle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget12windowHandleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:604
// index:0
// void windowTitleChanged(const class QString &)
func (this *QWidget) WindowTitleChanged(title unsafe.Pointer) {
	// 0: (, const QString & title), (title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget18windowTitleChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, title)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:605
// index:0
// void windowIconChanged(const class QIcon &)
func (this *QWidget) WindowIconChanged(icon unsafe.Pointer) {
	// 0: (, const QIcon & icon), (icon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17windowIconChangedERK5QIcon", ffiqt.FFI_TYPE_VOID, this.cthis, icon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:606
// index:0
// void windowIconTextChanged(const class QString &)
func (this *QWidget) WindowIconTextChanged(iconText unsafe.Pointer) {
	// 0: (, const QString & iconText), (iconText)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget21windowIconTextChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, iconText)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:607
// index:0
// void customContextMenuRequested(const class QPoint &)
func (this *QWidget) CustomContextMenuRequested(pos unsafe.Pointer) {
	// 0: (, const QPoint & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget26customContextMenuRequestedERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:660
// index:0
// virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QWidget) InputMethodQuery(arg0 int) {
	// 0: (, Qt::InputMethodQuery arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:662
// index:0
// Qt::InputMethodHints inputMethodHints()
func (this *QWidget) InputMethodHints() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget16inputMethodHintsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
