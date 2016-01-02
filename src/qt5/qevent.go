package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtGui/qevent.h
// dst-file: /src/gui/qevent.go
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
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QString QWhatsThisClickedEvent::href();
extern void demth_ZNK22QWhatsThisClickedEvent4hrefEv(void* qthis);
  // proto:  void QWhatsThisClickedEvent::~QWhatsThisClickedEvent();
extern void _ZN22QWhatsThisClickedEventD0Ev(void* qthis);
  // proto:  void QWhatsThisClickedEvent::QWhatsThisClickedEvent(const QString & href);
extern void* dector_ZN22QWhatsThisClickedEventC1ERK7QString(void* arg0);
extern void _ZN22QWhatsThisClickedEventC1ERK7QString(void* qthis, void* arg0);
  // proto:  void QExposeEvent::QExposeEvent(const QRegion & rgn);
extern void* dector_ZN12QExposeEventC1ERK7QRegion(void* arg0);
extern void _ZN12QExposeEventC1ERK7QRegion(void* qthis, void* arg0);
  // proto:  const QRegion & QExposeEvent::region();
extern void demth_ZNK12QExposeEvent6regionEv(void* qthis);
  // proto:  void QExposeEvent::~QExposeEvent();
extern void _ZN12QExposeEventD0Ev(void* qthis);
  // proto:  const QString & QInputMethodEvent::preeditString();
extern void demth_ZNK17QInputMethodEvent13preeditStringEv(void* qthis);
  // proto:  void QInputMethodEvent::QInputMethodEvent();
extern void* dector_ZN17QInputMethodEventC1Ev();
extern void _ZN17QInputMethodEventC1Ev(void* qthis);
  // proto:  int QInputMethodEvent::replacementStart();
extern void demth_ZNK17QInputMethodEvent16replacementStartEv(void* qthis);
  // proto:  void QInputMethodEvent::QInputMethodEvent(const QInputMethodEvent & other);
extern void* dector_ZN17QInputMethodEventC1ERKS_(void* arg0);
extern void _ZN17QInputMethodEventC1ERKS_(void* qthis, void* arg0);
  // proto:  const QString & QInputMethodEvent::commitString();
extern void demth_ZNK17QInputMethodEvent12commitStringEv(void* qthis);
  // proto:  void QInputMethodEvent::setCommitString(const QString & commitString, int replaceFrom, int replaceLength);
extern void _ZN17QInputMethodEvent15setCommitStringERK7QStringii(void* qthis, void* arg0, int arg1, int arg2);
  // proto:  int QInputMethodEvent::replacementLength();
extern void demth_ZNK17QInputMethodEvent17replacementLengthEv(void* qthis);
  // proto:  const QPoint & QHelpEvent::globalPos();
extern void demth_ZNK10QHelpEvent9globalPosEv(void* qthis);
  // proto:  int QHelpEvent::globalX();
extern void demth_ZNK10QHelpEvent7globalXEv(void* qthis);
  // proto:  const QPoint & QHelpEvent::pos();
extern void demth_ZNK10QHelpEvent3posEv(void* qthis);
  // proto:  int QHelpEvent::y();
extern void demth_ZNK10QHelpEvent1yEv(void* qthis);
  // proto:  int QHelpEvent::globalY();
extern void demth_ZNK10QHelpEvent7globalYEv(void* qthis);
  // proto:  int QHelpEvent::x();
extern void demth_ZNK10QHelpEvent1xEv(void* qthis);
  // proto:  void QHelpEvent::~QHelpEvent();
extern void _ZN10QHelpEventD0Ev(void* qthis);
  // proto:  void QActionEvent::QActionEvent(int type, QAction * action, QAction * before);
extern void* dector_ZN12QActionEventC1EiP7QActionS1_(int arg0, void* arg1, void* arg2);
extern void _ZN12QActionEventC1EiP7QActionS1_(void* qthis, int arg0, void* arg1, void* arg2);
  // proto:  QAction * QActionEvent::before();
extern void demth_ZNK12QActionEvent6beforeEv(void* qthis);
  // proto:  QAction * QActionEvent::action();
extern void demth_ZNK12QActionEvent6actionEv(void* qthis);
  // proto:  void QActionEvent::~QActionEvent();
extern void _ZN12QActionEventD0Ev(void* qthis);
  // proto:  QPoint QMouseEvent::globalPos();
extern void demth_ZNK11QMouseEvent9globalPosEv(void* qthis);
  // proto:  int QMouseEvent::y();
extern void demth_ZNK11QMouseEvent1yEv(void* qthis);
  // proto:  const QPointF & QMouseEvent::screenPos();
extern void demth_ZNK11QMouseEvent9screenPosEv(void* qthis);
  // proto:  int QMouseEvent::x();
extern void demth_ZNK11QMouseEvent1xEv(void* qthis);
  // proto:  const QPointF & QMouseEvent::localPos();
extern void _ZNK11QMouseEvent8localPosEv(void* qthis);
  // proto:  int QMouseEvent::globalX();
extern void demth_ZNK11QMouseEvent7globalXEv(void* qthis);
  // proto:  const QPointF & QMouseEvent::windowPos();
extern void _ZNK11QMouseEvent9windowPosEv(void* qthis);
  // proto:  void QMouseEvent::~QMouseEvent();
extern void _ZN11QMouseEventD0Ev(void* qthis);
  // proto:  int QMouseEvent::globalY();
extern void demth_ZNK11QMouseEvent7globalYEv(void* qthis);
  // proto:  QPoint QMouseEvent::pos();
extern void demth_ZNK11QMouseEvent3posEv(void* qthis);
  // proto:  void QFileOpenEvent::QFileOpenEvent(const QString & file);
extern void* dector_ZN14QFileOpenEventC1ERK7QString(void* arg0);
extern void _ZN14QFileOpenEventC1ERK7QString(void* qthis, void* arg0);
  // proto:  void QFileOpenEvent::~QFileOpenEvent();
extern void _ZN14QFileOpenEventD0Ev(void* qthis);
  // proto:  void QFileOpenEvent::QFileOpenEvent(const QUrl & url);
extern void* dector_ZN14QFileOpenEventC1ERK4QUrl(void* arg0);
extern void _ZN14QFileOpenEventC1ERK4QUrl(void* qthis, void* arg0);
  // proto:  QString QFileOpenEvent::file();
extern void demth_ZNK14QFileOpenEvent4fileEv(void* qthis);
  // proto:  QUrl QFileOpenEvent::url();
extern void _ZNK14QFileOpenEvent3urlEv(void* qthis);
  // proto:  void QToolBarChangeEvent::QToolBarChangeEvent(bool t);
extern void* dector_ZN19QToolBarChangeEventC1Eb(bool arg0);
extern void _ZN19QToolBarChangeEventC1Eb(void* qthis, bool arg0);
  // proto:  void QToolBarChangeEvent::~QToolBarChangeEvent();
extern void _ZN19QToolBarChangeEventD0Ev(void* qthis);
  // proto:  bool QToolBarChangeEvent::toggle();
extern void demth_ZNK19QToolBarChangeEvent6toggleEv(void* qthis);
  // proto:  int QTabletEvent::x();
extern void demth_ZNK12QTabletEvent1xEv(void* qthis);
  // proto:  int QTabletEvent::xTilt();
extern void demth_ZNK12QTabletEvent5xTiltEv(void* qthis);
  // proto:  qint64 QTabletEvent::uniqueId();
extern void demth_ZNK12QTabletEvent8uniqueIdEv(void* qthis);
  // proto:  const QPointF & QTabletEvent::globalPosF();
extern void demth_ZNK12QTabletEvent10globalPosFEv(void* qthis);
  // proto:  int QTabletEvent::z();
extern void demth_ZNK12QTabletEvent1zEv(void* qthis);
  // proto:  int QTabletEvent::y();
extern void demth_ZNK12QTabletEvent1yEv(void* qthis);
  // proto:  QPoint QTabletEvent::pos();
extern void demth_ZNK12QTabletEvent3posEv(void* qthis);
  // proto:  qreal QTabletEvent::rotation();
extern void demth_ZNK12QTabletEvent8rotationEv(void* qthis);
  // proto:  QPoint QTabletEvent::globalPos();
extern void demth_ZNK12QTabletEvent9globalPosEv(void* qthis);
  // proto:  void QTabletEvent::~QTabletEvent();
extern void _ZN12QTabletEventD0Ev(void* qthis);
  // proto:  qreal QTabletEvent::tangentialPressure();
extern void demth_ZNK12QTabletEvent18tangentialPressureEv(void* qthis);
  // proto:  qreal QTabletEvent::hiResGlobalX();
extern void demth_ZNK12QTabletEvent12hiResGlobalXEv(void* qthis);
  // proto:  int QTabletEvent::globalY();
extern void demth_ZNK12QTabletEvent7globalYEv(void* qthis);
  // proto:  qreal QTabletEvent::hiResGlobalY();
extern void demth_ZNK12QTabletEvent12hiResGlobalYEv(void* qthis);
  // proto:  int QTabletEvent::globalX();
extern void demth_ZNK12QTabletEvent7globalXEv(void* qthis);
  // proto:  const QPointF & QTabletEvent::posF();
extern void demth_ZNK12QTabletEvent4posFEv(void* qthis);
  // proto:  qreal QTabletEvent::pressure();
extern void demth_ZNK12QTabletEvent8pressureEv(void* qthis);
  // proto:  int QTabletEvent::yTilt();
extern void demth_ZNK12QTabletEvent5yTiltEv(void* qthis);
  // proto:  void QTouchEvent::setDevice(QTouchDevice * adevice);
extern void demth_ZN11QTouchEvent9setDeviceEP12QTouchDevice(void* qthis, void* arg0);
  // proto:  QWindow * QTouchEvent::window();
extern void demth_ZNK11QTouchEvent6windowEv(void* qthis);
  // proto:  QTouchDevice * QTouchEvent::device();
extern void demth_ZNK11QTouchEvent6deviceEv(void* qthis);
  // proto:  QObject * QTouchEvent::target();
extern void demth_ZNK11QTouchEvent6targetEv(void* qthis);
  // proto:  void QTouchEvent::~QTouchEvent();
extern void _ZN11QTouchEventD0Ev(void* qthis);
  // proto:  void QTouchEvent::setWindow(QWindow * awindow);
extern void demth_ZN11QTouchEvent9setWindowEP7QWindow(void* qthis, void* arg0);
  // proto:  void QTouchEvent::setTarget(QObject * atarget);
extern void demth_ZN11QTouchEvent9setTargetEP7QObject(void* qthis, void* arg0);
  // proto:  QScreen * QScreenOrientationChangeEvent::screen();
extern void _ZNK29QScreenOrientationChangeEvent6screenEv(void* qthis);
  // proto:  void QScreenOrientationChangeEvent::~QScreenOrientationChangeEvent();
extern void _ZN29QScreenOrientationChangeEventD0Ev(void* qthis);
  // proto:  void QIconDragEvent::~QIconDragEvent();
extern void _ZN14QIconDragEventD0Ev(void* qthis);
  // proto:  void QIconDragEvent::QIconDragEvent();
extern void* dector_ZN14QIconDragEventC1Ev();
extern void _ZN14QIconDragEventC1Ev(void* qthis);
  // proto:  void QCloseEvent::~QCloseEvent();
extern void _ZN11QCloseEventD0Ev(void* qthis);
  // proto:  void QCloseEvent::QCloseEvent();
extern void* dector_ZN11QCloseEventC1Ev();
extern void _ZN11QCloseEventC1Ev(void* qthis);
  // proto:  void QDragEnterEvent::~QDragEnterEvent();
extern void _ZN15QDragEnterEventD0Ev(void* qthis);
  // proto:  int QWheelEvent::x();
extern void demth_ZNK11QWheelEvent1xEv(void* qthis);
  // proto:  QPoint QWheelEvent::angleDelta();
extern void demth_ZNK11QWheelEvent10angleDeltaEv(void* qthis);
  // proto:  QPoint QWheelEvent::pos();
extern void demth_ZNK11QWheelEvent3posEv(void* qthis);
  // proto:  int QWheelEvent::globalY();
extern void demth_ZNK11QWheelEvent7globalYEv(void* qthis);
  // proto:  const QPointF & QWheelEvent::posF();
extern void demth_ZNK11QWheelEvent4posFEv(void* qthis);
  // proto:  int QWheelEvent::globalX();
extern void demth_ZNK11QWheelEvent7globalXEv(void* qthis);
  // proto:  int QWheelEvent::y();
extern void demth_ZNK11QWheelEvent1yEv(void* qthis);
  // proto:  void QWheelEvent::~QWheelEvent();
extern void _ZN11QWheelEventD0Ev(void* qthis);
  // proto:  QPoint QWheelEvent::pixelDelta();
extern void demth_ZNK11QWheelEvent10pixelDeltaEv(void* qthis);
  // proto:  int QWheelEvent::delta();
extern void demth_ZNK11QWheelEvent5deltaEv(void* qthis);
  // proto:  QPoint QWheelEvent::globalPos();
extern void demth_ZNK11QWheelEvent9globalPosEv(void* qthis);
  // proto:  const QPointF & QWheelEvent::globalPosF();
extern void demth_ZNK11QWheelEvent10globalPosFEv(void* qthis);
  // proto:  QPointF QScrollEvent::contentPos();
extern void _ZNK12QScrollEvent10contentPosEv(void* qthis);
  // proto:  QPointF QScrollEvent::overshootDistance();
extern void _ZNK12QScrollEvent17overshootDistanceEv(void* qthis);
  // proto:  void QScrollEvent::~QScrollEvent();
extern void _ZN12QScrollEventD0Ev(void* qthis);
  // proto:  void QHoverEvent::~QHoverEvent();
extern void _ZN11QHoverEventD0Ev(void* qthis);
  // proto:  const QPointF & QHoverEvent::posF();
extern void demth_ZNK11QHoverEvent4posFEv(void* qthis);
  // proto:  QPoint QHoverEvent::oldPos();
extern void demth_ZNK11QHoverEvent6oldPosEv(void* qthis);
  // proto:  const QPointF & QHoverEvent::oldPosF();
extern void demth_ZNK11QHoverEvent7oldPosFEv(void* qthis);
  // proto:  QPoint QHoverEvent::pos();
extern void demth_ZNK11QHoverEvent3posEv(void* qthis);
  // proto:  void QDragMoveEvent::accept(const QRect & r);
extern void demth_ZN14QDragMoveEvent6acceptERK5QRect(void* qthis, void* arg0);
  // proto:  QRect QDragMoveEvent::answerRect();
extern void demth_ZNK14QDragMoveEvent10answerRectEv(void* qthis);
  // proto:  void QDragMoveEvent::ignore(const QRect & r);
extern void demth_ZN14QDragMoveEvent6ignoreERK5QRect(void* qthis, void* arg0);
  // proto:  void QDragMoveEvent::ignore();
extern void demth_ZN14QDragMoveEvent6ignoreEv(void* qthis);
  // proto:  void QDragMoveEvent::~QDragMoveEvent();
extern void _ZN14QDragMoveEventD0Ev(void* qthis);
  // proto:  void QDragMoveEvent::accept();
extern void demth_ZN14QDragMoveEvent6acceptEv(void* qthis);
  // proto:  void QShowEvent::~QShowEvent();
extern void _ZN10QShowEventD0Ev(void* qthis);
  // proto:  void QShowEvent::QShowEvent();
extern void* dector_ZN10QShowEventC1Ev();
extern void _ZN10QShowEventC1Ev(void* qthis);
  // proto:  void QPlatformSurfaceEvent::~QPlatformSurfaceEvent();
extern void _ZN21QPlatformSurfaceEventD0Ev(void* qthis);
  // proto:  void QPaintEvent::~QPaintEvent();
extern void _ZN11QPaintEventD0Ev(void* qthis);
  // proto:  const QRect & QPaintEvent::rect();
extern void demth_ZNK11QPaintEvent4rectEv(void* qthis);
  // proto:  void QPaintEvent::QPaintEvent(const QRect & paintRect);
extern void* dector_ZN11QPaintEventC1ERK5QRect(void* arg0);
extern void _ZN11QPaintEventC1ERK5QRect(void* qthis, void* arg0);
  // proto:  const QRegion & QPaintEvent::region();
extern void demth_ZNK11QPaintEvent6regionEv(void* qthis);
  // proto:  void QPaintEvent::QPaintEvent(const QRegion & paintRegion);
extern void* dector_ZN11QPaintEventC1ERK7QRegion(void* arg0);
extern void _ZN11QPaintEventC1ERK7QRegion(void* qthis, void* arg0);
  // proto:  bool QFocusEvent::lostFocus();
extern void demth_ZNK11QFocusEvent9lostFocusEv(void* qthis);
  // proto:  bool QFocusEvent::gotFocus();
extern void demth_ZNK11QFocusEvent8gotFocusEv(void* qthis);
  // proto:  void QFocusEvent::~QFocusEvent();
extern void _ZN11QFocusEventD0Ev(void* qthis);
  // proto:  const QPointF & QNativeGestureEvent::localPos();
extern void _ZNK19QNativeGestureEvent8localPosEv(void* qthis);
  // proto:  const QPointF & QNativeGestureEvent::screenPos();
extern void _ZNK19QNativeGestureEvent9screenPosEv(void* qthis);
  // proto:  const QPoint QNativeGestureEvent::pos();
extern void demth_ZNK19QNativeGestureEvent3posEv(void* qthis);
  // proto:  const QPoint QNativeGestureEvent::globalPos();
extern void demth_ZNK19QNativeGestureEvent9globalPosEv(void* qthis);
  // proto:  qreal QNativeGestureEvent::value();
extern void _ZNK19QNativeGestureEvent5valueEv(void* qthis);
  // proto:  const QPointF & QNativeGestureEvent::windowPos();
extern void _ZNK19QNativeGestureEvent9windowPosEv(void* qthis);
  // proto:  const QSize & QResizeEvent::oldSize();
extern void demth_ZNK12QResizeEvent7oldSizeEv(void* qthis);
  // proto:  const QSize & QResizeEvent::size();
extern void demth_ZNK12QResizeEvent4sizeEv(void* qthis);
  // proto:  void QResizeEvent::~QResizeEvent();
extern void _ZN12QResizeEventD0Ev(void* qthis);
  // proto:  void QResizeEvent::QResizeEvent(const QSize & size, const QSize & oldSize);
extern void* dector_ZN12QResizeEventC1ERK5QSizeS2_(void* arg0, void* arg1);
extern void _ZN12QResizeEventC1ERK5QSizeS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QStatusTipEvent::~QStatusTipEvent();
extern void _ZN15QStatusTipEventD0Ev(void* qthis);
  // proto:  QString QStatusTipEvent::tip();
extern void demth_ZNK15QStatusTipEvent3tipEv(void* qthis);
  // proto:  void QStatusTipEvent::QStatusTipEvent(const QString & tip);
extern void* dector_ZN15QStatusTipEventC1ERK7QString(void* arg0);
extern void _ZN15QStatusTipEventC1ERK7QString(void* qthis, void* arg0);
  // proto:  int QEnterEvent::y();
extern void demth_ZNK11QEnterEvent1yEv(void* qthis);
  // proto:  QPoint QEnterEvent::pos();
extern void demth_ZNK11QEnterEvent3posEv(void* qthis);
  // proto:  void QEnterEvent::~QEnterEvent();
extern void _ZN11QEnterEventD0Ev(void* qthis);
  // proto:  const QPointF & QEnterEvent::screenPos();
extern void _ZNK11QEnterEvent9screenPosEv(void* qthis);
  // proto:  const QPointF & QEnterEvent::localPos();
extern void _ZNK11QEnterEvent8localPosEv(void* qthis);
  // proto:  const QPointF & QEnterEvent::windowPos();
extern void _ZNK11QEnterEvent9windowPosEv(void* qthis);
  // proto:  int QEnterEvent::globalX();
extern void demth_ZNK11QEnterEvent7globalXEv(void* qthis);
  // proto:  int QEnterEvent::x();
extern void demth_ZNK11QEnterEvent1xEv(void* qthis);
  // proto:  QPoint QEnterEvent::globalPos();
extern void demth_ZNK11QEnterEvent9globalPosEv(void* qthis);
  // proto:  int QEnterEvent::globalY();
extern void demth_ZNK11QEnterEvent7globalYEv(void* qthis);
  // proto:  void QEnterEvent::QEnterEvent(const QPointF & localPos, const QPointF & windowPos, const QPointF & screenPos);
extern void* dector_ZN11QEnterEventC1ERK7QPointFS2_S2_(void* arg0, void* arg1, void* arg2);
extern void _ZN11QEnterEventC1ERK7QPointFS2_S2_(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QMoveEvent::~QMoveEvent();
extern void _ZN10QMoveEventD0Ev(void* qthis);
  // proto:  const QPoint & QMoveEvent::oldPos();
extern void demth_ZNK10QMoveEvent6oldPosEv(void* qthis);
  // proto:  void QMoveEvent::QMoveEvent(const QPoint & pos, const QPoint & oldPos);
extern void* dector_ZN10QMoveEventC1ERK6QPointS2_(void* arg0, void* arg1);
extern void _ZN10QMoveEventC1ERK6QPointS2_(void* qthis, void* arg0, void* arg1);
  // proto:  const QPoint & QMoveEvent::pos();
extern void demth_ZNK10QMoveEvent3posEv(void* qthis);
  // proto:  void QHideEvent::QHideEvent();
extern void* dector_ZN10QHideEventC1Ev();
extern void _ZN10QHideEventC1Ev(void* qthis);
  // proto:  void QHideEvent::~QHideEvent();
extern void _ZN10QHideEventD0Ev(void* qthis);
  // proto:  void QDragLeaveEvent::~QDragLeaveEvent();
extern void _ZN15QDragLeaveEventD0Ev(void* qthis);
  // proto:  void QDragLeaveEvent::QDragLeaveEvent();
extern void* dector_ZN15QDragLeaveEventC1Ev();
extern void _ZN15QDragLeaveEventC1Ev(void* qthis);
  // proto:  void QDropEvent::~QDropEvent();
extern void _ZN10QDropEventD0Ev(void* qthis);
  // proto:  QPoint QDropEvent::pos();
extern void demth_ZNK10QDropEvent3posEv(void* qthis);
  // proto:  QObject * QDropEvent::source();
extern void _ZNK10QDropEvent6sourceEv(void* qthis);
  // proto:  const QMimeData * QDropEvent::mimeData();
extern void demth_ZNK10QDropEvent8mimeDataEv(void* qthis);
  // proto:  void QDropEvent::acceptProposedAction();
extern void demth_ZN10QDropEvent20acceptProposedActionEv(void* qthis);
  // proto:  const QPointF & QDropEvent::posF();
extern void demth_ZNK10QDropEvent4posFEv(void* qthis);
  // proto:  void QInputEvent::setTimestamp(ulong atimestamp);
extern void demth_ZN11QInputEvent12setTimestampEm(void* qthis, unsigned long arg0);
  // proto:  ulong QInputEvent::timestamp();
extern void demth_ZNK11QInputEvent9timestampEv(void* qthis);
  // proto:  void QInputEvent::~QInputEvent();
extern void _ZN11QInputEventD0Ev(void* qthis);
  // proto:  int QKeyEvent::count();
extern void demth_ZNK9QKeyEvent5countEv(void* qthis);
  // proto:  void QKeyEvent::~QKeyEvent();
extern void _ZN9QKeyEventD0Ev(void* qthis);
  // proto:  QString QKeyEvent::text();
extern void demth_ZNK9QKeyEvent4textEv(void* qthis);
  // proto:  quint32 QKeyEvent::nativeVirtualKey();
extern void demth_ZNK9QKeyEvent16nativeVirtualKeyEv(void* qthis);
  // proto:  bool QKeyEvent::isAutoRepeat();
extern void demth_ZNK9QKeyEvent12isAutoRepeatEv(void* qthis);
  // proto:  int QKeyEvent::key();
extern void _ZNK9QKeyEvent3keyEv(void* qthis);
  // proto:  quint32 QKeyEvent::nativeModifiers();
extern void demth_ZNK9QKeyEvent15nativeModifiersEv(void* qthis);
  // proto:  quint32 QKeyEvent::nativeScanCode();
extern void demth_ZNK9QKeyEvent14nativeScanCodeEv(void* qthis);
  // proto:  const QPoint & QContextMenuEvent::globalPos();
extern void demth_ZNK17QContextMenuEvent9globalPosEv(void* qthis);
  // proto:  int QContextMenuEvent::globalY();
extern void demth_ZNK17QContextMenuEvent7globalYEv(void* qthis);
  // proto:  int QContextMenuEvent::globalX();
extern void demth_ZNK17QContextMenuEvent7globalXEv(void* qthis);
  // proto:  const QPoint & QContextMenuEvent::pos();
extern void demth_ZNK17QContextMenuEvent3posEv(void* qthis);
  // proto:  int QContextMenuEvent::y();
extern void demth_ZNK17QContextMenuEvent1yEv(void* qthis);
  // proto:  int QContextMenuEvent::x();
extern void demth_ZNK17QContextMenuEvent1xEv(void* qthis);
  // proto:  void QContextMenuEvent::~QContextMenuEvent();
extern void _ZN17QContextMenuEventD0Ev(void* qthis);
  // proto:  void QScrollPrepareEvent::setContentPosRange(const QRectF & rect);
extern void _ZN19QScrollPrepareEvent18setContentPosRangeERK6QRectF(void* qthis, void* arg0);
  // proto:  void QScrollPrepareEvent::setContentPos(const QPointF & pos);
extern void _ZN19QScrollPrepareEvent13setContentPosERK7QPointF(void* qthis, void* arg0);
  // proto:  QRectF QScrollPrepareEvent::contentPosRange();
extern void _ZNK19QScrollPrepareEvent15contentPosRangeEv(void* qthis);
  // proto:  QPointF QScrollPrepareEvent::contentPos();
extern void _ZNK19QScrollPrepareEvent10contentPosEv(void* qthis);
  // proto:  void QScrollPrepareEvent::setViewportSize(const QSizeF & size);
extern void _ZN19QScrollPrepareEvent15setViewportSizeERK6QSizeF(void* qthis, void* arg0);
  // proto:  void QScrollPrepareEvent::QScrollPrepareEvent(const QPointF & startPos);
extern void* dector_ZN19QScrollPrepareEventC1ERK7QPointF(void* arg0);
extern void _ZN19QScrollPrepareEventC1ERK7QPointF(void* qthis, void* arg0);
  // proto:  QPointF QScrollPrepareEvent::startPos();
extern void _ZNK19QScrollPrepareEvent8startPosEv(void* qthis);
  // proto:  QSizeF QScrollPrepareEvent::viewportSize();
extern void _ZNK19QScrollPrepareEvent12viewportSizeEv(void* qthis);
  // proto:  void QScrollPrepareEvent::~QScrollPrepareEvent();
extern void _ZN19QScrollPrepareEventD0Ev(void* qthis);
  // proto:  const QKeySequence & QShortcutEvent::key();
extern void demth_ZNK14QShortcutEvent3keyEv(void* qthis);
  // proto:  void QShortcutEvent::~QShortcutEvent();
extern void _ZN14QShortcutEventD0Ev(void* qthis);
  // proto:  bool QShortcutEvent::isAmbiguous();
extern void demth_ZNK14QShortcutEvent11isAmbiguousEv(void* qthis);
  // proto:  void QShortcutEvent::QShortcutEvent(const QKeySequence & key, int id, bool ambiguous);
extern void* dector_ZN14QShortcutEventC1ERK12QKeySequenceib(void* arg0, int arg1, bool arg2);
extern void _ZN14QShortcutEventC1ERK12QKeySequenceib(void* qthis, void* arg0, int arg1, bool arg2);
  // proto:  int QShortcutEvent::shortcutId();
extern void demth_ZNK14QShortcutEvent10shortcutIdEv(void* qthis);
  // proto:  bool QWindowStateChangeEvent::isOverride();
extern void _ZNK23QWindowStateChangeEvent10isOverrideEv(void* qthis);
  // proto:  void QWindowStateChangeEvent::~QWindowStateChangeEvent();
extern void _ZN23QWindowStateChangeEventD0Ev(void* qthis);
  // proto:  void QInputMethodQueryEvent::~QInputMethodQueryEvent();
extern void _ZN22QInputMethodQueryEventD0Ev(void* qthis);
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

// class sizeof(QWhatsThisClickedEvent)=32
type QWhatsThisClickedEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QExposeEvent)=32
type QExposeEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QInputMethodEvent)=1
type QInputMethodEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QHelpEvent)=40
type QHelpEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QActionEvent)=40
type QActionEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QMouseEvent)=1
type QMouseEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QFileOpenEvent)=40
type QFileOpenEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QToolBarChangeEvent)=24
type QToolBarChangeEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTabletEvent)=1
type QTabletEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTouchEvent)=1
type QTouchEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QScreenOrientationChangeEvent)=40
type QScreenOrientationChangeEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QIconDragEvent)=24
type QIconDragEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QCloseEvent)=24
type QCloseEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QDragEnterEvent)=1
type QDragEnterEvent struct {
  /*qbase*/ QDragMoveEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QWheelEvent)=1
type QWheelEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QScrollEvent)=64
type QScrollEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QHoverEvent)=1
type QHoverEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QDragMoveEvent)=1
type QDragMoveEvent struct {
  /*qbase*/ QDropEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QShowEvent)=24
type QShowEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QPlatformSurfaceEvent)=24
type QPlatformSurfaceEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QPaintEvent)=56
type QPaintEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QFocusEvent)=24
type QFocusEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QNativeGestureEvent)=1
type QNativeGestureEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QResizeEvent)=40
type QResizeEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStatusTipEvent)=32
type QStatusTipEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QEnterEvent)=72
type QEnterEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QMoveEvent)=40
type QMoveEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QHideEvent)=24
type QHideEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QDragLeaveEvent)=24
type QDragLeaveEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QDropEvent)=1
type QDropEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QInputEvent)=1
type QInputEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QApplicationStateChangeEvent)=24
type QApplicationStateChangeEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QKeyEvent)=1
type QKeyEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QContextMenuEvent)=1
type QContextMenuEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QScrollPrepareEvent)=112
type QScrollPrepareEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QShortcutEvent)=40
type QShortcutEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QWindowStateChangeEvent)=1
type QWindowStateChangeEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QInputMethodQueryEvent)=1
type QInputMethodQueryEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QString QWhatsThisClickedEvent::href();
func (this *QWhatsThisClickedEvent) href(args ...interface{}) () {
  // href()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QWhatsThisClickedEvent4hrefEv
  default:
    qtrt.ErrorResolve("QWhatsThisClickedEvent", "href", args)
  }

}

  // proto:  void QWhatsThisClickedEvent::~QWhatsThisClickedEvent();
func (this *QWhatsThisClickedEvent) FreeQWhatsThisClickedEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWhatsThisClickedEvent", "~QWhatsThisClickedEvent", args)
  }

}

  // proto:  void QWhatsThisClickedEvent::QWhatsThisClickedEvent(const QString & href);
func NewQWhatsThisClickedEvent(args ...interface{}) QWhatsThisClickedEvent {
  return QWhatsThisClickedEvent{}
}

  // proto:  void QExposeEvent::QExposeEvent(const QRegion & rgn);
func NewQExposeEvent(args ...interface{}) QExposeEvent {
  return QExposeEvent{}
}

  // proto:  const QRegion & QExposeEvent::region();
func (this *QExposeEvent) region(args ...interface{}) () {
  // region()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QExposeEvent6regionEv
  default:
    qtrt.ErrorResolve("QExposeEvent", "region", args)
  }

}

  // proto:  void QExposeEvent::~QExposeEvent();
func (this *QExposeEvent) FreeQExposeEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QExposeEvent", "~QExposeEvent", args)
  }

}

  // proto:  const QString & QInputMethodEvent::preeditString();
func (this *QInputMethodEvent) preeditString(args ...interface{}) () {
  // preeditString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QInputMethodEvent13preeditStringEv
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "preeditString", args)
  }

}

  // proto:  void QInputMethodEvent::QInputMethodEvent();
func NewQInputMethodEvent(args ...interface{}) QInputMethodEvent {
  return QInputMethodEvent{}
}

  // proto:  int QInputMethodEvent::replacementStart();
func (this *QInputMethodEvent) replacementStart(args ...interface{}) () {
  // replacementStart()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QInputMethodEvent16replacementStartEv
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "replacementStart", args)
  }

}

  // proto:  const QString & QInputMethodEvent::commitString();
func (this *QInputMethodEvent) commitString(args ...interface{}) () {
  // commitString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QInputMethodEvent12commitStringEv
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "commitString", args)
  }

}

  // proto:  void QInputMethodEvent::setCommitString(const QString & commitString, int replaceFrom, int replaceLength);
func (this *QInputMethodEvent) setCommitString(args ...interface{}) () {
  // setCommitString(const class QString &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QInputMethodEvent15setCommitStringERK7QStringii
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "setCommitString", args)
  }

}

  // proto:  int QInputMethodEvent::replacementLength();
func (this *QInputMethodEvent) replacementLength(args ...interface{}) () {
  // replacementLength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QInputMethodEvent17replacementLengthEv
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "replacementLength", args)
  }

}

  // proto:  const QPoint & QHelpEvent::globalPos();
func (this *QHelpEvent) globalPos(args ...interface{}) () {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent9globalPosEv
  default:
    qtrt.ErrorResolve("QHelpEvent", "globalPos", args)
  }

}

  // proto:  int QHelpEvent::globalX();
func (this *QHelpEvent) globalX(args ...interface{}) () {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent7globalXEv
  default:
    qtrt.ErrorResolve("QHelpEvent", "globalX", args)
  }

}

  // proto:  const QPoint & QHelpEvent::pos();
func (this *QHelpEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent3posEv
  default:
    qtrt.ErrorResolve("QHelpEvent", "pos", args)
  }

}

  // proto:  int QHelpEvent::y();
func (this *QHelpEvent) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent1yEv
  default:
    qtrt.ErrorResolve("QHelpEvent", "y", args)
  }

}

  // proto:  int QHelpEvent::globalY();
func (this *QHelpEvent) globalY(args ...interface{}) () {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent7globalYEv
  default:
    qtrt.ErrorResolve("QHelpEvent", "globalY", args)
  }

}

  // proto:  int QHelpEvent::x();
func (this *QHelpEvent) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent1xEv
  default:
    qtrt.ErrorResolve("QHelpEvent", "x", args)
  }

}

  // proto:  void QHelpEvent::~QHelpEvent();
func (this *QHelpEvent) FreeQHelpEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QHelpEvent", "~QHelpEvent", args)
  }

}

  // proto:  void QActionEvent::QActionEvent(int type, QAction * action, QAction * before);
func NewQActionEvent(args ...interface{}) QActionEvent {
  return QActionEvent{}
}

  // proto:  QAction * QActionEvent::before();
func (this *QActionEvent) before(args ...interface{}) () {
  // before()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QActionEvent6beforeEv
  default:
    qtrt.ErrorResolve("QActionEvent", "before", args)
  }

}

  // proto:  QAction * QActionEvent::action();
func (this *QActionEvent) action(args ...interface{}) () {
  // action()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QActionEvent6actionEv
  default:
    qtrt.ErrorResolve("QActionEvent", "action", args)
  }

}

  // proto:  void QActionEvent::~QActionEvent();
func (this *QActionEvent) FreeQActionEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QActionEvent", "~QActionEvent", args)
  }

}

  // proto:  QPoint QMouseEvent::globalPos();
func (this *QMouseEvent) globalPos(args ...interface{}) () {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent9globalPosEv
  default:
    qtrt.ErrorResolve("QMouseEvent", "globalPos", args)
  }

}

  // proto:  int QMouseEvent::y();
func (this *QMouseEvent) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent1yEv
  default:
    qtrt.ErrorResolve("QMouseEvent", "y", args)
  }

}

  // proto:  const QPointF & QMouseEvent::screenPos();
func (this *QMouseEvent) screenPos(args ...interface{}) () {
  // screenPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent9screenPosEv
  default:
    qtrt.ErrorResolve("QMouseEvent", "screenPos", args)
  }

}

  // proto:  int QMouseEvent::x();
func (this *QMouseEvent) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent1xEv
  default:
    qtrt.ErrorResolve("QMouseEvent", "x", args)
  }

}

  // proto:  const QPointF & QMouseEvent::localPos();
func (this *QMouseEvent) localPos(args ...interface{}) () {
  // localPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent8localPosEv
  default:
    qtrt.ErrorResolve("QMouseEvent", "localPos", args)
  }

}

  // proto:  int QMouseEvent::globalX();
func (this *QMouseEvent) globalX(args ...interface{}) () {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent7globalXEv
  default:
    qtrt.ErrorResolve("QMouseEvent", "globalX", args)
  }

}

  // proto:  const QPointF & QMouseEvent::windowPos();
func (this *QMouseEvent) windowPos(args ...interface{}) () {
  // windowPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent9windowPosEv
  default:
    qtrt.ErrorResolve("QMouseEvent", "windowPos", args)
  }

}

  // proto:  void QMouseEvent::~QMouseEvent();
func (this *QMouseEvent) FreeQMouseEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMouseEvent", "~QMouseEvent", args)
  }

}

  // proto:  int QMouseEvent::globalY();
func (this *QMouseEvent) globalY(args ...interface{}) () {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent7globalYEv
  default:
    qtrt.ErrorResolve("QMouseEvent", "globalY", args)
  }

}

  // proto:  QPoint QMouseEvent::pos();
func (this *QMouseEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent3posEv
  default:
    qtrt.ErrorResolve("QMouseEvent", "pos", args)
  }

}

  // proto:  void QFileOpenEvent::QFileOpenEvent(const QString & file);
func NewQFileOpenEvent(args ...interface{}) QFileOpenEvent {
  return QFileOpenEvent{}
}

  // proto:  void QFileOpenEvent::~QFileOpenEvent();
func (this *QFileOpenEvent) FreeQFileOpenEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFileOpenEvent", "~QFileOpenEvent", args)
  }

}

  // proto:  QString QFileOpenEvent::file();
func (this *QFileOpenEvent) file(args ...interface{}) () {
  // file()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QFileOpenEvent4fileEv
  default:
    qtrt.ErrorResolve("QFileOpenEvent", "file", args)
  }

}

  // proto:  QUrl QFileOpenEvent::url();
func (this *QFileOpenEvent) url(args ...interface{}) () {
  // url()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QFileOpenEvent3urlEv
  default:
    qtrt.ErrorResolve("QFileOpenEvent", "url", args)
  }

}

  // proto:  void QToolBarChangeEvent::QToolBarChangeEvent(bool t);
func NewQToolBarChangeEvent(args ...interface{}) QToolBarChangeEvent {
  return QToolBarChangeEvent{}
}

  // proto:  void QToolBarChangeEvent::~QToolBarChangeEvent();
func (this *QToolBarChangeEvent) FreeQToolBarChangeEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QToolBarChangeEvent", "~QToolBarChangeEvent", args)
  }

}

  // proto:  bool QToolBarChangeEvent::toggle();
func (this *QToolBarChangeEvent) toggle(args ...interface{}) () {
  // toggle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QToolBarChangeEvent6toggleEv
  default:
    qtrt.ErrorResolve("QToolBarChangeEvent", "toggle", args)
  }

}

  // proto:  int QTabletEvent::x();
func (this *QTabletEvent) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent1xEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "x", args)
  }

}

  // proto:  int QTabletEvent::xTilt();
func (this *QTabletEvent) xTilt(args ...interface{}) () {
  // xTilt()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent5xTiltEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "xTilt", args)
  }

}

  // proto:  qint64 QTabletEvent::uniqueId();
func (this *QTabletEvent) uniqueId(args ...interface{}) () {
  // uniqueId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent8uniqueIdEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "uniqueId", args)
  }

}

  // proto:  const QPointF & QTabletEvent::globalPosF();
func (this *QTabletEvent) globalPosF(args ...interface{}) () {
  // globalPosF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent10globalPosFEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "globalPosF", args)
  }

}

  // proto:  int QTabletEvent::z();
func (this *QTabletEvent) z(args ...interface{}) () {
  // z()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent1zEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "z", args)
  }

}

  // proto:  int QTabletEvent::y();
func (this *QTabletEvent) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent1yEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "y", args)
  }

}

  // proto:  QPoint QTabletEvent::pos();
func (this *QTabletEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent3posEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "pos", args)
  }

}

  // proto:  qreal QTabletEvent::rotation();
func (this *QTabletEvent) rotation(args ...interface{}) () {
  // rotation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent8rotationEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "rotation", args)
  }

}

  // proto:  QPoint QTabletEvent::globalPos();
func (this *QTabletEvent) globalPos(args ...interface{}) () {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent9globalPosEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "globalPos", args)
  }

}

  // proto:  void QTabletEvent::~QTabletEvent();
func (this *QTabletEvent) FreeQTabletEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTabletEvent", "~QTabletEvent", args)
  }

}

  // proto:  qreal QTabletEvent::tangentialPressure();
func (this *QTabletEvent) tangentialPressure(args ...interface{}) () {
  // tangentialPressure()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent18tangentialPressureEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "tangentialPressure", args)
  }

}

  // proto:  qreal QTabletEvent::hiResGlobalX();
func (this *QTabletEvent) hiResGlobalX(args ...interface{}) () {
  // hiResGlobalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent12hiResGlobalXEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "hiResGlobalX", args)
  }

}

  // proto:  int QTabletEvent::globalY();
func (this *QTabletEvent) globalY(args ...interface{}) () {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent7globalYEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "globalY", args)
  }

}

  // proto:  qreal QTabletEvent::hiResGlobalY();
func (this *QTabletEvent) hiResGlobalY(args ...interface{}) () {
  // hiResGlobalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent12hiResGlobalYEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "hiResGlobalY", args)
  }

}

  // proto:  int QTabletEvent::globalX();
func (this *QTabletEvent) globalX(args ...interface{}) () {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent7globalXEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "globalX", args)
  }

}

  // proto:  const QPointF & QTabletEvent::posF();
func (this *QTabletEvent) posF(args ...interface{}) () {
  // posF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent4posFEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "posF", args)
  }

}

  // proto:  qreal QTabletEvent::pressure();
func (this *QTabletEvent) pressure(args ...interface{}) () {
  // pressure()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent8pressureEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "pressure", args)
  }

}

  // proto:  int QTabletEvent::yTilt();
func (this *QTabletEvent) yTilt(args ...interface{}) () {
  // yTilt()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent5yTiltEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "yTilt", args)
  }

}

  // proto:  void QTouchEvent::setDevice(QTouchDevice * adevice);
func (this *QTouchEvent) setDevice(args ...interface{}) () {
  // setDevice(class QTouchDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTouchDevice{}) // "QTouchDevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTouchEvent9setDeviceEP12QTouchDevice
  default:
    qtrt.ErrorResolve("QTouchEvent", "setDevice", args)
  }

}

  // proto:  QWindow * QTouchEvent::window();
func (this *QTouchEvent) window(args ...interface{}) () {
  // window()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTouchEvent6windowEv
  default:
    qtrt.ErrorResolve("QTouchEvent", "window", args)
  }

}

  // proto:  QTouchDevice * QTouchEvent::device();
func (this *QTouchEvent) device(args ...interface{}) () {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTouchEvent6deviceEv
  default:
    qtrt.ErrorResolve("QTouchEvent", "device", args)
  }

}

  // proto:  QObject * QTouchEvent::target();
func (this *QTouchEvent) target(args ...interface{}) () {
  // target()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTouchEvent6targetEv
  default:
    qtrt.ErrorResolve("QTouchEvent", "target", args)
  }

}

  // proto:  void QTouchEvent::~QTouchEvent();
func (this *QTouchEvent) FreeQTouchEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTouchEvent", "~QTouchEvent", args)
  }

}

  // proto:  void QTouchEvent::setWindow(QWindow * awindow);
func (this *QTouchEvent) setWindow(args ...interface{}) () {
  // setWindow(class QWindow *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWindow{}) // "QWindow *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTouchEvent9setWindowEP7QWindow
  default:
    qtrt.ErrorResolve("QTouchEvent", "setWindow", args)
  }

}

  // proto:  void QTouchEvent::setTarget(QObject * atarget);
func (this *QTouchEvent) setTarget(args ...interface{}) () {
  // setTarget(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTouchEvent9setTargetEP7QObject
  default:
    qtrt.ErrorResolve("QTouchEvent", "setTarget", args)
  }

}

  // proto:  QScreen * QScreenOrientationChangeEvent::screen();
func (this *QScreenOrientationChangeEvent) screen(args ...interface{}) () {
  // screen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK29QScreenOrientationChangeEvent6screenEv
  default:
    qtrt.ErrorResolve("QScreenOrientationChangeEvent", "screen", args)
  }

}

  // proto:  void QScreenOrientationChangeEvent::~QScreenOrientationChangeEvent();
func (this *QScreenOrientationChangeEvent) FreeQScreenOrientationChangeEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScreenOrientationChangeEvent", "~QScreenOrientationChangeEvent", args)
  }

}

  // proto:  void QIconDragEvent::~QIconDragEvent();
func (this *QIconDragEvent) FreeQIconDragEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QIconDragEvent", "~QIconDragEvent", args)
  }

}

  // proto:  void QIconDragEvent::QIconDragEvent();
func NewQIconDragEvent(args ...interface{}) QIconDragEvent {
  return QIconDragEvent{}
}

  // proto:  void QCloseEvent::~QCloseEvent();
func (this *QCloseEvent) FreeQCloseEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCloseEvent", "~QCloseEvent", args)
  }

}

  // proto:  void QCloseEvent::QCloseEvent();
func NewQCloseEvent(args ...interface{}) QCloseEvent {
  return QCloseEvent{}
}

  // proto:  void QDragEnterEvent::~QDragEnterEvent();
func (this *QDragEnterEvent) FreeQDragEnterEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDragEnterEvent", "~QDragEnterEvent", args)
  }

}

  // proto:  int QWheelEvent::x();
func (this *QWheelEvent) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent1xEv
  default:
    qtrt.ErrorResolve("QWheelEvent", "x", args)
  }

}

  // proto:  QPoint QWheelEvent::angleDelta();
func (this *QWheelEvent) angleDelta(args ...interface{}) () {
  // angleDelta()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent10angleDeltaEv
  default:
    qtrt.ErrorResolve("QWheelEvent", "angleDelta", args)
  }

}

  // proto:  QPoint QWheelEvent::pos();
func (this *QWheelEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent3posEv
  default:
    qtrt.ErrorResolve("QWheelEvent", "pos", args)
  }

}

  // proto:  int QWheelEvent::globalY();
func (this *QWheelEvent) globalY(args ...interface{}) () {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent7globalYEv
  default:
    qtrt.ErrorResolve("QWheelEvent", "globalY", args)
  }

}

  // proto:  const QPointF & QWheelEvent::posF();
func (this *QWheelEvent) posF(args ...interface{}) () {
  // posF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent4posFEv
  default:
    qtrt.ErrorResolve("QWheelEvent", "posF", args)
  }

}

  // proto:  int QWheelEvent::globalX();
func (this *QWheelEvent) globalX(args ...interface{}) () {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent7globalXEv
  default:
    qtrt.ErrorResolve("QWheelEvent", "globalX", args)
  }

}

  // proto:  int QWheelEvent::y();
func (this *QWheelEvent) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent1yEv
  default:
    qtrt.ErrorResolve("QWheelEvent", "y", args)
  }

}

  // proto:  void QWheelEvent::~QWheelEvent();
func (this *QWheelEvent) FreeQWheelEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWheelEvent", "~QWheelEvent", args)
  }

}

  // proto:  QPoint QWheelEvent::pixelDelta();
func (this *QWheelEvent) pixelDelta(args ...interface{}) () {
  // pixelDelta()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent10pixelDeltaEv
  default:
    qtrt.ErrorResolve("QWheelEvent", "pixelDelta", args)
  }

}

  // proto:  int QWheelEvent::delta();
func (this *QWheelEvent) delta(args ...interface{}) () {
  // delta()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent5deltaEv
  default:
    qtrt.ErrorResolve("QWheelEvent", "delta", args)
  }

}

  // proto:  QPoint QWheelEvent::globalPos();
func (this *QWheelEvent) globalPos(args ...interface{}) () {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent9globalPosEv
  default:
    qtrt.ErrorResolve("QWheelEvent", "globalPos", args)
  }

}

  // proto:  const QPointF & QWheelEvent::globalPosF();
func (this *QWheelEvent) globalPosF(args ...interface{}) () {
  // globalPosF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent10globalPosFEv
  default:
    qtrt.ErrorResolve("QWheelEvent", "globalPosF", args)
  }

}

  // proto:  QPointF QScrollEvent::contentPos();
func (this *QScrollEvent) contentPos(args ...interface{}) () {
  // contentPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QScrollEvent10contentPosEv
  default:
    qtrt.ErrorResolve("QScrollEvent", "contentPos", args)
  }

}

  // proto:  QPointF QScrollEvent::overshootDistance();
func (this *QScrollEvent) overshootDistance(args ...interface{}) () {
  // overshootDistance()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QScrollEvent17overshootDistanceEv
  default:
    qtrt.ErrorResolve("QScrollEvent", "overshootDistance", args)
  }

}

  // proto:  void QScrollEvent::~QScrollEvent();
func (this *QScrollEvent) FreeQScrollEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScrollEvent", "~QScrollEvent", args)
  }

}

  // proto:  void QHoverEvent::~QHoverEvent();
func (this *QHoverEvent) FreeQHoverEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QHoverEvent", "~QHoverEvent", args)
  }

}

  // proto:  const QPointF & QHoverEvent::posF();
func (this *QHoverEvent) posF(args ...interface{}) () {
  // posF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHoverEvent4posFEv
  default:
    qtrt.ErrorResolve("QHoverEvent", "posF", args)
  }

}

  // proto:  QPoint QHoverEvent::oldPos();
func (this *QHoverEvent) oldPos(args ...interface{}) () {
  // oldPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHoverEvent6oldPosEv
  default:
    qtrt.ErrorResolve("QHoverEvent", "oldPos", args)
  }

}

  // proto:  const QPointF & QHoverEvent::oldPosF();
func (this *QHoverEvent) oldPosF(args ...interface{}) () {
  // oldPosF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHoverEvent7oldPosFEv
  default:
    qtrt.ErrorResolve("QHoverEvent", "oldPosF", args)
  }

}

  // proto:  QPoint QHoverEvent::pos();
func (this *QHoverEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHoverEvent3posEv
  default:
    qtrt.ErrorResolve("QHoverEvent", "pos", args)
  }

}

  // proto:  void QDragMoveEvent::accept(const QRect & r);
func (this *QDragMoveEvent) accept(args ...interface{}) () {
  // accept(const class QRect &)
  // accept()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDragMoveEvent6acceptERK5QRect
  case 1:
    // invoke: _ZN14QDragMoveEvent6acceptEv
  default:
    qtrt.ErrorResolve("QDragMoveEvent", "accept", args)
  }

}

  // proto:  QRect QDragMoveEvent::answerRect();
func (this *QDragMoveEvent) answerRect(args ...interface{}) () {
  // answerRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDragMoveEvent10answerRectEv
  default:
    qtrt.ErrorResolve("QDragMoveEvent", "answerRect", args)
  }

}

  // proto:  void QDragMoveEvent::ignore(const QRect & r);
func (this *QDragMoveEvent) ignore(args ...interface{}) () {
  // ignore(const class QRect &)
  // ignore()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDragMoveEvent6ignoreERK5QRect
  case 1:
    // invoke: _ZN14QDragMoveEvent6ignoreEv
  default:
    qtrt.ErrorResolve("QDragMoveEvent", "ignore", args)
  }

}

  // proto:  void QDragMoveEvent::~QDragMoveEvent();
func (this *QDragMoveEvent) FreeQDragMoveEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDragMoveEvent", "~QDragMoveEvent", args)
  }

}

  // proto:  void QShowEvent::~QShowEvent();
func (this *QShowEvent) FreeQShowEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QShowEvent", "~QShowEvent", args)
  }

}

  // proto:  void QShowEvent::QShowEvent();
func NewQShowEvent(args ...interface{}) QShowEvent {
  return QShowEvent{}
}

  // proto:  void QPlatformSurfaceEvent::~QPlatformSurfaceEvent();
func (this *QPlatformSurfaceEvent) FreeQPlatformSurfaceEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPlatformSurfaceEvent", "~QPlatformSurfaceEvent", args)
  }

}

  // proto:  void QPaintEvent::~QPaintEvent();
func (this *QPaintEvent) FreeQPaintEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPaintEvent", "~QPaintEvent", args)
  }

}

  // proto:  const QRect & QPaintEvent::rect();
func (this *QPaintEvent) rect(args ...interface{}) () {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPaintEvent4rectEv
  default:
    qtrt.ErrorResolve("QPaintEvent", "rect", args)
  }

}

  // proto:  void QPaintEvent::QPaintEvent(const QRect & paintRect);
func NewQPaintEvent(args ...interface{}) QPaintEvent {
  return QPaintEvent{}
}

  // proto:  const QRegion & QPaintEvent::region();
func (this *QPaintEvent) region(args ...interface{}) () {
  // region()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPaintEvent6regionEv
  default:
    qtrt.ErrorResolve("QPaintEvent", "region", args)
  }

}

  // proto:  bool QFocusEvent::lostFocus();
func (this *QFocusEvent) lostFocus(args ...interface{}) () {
  // lostFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFocusEvent9lostFocusEv
  default:
    qtrt.ErrorResolve("QFocusEvent", "lostFocus", args)
  }

}

  // proto:  bool QFocusEvent::gotFocus();
func (this *QFocusEvent) gotFocus(args ...interface{}) () {
  // gotFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFocusEvent8gotFocusEv
  default:
    qtrt.ErrorResolve("QFocusEvent", "gotFocus", args)
  }

}

  // proto:  void QFocusEvent::~QFocusEvent();
func (this *QFocusEvent) FreeQFocusEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFocusEvent", "~QFocusEvent", args)
  }

}

  // proto:  const QPointF & QNativeGestureEvent::localPos();
func (this *QNativeGestureEvent) localPos(args ...interface{}) () {
  // localPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent8localPosEv
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "localPos", args)
  }

}

  // proto:  const QPointF & QNativeGestureEvent::screenPos();
func (this *QNativeGestureEvent) screenPos(args ...interface{}) () {
  // screenPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent9screenPosEv
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "screenPos", args)
  }

}

  // proto:  const QPoint QNativeGestureEvent::pos();
func (this *QNativeGestureEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent3posEv
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "pos", args)
  }

}

  // proto:  const QPoint QNativeGestureEvent::globalPos();
func (this *QNativeGestureEvent) globalPos(args ...interface{}) () {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent9globalPosEv
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "globalPos", args)
  }

}

  // proto:  qreal QNativeGestureEvent::value();
func (this *QNativeGestureEvent) value(args ...interface{}) () {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent5valueEv
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "value", args)
  }

}

  // proto:  const QPointF & QNativeGestureEvent::windowPos();
func (this *QNativeGestureEvent) windowPos(args ...interface{}) () {
  // windowPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent9windowPosEv
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "windowPos", args)
  }

}

  // proto:  const QSize & QResizeEvent::oldSize();
func (this *QResizeEvent) oldSize(args ...interface{}) () {
  // oldSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QResizeEvent7oldSizeEv
  default:
    qtrt.ErrorResolve("QResizeEvent", "oldSize", args)
  }

}

  // proto:  const QSize & QResizeEvent::size();
func (this *QResizeEvent) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QResizeEvent4sizeEv
  default:
    qtrt.ErrorResolve("QResizeEvent", "size", args)
  }

}

  // proto:  void QResizeEvent::~QResizeEvent();
func (this *QResizeEvent) FreeQResizeEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QResizeEvent", "~QResizeEvent", args)
  }

}

  // proto:  void QResizeEvent::QResizeEvent(const QSize & size, const QSize & oldSize);
func NewQResizeEvent(args ...interface{}) QResizeEvent {
  return QResizeEvent{}
}

  // proto:  void QStatusTipEvent::~QStatusTipEvent();
func (this *QStatusTipEvent) FreeQStatusTipEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStatusTipEvent", "~QStatusTipEvent", args)
  }

}

  // proto:  QString QStatusTipEvent::tip();
func (this *QStatusTipEvent) tip(args ...interface{}) () {
  // tip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QStatusTipEvent3tipEv
  default:
    qtrt.ErrorResolve("QStatusTipEvent", "tip", args)
  }

}

  // proto:  void QStatusTipEvent::QStatusTipEvent(const QString & tip);
func NewQStatusTipEvent(args ...interface{}) QStatusTipEvent {
  return QStatusTipEvent{}
}

  // proto:  int QEnterEvent::y();
func (this *QEnterEvent) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent1yEv
  default:
    qtrt.ErrorResolve("QEnterEvent", "y", args)
  }

}

  // proto:  QPoint QEnterEvent::pos();
func (this *QEnterEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent3posEv
  default:
    qtrt.ErrorResolve("QEnterEvent", "pos", args)
  }

}

  // proto:  void QEnterEvent::~QEnterEvent();
func (this *QEnterEvent) FreeQEnterEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QEnterEvent", "~QEnterEvent", args)
  }

}

  // proto:  const QPointF & QEnterEvent::screenPos();
func (this *QEnterEvent) screenPos(args ...interface{}) () {
  // screenPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent9screenPosEv
  default:
    qtrt.ErrorResolve("QEnterEvent", "screenPos", args)
  }

}

  // proto:  const QPointF & QEnterEvent::localPos();
func (this *QEnterEvent) localPos(args ...interface{}) () {
  // localPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent8localPosEv
  default:
    qtrt.ErrorResolve("QEnterEvent", "localPos", args)
  }

}

  // proto:  const QPointF & QEnterEvent::windowPos();
func (this *QEnterEvent) windowPos(args ...interface{}) () {
  // windowPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent9windowPosEv
  default:
    qtrt.ErrorResolve("QEnterEvent", "windowPos", args)
  }

}

  // proto:  int QEnterEvent::globalX();
func (this *QEnterEvent) globalX(args ...interface{}) () {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent7globalXEv
  default:
    qtrt.ErrorResolve("QEnterEvent", "globalX", args)
  }

}

  // proto:  int QEnterEvent::x();
func (this *QEnterEvent) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent1xEv
  default:
    qtrt.ErrorResolve("QEnterEvent", "x", args)
  }

}

  // proto:  QPoint QEnterEvent::globalPos();
func (this *QEnterEvent) globalPos(args ...interface{}) () {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent9globalPosEv
  default:
    qtrt.ErrorResolve("QEnterEvent", "globalPos", args)
  }

}

  // proto:  int QEnterEvent::globalY();
func (this *QEnterEvent) globalY(args ...interface{}) () {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent7globalYEv
  default:
    qtrt.ErrorResolve("QEnterEvent", "globalY", args)
  }

}

  // proto:  void QEnterEvent::QEnterEvent(const QPointF & localPos, const QPointF & windowPos, const QPointF & screenPos);
func NewQEnterEvent(args ...interface{}) QEnterEvent {
  return QEnterEvent{}
}

  // proto:  void QMoveEvent::~QMoveEvent();
func (this *QMoveEvent) FreeQMoveEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMoveEvent", "~QMoveEvent", args)
  }

}

  // proto:  const QPoint & QMoveEvent::oldPos();
func (this *QMoveEvent) oldPos(args ...interface{}) () {
  // oldPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMoveEvent6oldPosEv
  default:
    qtrt.ErrorResolve("QMoveEvent", "oldPos", args)
  }

}

  // proto:  void QMoveEvent::QMoveEvent(const QPoint & pos, const QPoint & oldPos);
func NewQMoveEvent(args ...interface{}) QMoveEvent {
  return QMoveEvent{}
}

  // proto:  const QPoint & QMoveEvent::pos();
func (this *QMoveEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMoveEvent3posEv
  default:
    qtrt.ErrorResolve("QMoveEvent", "pos", args)
  }

}

  // proto:  void QHideEvent::QHideEvent();
func NewQHideEvent(args ...interface{}) QHideEvent {
  return QHideEvent{}
}

  // proto:  void QHideEvent::~QHideEvent();
func (this *QHideEvent) FreeQHideEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QHideEvent", "~QHideEvent", args)
  }

}

  // proto:  void QDragLeaveEvent::~QDragLeaveEvent();
func (this *QDragLeaveEvent) FreeQDragLeaveEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDragLeaveEvent", "~QDragLeaveEvent", args)
  }

}

  // proto:  void QDragLeaveEvent::QDragLeaveEvent();
func NewQDragLeaveEvent(args ...interface{}) QDragLeaveEvent {
  return QDragLeaveEvent{}
}

  // proto:  void QDropEvent::~QDropEvent();
func (this *QDropEvent) FreeQDropEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDropEvent", "~QDropEvent", args)
  }

}

  // proto:  QPoint QDropEvent::pos();
func (this *QDropEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent3posEv
  default:
    qtrt.ErrorResolve("QDropEvent", "pos", args)
  }

}

  // proto:  QObject * QDropEvent::source();
func (this *QDropEvent) source(args ...interface{}) () {
  // source()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent6sourceEv
  default:
    qtrt.ErrorResolve("QDropEvent", "source", args)
  }

}

  // proto:  const QMimeData * QDropEvent::mimeData();
func (this *QDropEvent) mimeData(args ...interface{}) () {
  // mimeData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent8mimeDataEv
  default:
    qtrt.ErrorResolve("QDropEvent", "mimeData", args)
  }

}

  // proto:  void QDropEvent::acceptProposedAction();
func (this *QDropEvent) acceptProposedAction(args ...interface{}) () {
  // acceptProposedAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QDropEvent20acceptProposedActionEv
  default:
    qtrt.ErrorResolve("QDropEvent", "acceptProposedAction", args)
  }

}

  // proto:  const QPointF & QDropEvent::posF();
func (this *QDropEvent) posF(args ...interface{}) () {
  // posF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent4posFEv
  default:
    qtrt.ErrorResolve("QDropEvent", "posF", args)
  }

}

  // proto:  void QInputEvent::setTimestamp(ulong atimestamp);
func (this *QInputEvent) setTimestamp(args ...interface{}) () {
  // setTimestamp(ulong)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "ulong"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QInputEvent12setTimestampEm
  default:
    qtrt.ErrorResolve("QInputEvent", "setTimestamp", args)
  }

}

  // proto:  ulong QInputEvent::timestamp();
func (this *QInputEvent) timestamp(args ...interface{}) () {
  // timestamp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QInputEvent9timestampEv
  default:
    qtrt.ErrorResolve("QInputEvent", "timestamp", args)
  }

}

  // proto:  void QInputEvent::~QInputEvent();
func (this *QInputEvent) FreeQInputEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QInputEvent", "~QInputEvent", args)
  }

}

  // proto:  int QKeyEvent::count();
func (this *QKeyEvent) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent5countEv
  default:
    qtrt.ErrorResolve("QKeyEvent", "count", args)
  }

}

  // proto:  void QKeyEvent::~QKeyEvent();
func (this *QKeyEvent) FreeQKeyEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QKeyEvent", "~QKeyEvent", args)
  }

}

  // proto:  QString QKeyEvent::text();
func (this *QKeyEvent) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent4textEv
  default:
    qtrt.ErrorResolve("QKeyEvent", "text", args)
  }

}

  // proto:  quint32 QKeyEvent::nativeVirtualKey();
func (this *QKeyEvent) nativeVirtualKey(args ...interface{}) () {
  // nativeVirtualKey()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent16nativeVirtualKeyEv
  default:
    qtrt.ErrorResolve("QKeyEvent", "nativeVirtualKey", args)
  }

}

  // proto:  bool QKeyEvent::isAutoRepeat();
func (this *QKeyEvent) isAutoRepeat(args ...interface{}) () {
  // isAutoRepeat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent12isAutoRepeatEv
  default:
    qtrt.ErrorResolve("QKeyEvent", "isAutoRepeat", args)
  }

}

  // proto:  int QKeyEvent::key();
func (this *QKeyEvent) key(args ...interface{}) () {
  // key()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent3keyEv
  default:
    qtrt.ErrorResolve("QKeyEvent", "key", args)
  }

}

  // proto:  quint32 QKeyEvent::nativeModifiers();
func (this *QKeyEvent) nativeModifiers(args ...interface{}) () {
  // nativeModifiers()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent15nativeModifiersEv
  default:
    qtrt.ErrorResolve("QKeyEvent", "nativeModifiers", args)
  }

}

  // proto:  quint32 QKeyEvent::nativeScanCode();
func (this *QKeyEvent) nativeScanCode(args ...interface{}) () {
  // nativeScanCode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent14nativeScanCodeEv
  default:
    qtrt.ErrorResolve("QKeyEvent", "nativeScanCode", args)
  }

}

  // proto:  const QPoint & QContextMenuEvent::globalPos();
func (this *QContextMenuEvent) globalPos(args ...interface{}) () {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent9globalPosEv
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "globalPos", args)
  }

}

  // proto:  int QContextMenuEvent::globalY();
func (this *QContextMenuEvent) globalY(args ...interface{}) () {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent7globalYEv
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "globalY", args)
  }

}

  // proto:  int QContextMenuEvent::globalX();
func (this *QContextMenuEvent) globalX(args ...interface{}) () {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent7globalXEv
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "globalX", args)
  }

}

  // proto:  const QPoint & QContextMenuEvent::pos();
func (this *QContextMenuEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent3posEv
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "pos", args)
  }

}

  // proto:  int QContextMenuEvent::y();
func (this *QContextMenuEvent) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent1yEv
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "y", args)
  }

}

  // proto:  int QContextMenuEvent::x();
func (this *QContextMenuEvent) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent1xEv
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "x", args)
  }

}

  // proto:  void QContextMenuEvent::~QContextMenuEvent();
func (this *QContextMenuEvent) FreeQContextMenuEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "~QContextMenuEvent", args)
  }

}

  // proto:  void QScrollPrepareEvent::setContentPosRange(const QRectF & rect);
func (this *QScrollPrepareEvent) setContentPosRange(args ...interface{}) () {
  // setContentPosRange(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QScrollPrepareEvent18setContentPosRangeERK6QRectF
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "setContentPosRange", args)
  }

}

  // proto:  void QScrollPrepareEvent::setContentPos(const QPointF & pos);
func (this *QScrollPrepareEvent) setContentPos(args ...interface{}) () {
  // setContentPos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QScrollPrepareEvent13setContentPosERK7QPointF
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "setContentPos", args)
  }

}

  // proto:  QRectF QScrollPrepareEvent::contentPosRange();
func (this *QScrollPrepareEvent) contentPosRange(args ...interface{}) () {
  // contentPosRange()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QScrollPrepareEvent15contentPosRangeEv
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "contentPosRange", args)
  }

}

  // proto:  QPointF QScrollPrepareEvent::contentPos();
func (this *QScrollPrepareEvent) contentPos(args ...interface{}) () {
  // contentPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QScrollPrepareEvent10contentPosEv
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "contentPos", args)
  }

}

  // proto:  void QScrollPrepareEvent::setViewportSize(const QSizeF & size);
func (this *QScrollPrepareEvent) setViewportSize(args ...interface{}) () {
  // setViewportSize(const class QSizeF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QScrollPrepareEvent15setViewportSizeERK6QSizeF
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "setViewportSize", args)
  }

}

  // proto:  void QScrollPrepareEvent::QScrollPrepareEvent(const QPointF & startPos);
func NewQScrollPrepareEvent(args ...interface{}) QScrollPrepareEvent {
  return QScrollPrepareEvent{}
}

  // proto:  QPointF QScrollPrepareEvent::startPos();
func (this *QScrollPrepareEvent) startPos(args ...interface{}) () {
  // startPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QScrollPrepareEvent8startPosEv
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "startPos", args)
  }

}

  // proto:  QSizeF QScrollPrepareEvent::viewportSize();
func (this *QScrollPrepareEvent) viewportSize(args ...interface{}) () {
  // viewportSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QScrollPrepareEvent12viewportSizeEv
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "viewportSize", args)
  }

}

  // proto:  void QScrollPrepareEvent::~QScrollPrepareEvent();
func (this *QScrollPrepareEvent) FreeQScrollPrepareEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "~QScrollPrepareEvent", args)
  }

}

  // proto:  const QKeySequence & QShortcutEvent::key();
func (this *QShortcutEvent) key(args ...interface{}) () {
  // key()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QShortcutEvent3keyEv
  default:
    qtrt.ErrorResolve("QShortcutEvent", "key", args)
  }

}

  // proto:  void QShortcutEvent::~QShortcutEvent();
func (this *QShortcutEvent) FreeQShortcutEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QShortcutEvent", "~QShortcutEvent", args)
  }

}

  // proto:  bool QShortcutEvent::isAmbiguous();
func (this *QShortcutEvent) isAmbiguous(args ...interface{}) () {
  // isAmbiguous()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QShortcutEvent11isAmbiguousEv
  default:
    qtrt.ErrorResolve("QShortcutEvent", "isAmbiguous", args)
  }

}

  // proto:  void QShortcutEvent::QShortcutEvent(const QKeySequence & key, int id, bool ambiguous);
func NewQShortcutEvent(args ...interface{}) QShortcutEvent {
  return QShortcutEvent{}
}

  // proto:  int QShortcutEvent::shortcutId();
func (this *QShortcutEvent) shortcutId(args ...interface{}) () {
  // shortcutId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QShortcutEvent10shortcutIdEv
  default:
    qtrt.ErrorResolve("QShortcutEvent", "shortcutId", args)
  }

}

  // proto:  bool QWindowStateChangeEvent::isOverride();
func (this *QWindowStateChangeEvent) isOverride(args ...interface{}) () {
  // isOverride()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QWindowStateChangeEvent10isOverrideEv
  default:
    qtrt.ErrorResolve("QWindowStateChangeEvent", "isOverride", args)
  }

}

  // proto:  void QWindowStateChangeEvent::~QWindowStateChangeEvent();
func (this *QWindowStateChangeEvent) FreeQWindowStateChangeEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWindowStateChangeEvent", "~QWindowStateChangeEvent", args)
  }

}

  // proto:  void QInputMethodQueryEvent::~QInputMethodQueryEvent();
func (this *QInputMethodQueryEvent) FreeQInputMethodQueryEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QInputMethodQueryEvent", "~QInputMethodQueryEvent", args)
  }

}

// <= body block end

