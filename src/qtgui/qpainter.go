package qtgui
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
// src-file: /QtGui/qpainter.h
// dst-file: /src/gui/qpainter.go
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
  // proto:  void QPainter::restore();
extern void C_ZN8QPainter7restoreEv(void* qthis); // 4
  // proto:  QRect QPainter::boundingRect(const QRect & rect, int flags, const QString & text);
extern void* C_ZN8QPainter12boundingRectERK5QRectiRK7QString(void* qthis, void* arg0, int32_t arg1, void* arg2); // 4
  // proto:  QRectF QPainter::boundingRect(const QRectF & rect, int flags, const QString & text);
extern void* C_ZN8QPainter12boundingRectERK6QRectFiRK7QString(void* qthis, void* arg0, int32_t arg1, void* arg2); // 4
  // proto:  QRect QPainter::boundingRect(int x, int y, int w, int h, int flags, const QString & text);
extern void* C_ZN8QPainter12boundingRectEiiiiiRK7QString(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4, void* arg5); // 2
  // proto:  QRectF QPainter::boundingRect(const QRectF & rect, const QString & text, const QTextOption & o);
extern void* C_ZN8QPainter12boundingRectERK6QRectFRK7QStringRK11QTextOption(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QPainter::drawLine(int x1, int y1, int x2, int y2);
extern void C_ZN8QPainter8drawLineEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  void QPainter::drawLine(const QPoint & p1, const QPoint & p2);
extern void C_ZN8QPainter8drawLineERK6QPointS2_(void* qthis, void* arg0, void* arg1); // 2
  // proto:  void QPainter::drawLine(const QLineF & line);
extern void C_ZN8QPainter8drawLineERK6QLineF(void* qthis, void* arg0); // 2
  // proto:  void QPainter::drawLine(const QPointF & p1, const QPointF & p2);
extern void C_ZN8QPainter8drawLineERK7QPointFS2_(void* qthis, void* arg0, void* arg1); // 2
  // proto:  void QPainter::drawLine(const QLine & line);
extern void C_ZN8QPainter8drawLineERK5QLine(void* qthis, void* arg0); // 2
  // proto:  void QPainter::drawRects(const QRect * rects, int rectCount);
extern void C_ZN8QPainter9drawRectsEPK5QRecti(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QPainter::drawRects(const QRectF * rects, int rectCount);
extern void C_ZN8QPainter9drawRectsEPK6QRectFi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QPainter::setPen(const QColor & color);
extern void C_ZN8QPainter6setPenERK6QColor(void* qthis, void* arg0); // 4
  // proto:  void QPainter::setPen(const QPen & pen);
extern void C_ZN8QPainter6setPenERK4QPen(void* qthis, void* arg0); // 4
  // proto:  const QFont & QPainter::font();
extern void* C_ZNK8QPainter4fontEv(void* qthis); // 4
  // proto:  QPoint QPainter::brushOrigin();
extern void* C_ZNK8QPainter11brushOriginEv(void* qthis); // 4
  // proto:  void QPainter::drawPixmap(const QRect & r, const QPixmap & pm);
extern void C_ZN8QPainter10drawPixmapERK5QRectRK7QPixmap(void* qthis, void* arg0, void* arg1); // 2
  // proto:  void QPainter::drawPixmap(const QPointF & p, const QPixmap & pm, const QRectF & sr);
extern void C_ZN8QPainter10drawPixmapERK7QPointFRK7QPixmapRK6QRectF(void* qthis, void* arg0, void* arg1, void* arg2); // 2
  // proto:  void QPainter::drawPixmap(int x, int y, const QPixmap & pm, int sx, int sy, int sw, int sh);
extern void C_ZN8QPainter10drawPixmapEiiRK7QPixmapiiii(void* qthis, int32_t arg0, int32_t arg1, void* arg2, int32_t arg3, int32_t arg4, int32_t arg5, int32_t arg6); // 2
  // proto:  void QPainter::drawPixmap(const QPoint & p, const QPixmap & pm);
extern void C_ZN8QPainter10drawPixmapERK6QPointRK7QPixmap(void* qthis, void* arg0, void* arg1); // 2
  // proto:  void QPainter::drawPixmap(int x, int y, int w, int h, const QPixmap & pm, int sx, int sy, int sw, int sh);
extern void C_ZN8QPainter10drawPixmapEiiiiRK7QPixmapiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, void* arg4, int32_t arg5, int32_t arg6, int32_t arg7, int32_t arg8); // 2
  // proto:  void QPainter::drawPixmap(int x, int y, const QPixmap & pm);
extern void C_ZN8QPainter10drawPixmapEiiRK7QPixmap(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 2
  // proto:  void QPainter::drawPixmap(const QRect & targetRect, const QPixmap & pixmap, const QRect & sourceRect);
extern void C_ZN8QPainter10drawPixmapERK5QRectRK7QPixmapS2_(void* qthis, void* arg0, void* arg1, void* arg2); // 2
  // proto:  void QPainter::drawPixmap(const QPointF & p, const QPixmap & pm);
extern void C_ZN8QPainter10drawPixmapERK7QPointFRK7QPixmap(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QPainter::drawPixmap(const QRectF & targetRect, const QPixmap & pixmap, const QRectF & sourceRect);
extern void C_ZN8QPainter10drawPixmapERK6QRectFRK7QPixmapS2_(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QPainter::drawPixmap(const QPoint & p, const QPixmap & pm, const QRect & sr);
extern void C_ZN8QPainter10drawPixmapERK6QPointRK7QPixmapRK5QRect(void* qthis, void* arg0, void* arg1, void* arg2); // 2
  // proto:  void QPainter::drawPixmap(int x, int y, int w, int h, const QPixmap & pm);
extern void C_ZN8QPainter10drawPixmapEiiiiRK7QPixmap(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, void* arg4); // 2
  // proto:  void QPainter::drawText(const QRectF & r, const QString & text, const QTextOption & o);
extern void C_ZN8QPainter8drawTextERK6QRectFRK7QStringRK11QTextOption(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QPainter::drawText(const QRectF & r, int flags, const QString & text, QRectF * br);
extern void C_ZN8QPainter8drawTextERK6QRectFiRK7QStringPS0_(void* qthis, void* arg0, int32_t arg1, void* arg2, void* arg3); // 4
  // proto:  void QPainter::drawText(int x, int y, const QString & s);
extern void C_ZN8QPainter8drawTextEiiRK7QString(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 2
  // proto:  void QPainter::drawText(int x, int y, int w, int h, int flags, const QString & text, QRect * br);
extern void C_ZN8QPainter8drawTextEiiiiiRK7QStringP5QRect(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4, void* arg5, void* arg6); // 2
  // proto:  void QPainter::drawText(const QRect & r, int flags, const QString & text, QRect * br);
extern void C_ZN8QPainter8drawTextERK5QRectiRK7QStringPS0_(void* qthis, void* arg0, int32_t arg1, void* arg2, void* arg3); // 4
  // proto:  void QPainter::drawText(const QPoint & p, const QString & s);
extern void C_ZN8QPainter8drawTextERK6QPointRK7QString(void* qthis, void* arg0, void* arg1); // 2
  // proto:  void QPainter::drawText(const QPointF & p, const QString & str, int tf, int justificationPadding);
extern void C_ZN8QPainter8drawTextERK7QPointFRK7QStringii(void* qthis, void* arg0, void* arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QPainter::drawText(const QPointF & p, const QString & s);
extern void C_ZN8QPainter8drawTextERK7QPointFRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QPainter::CompositionMode QPainter::compositionMode();
extern void C_ZNK8QPainter15compositionModeEv(void* qthis); // 4
  // proto:  void QPainter::beginNativePainting();
extern void C_ZN8QPainter19beginNativePaintingEv(void* qthis); // 4
  // proto:  RenderHints QPainter::renderHints();
extern void C_ZNK8QPainter11renderHintsEv(void* qthis); // 4
  // proto:  QRect QPainter::window();
extern void* C_ZNK8QPainter6windowEv(void* qthis); // 4
  // proto:  void QPainter::strokePath(const QPainterPath & path, const QPen & pen);
extern void C_ZN8QPainter10strokePathERK12QPainterPathRK4QPen(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QPainter::setBrushOrigin(int x, int y);
extern void C_ZN8QPainter14setBrushOriginEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  void QPainter::setBrushOrigin(const QPoint & );
extern void C_ZN8QPainter14setBrushOriginERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QPainter::setBrushOrigin(const QPointF & );
extern void C_ZN8QPainter14setBrushOriginERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QPainter::save();
extern void C_ZN8QPainter4saveEv(void* qthis); // 4
  // proto:  void QPainter::setViewport(const QRect & viewport);
extern void C_ZN8QPainter11setViewportERK5QRect(void* qthis, void* arg0); // 4
  // proto:  void QPainter::setViewport(int x, int y, int w, int h);
extern void C_ZN8QPainter11setViewportEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  void QPainter::shear(qreal sh, qreal sv);
extern void C_ZN8QPainter5shearEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  QRectF QPainter::clipBoundingRect();
extern void* C_ZNK8QPainter16clipBoundingRectEv(void* qthis); // 4
  // proto:  QFontMetrics QPainter::fontMetrics();
extern void* C_ZNK8QPainter11fontMetricsEv(void* qthis); // 4
  // proto:  void QPainter::resetTransform();
extern void C_ZN8QPainter14resetTransformEv(void* qthis); // 4
  // proto:  bool QPainter::hasClipping();
extern bool C_ZNK8QPainter11hasClippingEv(void* qthis); // 4
  // proto:  const QBrush & QPainter::background();
extern void* C_ZNK8QPainter10backgroundEv(void* qthis); // 4
  // proto:  const QMatrix & QPainter::deviceMatrix();
extern void* C_ZNK8QPainter12deviceMatrixEv(void* qthis); // 4
  // proto:  void QPainter::rotate(qreal a);
extern void C_ZN8QPainter6rotateEd(void* qthis, double arg0); // 4
  // proto:  Qt::BGMode QPainter::backgroundMode();
extern void C_ZNK8QPainter14backgroundModeEv(void* qthis); // 4
  // proto:  void QPainter::drawConvexPolygon(const QPolygon & polygon);
extern void C_ZN8QPainter17drawConvexPolygonERK8QPolygon(void* qthis, void* arg0); // 2
  // proto:  void QPainter::drawConvexPolygon(const QPoint * points, int pointCount);
extern void C_ZN8QPainter17drawConvexPolygonEPK6QPointi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QPainter::drawConvexPolygon(const QPolygonF & polygon);
extern void C_ZN8QPainter17drawConvexPolygonERK9QPolygonF(void* qthis, void* arg0); // 2
  // proto:  void QPainter::drawConvexPolygon(const QPointF * points, int pointCount);
extern void C_ZN8QPainter17drawConvexPolygonEPK7QPointFi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QPainter::drawArc(const QRect & , int a, int alen);
extern void C_ZN8QPainter7drawArcERK5QRectii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 2
  // proto:  void QPainter::drawArc(int x, int y, int w, int h, int a, int alen);
extern void C_ZN8QPainter7drawArcEiiiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4, int32_t arg5); // 2
  // proto:  void QPainter::drawArc(const QRectF & rect, int a, int alen);
extern void C_ZN8QPainter7drawArcERK6QRectFii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QPainter::setTransform(const QTransform & transform, bool combine);
extern void C_ZN8QPainter12setTransformERK10QTransformb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  void QPainter::drawPoints(const QPointF * points, int pointCount);
extern void C_ZN8QPainter10drawPointsEPK7QPointFi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QPainter::drawPoints(const QPolygonF & points);
extern void C_ZN8QPainter10drawPointsERK9QPolygonF(void* qthis, void* arg0); // 2
  // proto:  void QPainter::drawPoints(const QPoint * points, int pointCount);
extern void C_ZN8QPainter10drawPointsEPK6QPointi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QPainter::drawPoints(const QPolygon & points);
extern void C_ZN8QPainter10drawPointsERK8QPolygon(void* qthis, void* arg0); // 2
  // proto:  void QPainter::setMatrixEnabled(bool enabled);
extern void C_ZN8QPainter16setMatrixEnabledEb(void* qthis, bool arg0); // 4
  // proto:  QPaintEngine * QPainter::paintEngine();
extern void* C_ZNK8QPainter11paintEngineEv(void* qthis); // 4
  // proto:  QRegion QPainter::clipRegion();
extern void* C_ZNK8QPainter10clipRegionEv(void* qthis); // 4
  // proto:  void QPainter::setWindow(int x, int y, int w, int h);
extern void C_ZN8QPainter9setWindowEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  void QPainter::setWindow(const QRect & window);
extern void C_ZN8QPainter9setWindowERK5QRect(void* qthis, void* arg0); // 4
  // proto:  void QPainter::scale(qreal sx, qreal sy);
extern void C_ZN8QPainter5scaleEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  void QPainter::drawImage(const QRect & r, const QImage & image);
extern void C_ZN8QPainter9drawImageERK5QRectRK6QImage(void* qthis, void* arg0, void* arg1); // 2
  // proto:  void QPainter::drawImage(const QPoint & p, const QImage & image);
extern void C_ZN8QPainter9drawImageERK6QPointRK6QImage(void* qthis, void* arg0, void* arg1); // 2
  // proto:  void QPainter::drawImage(const QRectF & r, const QImage & image);
extern void C_ZN8QPainter9drawImageERK6QRectFRK6QImage(void* qthis, void* arg0, void* arg1); // 2
  // proto:  void QPainter::drawImage(const QPointF & p, const QImage & image);
extern void C_ZN8QPainter9drawImageERK7QPointFRK6QImage(void* qthis, void* arg0, void* arg1); // 4
  // proto:  const QMatrix & QPainter::matrix();
extern void* C_ZNK8QPainter6matrixEv(void* qthis); // 4
  // proto:  const QPen & QPainter::pen();
extern void* C_ZNK8QPainter3penEv(void* qthis); // 4
  // proto:  void QPainter::drawPicture(const QPoint & p, const QPicture & picture);
extern void C_ZN8QPainter11drawPictureERK6QPointRK8QPicture(void* qthis, void* arg0, void* arg1); // 2
  // proto:  void QPainter::drawPicture(int x, int y, const QPicture & picture);
extern void C_ZN8QPainter11drawPictureEiiRK8QPicture(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 2
  // proto:  void QPainter::drawPicture(const QPointF & p, const QPicture & picture);
extern void C_ZN8QPainter11drawPictureERK7QPointFRK8QPicture(void* qthis, void* arg0, void* arg1); // 4
  // proto:  const QMatrix & QPainter::worldMatrix();
extern void* C_ZNK8QPainter11worldMatrixEv(void* qthis); // 4
  // proto:  const QTransform & QPainter::deviceTransform();
extern void* C_ZNK8QPainter15deviceTransformEv(void* qthis); // 4
  // proto:  void QPainter::drawPolyline(const QPoint * points, int pointCount);
extern void C_ZN8QPainter12drawPolylineEPK6QPointi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QPainter::drawPolyline(const QPolygon & polygon);
extern void C_ZN8QPainter12drawPolylineERK8QPolygon(void* qthis, void* arg0); // 2
  // proto:  void QPainter::drawPolyline(const QPolygonF & polyline);
extern void C_ZN8QPainter12drawPolylineERK9QPolygonF(void* qthis, void* arg0); // 2
  // proto:  void QPainter::drawPolyline(const QPointF * points, int pointCount);
extern void C_ZN8QPainter12drawPolylineEPK7QPointFi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QPainter::drawRect(const QRect & rect);
extern void C_ZN8QPainter8drawRectERK5QRect(void* qthis, void* arg0); // 2
  // proto:  void QPainter::drawRect(int x1, int y1, int w, int h);
extern void C_ZN8QPainter8drawRectEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  void QPainter::drawRect(const QRectF & rect);
extern void C_ZN8QPainter8drawRectERK6QRectF(void* qthis, void* arg0); // 2
  // proto:  void QPainter::drawLines(const QLineF * lines, int lineCount);
extern void C_ZN8QPainter9drawLinesEPK6QLineFi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QPainter::drawLines(const QPointF * pointPairs, int lineCount);
extern void C_ZN8QPainter9drawLinesEPK7QPointFi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QPainter::drawLines(const QLine * lines, int lineCount);
extern void C_ZN8QPainter9drawLinesEPK5QLinei(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QPainter::drawLines(const QPoint * pointPairs, int lineCount);
extern void C_ZN8QPainter9drawLinesEPK6QPointi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  bool QPainter::isActive();
extern bool C_ZNK8QPainter8isActiveEv(void* qthis); // 4
  // proto:  void QPainter::setWorldMatrix(const QMatrix & matrix, bool combine);
extern void C_ZN8QPainter14setWorldMatrixERK7QMatrixb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  void QPainter::setMatrix(const QMatrix & matrix, bool combine);
extern void C_ZN8QPainter9setMatrixERK7QMatrixb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  void QPainter::endNativePainting();
extern void C_ZN8QPainter17endNativePaintingEv(void* qthis); // 4
  // proto:  const QBrush & QPainter::brush();
extern void* C_ZNK8QPainter5brushEv(void* qthis); // 4
  // proto:  void QPainter::fillPath(const QPainterPath & path, const QBrush & brush);
extern void C_ZN8QPainter8fillPathERK12QPainterPathRK6QBrush(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QPainter::worldMatrixEnabled();
extern bool C_ZNK8QPainter18worldMatrixEnabledEv(void* qthis); // 4
  // proto:  Qt::LayoutDirection QPainter::layoutDirection();
extern void C_ZNK8QPainter15layoutDirectionEv(void* qthis); // 4
  // proto:  const QTransform & QPainter::transform();
extern void* C_ZNK8QPainter9transformEv(void* qthis); // 4
  // proto:  bool QPainter::viewTransformEnabled();
extern bool C_ZNK8QPainter20viewTransformEnabledEv(void* qthis); // 4
  // proto:  void QPainter::setBackground(const QBrush & bg);
extern void C_ZN8QPainter13setBackgroundERK6QBrush(void* qthis, void* arg0); // 4
  // proto:  void QPainter::drawTiledPixmap(const QRect & , const QPixmap & , const QPoint & );
extern void C_ZN8QPainter15drawTiledPixmapERK5QRectRK7QPixmapRK6QPoint(void* qthis, void* arg0, void* arg1, void* arg2); // 2
  // proto:  void QPainter::drawTiledPixmap(int x, int y, int w, int h, const QPixmap & , int sx, int sy);
extern void C_ZN8QPainter15drawTiledPixmapEiiiiRK7QPixmapii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, void* arg4, int32_t arg5, int32_t arg6); // 2
  // proto:  void QPainter::drawTiledPixmap(const QRectF & rect, const QPixmap & pm, const QPointF & offset);
extern void C_ZN8QPainter15drawTiledPixmapERK6QRectFRK7QPixmapRK7QPointF(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QPainter::translate(const QPointF & offset);
extern void C_ZN8QPainter9translateERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QPainter::translate(const QPoint & offset);
extern void C_ZN8QPainter9translateERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QPainter::translate(qreal dx, qreal dy);
extern void C_ZN8QPainter9translateEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QPainter::drawStaticText(const QPointF & topLeftPosition, const QStaticText & staticText);
extern void C_ZN8QPainter14drawStaticTextERK7QPointFRK11QStaticText(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QPainter::drawStaticText(const QPoint & topLeftPosition, const QStaticText & staticText);
extern void C_ZN8QPainter14drawStaticTextERK6QPointRK11QStaticText(void* qthis, void* arg0, void* arg1); // 2
  // proto:  void QPainter::drawStaticText(int left, int top, const QStaticText & staticText);
extern void C_ZN8QPainter14drawStaticTextEiiRK11QStaticText(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 2
  // proto:  void QPainter::setWorldTransform(const QTransform & matrix, bool combine);
extern void C_ZN8QPainter17setWorldTransformERK10QTransformb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  QRect QPainter::viewport();
extern void* C_ZNK8QPainter8viewportEv(void* qthis); // 4
  // proto:  void QPainter::drawEllipse(const QRect & r);
extern void C_ZN8QPainter11drawEllipseERK5QRect(void* qthis, void* arg0); // 4
  // proto:  void QPainter::drawEllipse(const QPoint & center, int rx, int ry);
extern void C_ZN8QPainter11drawEllipseERK6QPointii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 2
  // proto:  void QPainter::drawEllipse(const QRectF & r);
extern void C_ZN8QPainter11drawEllipseERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QPainter::drawEllipse(int x, int y, int w, int h);
extern void C_ZN8QPainter11drawEllipseEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  void QPainter::drawEllipse(const QPointF & center, qreal rx, qreal ry);
extern void C_ZN8QPainter11drawEllipseERK7QPointFdd(void* qthis, void* arg0, double arg1, double arg2); // 2
  // proto:  void QPainter::eraseRect(const QRectF & );
extern void C_ZN8QPainter9eraseRectERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QPainter::eraseRect(int x, int y, int w, int h);
extern void C_ZN8QPainter9eraseRectEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  void QPainter::eraseRect(const QRect & );
extern void C_ZN8QPainter9eraseRectERK5QRect(void* qthis, void* arg0); // 2
  // proto:  void QPainter::drawTextItem(const QPoint & p, const QTextItem & ti);
extern void C_ZN8QPainter12drawTextItemERK6QPointRK9QTextItem(void* qthis, void* arg0, void* arg1); // 2
  // proto:  void QPainter::drawTextItem(int x, int y, const QTextItem & ti);
extern void C_ZN8QPainter12drawTextItemEiiRK9QTextItem(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 2
  // proto:  void QPainter::drawTextItem(const QPointF & p, const QTextItem & ti);
extern void C_ZN8QPainter12drawTextItemERK7QPointFRK9QTextItem(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QPainter::drawPie(int x, int y, int w, int h, int a, int alen);
extern void C_ZN8QPainter7drawPieEiiiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4, int32_t arg5); // 2
  // proto:  void QPainter::drawPie(const QRectF & rect, int a, int alen);
extern void C_ZN8QPainter7drawPieERK6QRectFii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QPainter::drawPie(const QRect & , int a, int alen);
extern void C_ZN8QPainter7drawPieERK5QRectii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 2
  // proto:  QMatrix QPainter::combinedMatrix();
extern void* C_ZNK8QPainter14combinedMatrixEv(void* qthis); // 4
  // proto:  void QPainter::drawGlyphRun(const QPointF & position, const QGlyphRun & glyphRun);
extern void C_ZN8QPainter12drawGlyphRunERK7QPointFRK9QGlyphRun(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QPainter::drawChord(int x, int y, int w, int h, int a, int alen);
extern void C_ZN8QPainter9drawChordEiiiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4, int32_t arg5); // 2
  // proto:  void QPainter::drawChord(const QRectF & rect, int a, int alen);
extern void C_ZN8QPainter9drawChordERK6QRectFii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QPainter::drawChord(const QRect & , int a, int alen);
extern void C_ZN8QPainter9drawChordERK5QRectii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 2
  // proto:  void QPainter::setBrush(const QBrush & brush);
extern void C_ZN8QPainter8setBrushERK6QBrush(void* qthis, void* arg0); // 4
  // proto:  void QPainter::resetMatrix();
extern void C_ZN8QPainter11resetMatrixEv(void* qthis); // 4
  // proto:  void QPainter::initFrom(const QPaintDevice * device);
extern void C_ZN8QPainter8initFromEPK12QPaintDevice(void* qthis, void* arg0); // 4
  // proto:  void QPainter::setWorldMatrixEnabled(bool enabled);
extern void C_ZN8QPainter21setWorldMatrixEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QPainter::drawRoundRect(const QRectF & r, int xround, int yround);
extern void C_ZN8QPainter13drawRoundRectERK6QRectFii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QPainter::drawRoundRect(int x, int y, int w, int h, int , int );
extern void C_ZN8QPainter13drawRoundRectEiiiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4, int32_t arg5); // 2
  // proto:  void QPainter::drawRoundRect(const QRect & r, int xround, int yround);
extern void C_ZN8QPainter13drawRoundRectERK5QRectii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 2
  // proto:  void QPainter::setViewTransformEnabled(bool enable);
extern void C_ZN8QPainter23setViewTransformEnabledEb(void* qthis, bool arg0); // 4
  // proto:  QFontInfo QPainter::fontInfo();
extern void* C_ZNK8QPainter8fontInfoEv(void* qthis); // 4
  // proto:  void QPainter::~QPainter();
extern void C_ZN8QPainterD2Ev(void* qthis); // 4
  // proto:  void QPainter::drawPath(const QPainterPath & path);
extern void C_ZN8QPainter8drawPathERK12QPainterPath(void* qthis, void* arg0); // 4
  // proto:  bool QPainter::end();
extern bool C_ZN8QPainter3endEv(void* qthis); // 4
  // proto:  void QPainter::setOpacity(qreal opacity);
extern void C_ZN8QPainter10setOpacityEd(void* qthis, double arg0); // 4
  // proto:  const QTransform & QPainter::worldTransform();
extern void* C_ZNK8QPainter14worldTransformEv(void* qthis); // 4
  // proto:  void QPainter::fillRect(int x, int y, int w, int h, const QBrush & );
extern void C_ZN8QPainter8fillRectEiiiiRK6QBrush(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, void* arg4); // 2
  // proto:  void QPainter::fillRect(const QRectF & , const QBrush & );
extern void C_ZN8QPainter8fillRectERK6QRectFRK6QBrush(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QPainter::fillRect(const QRectF & , const QColor & color);
extern void C_ZN8QPainter8fillRectERK6QRectFRK6QColor(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QPainter::fillRect(int x, int y, int w, int h, const QColor & color);
extern void C_ZN8QPainter8fillRectEiiiiRK6QColor(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, void* arg4); // 2
  // proto:  void QPainter::fillRect(const QRect & , const QColor & color);
extern void C_ZN8QPainter8fillRectERK5QRectRK6QColor(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QPainter::fillRect(const QRect & , const QBrush & );
extern void C_ZN8QPainter8fillRectERK5QRectRK6QBrush(void* qthis, void* arg0, void* arg1); // 4
  // proto: static void QPainter::restoreRedirected(const QPaintDevice * device);
extern void C_ZN8QPainter17restoreRedirectedEPK12QPaintDevice(void* arg0); // 4
  // proto:  void QPainter::QPainter();
extern void* C_ZN8QPainterC2Ev(); // 3
  // proto:  void QPainter::QPainter(QPaintDevice * );
extern void* C_ZN8QPainterC2EP12QPaintDevice(void* arg0); // 3
  // proto:  qreal QPainter::opacity();
extern double C_ZNK8QPainter7opacityEv(void* qthis); // 4
  // proto:  bool QPainter::begin(QPaintDevice * );
extern bool C_ZN8QPainter5beginEP12QPaintDevice(void* qthis, void* arg0); // 4
  // proto:  QTransform QPainter::combinedTransform();
extern void* C_ZNK8QPainter17combinedTransformEv(void* qthis); // 4
  // proto:  QPaintDevice * QPainter::device();
extern void* C_ZNK8QPainter6deviceEv(void* qthis); // 4
  // proto: static void QPainter::setRedirected(const QPaintDevice * device, QPaintDevice * replacement, const QPoint & offset);
extern void C_ZN8QPainter13setRedirectedEPK12QPaintDevicePS0_RK6QPoint(void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QPainter::setClipping(bool enable);
extern void C_ZN8QPainter11setClippingEb(void* qthis, bool arg0); // 4
  // proto:  bool QPainter::matrixEnabled();
extern bool C_ZNK8QPainter13matrixEnabledEv(void* qthis); // 4
  // proto:  void QPainter::drawPoint(const QPointF & pt);
extern void C_ZN8QPainter9drawPointERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  void QPainter::drawPoint(int x, int y);
extern void C_ZN8QPainter9drawPointEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  void QPainter::drawPoint(const QPoint & p);
extern void C_ZN8QPainter9drawPointERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  QPainterPath QPainter::clipPath();
extern void* C_ZNK8QPainter8clipPathEv(void* qthis); // 4
  // proto:  void QPainter::setFont(const QFont & f);
extern void C_ZN8QPainter7setFontERK5QFont(void* qthis, void* arg0); // 4
  // proto: static QPaintDevice * QPainter::redirected(const QPaintDevice * device, QPoint * offset);
extern void* C_ZN8QPainter10redirectedEPK12QPaintDeviceP6QPoint(void* arg0, void* arg1); // 4
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

// class sizeof(QPainter)=1
type QPainter struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// restore()
func (this *QPainter) Restore(args ...interface{}) () {
  // restore()
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
    // invoke: _ZN8QPainter7restoreEv
    // invoke: void restore()
    C.C_ZN8QPainter7restoreEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "restore", args)
  }

  return
}

// boundingRect(const class QRect &, int, const class QString &)
func (this *QPainter) Boundingrect(args ...interface{}) (ret interface{}) {
  // boundingRect(const class QRect &, int, const class QString &)
  // boundingRect(const class QRectF &, int, const class QString &)
  // boundingRect(int, int, int, int, int, const class QString &)
  // boundingRect(const class QRectF &, const class QString &, const class QTextOption &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[2][3] = qtrt.Int32Ty(false) // "int"
  vtys[2][4] = qtrt.Int32Ty(false) // "int"
  vtys[2][5] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[3][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[3][2] = reflect.TypeOf(QTextOption{}) // "const QTextOption &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter12boundingRectERK5QRectiRK7QString
    // invoke: QRect boundingRect(const class QRect &, int, const class QString &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN8QPainter12boundingRectERK5QRectiRK7QString(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN8QPainter12boundingRectERK6QRectFiRK7QString
    // invoke: QRectF boundingRect(const class QRectF &, int, const class QString &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN8QPainter12boundingRectERK6QRectFiRK7QString(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZN8QPainter12boundingRectEiiiiiRK7QString
    // invoke: QRect boundingRect(int, int, int, int, int, const class QString &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg5)}
    var ret0 = C.C_ZN8QPainter12boundingRectEiiiiiRK7QString(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZN8QPainter12boundingRectERK6QRectFRK7QStringRK11QTextOption
    // invoke: QRectF boundingRect(const class QRectF &, const class QString &, const class QTextOption &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QTextOption).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN8QPainter12boundingRectERK6QRectFRK7QStringRK11QTextOption(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "boundingRect", args)
  }

  return
}

// drawLine(int, int, int, int)
func (this *QPainter) Drawline(args ...interface{}) () {
  // drawLine(int, int, int, int)
  // drawLine(const class QPoint &, const class QPoint &)
  // drawLine(const class QLineF &)
  // drawLine(const class QPointF &, const class QPointF &)
  // drawLine(const class QLine &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[1][1] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QLineF{}) // "const QLineF &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[3][1] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(qtcore.QLine{}) // "const QLine &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8drawLineEiiii
    // invoke: void drawLine(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN8QPainter8drawLineEiiii(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN8QPainter8drawLineERK6QPointS2_
    // invoke: void drawLine(const class QPoint &, const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter8drawLineERK6QPointS2_(this.Qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN8QPainter8drawLineERK6QLineF
    // invoke: void drawLine(const class QLineF &)
    var arg0 = args[0].(*qtcore.QLineF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter8drawLineERK6QLineF(this.Qclsinst, arg0)
  case 3:
    // invoke: _ZN8QPainter8drawLineERK7QPointFS2_
    // invoke: void drawLine(const class QPointF &, const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter8drawLineERK7QPointFS2_(this.Qclsinst, arg0, arg1)
  case 4:
    // invoke: _ZN8QPainter8drawLineERK5QLine
    // invoke: void drawLine(const class QLine &)
    var arg0 = args[0].(*qtcore.QLine).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter8drawLineERK5QLine(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "drawLine", args)
  }

  return
}

// drawRects(const class QRect *, int)
func (this *QPainter) Drawrects(args ...interface{}) () {
  // drawRects(const class QRect *, int)
  // drawRects(const class QRectF *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9drawRectsEPK5QRecti
    // invoke: void drawRects(const class QRect *, int)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter9drawRectsEPK5QRecti(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPainter9drawRectsEPK6QRectFi
    // invoke: void drawRects(const class QRectF *, int)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter9drawRectsEPK6QRectFi(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "drawRects", args)
  }

  return
}

// setPen(const class QColor &)
func (this *QPainter) Setpen(args ...interface{}) () {
  // setPen(const class QColor &)
  // setPen(const class QPen &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPen{}) // "const QPen &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter6setPenERK6QColor
    // invoke: void setPen(const class QColor &)
    var arg0 = args[0].(*QColor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter6setPenERK6QColor(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN8QPainter6setPenERK4QPen
    // invoke: void setPen(const class QPen &)
    var arg0 = args[0].(*QPen).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter6setPenERK4QPen(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "setPen", args)
  }

  return
}

// font()
func (this *QPainter) Font(args ...interface{}) (ret interface{}) {
  // font()
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
    // invoke: _ZNK8QPainter4fontEv
    // invoke: const QFont & font()
    var ret0 = C.C_ZNK8QPainter4fontEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QFont{}) // "const QFont &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "font", args)
  }

  return
}

// brushOrigin()
func (this *QPainter) Brushorigin(args ...interface{}) (ret interface{}) {
  // brushOrigin()
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
    // invoke: _ZNK8QPainter11brushOriginEv
    // invoke: QPoint brushOrigin()
    var ret0 = C.C_ZNK8QPainter11brushOriginEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "brushOrigin", args)
  }

  return
}

// drawPixmap(const class QRect &, const class QPixmap &)
func (this *QPainter) Drawpixmap(args ...interface{}) () {
  // drawPixmap(const class QRect &, const class QPixmap &)
  // drawPixmap(const class QPointF &, const class QPixmap &, const class QRectF &)
  // drawPixmap(int, int, const class QPixmap &, int, int, int, int)
  // drawPixmap(const class QPoint &, const class QPixmap &)
  // drawPixmap(int, int, int, int, const class QPixmap &, int, int, int, int)
  // drawPixmap(int, int, const class QPixmap &)
  // drawPixmap(const class QRect &, const class QPixmap &, const class QRect &)
  // drawPixmap(const class QPointF &, const class QPixmap &)
  // drawPixmap(const class QRectF &, const class QPixmap &, const class QRectF &)
  // drawPixmap(const class QPoint &, const class QPixmap &, const class QRect &)
  // drawPixmap(int, int, int, int, const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[0][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[1][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[1][2] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[2][3] = qtrt.Int32Ty(false) // "int"
  vtys[2][4] = qtrt.Int32Ty(false) // "int"
  vtys[2][5] = qtrt.Int32Ty(false) // "int"
  vtys[2][6] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[3][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"
  vtys[4][2] = qtrt.Int32Ty(false) // "int"
  vtys[4][3] = qtrt.Int32Ty(false) // "int"
  vtys[4][4] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[4][5] = qtrt.Int32Ty(false) // "int"
  vtys[4][6] = qtrt.Int32Ty(false) // "int"
  vtys[4][7] = qtrt.Int32Ty(false) // "int"
  vtys[4][8] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int32Ty(false) // "int"
  vtys[5][1] = qtrt.Int32Ty(false) // "int"
  vtys[5][2] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[6][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[6][2] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[7][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[8][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[8][2] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[9][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[9][2] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = qtrt.Int32Ty(false) // "int"
  vtys[10][1] = qtrt.Int32Ty(false) // "int"
  vtys[10][2] = qtrt.Int32Ty(false) // "int"
  vtys[10][3] = qtrt.Int32Ty(false) // "int"
  vtys[10][4] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter10drawPixmapERK5QRectRK7QPixmap
    // invoke: void drawPixmap(const class QRect &, const class QPixmap &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPixmap).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter10drawPixmapERK5QRectRK7QPixmap(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPainter10drawPixmapERK7QPointFRK7QPixmapRK6QRectF
    // invoke: void drawPixmap(const class QPointF &, const class QPixmap &, const class QRectF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPixmap).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter10drawPixmapERK7QPointFRK7QPixmapRK6QRectF(this.Qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN8QPainter10drawPixmapEiiRK7QPixmapiiii
    // invoke: void drawPixmap(int, int, const class QPixmap &, int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QPixmap).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(qtrt.PrimConv(args[5], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg5)}
    var arg6 = C.int32_t(qtrt.PrimConv(args[6], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg6)}
    C.C_ZN8QPainter10drawPixmapEiiRK7QPixmapiiii(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
  case 3:
    // invoke: _ZN8QPainter10drawPixmapERK6QPointRK7QPixmap
    // invoke: void drawPixmap(const class QPoint &, const class QPixmap &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPixmap).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter10drawPixmapERK6QPointRK7QPixmap(this.Qclsinst, arg0, arg1)
  case 4:
    // invoke: _ZN8QPainter10drawPixmapEiiiiRK7QPixmapiiii
    // invoke: void drawPixmap(int, int, int, int, const class QPixmap &, int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*QPixmap).Qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(qtrt.PrimConv(args[5], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg5)}
    var arg6 = C.int32_t(qtrt.PrimConv(args[6], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg6)}
    var arg7 = C.int32_t(qtrt.PrimConv(args[7], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg7)}
    var arg8 = C.int32_t(qtrt.PrimConv(args[8], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg8)}
    C.C_ZN8QPainter10drawPixmapEiiiiRK7QPixmapiiii(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
  case 5:
    // invoke: _ZN8QPainter10drawPixmapEiiRK7QPixmap
    // invoke: void drawPixmap(int, int, const class QPixmap &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QPixmap).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter10drawPixmapEiiRK7QPixmap(this.Qclsinst, arg0, arg1, arg2)
  case 6:
    // invoke: _ZN8QPainter10drawPixmapERK5QRectRK7QPixmapS2_
    // invoke: void drawPixmap(const class QRect &, const class QPixmap &, const class QRect &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPixmap).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter10drawPixmapERK5QRectRK7QPixmapS2_(this.Qclsinst, arg0, arg1, arg2)
  case 7:
    // invoke: _ZN8QPainter10drawPixmapERK7QPointFRK7QPixmap
    // invoke: void drawPixmap(const class QPointF &, const class QPixmap &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPixmap).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter10drawPixmapERK7QPointFRK7QPixmap(this.Qclsinst, arg0, arg1)
  case 8:
    // invoke: _ZN8QPainter10drawPixmapERK6QRectFRK7QPixmapS2_
    // invoke: void drawPixmap(const class QRectF &, const class QPixmap &, const class QRectF &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPixmap).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter10drawPixmapERK6QRectFRK7QPixmapS2_(this.Qclsinst, arg0, arg1, arg2)
  case 9:
    // invoke: _ZN8QPainter10drawPixmapERK6QPointRK7QPixmapRK5QRect
    // invoke: void drawPixmap(const class QPoint &, const class QPixmap &, const class QRect &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPixmap).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter10drawPixmapERK6QPointRK7QPixmapRK5QRect(this.Qclsinst, arg0, arg1, arg2)
  case 10:
    // invoke: _ZN8QPainter10drawPixmapEiiiiRK7QPixmap
    // invoke: void drawPixmap(int, int, int, int, const class QPixmap &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*QPixmap).Qclsinst
    if false {fmt.Println(arg4)}
    C.C_ZN8QPainter10drawPixmapEiiiiRK7QPixmap(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QPainter", "drawPixmap", args)
  }

  return
}

// drawText(const class QRectF &, const class QString &, const class QTextOption &)
func (this *QPainter) Drawtext(args ...interface{}) () {
  // drawText(const class QRectF &, const class QString &, const class QTextOption &)
  // drawText(const class QRectF &, int, const class QString &, class QRectF *)
  // drawText(int, int, const class QString &)
  // drawText(int, int, int, int, int, const class QString &, class QRect *)
  // drawText(const class QRect &, int, const class QString &, class QRect *)
  // drawText(const class QPoint &, const class QString &)
  // drawText(const class QPointF &, const class QString &, int, int)
  // drawText(const class QPointF &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QTextOption{}) // "const QTextOption &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][3] = reflect.TypeOf(qtcore.QRectF{}) // "QRectF *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[3][3] = qtrt.Int32Ty(false) // "int"
  vtys[3][4] = qtrt.Int32Ty(false) // "int"
  vtys[3][5] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[3][6] = reflect.TypeOf(qtcore.QRect{}) // "QRect *"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"
  vtys[4][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[4][3] = reflect.TypeOf(qtcore.QRect{}) // "QRect *"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[5][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[6][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[6][2] = qtrt.Int32Ty(false) // "int"
  vtys[6][3] = qtrt.Int32Ty(false) // "int"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[7][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8drawTextERK6QRectFRK7QStringRK11QTextOption
    // invoke: void drawText(const class QRectF &, const class QString &, const class QTextOption &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QTextOption).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter8drawTextERK6QRectFRK7QStringRK11QTextOption(this.Qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN8QPainter8drawTextERK6QRectFiRK7QStringPS0_
    // invoke: void drawText(const class QRectF &, int, const class QString &, class QRectF *)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg3)}
    C.C_ZN8QPainter8drawTextERK6QRectFiRK7QStringPS0_(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 2:
    // invoke: _ZN8QPainter8drawTextEiiRK7QString
    // invoke: void drawText(int, int, const class QString &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter8drawTextEiiRK7QString(this.Qclsinst, arg0, arg1, arg2)
  case 3:
    // invoke: _ZN8QPainter8drawTextEiiiiiRK7QStringP5QRect
    // invoke: void drawText(int, int, int, int, int, const class QString &, class QRect *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg5)}
    var arg6 = args[6].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg6)}
    C.C_ZN8QPainter8drawTextEiiiiiRK7QStringP5QRect(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
  case 4:
    // invoke: _ZN8QPainter8drawTextERK5QRectiRK7QStringPS0_
    // invoke: void drawText(const class QRect &, int, const class QString &, class QRect *)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg3)}
    C.C_ZN8QPainter8drawTextERK5QRectiRK7QStringPS0_(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 5:
    // invoke: _ZN8QPainter8drawTextERK6QPointRK7QString
    // invoke: void drawText(const class QPoint &, const class QString &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter8drawTextERK6QPointRK7QString(this.Qclsinst, arg0, arg1)
  case 6:
    // invoke: _ZN8QPainter8drawTextERK7QPointFRK7QStringii
    // invoke: void drawText(const class QPointF &, const class QString &, int, int)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN8QPainter8drawTextERK7QPointFRK7QStringii(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 7:
    // invoke: _ZN8QPainter8drawTextERK7QPointFRK7QString
    // invoke: void drawText(const class QPointF &, const class QString &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter8drawTextERK7QPointFRK7QString(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "drawText", args)
  }

  return
}

// compositionMode()
func (this *QPainter) Compositionmode(args ...interface{}) () {
  // compositionMode()
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
    // invoke: _ZNK8QPainter15compositionModeEv
    // invoke: QPainter::CompositionMode compositionMode()
    C.C_ZNK8QPainter15compositionModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "compositionMode", args)
  }

  return
}

// beginNativePainting()
func (this *QPainter) Beginnativepainting(args ...interface{}) () {
  // beginNativePainting()
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
    // invoke: _ZN8QPainter19beginNativePaintingEv
    // invoke: void beginNativePainting()
    C.C_ZN8QPainter19beginNativePaintingEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "beginNativePainting", args)
  }

  return
}

// renderHints()
func (this *QPainter) Renderhints(args ...interface{}) () {
  // renderHints()
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
    // invoke: _ZNK8QPainter11renderHintsEv
    // invoke: RenderHints renderHints()
    C.C_ZNK8QPainter11renderHintsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "renderHints", args)
  }

  return
}

// window()
func (this *QPainter) Window(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK8QPainter6windowEv
    // invoke: QRect window()
    var ret0 = C.C_ZNK8QPainter6windowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "window", args)
  }

  return
}

// strokePath(const class QPainterPath &, const class QPen &)
func (this *QPainter) Strokepath(args ...interface{}) () {
  // strokePath(const class QPainterPath &, const class QPen &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[0][1] = reflect.TypeOf(QPen{}) // "const QPen &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter10strokePathERK12QPainterPathRK4QPen
    // invoke: void strokePath(const class QPainterPath &, const class QPen &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPen).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter10strokePathERK12QPainterPathRK4QPen(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "strokePath", args)
  }

  return
}

// setBrushOrigin(int, int)
func (this *QPainter) Setbrushorigin(args ...interface{}) () {
  // setBrushOrigin(int, int)
  // setBrushOrigin(const class QPoint &)
  // setBrushOrigin(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter14setBrushOriginEii
    // invoke: void setBrushOrigin(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter14setBrushOriginEii(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPainter14setBrushOriginERK6QPoint
    // invoke: void setBrushOrigin(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter14setBrushOriginERK6QPoint(this.Qclsinst, arg0)
  case 2:
    // invoke: _ZN8QPainter14setBrushOriginERK7QPointF
    // invoke: void setBrushOrigin(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter14setBrushOriginERK7QPointF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "setBrushOrigin", args)
  }

  return
}

// save()
func (this *QPainter) Save(args ...interface{}) () {
  // save()
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
    // invoke: _ZN8QPainter4saveEv
    // invoke: void save()
    C.C_ZN8QPainter4saveEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "save", args)
  }

  return
}

// setViewport(const class QRect &)
func (this *QPainter) Setviewport(args ...interface{}) () {
  // setViewport(const class QRect &)
  // setViewport(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter11setViewportERK5QRect
    // invoke: void setViewport(const class QRect &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter11setViewportERK5QRect(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN8QPainter11setViewportEiiii
    // invoke: void setViewport(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN8QPainter11setViewportEiiii(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QPainter", "setViewport", args)
  }

  return
}

// shear(qreal, qreal)
func (this *QPainter) Shear(args ...interface{}) () {
  // shear(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter5shearEdd
    // invoke: void shear(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter5shearEdd(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "shear", args)
  }

  return
}

// clipBoundingRect()
func (this *QPainter) Clipboundingrect(args ...interface{}) (ret interface{}) {
  // clipBoundingRect()
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
    // invoke: _ZNK8QPainter16clipBoundingRectEv
    // invoke: QRectF clipBoundingRect()
    var ret0 = C.C_ZNK8QPainter16clipBoundingRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "clipBoundingRect", args)
  }

  return
}

// fontMetrics()
func (this *QPainter) Fontmetrics(args ...interface{}) (ret interface{}) {
  // fontMetrics()
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
    // invoke: _ZNK8QPainter11fontMetricsEv
    // invoke: QFontMetrics fontMetrics()
    var ret0 = C.C_ZNK8QPainter11fontMetricsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QFontMetrics{}) // "QFontMetrics"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "fontMetrics", args)
  }

  return
}

// resetTransform()
func (this *QPainter) Resettransform(args ...interface{}) () {
  // resetTransform()
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
    // invoke: _ZN8QPainter14resetTransformEv
    // invoke: void resetTransform()
    C.C_ZN8QPainter14resetTransformEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "resetTransform", args)
  }

  return
}

// hasClipping()
func (this *QPainter) Hasclipping(args ...interface{}) (ret interface{}) {
  // hasClipping()
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
    // invoke: _ZNK8QPainter11hasClippingEv
    // invoke: bool hasClipping()
    var ret0 = C.C_ZNK8QPainter11hasClippingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "hasClipping", args)
  }

  return
}

// background()
func (this *QPainter) Background(args ...interface{}) (ret interface{}) {
  // background()
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
    // invoke: _ZNK8QPainter10backgroundEv
    // invoke: const QBrush & background()
    var ret0 = C.C_ZNK8QPainter10backgroundEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "background", args)
  }

  return
}

// deviceMatrix()
func (this *QPainter) Devicematrix(args ...interface{}) (ret interface{}) {
  // deviceMatrix()
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
    // invoke: _ZNK8QPainter12deviceMatrixEv
    // invoke: const QMatrix & deviceMatrix()
    var ret0 = C.C_ZNK8QPainter12deviceMatrixEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "deviceMatrix", args)
  }

  return
}

// rotate(qreal)
func (this *QPainter) Rotate(args ...interface{}) () {
  // rotate(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter6rotateEd
    // invoke: void rotate(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter6rotateEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "rotate", args)
  }

  return
}

// backgroundMode()
func (this *QPainter) Backgroundmode(args ...interface{}) () {
  // backgroundMode()
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
    // invoke: _ZNK8QPainter14backgroundModeEv
    // invoke: Qt::BGMode backgroundMode()
    C.C_ZNK8QPainter14backgroundModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "backgroundMode", args)
  }

  return
}

// drawConvexPolygon(const class QPolygon &)
func (this *QPainter) Drawconvexpolygon(args ...interface{}) () {
  // drawConvexPolygon(const class QPolygon &)
  // drawConvexPolygon(const class QPoint *, int)
  // drawConvexPolygon(const class QPolygonF &)
  // drawConvexPolygon(const class QPointF *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygon{}) // "const QPolygon &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF *"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter17drawConvexPolygonERK8QPolygon
    // invoke: void drawConvexPolygon(const class QPolygon &)
    var arg0 = args[0].(*QPolygon).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter17drawConvexPolygonERK8QPolygon(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN8QPainter17drawConvexPolygonEPK6QPointi
    // invoke: void drawConvexPolygon(const class QPoint *, int)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter17drawConvexPolygonEPK6QPointi(this.Qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN8QPainter17drawConvexPolygonERK9QPolygonF
    // invoke: void drawConvexPolygon(const class QPolygonF &)
    var arg0 = args[0].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter17drawConvexPolygonERK9QPolygonF(this.Qclsinst, arg0)
  case 3:
    // invoke: _ZN8QPainter17drawConvexPolygonEPK7QPointFi
    // invoke: void drawConvexPolygon(const class QPointF *, int)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter17drawConvexPolygonEPK7QPointFi(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "drawConvexPolygon", args)
  }

  return
}

// drawArc(const class QRect &, int, int)
func (this *QPainter) Drawarc(args ...interface{}) () {
  // drawArc(const class QRect &, int, int)
  // drawArc(int, int, int, int, int, int)
  // drawArc(const class QRectF &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[1][4] = qtrt.Int32Ty(false) // "int"
  vtys[1][5] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter7drawArcERK5QRectii
    // invoke: void drawArc(const class QRect &, int, int)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter7drawArcERK5QRectii(this.Qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN8QPainter7drawArcEiiiiii
    // invoke: void drawArc(int, int, int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(qtrt.PrimConv(args[5], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg5)}
    C.C_ZN8QPainter7drawArcEiiiiii(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  case 2:
    // invoke: _ZN8QPainter7drawArcERK6QRectFii
    // invoke: void drawArc(const class QRectF &, int, int)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter7drawArcERK6QRectFii(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPainter", "drawArc", args)
  }

  return
}

// setTransform(const class QTransform &, _Bool)
func (this *QPainter) Settransform(args ...interface{}) () {
  // setTransform(const class QTransform &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter12setTransformERK10QTransformb
    // invoke: void setTransform(const class QTransform &, _Bool)
    var arg0 = args[0].(*QTransform).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter12setTransformERK10QTransformb(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "setTransform", args)
  }

  return
}

// drawPoints(const class QPointF *, int)
func (this *QPainter) Drawpoints(args ...interface{}) () {
  // drawPoints(const class QPointF *, int)
  // drawPoints(const class QPolygonF &)
  // drawPoints(const class QPoint *, int)
  // drawPoints(const class QPolygon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint *"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPolygon{}) // "const QPolygon &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter10drawPointsEPK7QPointFi
    // invoke: void drawPoints(const class QPointF *, int)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter10drawPointsEPK7QPointFi(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPainter10drawPointsERK9QPolygonF
    // invoke: void drawPoints(const class QPolygonF &)
    var arg0 = args[0].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter10drawPointsERK9QPolygonF(this.Qclsinst, arg0)
  case 2:
    // invoke: _ZN8QPainter10drawPointsEPK6QPointi
    // invoke: void drawPoints(const class QPoint *, int)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter10drawPointsEPK6QPointi(this.Qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZN8QPainter10drawPointsERK8QPolygon
    // invoke: void drawPoints(const class QPolygon &)
    var arg0 = args[0].(*QPolygon).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter10drawPointsERK8QPolygon(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "drawPoints", args)
  }

  return
}

// setMatrixEnabled(_Bool)
func (this *QPainter) Setmatrixenabled(args ...interface{}) () {
  // setMatrixEnabled(_Bool)
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
    // invoke: _ZN8QPainter16setMatrixEnabledEb
    // invoke: void setMatrixEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter16setMatrixEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "setMatrixEnabled", args)
  }

  return
}

// paintEngine()
func (this *QPainter) Paintengine(args ...interface{}) (ret interface{}) {
  // paintEngine()
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
    // invoke: _ZNK8QPainter11paintEngineEv
    // invoke: QPaintEngine * paintEngine()
    var ret0 = C.C_ZNK8QPainter11paintEngineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPaintEngine{}) // "QPaintEngine *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "paintEngine", args)
  }

  return
}

// clipRegion()
func (this *QPainter) Clipregion(args ...interface{}) (ret interface{}) {
  // clipRegion()
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
    // invoke: _ZNK8QPainter10clipRegionEv
    // invoke: QRegion clipRegion()
    var ret0 = C.C_ZNK8QPainter10clipRegionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegion{}) // "QRegion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "clipRegion", args)
  }

  return
}

// setWindow(int, int, int, int)
func (this *QPainter) Setwindow(args ...interface{}) () {
  // setWindow(int, int, int, int)
  // setWindow(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
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
    // invoke: _ZN8QPainter9setWindowEiiii
    // invoke: void setWindow(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN8QPainter9setWindowEiiii(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN8QPainter9setWindowERK5QRect
    // invoke: void setWindow(const class QRect &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter9setWindowERK5QRect(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "setWindow", args)
  }

  return
}

// scale(qreal, qreal)
func (this *QPainter) Scale(args ...interface{}) () {
  // scale(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter5scaleEdd
    // invoke: void scale(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter5scaleEdd(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "scale", args)
  }

  return
}

// drawImage(const class QRect &, const class QImage &)
func (this *QPainter) Drawimage(args ...interface{}) () {
  // drawImage(const class QRect &, const class QImage &)
  // drawImage(const class QPoint &, const class QImage &)
  // drawImage(const class QRectF &, const class QImage &)
  // drawImage(const class QPointF &, const class QImage &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[0][1] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[1][1] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[2][1] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[3][1] = reflect.TypeOf(QImage{}) // "const QImage &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9drawImageERK5QRectRK6QImage
    // invoke: void drawImage(const class QRect &, const class QImage &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QImage).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter9drawImageERK5QRectRK6QImage(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPainter9drawImageERK6QPointRK6QImage
    // invoke: void drawImage(const class QPoint &, const class QImage &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QImage).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter9drawImageERK6QPointRK6QImage(this.Qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN8QPainter9drawImageERK6QRectFRK6QImage
    // invoke: void drawImage(const class QRectF &, const class QImage &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QImage).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter9drawImageERK6QRectFRK6QImage(this.Qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZN8QPainter9drawImageERK7QPointFRK6QImage
    // invoke: void drawImage(const class QPointF &, const class QImage &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QImage).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter9drawImageERK7QPointFRK6QImage(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "drawImage", args)
  }

  return
}

// matrix()
func (this *QPainter) Matrix(args ...interface{}) (ret interface{}) {
  // matrix()
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
    // invoke: _ZNK8QPainter6matrixEv
    // invoke: const QMatrix & matrix()
    var ret0 = C.C_ZNK8QPainter6matrixEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "matrix", args)
  }

  return
}

// pen()
func (this *QPainter) Pen(args ...interface{}) (ret interface{}) {
  // pen()
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
    // invoke: _ZNK8QPainter3penEv
    // invoke: const QPen & pen()
    var ret0 = C.C_ZNK8QPainter3penEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPen{}) // "const QPen &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "pen", args)
  }

  return
}

// drawPicture(const class QPoint &, const class QPicture &)
func (this *QPainter) Drawpicture(args ...interface{}) () {
  // drawPicture(const class QPoint &, const class QPicture &)
  // drawPicture(int, int, const class QPicture &)
  // drawPicture(const class QPointF &, const class QPicture &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[0][1] = reflect.TypeOf(QPicture{}) // "const QPicture &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QPicture{}) // "const QPicture &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[2][1] = reflect.TypeOf(QPicture{}) // "const QPicture &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter11drawPictureERK6QPointRK8QPicture
    // invoke: void drawPicture(const class QPoint &, const class QPicture &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPicture).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter11drawPictureERK6QPointRK8QPicture(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPainter11drawPictureEiiRK8QPicture
    // invoke: void drawPicture(int, int, const class QPicture &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QPicture).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter11drawPictureEiiRK8QPicture(this.Qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN8QPainter11drawPictureERK7QPointFRK8QPicture
    // invoke: void drawPicture(const class QPointF &, const class QPicture &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPicture).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter11drawPictureERK7QPointFRK8QPicture(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "drawPicture", args)
  }

  return
}

// worldMatrix()
func (this *QPainter) Worldmatrix(args ...interface{}) (ret interface{}) {
  // worldMatrix()
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
    // invoke: _ZNK8QPainter11worldMatrixEv
    // invoke: const QMatrix & worldMatrix()
    var ret0 = C.C_ZNK8QPainter11worldMatrixEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "worldMatrix", args)
  }

  return
}

// deviceTransform()
func (this *QPainter) Devicetransform(args ...interface{}) (ret interface{}) {
  // deviceTransform()
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
    // invoke: _ZNK8QPainter15deviceTransformEv
    // invoke: const QTransform & deviceTransform()
    var ret0 = C.C_ZNK8QPainter15deviceTransformEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTransform{}) // "const QTransform &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "deviceTransform", args)
  }

  return
}

// drawPolyline(const class QPoint *, int)
func (this *QPainter) Drawpolyline(args ...interface{}) () {
  // drawPolyline(const class QPoint *, int)
  // drawPolyline(const class QPolygon &)
  // drawPolyline(const class QPolygonF &)
  // drawPolyline(const class QPointF *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPolygon{}) // "const QPolygon &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF *"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter12drawPolylineEPK6QPointi
    // invoke: void drawPolyline(const class QPoint *, int)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter12drawPolylineEPK6QPointi(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPainter12drawPolylineERK8QPolygon
    // invoke: void drawPolyline(const class QPolygon &)
    var arg0 = args[0].(*QPolygon).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter12drawPolylineERK8QPolygon(this.Qclsinst, arg0)
  case 2:
    // invoke: _ZN8QPainter12drawPolylineERK9QPolygonF
    // invoke: void drawPolyline(const class QPolygonF &)
    var arg0 = args[0].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter12drawPolylineERK9QPolygonF(this.Qclsinst, arg0)
  case 3:
    // invoke: _ZN8QPainter12drawPolylineEPK7QPointFi
    // invoke: void drawPolyline(const class QPointF *, int)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter12drawPolylineEPK7QPointFi(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "drawPolyline", args)
  }

  return
}

// drawRect(const class QRect &)
func (this *QPainter) Drawrect(args ...interface{}) () {
  // drawRect(const class QRect &)
  // drawRect(int, int, int, int)
  // drawRect(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8drawRectERK5QRect
    // invoke: void drawRect(const class QRect &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter8drawRectERK5QRect(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN8QPainter8drawRectEiiii
    // invoke: void drawRect(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN8QPainter8drawRectEiiii(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 2:
    // invoke: _ZN8QPainter8drawRectERK6QRectF
    // invoke: void drawRect(const class QRectF &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter8drawRectERK6QRectF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "drawRect", args)
  }

  return
}

// drawLines(const class QLineF *, int)
func (this *QPainter) Drawlines(args ...interface{}) () {
  // drawLines(const class QLineF *, int)
  // drawLines(const class QPointF *, int)
  // drawLines(const class QLine *, int)
  // drawLines(const class QPoint *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QLineF{}) // "const QLineF *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QLine{}) // "const QLine *"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint *"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9drawLinesEPK6QLineFi
    // invoke: void drawLines(const class QLineF *, int)
    var arg0 = args[0].(*qtcore.QLineF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter9drawLinesEPK6QLineFi(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPainter9drawLinesEPK7QPointFi
    // invoke: void drawLines(const class QPointF *, int)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter9drawLinesEPK7QPointFi(this.Qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN8QPainter9drawLinesEPK5QLinei
    // invoke: void drawLines(const class QLine *, int)
    var arg0 = args[0].(*qtcore.QLine).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter9drawLinesEPK5QLinei(this.Qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZN8QPainter9drawLinesEPK6QPointi
    // invoke: void drawLines(const class QPoint *, int)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter9drawLinesEPK6QPointi(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "drawLines", args)
  }

  return
}

// isActive()
func (this *QPainter) Isactive(args ...interface{}) (ret interface{}) {
  // isActive()
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
    // invoke: _ZNK8QPainter8isActiveEv
    // invoke: bool isActive()
    var ret0 = C.C_ZNK8QPainter8isActiveEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "isActive", args)
  }

  return
}

// setWorldMatrix(const class QMatrix &, _Bool)
func (this *QPainter) Setworldmatrix(args ...interface{}) () {
  // setWorldMatrix(const class QMatrix &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter14setWorldMatrixERK7QMatrixb
    // invoke: void setWorldMatrix(const class QMatrix &, _Bool)
    var arg0 = args[0].(*QMatrix).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter14setWorldMatrixERK7QMatrixb(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "setWorldMatrix", args)
  }

  return
}

// setMatrix(const class QMatrix &, _Bool)
func (this *QPainter) Setmatrix(args ...interface{}) () {
  // setMatrix(const class QMatrix &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9setMatrixERK7QMatrixb
    // invoke: void setMatrix(const class QMatrix &, _Bool)
    var arg0 = args[0].(*QMatrix).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter9setMatrixERK7QMatrixb(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "setMatrix", args)
  }

  return
}

// endNativePainting()
func (this *QPainter) Endnativepainting(args ...interface{}) () {
  // endNativePainting()
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
    // invoke: _ZN8QPainter17endNativePaintingEv
    // invoke: void endNativePainting()
    C.C_ZN8QPainter17endNativePaintingEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "endNativePainting", args)
  }

  return
}

// brush()
func (this *QPainter) Brush(args ...interface{}) (ret interface{}) {
  // brush()
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
    // invoke: _ZNK8QPainter5brushEv
    // invoke: const QBrush & brush()
    var ret0 = C.C_ZNK8QPainter5brushEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "brush", args)
  }

  return
}

// fillPath(const class QPainterPath &, const class QBrush &)
func (this *QPainter) Fillpath(args ...interface{}) () {
  // fillPath(const class QPainterPath &, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[0][1] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8fillPathERK12QPainterPathRK6QBrush
    // invoke: void fillPath(const class QPainterPath &, const class QBrush &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QBrush).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter8fillPathERK12QPainterPathRK6QBrush(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "fillPath", args)
  }

  return
}

// worldMatrixEnabled()
func (this *QPainter) Worldmatrixenabled(args ...interface{}) (ret interface{}) {
  // worldMatrixEnabled()
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
    // invoke: _ZNK8QPainter18worldMatrixEnabledEv
    // invoke: bool worldMatrixEnabled()
    var ret0 = C.C_ZNK8QPainter18worldMatrixEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "worldMatrixEnabled", args)
  }

  return
}

// layoutDirection()
func (this *QPainter) Layoutdirection(args ...interface{}) () {
  // layoutDirection()
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
    // invoke: _ZNK8QPainter15layoutDirectionEv
    // invoke: Qt::LayoutDirection layoutDirection()
    C.C_ZNK8QPainter15layoutDirectionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "layoutDirection", args)
  }

  return
}

// transform()
func (this *QPainter) Transform(args ...interface{}) (ret interface{}) {
  // transform()
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
    // invoke: _ZNK8QPainter9transformEv
    // invoke: const QTransform & transform()
    var ret0 = C.C_ZNK8QPainter9transformEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTransform{}) // "const QTransform &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "transform", args)
  }

  return
}

// viewTransformEnabled()
func (this *QPainter) Viewtransformenabled(args ...interface{}) (ret interface{}) {
  // viewTransformEnabled()
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
    // invoke: _ZNK8QPainter20viewTransformEnabledEv
    // invoke: bool viewTransformEnabled()
    var ret0 = C.C_ZNK8QPainter20viewTransformEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "viewTransformEnabled", args)
  }

  return
}

// setBackground(const class QBrush &)
func (this *QPainter) Setbackground(args ...interface{}) () {
  // setBackground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter13setBackgroundERK6QBrush
    // invoke: void setBackground(const class QBrush &)
    var arg0 = args[0].(*QBrush).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter13setBackgroundERK6QBrush(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "setBackground", args)
  }

  return
}

// drawTiledPixmap(const class QRect &, const class QPixmap &, const class QPoint &)
func (this *QPainter) Drawtiledpixmap(args ...interface{}) () {
  // drawTiledPixmap(const class QRect &, const class QPixmap &, const class QPoint &)
  // drawTiledPixmap(int, int, int, int, const class QPixmap &, int, int)
  // drawTiledPixmap(const class QRectF &, const class QPixmap &, const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[0][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[0][2] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[1][4] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[1][5] = qtrt.Int32Ty(false) // "int"
  vtys[1][6] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[2][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[2][2] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter15drawTiledPixmapERK5QRectRK7QPixmapRK6QPoint
    // invoke: void drawTiledPixmap(const class QRect &, const class QPixmap &, const class QPoint &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPixmap).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter15drawTiledPixmapERK5QRectRK7QPixmapRK6QPoint(this.Qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN8QPainter15drawTiledPixmapEiiiiRK7QPixmapii
    // invoke: void drawTiledPixmap(int, int, int, int, const class QPixmap &, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*QPixmap).Qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(qtrt.PrimConv(args[5], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg5)}
    var arg6 = C.int32_t(qtrt.PrimConv(args[6], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg6)}
    C.C_ZN8QPainter15drawTiledPixmapEiiiiRK7QPixmapii(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
  case 2:
    // invoke: _ZN8QPainter15drawTiledPixmapERK6QRectFRK7QPixmapRK7QPointF
    // invoke: void drawTiledPixmap(const class QRectF &, const class QPixmap &, const class QPointF &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPixmap).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter15drawTiledPixmapERK6QRectFRK7QPixmapRK7QPointF(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPainter", "drawTiledPixmap", args)
  }

  return
}

// translate(const class QPointF &)
func (this *QPainter) Translate(args ...interface{}) () {
  // translate(const class QPointF &)
  // translate(const class QPoint &)
  // translate(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9translateERK7QPointF
    // invoke: void translate(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter9translateERK7QPointF(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN8QPainter9translateERK6QPoint
    // invoke: void translate(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter9translateERK6QPoint(this.Qclsinst, arg0)
  case 2:
    // invoke: _ZN8QPainter9translateEdd
    // invoke: void translate(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter9translateEdd(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "translate", args)
  }

  return
}

// drawStaticText(const class QPointF &, const class QStaticText &)
func (this *QPainter) Drawstatictext(args ...interface{}) () {
  // drawStaticText(const class QPointF &, const class QStaticText &)
  // drawStaticText(const class QPoint &, const class QStaticText &)
  // drawStaticText(int, int, const class QStaticText &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QStaticText{}) // "const QStaticText &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[1][1] = reflect.TypeOf(QStaticText{}) // "const QStaticText &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = reflect.TypeOf(QStaticText{}) // "const QStaticText &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter14drawStaticTextERK7QPointFRK11QStaticText
    // invoke: void drawStaticText(const class QPointF &, const class QStaticText &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStaticText).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter14drawStaticTextERK7QPointFRK11QStaticText(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPainter14drawStaticTextERK6QPointRK11QStaticText
    // invoke: void drawStaticText(const class QPoint &, const class QStaticText &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStaticText).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter14drawStaticTextERK6QPointRK11QStaticText(this.Qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN8QPainter14drawStaticTextEiiRK11QStaticText
    // invoke: void drawStaticText(int, int, const class QStaticText &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QStaticText).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter14drawStaticTextEiiRK11QStaticText(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPainter", "drawStaticText", args)
  }

  return
}

// setWorldTransform(const class QTransform &, _Bool)
func (this *QPainter) Setworldtransform(args ...interface{}) () {
  // setWorldTransform(const class QTransform &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter17setWorldTransformERK10QTransformb
    // invoke: void setWorldTransform(const class QTransform &, _Bool)
    var arg0 = args[0].(*QTransform).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter17setWorldTransformERK10QTransformb(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "setWorldTransform", args)
  }

  return
}

// viewport()
func (this *QPainter) Viewport(args ...interface{}) (ret interface{}) {
  // viewport()
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
    // invoke: _ZNK8QPainter8viewportEv
    // invoke: QRect viewport()
    var ret0 = C.C_ZNK8QPainter8viewportEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "viewport", args)
  }

  return
}

// drawEllipse(const class QRect &)
func (this *QPainter) Drawellipse(args ...interface{}) () {
  // drawEllipse(const class QRect &)
  // drawEllipse(const class QPoint &, int, int)
  // drawEllipse(const class QRectF &)
  // drawEllipse(int, int, int, int)
  // drawEllipse(const class QPointF &, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[3][3] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[4][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[4][2] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter11drawEllipseERK5QRect
    // invoke: void drawEllipse(const class QRect &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter11drawEllipseERK5QRect(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN8QPainter11drawEllipseERK6QPointii
    // invoke: void drawEllipse(const class QPoint &, int, int)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter11drawEllipseERK6QPointii(this.Qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN8QPainter11drawEllipseERK6QRectF
    // invoke: void drawEllipse(const class QRectF &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter11drawEllipseERK6QRectF(this.Qclsinst, arg0)
  case 3:
    // invoke: _ZN8QPainter11drawEllipseEiiii
    // invoke: void drawEllipse(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN8QPainter11drawEllipseEiiii(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 4:
    // invoke: _ZN8QPainter11drawEllipseERK7QPointFdd
    // invoke: void drawEllipse(const class QPointF &, qreal, qreal)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter11drawEllipseERK7QPointFdd(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPainter", "drawEllipse", args)
  }

  return
}

// eraseRect(const class QRectF &)
func (this *QPainter) Eraserect(args ...interface{}) () {
  // eraseRect(const class QRectF &)
  // eraseRect(int, int, int, int)
  // eraseRect(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9eraseRectERK6QRectF
    // invoke: void eraseRect(const class QRectF &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter9eraseRectERK6QRectF(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN8QPainter9eraseRectEiiii
    // invoke: void eraseRect(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN8QPainter9eraseRectEiiii(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 2:
    // invoke: _ZN8QPainter9eraseRectERK5QRect
    // invoke: void eraseRect(const class QRect &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter9eraseRectERK5QRect(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "eraseRect", args)
  }

  return
}

// drawTextItem(const class QPoint &, const class QTextItem &)
func (this *QPainter) Drawtextitem(args ...interface{}) () {
  // drawTextItem(const class QPoint &, const class QTextItem &)
  // drawTextItem(int, int, const class QTextItem &)
  // drawTextItem(const class QPointF &, const class QTextItem &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[0][1] = reflect.TypeOf(QTextItem{}) // "const QTextItem &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QTextItem{}) // "const QTextItem &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[2][1] = reflect.TypeOf(QTextItem{}) // "const QTextItem &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter12drawTextItemERK6QPointRK9QTextItem
    // invoke: void drawTextItem(const class QPoint &, const class QTextItem &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QTextItem).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter12drawTextItemERK6QPointRK9QTextItem(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPainter12drawTextItemEiiRK9QTextItem
    // invoke: void drawTextItem(int, int, const class QTextItem &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QTextItem).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter12drawTextItemEiiRK9QTextItem(this.Qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN8QPainter12drawTextItemERK7QPointFRK9QTextItem
    // invoke: void drawTextItem(const class QPointF &, const class QTextItem &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QTextItem).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter12drawTextItemERK7QPointFRK9QTextItem(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "drawTextItem", args)
  }

  return
}

// drawPie(int, int, int, int, int, int)
func (this *QPainter) Drawpie(args ...interface{}) () {
  // drawPie(int, int, int, int, int, int)
  // drawPie(const class QRectF &, int, int)
  // drawPie(const class QRect &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"
  vtys[0][5] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter7drawPieEiiiiii
    // invoke: void drawPie(int, int, int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(qtrt.PrimConv(args[5], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg5)}
    C.C_ZN8QPainter7drawPieEiiiiii(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  case 1:
    // invoke: _ZN8QPainter7drawPieERK6QRectFii
    // invoke: void drawPie(const class QRectF &, int, int)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter7drawPieERK6QRectFii(this.Qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN8QPainter7drawPieERK5QRectii
    // invoke: void drawPie(const class QRect &, int, int)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter7drawPieERK5QRectii(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPainter", "drawPie", args)
  }

  return
}

// combinedMatrix()
func (this *QPainter) Combinedmatrix(args ...interface{}) (ret interface{}) {
  // combinedMatrix()
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
    // invoke: _ZNK8QPainter14combinedMatrixEv
    // invoke: QMatrix combinedMatrix()
    var ret0 = C.C_ZNK8QPainter14combinedMatrixEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMatrix{}) // "QMatrix"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "combinedMatrix", args)
  }

  return
}

// drawGlyphRun(const class QPointF &, const class QGlyphRun &)
func (this *QPainter) Drawglyphrun(args ...interface{}) () {
  // drawGlyphRun(const class QPointF &, const class QGlyphRun &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QGlyphRun{}) // "const QGlyphRun &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter12drawGlyphRunERK7QPointFRK9QGlyphRun
    // invoke: void drawGlyphRun(const class QPointF &, const class QGlyphRun &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QGlyphRun).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter12drawGlyphRunERK7QPointFRK9QGlyphRun(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "drawGlyphRun", args)
  }

  return
}

// drawChord(int, int, int, int, int, int)
func (this *QPainter) Drawchord(args ...interface{}) () {
  // drawChord(int, int, int, int, int, int)
  // drawChord(const class QRectF &, int, int)
  // drawChord(const class QRect &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"
  vtys[0][5] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9drawChordEiiiiii
    // invoke: void drawChord(int, int, int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(qtrt.PrimConv(args[5], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg5)}
    C.C_ZN8QPainter9drawChordEiiiiii(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  case 1:
    // invoke: _ZN8QPainter9drawChordERK6QRectFii
    // invoke: void drawChord(const class QRectF &, int, int)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter9drawChordERK6QRectFii(this.Qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN8QPainter9drawChordERK5QRectii
    // invoke: void drawChord(const class QRect &, int, int)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter9drawChordERK5QRectii(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPainter", "drawChord", args)
  }

  return
}

// setBrush(const class QBrush &)
func (this *QPainter) Setbrush(args ...interface{}) () {
  // setBrush(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8setBrushERK6QBrush
    // invoke: void setBrush(const class QBrush &)
    var arg0 = args[0].(*QBrush).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter8setBrushERK6QBrush(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "setBrush", args)
  }

  return
}

// resetMatrix()
func (this *QPainter) Resetmatrix(args ...interface{}) () {
  // resetMatrix()
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
    // invoke: _ZN8QPainter11resetMatrixEv
    // invoke: void resetMatrix()
    C.C_ZN8QPainter11resetMatrixEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "resetMatrix", args)
  }

  return
}

// initFrom(const class QPaintDevice *)
func (this *QPainter) Initfrom(args ...interface{}) () {
  // initFrom(const class QPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPaintDevice{}) // "const QPaintDevice *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8initFromEPK12QPaintDevice
    // invoke: void initFrom(const class QPaintDevice *)
    var arg0 = args[0].(*QPaintDevice).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter8initFromEPK12QPaintDevice(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "initFrom", args)
  }

  return
}

// setWorldMatrixEnabled(_Bool)
func (this *QPainter) Setworldmatrixenabled(args ...interface{}) () {
  // setWorldMatrixEnabled(_Bool)
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
    // invoke: _ZN8QPainter21setWorldMatrixEnabledEb
    // invoke: void setWorldMatrixEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter21setWorldMatrixEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "setWorldMatrixEnabled", args)
  }

  return
}

// drawRoundRect(const class QRectF &, int, int)
func (this *QPainter) Drawroundrect(args ...interface{}) () {
  // drawRoundRect(const class QRectF &, int, int)
  // drawRoundRect(int, int, int, int, int, int)
  // drawRoundRect(const class QRect &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[1][4] = qtrt.Int32Ty(false) // "int"
  vtys[1][5] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter13drawRoundRectERK6QRectFii
    // invoke: void drawRoundRect(const class QRectF &, int, int)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter13drawRoundRectERK6QRectFii(this.Qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN8QPainter13drawRoundRectEiiiiii
    // invoke: void drawRoundRect(int, int, int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(qtrt.PrimConv(args[5], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg5)}
    C.C_ZN8QPainter13drawRoundRectEiiiiii(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  case 2:
    // invoke: _ZN8QPainter13drawRoundRectERK5QRectii
    // invoke: void drawRoundRect(const class QRect &, int, int)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter13drawRoundRectERK5QRectii(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPainter", "drawRoundRect", args)
  }

  return
}

// setViewTransformEnabled(_Bool)
func (this *QPainter) Setviewtransformenabled(args ...interface{}) () {
  // setViewTransformEnabled(_Bool)
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
    // invoke: _ZN8QPainter23setViewTransformEnabledEb
    // invoke: void setViewTransformEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter23setViewTransformEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "setViewTransformEnabled", args)
  }

  return
}

// fontInfo()
func (this *QPainter) Fontinfo(args ...interface{}) (ret interface{}) {
  // fontInfo()
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
    // invoke: _ZNK8QPainter8fontInfoEv
    // invoke: QFontInfo fontInfo()
    var ret0 = C.C_ZNK8QPainter8fontInfoEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QFontInfo{}) // "QFontInfo"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "fontInfo", args)
  }

  return
}

// ~QPainter()
func (this *QPainter) Freeqpainter(args ...interface{}) () {
  // ~QPainter()
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
    // invoke: _ZN8QPainterD0Ev
    // invoke: void ~QPainter()
    C.C_ZN8QPainterD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "~QPainter", args)
  }

  return
}

// drawPath(const class QPainterPath &)
func (this *QPainter) Drawpath(args ...interface{}) () {
  // drawPath(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8drawPathERK12QPainterPath
    // invoke: void drawPath(const class QPainterPath &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter8drawPathERK12QPainterPath(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "drawPath", args)
  }

  return
}

// end()
func (this *QPainter) End(args ...interface{}) (ret interface{}) {
  // end()
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
    // invoke: _ZN8QPainter3endEv
    // invoke: bool end()
    var ret0 = C.C_ZN8QPainter3endEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "end", args)
  }

  return
}

// setOpacity(qreal)
func (this *QPainter) Setopacity(args ...interface{}) () {
  // setOpacity(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter10setOpacityEd
    // invoke: void setOpacity(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter10setOpacityEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "setOpacity", args)
  }

  return
}

// worldTransform()
func (this *QPainter) Worldtransform(args ...interface{}) (ret interface{}) {
  // worldTransform()
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
    // invoke: _ZNK8QPainter14worldTransformEv
    // invoke: const QTransform & worldTransform()
    var ret0 = C.C_ZNK8QPainter14worldTransformEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTransform{}) // "const QTransform &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "worldTransform", args)
  }

  return
}

// fillRect(int, int, int, int, const class QBrush &)
func (this *QPainter) Fillrect(args ...interface{}) () {
  // fillRect(int, int, int, int, const class QBrush &)
  // fillRect(const class QRectF &, const class QBrush &)
  // fillRect(const class QRectF &, const class QColor &)
  // fillRect(int, int, int, int, const class QColor &)
  // fillRect(const class QRect &, const class QColor &)
  // fillRect(const class QRect &, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[1][1] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[2][1] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[3][3] = qtrt.Int32Ty(false) // "int"
  vtys[3][4] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[4][1] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[5][1] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8fillRectEiiiiRK6QBrush
    // invoke: void fillRect(int, int, int, int, const class QBrush &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*QBrush).Qclsinst
    if false {fmt.Println(arg4)}
    C.C_ZN8QPainter8fillRectEiiiiRK6QBrush(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
  case 1:
    // invoke: _ZN8QPainter8fillRectERK6QRectFRK6QBrush
    // invoke: void fillRect(const class QRectF &, const class QBrush &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QBrush).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter8fillRectERK6QRectFRK6QBrush(this.Qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN8QPainter8fillRectERK6QRectFRK6QColor
    // invoke: void fillRect(const class QRectF &, const class QColor &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QColor).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter8fillRectERK6QRectFRK6QColor(this.Qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZN8QPainter8fillRectEiiiiRK6QColor
    // invoke: void fillRect(int, int, int, int, const class QColor &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*QColor).Qclsinst
    if false {fmt.Println(arg4)}
    C.C_ZN8QPainter8fillRectEiiiiRK6QColor(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
  case 4:
    // invoke: _ZN8QPainter8fillRectERK5QRectRK6QColor
    // invoke: void fillRect(const class QRect &, const class QColor &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QColor).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter8fillRectERK5QRectRK6QColor(this.Qclsinst, arg0, arg1)
  case 5:
    // invoke: _ZN8QPainter8fillRectERK5QRectRK6QBrush
    // invoke: void fillRect(const class QRect &, const class QBrush &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QBrush).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter8fillRectERK5QRectRK6QBrush(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "fillRect", args)
  }

  return
}

// restoreRedirected(const class QPaintDevice *)
func (this *QPainter) Restoreredirected_S(args ...interface{}) () {
  // restoreRedirected(const class QPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPaintDevice{}) // "const QPaintDevice *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter17restoreRedirectedEPK12QPaintDevice
    // invoke: void restoreRedirected(const class QPaintDevice *)
    var arg0 = args[0].(*QPaintDevice).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter17restoreRedirectedEPK12QPaintDevice(arg0)
  default:
    qtrt.ErrorResolve("QPainter", "restoreRedirected", args)
  }

  return
}

// QPainter()
func NewQPainter(args ...interface{}) *QPainter {
  // QPainter()
  // QPainter(class QPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainterC1Ev
    // invoke: void QPainter()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QPainterC2Ev()
    return &QPainter{Qclsinst:qthis}
  case 1:
    // invoke: _ZN8QPainterC1EP12QPaintDevice
    // invoke: void QPainter(class QPaintDevice *)
    var arg0 = args[0].(*QPaintDevice).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QPainterC2EP12QPaintDevice(arg0)
    return &QPainter{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPainter", "QPainter", args)
  }

  return nil // QPainter{Qclsinst:qthis}
}

// opacity()
func (this *QPainter) Opacity(args ...interface{}) (ret interface{}) {
  // opacity()
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
    // invoke: _ZNK8QPainter7opacityEv
    // invoke: qreal opacity()
    var ret0 = C.C_ZNK8QPainter7opacityEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "opacity", args)
  }

  return
}

// begin(class QPaintDevice *)
func (this *QPainter) Begin(args ...interface{}) (ret interface{}) {
  // begin(class QPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter5beginEP12QPaintDevice
    // invoke: bool begin(class QPaintDevice *)
    var arg0 = args[0].(*QPaintDevice).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN8QPainter5beginEP12QPaintDevice(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "begin", args)
  }

  return
}

// combinedTransform()
func (this *QPainter) Combinedtransform(args ...interface{}) (ret interface{}) {
  // combinedTransform()
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
    // invoke: _ZNK8QPainter17combinedTransformEv
    // invoke: QTransform combinedTransform()
    var ret0 = C.C_ZNK8QPainter17combinedTransformEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTransform{}) // "QTransform"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "combinedTransform", args)
  }

  return
}

// device()
func (this *QPainter) Device(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK8QPainter6deviceEv
    // invoke: QPaintDevice * device()
    var ret0 = C.C_ZNK8QPainter6deviceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "device", args)
  }

  return
}

// setRedirected(const class QPaintDevice *, class QPaintDevice *, const class QPoint &)
func (this *QPainter) Setredirected_S(args ...interface{}) () {
  // setRedirected(const class QPaintDevice *, class QPaintDevice *, const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPaintDevice{}) // "const QPaintDevice *"
  vtys[0][1] = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"
  vtys[0][2] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter13setRedirectedEPK12QPaintDevicePS0_RK6QPoint
    // invoke: void setRedirected(const class QPaintDevice *, class QPaintDevice *, const class QPoint &)
    var arg0 = args[0].(*QPaintDevice).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPaintDevice).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN8QPainter13setRedirectedEPK12QPaintDevicePS0_RK6QPoint(arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPainter", "setRedirected", args)
  }

  return
}

// setClipping(_Bool)
func (this *QPainter) Setclipping(args ...interface{}) () {
  // setClipping(_Bool)
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
    // invoke: _ZN8QPainter11setClippingEb
    // invoke: void setClipping(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter11setClippingEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "setClipping", args)
  }

  return
}

// matrixEnabled()
func (this *QPainter) Matrixenabled(args ...interface{}) (ret interface{}) {
  // matrixEnabled()
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
    // invoke: _ZNK8QPainter13matrixEnabledEv
    // invoke: bool matrixEnabled()
    var ret0 = C.C_ZNK8QPainter13matrixEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "matrixEnabled", args)
  }

  return
}

// drawPoint(const class QPointF &)
func (this *QPainter) Drawpoint(args ...interface{}) () {
  // drawPoint(const class QPointF &)
  // drawPoint(int, int)
  // drawPoint(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9drawPointERK7QPointF
    // invoke: void drawPoint(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter9drawPointERK7QPointF(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN8QPainter9drawPointEii
    // invoke: void drawPoint(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPainter9drawPointEii(this.Qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN8QPainter9drawPointERK6QPoint
    // invoke: void drawPoint(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter9drawPointERK6QPoint(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "drawPoint", args)
  }

  return
}

// clipPath()
func (this *QPainter) Clippath(args ...interface{}) (ret interface{}) {
  // clipPath()
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
    // invoke: _ZNK8QPainter8clipPathEv
    // invoke: QPainterPath clipPath()
    var ret0 = C.C_ZNK8QPainter8clipPathEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "clipPath", args)
  }

  return
}

// setFont(const class QFont &)
func (this *QPainter) Setfont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter7setFontERK5QFont
    // invoke: void setFont(const class QFont &)
    var arg0 = args[0].(*QFont).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPainter7setFontERK5QFont(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "setFont", args)
  }

  return
}

// redirected(const class QPaintDevice *, class QPoint *)
func (this *QPainter) Redirected_S(args ...interface{}) (ret interface{}) {
  // redirected(const class QPaintDevice *, class QPoint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPaintDevice{}) // "const QPaintDevice *"
  vtys[0][1] = reflect.TypeOf(qtcore.QPoint{}) // "QPoint *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter10redirectedEPK12QPaintDeviceP6QPoint
    // invoke: QPaintDevice * redirected(const class QPaintDevice *, class QPoint *)
    var arg0 = args[0].(*QPaintDevice).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN8QPainter10redirectedEPK12QPaintDeviceP6QPoint(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainter", "redirected", args)
  }

  return
}

// <= body block end

