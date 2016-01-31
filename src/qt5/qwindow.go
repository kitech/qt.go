package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtGui/qwindow.h
// dst-file: /src/gui/qwindow.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QWindow::~QWindow();
extern void C_ZN7QWindowD2Ev(void* qthis); // 4
  // proto:  QWindow * QWindow::transientParent();
extern void C_ZNK7QWindow15transientParentEv(void* qthis); // 4
  // proto:  void QWindow::show();
extern void C_ZN7QWindow4showEv(void* qthis); // 4
  // proto:  QObject * QWindow::focusObject();
extern void C_ZNK7QWindow11focusObjectEv(void* qthis); // 4
  // proto:  void QWindow::setScreen(QScreen * screen);
extern void C_ZN7QWindow9setScreenEP7QScreen(void* qthis, void* arg0); // 4
  // proto:  QString QWindow::title();
extern void C_ZNK7QWindow5titleEv(void* qthis); // 4
  // proto:  void QWindow::resize(const QSize & newSize);
extern void C_ZN7QWindow6resizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  void QWindow::resize(int w, int h);
extern void C_ZN7QWindow6resizeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QWindow::setIcon(const QIcon & icon);
extern void C_ZN7QWindow7setIconERK5QIcon(void* qthis, void* arg0); // 4
  // proto:  QPlatformWindow * QWindow::handle();
extern void C_ZNK7QWindow6handleEv(void* qthis); // 4
  // proto:  QSurfaceFormat QWindow::format();
extern void C_ZNK7QWindow6formatEv(void* qthis); // 4
  // proto:  void QWindow::setGeometry(int posx, int posy, int w, int h);
extern void C_ZN7QWindow11setGeometryEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QWindow::setGeometry(const QRect & rect);
extern void C_ZN7QWindow11setGeometryERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QPoint QWindow::mapFromGlobal(const QPoint & pos);
extern void C_ZNK7QWindow13mapFromGlobalERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QRect QWindow::frameGeometry();
extern void C_ZNK7QWindow13frameGeometryEv(void* qthis); // 4
  // proto:  const QMetaObject * QWindow::metaObject();
extern void C_ZNK7QWindow10metaObjectEv(void* qthis); // 4
  // proto:  QSurfaceFormat QWindow::requestedFormat();
extern void C_ZNK7QWindow15requestedFormatEv(void* qthis); // 4
  // proto:  QSize QWindow::minimumSize();
extern void C_ZNK7QWindow11minimumSizeEv(void* qthis); // 4
  // proto:  Qt::WindowState QWindow::windowState();
extern void C_ZNK7QWindow11windowStateEv(void* qthis); // 4
  // proto:  void QWindow::setMaximumHeight(int h);
extern void C_ZN7QWindow16setMaximumHeightEi(void* qthis, int32_t arg0); // 4
  // proto:  int QWindow::x();
extern void C_ZNK7QWindow1xEv(void* qthis); // 2
  // proto:  QSize QWindow::maximumSize();
extern void C_ZNK7QWindow11maximumSizeEv(void* qthis); // 4
  // proto:  int QWindow::minimumHeight();
extern void C_ZNK7QWindow13minimumHeightEv(void* qthis); // 2
  // proto:  bool QWindow::close();
extern void C_ZN7QWindow5closeEv(void* qthis); // 4
  // proto:  void QWindow::requestActivate();
extern void C_ZN7QWindow15requestActivateEv(void* qthis); // 4
  // proto:  void QWindow::setMaximumSize(const QSize & size);
extern void C_ZN7QWindow14setMaximumSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  void QWindow::hide();
extern void C_ZN7QWindow4hideEv(void* qthis); // 4
  // proto:  void QWindow::setParent(QWindow * parent);
extern void C_ZN7QWindow9setParentEPS_(void* qthis, void* arg0); // 4
  // proto:  void QWindow::requestUpdate();
extern void C_ZN7QWindow13requestUpdateEv(void* qthis); // 4
  // proto:  void QWindow::setTransientParent(QWindow * parent);
extern void C_ZN7QWindow18setTransientParentEPS_(void* qthis, void* arg0); // 4
  // proto:  QWindow * QWindow::parent();
extern void C_ZNK7QWindow6parentEv(void* qthis); // 4
  // proto:  QScreen * QWindow::screen();
extern void C_ZNK7QWindow6screenEv(void* qthis); // 4
  // proto:  QAccessibleInterface * QWindow::accessibleRoot();
extern void C_ZNK7QWindow14accessibleRootEv(void* qthis); // 4
  // proto:  void QWindow::setPosition(int posx, int posy);
extern void C_ZN7QWindow11setPositionEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QWindow::setPosition(const QPoint & pt);
extern void C_ZN7QWindow11setPositionERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  bool QWindow::isActive();
extern void C_ZNK7QWindow8isActiveEv(void* qthis); // 4
  // proto:  QRect QWindow::geometry();
extern void C_ZNK7QWindow8geometryEv(void* qthis); // 4
  // proto:  bool QWindow::isVisible();
extern void C_ZNK7QWindow9isVisibleEv(void* qthis); // 4
  // proto:  QPoint QWindow::mapToGlobal(const QPoint & pos);
extern void C_ZNK7QWindow11mapToGlobalERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QWindow::setMaximumWidth(int w);
extern void C_ZN7QWindow15setMaximumWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QWindow::setMinimumSize(const QSize & size);
extern void C_ZN7QWindow14setMinimumSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  int QWindow::height();
extern void C_ZNK7QWindow6heightEv(void* qthis); // 2
  // proto:  qreal QWindow::devicePixelRatio();
extern void C_ZNK7QWindow16devicePixelRatioEv(void* qthis); // 4
  // proto:  void QWindow::setWidth(int arg);
extern void C_ZN7QWindow8setWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QWindow::size();
extern void C_ZNK7QWindow4sizeEv(void* qthis); // 2
  // proto:  void QWindow::unsetCursor();
extern void C_ZN7QWindow11unsetCursorEv(void* qthis); // 4
  // proto:  WId QWindow::winId();
extern void C_ZNK7QWindow5winIdEv(void* qthis); // 4
  // proto:  void QWindow::raise();
extern void C_ZN7QWindow5raiseEv(void* qthis); // 4
  // proto:  void QWindow::setFormat(const QSurfaceFormat & format);
extern void C_ZN7QWindow9setFormatERK14QSurfaceFormat(void* qthis, void* arg0); // 4
  // proto:  void QWindow::create();
extern void C_ZN7QWindow6createEv(void* qthis); // 4
  // proto:  QSize QWindow::sizeIncrement();
extern void C_ZNK7QWindow13sizeIncrementEv(void* qthis); // 4
  // proto:  int QWindow::width();
extern void C_ZNK7QWindow5widthEv(void* qthis); // 2
  // proto:  Qt::WindowType QWindow::type();
extern void C_ZNK7QWindow4typeEv(void* qthis); // 4
  // proto:  QSurface::SurfaceType QWindow::surfaceType();
extern void C_ZNK7QWindow11surfaceTypeEv(void* qthis); // 4
  // proto:  void QWindow::showMaximized();
extern void C_ZN7QWindow13showMaximizedEv(void* qthis); // 4
  // proto:  int QWindow::maximumWidth();
extern void C_ZNK7QWindow12maximumWidthEv(void* qthis); // 2
  // proto:  QWindow::Visibility QWindow::visibility();
extern void C_ZNK7QWindow10visibilityEv(void* qthis); // 4
  // proto:  void QWindow::setMask(const QRegion & region);
extern void C_ZN7QWindow7setMaskERK7QRegion(void* qthis, void* arg0); // 4
  // proto:  void QWindow::setSizeIncrement(const QSize & size);
extern void C_ZN7QWindow16setSizeIncrementERK5QSize(void* qthis, void* arg0); // 4
  // proto:  void QWindow::QWindow(QWindow * parent);
extern void C_ZN7QWindowC2EPS_(void* qthis, void* arg0); // 3
  // proto:  void QWindow::QWindow(QScreen * screen);
extern void C_ZN7QWindowC2EP7QScreen(void* qthis, void* arg0); // 3
  // proto:  bool QWindow::setKeyboardGrabEnabled(bool grab);
extern void C_ZN7QWindow22setKeyboardGrabEnabledEb(void* qthis, bool arg0); // 4
  // proto:  bool QWindow::isTopLevel();
extern void C_ZNK7QWindow10isTopLevelEv(void* qthis); // 4
  // proto:  bool QWindow::isModal();
extern void C_ZNK7QWindow7isModalEv(void* qthis); // 4
  // proto:  void QWindow::setMinimumWidth(int w);
extern void C_ZN7QWindow15setMinimumWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  QMargins QWindow::frameMargins();
extern void C_ZNK7QWindow12frameMarginsEv(void* qthis); // 4
  // proto:  QCursor QWindow::cursor();
extern void C_ZNK7QWindow6cursorEv(void* qthis); // 4
  // proto:  void QWindow::setVisible(bool visible);
extern void C_ZN7QWindow10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  QIcon QWindow::icon();
extern void C_ZNK7QWindow4iconEv(void* qthis); // 4
  // proto:  Qt::ScreenOrientation QWindow::contentOrientation();
extern void C_ZNK7QWindow18contentOrientationEv(void* qthis); // 4
  // proto:  void QWindow::setTitle(const QString & );
extern void C_ZN7QWindow8setTitleERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QWindow::showMinimized();
extern void C_ZN7QWindow13showMinimizedEv(void* qthis); // 4
  // proto:  QSize QWindow::baseSize();
extern void C_ZNK7QWindow8baseSizeEv(void* qthis); // 4
  // proto:  Qt::WindowModality QWindow::modality();
extern void C_ZNK7QWindow8modalityEv(void* qthis); // 4
  // proto:  void QWindow::showFullScreen();
extern void C_ZN7QWindow14showFullScreenEv(void* qthis); // 4
  // proto:  void QWindow::setMinimumHeight(int h);
extern void C_ZN7QWindow16setMinimumHeightEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QWindow::filePath();
extern void C_ZNK7QWindow8filePathEv(void* qthis); // 4
  // proto:  void QWindow::setOpacity(qreal level);
extern void C_ZN7QWindow10setOpacityEd(void* qthis, double arg0); // 4
  // proto:  void QWindow::showNormal();
extern void C_ZN7QWindow10showNormalEv(void* qthis); // 4
  // proto:  void QWindow::destroy();
extern void C_ZN7QWindow7destroyEv(void* qthis); // 4
  // proto:  void QWindow::setBaseSize(const QSize & size);
extern void C_ZN7QWindow11setBaseSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  qreal QWindow::opacity();
extern void C_ZNK7QWindow7opacityEv(void* qthis); // 4
  // proto:  int QWindow::minimumWidth();
extern void C_ZNK7QWindow12minimumWidthEv(void* qthis); // 2
  // proto:  int QWindow::maximumHeight();
extern void C_ZNK7QWindow13maximumHeightEv(void* qthis); // 2
  // proto:  QPoint QWindow::framePosition();
extern void C_ZNK7QWindow13framePositionEv(void* qthis); // 4
  // proto:  bool QWindow::setMouseGrabEnabled(bool grab);
extern void C_ZN7QWindow19setMouseGrabEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QWindow::alert(int msec);
extern void C_ZN7QWindow5alertEi(void* qthis, int32_t arg0); // 4
  // proto:  void QWindow::setX(int arg);
extern void C_ZN7QWindow4setXEi(void* qthis, int32_t arg0); // 4
  // proto:  void QWindow::setY(int arg);
extern void C_ZN7QWindow4setYEi(void* qthis, int32_t arg0); // 4
  // proto:  void QWindow::setCursor(const QCursor & );
extern void C_ZN7QWindow9setCursorERK7QCursor(void* qthis, void* arg0); // 4
  // proto:  void QWindow::setHeight(int arg);
extern void C_ZN7QWindow9setHeightEi(void* qthis, int32_t arg0); // 4
  // proto:  void QWindow::lower();
extern void C_ZN7QWindow5lowerEv(void* qthis); // 4
  // proto:  void QWindow::setFilePath(const QString & filePath);
extern void C_ZN7QWindow11setFilePathERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QWindow::isExposed();
extern void C_ZNK7QWindow9isExposedEv(void* qthis); // 4
  // proto:  QRegion QWindow::mask();
extern void C_ZNK7QWindow4maskEv(void* qthis); // 4
  // proto:  Qt::WindowFlags QWindow::flags();
extern void C_ZNK7QWindow5flagsEv(void* qthis); // 4
  // proto:  void QWindow::setFramePosition(const QPoint & point);
extern void C_ZN7QWindow16setFramePositionERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  int QWindow::y();
extern void C_ZNK7QWindow1yEv(void* qthis); // 2
  // proto:  QPoint QWindow::position();
extern void C_ZNK7QWindow8positionEv(void* qthis); // 2
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QWindow)=1
type QWindow struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _modalityChanged QWindow_modalityChanged_signal;
//  _activeChanged QWindow_activeChanged_signal;
//  _heightChanged QWindow_heightChanged_signal;
//  _contentOrientationChanged QWindow_contentOrientationChanged_signal;
//  _minimumWidthChanged QWindow_minimumWidthChanged_signal;
//  _opacityChanged QWindow_opacityChanged_signal;
//  _visibleChanged QWindow_visibleChanged_signal;
//  _screenChanged QWindow_screenChanged_signal;
//  _maximumHeightChanged QWindow_maximumHeightChanged_signal;
//  _yChanged QWindow_yChanged_signal;
//  _widthChanged QWindow_widthChanged_signal;
//  _windowStateChanged QWindow_windowStateChanged_signal;
//  _windowTitleChanged QWindow_windowTitleChanged_signal;
//  _visibilityChanged QWindow_visibilityChanged_signal;
//  _minimumHeightChanged QWindow_minimumHeightChanged_signal;
//  _xChanged QWindow_xChanged_signal;
//  _focusObjectChanged QWindow_focusObjectChanged_signal;
//  _maximumWidthChanged QWindow_maximumWidthChanged_signal;
}

// ~QWindow()
func (this *QWindow) FreeQWindow(args ...interface{}) () {
  // ~QWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindowD0Ev
    // invoke: void ~QWindow()
    C.C_ZN7QWindowD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "~QWindow", args)
  }

}

// transientParent()
func (this *QWindow) transientParent(args ...interface{}) () {
  // transientParent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow15transientParentEv
    // invoke: QWindow * transientParent()
    C.C_ZNK7QWindow15transientParentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "transientParent", args)
  }

}

// show()
func (this *QWindow) show(args ...interface{}) () {
  // show()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow4showEv
    // invoke: void show()
    C.C_ZN7QWindow4showEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "show", args)
  }

}

// focusObject()
func (this *QWindow) focusObject(args ...interface{}) () {
  // focusObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow11focusObjectEv
    // invoke: QObject * focusObject()
    C.C_ZNK7QWindow11focusObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "focusObject", args)
  }

}

// setScreen(class QScreen *)
func (this *QWindow) setScreen(args ...interface{}) () {
  // setScreen(class QScreen *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QScreen{}) // "QScreen *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow9setScreenEP7QScreen
    // invoke: void setScreen(class QScreen *)
    var arg0 = args[0].(QScreen).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow9setScreenEP7QScreen(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setScreen", args)
  }

}

// title()
func (this *QWindow) title(args ...interface{}) () {
  // title()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow5titleEv
    // invoke: QString title()
    C.C_ZNK7QWindow5titleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "title", args)
  }

}

// resize(const class QSize &)
func (this *QWindow) resize(args ...interface{}) () {
  // resize(const class QSize &)
  // resize(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow6resizeERK5QSize
    // invoke: void resize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow6resizeERK5QSize(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN7QWindow6resizeEii
    // invoke: void resize(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN7QWindow6resizeEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QWindow", "resize", args)
  }

}

// setIcon(const class QIcon &)
func (this *QWindow) setIcon(args ...interface{}) () {
  // setIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow7setIconERK5QIcon
    // invoke: void setIcon(const class QIcon &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow7setIconERK5QIcon(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setIcon", args)
  }

}

// handle()
func (this *QWindow) handle(args ...interface{}) () {
  // handle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow6handleEv
    // invoke: QPlatformWindow * handle()
    C.C_ZNK7QWindow6handleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "handle", args)
  }

}

// format()
func (this *QWindow) format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow6formatEv
    // invoke: QSurfaceFormat format()
    C.C_ZNK7QWindow6formatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "format", args)
  }

}

// setGeometry(int, int, int, int)
func (this *QWindow) setGeometry(args ...interface{}) () {
  // setGeometry(int, int, int, int)
  // setGeometry(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow11setGeometryEiiii
    // invoke: void setGeometry(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN7QWindow11setGeometryEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN7QWindow11setGeometryERK5QRect
    // invoke: void setGeometry(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow11setGeometryERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setGeometry", args)
  }

}

// mapFromGlobal(const class QPoint &)
func (this *QWindow) mapFromGlobal(args ...interface{}) () {
  // mapFromGlobal(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow13mapFromGlobalERK6QPoint
    // invoke: QPoint mapFromGlobal(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK7QWindow13mapFromGlobalERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "mapFromGlobal", args)
  }

}

// frameGeometry()
func (this *QWindow) frameGeometry(args ...interface{}) () {
  // frameGeometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow13frameGeometryEv
    // invoke: QRect frameGeometry()
    C.C_ZNK7QWindow13frameGeometryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "frameGeometry", args)
  }

}

// metaObject()
func (this *QWindow) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK7QWindow10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "metaObject", args)
  }

}

// requestedFormat()
func (this *QWindow) requestedFormat(args ...interface{}) () {
  // requestedFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow15requestedFormatEv
    // invoke: QSurfaceFormat requestedFormat()
    C.C_ZNK7QWindow15requestedFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "requestedFormat", args)
  }

}

// minimumSize()
func (this *QWindow) minimumSize(args ...interface{}) () {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow11minimumSizeEv
    // invoke: QSize minimumSize()
    C.C_ZNK7QWindow11minimumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "minimumSize", args)
  }

}

// windowState()
func (this *QWindow) windowState(args ...interface{}) () {
  // windowState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow11windowStateEv
    // invoke: Qt::WindowState windowState()
    C.C_ZNK7QWindow11windowStateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "windowState", args)
  }

}

// setMaximumHeight(int)
func (this *QWindow) setMaximumHeight(args ...interface{}) () {
  // setMaximumHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow16setMaximumHeightEi
    // invoke: void setMaximumHeight(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow16setMaximumHeightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setMaximumHeight", args)
  }

}

// x()
func (this *QWindow) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow1xEv
    // invoke: int x()
    C.C_ZNK7QWindow1xEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "x", args)
  }

}

// maximumSize()
func (this *QWindow) maximumSize(args ...interface{}) () {
  // maximumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow11maximumSizeEv
    // invoke: QSize maximumSize()
    C.C_ZNK7QWindow11maximumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "maximumSize", args)
  }

}

// minimumHeight()
func (this *QWindow) minimumHeight(args ...interface{}) () {
  // minimumHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow13minimumHeightEv
    // invoke: int minimumHeight()
    C.C_ZNK7QWindow13minimumHeightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "minimumHeight", args)
  }

}

// close()
func (this *QWindow) close(args ...interface{}) () {
  // close()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow5closeEv
    // invoke: bool close()
    C.C_ZN7QWindow5closeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "close", args)
  }

}

// requestActivate()
func (this *QWindow) requestActivate(args ...interface{}) () {
  // requestActivate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow15requestActivateEv
    // invoke: void requestActivate()
    C.C_ZN7QWindow15requestActivateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "requestActivate", args)
  }

}

// setMaximumSize(const class QSize &)
func (this *QWindow) setMaximumSize(args ...interface{}) () {
  // setMaximumSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow14setMaximumSizeERK5QSize
    // invoke: void setMaximumSize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow14setMaximumSizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setMaximumSize", args)
  }

}

// hide()
func (this *QWindow) hide(args ...interface{}) () {
  // hide()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow4hideEv
    // invoke: void hide()
    C.C_ZN7QWindow4hideEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "hide", args)
  }

}

// setParent(class QWindow *)
func (this *QWindow) setParent(args ...interface{}) () {
  // setParent(class QWindow *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWindow{}) // "QWindow *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow9setParentEPS_
    // invoke: void setParent(class QWindow *)
    var arg0 = args[0].(QWindow).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow9setParentEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setParent", args)
  }

}

// requestUpdate()
func (this *QWindow) requestUpdate(args ...interface{}) () {
  // requestUpdate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow13requestUpdateEv
    // invoke: void requestUpdate()
    C.C_ZN7QWindow13requestUpdateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "requestUpdate", args)
  }

}

// setTransientParent(class QWindow *)
func (this *QWindow) setTransientParent(args ...interface{}) () {
  // setTransientParent(class QWindow *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWindow{}) // "QWindow *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow18setTransientParentEPS_
    // invoke: void setTransientParent(class QWindow *)
    var arg0 = args[0].(QWindow).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow18setTransientParentEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setTransientParent", args)
  }

}

// parent()
func (this *QWindow) parent(args ...interface{}) () {
  // parent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow6parentEv
    // invoke: QWindow * parent()
    C.C_ZNK7QWindow6parentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "parent", args)
  }

}

// screen()
func (this *QWindow) screen(args ...interface{}) () {
  // screen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow6screenEv
    // invoke: QScreen * screen()
    C.C_ZNK7QWindow6screenEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "screen", args)
  }

}

// accessibleRoot()
func (this *QWindow) accessibleRoot(args ...interface{}) () {
  // accessibleRoot()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow14accessibleRootEv
    // invoke: QAccessibleInterface * accessibleRoot()
    C.C_ZNK7QWindow14accessibleRootEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "accessibleRoot", args)
  }

}

// setPosition(int, int)
func (this *QWindow) setPosition(args ...interface{}) () {
  // setPosition(int, int)
  // setPosition(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow11setPositionEii
    // invoke: void setPosition(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN7QWindow11setPositionEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN7QWindow11setPositionERK6QPoint
    // invoke: void setPosition(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow11setPositionERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setPosition", args)
  }

}

// isActive()
func (this *QWindow) isActive(args ...interface{}) () {
  // isActive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow8isActiveEv
    // invoke: bool isActive()
    C.C_ZNK7QWindow8isActiveEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "isActive", args)
  }

}

// geometry()
func (this *QWindow) geometry(args ...interface{}) () {
  // geometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow8geometryEv
    // invoke: QRect geometry()
    C.C_ZNK7QWindow8geometryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "geometry", args)
  }

}

// isVisible()
func (this *QWindow) isVisible(args ...interface{}) () {
  // isVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow9isVisibleEv
    // invoke: bool isVisible()
    C.C_ZNK7QWindow9isVisibleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "isVisible", args)
  }

}

// mapToGlobal(const class QPoint &)
func (this *QWindow) mapToGlobal(args ...interface{}) () {
  // mapToGlobal(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow11mapToGlobalERK6QPoint
    // invoke: QPoint mapToGlobal(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK7QWindow11mapToGlobalERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "mapToGlobal", args)
  }

}

// setMaximumWidth(int)
func (this *QWindow) setMaximumWidth(args ...interface{}) () {
  // setMaximumWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow15setMaximumWidthEi
    // invoke: void setMaximumWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow15setMaximumWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setMaximumWidth", args)
  }

}

// setMinimumSize(const class QSize &)
func (this *QWindow) setMinimumSize(args ...interface{}) () {
  // setMinimumSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow14setMinimumSizeERK5QSize
    // invoke: void setMinimumSize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow14setMinimumSizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setMinimumSize", args)
  }

}

// height()
func (this *QWindow) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow6heightEv
    // invoke: int height()
    C.C_ZNK7QWindow6heightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "height", args)
  }

}

// devicePixelRatio()
func (this *QWindow) devicePixelRatio(args ...interface{}) () {
  // devicePixelRatio()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow16devicePixelRatioEv
    // invoke: qreal devicePixelRatio()
    C.C_ZNK7QWindow16devicePixelRatioEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "devicePixelRatio", args)
  }

}

// setWidth(int)
func (this *QWindow) setWidth(args ...interface{}) () {
  // setWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow8setWidthEi
    // invoke: void setWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow8setWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setWidth", args)
  }

}

// size()
func (this *QWindow) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow4sizeEv
    // invoke: QSize size()
    C.C_ZNK7QWindow4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "size", args)
  }

}

// unsetCursor()
func (this *QWindow) unsetCursor(args ...interface{}) () {
  // unsetCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow11unsetCursorEv
    // invoke: void unsetCursor()
    C.C_ZN7QWindow11unsetCursorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "unsetCursor", args)
  }

}

// winId()
func (this *QWindow) winId(args ...interface{}) () {
  // winId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow5winIdEv
    // invoke: WId winId()
    C.C_ZNK7QWindow5winIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "winId", args)
  }

}

// raise()
func (this *QWindow) raise(args ...interface{}) () {
  // raise()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow5raiseEv
    // invoke: void raise()
    C.C_ZN7QWindow5raiseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "raise", args)
  }

}

// setFormat(const class QSurfaceFormat &)
func (this *QWindow) setFormat(args ...interface{}) () {
  // setFormat(const class QSurfaceFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSurfaceFormat{}) // "const QSurfaceFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow9setFormatERK14QSurfaceFormat
    // invoke: void setFormat(const class QSurfaceFormat &)
    var arg0 = args[0].(QSurfaceFormat).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow9setFormatERK14QSurfaceFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setFormat", args)
  }

}

// create()
func (this *QWindow) create(args ...interface{}) () {
  // create()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow6createEv
    // invoke: void create()
    C.C_ZN7QWindow6createEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "create", args)
  }

}

// sizeIncrement()
func (this *QWindow) sizeIncrement(args ...interface{}) () {
  // sizeIncrement()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow13sizeIncrementEv
    // invoke: QSize sizeIncrement()
    C.C_ZNK7QWindow13sizeIncrementEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "sizeIncrement", args)
  }

}

// width()
func (this *QWindow) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow5widthEv
    // invoke: int width()
    C.C_ZNK7QWindow5widthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "width", args)
  }

}

// type()
func (this *QWindow) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow4typeEv
    // invoke: Qt::WindowType type()
    C.C_ZNK7QWindow4typeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "type", args)
  }

}

// surfaceType()
func (this *QWindow) surfaceType(args ...interface{}) () {
  // surfaceType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow11surfaceTypeEv
    // invoke: QSurface::SurfaceType surfaceType()
    C.C_ZNK7QWindow11surfaceTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "surfaceType", args)
  }

}

// showMaximized()
func (this *QWindow) showMaximized(args ...interface{}) () {
  // showMaximized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow13showMaximizedEv
    // invoke: void showMaximized()
    C.C_ZN7QWindow13showMaximizedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "showMaximized", args)
  }

}

// maximumWidth()
func (this *QWindow) maximumWidth(args ...interface{}) () {
  // maximumWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow12maximumWidthEv
    // invoke: int maximumWidth()
    C.C_ZNK7QWindow12maximumWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "maximumWidth", args)
  }

}

// visibility()
func (this *QWindow) visibility(args ...interface{}) () {
  // visibility()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow10visibilityEv
    // invoke: QWindow::Visibility visibility()
    C.C_ZNK7QWindow10visibilityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "visibility", args)
  }

}

// setMask(const class QRegion &)
func (this *QWindow) setMask(args ...interface{}) () {
  // setMask(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow7setMaskERK7QRegion
    // invoke: void setMask(const class QRegion &)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow7setMaskERK7QRegion(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setMask", args)
  }

}

// setSizeIncrement(const class QSize &)
func (this *QWindow) setSizeIncrement(args ...interface{}) () {
  // setSizeIncrement(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow16setSizeIncrementERK5QSize
    // invoke: void setSizeIncrement(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow16setSizeIncrementERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setSizeIncrement", args)
  }

}

// QWindow(class QWindow *)
func NewQWindow(args ...interface{}) QWindow {
  // QWindow(class QWindow *)
  // QWindow(class QScreen *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWindow{}) // "QWindow *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QScreen{}) // "QScreen *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindowC1EPS_
    // invoke: void QWindow(class QWindow *)
    var arg0 = args[0].(QWindow).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN7QWindowC2EPS_(qthis, arg0)
  case 1:
    // invoke: _ZN7QWindowC1EP7QScreen
    // invoke: void QWindow(class QScreen *)
    var arg0 = args[0].(QScreen).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN7QWindowC2EP7QScreen(qthis, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "QWindow", args)
  }

  return QWindow{}
}

// setKeyboardGrabEnabled(_Bool)
func (this *QWindow) setKeyboardGrabEnabled(args ...interface{}) () {
  // setKeyboardGrabEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow22setKeyboardGrabEnabledEb
    // invoke: bool setKeyboardGrabEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow22setKeyboardGrabEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setKeyboardGrabEnabled", args)
  }

}

// isTopLevel()
func (this *QWindow) isTopLevel(args ...interface{}) () {
  // isTopLevel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow10isTopLevelEv
    // invoke: bool isTopLevel()
    C.C_ZNK7QWindow10isTopLevelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "isTopLevel", args)
  }

}

// isModal()
func (this *QWindow) isModal(args ...interface{}) () {
  // isModal()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow7isModalEv
    // invoke: bool isModal()
    C.C_ZNK7QWindow7isModalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "isModal", args)
  }

}

// setMinimumWidth(int)
func (this *QWindow) setMinimumWidth(args ...interface{}) () {
  // setMinimumWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow15setMinimumWidthEi
    // invoke: void setMinimumWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow15setMinimumWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setMinimumWidth", args)
  }

}

// frameMargins()
func (this *QWindow) frameMargins(args ...interface{}) () {
  // frameMargins()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow12frameMarginsEv
    // invoke: QMargins frameMargins()
    C.C_ZNK7QWindow12frameMarginsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "frameMargins", args)
  }

}

// cursor()
func (this *QWindow) cursor(args ...interface{}) () {
  // cursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow6cursorEv
    // invoke: QCursor cursor()
    C.C_ZNK7QWindow6cursorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "cursor", args)
  }

}

// setVisible(_Bool)
func (this *QWindow) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow10setVisibleEb
    // invoke: void setVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow10setVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setVisible", args)
  }

}

// icon()
func (this *QWindow) icon(args ...interface{}) () {
  // icon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow4iconEv
    // invoke: QIcon icon()
    C.C_ZNK7QWindow4iconEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "icon", args)
  }

}

// contentOrientation()
func (this *QWindow) contentOrientation(args ...interface{}) () {
  // contentOrientation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow18contentOrientationEv
    // invoke: Qt::ScreenOrientation contentOrientation()
    C.C_ZNK7QWindow18contentOrientationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "contentOrientation", args)
  }

}

// setTitle(const class QString &)
func (this *QWindow) setTitle(args ...interface{}) () {
  // setTitle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow8setTitleERK7QString
    // invoke: void setTitle(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow8setTitleERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setTitle", args)
  }

}

// showMinimized()
func (this *QWindow) showMinimized(args ...interface{}) () {
  // showMinimized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow13showMinimizedEv
    // invoke: void showMinimized()
    C.C_ZN7QWindow13showMinimizedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "showMinimized", args)
  }

}

// baseSize()
func (this *QWindow) baseSize(args ...interface{}) () {
  // baseSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow8baseSizeEv
    // invoke: QSize baseSize()
    C.C_ZNK7QWindow8baseSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "baseSize", args)
  }

}

// modality()
func (this *QWindow) modality(args ...interface{}) () {
  // modality()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow8modalityEv
    // invoke: Qt::WindowModality modality()
    C.C_ZNK7QWindow8modalityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "modality", args)
  }

}

// showFullScreen()
func (this *QWindow) showFullScreen(args ...interface{}) () {
  // showFullScreen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow14showFullScreenEv
    // invoke: void showFullScreen()
    C.C_ZN7QWindow14showFullScreenEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "showFullScreen", args)
  }

}

// setMinimumHeight(int)
func (this *QWindow) setMinimumHeight(args ...interface{}) () {
  // setMinimumHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow16setMinimumHeightEi
    // invoke: void setMinimumHeight(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow16setMinimumHeightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setMinimumHeight", args)
  }

}

// filePath()
func (this *QWindow) filePath(args ...interface{}) () {
  // filePath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow8filePathEv
    // invoke: QString filePath()
    C.C_ZNK7QWindow8filePathEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "filePath", args)
  }

}

// setOpacity(qreal)
func (this *QWindow) setOpacity(args ...interface{}) () {
  // setOpacity(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow10setOpacityEd
    // invoke: void setOpacity(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow10setOpacityEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setOpacity", args)
  }

}

// showNormal()
func (this *QWindow) showNormal(args ...interface{}) () {
  // showNormal()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow10showNormalEv
    // invoke: void showNormal()
    C.C_ZN7QWindow10showNormalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "showNormal", args)
  }

}

// destroy()
func (this *QWindow) destroy(args ...interface{}) () {
  // destroy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow7destroyEv
    // invoke: void destroy()
    C.C_ZN7QWindow7destroyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "destroy", args)
  }

}

// setBaseSize(const class QSize &)
func (this *QWindow) setBaseSize(args ...interface{}) () {
  // setBaseSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow11setBaseSizeERK5QSize
    // invoke: void setBaseSize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow11setBaseSizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setBaseSize", args)
  }

}

// opacity()
func (this *QWindow) opacity(args ...interface{}) () {
  // opacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow7opacityEv
    // invoke: qreal opacity()
    C.C_ZNK7QWindow7opacityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "opacity", args)
  }

}

// minimumWidth()
func (this *QWindow) minimumWidth(args ...interface{}) () {
  // minimumWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow12minimumWidthEv
    // invoke: int minimumWidth()
    C.C_ZNK7QWindow12minimumWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "minimumWidth", args)
  }

}

// maximumHeight()
func (this *QWindow) maximumHeight(args ...interface{}) () {
  // maximumHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow13maximumHeightEv
    // invoke: int maximumHeight()
    C.C_ZNK7QWindow13maximumHeightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "maximumHeight", args)
  }

}

// framePosition()
func (this *QWindow) framePosition(args ...interface{}) () {
  // framePosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow13framePositionEv
    // invoke: QPoint framePosition()
    C.C_ZNK7QWindow13framePositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "framePosition", args)
  }

}

// setMouseGrabEnabled(_Bool)
func (this *QWindow) setMouseGrabEnabled(args ...interface{}) () {
  // setMouseGrabEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow19setMouseGrabEnabledEb
    // invoke: bool setMouseGrabEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow19setMouseGrabEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setMouseGrabEnabled", args)
  }

}

// alert(int)
func (this *QWindow) alert(args ...interface{}) () {
  // alert(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow5alertEi
    // invoke: void alert(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow5alertEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "alert", args)
  }

}

// setX(int)
func (this *QWindow) setX(args ...interface{}) () {
  // setX(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow4setXEi
    // invoke: void setX(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow4setXEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setX", args)
  }

}

// setY(int)
func (this *QWindow) setY(args ...interface{}) () {
  // setY(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow4setYEi
    // invoke: void setY(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow4setYEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setY", args)
  }

}

// setCursor(const class QCursor &)
func (this *QWindow) setCursor(args ...interface{}) () {
  // setCursor(const class QCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCursor{}) // "const QCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow9setCursorERK7QCursor
    // invoke: void setCursor(const class QCursor &)
    var arg0 = args[0].(QCursor).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow9setCursorERK7QCursor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setCursor", args)
  }

}

// setHeight(int)
func (this *QWindow) setHeight(args ...interface{}) () {
  // setHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow9setHeightEi
    // invoke: void setHeight(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow9setHeightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setHeight", args)
  }

}

// lower()
func (this *QWindow) lower(args ...interface{}) () {
  // lower()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow5lowerEv
    // invoke: void lower()
    C.C_ZN7QWindow5lowerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "lower", args)
  }

}

// setFilePath(const class QString &)
func (this *QWindow) setFilePath(args ...interface{}) () {
  // setFilePath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow11setFilePathERK7QString
    // invoke: void setFilePath(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow11setFilePathERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setFilePath", args)
  }

}

// isExposed()
func (this *QWindow) isExposed(args ...interface{}) () {
  // isExposed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow9isExposedEv
    // invoke: bool isExposed()
    C.C_ZNK7QWindow9isExposedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "isExposed", args)
  }

}

// mask()
func (this *QWindow) mask(args ...interface{}) () {
  // mask()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow4maskEv
    // invoke: QRegion mask()
    C.C_ZNK7QWindow4maskEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "mask", args)
  }

}

// flags()
func (this *QWindow) flags(args ...interface{}) () {
  // flags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow5flagsEv
    // invoke: Qt::WindowFlags flags()
    C.C_ZNK7QWindow5flagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "flags", args)
  }

}

// setFramePosition(const class QPoint &)
func (this *QWindow) setFramePosition(args ...interface{}) () {
  // setFramePosition(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow16setFramePositionERK6QPoint
    // invoke: void setFramePosition(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWindow16setFramePositionERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setFramePosition", args)
  }

}

// y()
func (this *QWindow) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow1yEv
    // invoke: int y()
    C.C_ZNK7QWindow1yEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "y", args)
  }

}

// position()
func (this *QWindow) position(args ...interface{}) () {
  // position()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow8positionEv
    // invoke: QPoint position()
    C.C_ZNK7QWindow8positionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "position", args)
  }

}

// <= body block end

