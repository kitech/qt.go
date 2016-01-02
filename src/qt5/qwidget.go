package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtWidgets/qwidget.h
// dst-file: /src/widgets/qwidget.go
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
  // proto:  void QWidget::setGraphicsEffect(QGraphicsEffect * effect);
extern void _ZN7QWidget17setGraphicsEffectEP15QGraphicsEffect(void* qthis, void* arg0);
  // proto:  QString QWidget::accessibleDescription();
extern void _ZNK7QWidget21accessibleDescriptionEv(void* qthis);
  // proto:  QGraphicsEffect * QWidget::graphicsEffect();
extern void _ZNK7QWidget14graphicsEffectEv(void* qthis);
  // proto:  bool QWidget::isFullScreen();
extern void _ZNK7QWidget12isFullScreenEv(void* qthis);
  // proto:  QSize QWidget::maximumSize();
extern void _ZNK7QWidget11maximumSizeEv(void* qthis);
  // proto:  bool QWidget::isEnabledToTLW();
extern void _ZNK7QWidget14isEnabledToTLWEv(void* qthis);
  // proto:  void QWidget::setStatusTip(const QString & );
extern void _ZN7QWidget12setStatusTipERK7QString(void* qthis, void* arg0);
  // proto:  void QWidget::setSizeIncrement(const QSize & );
extern void _ZN7QWidget16setSizeIncrementERK5QSize(void* qthis, void* arg0);
  // proto:  QWidget * QWidget::focusWidget();
extern void _ZNK7QWidget11focusWidgetEv(void* qthis);
  // proto:  bool QWidget::isTopLevel();
extern void _ZNK7QWidget10isTopLevelEv(void* qthis);
  // proto:  bool QWidget::acceptDrops();
extern void _ZNK7QWidget11acceptDropsEv(void* qthis);
  // proto:  bool QWidget::isWindow();
extern void _ZNK7QWidget8isWindowEv(void* qthis);
  // proto:  void QWidget::setFixedSize(const QSize & );
extern void _ZN7QWidget12setFixedSizeERK5QSize(void* qthis, void* arg0);
  // proto:  bool QWidget::isVisible();
extern void _ZNK7QWidget9isVisibleEv(void* qthis);
  // proto:  int QWidget::minimumHeight();
extern void _ZNK7QWidget13minimumHeightEv(void* qthis);
  // proto:  QSize QWidget::sizeIncrement();
extern void _ZNK7QWidget13sizeIncrementEv(void* qthis);
  // proto:  void QWidget::repaint(const QRect & );
extern void _ZN7QWidget7repaintERK5QRect(void* qthis, void* arg0);
  // proto:  void QWidget::update(int x, int y, int w, int h);
extern void demth_ZN7QWidget6updateEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  QString QWidget::windowFilePath();
extern void _ZNK7QWidget14windowFilePathEv(void* qthis);
  // proto:  void QWidget::clearMask();
extern void _ZN7QWidget9clearMaskEv(void* qthis);
  // proto:  QPoint QWidget::mapFromParent(const QPoint & );
extern void _ZNK7QWidget13mapFromParentERK6QPoint(void* qthis, void* arg0);
  // proto:  QRect QWidget::rect();
extern void demth_ZNK7QWidget4rectEv(void* qthis);
  // proto:  void QWidget::unsetLayoutDirection();
extern void _ZN7QWidget20unsetLayoutDirectionEv(void* qthis);
  // proto:  void QWidget::setMinimumSize(const QSize & );
extern void _ZN7QWidget14setMinimumSizeERK5QSize(void* qthis, void* arg0);
  // proto:  bool QWidget::isActiveWindow();
extern void _ZNK7QWidget14isActiveWindowEv(void* qthis);
  // proto:  void QWidget::grabKeyboard();
extern void _ZN7QWidget12grabKeyboardEv(void* qthis);
  // proto:  QSize QWidget::frameSize();
extern void _ZNK7QWidget9frameSizeEv(void* qthis);
  // proto:  void QWidget::setFocus();
extern void demth_ZN7QWidget8setFocusEv(void* qthis);
  // proto:  QPoint QWidget::mapToParent(const QPoint & );
extern void _ZNK7QWidget11mapToParentERK6QPoint(void* qthis, void* arg0);
  // proto:  void QWidget::updateGeometry();
extern void _ZN7QWidget14updateGeometryEv(void* qthis);
  // proto:  void QWidget::repaint(const QRegion & );
extern void _ZN7QWidget7repaintERK7QRegion(void* qthis, void* arg0);
  // proto:  void QWidget::insertAction(QAction * before, QAction * action);
extern void _ZN7QWidget12insertActionEP7QActionS1_(void* qthis, void* arg0, void* arg1);
  // proto:  void QWidget::setWindowRole(const QString & );
extern void _ZN7QWidget13setWindowRoleERK7QString(void* qthis, void* arg0);
  // proto:  int QWidget::toolTipDuration();
extern void _ZNK7QWidget15toolTipDurationEv(void* qthis);
  // proto:  void QWidget::setPalette(const QPalette & );
extern void _ZN7QWidget10setPaletteERK8QPalette(void* qthis, void* arg0);
  // proto:  void QWidget::setStyleSheet(const QString & styleSheet);
extern void _ZN7QWidget13setStyleSheetERK7QString(void* qthis, void* arg0);
  // proto:  QString QWidget::windowIconText();
extern void _ZNK7QWidget14windowIconTextEv(void* qthis);
  // proto:  void QWidget::releaseMouse();
extern void _ZN7QWidget12releaseMouseEv(void* qthis);
  // proto:  bool QWidget::isModal();
extern void _ZNK7QWidget7isModalEv(void* qthis);
  // proto:  void QWidget::setStyle(QStyle * );
extern void _ZN7QWidget8setStyleEP6QStyle(void* qthis, void* arg0);
  // proto:  void QWidget::repaint();
extern void _ZN7QWidget7repaintEv(void* qthis);
  // proto:  void QWidget::setBaseSize(int basew, int baseh);
extern void _ZN7QWidget11setBaseSizeEii(void* qthis, int arg0, int arg1);
  // proto:  bool QWidget::isWindowModified();
extern void _ZNK7QWidget16isWindowModifiedEv(void* qthis);
  // proto:  const QRect & QWidget::geometry();
extern void _ZNK7QWidget8geometryEv(void* qthis);
  // proto:  void QWidget::setAccessibleDescription(const QString & description);
extern void _ZN7QWidget24setAccessibleDescriptionERK7QString(void* qthis, void* arg0);
  // proto:  QWidget * QWidget::nextInFocusChain();
extern void _ZNK7QWidget16nextInFocusChainEv(void* qthis);
  // proto: static void QWidget::setTabOrder(QWidget * , QWidget * );
extern void _ZN7QWidget11setTabOrderEPS_S0_(void* arg0, void* arg1);
  // proto:  QRect QWidget::frameGeometry();
extern void _ZNK7QWidget13frameGeometryEv(void* qthis);
  // proto:  QSize QWidget::sizeHint();
extern void _ZNK7QWidget8sizeHintEv(void* qthis);
  // proto:  void QWidget::setMinimumWidth(int minw);
extern void _ZN7QWidget15setMinimumWidthEi(void* qthis, int arg0);
  // proto:  bool QWidget::isVisibleTo(const QWidget * );
extern void _ZNK7QWidget11isVisibleToEPKS_(void* qthis, void* arg0);
  // proto:  void QWidget::setMaximumSize(int maxw, int maxh);
extern void _ZN7QWidget14setMaximumSizeEii(void* qthis, int arg0, int arg1);
  // proto:  bool QWidget::hasMouseTracking();
extern void _ZNK7QWidget16hasMouseTrackingEv(void* qthis);
  // proto:  void QWidget::update(const QRect & );
extern void _ZN7QWidget6updateERK5QRect(void* qthis, void* arg0);
  // proto:  bool QWidget::isHidden();
extern void demth_ZNK7QWidget8isHiddenEv(void* qthis);
  // proto:  int QWidget::devType();
extern void _ZNK7QWidget7devTypeEv(void* qthis);
  // proto:  QWidget * QWidget::childAt(int x, int y);
extern void demth_ZNK7QWidget7childAtEii(void* qthis, int arg0, int arg1);
  // proto:  void QWidget::activateWindow();
extern void _ZN7QWidget14activateWindowEv(void* qthis);
  // proto:  QRect QWidget::normalGeometry();
extern void _ZNK7QWidget14normalGeometryEv(void* qthis);
  // proto:  QString QWidget::windowTitle();
extern void _ZNK7QWidget11windowTitleEv(void* qthis);
  // proto:  void QWidget::grabMouse(const QCursor & );
extern void _ZN7QWidget9grabMouseERK7QCursor(void* qthis, void* arg0);
  // proto:  QPixmap QWidget::grab(const QRect & rectangle);
extern void _ZN7QWidget4grabERK5QRect(void* qthis, void* arg0);
  // proto:  void QWidget::setVisible(bool visible);
extern void _ZN7QWidget10setVisibleEb(void* qthis, bool arg0);
  // proto:  bool QWidget::isEnabledTo(const QWidget * );
extern void _ZNK7QWidget11isEnabledToEPKS_(void* qthis, void* arg0);
  // proto:  bool QWidget::isLeftToRight();
extern void demth_ZNK7QWidget13isLeftToRightEv(void* qthis);
  // proto:  void QWidget::setGeometry(const QRect & );
extern void _ZN7QWidget11setGeometryERK5QRect(void* qthis, void* arg0);
  // proto:  void QWidget::unsetLocale();
extern void _ZN7QWidget11unsetLocaleEv(void* qthis);
  // proto:  void QWidget::showNormal();
extern void _ZN7QWidget10showNormalEv(void* qthis);
  // proto:  int QWidget::y();
extern void _ZNK7QWidget1yEv(void* qthis);
  // proto:  int QWidget::width();
extern void demth_ZNK7QWidget5widthEv(void* qthis);
  // proto:  bool QWidget::isMaximized();
extern void _ZNK7QWidget11isMaximizedEv(void* qthis);
  // proto:  void QWidget::resize(const QSize & );
extern void _ZN7QWidget6resizeERK5QSize(void* qthis, void* arg0);
  // proto:  QWindow * QWidget::windowHandle();
extern void _ZNK7QWidget12windowHandleEv(void* qthis);
  // proto:  QString QWidget::accessibleName();
extern void _ZNK7QWidget14accessibleNameEv(void* qthis);
  // proto:  void QWidget::setContentsMargins(const QMargins & margins);
extern void _ZN7QWidget18setContentsMarginsERK8QMargins(void* qthis, void* arg0);
  // proto:  QByteArray QWidget::saveGeometry();
extern void _ZNK7QWidget12saveGeometryEv(void* qthis);
  // proto:  bool QWidget::isEnabled();
extern void _ZNK7QWidget9isEnabledEv(void* qthis);
  // proto:  void QWidget::setFixedHeight(int h);
extern void _ZN7QWidget14setFixedHeightEi(void* qthis, int arg0);
  // proto:  QRegion QWidget::mask();
extern void _ZNK7QWidget4maskEv(void* qthis);
  // proto:  void QWidget::stackUnder(QWidget * );
extern void _ZN7QWidget10stackUnderEPS_(void* qthis, void* arg0);
  // proto:  QPaintEngine * QWidget::paintEngine();
extern void _ZNK7QWidget11paintEngineEv(void* qthis);
  // proto:  void QWidget::setAcceptDrops(bool on);
extern void _ZN7QWidget14setAcceptDropsEb(void* qthis, bool arg0);
  // proto:  void QWidget::move(const QPoint & );
extern void _ZN7QWidget4moveERK6QPoint(void* qthis, void* arg0);
  // proto:  QList<QAction *> QWidget::actions();
extern void _ZNK7QWidget7actionsEv(void* qthis);
  // proto:  void QWidget::show();
extern void _ZN7QWidget4showEv(void* qthis);
  // proto: static QWidget * QWidget::keyboardGrabber();
extern void _ZN7QWidget15keyboardGrabberEv();
  // proto:  QPoint QWidget::mapTo(const QWidget * , const QPoint & );
extern void _ZNK7QWidget5mapToEPKS_RK6QPoint(void* qthis, void* arg0, void* arg1);
  // proto:  int QWidget::minimumWidth();
extern void _ZNK7QWidget12minimumWidthEv(void* qthis);
  // proto:  QFontInfo QWidget::fontInfo();
extern void _ZNK7QWidget8fontInfoEv(void* qthis);
  // proto:  bool QWidget::autoFillBackground();
extern void _ZNK7QWidget18autoFillBackgroundEv(void* qthis);
  // proto:  void QWidget::scroll(int dx, int dy, const QRect & );
extern void _ZN7QWidget6scrollEiiRK5QRect(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  QFontMetrics QWidget::fontMetrics();
extern void _ZNK7QWidget11fontMetricsEv(void* qthis);
  // proto:  void QWidget::adjustSize();
extern void _ZN7QWidget10adjustSizeEv(void* qthis);
  // proto:  QWidget * QWidget::previousInFocusChain();
extern void _ZNK7QWidget20previousInFocusChainEv(void* qthis);
  // proto:  bool QWidget::updatesEnabled();
extern void demth_ZNK7QWidget14updatesEnabledEv(void* qthis);
  // proto:  void QWidget::setMaximumHeight(int maxh);
extern void _ZN7QWidget16setMaximumHeightEi(void* qthis, int arg0);
  // proto:  void QWidget::showMaximized();
extern void _ZN7QWidget13showMaximizedEv(void* qthis);
  // proto:  QPoint QWidget::mapFrom(const QWidget * , const QPoint & );
extern void _ZNK7QWidget7mapFromEPKS_RK6QPoint(void* qthis, void* arg0, void* arg1);
  // proto:  int QWidget::x();
extern void _ZNK7QWidget1xEv(void* qthis);
  // proto:  void QWidget::clearFocus();
extern void _ZN7QWidget10clearFocusEv(void* qthis);
  // proto: static QWidget * QWidget::find(WId );
extern void _ZN7QWidget4findEi(uint32_t* arg0);
  // proto:  const QPalette & QWidget::palette();
extern void _ZNK7QWidget7paletteEv(void* qthis);
  // proto:  void QWidget::setSizePolicy(QSizePolicy );
extern void _ZN7QWidget13setSizePolicyE11QSizePolicy(void* qthis, void* arg0);
  // proto:  void QWidget::setMask(const QRegion & );
extern void _ZN7QWidget7setMaskERK7QRegion(void* qthis, void* arg0);
  // proto:  void QWidget::setMaximumWidth(int maxw);
extern void _ZN7QWidget15setMaximumWidthEi(void* qthis, int arg0);
  // proto:  void QWidget::setWindowIconText(const QString & );
extern void _ZN7QWidget17setWindowIconTextERK7QString(void* qthis, void* arg0);
  // proto:  void QWidget::setWindowIcon(const QIcon & icon);
extern void _ZN7QWidget13setWindowIconERK5QIcon(void* qthis, void* arg0);
  // proto:  void QWidget::~QWidget();
extern void _ZN7QWidgetD0Ev(void* qthis);
  // proto:  void QWidget::getContentsMargins(int * left, int * top, int * right, int * bottom);
extern void _ZNK7QWidget18getContentsMarginsEPiS0_S0_S0_(void* qthis, int* arg0, int* arg1, int* arg2, int* arg3);
  // proto:  QSize QWidget::minimumSizeHint();
extern void _ZNK7QWidget15minimumSizeHintEv(void* qthis);
  // proto:  void QWidget::setWindowModified(bool );
extern void _ZN7QWidget17setWindowModifiedEb(void* qthis, bool arg0);
  // proto:  bool QWidget::restoreGeometry(const QByteArray & geometry);
extern void _ZN7QWidget15restoreGeometryERK10QByteArray(void* qthis, void* arg0);
  // proto:  QLayout * QWidget::layout();
extern void _ZNK7QWidget6layoutEv(void* qthis);
  // proto:  QRect QWidget::contentsRect();
extern void _ZNK7QWidget12contentsRectEv(void* qthis);
  // proto:  QBackingStore * QWidget::backingStore();
extern void _ZNK7QWidget12backingStoreEv(void* qthis);
  // proto:  QWidget * QWidget::focusProxy();
extern void _ZNK7QWidget10focusProxyEv(void* qthis);
  // proto:  QString QWidget::styleSheet();
extern void _ZNK7QWidget10styleSheetEv(void* qthis);
  // proto:  QWidget * QWidget::childAt(const QPoint & p);
extern void _ZNK7QWidget7childAtERK6QPoint(void* qthis, void* arg0);
  // proto:  void QWidget::repaint(int x, int y, int w, int h);
extern void _ZN7QWidget7repaintEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  QString QWidget::whatsThis();
extern void _ZNK7QWidget9whatsThisEv(void* qthis);
  // proto:  const QFont & QWidget::font();
extern void _ZNK7QWidget4fontEv(void* qthis);
  // proto:  void QWidget::setMinimumSize(int minw, int minh);
extern void _ZN7QWidget14setMinimumSizeEii(void* qthis, int arg0, int arg1);
  // proto:  const QMetaObject * QWidget::metaObject();
extern void _ZNK7QWidget10metaObjectEv(void* qthis);
  // proto:  void QWidget::setMinimumHeight(int minh);
extern void _ZN7QWidget16setMinimumHeightEi(void* qthis, int arg0);
  // proto:  void QWidget::update(const QRegion & );
extern void _ZN7QWidget6updateERK7QRegion(void* qthis, void* arg0);
  // proto:  qreal QWidget::windowOpacity();
extern void _ZNK7QWidget13windowOpacityEv(void* qthis);
  // proto:  QRegion QWidget::childrenRegion();
extern void _ZNK7QWidget14childrenRegionEv(void* qthis);
  // proto:  void QWidget::setWindowFilePath(const QString & filePath);
extern void _ZN7QWidget17setWindowFilePathERK7QString(void* qthis, void* arg0);
  // proto:  void QWidget::setShortcutEnabled(int id, bool enable);
extern void _ZN7QWidget18setShortcutEnabledEib(void* qthis, int arg0, bool arg1);
  // proto:  void QWidget::raise();
extern void _ZN7QWidget5raiseEv(void* qthis);
  // proto:  QString QWidget::statusTip();
extern void _ZNK7QWidget9statusTipEv(void* qthis);
  // proto:  QRect QWidget::childrenRect();
extern void _ZNK7QWidget12childrenRectEv(void* qthis);
  // proto:  void QWidget::setParent(QWidget * parent);
extern void _ZN7QWidget9setParentEPS_(void* qthis, void* arg0);
  // proto:  QRegion QWidget::visibleRegion();
extern void _ZNK7QWidget13visibleRegionEv(void* qthis);
  // proto:  QLocale QWidget::locale();
extern void _ZNK7QWidget6localeEv(void* qthis);
  // proto:  void QWidget::releaseKeyboard();
extern void _ZN7QWidget15releaseKeyboardEv(void* qthis);
  // proto: static QWidget * QWidget::mouseGrabber();
extern void _ZN7QWidget12mouseGrabberEv();
  // proto:  void QWidget::setFixedWidth(int w);
extern void _ZN7QWidget13setFixedWidthEi(void* qthis, int arg0);
  // proto:  void QWidget::addAction(QAction * action);
extern void _ZN7QWidget9addActionEP7QAction(void* qthis, void* arg0);
  // proto:  void QWidget::setDisabled(bool );
extern void _ZN7QWidget11setDisabledEb(void* qthis, bool arg0);
  // proto:  QIcon QWidget::windowIcon();
extern void _ZNK7QWidget10windowIconEv(void* qthis);
  // proto:  void QWidget::setContentsMargins(int left, int top, int right, int bottom);
extern void _ZN7QWidget18setContentsMarginsEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  QString QWidget::windowRole();
extern void _ZNK7QWidget10windowRoleEv(void* qthis);
  // proto:  void QWidget::setShortcutAutoRepeat(int id, bool enable);
extern void _ZN7QWidget21setShortcutAutoRepeatEib(void* qthis, int arg0, bool arg1);
  // proto:  void QWidget::showFullScreen();
extern void _ZN7QWidget14showFullScreenEv(void* qthis);
  // proto:  void QWidget::grabMouse();
extern void _ZN7QWidget9grabMouseEv(void* qthis);
  // proto:  void QWidget::setMaximumSize(const QSize & );
extern void _ZN7QWidget14setMaximumSizeERK5QSize(void* qthis, void* arg0);
  // proto:  QPoint QWidget::mapToGlobal(const QPoint & );
extern void _ZNK7QWidget11mapToGlobalERK6QPoint(void* qthis, void* arg0);
  // proto:  QString QWidget::toolTip();
extern void _ZNK7QWidget7toolTipEv(void* qthis);
  // proto:  void QWidget::setWhatsThis(const QString & );
extern void _ZN7QWidget12setWhatsThisERK7QString(void* qthis, void* arg0);
  // proto:  void QWidget::resize(int w, int h);
extern void _ZN7QWidget6resizeEii(void* qthis, int arg0, int arg1);
  // proto:  QWidget * QWidget::parentWidget();
extern void _ZNK7QWidget12parentWidgetEv(void* qthis);
  // proto:  QPoint QWidget::pos();
extern void _ZNK7QWidget3posEv(void* qthis);
  // proto:  void QWidget::setAutoFillBackground(bool enabled);
extern void _ZN7QWidget21setAutoFillBackgroundEb(void* qthis, bool arg0);
  // proto:  bool QWidget::hasFocus();
extern void _ZNK7QWidget8hasFocusEv(void* qthis);
  // proto:  QSize QWidget::baseSize();
extern void _ZNK7QWidget8baseSizeEv(void* qthis);
  // proto:  void QWidget::setMask(const QBitmap & );
extern void _ZN7QWidget7setMaskERK7QBitmap(void* qthis, void* arg0);
  // proto:  void QWidget::ensurePolished();
extern void _ZNK7QWidget14ensurePolishedEv(void* qthis);
  // proto:  void QWidget::setWindowTitle(const QString & );
extern void _ZN7QWidget14setWindowTitleERK7QString(void* qthis, void* arg0);
  // proto:  QWidget * QWidget::window();
extern void _ZNK7QWidget6windowEv(void* qthis);
  // proto:  void QWidget::scroll(int dx, int dy);
extern void _ZN7QWidget6scrollEii(void* qthis, int arg0, int arg1);
  // proto:  void QWidget::releaseShortcut(int id);
extern void _ZN7QWidget15releaseShortcutEi(void* qthis, int arg0);
  // proto:  void QWidget::setToolTipDuration(int msec);
extern void _ZN7QWidget18setToolTipDurationEi(void* qthis, int arg0);
  // proto:  void QWidget::setGeometry(int x, int y, int w, int h);
extern void demth_ZN7QWidget11setGeometryEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  void QWidget::setSizeIncrement(int w, int h);
extern void _ZN7QWidget16setSizeIncrementEii(void* qthis, int arg0, int arg1);
  // proto:  void QWidget::setUpdatesEnabled(bool enable);
extern void _ZN7QWidget17setUpdatesEnabledEb(void* qthis, bool arg0);
  // proto:  void QWidget::lower();
extern void _ZN7QWidget5lowerEv(void* qthis);
  // proto:  void QWidget::setMouseTracking(bool enable);
extern void _ZN7QWidget16setMouseTrackingEb(void* qthis, bool arg0);
  // proto:  void QWidget::setBaseSize(const QSize & );
extern void _ZN7QWidget11setBaseSizeERK5QSize(void* qthis, void* arg0);
  // proto:  void QWidget::hide();
extern void _ZN7QWidget4hideEv(void* qthis);
  // proto:  void QWidget::removeAction(QAction * action);
extern void _ZN7QWidget12removeActionEP7QAction(void* qthis, void* arg0);
  // proto:  void QWidget::setFocusProxy(QWidget * );
extern void _ZN7QWidget13setFocusProxyEPS_(void* qthis, void* arg0);
  // proto:  bool QWidget::close();
extern void _ZN7QWidget5closeEv(void* qthis);
  // proto:  void QWidget::showMinimized();
extern void _ZN7QWidget13showMinimizedEv(void* qthis);
  // proto:  void QWidget::setFixedSize(int w, int h);
extern void _ZN7QWidget12setFixedSizeEii(void* qthis, int arg0, int arg1);
  // proto:  QSize QWidget::minimumSize();
extern void _ZNK7QWidget11minimumSizeEv(void* qthis);
  // proto:  void QWidget::setEnabled(bool );
extern void _ZN7QWidget10setEnabledEb(void* qthis, bool arg0);
  // proto:  int QWidget::maximumHeight();
extern void _ZNK7QWidget13maximumHeightEv(void* qthis);
  // proto:  void QWidget::move(int x, int y);
extern void _ZN7QWidget4moveEii(void* qthis, int arg0, int arg1);
  // proto:  bool QWidget::isAncestorOf(const QWidget * child);
extern void _ZNK7QWidget12isAncestorOfEPKS_(void* qthis, void* arg0);
  // proto:  void QWidget::QWidget(const QWidget & );
extern void* dector_ZN7QWidgetC1ERKS_(void* arg0);
extern void _ZN7QWidgetC1ERKS_(void* qthis, void* arg0);
  // proto:  QCursor QWidget::cursor();
extern void _ZNK7QWidget6cursorEv(void* qthis);
  // proto:  QPoint QWidget::mapFromGlobal(const QPoint & );
extern void _ZNK7QWidget13mapFromGlobalERK6QPoint(void* qthis, void* arg0);
  // proto:  void QWidget::setToolTip(const QString & );
extern void _ZN7QWidget10setToolTipERK7QString(void* qthis, void* arg0);
  // proto:  QSizePolicy QWidget::sizePolicy();
extern void _ZNK7QWidget10sizePolicyEv(void* qthis);
  // proto:  bool QWidget::hasHeightForWidth();
extern void _ZNK7QWidget17hasHeightForWidthEv(void* qthis);
  // proto:  QGraphicsProxyWidget * QWidget::graphicsProxyWidget();
extern void _ZNK7QWidget19graphicsProxyWidgetEv(void* qthis);
  // proto:  QMargins QWidget::contentsMargins();
extern void _ZNK7QWidget15contentsMarginsEv(void* qthis);
  // proto:  QWidget * QWidget::topLevelWidget();
extern void demth_ZNK7QWidget14topLevelWidgetEv(void* qthis);
  // proto:  void QWidget::setLayout(QLayout * );
extern void _ZN7QWidget9setLayoutEP7QLayout(void* qthis, void* arg0);
  // proto:  bool QWidget::underMouse();
extern void _ZNK7QWidget10underMouseEv(void* qthis);
  // proto:  int QWidget::heightForWidth(int );
extern void _ZNK7QWidget14heightForWidthEi(void* qthis, int arg0);
  // proto:  void QWidget::setFont(const QFont & );
extern void _ZN7QWidget7setFontERK5QFont(void* qthis, void* arg0);
  // proto:  QWidget * QWidget::nativeParentWidget();
extern void _ZNK7QWidget18nativeParentWidgetEv(void* qthis);
  // proto:  void QWidget::setLocale(const QLocale & locale);
extern void _ZN7QWidget9setLocaleERK7QLocale(void* qthis, void* arg0);
  // proto:  int QWidget::height();
extern void demth_ZNK7QWidget6heightEv(void* qthis);
  // proto:  void QWidget::setHidden(bool hidden);
extern void _ZN7QWidget9setHiddenEb(void* qthis, bool arg0);
  // proto:  QSize QWidget::size();
extern void _ZNK7QWidget4sizeEv(void* qthis);
  // proto:  int QWidget::maximumWidth();
extern void _ZNK7QWidget12maximumWidthEv(void* qthis);
  // proto:  bool QWidget::isMinimized();
extern void _ZNK7QWidget11isMinimizedEv(void* qthis);
  // proto:  void QWidget::update();
extern void _ZN7QWidget6updateEv(void* qthis);
  // proto:  void QWidget::setCursor(const QCursor & );
extern void _ZN7QWidget9setCursorERK7QCursor(void* qthis, void* arg0);
  // proto:  QStyle * QWidget::style();
extern void _ZNK7QWidget5styleEv(void* qthis);
  // proto:  void QWidget::createWinId();
extern void _ZN7QWidget11createWinIdEv(void* qthis);
  // proto:  void QWidget::setWindowOpacity(qreal level);
extern void _ZN7QWidget16setWindowOpacityEd(void* qthis, double arg0);
  // proto:  bool QWidget::isRightToLeft();
extern void demth_ZNK7QWidget13isRightToLeftEv(void* qthis);
  // proto:  void QWidget::setAccessibleName(const QString & name);
extern void _ZN7QWidget17setAccessibleNameERK7QString(void* qthis, void* arg0);
  // proto:  void QWidget::unsetCursor();
extern void _ZN7QWidget11unsetCursorEv(void* qthis);
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

// class sizeof(QWidget)=1
type QWidget struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _windowIconChanged QWidget_windowIconChanged_signal;
//  _windowTitleChanged QWidget_windowTitleChanged_signal;
//  _customContextMenuRequested QWidget_customContextMenuRequested_signal;
//  _windowIconTextChanged QWidget_windowIconTextChanged_signal;
}

// class sizeof(QWidgetData)=1
type QWidgetData struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QWidget::setGraphicsEffect(QGraphicsEffect * effect);
func (this *QWidget) setGraphicsEffect(args ...interface{}) () {
  // setGraphicsEffect(class QGraphicsEffect *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsEffect{}) // "QGraphicsEffect *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget17setGraphicsEffectEP15QGraphicsEffect
    var arg0 = args[0].(QGraphicsEffect).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setGraphicsEffect", args)
  }

}

  // proto:  QString QWidget::accessibleDescription();
func (this *QWidget) accessibleDescription(args ...interface{}) () {
  // accessibleDescription()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget21accessibleDescriptionEv
  default:
    qtrt.ErrorResolve("QWidget", "accessibleDescription", args)
  }

}

  // proto:  QGraphicsEffect * QWidget::graphicsEffect();
func (this *QWidget) graphicsEffect(args ...interface{}) () {
  // graphicsEffect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14graphicsEffectEv
  default:
    qtrt.ErrorResolve("QWidget", "graphicsEffect", args)
  }

}

  // proto:  bool QWidget::isFullScreen();
func (this *QWidget) isFullScreen(args ...interface{}) () {
  // isFullScreen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget12isFullScreenEv
  default:
    qtrt.ErrorResolve("QWidget", "isFullScreen", args)
  }

}

  // proto:  QSize QWidget::maximumSize();
func (this *QWidget) maximumSize(args ...interface{}) () {
  // maximumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11maximumSizeEv
  default:
    qtrt.ErrorResolve("QWidget", "maximumSize", args)
  }

}

  // proto:  bool QWidget::isEnabledToTLW();
func (this *QWidget) isEnabledToTLW(args ...interface{}) () {
  // isEnabledToTLW()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14isEnabledToTLWEv
  default:
    qtrt.ErrorResolve("QWidget", "isEnabledToTLW", args)
  }

}

  // proto:  void QWidget::setStatusTip(const QString & );
func (this *QWidget) setStatusTip(args ...interface{}) () {
  // setStatusTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget12setStatusTipERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setStatusTip", args)
  }

}

  // proto:  void QWidget::setSizeIncrement(const QSize & );
func (this *QWidget) setSizeIncrement(args ...interface{}) () {
  // setSizeIncrement(const class QSize &)
  // setSizeIncrement(int, int)
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
    // invoke: _ZN7QWidget16setSizeIncrementERK5QSize
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN7QWidget16setSizeIncrementEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QWidget", "setSizeIncrement", args)
  }

}

  // proto:  QWidget * QWidget::focusWidget();
func (this *QWidget) focusWidget(args ...interface{}) () {
  // focusWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11focusWidgetEv
  default:
    qtrt.ErrorResolve("QWidget", "focusWidget", args)
  }

}

  // proto:  bool QWidget::isTopLevel();
func (this *QWidget) isTopLevel(args ...interface{}) () {
  // isTopLevel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget10isTopLevelEv
  default:
    qtrt.ErrorResolve("QWidget", "isTopLevel", args)
  }

}

  // proto:  bool QWidget::acceptDrops();
func (this *QWidget) acceptDrops(args ...interface{}) () {
  // acceptDrops()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11acceptDropsEv
  default:
    qtrt.ErrorResolve("QWidget", "acceptDrops", args)
  }

}

  // proto:  bool QWidget::isWindow();
func (this *QWidget) isWindow(args ...interface{}) () {
  // isWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget8isWindowEv
  default:
    qtrt.ErrorResolve("QWidget", "isWindow", args)
  }

}

  // proto:  void QWidget::setFixedSize(const QSize & );
func (this *QWidget) setFixedSize(args ...interface{}) () {
  // setFixedSize(const class QSize &)
  // setFixedSize(int, int)
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
    // invoke: _ZN7QWidget12setFixedSizeERK5QSize
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN7QWidget12setFixedSizeEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QWidget", "setFixedSize", args)
  }

}

  // proto:  bool QWidget::isVisible();
func (this *QWidget) isVisible(args ...interface{}) () {
  // isVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget9isVisibleEv
  default:
    qtrt.ErrorResolve("QWidget", "isVisible", args)
  }

}

  // proto:  int QWidget::minimumHeight();
func (this *QWidget) minimumHeight(args ...interface{}) () {
  // minimumHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget13minimumHeightEv
  default:
    qtrt.ErrorResolve("QWidget", "minimumHeight", args)
  }

}

  // proto:  QSize QWidget::sizeIncrement();
func (this *QWidget) sizeIncrement(args ...interface{}) () {
  // sizeIncrement()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget13sizeIncrementEv
  default:
    qtrt.ErrorResolve("QWidget", "sizeIncrement", args)
  }

}

  // proto:  void QWidget::repaint(const QRect & );
func (this *QWidget) repaint(args ...interface{}) () {
  // repaint(const class QRect &)
  // repaint(const class QRegion &)
  // repaint()
  // repaint(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[3][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget7repaintERK5QRect
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN7QWidget7repaintERK7QRegion
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
  case 2:
    // invoke: _ZN7QWidget7repaintEv
  case 3:
    // invoke: _ZN7QWidget7repaintEiiii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QWidget", "repaint", args)
  }

}

  // proto:  void QWidget::update(int x, int y, int w, int h);
func (this *QWidget) update(args ...interface{}) () {
  // update(int, int, int, int)
  // update(const class QRect &)
  // update(const class QRegion &)
  // update()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  vtys[3] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget6updateEiiii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  case 1:
    // invoke: _ZN7QWidget6updateERK5QRect
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
  case 2:
    // invoke: _ZN7QWidget6updateERK7QRegion
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
  case 3:
    // invoke: _ZN7QWidget6updateEv
  default:
    qtrt.ErrorResolve("QWidget", "update", args)
  }

}

  // proto:  QString QWidget::windowFilePath();
func (this *QWidget) windowFilePath(args ...interface{}) () {
  // windowFilePath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14windowFilePathEv
  default:
    qtrt.ErrorResolve("QWidget", "windowFilePath", args)
  }

}

  // proto:  void QWidget::clearMask();
func (this *QWidget) clearMask(args ...interface{}) () {
  // clearMask()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget9clearMaskEv
  default:
    qtrt.ErrorResolve("QWidget", "clearMask", args)
  }

}

  // proto:  QPoint QWidget::mapFromParent(const QPoint & );
func (this *QWidget) mapFromParent(args ...interface{}) () {
  // mapFromParent(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget13mapFromParentERK6QPoint
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "mapFromParent", args)
  }

}

  // proto:  QRect QWidget::rect();
func (this *QWidget) rect(args ...interface{}) () {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget4rectEv
  default:
    qtrt.ErrorResolve("QWidget", "rect", args)
  }

}

  // proto:  void QWidget::unsetLayoutDirection();
func (this *QWidget) unsetLayoutDirection(args ...interface{}) () {
  // unsetLayoutDirection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget20unsetLayoutDirectionEv
  default:
    qtrt.ErrorResolve("QWidget", "unsetLayoutDirection", args)
  }

}

  // proto:  void QWidget::setMinimumSize(const QSize & );
func (this *QWidget) setMinimumSize(args ...interface{}) () {
  // setMinimumSize(const class QSize &)
  // setMinimumSize(int, int)
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
    // invoke: _ZN7QWidget14setMinimumSizeERK5QSize
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN7QWidget14setMinimumSizeEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QWidget", "setMinimumSize", args)
  }

}

  // proto:  bool QWidget::isActiveWindow();
func (this *QWidget) isActiveWindow(args ...interface{}) () {
  // isActiveWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14isActiveWindowEv
  default:
    qtrt.ErrorResolve("QWidget", "isActiveWindow", args)
  }

}

  // proto:  void QWidget::grabKeyboard();
func (this *QWidget) grabKeyboard(args ...interface{}) () {
  // grabKeyboard()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget12grabKeyboardEv
  default:
    qtrt.ErrorResolve("QWidget", "grabKeyboard", args)
  }

}

  // proto:  QSize QWidget::frameSize();
func (this *QWidget) frameSize(args ...interface{}) () {
  // frameSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget9frameSizeEv
  default:
    qtrt.ErrorResolve("QWidget", "frameSize", args)
  }

}

  // proto:  void QWidget::setFocus();
func (this *QWidget) setFocus(args ...interface{}) () {
  // setFocus(Qt::FocusReason)
  // setFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "Qt::FocusReason"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget8setFocusEN2Qt11FocusReasonE
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN7QWidget8setFocusEv
  default:
    qtrt.ErrorResolve("QWidget", "setFocus", args)
  }

}

  // proto:  QPoint QWidget::mapToParent(const QPoint & );
func (this *QWidget) mapToParent(args ...interface{}) () {
  // mapToParent(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11mapToParentERK6QPoint
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "mapToParent", args)
  }

}

  // proto:  void QWidget::updateGeometry();
func (this *QWidget) updateGeometry(args ...interface{}) () {
  // updateGeometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget14updateGeometryEv
  default:
    qtrt.ErrorResolve("QWidget", "updateGeometry", args)
  }

}

  // proto:  void QWidget::insertAction(QAction * before, QAction * action);
func (this *QWidget) insertAction(args ...interface{}) () {
  // insertAction(class QAction *, class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"
  vtys[0][1] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget12insertActionEP7QActionS1_
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QAction).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QWidget", "insertAction", args)
  }

}

  // proto:  void QWidget::setWindowRole(const QString & );
func (this *QWidget) setWindowRole(args ...interface{}) () {
  // setWindowRole(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget13setWindowRoleERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setWindowRole", args)
  }

}

  // proto:  int QWidget::toolTipDuration();
func (this *QWidget) toolTipDuration(args ...interface{}) () {
  // toolTipDuration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget15toolTipDurationEv
  default:
    qtrt.ErrorResolve("QWidget", "toolTipDuration", args)
  }

}

  // proto:  void QWidget::setPalette(const QPalette & );
func (this *QWidget) setPalette(args ...interface{}) () {
  // setPalette(const class QPalette &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPalette{}) // "const QPalette &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget10setPaletteERK8QPalette
    var arg0 = args[0].(QPalette).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setPalette", args)
  }

}

  // proto:  void QWidget::setStyleSheet(const QString & styleSheet);
func (this *QWidget) setStyleSheet(args ...interface{}) () {
  // setStyleSheet(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget13setStyleSheetERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setStyleSheet", args)
  }

}

  // proto:  QString QWidget::windowIconText();
func (this *QWidget) windowIconText(args ...interface{}) () {
  // windowIconText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14windowIconTextEv
  default:
    qtrt.ErrorResolve("QWidget", "windowIconText", args)
  }

}

  // proto:  void QWidget::releaseMouse();
func (this *QWidget) releaseMouse(args ...interface{}) () {
  // releaseMouse()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget12releaseMouseEv
  default:
    qtrt.ErrorResolve("QWidget", "releaseMouse", args)
  }

}

  // proto:  bool QWidget::isModal();
func (this *QWidget) isModal(args ...interface{}) () {
  // isModal()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget7isModalEv
  default:
    qtrt.ErrorResolve("QWidget", "isModal", args)
  }

}

  // proto:  void QWidget::setStyle(QStyle * );
func (this *QWidget) setStyle(args ...interface{}) () {
  // setStyle(class QStyle *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyle{}) // "QStyle *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget8setStyleEP6QStyle
    var arg0 = args[0].(QStyle).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setStyle", args)
  }

}

  // proto:  void QWidget::setBaseSize(int basew, int baseh);
func (this *QWidget) setBaseSize(args ...interface{}) () {
  // setBaseSize(int, int)
  // setBaseSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget11setBaseSizeEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN7QWidget11setBaseSizeERK5QSize
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setBaseSize", args)
  }

}

  // proto:  bool QWidget::isWindowModified();
func (this *QWidget) isWindowModified(args ...interface{}) () {
  // isWindowModified()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget16isWindowModifiedEv
  default:
    qtrt.ErrorResolve("QWidget", "isWindowModified", args)
  }

}

  // proto:  const QRect & QWidget::geometry();
func (this *QWidget) geometry(args ...interface{}) () {
  // geometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget8geometryEv
  default:
    qtrt.ErrorResolve("QWidget", "geometry", args)
  }

}

  // proto:  void QWidget::setAccessibleDescription(const QString & description);
func (this *QWidget) setAccessibleDescription(args ...interface{}) () {
  // setAccessibleDescription(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget24setAccessibleDescriptionERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setAccessibleDescription", args)
  }

}

  // proto:  QWidget * QWidget::nextInFocusChain();
func (this *QWidget) nextInFocusChain(args ...interface{}) () {
  // nextInFocusChain()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget16nextInFocusChainEv
  default:
    qtrt.ErrorResolve("QWidget", "nextInFocusChain", args)
  }

}

  // proto: static void QWidget::setTabOrder(QWidget * , QWidget * );
func (this *QWidget) setTabOrder_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWidget", "setTabOrder", args)
  }

}

  // proto:  QRect QWidget::frameGeometry();
func (this *QWidget) frameGeometry(args ...interface{}) () {
  // frameGeometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget13frameGeometryEv
  default:
    qtrt.ErrorResolve("QWidget", "frameGeometry", args)
  }

}

  // proto:  QSize QWidget::sizeHint();
func (this *QWidget) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget8sizeHintEv
  default:
    qtrt.ErrorResolve("QWidget", "sizeHint", args)
  }

}

  // proto:  void QWidget::setMinimumWidth(int minw);
func (this *QWidget) setMinimumWidth(args ...interface{}) () {
  // setMinimumWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget15setMinimumWidthEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setMinimumWidth", args)
  }

}

  // proto:  bool QWidget::isVisibleTo(const QWidget * );
func (this *QWidget) isVisibleTo(args ...interface{}) () {
  // isVisibleTo(const class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11isVisibleToEPKS_
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "isVisibleTo", args)
  }

}

  // proto:  void QWidget::setMaximumSize(int maxw, int maxh);
func (this *QWidget) setMaximumSize(args ...interface{}) () {
  // setMaximumSize(int, int)
  // setMaximumSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget14setMaximumSizeEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN7QWidget14setMaximumSizeERK5QSize
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setMaximumSize", args)
  }

}

  // proto:  bool QWidget::hasMouseTracking();
func (this *QWidget) hasMouseTracking(args ...interface{}) () {
  // hasMouseTracking()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget16hasMouseTrackingEv
  default:
    qtrt.ErrorResolve("QWidget", "hasMouseTracking", args)
  }

}

  // proto:  bool QWidget::isHidden();
func (this *QWidget) isHidden(args ...interface{}) () {
  // isHidden()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget8isHiddenEv
  default:
    qtrt.ErrorResolve("QWidget", "isHidden", args)
  }

}

  // proto:  int QWidget::devType();
func (this *QWidget) devType(args ...interface{}) () {
  // devType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget7devTypeEv
  default:
    qtrt.ErrorResolve("QWidget", "devType", args)
  }

}

  // proto:  QWidget * QWidget::childAt(int x, int y);
func (this *QWidget) childAt(args ...interface{}) () {
  // childAt(int, int)
  // childAt(const class QPoint &)
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
    // invoke: _ZNK7QWidget7childAtEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZNK7QWidget7childAtERK6QPoint
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "childAt", args)
  }

}

  // proto:  void QWidget::activateWindow();
func (this *QWidget) activateWindow(args ...interface{}) () {
  // activateWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget14activateWindowEv
  default:
    qtrt.ErrorResolve("QWidget", "activateWindow", args)
  }

}

  // proto:  QRect QWidget::normalGeometry();
func (this *QWidget) normalGeometry(args ...interface{}) () {
  // normalGeometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14normalGeometryEv
  default:
    qtrt.ErrorResolve("QWidget", "normalGeometry", args)
  }

}

  // proto:  QString QWidget::windowTitle();
func (this *QWidget) windowTitle(args ...interface{}) () {
  // windowTitle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11windowTitleEv
  default:
    qtrt.ErrorResolve("QWidget", "windowTitle", args)
  }

}

  // proto:  void QWidget::grabMouse(const QCursor & );
func (this *QWidget) grabMouse(args ...interface{}) () {
  // grabMouse(const class QCursor &)
  // grabMouse()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCursor{}) // "const QCursor &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget9grabMouseERK7QCursor
    var arg0 = args[0].(QCursor).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN7QWidget9grabMouseEv
  default:
    qtrt.ErrorResolve("QWidget", "grabMouse", args)
  }

}

  // proto:  QPixmap QWidget::grab(const QRect & rectangle);
func (this *QWidget) grab(args ...interface{}) () {
  // grab(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget4grabERK5QRect
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "grab", args)
  }

}

  // proto:  void QWidget::setVisible(bool visible);
func (this *QWidget) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget10setVisibleEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setVisible", args)
  }

}

  // proto:  bool QWidget::isEnabledTo(const QWidget * );
func (this *QWidget) isEnabledTo(args ...interface{}) () {
  // isEnabledTo(const class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11isEnabledToEPKS_
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "isEnabledTo", args)
  }

}

  // proto:  bool QWidget::isLeftToRight();
func (this *QWidget) isLeftToRight(args ...interface{}) () {
  // isLeftToRight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget13isLeftToRightEv
  default:
    qtrt.ErrorResolve("QWidget", "isLeftToRight", args)
  }

}

  // proto:  void QWidget::setGeometry(const QRect & );
func (this *QWidget) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRect &)
  // setGeometry(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget11setGeometryERK5QRect
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN7QWidget11setGeometryEiiii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QWidget", "setGeometry", args)
  }

}

  // proto:  void QWidget::unsetLocale();
func (this *QWidget) unsetLocale(args ...interface{}) () {
  // unsetLocale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget11unsetLocaleEv
  default:
    qtrt.ErrorResolve("QWidget", "unsetLocale", args)
  }

}

  // proto:  void QWidget::showNormal();
func (this *QWidget) showNormal(args ...interface{}) () {
  // showNormal()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget10showNormalEv
  default:
    qtrt.ErrorResolve("QWidget", "showNormal", args)
  }

}

  // proto:  int QWidget::y();
func (this *QWidget) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget1yEv
  default:
    qtrt.ErrorResolve("QWidget", "y", args)
  }

}

  // proto:  int QWidget::width();
func (this *QWidget) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget5widthEv
  default:
    qtrt.ErrorResolve("QWidget", "width", args)
  }

}

  // proto:  bool QWidget::isMaximized();
func (this *QWidget) isMaximized(args ...interface{}) () {
  // isMaximized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11isMaximizedEv
  default:
    qtrt.ErrorResolve("QWidget", "isMaximized", args)
  }

}

  // proto:  void QWidget::resize(const QSize & );
func (this *QWidget) resize(args ...interface{}) () {
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
    // invoke: _ZN7QWidget6resizeERK5QSize
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN7QWidget6resizeEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QWidget", "resize", args)
  }

}

  // proto:  QWindow * QWidget::windowHandle();
func (this *QWidget) windowHandle(args ...interface{}) () {
  // windowHandle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget12windowHandleEv
  default:
    qtrt.ErrorResolve("QWidget", "windowHandle", args)
  }

}

  // proto:  QString QWidget::accessibleName();
func (this *QWidget) accessibleName(args ...interface{}) () {
  // accessibleName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14accessibleNameEv
  default:
    qtrt.ErrorResolve("QWidget", "accessibleName", args)
  }

}

  // proto:  void QWidget::setContentsMargins(const QMargins & margins);
func (this *QWidget) setContentsMargins(args ...interface{}) () {
  // setContentsMargins(const class QMargins &)
  // setContentsMargins(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMargins{}) // "const QMargins &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget18setContentsMarginsERK8QMargins
    var arg0 = args[0].(QMargins).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN7QWidget18setContentsMarginsEiiii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QWidget", "setContentsMargins", args)
  }

}

  // proto:  QByteArray QWidget::saveGeometry();
func (this *QWidget) saveGeometry(args ...interface{}) () {
  // saveGeometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget12saveGeometryEv
  default:
    qtrt.ErrorResolve("QWidget", "saveGeometry", args)
  }

}

  // proto:  bool QWidget::isEnabled();
func (this *QWidget) isEnabled(args ...interface{}) () {
  // isEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget9isEnabledEv
  default:
    qtrt.ErrorResolve("QWidget", "isEnabled", args)
  }

}

  // proto:  void QWidget::setFixedHeight(int h);
func (this *QWidget) setFixedHeight(args ...interface{}) () {
  // setFixedHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget14setFixedHeightEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setFixedHeight", args)
  }

}

  // proto:  QRegion QWidget::mask();
func (this *QWidget) mask(args ...interface{}) () {
  // mask()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget4maskEv
  default:
    qtrt.ErrorResolve("QWidget", "mask", args)
  }

}

  // proto:  void QWidget::stackUnder(QWidget * );
func (this *QWidget) stackUnder(args ...interface{}) () {
  // stackUnder(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget10stackUnderEPS_
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "stackUnder", args)
  }

}

  // proto:  QPaintEngine * QWidget::paintEngine();
func (this *QWidget) paintEngine(args ...interface{}) () {
  // paintEngine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11paintEngineEv
  default:
    qtrt.ErrorResolve("QWidget", "paintEngine", args)
  }

}

  // proto:  void QWidget::setAcceptDrops(bool on);
func (this *QWidget) setAcceptDrops(args ...interface{}) () {
  // setAcceptDrops(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget14setAcceptDropsEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setAcceptDrops", args)
  }

}

  // proto:  void QWidget::move(const QPoint & );
func (this *QWidget) move_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWidget", "move", args)
  }

}

  // proto:  QList<QAction *> QWidget::actions();
func (this *QWidget) actions(args ...interface{}) () {
  // actions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget7actionsEv
  default:
    qtrt.ErrorResolve("QWidget", "actions", args)
  }

}

  // proto:  void QWidget::show();
func (this *QWidget) show(args ...interface{}) () {
  // show()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget4showEv
  default:
    qtrt.ErrorResolve("QWidget", "show", args)
  }

}

  // proto: static QWidget * QWidget::keyboardGrabber();
func (this *QWidget) keyboardGrabber_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWidget", "keyboardGrabber", args)
  }

}

  // proto:  QPoint QWidget::mapTo(const QWidget * , const QPoint & );
func (this *QWidget) mapTo(args ...interface{}) () {
  // mapTo(const class QWidget *, const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"
  vtys[0][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget5mapToEPKS_RK6QPoint
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPoint).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QWidget", "mapTo", args)
  }

}

  // proto:  int QWidget::minimumWidth();
func (this *QWidget) minimumWidth(args ...interface{}) () {
  // minimumWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget12minimumWidthEv
  default:
    qtrt.ErrorResolve("QWidget", "minimumWidth", args)
  }

}

  // proto:  QFontInfo QWidget::fontInfo();
func (this *QWidget) fontInfo(args ...interface{}) () {
  // fontInfo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget8fontInfoEv
  default:
    qtrt.ErrorResolve("QWidget", "fontInfo", args)
  }

}

  // proto:  bool QWidget::autoFillBackground();
func (this *QWidget) autoFillBackground(args ...interface{}) () {
  // autoFillBackground()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget18autoFillBackgroundEv
  default:
    qtrt.ErrorResolve("QWidget", "autoFillBackground", args)
  }

}

  // proto:  void QWidget::scroll(int dx, int dy, const QRect & );
func (this *QWidget) scroll(args ...interface{}) () {
  // scroll(int, int, const class QRect &)
  // scroll(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget6scrollEiiRK5QRect
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QRect).qclsinst
    if false {fmt.Println(arg2)}
  case 1:
    // invoke: _ZN7QWidget6scrollEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QWidget", "scroll", args)
  }

}

  // proto:  QFontMetrics QWidget::fontMetrics();
func (this *QWidget) fontMetrics(args ...interface{}) () {
  // fontMetrics()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11fontMetricsEv
  default:
    qtrt.ErrorResolve("QWidget", "fontMetrics", args)
  }

}

  // proto:  void QWidget::adjustSize();
func (this *QWidget) adjustSize(args ...interface{}) () {
  // adjustSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget10adjustSizeEv
  default:
    qtrt.ErrorResolve("QWidget", "adjustSize", args)
  }

}

  // proto:  QWidget * QWidget::previousInFocusChain();
func (this *QWidget) previousInFocusChain(args ...interface{}) () {
  // previousInFocusChain()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget20previousInFocusChainEv
  default:
    qtrt.ErrorResolve("QWidget", "previousInFocusChain", args)
  }

}

  // proto:  bool QWidget::updatesEnabled();
func (this *QWidget) updatesEnabled(args ...interface{}) () {
  // updatesEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14updatesEnabledEv
  default:
    qtrt.ErrorResolve("QWidget", "updatesEnabled", args)
  }

}

  // proto:  void QWidget::setMaximumHeight(int maxh);
func (this *QWidget) setMaximumHeight(args ...interface{}) () {
  // setMaximumHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget16setMaximumHeightEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setMaximumHeight", args)
  }

}

  // proto:  void QWidget::showMaximized();
func (this *QWidget) showMaximized(args ...interface{}) () {
  // showMaximized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget13showMaximizedEv
  default:
    qtrt.ErrorResolve("QWidget", "showMaximized", args)
  }

}

  // proto:  QPoint QWidget::mapFrom(const QWidget * , const QPoint & );
func (this *QWidget) mapFrom(args ...interface{}) () {
  // mapFrom(const class QWidget *, const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"
  vtys[0][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget7mapFromEPKS_RK6QPoint
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPoint).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QWidget", "mapFrom", args)
  }

}

  // proto:  int QWidget::x();
func (this *QWidget) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget1xEv
  default:
    qtrt.ErrorResolve("QWidget", "x", args)
  }

}

  // proto:  void QWidget::clearFocus();
func (this *QWidget) clearFocus(args ...interface{}) () {
  // clearFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget10clearFocusEv
  default:
    qtrt.ErrorResolve("QWidget", "clearFocus", args)
  }

}

  // proto: static QWidget * QWidget::find(WId );
func (this *QWidget) find_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWidget", "find", args)
  }

}

  // proto:  const QPalette & QWidget::palette();
func (this *QWidget) palette(args ...interface{}) () {
  // palette()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget7paletteEv
  default:
    qtrt.ErrorResolve("QWidget", "palette", args)
  }

}

  // proto:  void QWidget::setSizePolicy(QSizePolicy );
func (this *QWidget) setSizePolicy(args ...interface{}) () {
  // setSizePolicy(class QSizePolicy::Policy, class QSizePolicy::Policy)
  // setSizePolicy(class QSizePolicy)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "QSizePolicy::Policy"
  vtys[0][1] = qtrt.Int32Ty(false) // "QSizePolicy::Policy"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QSizePolicy{}) // "QSizePolicy"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget13setSizePolicyEN11QSizePolicy6PolicyES1_
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN7QWidget13setSizePolicyE11QSizePolicy
    var arg0 = args[0].(QSizePolicy).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setSizePolicy", args)
  }

}

  // proto:  void QWidget::setMask(const QRegion & );
func (this *QWidget) setMask(args ...interface{}) () {
  // setMask(const class QRegion &)
  // setMask(const class QBitmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QBitmap{}) // "const QBitmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget7setMaskERK7QRegion
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN7QWidget7setMaskERK7QBitmap
    var arg0 = args[0].(QBitmap).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setMask", args)
  }

}

  // proto:  void QWidget::setMaximumWidth(int maxw);
func (this *QWidget) setMaximumWidth(args ...interface{}) () {
  // setMaximumWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget15setMaximumWidthEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setMaximumWidth", args)
  }

}

  // proto:  void QWidget::setWindowIconText(const QString & );
func (this *QWidget) setWindowIconText(args ...interface{}) () {
  // setWindowIconText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget17setWindowIconTextERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setWindowIconText", args)
  }

}

  // proto:  void QWidget::setWindowIcon(const QIcon & icon);
func (this *QWidget) setWindowIcon(args ...interface{}) () {
  // setWindowIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget13setWindowIconERK5QIcon
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setWindowIcon", args)
  }

}

  // proto:  void QWidget::~QWidget();
func (this *QWidget) FreeQWidget(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWidget", "~QWidget", args)
  }

}

  // proto:  void QWidget::getContentsMargins(int * left, int * top, int * right, int * bottom);
func (this *QWidget) getContentsMargins(args ...interface{}) () {
  // getContentsMargins(int *, int *, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget18getContentsMarginsEPiS0_S0_S0_
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QWidget", "getContentsMargins", args)
  }

}

  // proto:  QSize QWidget::minimumSizeHint();
func (this *QWidget) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget15minimumSizeHintEv
  default:
    qtrt.ErrorResolve("QWidget", "minimumSizeHint", args)
  }

}

  // proto:  void QWidget::setWindowModified(bool );
func (this *QWidget) setWindowModified(args ...interface{}) () {
  // setWindowModified(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget17setWindowModifiedEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setWindowModified", args)
  }

}

  // proto:  bool QWidget::restoreGeometry(const QByteArray & geometry);
func (this *QWidget) restoreGeometry(args ...interface{}) () {
  // restoreGeometry(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget15restoreGeometryERK10QByteArray
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "restoreGeometry", args)
  }

}

  // proto:  QLayout * QWidget::layout();
func (this *QWidget) layout(args ...interface{}) () {
  // layout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget6layoutEv
  default:
    qtrt.ErrorResolve("QWidget", "layout", args)
  }

}

  // proto:  QRect QWidget::contentsRect();
func (this *QWidget) contentsRect(args ...interface{}) () {
  // contentsRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget12contentsRectEv
  default:
    qtrt.ErrorResolve("QWidget", "contentsRect", args)
  }

}

  // proto:  QBackingStore * QWidget::backingStore();
func (this *QWidget) backingStore(args ...interface{}) () {
  // backingStore()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget12backingStoreEv
  default:
    qtrt.ErrorResolve("QWidget", "backingStore", args)
  }

}

  // proto:  QWidget * QWidget::focusProxy();
func (this *QWidget) focusProxy(args ...interface{}) () {
  // focusProxy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget10focusProxyEv
  default:
    qtrt.ErrorResolve("QWidget", "focusProxy", args)
  }

}

  // proto:  QString QWidget::styleSheet();
func (this *QWidget) styleSheet(args ...interface{}) () {
  // styleSheet()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget10styleSheetEv
  default:
    qtrt.ErrorResolve("QWidget", "styleSheet", args)
  }

}

  // proto:  QString QWidget::whatsThis();
func (this *QWidget) whatsThis(args ...interface{}) () {
  // whatsThis()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget9whatsThisEv
  default:
    qtrt.ErrorResolve("QWidget", "whatsThis", args)
  }

}

  // proto:  const QFont & QWidget::font();
func (this *QWidget) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget4fontEv
  default:
    qtrt.ErrorResolve("QWidget", "font", args)
  }

}

  // proto:  const QMetaObject * QWidget::metaObject();
func (this *QWidget) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget10metaObjectEv
  default:
    qtrt.ErrorResolve("QWidget", "metaObject", args)
  }

}

  // proto:  void QWidget::setMinimumHeight(int minh);
func (this *QWidget) setMinimumHeight(args ...interface{}) () {
  // setMinimumHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget16setMinimumHeightEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setMinimumHeight", args)
  }

}

  // proto:  qreal QWidget::windowOpacity();
func (this *QWidget) windowOpacity(args ...interface{}) () {
  // windowOpacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget13windowOpacityEv
  default:
    qtrt.ErrorResolve("QWidget", "windowOpacity", args)
  }

}

  // proto:  QRegion QWidget::childrenRegion();
func (this *QWidget) childrenRegion(args ...interface{}) () {
  // childrenRegion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14childrenRegionEv
  default:
    qtrt.ErrorResolve("QWidget", "childrenRegion", args)
  }

}

  // proto:  void QWidget::setWindowFilePath(const QString & filePath);
func (this *QWidget) setWindowFilePath(args ...interface{}) () {
  // setWindowFilePath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget17setWindowFilePathERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setWindowFilePath", args)
  }

}

  // proto:  void QWidget::setShortcutEnabled(int id, bool enable);
func (this *QWidget) setShortcutEnabled(args ...interface{}) () {
  // setShortcutEnabled(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget18setShortcutEnabledEib
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int8_t(args[1].(int8))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QWidget", "setShortcutEnabled", args)
  }

}

  // proto:  void QWidget::raise();
func (this *QWidget) raise(args ...interface{}) () {
  // raise()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget5raiseEv
  default:
    qtrt.ErrorResolve("QWidget", "raise", args)
  }

}

  // proto:  QString QWidget::statusTip();
func (this *QWidget) statusTip(args ...interface{}) () {
  // statusTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget9statusTipEv
  default:
    qtrt.ErrorResolve("QWidget", "statusTip", args)
  }

}

  // proto:  QRect QWidget::childrenRect();
func (this *QWidget) childrenRect(args ...interface{}) () {
  // childrenRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget12childrenRectEv
  default:
    qtrt.ErrorResolve("QWidget", "childrenRect", args)
  }

}

  // proto:  void QWidget::setParent(QWidget * parent);
func (this *QWidget) setParent(args ...interface{}) () {
  // setParent(class QWidget *, Qt::WindowFlags)
  // setParent(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = qtrt.Int64Ty(false) // "Qt::WindowFlags"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget9setParentEPS_6QFlagsIN2Qt10WindowTypeEE
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN7QWidget9setParentEPS_
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setParent", args)
  }

}

  // proto:  QRegion QWidget::visibleRegion();
func (this *QWidget) visibleRegion(args ...interface{}) () {
  // visibleRegion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget13visibleRegionEv
  default:
    qtrt.ErrorResolve("QWidget", "visibleRegion", args)
  }

}

  // proto:  QLocale QWidget::locale();
func (this *QWidget) locale(args ...interface{}) () {
  // locale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget6localeEv
  default:
    qtrt.ErrorResolve("QWidget", "locale", args)
  }

}

  // proto:  void QWidget::releaseKeyboard();
func (this *QWidget) releaseKeyboard(args ...interface{}) () {
  // releaseKeyboard()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget15releaseKeyboardEv
  default:
    qtrt.ErrorResolve("QWidget", "releaseKeyboard", args)
  }

}

  // proto: static QWidget * QWidget::mouseGrabber();
func (this *QWidget) mouseGrabber_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWidget", "mouseGrabber", args)
  }

}

  // proto:  void QWidget::setFixedWidth(int w);
func (this *QWidget) setFixedWidth(args ...interface{}) () {
  // setFixedWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget13setFixedWidthEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setFixedWidth", args)
  }

}

  // proto:  void QWidget::addAction(QAction * action);
func (this *QWidget) addAction(args ...interface{}) () {
  // addAction(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget9addActionEP7QAction
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "addAction", args)
  }

}

  // proto:  void QWidget::setDisabled(bool );
func (this *QWidget) setDisabled(args ...interface{}) () {
  // setDisabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget11setDisabledEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setDisabled", args)
  }

}

  // proto:  QIcon QWidget::windowIcon();
func (this *QWidget) windowIcon(args ...interface{}) () {
  // windowIcon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget10windowIconEv
  default:
    qtrt.ErrorResolve("QWidget", "windowIcon", args)
  }

}

  // proto:  QString QWidget::windowRole();
func (this *QWidget) windowRole(args ...interface{}) () {
  // windowRole()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget10windowRoleEv
  default:
    qtrt.ErrorResolve("QWidget", "windowRole", args)
  }

}

  // proto:  void QWidget::setShortcutAutoRepeat(int id, bool enable);
func (this *QWidget) setShortcutAutoRepeat(args ...interface{}) () {
  // setShortcutAutoRepeat(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget21setShortcutAutoRepeatEib
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int8_t(args[1].(int8))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QWidget", "setShortcutAutoRepeat", args)
  }

}

  // proto:  void QWidget::showFullScreen();
func (this *QWidget) showFullScreen(args ...interface{}) () {
  // showFullScreen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget14showFullScreenEv
  default:
    qtrt.ErrorResolve("QWidget", "showFullScreen", args)
  }

}

  // proto:  QPoint QWidget::mapToGlobal(const QPoint & );
func (this *QWidget) mapToGlobal(args ...interface{}) () {
  // mapToGlobal(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11mapToGlobalERK6QPoint
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "mapToGlobal", args)
  }

}

  // proto:  QString QWidget::toolTip();
func (this *QWidget) toolTip(args ...interface{}) () {
  // toolTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget7toolTipEv
  default:
    qtrt.ErrorResolve("QWidget", "toolTip", args)
  }

}

  // proto:  void QWidget::setWhatsThis(const QString & );
func (this *QWidget) setWhatsThis(args ...interface{}) () {
  // setWhatsThis(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget12setWhatsThisERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setWhatsThis", args)
  }

}

  // proto:  QWidget * QWidget::parentWidget();
func (this *QWidget) parentWidget(args ...interface{}) () {
  // parentWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget12parentWidgetEv
  default:
    qtrt.ErrorResolve("QWidget", "parentWidget", args)
  }

}

  // proto:  QPoint QWidget::pos();
func (this *QWidget) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget3posEv
  default:
    qtrt.ErrorResolve("QWidget", "pos", args)
  }

}

  // proto:  void QWidget::setAutoFillBackground(bool enabled);
func (this *QWidget) setAutoFillBackground(args ...interface{}) () {
  // setAutoFillBackground(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget21setAutoFillBackgroundEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setAutoFillBackground", args)
  }

}

  // proto:  bool QWidget::hasFocus();
func (this *QWidget) hasFocus(args ...interface{}) () {
  // hasFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget8hasFocusEv
  default:
    qtrt.ErrorResolve("QWidget", "hasFocus", args)
  }

}

  // proto:  QSize QWidget::baseSize();
func (this *QWidget) baseSize(args ...interface{}) () {
  // baseSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget8baseSizeEv
  default:
    qtrt.ErrorResolve("QWidget", "baseSize", args)
  }

}

  // proto:  void QWidget::ensurePolished();
func (this *QWidget) ensurePolished(args ...interface{}) () {
  // ensurePolished()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14ensurePolishedEv
  default:
    qtrt.ErrorResolve("QWidget", "ensurePolished", args)
  }

}

  // proto:  void QWidget::setWindowTitle(const QString & );
func (this *QWidget) setWindowTitle(args ...interface{}) () {
  // setWindowTitle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget14setWindowTitleERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setWindowTitle", args)
  }

}

  // proto:  QWidget * QWidget::window();
func (this *QWidget) window(args ...interface{}) () {
  // window()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget6windowEv
  default:
    qtrt.ErrorResolve("QWidget", "window", args)
  }

}

  // proto:  void QWidget::releaseShortcut(int id);
func (this *QWidget) releaseShortcut(args ...interface{}) () {
  // releaseShortcut(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget15releaseShortcutEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "releaseShortcut", args)
  }

}

  // proto:  void QWidget::setToolTipDuration(int msec);
func (this *QWidget) setToolTipDuration(args ...interface{}) () {
  // setToolTipDuration(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget18setToolTipDurationEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setToolTipDuration", args)
  }

}

  // proto:  void QWidget::setUpdatesEnabled(bool enable);
func (this *QWidget) setUpdatesEnabled(args ...interface{}) () {
  // setUpdatesEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget17setUpdatesEnabledEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setUpdatesEnabled", args)
  }

}

  // proto:  void QWidget::lower();
func (this *QWidget) lower(args ...interface{}) () {
  // lower()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget5lowerEv
  default:
    qtrt.ErrorResolve("QWidget", "lower", args)
  }

}

  // proto:  void QWidget::setMouseTracking(bool enable);
func (this *QWidget) setMouseTracking(args ...interface{}) () {
  // setMouseTracking(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget16setMouseTrackingEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setMouseTracking", args)
  }

}

  // proto:  void QWidget::hide();
func (this *QWidget) hide(args ...interface{}) () {
  // hide()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget4hideEv
  default:
    qtrt.ErrorResolve("QWidget", "hide", args)
  }

}

  // proto:  void QWidget::removeAction(QAction * action);
func (this *QWidget) removeAction(args ...interface{}) () {
  // removeAction(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget12removeActionEP7QAction
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "removeAction", args)
  }

}

  // proto:  void QWidget::setFocusProxy(QWidget * );
func (this *QWidget) setFocusProxy(args ...interface{}) () {
  // setFocusProxy(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget13setFocusProxyEPS_
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setFocusProxy", args)
  }

}

  // proto:  bool QWidget::close();
func (this *QWidget) close(args ...interface{}) () {
  // close()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget5closeEv
  default:
    qtrt.ErrorResolve("QWidget", "close", args)
  }

}

  // proto:  void QWidget::showMinimized();
func (this *QWidget) showMinimized(args ...interface{}) () {
  // showMinimized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget13showMinimizedEv
  default:
    qtrt.ErrorResolve("QWidget", "showMinimized", args)
  }

}

  // proto:  QSize QWidget::minimumSize();
func (this *QWidget) minimumSize(args ...interface{}) () {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11minimumSizeEv
  default:
    qtrt.ErrorResolve("QWidget", "minimumSize", args)
  }

}

  // proto:  void QWidget::setEnabled(bool );
func (this *QWidget) setEnabled(args ...interface{}) () {
  // setEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget10setEnabledEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setEnabled", args)
  }

}

  // proto:  int QWidget::maximumHeight();
func (this *QWidget) maximumHeight(args ...interface{}) () {
  // maximumHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget13maximumHeightEv
  default:
    qtrt.ErrorResolve("QWidget", "maximumHeight", args)
  }

}

  // proto:  bool QWidget::isAncestorOf(const QWidget * child);
func (this *QWidget) isAncestorOf(args ...interface{}) () {
  // isAncestorOf(const class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget12isAncestorOfEPKS_
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "isAncestorOf", args)
  }

}

  // proto:  void QWidget::QWidget(const QWidget & );
func NewQWidget(args ...interface{}) QWidget {
  return QWidget{}
}

  // proto:  QCursor QWidget::cursor();
func (this *QWidget) cursor(args ...interface{}) () {
  // cursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget6cursorEv
  default:
    qtrt.ErrorResolve("QWidget", "cursor", args)
  }

}

  // proto:  QPoint QWidget::mapFromGlobal(const QPoint & );
func (this *QWidget) mapFromGlobal(args ...interface{}) () {
  // mapFromGlobal(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget13mapFromGlobalERK6QPoint
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "mapFromGlobal", args)
  }

}

  // proto:  void QWidget::setToolTip(const QString & );
func (this *QWidget) setToolTip(args ...interface{}) () {
  // setToolTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget10setToolTipERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setToolTip", args)
  }

}

  // proto:  QSizePolicy QWidget::sizePolicy();
func (this *QWidget) sizePolicy(args ...interface{}) () {
  // sizePolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget10sizePolicyEv
  default:
    qtrt.ErrorResolve("QWidget", "sizePolicy", args)
  }

}

  // proto:  bool QWidget::hasHeightForWidth();
func (this *QWidget) hasHeightForWidth(args ...interface{}) () {
  // hasHeightForWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget17hasHeightForWidthEv
  default:
    qtrt.ErrorResolve("QWidget", "hasHeightForWidth", args)
  }

}

  // proto:  QGraphicsProxyWidget * QWidget::graphicsProxyWidget();
func (this *QWidget) graphicsProxyWidget(args ...interface{}) () {
  // graphicsProxyWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget19graphicsProxyWidgetEv
  default:
    qtrt.ErrorResolve("QWidget", "graphicsProxyWidget", args)
  }

}

  // proto:  QMargins QWidget::contentsMargins();
func (this *QWidget) contentsMargins(args ...interface{}) () {
  // contentsMargins()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget15contentsMarginsEv
  default:
    qtrt.ErrorResolve("QWidget", "contentsMargins", args)
  }

}

  // proto:  QWidget * QWidget::topLevelWidget();
func (this *QWidget) topLevelWidget(args ...interface{}) () {
  // topLevelWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14topLevelWidgetEv
  default:
    qtrt.ErrorResolve("QWidget", "topLevelWidget", args)
  }

}

  // proto:  void QWidget::setLayout(QLayout * );
func (this *QWidget) setLayout(args ...interface{}) () {
  // setLayout(class QLayout *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLayout{}) // "QLayout *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget9setLayoutEP7QLayout
    var arg0 = args[0].(QLayout).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setLayout", args)
  }

}

  // proto:  bool QWidget::underMouse();
func (this *QWidget) underMouse(args ...interface{}) () {
  // underMouse()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget10underMouseEv
  default:
    qtrt.ErrorResolve("QWidget", "underMouse", args)
  }

}

  // proto:  int QWidget::heightForWidth(int );
func (this *QWidget) heightForWidth(args ...interface{}) () {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14heightForWidthEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "heightForWidth", args)
  }

}

  // proto:  void QWidget::setFont(const QFont & );
func (this *QWidget) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget7setFontERK5QFont
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setFont", args)
  }

}

  // proto:  QWidget * QWidget::nativeParentWidget();
func (this *QWidget) nativeParentWidget(args ...interface{}) () {
  // nativeParentWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget18nativeParentWidgetEv
  default:
    qtrt.ErrorResolve("QWidget", "nativeParentWidget", args)
  }

}

  // proto:  void QWidget::setLocale(const QLocale & locale);
func (this *QWidget) setLocale(args ...interface{}) () {
  // setLocale(const class QLocale &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLocale{}) // "const QLocale &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget9setLocaleERK7QLocale
    var arg0 = args[0].(QLocale).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setLocale", args)
  }

}

  // proto:  int QWidget::height();
func (this *QWidget) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget6heightEv
  default:
    qtrt.ErrorResolve("QWidget", "height", args)
  }

}

  // proto:  void QWidget::setHidden(bool hidden);
func (this *QWidget) setHidden(args ...interface{}) () {
  // setHidden(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget9setHiddenEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setHidden", args)
  }

}

  // proto:  QSize QWidget::size();
func (this *QWidget) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget4sizeEv
  default:
    qtrt.ErrorResolve("QWidget", "size", args)
  }

}

  // proto:  int QWidget::maximumWidth();
func (this *QWidget) maximumWidth(args ...interface{}) () {
  // maximumWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget12maximumWidthEv
  default:
    qtrt.ErrorResolve("QWidget", "maximumWidth", args)
  }

}

  // proto:  bool QWidget::isMinimized();
func (this *QWidget) isMinimized(args ...interface{}) () {
  // isMinimized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11isMinimizedEv
  default:
    qtrt.ErrorResolve("QWidget", "isMinimized", args)
  }

}

  // proto:  void QWidget::setCursor(const QCursor & );
func (this *QWidget) setCursor(args ...interface{}) () {
  // setCursor(const class QCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCursor{}) // "const QCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget9setCursorERK7QCursor
    var arg0 = args[0].(QCursor).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setCursor", args)
  }

}

  // proto:  QStyle * QWidget::style();
func (this *QWidget) style(args ...interface{}) () {
  // style()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget5styleEv
  default:
    qtrt.ErrorResolve("QWidget", "style", args)
  }

}

  // proto:  void QWidget::createWinId();
func (this *QWidget) createWinId(args ...interface{}) () {
  // createWinId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget11createWinIdEv
  default:
    qtrt.ErrorResolve("QWidget", "createWinId", args)
  }

}

  // proto:  void QWidget::setWindowOpacity(qreal level);
func (this *QWidget) setWindowOpacity(args ...interface{}) () {
  // setWindowOpacity(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget16setWindowOpacityEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setWindowOpacity", args)
  }

}

  // proto:  bool QWidget::isRightToLeft();
func (this *QWidget) isRightToLeft(args ...interface{}) () {
  // isRightToLeft()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget13isRightToLeftEv
  default:
    qtrt.ErrorResolve("QWidget", "isRightToLeft", args)
  }

}

  // proto:  void QWidget::setAccessibleName(const QString & name);
func (this *QWidget) setAccessibleName(args ...interface{}) () {
  // setAccessibleName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget17setAccessibleNameERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QWidget", "setAccessibleName", args)
  }

}

  // proto:  void QWidget::unsetCursor();
func (this *QWidget) unsetCursor(args ...interface{}) () {
  // unsetCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget11unsetCursorEv
  default:
    qtrt.ErrorResolve("QWidget", "unsetCursor", args)
  }

}

// <= body block end

