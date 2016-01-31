package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
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
extern void* C_ZNK7QWindow15transientParentEv(void* qthis); // 4
  // proto:  void QWindow::show();
extern void C_ZN7QWindow4showEv(void* qthis); // 4
  // proto:  QObject * QWindow::focusObject();
extern void* C_ZNK7QWindow11focusObjectEv(void* qthis); // 4
  // proto:  void QWindow::setScreen(QScreen * screen);
extern void C_ZN7QWindow9setScreenEP7QScreen(void* qthis, void* arg0); // 4
  // proto:  QString QWindow::title();
extern void* C_ZNK7QWindow5titleEv(void* qthis); // 4
  // proto:  void QWindow::resize(const QSize & newSize);
extern void C_ZN7QWindow6resizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  void QWindow::resize(int w, int h);
extern void C_ZN7QWindow6resizeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QWindow::setIcon(const QIcon & icon);
extern void C_ZN7QWindow7setIconERK5QIcon(void* qthis, void* arg0); // 4
  // proto:  QPlatformWindow * QWindow::handle();
extern void C_ZNK7QWindow6handleEv(void* qthis); // 4
  // proto:  QSurfaceFormat QWindow::format();
extern void* C_ZNK7QWindow6formatEv(void* qthis); // 4
  // proto:  void QWindow::setGeometry(int posx, int posy, int w, int h);
extern void C_ZN7QWindow11setGeometryEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QWindow::setGeometry(const QRect & rect);
extern void C_ZN7QWindow11setGeometryERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QPoint QWindow::mapFromGlobal(const QPoint & pos);
extern void* C_ZNK7QWindow13mapFromGlobalERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QRect QWindow::frameGeometry();
extern void* C_ZNK7QWindow13frameGeometryEv(void* qthis); // 4
  // proto:  const QMetaObject * QWindow::metaObject();
extern void C_ZNK7QWindow10metaObjectEv(void* qthis); // 4
  // proto:  QSurfaceFormat QWindow::requestedFormat();
extern void* C_ZNK7QWindow15requestedFormatEv(void* qthis); // 4
  // proto:  QSize QWindow::minimumSize();
extern void* C_ZNK7QWindow11minimumSizeEv(void* qthis); // 4
  // proto:  Qt::WindowState QWindow::windowState();
extern void C_ZNK7QWindow11windowStateEv(void* qthis); // 4
  // proto:  void QWindow::setMaximumHeight(int h);
extern void C_ZN7QWindow16setMaximumHeightEi(void* qthis, int32_t arg0); // 4
  // proto:  int QWindow::x();
extern void C_ZNK7QWindow1xEv(void* qthis); // 2
  // proto:  QSize QWindow::maximumSize();
extern void* C_ZNK7QWindow11maximumSizeEv(void* qthis); // 4
  // proto:  int QWindow::minimumHeight();
extern int32_t C_ZNK7QWindow13minimumHeightEv(void* qthis); // 2
  // proto:  bool QWindow::close();
extern bool C_ZN7QWindow5closeEv(void* qthis); // 4
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
extern void* C_ZNK7QWindow6parentEv(void* qthis); // 4
  // proto:  QScreen * QWindow::screen();
extern void* C_ZNK7QWindow6screenEv(void* qthis); // 4
  // proto:  QAccessibleInterface * QWindow::accessibleRoot();
extern void* C_ZNK7QWindow14accessibleRootEv(void* qthis); // 4
  // proto:  void QWindow::setPosition(int posx, int posy);
extern void C_ZN7QWindow11setPositionEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QWindow::setPosition(const QPoint & pt);
extern void C_ZN7QWindow11setPositionERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  bool QWindow::isActive();
extern bool C_ZNK7QWindow8isActiveEv(void* qthis); // 4
  // proto:  QRect QWindow::geometry();
extern void* C_ZNK7QWindow8geometryEv(void* qthis); // 4
  // proto:  bool QWindow::isVisible();
extern bool C_ZNK7QWindow9isVisibleEv(void* qthis); // 4
  // proto:  QPoint QWindow::mapToGlobal(const QPoint & pos);
extern void* C_ZNK7QWindow11mapToGlobalERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QWindow::setMaximumWidth(int w);
extern void C_ZN7QWindow15setMaximumWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QWindow::setMinimumSize(const QSize & size);
extern void C_ZN7QWindow14setMinimumSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  int QWindow::height();
extern int32_t C_ZNK7QWindow6heightEv(void* qthis); // 2
  // proto:  qreal QWindow::devicePixelRatio();
extern double C_ZNK7QWindow16devicePixelRatioEv(void* qthis); // 4
  // proto:  void QWindow::setWidth(int arg);
extern void C_ZN7QWindow8setWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QWindow::size();
extern void* C_ZNK7QWindow4sizeEv(void* qthis); // 2
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
extern void* C_ZNK7QWindow13sizeIncrementEv(void* qthis); // 4
  // proto:  int QWindow::width();
extern int32_t C_ZNK7QWindow5widthEv(void* qthis); // 2
  // proto:  Qt::WindowType QWindow::type();
extern void C_ZNK7QWindow4typeEv(void* qthis); // 4
  // proto:  QSurface::SurfaceType QWindow::surfaceType();
extern void C_ZNK7QWindow11surfaceTypeEv(void* qthis); // 4
  // proto:  void QWindow::showMaximized();
extern void C_ZN7QWindow13showMaximizedEv(void* qthis); // 4
  // proto:  int QWindow::maximumWidth();
extern int32_t C_ZNK7QWindow12maximumWidthEv(void* qthis); // 2
  // proto:  QWindow::Visibility QWindow::visibility();
extern void C_ZNK7QWindow10visibilityEv(void* qthis); // 4
  // proto:  void QWindow::setMask(const QRegion & region);
extern void C_ZN7QWindow7setMaskERK7QRegion(void* qthis, void* arg0); // 4
  // proto:  void QWindow::setSizeIncrement(const QSize & size);
extern void C_ZN7QWindow16setSizeIncrementERK5QSize(void* qthis, void* arg0); // 4
  // proto:  void QWindow::QWindow(QWindow * parent);
extern void* C_ZN7QWindowC2EPS_(void* arg0); // 3
  // proto:  void QWindow::QWindow(QScreen * screen);
extern void* C_ZN7QWindowC2EP7QScreen(void* arg0); // 3
  // proto:  bool QWindow::setKeyboardGrabEnabled(bool grab);
extern bool C_ZN7QWindow22setKeyboardGrabEnabledEb(void* qthis, bool arg0); // 4
  // proto:  bool QWindow::isTopLevel();
extern bool C_ZNK7QWindow10isTopLevelEv(void* qthis); // 4
  // proto:  bool QWindow::isModal();
extern bool C_ZNK7QWindow7isModalEv(void* qthis); // 4
  // proto:  void QWindow::setMinimumWidth(int w);
extern void C_ZN7QWindow15setMinimumWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  QMargins QWindow::frameMargins();
extern void* C_ZNK7QWindow12frameMarginsEv(void* qthis); // 4
  // proto:  QCursor QWindow::cursor();
extern void* C_ZNK7QWindow6cursorEv(void* qthis); // 4
  // proto:  void QWindow::setVisible(bool visible);
extern void C_ZN7QWindow10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  QIcon QWindow::icon();
extern void* C_ZNK7QWindow4iconEv(void* qthis); // 4
  // proto:  Qt::ScreenOrientation QWindow::contentOrientation();
extern void C_ZNK7QWindow18contentOrientationEv(void* qthis); // 4
  // proto:  void QWindow::setTitle(const QString & );
extern void C_ZN7QWindow8setTitleERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QWindow::showMinimized();
extern void C_ZN7QWindow13showMinimizedEv(void* qthis); // 4
  // proto:  QSize QWindow::baseSize();
extern void* C_ZNK7QWindow8baseSizeEv(void* qthis); // 4
  // proto:  Qt::WindowModality QWindow::modality();
extern void C_ZNK7QWindow8modalityEv(void* qthis); // 4
  // proto:  void QWindow::showFullScreen();
extern void C_ZN7QWindow14showFullScreenEv(void* qthis); // 4
  // proto:  void QWindow::setMinimumHeight(int h);
extern void C_ZN7QWindow16setMinimumHeightEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QWindow::filePath();
extern void* C_ZNK7QWindow8filePathEv(void* qthis); // 4
  // proto:  void QWindow::setOpacity(qreal level);
extern void C_ZN7QWindow10setOpacityEd(void* qthis, double arg0); // 4
  // proto:  void QWindow::showNormal();
extern void C_ZN7QWindow10showNormalEv(void* qthis); // 4
  // proto:  void QWindow::destroy();
extern void C_ZN7QWindow7destroyEv(void* qthis); // 4
  // proto:  void QWindow::setBaseSize(const QSize & size);
extern void C_ZN7QWindow11setBaseSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  qreal QWindow::opacity();
extern double C_ZNK7QWindow7opacityEv(void* qthis); // 4
  // proto:  int QWindow::minimumWidth();
extern int32_t C_ZNK7QWindow12minimumWidthEv(void* qthis); // 2
  // proto:  int QWindow::maximumHeight();
extern int32_t C_ZNK7QWindow13maximumHeightEv(void* qthis); // 2
  // proto:  QPoint QWindow::framePosition();
extern void* C_ZNK7QWindow13framePositionEv(void* qthis); // 4
  // proto:  bool QWindow::setMouseGrabEnabled(bool grab);
extern bool C_ZN7QWindow19setMouseGrabEnabledEb(void* qthis, bool arg0); // 4
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
extern bool C_ZNK7QWindow9isExposedEv(void* qthis); // 4
  // proto:  QRegion QWindow::mask();
extern void* C_ZNK7QWindow4maskEv(void* qthis); // 4
  // proto:  Qt::WindowFlags QWindow::flags();
extern void C_ZNK7QWindow5flagsEv(void* qthis); // 4
  // proto:  void QWindow::setFramePosition(const QPoint & point);
extern void C_ZN7QWindow16setFramePositionERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  int QWindow::y();
extern void C_ZNK7QWindow1yEv(void* qthis); // 2
  // proto:  QPoint QWindow::position();
extern void* C_ZNK7QWindow8positionEv(void* qthis); // 2
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
func (this *QWindow) Freeqwindow(args ...interface{}) () {
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

  return
}

// transientParent()
func (this *QWindow) Transientparent(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow15transientParentEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWindow{}) // "QWindow *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "transientParent", args)
  }

  return
}

// show()
func (this *QWindow) Show(args ...interface{}) () {
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

  return
}

// focusObject()
func (this *QWindow) Focusobject(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow11focusObjectEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QObject{}) // "QObject *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "focusObject", args)
  }

  return
}

// setScreen(class QScreen *)
func (this *QWindow) Setscreen(args ...interface{}) () {
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

  return
}

// title()
func (this *QWindow) Title(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow5titleEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "title", args)
  }

  return
}

// resize(const class QSize &)
func (this *QWindow) Resize(args ...interface{}) () {
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

  return
}

// setIcon(const class QIcon &)
func (this *QWindow) Seticon(args ...interface{}) () {
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

  return
}

// handle()
func (this *QWindow) Handle(args ...interface{}) () {
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

  return
}

// format()
func (this *QWindow) Format(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow6formatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSurfaceFormat{}) // "QSurfaceFormat"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "format", args)
  }

  return
}

// setGeometry(int, int, int, int)
func (this *QWindow) Setgeometry(args ...interface{}) () {
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

  return
}

// mapFromGlobal(const class QPoint &)
func (this *QWindow) Mapfromglobal(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow13mapFromGlobalERK6QPoint(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "mapFromGlobal", args)
  }

  return
}

// frameGeometry()
func (this *QWindow) Framegeometry(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow13frameGeometryEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "frameGeometry", args)
  }

  return
}

// metaObject()
func (this *QWindow) Metaobject(args ...interface{}) () {
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

  return
}

// requestedFormat()
func (this *QWindow) Requestedformat(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow15requestedFormatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSurfaceFormat{}) // "QSurfaceFormat"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "requestedFormat", args)
  }

  return
}

// minimumSize()
func (this *QWindow) Minimumsize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow11minimumSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "minimumSize", args)
  }

  return
}

// windowState()
func (this *QWindow) Windowstate(args ...interface{}) () {
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

  return
}

// setMaximumHeight(int)
func (this *QWindow) Setmaximumheight(args ...interface{}) () {
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

  return
}

// x()
func (this *QWindow) X(args ...interface{}) () {
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

  return
}

// maximumSize()
func (this *QWindow) Maximumsize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow11maximumSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "maximumSize", args)
  }

  return
}

// minimumHeight()
func (this *QWindow) Minimumheight(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow13minimumHeightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "minimumHeight", args)
  }

  return
}

// close()
func (this *QWindow) Close(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN7QWindow5closeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "close", args)
  }

  return
}

// requestActivate()
func (this *QWindow) Requestactivate(args ...interface{}) () {
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

  return
}

// setMaximumSize(const class QSize &)
func (this *QWindow) Setmaximumsize(args ...interface{}) () {
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

  return
}

// hide()
func (this *QWindow) Hide(args ...interface{}) () {
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

  return
}

// setParent(class QWindow *)
func (this *QWindow) Setparent(args ...interface{}) () {
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

  return
}

// requestUpdate()
func (this *QWindow) Requestupdate(args ...interface{}) () {
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

  return
}

// setTransientParent(class QWindow *)
func (this *QWindow) Settransientparent(args ...interface{}) () {
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

  return
}

// parent()
func (this *QWindow) Parent(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow6parentEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWindow{}) // "QWindow *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "parent", args)
  }

  return
}

// screen()
func (this *QWindow) Screen(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow6screenEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QScreen{}) // "QScreen *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "screen", args)
  }

  return
}

// accessibleRoot()
func (this *QWindow) Accessibleroot(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow14accessibleRootEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "accessibleRoot", args)
  }

  return
}

// setPosition(int, int)
func (this *QWindow) Setposition(args ...interface{}) () {
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

  return
}

// isActive()
func (this *QWindow) Isactive(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow8isActiveEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "isActive", args)
  }

  return
}

// geometry()
func (this *QWindow) Geometry(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow8geometryEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "geometry", args)
  }

  return
}

// isVisible()
func (this *QWindow) Isvisible(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow9isVisibleEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "isVisible", args)
  }

  return
}

// mapToGlobal(const class QPoint &)
func (this *QWindow) Maptoglobal(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow11mapToGlobalERK6QPoint(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "mapToGlobal", args)
  }

  return
}

// setMaximumWidth(int)
func (this *QWindow) Setmaximumwidth(args ...interface{}) () {
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

  return
}

// setMinimumSize(const class QSize &)
func (this *QWindow) Setminimumsize(args ...interface{}) () {
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

  return
}

// height()
func (this *QWindow) Height(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow6heightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "height", args)
  }

  return
}

// devicePixelRatio()
func (this *QWindow) Devicepixelratio(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow16devicePixelRatioEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "devicePixelRatio", args)
  }

  return
}

// setWidth(int)
func (this *QWindow) Setwidth(args ...interface{}) () {
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

  return
}

// size()
func (this *QWindow) Size(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow4sizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "size", args)
  }

  return
}

// unsetCursor()
func (this *QWindow) Unsetcursor(args ...interface{}) () {
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

  return
}

// winId()
func (this *QWindow) Winid(args ...interface{}) () {
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

  return
}

// raise()
func (this *QWindow) Raise(args ...interface{}) () {
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

  return
}

// setFormat(const class QSurfaceFormat &)
func (this *QWindow) Setformat(args ...interface{}) () {
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

  return
}

// create()
func (this *QWindow) Create(args ...interface{}) () {
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

  return
}

// sizeIncrement()
func (this *QWindow) Sizeincrement(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow13sizeIncrementEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "sizeIncrement", args)
  }

  return
}

// width()
func (this *QWindow) Width(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow5widthEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "width", args)
  }

  return
}

// type()
func (this *QWindow) Type_(args ...interface{}) () {
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

  return
}

// surfaceType()
func (this *QWindow) Surfacetype(args ...interface{}) () {
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

  return
}

// showMaximized()
func (this *QWindow) Showmaximized(args ...interface{}) () {
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

  return
}

// maximumWidth()
func (this *QWindow) Maximumwidth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow12maximumWidthEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "maximumWidth", args)
  }

  return
}

// visibility()
func (this *QWindow) Visibility(args ...interface{}) () {
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

  return
}

// setMask(const class QRegion &)
func (this *QWindow) Setmask(args ...interface{}) () {
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

  return
}

// setSizeIncrement(const class QSize &)
func (this *QWindow) Setsizeincrement(args ...interface{}) () {
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

  return
}

// QWindow(class QWindow *)
func NewQWindow(args ...interface{}) *QWindow {
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
    qthis = C.C_ZN7QWindowC2EPS_(arg0)
    return &QWindow{qclsinst:qthis}
  case 1:
    // invoke: _ZN7QWindowC1EP7QScreen
    // invoke: void QWindow(class QScreen *)
    var arg0 = args[0].(QScreen).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QWindowC2EP7QScreen(arg0)
    return &QWindow{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QWindow", "QWindow", args)
  }

  return nil // QWindow{qclsinst:qthis}
}

// setKeyboardGrabEnabled(_Bool)
func (this *QWindow) Setkeyboardgrabenabled(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN7QWindow22setKeyboardGrabEnabledEb(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "setKeyboardGrabEnabled", args)
  }

  return
}

// isTopLevel()
func (this *QWindow) Istoplevel(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow10isTopLevelEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "isTopLevel", args)
  }

  return
}

// isModal()
func (this *QWindow) Ismodal(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow7isModalEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "isModal", args)
  }

  return
}

// setMinimumWidth(int)
func (this *QWindow) Setminimumwidth(args ...interface{}) () {
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

  return
}

// frameMargins()
func (this *QWindow) Framemargins(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow12frameMarginsEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMargins{}) // "QMargins"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "frameMargins", args)
  }

  return
}

// cursor()
func (this *QWindow) Cursor(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow6cursorEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QCursor{}) // "QCursor"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "cursor", args)
  }

  return
}

// setVisible(_Bool)
func (this *QWindow) Setvisible(args ...interface{}) () {
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

  return
}

// icon()
func (this *QWindow) Icon(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow4iconEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QIcon{}) // "QIcon"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "icon", args)
  }

  return
}

// contentOrientation()
func (this *QWindow) Contentorientation(args ...interface{}) () {
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

  return
}

// setTitle(const class QString &)
func (this *QWindow) Settitle(args ...interface{}) () {
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

  return
}

// showMinimized()
func (this *QWindow) Showminimized(args ...interface{}) () {
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

  return
}

// baseSize()
func (this *QWindow) Basesize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow8baseSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "baseSize", args)
  }

  return
}

// modality()
func (this *QWindow) Modality(args ...interface{}) () {
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

  return
}

// showFullScreen()
func (this *QWindow) Showfullscreen(args ...interface{}) () {
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

  return
}

// setMinimumHeight(int)
func (this *QWindow) Setminimumheight(args ...interface{}) () {
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

  return
}

// filePath()
func (this *QWindow) Filepath(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow8filePathEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "filePath", args)
  }

  return
}

// setOpacity(qreal)
func (this *QWindow) Setopacity(args ...interface{}) () {
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

  return
}

// showNormal()
func (this *QWindow) Shownormal(args ...interface{}) () {
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

  return
}

// destroy()
func (this *QWindow) Destroy(args ...interface{}) () {
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

  return
}

// setBaseSize(const class QSize &)
func (this *QWindow) Setbasesize(args ...interface{}) () {
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

  return
}

// opacity()
func (this *QWindow) Opacity(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow7opacityEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "opacity", args)
  }

  return
}

// minimumWidth()
func (this *QWindow) Minimumwidth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow12minimumWidthEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "minimumWidth", args)
  }

  return
}

// maximumHeight()
func (this *QWindow) Maximumheight(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow13maximumHeightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "maximumHeight", args)
  }

  return
}

// framePosition()
func (this *QWindow) Frameposition(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow13framePositionEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "framePosition", args)
  }

  return
}

// setMouseGrabEnabled(_Bool)
func (this *QWindow) Setmousegrabenabled(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN7QWindow19setMouseGrabEnabledEb(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "setMouseGrabEnabled", args)
  }

  return
}

// alert(int)
func (this *QWindow) Alert(args ...interface{}) () {
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

  return
}

// setX(int)
func (this *QWindow) Setx(args ...interface{}) () {
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

  return
}

// setY(int)
func (this *QWindow) Sety(args ...interface{}) () {
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

  return
}

// setCursor(const class QCursor &)
func (this *QWindow) Setcursor(args ...interface{}) () {
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

  return
}

// setHeight(int)
func (this *QWindow) Setheight(args ...interface{}) () {
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

  return
}

// lower()
func (this *QWindow) Lower(args ...interface{}) () {
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

  return
}

// setFilePath(const class QString &)
func (this *QWindow) Setfilepath(args ...interface{}) () {
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

  return
}

// isExposed()
func (this *QWindow) Isexposed(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow9isExposedEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "isExposed", args)
  }

  return
}

// mask()
func (this *QWindow) Mask(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow4maskEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegion{}) // "QRegion"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "mask", args)
  }

  return
}

// flags()
func (this *QWindow) Flags(args ...interface{}) () {
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

  return
}

// setFramePosition(const class QPoint &)
func (this *QWindow) Setframeposition(args ...interface{}) () {
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

  return
}

// y()
func (this *QWindow) Y(args ...interface{}) () {
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

  return
}

// position()
func (this *QWindow) Position(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWindow8positionEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWindow", "position", args)
  }

  return
}

// <= body block end

