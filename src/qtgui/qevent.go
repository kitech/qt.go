package qtgui
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QWhatsThisClickedEvent::QWhatsThisClickedEvent(const QString & href);
extern void* C_ZN22QWhatsThisClickedEventC2ERK7QString(void* arg0); // 3
  // proto:  void QWhatsThisClickedEvent::~QWhatsThisClickedEvent();
extern void C_ZN22QWhatsThisClickedEventD2Ev(void* qthis); // 4
  // proto:  QString QWhatsThisClickedEvent::href();
extern void* C_ZNK22QWhatsThisClickedEvent4hrefEv(void* qthis); // 2
  // proto:  const QRegion & QExposeEvent::region();
extern void* C_ZNK12QExposeEvent6regionEv(void* qthis); // 2
  // proto:  void QExposeEvent::~QExposeEvent();
extern void C_ZN12QExposeEventD2Ev(void* qthis); // 4
  // proto:  void QExposeEvent::QExposeEvent(const QRegion & rgn);
extern void* C_ZN12QExposeEventC2ERK7QRegion(void* arg0); // 3
  // proto:  int QInputMethodEvent::replacementStart();
extern int32_t C_ZNK17QInputMethodEvent16replacementStartEv(void* qthis); // 2
  // proto:  int QInputMethodEvent::replacementLength();
extern int32_t C_ZNK17QInputMethodEvent17replacementLengthEv(void* qthis); // 2
  // proto:  void QInputMethodEvent::QInputMethodEvent(const QInputMethodEvent & other);
extern void* C_ZN17QInputMethodEventC2ERKS_(void* arg0); // 3
  // proto:  void QInputMethodEvent::QInputMethodEvent();
extern void* C_ZN17QInputMethodEventC2Ev(); // 3
  // proto:  void QInputMethodEvent::~QInputMethodEvent();
extern void C_ZN17QInputMethodEventD2Ev(void* qthis); // 4
  // proto:  const QString & QInputMethodEvent::preeditString();
extern void* C_ZNK17QInputMethodEvent13preeditStringEv(void* qthis); // 2
  // proto:  const QList<QInputMethodEvent::Attribute> & QInputMethodEvent::attributes();
extern void C_ZNK17QInputMethodEvent10attributesEv(void* qthis); // 2
  // proto:  const QString & QInputMethodEvent::commitString();
extern void* C_ZNK17QInputMethodEvent12commitStringEv(void* qthis); // 2
  // proto:  void QInputMethodEvent::setCommitString(const QString & commitString, int replaceFrom, int replaceLength);
extern void C_ZN17QInputMethodEvent15setCommitStringERK7QStringii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  const QPoint & QHelpEvent::globalPos();
extern void* C_ZNK10QHelpEvent9globalPosEv(void* qthis); // 2
  // proto:  const QPoint & QHelpEvent::pos();
extern void* C_ZNK10QHelpEvent3posEv(void* qthis); // 2
  // proto:  int QHelpEvent::y();
extern void C_ZNK10QHelpEvent1yEv(void* qthis); // 2
  // proto:  int QHelpEvent::x();
extern void C_ZNK10QHelpEvent1xEv(void* qthis); // 2
  // proto:  int QHelpEvent::globalX();
extern int32_t C_ZNK10QHelpEvent7globalXEv(void* qthis); // 2
  // proto:  int QHelpEvent::globalY();
extern int32_t C_ZNK10QHelpEvent7globalYEv(void* qthis); // 2
  // proto:  void QHelpEvent::~QHelpEvent();
extern void C_ZN10QHelpEventD2Ev(void* qthis); // 4
  // proto:  const QPointF & QMouseEvent::localPos();
extern void* C_ZNK11QMouseEvent8localPosEv(void* qthis); // 2
  // proto:  const QPointF & QMouseEvent::screenPos();
extern void* C_ZNK11QMouseEvent9screenPosEv(void* qthis); // 2
  // proto:  QPoint QMouseEvent::globalPos();
extern void* C_ZNK11QMouseEvent9globalPosEv(void* qthis); // 2
  // proto:  Qt::MouseButton QMouseEvent::button();
extern void C_ZNK11QMouseEvent6buttonEv(void* qthis); // 2
  // proto:  void QMouseEvent::~QMouseEvent();
extern void C_ZN11QMouseEventD2Ev(void* qthis); // 4
  // proto:  QPoint QMouseEvent::pos();
extern void* C_ZNK11QMouseEvent3posEv(void* qthis); // 2
  // proto:  Qt::MouseButtons QMouseEvent::buttons();
extern void C_ZNK11QMouseEvent7buttonsEv(void* qthis); // 2
  // proto:  Qt::MouseEventSource QMouseEvent::source();
extern void C_ZNK11QMouseEvent6sourceEv(void* qthis); // 4
  // proto:  Qt::MouseEventFlags QMouseEvent::flags();
extern void C_ZNK11QMouseEvent5flagsEv(void* qthis); // 4
  // proto:  int QMouseEvent::globalY();
extern int32_t C_ZNK11QMouseEvent7globalYEv(void* qthis); // 2
  // proto:  int QMouseEvent::y();
extern void C_ZNK11QMouseEvent1yEv(void* qthis); // 2
  // proto:  int QMouseEvent::x();
extern void C_ZNK11QMouseEvent1xEv(void* qthis); // 2
  // proto:  const QPointF & QMouseEvent::windowPos();
extern void* C_ZNK11QMouseEvent9windowPosEv(void* qthis); // 2
  // proto:  int QMouseEvent::globalX();
extern int32_t C_ZNK11QMouseEvent7globalXEv(void* qthis); // 2
  // proto:  QUrl QFileOpenEvent::url();
extern void* C_ZNK14QFileOpenEvent3urlEv(void* qthis); // 2
  // proto:  void QFileOpenEvent::~QFileOpenEvent();
extern void C_ZN14QFileOpenEventD2Ev(void* qthis); // 4
  // proto:  QString QFileOpenEvent::file();
extern void* C_ZNK14QFileOpenEvent4fileEv(void* qthis); // 2
  // proto:  void QFileOpenEvent::QFileOpenEvent(const QString & file);
extern void* C_ZN14QFileOpenEventC2ERK7QString(void* arg0); // 3
  // proto:  void QFileOpenEvent::QFileOpenEvent(const QUrl & url);
extern void* C_ZN14QFileOpenEventC2ERK4QUrl(void* arg0); // 3
  // proto:  void QToolBarChangeEvent::~QToolBarChangeEvent();
extern void C_ZN19QToolBarChangeEventD2Ev(void* qthis); // 4
  // proto:  bool QToolBarChangeEvent::toggle();
extern bool C_ZNK19QToolBarChangeEvent6toggleEv(void* qthis); // 2
  // proto:  void QToolBarChangeEvent::QToolBarChangeEvent(bool t);
extern void* C_ZN19QToolBarChangeEventC2Eb(bool arg0); // 3
  // proto:  int QTabletEvent::xTilt();
extern int32_t C_ZNK12QTabletEvent5xTiltEv(void* qthis); // 2
  // proto:  QPoint QTabletEvent::pos();
extern void* C_ZNK12QTabletEvent3posEv(void* qthis); // 2
  // proto:  QTabletEvent::PointerType QTabletEvent::pointerType();
extern void C_ZNK12QTabletEvent11pointerTypeEv(void* qthis); // 2
  // proto:  Qt::MouseButtons QTabletEvent::buttons();
extern void C_ZNK12QTabletEvent7buttonsEv(void* qthis); // 4
  // proto:  qreal QTabletEvent::hiResGlobalX();
extern double C_ZNK12QTabletEvent12hiResGlobalXEv(void* qthis); // 2
  // proto:  qreal QTabletEvent::hiResGlobalY();
extern double C_ZNK12QTabletEvent12hiResGlobalYEv(void* qthis); // 2
  // proto:  void QTabletEvent::~QTabletEvent();
extern void C_ZN12QTabletEventD2Ev(void* qthis); // 4
  // proto:  qint64 QTabletEvent::uniqueId();
extern int64_t C_ZNK12QTabletEvent8uniqueIdEv(void* qthis); // 2
  // proto:  int QTabletEvent::globalX();
extern int32_t C_ZNK12QTabletEvent7globalXEv(void* qthis); // 2
  // proto:  int QTabletEvent::globalY();
extern int32_t C_ZNK12QTabletEvent7globalYEv(void* qthis); // 2
  // proto:  int QTabletEvent::yTilt();
extern int32_t C_ZNK12QTabletEvent5yTiltEv(void* qthis); // 2
  // proto:  const QPointF & QTabletEvent::globalPosF();
extern void* C_ZNK12QTabletEvent10globalPosFEv(void* qthis); // 2
  // proto:  qreal QTabletEvent::pressure();
extern double C_ZNK12QTabletEvent8pressureEv(void* qthis); // 2
  // proto:  QTabletEvent::TabletDevice QTabletEvent::device();
extern void C_ZNK12QTabletEvent6deviceEv(void* qthis); // 2
  // proto:  qreal QTabletEvent::rotation();
extern double C_ZNK12QTabletEvent8rotationEv(void* qthis); // 2
  // proto:  QPoint QTabletEvent::globalPos();
extern void* C_ZNK12QTabletEvent9globalPosEv(void* qthis); // 2
  // proto:  const QPointF & QTabletEvent::posF();
extern void* C_ZNK12QTabletEvent4posFEv(void* qthis); // 2
  // proto:  Qt::MouseButton QTabletEvent::button();
extern void C_ZNK12QTabletEvent6buttonEv(void* qthis); // 4
  // proto:  int QTabletEvent::y();
extern void C_ZNK12QTabletEvent1yEv(void* qthis); // 2
  // proto:  int QTabletEvent::x();
extern void C_ZNK12QTabletEvent1xEv(void* qthis); // 2
  // proto:  int QTabletEvent::z();
extern int32_t C_ZNK12QTabletEvent1zEv(void* qthis); // 2
  // proto:  qreal QTabletEvent::tangentialPressure();
extern double C_ZNK12QTabletEvent18tangentialPressureEv(void* qthis); // 2
  // proto:  Qt::TouchPointStates QTouchEvent::touchPointStates();
extern void C_ZNK11QTouchEvent16touchPointStatesEv(void* qthis); // 2
  // proto:  QObject * QTouchEvent::target();
extern void* C_ZNK11QTouchEvent6targetEv(void* qthis); // 2
  // proto:  void QTouchEvent::setTarget(QObject * atarget);
extern void C_ZN11QTouchEvent9setTargetEP7QObject(void* qthis, void* arg0); // 2
  // proto:  const QList<QTouchEvent::TouchPoint> & QTouchEvent::touchPoints();
extern void C_ZNK11QTouchEvent11touchPointsEv(void* qthis); // 2
  // proto:  void QTouchEvent::~QTouchEvent();
extern void C_ZN11QTouchEventD2Ev(void* qthis); // 4
  // proto:  QWindow * QTouchEvent::window();
extern void* C_ZNK11QTouchEvent6windowEv(void* qthis); // 2
  // proto:  void QTouchEvent::setDevice(QTouchDevice * adevice);
extern void C_ZN11QTouchEvent9setDeviceEP12QTouchDevice(void* qthis, void* arg0); // 2
  // proto:  QTouchDevice * QTouchEvent::device();
extern void* C_ZNK11QTouchEvent6deviceEv(void* qthis); // 2
  // proto:  void QTouchEvent::setWindow(QWindow * awindow);
extern void C_ZN11QTouchEvent9setWindowEP7QWindow(void* qthis, void* arg0); // 2
  // proto:  QScreen * QScreenOrientationChangeEvent::screen();
extern void* C_ZNK29QScreenOrientationChangeEvent6screenEv(void* qthis); // 4
  // proto:  Qt::ScreenOrientation QScreenOrientationChangeEvent::orientation();
extern void C_ZNK29QScreenOrientationChangeEvent11orientationEv(void* qthis); // 4
  // proto:  void QScreenOrientationChangeEvent::~QScreenOrientationChangeEvent();
extern void C_ZN29QScreenOrientationChangeEventD2Ev(void* qthis); // 4
  // proto:  void QIconDragEvent::~QIconDragEvent();
extern void C_ZN14QIconDragEventD2Ev(void* qthis); // 4
  // proto:  void QIconDragEvent::QIconDragEvent();
extern void* C_ZN14QIconDragEventC2Ev(); // 3
  // proto:  void QCloseEvent::QCloseEvent();
extern void* C_ZN11QCloseEventC2Ev(); // 3
  // proto:  void QCloseEvent::~QCloseEvent();
extern void C_ZN11QCloseEventD2Ev(void* qthis); // 4
  // proto:  void QDragEnterEvent::~QDragEnterEvent();
extern void C_ZN15QDragEnterEventD2Ev(void* qthis); // 4
  // proto:  QPoint QWheelEvent::pixelDelta();
extern void* C_ZNK11QWheelEvent10pixelDeltaEv(void* qthis); // 2
  // proto:  QPoint QWheelEvent::angleDelta();
extern void* C_ZNK11QWheelEvent10angleDeltaEv(void* qthis); // 2
  // proto:  void QWheelEvent::~QWheelEvent();
extern void C_ZN11QWheelEventD2Ev(void* qthis); // 4
  // proto:  Qt::Orientation QWheelEvent::orientation();
extern void C_ZNK11QWheelEvent11orientationEv(void* qthis); // 2
  // proto:  Qt::MouseEventSource QWheelEvent::source();
extern void C_ZNK11QWheelEvent6sourceEv(void* qthis); // 2
  // proto:  const QPointF & QWheelEvent::posF();
extern void* C_ZNK11QWheelEvent4posFEv(void* qthis); // 2
  // proto:  const QPointF & QWheelEvent::globalPosF();
extern void* C_ZNK11QWheelEvent10globalPosFEv(void* qthis); // 2
  // proto:  QPoint QWheelEvent::pos();
extern void* C_ZNK11QWheelEvent3posEv(void* qthis); // 2
  // proto:  bool QWheelEvent::inverted();
extern bool C_ZNK11QWheelEvent8invertedEv(void* qthis); // 2
  // proto:  Qt::MouseButtons QWheelEvent::buttons();
extern void C_ZNK11QWheelEvent7buttonsEv(void* qthis); // 2
  // proto:  int QWheelEvent::y();
extern void C_ZNK11QWheelEvent1yEv(void* qthis); // 2
  // proto:  int QWheelEvent::delta();
extern int32_t C_ZNK11QWheelEvent5deltaEv(void* qthis); // 2
  // proto:  Qt::ScrollPhase QWheelEvent::phase();
extern void C_ZNK11QWheelEvent5phaseEv(void* qthis); // 2
  // proto:  int QWheelEvent::globalY();
extern int32_t C_ZNK11QWheelEvent7globalYEv(void* qthis); // 2
  // proto:  int QWheelEvent::globalX();
extern int32_t C_ZNK11QWheelEvent7globalXEv(void* qthis); // 2
  // proto:  int QWheelEvent::x();
extern void C_ZNK11QWheelEvent1xEv(void* qthis); // 2
  // proto:  QPoint QWheelEvent::globalPos();
extern void* C_ZNK11QWheelEvent9globalPosEv(void* qthis); // 2
  // proto:  const QPointF & QHoverEvent::posF();
extern void* C_ZNK11QHoverEvent4posFEv(void* qthis); // 2
  // proto:  QPoint QHoverEvent::pos();
extern void* C_ZNK11QHoverEvent3posEv(void* qthis); // 2
  // proto:  QPoint QHoverEvent::oldPos();
extern void* C_ZNK11QHoverEvent6oldPosEv(void* qthis); // 2
  // proto:  const QPointF & QHoverEvent::oldPosF();
extern void* C_ZNK11QHoverEvent7oldPosFEv(void* qthis); // 2
  // proto:  void QHoverEvent::~QHoverEvent();
extern void C_ZN11QHoverEventD2Ev(void* qthis); // 4
  // proto:  void QDragMoveEvent::ignore();
extern void C_ZN14QDragMoveEvent6ignoreEv(void* qthis); // 2
  // proto:  void QDragMoveEvent::ignore(const QRect & r);
extern void C_ZN14QDragMoveEvent6ignoreERK5QRect(void* qthis, void* arg0); // 2
  // proto:  void QDragMoveEvent::~QDragMoveEvent();
extern void C_ZN14QDragMoveEventD2Ev(void* qthis); // 4
  // proto:  QRect QDragMoveEvent::answerRect();
extern void* C_ZNK14QDragMoveEvent10answerRectEv(void* qthis); // 2
  // proto:  void QDragMoveEvent::accept(const QRect & r);
extern void C_ZN14QDragMoveEvent6acceptERK5QRect(void* qthis, void* arg0); // 2
  // proto:  void QDragMoveEvent::accept();
extern void C_ZN14QDragMoveEvent6acceptEv(void* qthis); // 2
  // proto:  void QShowEvent::~QShowEvent();
extern void C_ZN10QShowEventD2Ev(void* qthis); // 4
  // proto:  void QShowEvent::QShowEvent();
extern void* C_ZN10QShowEventC2Ev(); // 3
  // proto:  QPlatformSurfaceEvent::SurfaceEventType QPlatformSurfaceEvent::surfaceEventType();
extern void C_ZNK21QPlatformSurfaceEvent16surfaceEventTypeEv(void* qthis); // 2
  // proto:  void QPlatformSurfaceEvent::~QPlatformSurfaceEvent();
extern void C_ZN21QPlatformSurfaceEventD2Ev(void* qthis); // 4
  // proto:  void QPaintEvent::QPaintEvent(const QRect & paintRect);
extern void* C_ZN11QPaintEventC2ERK5QRect(void* arg0); // 3
  // proto:  void QPaintEvent::QPaintEvent(const QRegion & paintRegion);
extern void* C_ZN11QPaintEventC2ERK7QRegion(void* arg0); // 3
  // proto:  void QPaintEvent::~QPaintEvent();
extern void C_ZN11QPaintEventD2Ev(void* qthis); // 4
  // proto:  const QRegion & QPaintEvent::region();
extern void* C_ZNK11QPaintEvent6regionEv(void* qthis); // 2
  // proto:  const QRect & QPaintEvent::rect();
extern void* C_ZNK11QPaintEvent4rectEv(void* qthis); // 2
  // proto:  Qt::FocusReason QFocusEvent::reason();
extern void C_ZNK11QFocusEvent6reasonEv(void* qthis); // 4
  // proto:  void QFocusEvent::~QFocusEvent();
extern void C_ZN11QFocusEventD2Ev(void* qthis); // 4
  // proto:  bool QFocusEvent::gotFocus();
extern bool C_ZNK11QFocusEvent8gotFocusEv(void* qthis); // 2
  // proto:  bool QFocusEvent::lostFocus();
extern bool C_ZNK11QFocusEvent9lostFocusEv(void* qthis); // 2
  // proto:  const QPointF & QNativeGestureEvent::localPos();
extern void* C_ZNK19QNativeGestureEvent8localPosEv(void* qthis); // 2
  // proto:  const QPointF & QNativeGestureEvent::screenPos();
extern void* C_ZNK19QNativeGestureEvent9screenPosEv(void* qthis); // 2
  // proto:  const QPoint QNativeGestureEvent::globalPos();
extern void* C_ZNK19QNativeGestureEvent9globalPosEv(void* qthis); // 2
  // proto:  const QPoint QNativeGestureEvent::pos();
extern void* C_ZNK19QNativeGestureEvent3posEv(void* qthis); // 2
  // proto:  qreal QNativeGestureEvent::value();
extern double C_ZNK19QNativeGestureEvent5valueEv(void* qthis); // 2
  // proto:  Qt::NativeGestureType QNativeGestureEvent::gestureType();
extern void C_ZNK19QNativeGestureEvent11gestureTypeEv(void* qthis); // 2
  // proto:  const QPointF & QNativeGestureEvent::windowPos();
extern void* C_ZNK19QNativeGestureEvent9windowPosEv(void* qthis); // 2
  // proto:  void QResizeEvent::~QResizeEvent();
extern void C_ZN12QResizeEventD2Ev(void* qthis); // 4
  // proto:  const QSize & QResizeEvent::oldSize();
extern void* C_ZNK12QResizeEvent7oldSizeEv(void* qthis); // 2
  // proto:  void QResizeEvent::QResizeEvent(const QSize & size, const QSize & oldSize);
extern void* C_ZN12QResizeEventC2ERK5QSizeS2_(void* arg0, void* arg1); // 3
  // proto:  const QSize & QResizeEvent::size();
extern void* C_ZNK12QResizeEvent4sizeEv(void* qthis); // 2
  // proto:  QString QStatusTipEvent::tip();
extern void* C_ZNK15QStatusTipEvent3tipEv(void* qthis); // 2
  // proto:  void QStatusTipEvent::~QStatusTipEvent();
extern void C_ZN15QStatusTipEventD2Ev(void* qthis); // 4
  // proto:  void QStatusTipEvent::QStatusTipEvent(const QString & tip);
extern void* C_ZN15QStatusTipEventC2ERK7QString(void* arg0); // 3
  // proto:  const QPointF & QEnterEvent::localPos();
extern void* C_ZNK11QEnterEvent8localPosEv(void* qthis); // 2
  // proto:  const QPointF & QEnterEvent::screenPos();
extern void* C_ZNK11QEnterEvent9screenPosEv(void* qthis); // 2
  // proto:  QPoint QEnterEvent::globalPos();
extern void* C_ZNK11QEnterEvent9globalPosEv(void* qthis); // 2
  // proto:  void QEnterEvent::QEnterEvent(const QPointF & localPos, const QPointF & windowPos, const QPointF & screenPos);
extern void* C_ZN11QEnterEventC2ERK7QPointFS2_S2_(void* arg0, void* arg1, void* arg2); // 3
  // proto:  QPoint QEnterEvent::pos();
extern void* C_ZNK11QEnterEvent3posEv(void* qthis); // 2
  // proto:  int QEnterEvent::globalX();
extern int32_t C_ZNK11QEnterEvent7globalXEv(void* qthis); // 2
  // proto:  void QEnterEvent::~QEnterEvent();
extern void C_ZN11QEnterEventD2Ev(void* qthis); // 4
  // proto:  int QEnterEvent::globalY();
extern int32_t C_ZNK11QEnterEvent7globalYEv(void* qthis); // 2
  // proto:  int QEnterEvent::y();
extern void C_ZNK11QEnterEvent1yEv(void* qthis); // 2
  // proto:  int QEnterEvent::x();
extern void C_ZNK11QEnterEvent1xEv(void* qthis); // 2
  // proto:  const QPointF & QEnterEvent::windowPos();
extern void* C_ZNK11QEnterEvent9windowPosEv(void* qthis); // 2
  // proto:  void QMoveEvent::QMoveEvent(const QPoint & pos, const QPoint & oldPos);
extern void* C_ZN10QMoveEventC2ERK6QPointS2_(void* arg0, void* arg1); // 3
  // proto:  void QMoveEvent::~QMoveEvent();
extern void C_ZN10QMoveEventD2Ev(void* qthis); // 4
  // proto:  const QPoint & QMoveEvent::oldPos();
extern void* C_ZNK10QMoveEvent6oldPosEv(void* qthis); // 2
  // proto:  const QPoint & QMoveEvent::pos();
extern void* C_ZNK10QMoveEvent3posEv(void* qthis); // 2
  // proto:  void QHideEvent::QHideEvent();
extern void* C_ZN10QHideEventC2Ev(); // 3
  // proto:  void QHideEvent::~QHideEvent();
extern void C_ZN10QHideEventD2Ev(void* qthis); // 4
  // proto:  void QDragLeaveEvent::~QDragLeaveEvent();
extern void C_ZN15QDragLeaveEventD2Ev(void* qthis); // 4
  // proto:  void QDragLeaveEvent::QDragLeaveEvent();
extern void* C_ZN15QDragLeaveEventC2Ev(); // 3
  // proto:  const QMimeData * QDropEvent::mimeData();
extern void* C_ZNK10QDropEvent8mimeDataEv(void* qthis); // 2
  // proto:  void QDropEvent::acceptProposedAction();
extern void C_ZN10QDropEvent20acceptProposedActionEv(void* qthis); // 2
  // proto:  Qt::DropActions QDropEvent::possibleActions();
extern void C_ZNK10QDropEvent15possibleActionsEv(void* qthis); // 2
  // proto:  const QPointF & QDropEvent::posF();
extern void* C_ZNK10QDropEvent4posFEv(void* qthis); // 2
  // proto:  QPoint QDropEvent::pos();
extern void* C_ZNK10QDropEvent3posEv(void* qthis); // 2
  // proto:  QObject * QDropEvent::source();
extern void* C_ZNK10QDropEvent6sourceEv(void* qthis); // 4
  // proto:  Qt::DropAction QDropEvent::proposedAction();
extern void C_ZNK10QDropEvent14proposedActionEv(void* qthis); // 2
  // proto:  Qt::DropAction QDropEvent::dropAction();
extern void C_ZNK10QDropEvent10dropActionEv(void* qthis); // 2
  // proto:  Qt::KeyboardModifiers QDropEvent::keyboardModifiers();
extern void C_ZNK10QDropEvent17keyboardModifiersEv(void* qthis); // 2
  // proto:  void QDropEvent::~QDropEvent();
extern void C_ZN10QDropEventD2Ev(void* qthis); // 4
  // proto:  Qt::MouseButtons QDropEvent::mouseButtons();
extern void C_ZNK10QDropEvent12mouseButtonsEv(void* qthis); // 2
  // proto:  Qt::KeyboardModifiers QInputEvent::modifiers();
extern void C_ZNK11QInputEvent9modifiersEv(void* qthis); // 2
  // proto:  ulong QInputEvent::timestamp();
extern int32_t C_ZNK11QInputEvent9timestampEv(void* qthis); // 2
  // proto:  void QInputEvent::setTimestamp(ulong atimestamp);
extern void C_ZN11QInputEvent12setTimestampEm(void* qthis, int32_t arg0); // 2
  // proto:  void QInputEvent::~QInputEvent();
extern void C_ZN11QInputEventD2Ev(void* qthis); // 4
  // proto:  Qt::ApplicationState QApplicationStateChangeEvent::applicationState();
extern void C_ZNK28QApplicationStateChangeEvent16applicationStateEv(void* qthis); // 4
  // proto:  QPointF QScrollEvent::contentPos();
extern void* C_ZNK12QScrollEvent10contentPosEv(void* qthis); // 4
  // proto:  QPointF QScrollEvent::overshootDistance();
extern void* C_ZNK12QScrollEvent17overshootDistanceEv(void* qthis); // 4
  // proto:  void QScrollEvent::~QScrollEvent();
extern void C_ZN12QScrollEventD2Ev(void* qthis); // 4
  // proto:  QScrollEvent::ScrollState QScrollEvent::scrollState();
extern void C_ZNK12QScrollEvent11scrollStateEv(void* qthis); // 4
  // proto:  int QKeyEvent::count();
extern int32_t C_ZNK9QKeyEvent5countEv(void* qthis); // 2
  // proto:  Qt::KeyboardModifiers QKeyEvent::modifiers();
extern void C_ZNK9QKeyEvent9modifiersEv(void* qthis); // 4
  // proto:  quint32 QKeyEvent::nativeModifiers();
extern int32_t C_ZNK9QKeyEvent15nativeModifiersEv(void* qthis); // 2
  // proto:  QString QKeyEvent::text();
extern void* C_ZNK9QKeyEvent4textEv(void* qthis); // 2
  // proto:  void QKeyEvent::~QKeyEvent();
extern void C_ZN9QKeyEventD2Ev(void* qthis); // 4
  // proto:  quint32 QKeyEvent::nativeScanCode();
extern int32_t C_ZNK9QKeyEvent14nativeScanCodeEv(void* qthis); // 2
  // proto:  bool QKeyEvent::isAutoRepeat();
extern bool C_ZNK9QKeyEvent12isAutoRepeatEv(void* qthis); // 2
  // proto:  int QKeyEvent::key();
extern int32_t C_ZNK9QKeyEvent3keyEv(void* qthis); // 2
  // proto:  quint32 QKeyEvent::nativeVirtualKey();
extern int32_t C_ZNK9QKeyEvent16nativeVirtualKeyEv(void* qthis); // 2
  // proto:  const QPoint & QContextMenuEvent::globalPos();
extern void* C_ZNK17QContextMenuEvent9globalPosEv(void* qthis); // 2
  // proto:  const QPoint & QContextMenuEvent::pos();
extern void* C_ZNK17QContextMenuEvent3posEv(void* qthis); // 2
  // proto:  int QContextMenuEvent::y();
extern void C_ZNK17QContextMenuEvent1yEv(void* qthis); // 2
  // proto:  QContextMenuEvent::Reason QContextMenuEvent::reason();
extern void C_ZNK17QContextMenuEvent6reasonEv(void* qthis); // 2
  // proto:  int QContextMenuEvent::x();
extern void C_ZNK17QContextMenuEvent1xEv(void* qthis); // 2
  // proto:  int QContextMenuEvent::globalX();
extern int32_t C_ZNK17QContextMenuEvent7globalXEv(void* qthis); // 2
  // proto:  int QContextMenuEvent::globalY();
extern int32_t C_ZNK17QContextMenuEvent7globalYEv(void* qthis); // 2
  // proto:  void QContextMenuEvent::~QContextMenuEvent();
extern void C_ZN17QContextMenuEventD2Ev(void* qthis); // 4
  // proto:  void QScrollPrepareEvent::setContentPosRange(const QRectF & rect);
extern void C_ZN19QScrollPrepareEvent18setContentPosRangeERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QPointF QScrollPrepareEvent::contentPos();
extern void* C_ZNK19QScrollPrepareEvent10contentPosEv(void* qthis); // 4
  // proto:  void QScrollPrepareEvent::setContentPos(const QPointF & pos);
extern void C_ZN19QScrollPrepareEvent13setContentPosERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QScrollPrepareEvent::~QScrollPrepareEvent();
extern void C_ZN19QScrollPrepareEventD2Ev(void* qthis); // 4
  // proto:  QSizeF QScrollPrepareEvent::viewportSize();
extern void* C_ZNK19QScrollPrepareEvent12viewportSizeEv(void* qthis); // 4
  // proto:  void QScrollPrepareEvent::setViewportSize(const QSizeF & size);
extern void C_ZN19QScrollPrepareEvent15setViewportSizeERK6QSizeF(void* qthis, void* arg0); // 4
  // proto:  void QScrollPrepareEvent::QScrollPrepareEvent(const QPointF & startPos);
extern void* C_ZN19QScrollPrepareEventC2ERK7QPointF(void* arg0); // 3
  // proto:  QRectF QScrollPrepareEvent::contentPosRange();
extern void* C_ZNK19QScrollPrepareEvent15contentPosRangeEv(void* qthis); // 4
  // proto:  QPointF QScrollPrepareEvent::startPos();
extern void* C_ZNK19QScrollPrepareEvent8startPosEv(void* qthis); // 4
  // proto:  void QShortcutEvent::~QShortcutEvent();
extern void C_ZN14QShortcutEventD2Ev(void* qthis); // 4
  // proto:  bool QShortcutEvent::isAmbiguous();
extern bool C_ZNK14QShortcutEvent11isAmbiguousEv(void* qthis); // 2
  // proto:  int QShortcutEvent::shortcutId();
extern int32_t C_ZNK14QShortcutEvent10shortcutIdEv(void* qthis); // 2
  // proto:  const QKeySequence & QShortcutEvent::key();
extern void* C_ZNK14QShortcutEvent3keyEv(void* qthis); // 2
  // proto:  void QShortcutEvent::QShortcutEvent(const QKeySequence & key, int id, bool ambiguous);
extern void* C_ZN14QShortcutEventC2ERK12QKeySequenceib(void* arg0, int32_t arg1, bool arg2); // 3
  // proto:  bool QWindowStateChangeEvent::isOverride();
extern bool C_ZNK23QWindowStateChangeEvent10isOverrideEv(void* qthis); // 4
  // proto:  Qt::WindowStates QWindowStateChangeEvent::oldState();
extern void C_ZNK23QWindowStateChangeEvent8oldStateEv(void* qthis); // 2
  // proto:  void QWindowStateChangeEvent::~QWindowStateChangeEvent();
extern void C_ZN23QWindowStateChangeEventD2Ev(void* qthis); // 4
  // proto:  void QInputMethodQueryEvent::~QInputMethodQueryEvent();
extern void C_ZN22QInputMethodQueryEventD2Ev(void* qthis); // 4
  // proto:  Qt::InputMethodQueries QInputMethodQueryEvent::queries();
extern void C_ZNK22QInputMethodQueryEvent7queriesEv(void* qthis); // 2
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QWhatsThisClickedEvent)=32
type QWhatsThisClickedEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QExposeEvent)=32
type QExposeEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QInputMethodEvent)=1
type QInputMethodEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QHelpEvent)=40
type QHelpEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QMouseEvent)=1
type QMouseEvent struct {
  /*qbase*/ QInputEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QFileOpenEvent)=40
type QFileOpenEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QToolBarChangeEvent)=24
type QToolBarChangeEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTabletEvent)=1
type QTabletEvent struct {
  /*qbase*/ QInputEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTouchEvent)=1
type QTouchEvent struct {
  /*qbase*/ QInputEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QScreenOrientationChangeEvent)=40
type QScreenOrientationChangeEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QIconDragEvent)=24
type QIconDragEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QCloseEvent)=24
type QCloseEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QDragEnterEvent)=1
type QDragEnterEvent struct {
  /*qbase*/ QDragMoveEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QWheelEvent)=1
type QWheelEvent struct {
  /*qbase*/ QInputEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QHoverEvent)=1
type QHoverEvent struct {
  /*qbase*/ QInputEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QDragMoveEvent)=1
type QDragMoveEvent struct {
  /*qbase*/ QDropEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QShowEvent)=24
type QShowEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QPlatformSurfaceEvent)=24
type QPlatformSurfaceEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QPaintEvent)=56
type QPaintEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QFocusEvent)=24
type QFocusEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QNativeGestureEvent)=1
type QNativeGestureEvent struct {
  /*qbase*/ QInputEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QResizeEvent)=40
type QResizeEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStatusTipEvent)=32
type QStatusTipEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QEnterEvent)=72
type QEnterEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QMoveEvent)=40
type QMoveEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QHideEvent)=24
type QHideEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QDragLeaveEvent)=24
type QDragLeaveEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QDropEvent)=1
type QDropEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QInputEvent)=1
type QInputEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QApplicationStateChangeEvent)=24
type QApplicationStateChangeEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QScrollEvent)=64
type QScrollEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QKeyEvent)=1
type QKeyEvent struct {
  /*qbase*/ QInputEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QContextMenuEvent)=1
type QContextMenuEvent struct {
  /*qbase*/ QInputEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QScrollPrepareEvent)=112
type QScrollPrepareEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QShortcutEvent)=40
type QShortcutEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QWindowStateChangeEvent)=1
type QWindowStateChangeEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QInputMethodQueryEvent)=1
type QInputMethodQueryEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// QWhatsThisClickedEvent(const class QString &)
func NewQWhatsThisClickedEvent(args ...interface{}) *QWhatsThisClickedEvent {
  // QWhatsThisClickedEvent(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QWhatsThisClickedEventC1ERK7QString
    // invoke: void QWhatsThisClickedEvent(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN22QWhatsThisClickedEventC2ERK7QString(arg0)
    return &QWhatsThisClickedEvent{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QWhatsThisClickedEvent", "QWhatsThisClickedEvent", args)
  }

  return nil // QWhatsThisClickedEvent{Qclsinst:qthis}
}

// ~QWhatsThisClickedEvent()
func (this *QWhatsThisClickedEvent) Freeqwhatsthisclickedevent(args ...interface{}) () {
  // ~QWhatsThisClickedEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QWhatsThisClickedEventD0Ev
    // invoke: void ~QWhatsThisClickedEvent()
    C.C_ZN22QWhatsThisClickedEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWhatsThisClickedEvent", "~QWhatsThisClickedEvent", args)
  }

  return
}

// href()
func (this *QWhatsThisClickedEvent) Href(args ...interface{}) (ret interface{}) {
  // href()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QWhatsThisClickedEvent4hrefEv
    // invoke: QString href()
    var ret0 = C.C_ZNK22QWhatsThisClickedEvent4hrefEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWhatsThisClickedEvent", "href", args)
  }

  return
}

// region()
func (this *QExposeEvent) Region(args ...interface{}) (ret interface{}) {
  // region()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QExposeEvent6regionEv
    // invoke: const QRegion & region()
    var ret0 = C.C_ZNK12QExposeEvent6regionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegion{}) // "const QRegion &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QExposeEvent", "region", args)
  }

  return
}

// ~QExposeEvent()
func (this *QExposeEvent) Freeqexposeevent(args ...interface{}) () {
  // ~QExposeEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QExposeEventD0Ev
    // invoke: void ~QExposeEvent()
    C.C_ZN12QExposeEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QExposeEvent", "~QExposeEvent", args)
  }

  return
}

// QExposeEvent(const class QRegion &)
func NewQExposeEvent(args ...interface{}) *QExposeEvent {
  // QExposeEvent(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QExposeEventC1ERK7QRegion
    // invoke: void QExposeEvent(const class QRegion &)
    var arg0 = args[0].(*QRegion).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QExposeEventC2ERK7QRegion(arg0)
    return &QExposeEvent{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QExposeEvent", "QExposeEvent", args)
  }

  return nil // QExposeEvent{Qclsinst:qthis}
}

// replacementStart()
func (this *QInputMethodEvent) Replacementstart(args ...interface{}) (ret interface{}) {
  // replacementStart()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QInputMethodEvent16replacementStartEv
    // invoke: int replacementStart()
    var ret0 = C.C_ZNK17QInputMethodEvent16replacementStartEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "replacementStart", args)
  }

  return
}

// replacementLength()
func (this *QInputMethodEvent) Replacementlength(args ...interface{}) (ret interface{}) {
  // replacementLength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QInputMethodEvent17replacementLengthEv
    // invoke: int replacementLength()
    var ret0 = C.C_ZNK17QInputMethodEvent17replacementLengthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "replacementLength", args)
  }

  return
}

// QInputMethodEvent(const class QInputMethodEvent &)
func NewQInputMethodEvent(args ...interface{}) *QInputMethodEvent {
  // QInputMethodEvent(const class QInputMethodEvent &)
  // QInputMethodEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QInputMethodEvent{}) // "const QInputMethodEvent &"
  vtys[1] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QInputMethodEventC1ERKS_
    // invoke: void QInputMethodEvent(const class QInputMethodEvent &)
    var arg0 = args[0].(*QInputMethodEvent).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QInputMethodEventC2ERKS_(arg0)
    return &QInputMethodEvent{Qclsinst:qthis}
  case 1:
    // invoke: _ZN17QInputMethodEventC1Ev
    // invoke: void QInputMethodEvent()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QInputMethodEventC2Ev()
    return &QInputMethodEvent{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "QInputMethodEvent", args)
  }

  return nil // QInputMethodEvent{Qclsinst:qthis}
}

// ~QInputMethodEvent()
func (this *QInputMethodEvent) Freeqinputmethodevent(args ...interface{}) () {
  // ~QInputMethodEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QInputMethodEventD0Ev
    // invoke: void ~QInputMethodEvent()
    C.C_ZN17QInputMethodEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "~QInputMethodEvent", args)
  }

  return
}

// preeditString()
func (this *QInputMethodEvent) Preeditstring(args ...interface{}) (ret interface{}) {
  // preeditString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QInputMethodEvent13preeditStringEv
    // invoke: const QString & preeditString()
    var ret0 = C.C_ZNK17QInputMethodEvent13preeditStringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "const QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "preeditString", args)
  }

  return
}

// attributes()
func (this *QInputMethodEvent) Attributes(args ...interface{}) () {
  // attributes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QInputMethodEvent10attributesEv
    // invoke: const QList<QInputMethodEvent::Attribute> & attributes()
    C.C_ZNK17QInputMethodEvent10attributesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "attributes", args)
  }

  return
}

// commitString()
func (this *QInputMethodEvent) Commitstring(args ...interface{}) (ret interface{}) {
  // commitString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QInputMethodEvent12commitStringEv
    // invoke: const QString & commitString()
    var ret0 = C.C_ZNK17QInputMethodEvent12commitStringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "const QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "commitString", args)
  }

  return
}

// setCommitString(const class QString &, int, int)
func (this *QInputMethodEvent) Setcommitstring(args ...interface{}) () {
  // setCommitString(const class QString &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QInputMethodEvent15setCommitStringERK7QStringii
    // invoke: void setCommitString(const class QString &, int, int)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN17QInputMethodEvent15setCommitStringERK7QStringii(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "setCommitString", args)
  }

  return
}

// globalPos()
func (this *QHelpEvent) Globalpos(args ...interface{}) (ret interface{}) {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent9globalPosEv
    // invoke: const QPoint & globalPos()
    var ret0 = C.C_ZNK10QHelpEvent9globalPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QHelpEvent", "globalPos", args)
  }

  return
}

// pos()
func (this *QHelpEvent) Pos(args ...interface{}) (ret interface{}) {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent3posEv
    // invoke: const QPoint & pos()
    var ret0 = C.C_ZNK10QHelpEvent3posEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QHelpEvent", "pos", args)
  }

  return
}

// y()
func (this *QHelpEvent) Y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent1yEv
    // invoke: int y()
    C.C_ZNK10QHelpEvent1yEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QHelpEvent", "y", args)
  }

  return
}

// x()
func (this *QHelpEvent) X(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent1xEv
    // invoke: int x()
    C.C_ZNK10QHelpEvent1xEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QHelpEvent", "x", args)
  }

  return
}

// globalX()
func (this *QHelpEvent) Globalx(args ...interface{}) (ret interface{}) {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent7globalXEv
    // invoke: int globalX()
    var ret0 = C.C_ZNK10QHelpEvent7globalXEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QHelpEvent", "globalX", args)
  }

  return
}

// globalY()
func (this *QHelpEvent) Globaly(args ...interface{}) (ret interface{}) {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent7globalYEv
    // invoke: int globalY()
    var ret0 = C.C_ZNK10QHelpEvent7globalYEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QHelpEvent", "globalY", args)
  }

  return
}

// ~QHelpEvent()
func (this *QHelpEvent) Freeqhelpevent(args ...interface{}) () {
  // ~QHelpEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QHelpEventD0Ev
    // invoke: void ~QHelpEvent()
    C.C_ZN10QHelpEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QHelpEvent", "~QHelpEvent", args)
  }

  return
}

// localPos()
func (this *QMouseEvent) Localpos(args ...interface{}) (ret interface{}) {
  // localPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent8localPosEv
    // invoke: const QPointF & localPos()
    var ret0 = C.C_ZNK11QMouseEvent8localPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMouseEvent", "localPos", args)
  }

  return
}

// screenPos()
func (this *QMouseEvent) Screenpos(args ...interface{}) (ret interface{}) {
  // screenPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent9screenPosEv
    // invoke: const QPointF & screenPos()
    var ret0 = C.C_ZNK11QMouseEvent9screenPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMouseEvent", "screenPos", args)
  }

  return
}

// globalPos()
func (this *QMouseEvent) Globalpos(args ...interface{}) (ret interface{}) {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent9globalPosEv
    // invoke: QPoint globalPos()
    var ret0 = C.C_ZNK11QMouseEvent9globalPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMouseEvent", "globalPos", args)
  }

  return
}

// button()
func (this *QMouseEvent) Button(args ...interface{}) () {
  // button()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent6buttonEv
    // invoke: Qt::MouseButton button()
    C.C_ZNK11QMouseEvent6buttonEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEvent", "button", args)
  }

  return
}

// ~QMouseEvent()
func (this *QMouseEvent) Freeqmouseevent(args ...interface{}) () {
  // ~QMouseEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMouseEventD0Ev
    // invoke: void ~QMouseEvent()
    C.C_ZN11QMouseEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEvent", "~QMouseEvent", args)
  }

  return
}

// pos()
func (this *QMouseEvent) Pos(args ...interface{}) (ret interface{}) {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent3posEv
    // invoke: QPoint pos()
    var ret0 = C.C_ZNK11QMouseEvent3posEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMouseEvent", "pos", args)
  }

  return
}

// buttons()
func (this *QMouseEvent) Buttons(args ...interface{}) () {
  // buttons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent7buttonsEv
    // invoke: Qt::MouseButtons buttons()
    C.C_ZNK11QMouseEvent7buttonsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEvent", "buttons", args)
  }

  return
}

// source()
func (this *QMouseEvent) Source(args ...interface{}) () {
  // source()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent6sourceEv
    // invoke: Qt::MouseEventSource source()
    C.C_ZNK11QMouseEvent6sourceEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEvent", "source", args)
  }

  return
}

// flags()
func (this *QMouseEvent) Flags(args ...interface{}) () {
  // flags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent5flagsEv
    // invoke: Qt::MouseEventFlags flags()
    C.C_ZNK11QMouseEvent5flagsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEvent", "flags", args)
  }

  return
}

// globalY()
func (this *QMouseEvent) Globaly(args ...interface{}) (ret interface{}) {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent7globalYEv
    // invoke: int globalY()
    var ret0 = C.C_ZNK11QMouseEvent7globalYEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMouseEvent", "globalY", args)
  }

  return
}

// y()
func (this *QMouseEvent) Y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent1yEv
    // invoke: int y()
    C.C_ZNK11QMouseEvent1yEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEvent", "y", args)
  }

  return
}

// x()
func (this *QMouseEvent) X(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent1xEv
    // invoke: int x()
    C.C_ZNK11QMouseEvent1xEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEvent", "x", args)
  }

  return
}

// windowPos()
func (this *QMouseEvent) Windowpos(args ...interface{}) (ret interface{}) {
  // windowPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent9windowPosEv
    // invoke: const QPointF & windowPos()
    var ret0 = C.C_ZNK11QMouseEvent9windowPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMouseEvent", "windowPos", args)
  }

  return
}

// globalX()
func (this *QMouseEvent) Globalx(args ...interface{}) (ret interface{}) {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent7globalXEv
    // invoke: int globalX()
    var ret0 = C.C_ZNK11QMouseEvent7globalXEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMouseEvent", "globalX", args)
  }

  return
}

// url()
func (this *QFileOpenEvent) Url(args ...interface{}) (ret interface{}) {
  // url()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QFileOpenEvent3urlEv
    // invoke: QUrl url()
    var ret0 = C.C_ZNK14QFileOpenEvent3urlEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QUrl{}) // "QUrl"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileOpenEvent", "url", args)
  }

  return
}

// ~QFileOpenEvent()
func (this *QFileOpenEvent) Freeqfileopenevent(args ...interface{}) () {
  // ~QFileOpenEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QFileOpenEventD0Ev
    // invoke: void ~QFileOpenEvent()
    C.C_ZN14QFileOpenEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileOpenEvent", "~QFileOpenEvent", args)
  }

  return
}

// file()
func (this *QFileOpenEvent) File(args ...interface{}) (ret interface{}) {
  // file()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QFileOpenEvent4fileEv
    // invoke: QString file()
    var ret0 = C.C_ZNK14QFileOpenEvent4fileEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileOpenEvent", "file", args)
  }

  return
}

// QFileOpenEvent(const class QString &)
func NewQFileOpenEvent(args ...interface{}) *QFileOpenEvent {
  // QFileOpenEvent(const class QString &)
  // QFileOpenEvent(const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QUrl{}) // "const QUrl &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QFileOpenEventC1ERK7QString
    // invoke: void QFileOpenEvent(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QFileOpenEventC2ERK7QString(arg0)
    return &QFileOpenEvent{Qclsinst:qthis}
  case 1:
    // invoke: _ZN14QFileOpenEventC1ERK4QUrl
    // invoke: void QFileOpenEvent(const class QUrl &)
    var arg0 = args[0].(*qtcore.QUrl).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QFileOpenEventC2ERK4QUrl(arg0)
    return &QFileOpenEvent{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QFileOpenEvent", "QFileOpenEvent", args)
  }

  return nil // QFileOpenEvent{Qclsinst:qthis}
}

// ~QToolBarChangeEvent()
func (this *QToolBarChangeEvent) Freeqtoolbarchangeevent(args ...interface{}) () {
  // ~QToolBarChangeEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QToolBarChangeEventD0Ev
    // invoke: void ~QToolBarChangeEvent()
    C.C_ZN19QToolBarChangeEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QToolBarChangeEvent", "~QToolBarChangeEvent", args)
  }

  return
}

// toggle()
func (this *QToolBarChangeEvent) Toggle(args ...interface{}) (ret interface{}) {
  // toggle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QToolBarChangeEvent6toggleEv
    // invoke: bool toggle()
    var ret0 = C.C_ZNK19QToolBarChangeEvent6toggleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBarChangeEvent", "toggle", args)
  }

  return
}

// QToolBarChangeEvent(_Bool)
func NewQToolBarChangeEvent(args ...interface{}) *QToolBarChangeEvent {
  // QToolBarChangeEvent(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QToolBarChangeEventC1Eb
    // invoke: void QToolBarChangeEvent(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QToolBarChangeEventC2Eb(arg0)
    return &QToolBarChangeEvent{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QToolBarChangeEvent", "QToolBarChangeEvent", args)
  }

  return nil // QToolBarChangeEvent{Qclsinst:qthis}
}

// xTilt()
func (this *QTabletEvent) Xtilt(args ...interface{}) (ret interface{}) {
  // xTilt()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent5xTiltEv
    // invoke: int xTilt()
    var ret0 = C.C_ZNK12QTabletEvent5xTiltEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabletEvent", "xTilt", args)
  }

  return
}

// pos()
func (this *QTabletEvent) Pos(args ...interface{}) (ret interface{}) {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent3posEv
    // invoke: QPoint pos()
    var ret0 = C.C_ZNK12QTabletEvent3posEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabletEvent", "pos", args)
  }

  return
}

// pointerType()
func (this *QTabletEvent) Pointertype(args ...interface{}) () {
  // pointerType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent11pointerTypeEv
    // invoke: QTabletEvent::PointerType pointerType()
    C.C_ZNK12QTabletEvent11pointerTypeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "pointerType", args)
  }

  return
}

// buttons()
func (this *QTabletEvent) Buttons(args ...interface{}) () {
  // buttons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent7buttonsEv
    // invoke: Qt::MouseButtons buttons()
    C.C_ZNK12QTabletEvent7buttonsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "buttons", args)
  }

  return
}

// hiResGlobalX()
func (this *QTabletEvent) Hiresglobalx(args ...interface{}) (ret interface{}) {
  // hiResGlobalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent12hiResGlobalXEv
    // invoke: qreal hiResGlobalX()
    var ret0 = C.C_ZNK12QTabletEvent12hiResGlobalXEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabletEvent", "hiResGlobalX", args)
  }

  return
}

// hiResGlobalY()
func (this *QTabletEvent) Hiresglobaly(args ...interface{}) (ret interface{}) {
  // hiResGlobalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent12hiResGlobalYEv
    // invoke: qreal hiResGlobalY()
    var ret0 = C.C_ZNK12QTabletEvent12hiResGlobalYEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabletEvent", "hiResGlobalY", args)
  }

  return
}

// ~QTabletEvent()
func (this *QTabletEvent) Freeqtabletevent(args ...interface{}) () {
  // ~QTabletEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTabletEventD0Ev
    // invoke: void ~QTabletEvent()
    C.C_ZN12QTabletEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "~QTabletEvent", args)
  }

  return
}

// uniqueId()
func (this *QTabletEvent) Uniqueid(args ...interface{}) (ret interface{}) {
  // uniqueId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent8uniqueIdEv
    // invoke: qint64 uniqueId()
    var ret0 = C.C_ZNK12QTabletEvent8uniqueIdEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabletEvent", "uniqueId", args)
  }

  return
}

// globalX()
func (this *QTabletEvent) Globalx(args ...interface{}) (ret interface{}) {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent7globalXEv
    // invoke: int globalX()
    var ret0 = C.C_ZNK12QTabletEvent7globalXEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabletEvent", "globalX", args)
  }

  return
}

// globalY()
func (this *QTabletEvent) Globaly(args ...interface{}) (ret interface{}) {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent7globalYEv
    // invoke: int globalY()
    var ret0 = C.C_ZNK12QTabletEvent7globalYEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabletEvent", "globalY", args)
  }

  return
}

// yTilt()
func (this *QTabletEvent) Ytilt(args ...interface{}) (ret interface{}) {
  // yTilt()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent5yTiltEv
    // invoke: int yTilt()
    var ret0 = C.C_ZNK12QTabletEvent5yTiltEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabletEvent", "yTilt", args)
  }

  return
}

// globalPosF()
func (this *QTabletEvent) Globalposf(args ...interface{}) (ret interface{}) {
  // globalPosF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent10globalPosFEv
    // invoke: const QPointF & globalPosF()
    var ret0 = C.C_ZNK12QTabletEvent10globalPosFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabletEvent", "globalPosF", args)
  }

  return
}

// pressure()
func (this *QTabletEvent) Pressure(args ...interface{}) (ret interface{}) {
  // pressure()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent8pressureEv
    // invoke: qreal pressure()
    var ret0 = C.C_ZNK12QTabletEvent8pressureEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabletEvent", "pressure", args)
  }

  return
}

// device()
func (this *QTabletEvent) Device(args ...interface{}) () {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent6deviceEv
    // invoke: QTabletEvent::TabletDevice device()
    C.C_ZNK12QTabletEvent6deviceEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "device", args)
  }

  return
}

// rotation()
func (this *QTabletEvent) Rotation(args ...interface{}) (ret interface{}) {
  // rotation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent8rotationEv
    // invoke: qreal rotation()
    var ret0 = C.C_ZNK12QTabletEvent8rotationEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabletEvent", "rotation", args)
  }

  return
}

// globalPos()
func (this *QTabletEvent) Globalpos(args ...interface{}) (ret interface{}) {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent9globalPosEv
    // invoke: QPoint globalPos()
    var ret0 = C.C_ZNK12QTabletEvent9globalPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabletEvent", "globalPos", args)
  }

  return
}

// posF()
func (this *QTabletEvent) Posf(args ...interface{}) (ret interface{}) {
  // posF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent4posFEv
    // invoke: const QPointF & posF()
    var ret0 = C.C_ZNK12QTabletEvent4posFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabletEvent", "posF", args)
  }

  return
}

// button()
func (this *QTabletEvent) Button(args ...interface{}) () {
  // button()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent6buttonEv
    // invoke: Qt::MouseButton button()
    C.C_ZNK12QTabletEvent6buttonEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "button", args)
  }

  return
}

// y()
func (this *QTabletEvent) Y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent1yEv
    // invoke: int y()
    C.C_ZNK12QTabletEvent1yEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "y", args)
  }

  return
}

// x()
func (this *QTabletEvent) X(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent1xEv
    // invoke: int x()
    C.C_ZNK12QTabletEvent1xEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "x", args)
  }

  return
}

// z()
func (this *QTabletEvent) Z(args ...interface{}) (ret interface{}) {
  // z()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent1zEv
    // invoke: int z()
    var ret0 = C.C_ZNK12QTabletEvent1zEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabletEvent", "z", args)
  }

  return
}

// tangentialPressure()
func (this *QTabletEvent) Tangentialpressure(args ...interface{}) (ret interface{}) {
  // tangentialPressure()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent18tangentialPressureEv
    // invoke: qreal tangentialPressure()
    var ret0 = C.C_ZNK12QTabletEvent18tangentialPressureEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabletEvent", "tangentialPressure", args)
  }

  return
}

// touchPointStates()
func (this *QTouchEvent) Touchpointstates(args ...interface{}) () {
  // touchPointStates()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTouchEvent16touchPointStatesEv
    // invoke: Qt::TouchPointStates touchPointStates()
    C.C_ZNK11QTouchEvent16touchPointStatesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTouchEvent", "touchPointStates", args)
  }

  return
}

// target()
func (this *QTouchEvent) Target(args ...interface{}) (ret interface{}) {
  // target()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTouchEvent6targetEv
    // invoke: QObject * target()
    var ret0 = C.C_ZNK11QTouchEvent6targetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTouchEvent", "target", args)
  }

  return
}

// setTarget(class QObject *)
func (this *QTouchEvent) Settarget(args ...interface{}) () {
  // setTarget(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTouchEvent9setTargetEP7QObject
    // invoke: void setTarget(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTouchEvent9setTargetEP7QObject(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTouchEvent", "setTarget", args)
  }

  return
}

// touchPoints()
func (this *QTouchEvent) Touchpoints(args ...interface{}) () {
  // touchPoints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTouchEvent11touchPointsEv
    // invoke: const QList<QTouchEvent::TouchPoint> & touchPoints()
    C.C_ZNK11QTouchEvent11touchPointsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTouchEvent", "touchPoints", args)
  }

  return
}

// ~QTouchEvent()
func (this *QTouchEvent) Freeqtouchevent(args ...interface{}) () {
  // ~QTouchEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTouchEventD0Ev
    // invoke: void ~QTouchEvent()
    C.C_ZN11QTouchEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTouchEvent", "~QTouchEvent", args)
  }

  return
}

// window()
func (this *QTouchEvent) Window(args ...interface{}) (ret interface{}) {
  // window()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTouchEvent6windowEv
    // invoke: QWindow * window()
    var ret0 = C.C_ZNK11QTouchEvent6windowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWindow{}) // "QWindow *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTouchEvent", "window", args)
  }

  return
}

// setDevice(class QTouchDevice *)
func (this *QTouchEvent) Setdevice(args ...interface{}) () {
  // setDevice(class QTouchDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTouchDevice{}) // "QTouchDevice *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTouchEvent9setDeviceEP12QTouchDevice
    // invoke: void setDevice(class QTouchDevice *)
    var arg0 = args[0].(*QTouchDevice).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTouchEvent9setDeviceEP12QTouchDevice(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTouchEvent", "setDevice", args)
  }

  return
}

// device()
func (this *QTouchEvent) Device(args ...interface{}) (ret interface{}) {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTouchEvent6deviceEv
    // invoke: QTouchDevice * device()
    var ret0 = C.C_ZNK11QTouchEvent6deviceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTouchDevice{}) // "QTouchDevice *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTouchEvent", "device", args)
  }

  return
}

// setWindow(class QWindow *)
func (this *QTouchEvent) Setwindow(args ...interface{}) () {
  // setWindow(class QWindow *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWindow{}) // "QWindow *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTouchEvent9setWindowEP7QWindow
    // invoke: void setWindow(class QWindow *)
    var arg0 = args[0].(*QWindow).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTouchEvent9setWindowEP7QWindow(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTouchEvent", "setWindow", args)
  }

  return
}

// screen()
func (this *QScreenOrientationChangeEvent) Screen(args ...interface{}) (ret interface{}) {
  // screen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK29QScreenOrientationChangeEvent6screenEv
    // invoke: QScreen * screen()
    var ret0 = C.C_ZNK29QScreenOrientationChangeEvent6screenEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QScreen{}) // "QScreen *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScreenOrientationChangeEvent", "screen", args)
  }

  return
}

// orientation()
func (this *QScreenOrientationChangeEvent) Orientation(args ...interface{}) () {
  // orientation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK29QScreenOrientationChangeEvent11orientationEv
    // invoke: Qt::ScreenOrientation orientation()
    C.C_ZNK29QScreenOrientationChangeEvent11orientationEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QScreenOrientationChangeEvent", "orientation", args)
  }

  return
}

// ~QScreenOrientationChangeEvent()
func (this *QScreenOrientationChangeEvent) Freeqscreenorientationchangeevent(args ...interface{}) () {
  // ~QScreenOrientationChangeEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN29QScreenOrientationChangeEventD0Ev
    // invoke: void ~QScreenOrientationChangeEvent()
    C.C_ZN29QScreenOrientationChangeEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QScreenOrientationChangeEvent", "~QScreenOrientationChangeEvent", args)
  }

  return
}

// ~QIconDragEvent()
func (this *QIconDragEvent) Freeqicondragevent(args ...interface{}) () {
  // ~QIconDragEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QIconDragEventD0Ev
    // invoke: void ~QIconDragEvent()
    C.C_ZN14QIconDragEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QIconDragEvent", "~QIconDragEvent", args)
  }

  return
}

// QIconDragEvent()
func NewQIconDragEvent(args ...interface{}) *QIconDragEvent {
  // QIconDragEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QIconDragEventC1Ev
    // invoke: void QIconDragEvent()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QIconDragEventC2Ev()
    return &QIconDragEvent{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QIconDragEvent", "QIconDragEvent", args)
  }

  return nil // QIconDragEvent{Qclsinst:qthis}
}

// QCloseEvent()
func NewQCloseEvent(args ...interface{}) *QCloseEvent {
  // QCloseEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QCloseEventC1Ev
    // invoke: void QCloseEvent()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QCloseEventC2Ev()
    return &QCloseEvent{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QCloseEvent", "QCloseEvent", args)
  }

  return nil // QCloseEvent{Qclsinst:qthis}
}

// ~QCloseEvent()
func (this *QCloseEvent) Freeqcloseevent(args ...interface{}) () {
  // ~QCloseEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QCloseEventD0Ev
    // invoke: void ~QCloseEvent()
    C.C_ZN11QCloseEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCloseEvent", "~QCloseEvent", args)
  }

  return
}

// ~QDragEnterEvent()
func (this *QDragEnterEvent) Freeqdragenterevent(args ...interface{}) () {
  // ~QDragEnterEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QDragEnterEventD0Ev
    // invoke: void ~QDragEnterEvent()
    C.C_ZN15QDragEnterEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDragEnterEvent", "~QDragEnterEvent", args)
  }

  return
}

// pixelDelta()
func (this *QWheelEvent) Pixeldelta(args ...interface{}) (ret interface{}) {
  // pixelDelta()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent10pixelDeltaEv
    // invoke: QPoint pixelDelta()
    var ret0 = C.C_ZNK11QWheelEvent10pixelDeltaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWheelEvent", "pixelDelta", args)
  }

  return
}

// angleDelta()
func (this *QWheelEvent) Angledelta(args ...interface{}) (ret interface{}) {
  // angleDelta()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent10angleDeltaEv
    // invoke: QPoint angleDelta()
    var ret0 = C.C_ZNK11QWheelEvent10angleDeltaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWheelEvent", "angleDelta", args)
  }

  return
}

// ~QWheelEvent()
func (this *QWheelEvent) Freeqwheelevent(args ...interface{}) () {
  // ~QWheelEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QWheelEventD0Ev
    // invoke: void ~QWheelEvent()
    C.C_ZN11QWheelEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "~QWheelEvent", args)
  }

  return
}

// orientation()
func (this *QWheelEvent) Orientation(args ...interface{}) () {
  // orientation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent11orientationEv
    // invoke: Qt::Orientation orientation()
    C.C_ZNK11QWheelEvent11orientationEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "orientation", args)
  }

  return
}

// source()
func (this *QWheelEvent) Source(args ...interface{}) () {
  // source()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent6sourceEv
    // invoke: Qt::MouseEventSource source()
    C.C_ZNK11QWheelEvent6sourceEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "source", args)
  }

  return
}

// posF()
func (this *QWheelEvent) Posf(args ...interface{}) (ret interface{}) {
  // posF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent4posFEv
    // invoke: const QPointF & posF()
    var ret0 = C.C_ZNK11QWheelEvent4posFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWheelEvent", "posF", args)
  }

  return
}

// globalPosF()
func (this *QWheelEvent) Globalposf(args ...interface{}) (ret interface{}) {
  // globalPosF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent10globalPosFEv
    // invoke: const QPointF & globalPosF()
    var ret0 = C.C_ZNK11QWheelEvent10globalPosFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWheelEvent", "globalPosF", args)
  }

  return
}

// pos()
func (this *QWheelEvent) Pos(args ...interface{}) (ret interface{}) {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent3posEv
    // invoke: QPoint pos()
    var ret0 = C.C_ZNK11QWheelEvent3posEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWheelEvent", "pos", args)
  }

  return
}

// inverted()
func (this *QWheelEvent) Inverted(args ...interface{}) (ret interface{}) {
  // inverted()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent8invertedEv
    // invoke: bool inverted()
    var ret0 = C.C_ZNK11QWheelEvent8invertedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWheelEvent", "inverted", args)
  }

  return
}

// buttons()
func (this *QWheelEvent) Buttons(args ...interface{}) () {
  // buttons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent7buttonsEv
    // invoke: Qt::MouseButtons buttons()
    C.C_ZNK11QWheelEvent7buttonsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "buttons", args)
  }

  return
}

// y()
func (this *QWheelEvent) Y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent1yEv
    // invoke: int y()
    C.C_ZNK11QWheelEvent1yEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "y", args)
  }

  return
}

// delta()
func (this *QWheelEvent) Delta(args ...interface{}) (ret interface{}) {
  // delta()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent5deltaEv
    // invoke: int delta()
    var ret0 = C.C_ZNK11QWheelEvent5deltaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWheelEvent", "delta", args)
  }

  return
}

// phase()
func (this *QWheelEvent) Phase(args ...interface{}) () {
  // phase()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent5phaseEv
    // invoke: Qt::ScrollPhase phase()
    C.C_ZNK11QWheelEvent5phaseEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "phase", args)
  }

  return
}

// globalY()
func (this *QWheelEvent) Globaly(args ...interface{}) (ret interface{}) {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent7globalYEv
    // invoke: int globalY()
    var ret0 = C.C_ZNK11QWheelEvent7globalYEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWheelEvent", "globalY", args)
  }

  return
}

// globalX()
func (this *QWheelEvent) Globalx(args ...interface{}) (ret interface{}) {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent7globalXEv
    // invoke: int globalX()
    var ret0 = C.C_ZNK11QWheelEvent7globalXEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWheelEvent", "globalX", args)
  }

  return
}

// x()
func (this *QWheelEvent) X(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent1xEv
    // invoke: int x()
    C.C_ZNK11QWheelEvent1xEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "x", args)
  }

  return
}

// globalPos()
func (this *QWheelEvent) Globalpos(args ...interface{}) (ret interface{}) {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent9globalPosEv
    // invoke: QPoint globalPos()
    var ret0 = C.C_ZNK11QWheelEvent9globalPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWheelEvent", "globalPos", args)
  }

  return
}

// posF()
func (this *QHoverEvent) Posf(args ...interface{}) (ret interface{}) {
  // posF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHoverEvent4posFEv
    // invoke: const QPointF & posF()
    var ret0 = C.C_ZNK11QHoverEvent4posFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QHoverEvent", "posF", args)
  }

  return
}

// pos()
func (this *QHoverEvent) Pos(args ...interface{}) (ret interface{}) {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHoverEvent3posEv
    // invoke: QPoint pos()
    var ret0 = C.C_ZNK11QHoverEvent3posEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QHoverEvent", "pos", args)
  }

  return
}

// oldPos()
func (this *QHoverEvent) Oldpos(args ...interface{}) (ret interface{}) {
  // oldPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHoverEvent6oldPosEv
    // invoke: QPoint oldPos()
    var ret0 = C.C_ZNK11QHoverEvent6oldPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QHoverEvent", "oldPos", args)
  }

  return
}

// oldPosF()
func (this *QHoverEvent) Oldposf(args ...interface{}) (ret interface{}) {
  // oldPosF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHoverEvent7oldPosFEv
    // invoke: const QPointF & oldPosF()
    var ret0 = C.C_ZNK11QHoverEvent7oldPosFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QHoverEvent", "oldPosF", args)
  }

  return
}

// ~QHoverEvent()
func (this *QHoverEvent) Freeqhoverevent(args ...interface{}) () {
  // ~QHoverEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHoverEventD0Ev
    // invoke: void ~QHoverEvent()
    C.C_ZN11QHoverEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QHoverEvent", "~QHoverEvent", args)
  }

  return
}

// ignore()
func (this *QDragMoveEvent) Ignore(args ...interface{}) () {
  // ignore()
  // ignore(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDragMoveEvent6ignoreEv
    // invoke: void ignore()
    C.C_ZN14QDragMoveEvent6ignoreEv(this.Qclsinst)
  case 1:
    // invoke: _ZN14QDragMoveEvent6ignoreERK5QRect
    // invoke: void ignore(const class QRect &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QDragMoveEvent6ignoreERK5QRect(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDragMoveEvent", "ignore", args)
  }

  return
}

// ~QDragMoveEvent()
func (this *QDragMoveEvent) Freeqdragmoveevent(args ...interface{}) () {
  // ~QDragMoveEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDragMoveEventD0Ev
    // invoke: void ~QDragMoveEvent()
    C.C_ZN14QDragMoveEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDragMoveEvent", "~QDragMoveEvent", args)
  }

  return
}

// answerRect()
func (this *QDragMoveEvent) Answerrect(args ...interface{}) (ret interface{}) {
  // answerRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDragMoveEvent10answerRectEv
    // invoke: QRect answerRect()
    var ret0 = C.C_ZNK14QDragMoveEvent10answerRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDragMoveEvent", "answerRect", args)
  }

  return
}

// accept(const class QRect &)
func (this *QDragMoveEvent) Accept(args ...interface{}) () {
  // accept(const class QRect &)
  // accept()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDragMoveEvent6acceptERK5QRect
    // invoke: void accept(const class QRect &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QDragMoveEvent6acceptERK5QRect(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN14QDragMoveEvent6acceptEv
    // invoke: void accept()
    C.C_ZN14QDragMoveEvent6acceptEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDragMoveEvent", "accept", args)
  }

  return
}

// ~QShowEvent()
func (this *QShowEvent) Freeqshowevent(args ...interface{}) () {
  // ~QShowEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QShowEventD0Ev
    // invoke: void ~QShowEvent()
    C.C_ZN10QShowEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QShowEvent", "~QShowEvent", args)
  }

  return
}

// QShowEvent()
func NewQShowEvent(args ...interface{}) *QShowEvent {
  // QShowEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QShowEventC1Ev
    // invoke: void QShowEvent()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QShowEventC2Ev()
    return &QShowEvent{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QShowEvent", "QShowEvent", args)
  }

  return nil // QShowEvent{Qclsinst:qthis}
}

// surfaceEventType()
func (this *QPlatformSurfaceEvent) Surfaceeventtype(args ...interface{}) () {
  // surfaceEventType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPlatformSurfaceEvent16surfaceEventTypeEv
    // invoke: QPlatformSurfaceEvent::SurfaceEventType surfaceEventType()
    C.C_ZNK21QPlatformSurfaceEvent16surfaceEventTypeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPlatformSurfaceEvent", "surfaceEventType", args)
  }

  return
}

// ~QPlatformSurfaceEvent()
func (this *QPlatformSurfaceEvent) Freeqplatformsurfaceevent(args ...interface{}) () {
  // ~QPlatformSurfaceEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QPlatformSurfaceEventD0Ev
    // invoke: void ~QPlatformSurfaceEvent()
    C.C_ZN21QPlatformSurfaceEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPlatformSurfaceEvent", "~QPlatformSurfaceEvent", args)
  }

  return
}

// QPaintEvent(const class QRect &)
func NewQPaintEvent(args ...interface{}) *QPaintEvent {
  // QPaintEvent(const class QRect &)
  // QPaintEvent(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPaintEventC1ERK5QRect
    // invoke: void QPaintEvent(const class QRect &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QPaintEventC2ERK5QRect(arg0)
    return &QPaintEvent{Qclsinst:qthis}
  case 1:
    // invoke: _ZN11QPaintEventC1ERK7QRegion
    // invoke: void QPaintEvent(const class QRegion &)
    var arg0 = args[0].(*QRegion).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QPaintEventC2ERK7QRegion(arg0)
    return &QPaintEvent{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPaintEvent", "QPaintEvent", args)
  }

  return nil // QPaintEvent{Qclsinst:qthis}
}

// ~QPaintEvent()
func (this *QPaintEvent) Freeqpaintevent(args ...interface{}) () {
  // ~QPaintEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPaintEventD0Ev
    // invoke: void ~QPaintEvent()
    C.C_ZN11QPaintEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEvent", "~QPaintEvent", args)
  }

  return
}

// region()
func (this *QPaintEvent) Region(args ...interface{}) (ret interface{}) {
  // region()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPaintEvent6regionEv
    // invoke: const QRegion & region()
    var ret0 = C.C_ZNK11QPaintEvent6regionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegion{}) // "const QRegion &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEvent", "region", args)
  }

  return
}

// rect()
func (this *QPaintEvent) Rect(args ...interface{}) (ret interface{}) {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPaintEvent4rectEv
    // invoke: const QRect & rect()
    var ret0 = C.C_ZNK11QPaintEvent4rectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEvent", "rect", args)
  }

  return
}

// reason()
func (this *QFocusEvent) Reason(args ...interface{}) () {
  // reason()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFocusEvent6reasonEv
    // invoke: Qt::FocusReason reason()
    C.C_ZNK11QFocusEvent6reasonEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFocusEvent", "reason", args)
  }

  return
}

// ~QFocusEvent()
func (this *QFocusEvent) Freeqfocusevent(args ...interface{}) () {
  // ~QFocusEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFocusEventD0Ev
    // invoke: void ~QFocusEvent()
    C.C_ZN11QFocusEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFocusEvent", "~QFocusEvent", args)
  }

  return
}

// gotFocus()
func (this *QFocusEvent) Gotfocus(args ...interface{}) (ret interface{}) {
  // gotFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFocusEvent8gotFocusEv
    // invoke: bool gotFocus()
    var ret0 = C.C_ZNK11QFocusEvent8gotFocusEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFocusEvent", "gotFocus", args)
  }

  return
}

// lostFocus()
func (this *QFocusEvent) Lostfocus(args ...interface{}) (ret interface{}) {
  // lostFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFocusEvent9lostFocusEv
    // invoke: bool lostFocus()
    var ret0 = C.C_ZNK11QFocusEvent9lostFocusEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFocusEvent", "lostFocus", args)
  }

  return
}

// localPos()
func (this *QNativeGestureEvent) Localpos(args ...interface{}) (ret interface{}) {
  // localPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent8localPosEv
    // invoke: const QPointF & localPos()
    var ret0 = C.C_ZNK19QNativeGestureEvent8localPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "localPos", args)
  }

  return
}

// screenPos()
func (this *QNativeGestureEvent) Screenpos(args ...interface{}) (ret interface{}) {
  // screenPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent9screenPosEv
    // invoke: const QPointF & screenPos()
    var ret0 = C.C_ZNK19QNativeGestureEvent9screenPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "screenPos", args)
  }

  return
}

// globalPos()
func (this *QNativeGestureEvent) Globalpos(args ...interface{}) (ret interface{}) {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent9globalPosEv
    // invoke: const QPoint globalPos()
    var ret0 = C.C_ZNK19QNativeGestureEvent9globalPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "globalPos", args)
  }

  return
}

// pos()
func (this *QNativeGestureEvent) Pos(args ...interface{}) (ret interface{}) {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent3posEv
    // invoke: const QPoint pos()
    var ret0 = C.C_ZNK19QNativeGestureEvent3posEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "pos", args)
  }

  return
}

// value()
func (this *QNativeGestureEvent) Value(args ...interface{}) (ret interface{}) {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent5valueEv
    // invoke: qreal value()
    var ret0 = C.C_ZNK19QNativeGestureEvent5valueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "value", args)
  }

  return
}

// gestureType()
func (this *QNativeGestureEvent) Gesturetype(args ...interface{}) () {
  // gestureType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent11gestureTypeEv
    // invoke: Qt::NativeGestureType gestureType()
    C.C_ZNK19QNativeGestureEvent11gestureTypeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "gestureType", args)
  }

  return
}

// windowPos()
func (this *QNativeGestureEvent) Windowpos(args ...interface{}) (ret interface{}) {
  // windowPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent9windowPosEv
    // invoke: const QPointF & windowPos()
    var ret0 = C.C_ZNK19QNativeGestureEvent9windowPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "windowPos", args)
  }

  return
}

// ~QResizeEvent()
func (this *QResizeEvent) Freeqresizeevent(args ...interface{}) () {
  // ~QResizeEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QResizeEventD0Ev
    // invoke: void ~QResizeEvent()
    C.C_ZN12QResizeEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QResizeEvent", "~QResizeEvent", args)
  }

  return
}

// oldSize()
func (this *QResizeEvent) Oldsize(args ...interface{}) (ret interface{}) {
  // oldSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QResizeEvent7oldSizeEv
    // invoke: const QSize & oldSize()
    var ret0 = C.C_ZNK12QResizeEvent7oldSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "const QSize &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QResizeEvent", "oldSize", args)
  }

  return
}

// QResizeEvent(const class QSize &, const class QSize &)
func NewQResizeEvent(args ...interface{}) *QResizeEvent {
  // QResizeEvent(const class QSize &, const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QSize{}) // "const QSize &"
  vtys[0][1] = reflect.TypeOf(qtcore.QSize{}) // "const QSize &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QResizeEventC1ERK5QSizeS2_
    // invoke: void QResizeEvent(const class QSize &, const class QSize &)
    var arg0 = args[0].(*qtcore.QSize).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QSize).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QResizeEventC2ERK5QSizeS2_(arg0, arg1)
    return &QResizeEvent{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QResizeEvent", "QResizeEvent", args)
  }

  return nil // QResizeEvent{Qclsinst:qthis}
}

// size()
func (this *QResizeEvent) Size(args ...interface{}) (ret interface{}) {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QResizeEvent4sizeEv
    // invoke: const QSize & size()
    var ret0 = C.C_ZNK12QResizeEvent4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "const QSize &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QResizeEvent", "size", args)
  }

  return
}

// tip()
func (this *QStatusTipEvent) Tip(args ...interface{}) (ret interface{}) {
  // tip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QStatusTipEvent3tipEv
    // invoke: QString tip()
    var ret0 = C.C_ZNK15QStatusTipEvent3tipEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStatusTipEvent", "tip", args)
  }

  return
}

// ~QStatusTipEvent()
func (this *QStatusTipEvent) Freeqstatustipevent(args ...interface{}) () {
  // ~QStatusTipEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QStatusTipEventD0Ev
    // invoke: void ~QStatusTipEvent()
    C.C_ZN15QStatusTipEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStatusTipEvent", "~QStatusTipEvent", args)
  }

  return
}

// QStatusTipEvent(const class QString &)
func NewQStatusTipEvent(args ...interface{}) *QStatusTipEvent {
  // QStatusTipEvent(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QStatusTipEventC1ERK7QString
    // invoke: void QStatusTipEvent(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QStatusTipEventC2ERK7QString(arg0)
    return &QStatusTipEvent{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStatusTipEvent", "QStatusTipEvent", args)
  }

  return nil // QStatusTipEvent{Qclsinst:qthis}
}

// localPos()
func (this *QEnterEvent) Localpos(args ...interface{}) (ret interface{}) {
  // localPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent8localPosEv
    // invoke: const QPointF & localPos()
    var ret0 = C.C_ZNK11QEnterEvent8localPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QEnterEvent", "localPos", args)
  }

  return
}

// screenPos()
func (this *QEnterEvent) Screenpos(args ...interface{}) (ret interface{}) {
  // screenPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent9screenPosEv
    // invoke: const QPointF & screenPos()
    var ret0 = C.C_ZNK11QEnterEvent9screenPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QEnterEvent", "screenPos", args)
  }

  return
}

// globalPos()
func (this *QEnterEvent) Globalpos(args ...interface{}) (ret interface{}) {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent9globalPosEv
    // invoke: QPoint globalPos()
    var ret0 = C.C_ZNK11QEnterEvent9globalPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QEnterEvent", "globalPos", args)
  }

  return
}

// QEnterEvent(const class QPointF &, const class QPointF &, const class QPointF &)
func NewQEnterEvent(args ...interface{}) *QEnterEvent {
  // QEnterEvent(const class QPointF &, const class QPointF &, const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[0][2] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QEnterEventC1ERK7QPointFS2_S2_
    // invoke: void QEnterEvent(const class QPointF &, const class QPointF &, const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QEnterEventC2ERK7QPointFS2_S2_(arg0, arg1, arg2)
    return &QEnterEvent{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QEnterEvent", "QEnterEvent", args)
  }

  return nil // QEnterEvent{Qclsinst:qthis}
}

// pos()
func (this *QEnterEvent) Pos(args ...interface{}) (ret interface{}) {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent3posEv
    // invoke: QPoint pos()
    var ret0 = C.C_ZNK11QEnterEvent3posEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QEnterEvent", "pos", args)
  }

  return
}

// globalX()
func (this *QEnterEvent) Globalx(args ...interface{}) (ret interface{}) {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent7globalXEv
    // invoke: int globalX()
    var ret0 = C.C_ZNK11QEnterEvent7globalXEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QEnterEvent", "globalX", args)
  }

  return
}

// ~QEnterEvent()
func (this *QEnterEvent) Freeqenterevent(args ...interface{}) () {
  // ~QEnterEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QEnterEventD0Ev
    // invoke: void ~QEnterEvent()
    C.C_ZN11QEnterEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QEnterEvent", "~QEnterEvent", args)
  }

  return
}

// globalY()
func (this *QEnterEvent) Globaly(args ...interface{}) (ret interface{}) {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent7globalYEv
    // invoke: int globalY()
    var ret0 = C.C_ZNK11QEnterEvent7globalYEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QEnterEvent", "globalY", args)
  }

  return
}

// y()
func (this *QEnterEvent) Y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent1yEv
    // invoke: int y()
    C.C_ZNK11QEnterEvent1yEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QEnterEvent", "y", args)
  }

  return
}

// x()
func (this *QEnterEvent) X(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent1xEv
    // invoke: int x()
    C.C_ZNK11QEnterEvent1xEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QEnterEvent", "x", args)
  }

  return
}

// windowPos()
func (this *QEnterEvent) Windowpos(args ...interface{}) (ret interface{}) {
  // windowPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent9windowPosEv
    // invoke: const QPointF & windowPos()
    var ret0 = C.C_ZNK11QEnterEvent9windowPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QEnterEvent", "windowPos", args)
  }

  return
}

// QMoveEvent(const class QPoint &, const class QPoint &)
func NewQMoveEvent(args ...interface{}) *QMoveEvent {
  // QMoveEvent(const class QPoint &, const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[0][1] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMoveEventC1ERK6QPointS2_
    // invoke: void QMoveEvent(const class QPoint &, const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QMoveEventC2ERK6QPointS2_(arg0, arg1)
    return &QMoveEvent{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QMoveEvent", "QMoveEvent", args)
  }

  return nil // QMoveEvent{Qclsinst:qthis}
}

// ~QMoveEvent()
func (this *QMoveEvent) Freeqmoveevent(args ...interface{}) () {
  // ~QMoveEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMoveEventD0Ev
    // invoke: void ~QMoveEvent()
    C.C_ZN10QMoveEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMoveEvent", "~QMoveEvent", args)
  }

  return
}

// oldPos()
func (this *QMoveEvent) Oldpos(args ...interface{}) (ret interface{}) {
  // oldPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMoveEvent6oldPosEv
    // invoke: const QPoint & oldPos()
    var ret0 = C.C_ZNK10QMoveEvent6oldPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMoveEvent", "oldPos", args)
  }

  return
}

// pos()
func (this *QMoveEvent) Pos(args ...interface{}) (ret interface{}) {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMoveEvent3posEv
    // invoke: const QPoint & pos()
    var ret0 = C.C_ZNK10QMoveEvent3posEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMoveEvent", "pos", args)
  }

  return
}

// QHideEvent()
func NewQHideEvent(args ...interface{}) *QHideEvent {
  // QHideEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QHideEventC1Ev
    // invoke: void QHideEvent()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QHideEventC2Ev()
    return &QHideEvent{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QHideEvent", "QHideEvent", args)
  }

  return nil // QHideEvent{Qclsinst:qthis}
}

// ~QHideEvent()
func (this *QHideEvent) Freeqhideevent(args ...interface{}) () {
  // ~QHideEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QHideEventD0Ev
    // invoke: void ~QHideEvent()
    C.C_ZN10QHideEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QHideEvent", "~QHideEvent", args)
  }

  return
}

// ~QDragLeaveEvent()
func (this *QDragLeaveEvent) Freeqdragleaveevent(args ...interface{}) () {
  // ~QDragLeaveEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QDragLeaveEventD0Ev
    // invoke: void ~QDragLeaveEvent()
    C.C_ZN15QDragLeaveEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDragLeaveEvent", "~QDragLeaveEvent", args)
  }

  return
}

// QDragLeaveEvent()
func NewQDragLeaveEvent(args ...interface{}) *QDragLeaveEvent {
  // QDragLeaveEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QDragLeaveEventC1Ev
    // invoke: void QDragLeaveEvent()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QDragLeaveEventC2Ev()
    return &QDragLeaveEvent{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QDragLeaveEvent", "QDragLeaveEvent", args)
  }

  return nil // QDragLeaveEvent{Qclsinst:qthis}
}

// mimeData()
func (this *QDropEvent) Mimedata(args ...interface{}) (ret interface{}) {
  // mimeData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent8mimeDataEv
    // invoke: const QMimeData * mimeData()
    var ret0 = C.C_ZNK10QDropEvent8mimeDataEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QMimeData{}) // "const QMimeData *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDropEvent", "mimeData", args)
  }

  return
}

// acceptProposedAction()
func (this *QDropEvent) Acceptproposedaction(args ...interface{}) () {
  // acceptProposedAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QDropEvent20acceptProposedActionEv
    // invoke: void acceptProposedAction()
    C.C_ZN10QDropEvent20acceptProposedActionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDropEvent", "acceptProposedAction", args)
  }

  return
}

// possibleActions()
func (this *QDropEvent) Possibleactions(args ...interface{}) () {
  // possibleActions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent15possibleActionsEv
    // invoke: Qt::DropActions possibleActions()
    C.C_ZNK10QDropEvent15possibleActionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDropEvent", "possibleActions", args)
  }

  return
}

// posF()
func (this *QDropEvent) Posf(args ...interface{}) (ret interface{}) {
  // posF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent4posFEv
    // invoke: const QPointF & posF()
    var ret0 = C.C_ZNK10QDropEvent4posFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDropEvent", "posF", args)
  }

  return
}

// pos()
func (this *QDropEvent) Pos(args ...interface{}) (ret interface{}) {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent3posEv
    // invoke: QPoint pos()
    var ret0 = C.C_ZNK10QDropEvent3posEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDropEvent", "pos", args)
  }

  return
}

// source()
func (this *QDropEvent) Source(args ...interface{}) (ret interface{}) {
  // source()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent6sourceEv
    // invoke: QObject * source()
    var ret0 = C.C_ZNK10QDropEvent6sourceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDropEvent", "source", args)
  }

  return
}

// proposedAction()
func (this *QDropEvent) Proposedaction(args ...interface{}) () {
  // proposedAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent14proposedActionEv
    // invoke: Qt::DropAction proposedAction()
    C.C_ZNK10QDropEvent14proposedActionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDropEvent", "proposedAction", args)
  }

  return
}

// dropAction()
func (this *QDropEvent) Dropaction(args ...interface{}) () {
  // dropAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent10dropActionEv
    // invoke: Qt::DropAction dropAction()
    C.C_ZNK10QDropEvent10dropActionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDropEvent", "dropAction", args)
  }

  return
}

// keyboardModifiers()
func (this *QDropEvent) Keyboardmodifiers(args ...interface{}) () {
  // keyboardModifiers()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent17keyboardModifiersEv
    // invoke: Qt::KeyboardModifiers keyboardModifiers()
    C.C_ZNK10QDropEvent17keyboardModifiersEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDropEvent", "keyboardModifiers", args)
  }

  return
}

// ~QDropEvent()
func (this *QDropEvent) Freeqdropevent(args ...interface{}) () {
  // ~QDropEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QDropEventD0Ev
    // invoke: void ~QDropEvent()
    C.C_ZN10QDropEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDropEvent", "~QDropEvent", args)
  }

  return
}

// mouseButtons()
func (this *QDropEvent) Mousebuttons(args ...interface{}) () {
  // mouseButtons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent12mouseButtonsEv
    // invoke: Qt::MouseButtons mouseButtons()
    C.C_ZNK10QDropEvent12mouseButtonsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDropEvent", "mouseButtons", args)
  }

  return
}

// modifiers()
func (this *QInputEvent) Modifiers(args ...interface{}) () {
  // modifiers()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QInputEvent9modifiersEv
    // invoke: Qt::KeyboardModifiers modifiers()
    C.C_ZNK11QInputEvent9modifiersEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QInputEvent", "modifiers", args)
  }

  return
}

// timestamp()
func (this *QInputEvent) Timestamp(args ...interface{}) (ret interface{}) {
  // timestamp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QInputEvent9timestampEv
    // invoke: ulong timestamp()
    var ret0 = C.C_ZNK11QInputEvent9timestampEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "ulong"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QInputEvent", "timestamp", args)
  }

  return
}

// setTimestamp(ulong)
func (this *QInputEvent) Settimestamp(args ...interface{}) () {
  // setTimestamp(ulong)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "ulong"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QInputEvent12setTimestampEm
    // invoke: void setTimestamp(ulong)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QInputEvent12setTimestampEm(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputEvent", "setTimestamp", args)
  }

  return
}

// ~QInputEvent()
func (this *QInputEvent) Freeqinputevent(args ...interface{}) () {
  // ~QInputEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QInputEventD0Ev
    // invoke: void ~QInputEvent()
    C.C_ZN11QInputEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QInputEvent", "~QInputEvent", args)
  }

  return
}

// applicationState()
func (this *QApplicationStateChangeEvent) Applicationstate(args ...interface{}) () {
  // applicationState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK28QApplicationStateChangeEvent16applicationStateEv
    // invoke: Qt::ApplicationState applicationState()
    C.C_ZNK28QApplicationStateChangeEvent16applicationStateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QApplicationStateChangeEvent", "applicationState", args)
  }

  return
}

// contentPos()
func (this *QScrollEvent) Contentpos(args ...interface{}) (ret interface{}) {
  // contentPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QScrollEvent10contentPosEv
    // invoke: QPointF contentPos()
    var ret0 = C.C_ZNK12QScrollEvent10contentPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScrollEvent", "contentPos", args)
  }

  return
}

// overshootDistance()
func (this *QScrollEvent) Overshootdistance(args ...interface{}) (ret interface{}) {
  // overshootDistance()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QScrollEvent17overshootDistanceEv
    // invoke: QPointF overshootDistance()
    var ret0 = C.C_ZNK12QScrollEvent17overshootDistanceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScrollEvent", "overshootDistance", args)
  }

  return
}

// ~QScrollEvent()
func (this *QScrollEvent) Freeqscrollevent(args ...interface{}) () {
  // ~QScrollEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QScrollEventD0Ev
    // invoke: void ~QScrollEvent()
    C.C_ZN12QScrollEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QScrollEvent", "~QScrollEvent", args)
  }

  return
}

// scrollState()
func (this *QScrollEvent) Scrollstate(args ...interface{}) () {
  // scrollState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QScrollEvent11scrollStateEv
    // invoke: QScrollEvent::ScrollState scrollState()
    C.C_ZNK12QScrollEvent11scrollStateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QScrollEvent", "scrollState", args)
  }

  return
}

// count()
func (this *QKeyEvent) Count(args ...interface{}) (ret interface{}) {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent5countEv
    // invoke: int count()
    var ret0 = C.C_ZNK9QKeyEvent5countEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QKeyEvent", "count", args)
  }

  return
}

// modifiers()
func (this *QKeyEvent) Modifiers(args ...interface{}) () {
  // modifiers()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent9modifiersEv
    // invoke: Qt::KeyboardModifiers modifiers()
    C.C_ZNK9QKeyEvent9modifiersEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QKeyEvent", "modifiers", args)
  }

  return
}

// nativeModifiers()
func (this *QKeyEvent) Nativemodifiers(args ...interface{}) (ret interface{}) {
  // nativeModifiers()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent15nativeModifiersEv
    // invoke: quint32 nativeModifiers()
    var ret0 = C.C_ZNK9QKeyEvent15nativeModifiersEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "quint32"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QKeyEvent", "nativeModifiers", args)
  }

  return
}

// text()
func (this *QKeyEvent) Text(args ...interface{}) (ret interface{}) {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent4textEv
    // invoke: QString text()
    var ret0 = C.C_ZNK9QKeyEvent4textEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QKeyEvent", "text", args)
  }

  return
}

// ~QKeyEvent()
func (this *QKeyEvent) Freeqkeyevent(args ...interface{}) () {
  // ~QKeyEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QKeyEventD0Ev
    // invoke: void ~QKeyEvent()
    C.C_ZN9QKeyEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QKeyEvent", "~QKeyEvent", args)
  }

  return
}

// nativeScanCode()
func (this *QKeyEvent) Nativescancode(args ...interface{}) (ret interface{}) {
  // nativeScanCode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent14nativeScanCodeEv
    // invoke: quint32 nativeScanCode()
    var ret0 = C.C_ZNK9QKeyEvent14nativeScanCodeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "quint32"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QKeyEvent", "nativeScanCode", args)
  }

  return
}

// isAutoRepeat()
func (this *QKeyEvent) Isautorepeat(args ...interface{}) (ret interface{}) {
  // isAutoRepeat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent12isAutoRepeatEv
    // invoke: bool isAutoRepeat()
    var ret0 = C.C_ZNK9QKeyEvent12isAutoRepeatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QKeyEvent", "isAutoRepeat", args)
  }

  return
}

// key()
func (this *QKeyEvent) Key(args ...interface{}) (ret interface{}) {
  // key()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent3keyEv
    // invoke: int key()
    var ret0 = C.C_ZNK9QKeyEvent3keyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QKeyEvent", "key", args)
  }

  return
}

// nativeVirtualKey()
func (this *QKeyEvent) Nativevirtualkey(args ...interface{}) (ret interface{}) {
  // nativeVirtualKey()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent16nativeVirtualKeyEv
    // invoke: quint32 nativeVirtualKey()
    var ret0 = C.C_ZNK9QKeyEvent16nativeVirtualKeyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "quint32"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QKeyEvent", "nativeVirtualKey", args)
  }

  return
}

// globalPos()
func (this *QContextMenuEvent) Globalpos(args ...interface{}) (ret interface{}) {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent9globalPosEv
    // invoke: const QPoint & globalPos()
    var ret0 = C.C_ZNK17QContextMenuEvent9globalPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "globalPos", args)
  }

  return
}

// pos()
func (this *QContextMenuEvent) Pos(args ...interface{}) (ret interface{}) {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent3posEv
    // invoke: const QPoint & pos()
    var ret0 = C.C_ZNK17QContextMenuEvent3posEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "pos", args)
  }

  return
}

// y()
func (this *QContextMenuEvent) Y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent1yEv
    // invoke: int y()
    C.C_ZNK17QContextMenuEvent1yEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "y", args)
  }

  return
}

// reason()
func (this *QContextMenuEvent) Reason(args ...interface{}) () {
  // reason()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent6reasonEv
    // invoke: QContextMenuEvent::Reason reason()
    C.C_ZNK17QContextMenuEvent6reasonEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "reason", args)
  }

  return
}

// x()
func (this *QContextMenuEvent) X(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent1xEv
    // invoke: int x()
    C.C_ZNK17QContextMenuEvent1xEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "x", args)
  }

  return
}

// globalX()
func (this *QContextMenuEvent) Globalx(args ...interface{}) (ret interface{}) {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent7globalXEv
    // invoke: int globalX()
    var ret0 = C.C_ZNK17QContextMenuEvent7globalXEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "globalX", args)
  }

  return
}

// globalY()
func (this *QContextMenuEvent) Globaly(args ...interface{}) (ret interface{}) {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent7globalYEv
    // invoke: int globalY()
    var ret0 = C.C_ZNK17QContextMenuEvent7globalYEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "globalY", args)
  }

  return
}

// ~QContextMenuEvent()
func (this *QContextMenuEvent) Freeqcontextmenuevent(args ...interface{}) () {
  // ~QContextMenuEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QContextMenuEventD0Ev
    // invoke: void ~QContextMenuEvent()
    C.C_ZN17QContextMenuEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "~QContextMenuEvent", args)
  }

  return
}

// setContentPosRange(const class QRectF &)
func (this *QScrollPrepareEvent) Setcontentposrange(args ...interface{}) () {
  // setContentPosRange(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QScrollPrepareEvent18setContentPosRangeERK6QRectF
    // invoke: void setContentPosRange(const class QRectF &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QScrollPrepareEvent18setContentPosRangeERK6QRectF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "setContentPosRange", args)
  }

  return
}

// contentPos()
func (this *QScrollPrepareEvent) Contentpos(args ...interface{}) (ret interface{}) {
  // contentPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QScrollPrepareEvent10contentPosEv
    // invoke: QPointF contentPos()
    var ret0 = C.C_ZNK19QScrollPrepareEvent10contentPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "contentPos", args)
  }

  return
}

// setContentPos(const class QPointF &)
func (this *QScrollPrepareEvent) Setcontentpos(args ...interface{}) () {
  // setContentPos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QScrollPrepareEvent13setContentPosERK7QPointF
    // invoke: void setContentPos(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QScrollPrepareEvent13setContentPosERK7QPointF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "setContentPos", args)
  }

  return
}

// ~QScrollPrepareEvent()
func (this *QScrollPrepareEvent) Freeqscrollprepareevent(args ...interface{}) () {
  // ~QScrollPrepareEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QScrollPrepareEventD0Ev
    // invoke: void ~QScrollPrepareEvent()
    C.C_ZN19QScrollPrepareEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "~QScrollPrepareEvent", args)
  }

  return
}

// viewportSize()
func (this *QScrollPrepareEvent) Viewportsize(args ...interface{}) (ret interface{}) {
  // viewportSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QScrollPrepareEvent12viewportSizeEv
    // invoke: QSizeF viewportSize()
    var ret0 = C.C_ZNK19QScrollPrepareEvent12viewportSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSizeF{}) // "QSizeF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "viewportSize", args)
  }

  return
}

// setViewportSize(const class QSizeF &)
func (this *QScrollPrepareEvent) Setviewportsize(args ...interface{}) () {
  // setViewportSize(const class QSizeF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QSizeF{}) // "const QSizeF &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QScrollPrepareEvent15setViewportSizeERK6QSizeF
    // invoke: void setViewportSize(const class QSizeF &)
    var arg0 = args[0].(*qtcore.QSizeF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QScrollPrepareEvent15setViewportSizeERK6QSizeF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "setViewportSize", args)
  }

  return
}

// QScrollPrepareEvent(const class QPointF &)
func NewQScrollPrepareEvent(args ...interface{}) *QScrollPrepareEvent {
  // QScrollPrepareEvent(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QScrollPrepareEventC1ERK7QPointF
    // invoke: void QScrollPrepareEvent(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QScrollPrepareEventC2ERK7QPointF(arg0)
    return &QScrollPrepareEvent{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "QScrollPrepareEvent", args)
  }

  return nil // QScrollPrepareEvent{Qclsinst:qthis}
}

// contentPosRange()
func (this *QScrollPrepareEvent) Contentposrange(args ...interface{}) (ret interface{}) {
  // contentPosRange()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QScrollPrepareEvent15contentPosRangeEv
    // invoke: QRectF contentPosRange()
    var ret0 = C.C_ZNK19QScrollPrepareEvent15contentPosRangeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "contentPosRange", args)
  }

  return
}

// startPos()
func (this *QScrollPrepareEvent) Startpos(args ...interface{}) (ret interface{}) {
  // startPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QScrollPrepareEvent8startPosEv
    // invoke: QPointF startPos()
    var ret0 = C.C_ZNK19QScrollPrepareEvent8startPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "startPos", args)
  }

  return
}

// ~QShortcutEvent()
func (this *QShortcutEvent) Freeqshortcutevent(args ...interface{}) () {
  // ~QShortcutEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QShortcutEventD0Ev
    // invoke: void ~QShortcutEvent()
    C.C_ZN14QShortcutEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QShortcutEvent", "~QShortcutEvent", args)
  }

  return
}

// isAmbiguous()
func (this *QShortcutEvent) Isambiguous(args ...interface{}) (ret interface{}) {
  // isAmbiguous()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QShortcutEvent11isAmbiguousEv
    // invoke: bool isAmbiguous()
    var ret0 = C.C_ZNK14QShortcutEvent11isAmbiguousEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QShortcutEvent", "isAmbiguous", args)
  }

  return
}

// shortcutId()
func (this *QShortcutEvent) Shortcutid(args ...interface{}) (ret interface{}) {
  // shortcutId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QShortcutEvent10shortcutIdEv
    // invoke: int shortcutId()
    var ret0 = C.C_ZNK14QShortcutEvent10shortcutIdEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QShortcutEvent", "shortcutId", args)
  }

  return
}

// key()
func (this *QShortcutEvent) Key(args ...interface{}) (ret interface{}) {
  // key()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QShortcutEvent3keyEv
    // invoke: const QKeySequence & key()
    var ret0 = C.C_ZNK14QShortcutEvent3keyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QKeySequence{}) // "const QKeySequence &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QShortcutEvent", "key", args)
  }

  return
}

// QShortcutEvent(const class QKeySequence &, int, _Bool)
func NewQShortcutEvent(args ...interface{}) *QShortcutEvent {
  // QShortcutEvent(const class QKeySequence &, int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QKeySequence{}) // "const QKeySequence &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QShortcutEventC1ERK12QKeySequenceib
    // invoke: void QShortcutEvent(const class QKeySequence &, int, _Bool)
    var arg0 = args[0].(*QKeySequence).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.bool(args[2].(bool))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QShortcutEventC2ERK12QKeySequenceib(arg0, arg1, arg2)
    return &QShortcutEvent{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QShortcutEvent", "QShortcutEvent", args)
  }

  return nil // QShortcutEvent{Qclsinst:qthis}
}

// isOverride()
func (this *QWindowStateChangeEvent) Isoverride(args ...interface{}) (ret interface{}) {
  // isOverride()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QWindowStateChangeEvent10isOverrideEv
    // invoke: bool isOverride()
    var ret0 = C.C_ZNK23QWindowStateChangeEvent10isOverrideEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWindowStateChangeEvent", "isOverride", args)
  }

  return
}

// oldState()
func (this *QWindowStateChangeEvent) Oldstate(args ...interface{}) () {
  // oldState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QWindowStateChangeEvent8oldStateEv
    // invoke: Qt::WindowStates oldState()
    C.C_ZNK23QWindowStateChangeEvent8oldStateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWindowStateChangeEvent", "oldState", args)
  }

  return
}

// ~QWindowStateChangeEvent()
func (this *QWindowStateChangeEvent) Freeqwindowstatechangeevent(args ...interface{}) () {
  // ~QWindowStateChangeEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QWindowStateChangeEventD0Ev
    // invoke: void ~QWindowStateChangeEvent()
    C.C_ZN23QWindowStateChangeEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWindowStateChangeEvent", "~QWindowStateChangeEvent", args)
  }

  return
}

// ~QInputMethodQueryEvent()
func (this *QInputMethodQueryEvent) Freeqinputmethodqueryevent(args ...interface{}) () {
  // ~QInputMethodQueryEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QInputMethodQueryEventD0Ev
    // invoke: void ~QInputMethodQueryEvent()
    C.C_ZN22QInputMethodQueryEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QInputMethodQueryEvent", "~QInputMethodQueryEvent", args)
  }

  return
}

// queries()
func (this *QInputMethodQueryEvent) Queries(args ...interface{}) () {
  // queries()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QInputMethodQueryEvent7queriesEv
    // invoke: Qt::InputMethodQueries queries()
    C.C_ZNK22QInputMethodQueryEvent7queriesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QInputMethodQueryEvent", "queries", args)
  }

  return
}

// <= body block end

