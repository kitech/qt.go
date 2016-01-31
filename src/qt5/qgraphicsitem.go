package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtWidgets/qgraphicsitem.h
// dst-file: /src/widgets/qgraphicsitem.go
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
  // proto:  bool QGraphicsTextItem::contains(const QPointF & point);
extern void C_ZNK17QGraphicsTextItem8containsERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QRectF QGraphicsTextItem::boundingRect();
extern void C_ZNK17QGraphicsTextItem12boundingRectEv(void* qthis); // 4
  // proto:  void QGraphicsTextItem::setPlainText(const QString & text);
extern void C_ZN17QGraphicsTextItem12setPlainTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsTextItem::setHtml(const QString & html);
extern void C_ZN17QGraphicsTextItem7setHtmlERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsTextItem::setTabChangesFocus(bool b);
extern void C_ZN17QGraphicsTextItem18setTabChangesFocusEb(void* qthis, bool arg0); // 4
  // proto:  void QGraphicsTextItem::setDocument(QTextDocument * document);
extern void C_ZN17QGraphicsTextItem11setDocumentEP13QTextDocument(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsTextItem::setOpenExternalLinks(bool open);
extern void C_ZN17QGraphicsTextItem20setOpenExternalLinksEb(void* qthis, bool arg0); // 4
  // proto:  QFont QGraphicsTextItem::font();
extern void C_ZNK17QGraphicsTextItem4fontEv(void* qthis); // 4
  // proto:  void QGraphicsTextItem::setDefaultTextColor(const QColor & c);
extern void C_ZN17QGraphicsTextItem19setDefaultTextColorERK6QColor(void* qthis, void* arg0); // 4
  // proto:  bool QGraphicsTextItem::openExternalLinks();
extern void C_ZNK17QGraphicsTextItem17openExternalLinksEv(void* qthis); // 4
  // proto:  void QGraphicsTextItem::setTextCursor(const QTextCursor & cursor);
extern void C_ZN17QGraphicsTextItem13setTextCursorERK11QTextCursor(void* qthis, void* arg0); // 4
  // proto:  QString QGraphicsTextItem::toPlainText();
extern void C_ZNK17QGraphicsTextItem11toPlainTextEv(void* qthis); // 4
  // proto:  bool QGraphicsTextItem::tabChangesFocus();
extern void C_ZNK17QGraphicsTextItem15tabChangesFocusEv(void* qthis); // 4
  // proto:  QTextCursor QGraphicsTextItem::textCursor();
extern void C_ZNK17QGraphicsTextItem10textCursorEv(void* qthis); // 4
  // proto:  void QGraphicsTextItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void C_ZN17QGraphicsTextItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QGraphicsTextItem::QGraphicsTextItem(QGraphicsItem * parent);
extern void* C_ZN17QGraphicsTextItemC2EP13QGraphicsItem(void* arg0); // 3
  // proto:  void QGraphicsTextItem::QGraphicsTextItem(const QString & text, QGraphicsItem * parent);
extern void* C_ZN17QGraphicsTextItemC2ERK7QStringP13QGraphicsItem(void* arg0, void* arg1); // 3
  // proto:  int QGraphicsTextItem::type();
extern void C_ZNK17QGraphicsTextItem4typeEv(void* qthis); // 4
  // proto:  Qt::TextInteractionFlags QGraphicsTextItem::textInteractionFlags();
extern void C_ZNK17QGraphicsTextItem20textInteractionFlagsEv(void* qthis); // 4
  // proto:  QString QGraphicsTextItem::toHtml();
extern void C_ZNK17QGraphicsTextItem6toHtmlEv(void* qthis); // 4
  // proto:  void QGraphicsTextItem::adjustSize();
extern void C_ZN17QGraphicsTextItem10adjustSizeEv(void* qthis); // 4
  // proto:  void QGraphicsTextItem::~QGraphicsTextItem();
extern void C_ZN17QGraphicsTextItemD2Ev(void* qthis); // 4
  // proto:  void QGraphicsTextItem::setTextWidth(qreal width);
extern void C_ZN17QGraphicsTextItem12setTextWidthEd(void* qthis, double arg0); // 4
  // proto:  bool QGraphicsTextItem::isObscuredBy(const QGraphicsItem * item);
extern void C_ZNK17QGraphicsTextItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  QTextDocument * QGraphicsTextItem::document();
extern void C_ZNK17QGraphicsTextItem8documentEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsTextItem::opaqueArea();
extern void C_ZNK17QGraphicsTextItem10opaqueAreaEv(void* qthis); // 4
  // proto:  QColor QGraphicsTextItem::defaultTextColor();
extern void C_ZNK17QGraphicsTextItem16defaultTextColorEv(void* qthis); // 4
  // proto:  const QMetaObject * QGraphicsTextItem::metaObject();
extern void C_ZNK17QGraphicsTextItem10metaObjectEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsTextItem::shape();
extern void C_ZNK17QGraphicsTextItem5shapeEv(void* qthis); // 4
  // proto:  void QGraphicsTextItem::setFont(const QFont & font);
extern void C_ZN17QGraphicsTextItem7setFontERK5QFont(void* qthis, void* arg0); // 4
  // proto:  qreal QGraphicsTextItem::textWidth();
extern void C_ZNK17QGraphicsTextItem9textWidthEv(void* qthis); // 4
  // proto:  QRectF QGraphicsPixmapItem::boundingRect();
extern void C_ZNK19QGraphicsPixmapItem12boundingRectEv(void* qthis); // 4
  // proto:  void QGraphicsPixmapItem::setPixmap(const QPixmap & pixmap);
extern void C_ZN19QGraphicsPixmapItem9setPixmapERK7QPixmap(void* qthis, void* arg0); // 4
  // proto:  QPainterPath QGraphicsPixmapItem::shape();
extern void C_ZNK19QGraphicsPixmapItem5shapeEv(void* qthis); // 4
  // proto:  Qt::TransformationMode QGraphicsPixmapItem::transformationMode();
extern void C_ZNK19QGraphicsPixmapItem18transformationModeEv(void* qthis); // 4
  // proto:  QPixmap QGraphicsPixmapItem::pixmap();
extern void C_ZNK19QGraphicsPixmapItem6pixmapEv(void* qthis); // 4
  // proto:  bool QGraphicsPixmapItem::contains(const QPointF & point);
extern void C_ZNK19QGraphicsPixmapItem8containsERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QGraphicsPixmapItem::ShapeMode QGraphicsPixmapItem::shapeMode();
extern void C_ZNK19QGraphicsPixmapItem9shapeModeEv(void* qthis); // 4
  // proto:  void QGraphicsPixmapItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void C_ZN19QGraphicsPixmapItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QGraphicsPixmapItem::setOffset(const QPointF & offset);
extern void C_ZN19QGraphicsPixmapItem9setOffsetERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsPixmapItem::setOffset(qreal x, qreal y);
extern void C_ZN19QGraphicsPixmapItem9setOffsetEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  int QGraphicsPixmapItem::type();
extern void C_ZNK19QGraphicsPixmapItem4typeEv(void* qthis); // 4
  // proto:  void QGraphicsPixmapItem::~QGraphicsPixmapItem();
extern void C_ZN19QGraphicsPixmapItemD2Ev(void* qthis); // 4
  // proto:  QPointF QGraphicsPixmapItem::offset();
extern void C_ZNK19QGraphicsPixmapItem6offsetEv(void* qthis); // 4
  // proto:  void QGraphicsPixmapItem::QGraphicsPixmapItem(QGraphicsItem * parent);
extern void* C_ZN19QGraphicsPixmapItemC2EP13QGraphicsItem(void* arg0); // 3
  // proto:  void QGraphicsPixmapItem::QGraphicsPixmapItem(const QPixmap & pixmap, QGraphicsItem * parent);
extern void* C_ZN19QGraphicsPixmapItemC2ERK7QPixmapP13QGraphicsItem(void* arg0, void* arg1); // 3
  // proto:  bool QGraphicsPixmapItem::isObscuredBy(const QGraphicsItem * item);
extern void C_ZNK19QGraphicsPixmapItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  QPainterPath QGraphicsPixmapItem::opaqueArea();
extern void C_ZNK19QGraphicsPixmapItem10opaqueAreaEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsRectItem::opaqueArea();
extern void C_ZNK17QGraphicsRectItem10opaqueAreaEv(void* qthis); // 4
  // proto:  QRectF QGraphicsRectItem::boundingRect();
extern void C_ZNK17QGraphicsRectItem12boundingRectEv(void* qthis); // 4
  // proto:  int QGraphicsRectItem::type();
extern void C_ZNK17QGraphicsRectItem4typeEv(void* qthis); // 4
  // proto:  bool QGraphicsRectItem::contains(const QPointF & point);
extern void C_ZNK17QGraphicsRectItem8containsERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsRectItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void C_ZN17QGraphicsRectItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  QPainterPath QGraphicsRectItem::shape();
extern void C_ZNK17QGraphicsRectItem5shapeEv(void* qthis); // 4
  // proto:  void QGraphicsRectItem::QGraphicsRectItem(QGraphicsItem * parent);
extern void* C_ZN17QGraphicsRectItemC2EP13QGraphicsItem(void* arg0); // 3
  // proto:  void QGraphicsRectItem::QGraphicsRectItem(const QRectF & rect, QGraphicsItem * parent);
extern void* C_ZN17QGraphicsRectItemC2ERK6QRectFP13QGraphicsItem(void* arg0, void* arg1); // 3
  // proto:  void QGraphicsRectItem::QGraphicsRectItem(qreal x, qreal y, qreal w, qreal h, QGraphicsItem * parent);
extern void* C_ZN17QGraphicsRectItemC2EddddP13QGraphicsItem(double arg0, double arg1, double arg2, double arg3, void* arg4); // 3
  // proto:  void QGraphicsRectItem::~QGraphicsRectItem();
extern void C_ZN17QGraphicsRectItemD2Ev(void* qthis); // 4
  // proto:  bool QGraphicsRectItem::isObscuredBy(const QGraphicsItem * item);
extern void C_ZNK17QGraphicsRectItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsRectItem::setRect(const QRectF & rect);
extern void C_ZN17QGraphicsRectItem7setRectERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsRectItem::setRect(qreal x, qreal y, qreal w, qreal h);
extern void C_ZN17QGraphicsRectItem7setRectEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  QRectF QGraphicsRectItem::rect();
extern void C_ZNK17QGraphicsRectItem4rectEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsEllipseItem::opaqueArea();
extern void C_ZNK20QGraphicsEllipseItem10opaqueAreaEv(void* qthis); // 4
  // proto:  QRectF QGraphicsEllipseItem::boundingRect();
extern void C_ZNK20QGraphicsEllipseItem12boundingRectEv(void* qthis); // 4
  // proto:  int QGraphicsEllipseItem::spanAngle();
extern void C_ZNK20QGraphicsEllipseItem9spanAngleEv(void* qthis); // 4
  // proto:  int QGraphicsEllipseItem::type();
extern void C_ZNK20QGraphicsEllipseItem4typeEv(void* qthis); // 4
  // proto:  bool QGraphicsEllipseItem::contains(const QPointF & point);
extern void C_ZNK20QGraphicsEllipseItem8containsERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsEllipseItem::setSpanAngle(int angle);
extern void C_ZN20QGraphicsEllipseItem12setSpanAngleEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsEllipseItem::setRect(const QRectF & rect);
extern void C_ZN20QGraphicsEllipseItem7setRectERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsEllipseItem::setRect(qreal x, qreal y, qreal w, qreal h);
extern void C_ZN20QGraphicsEllipseItem7setRectEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  void QGraphicsEllipseItem::setStartAngle(int angle);
extern void C_ZN20QGraphicsEllipseItem13setStartAngleEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsEllipseItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void C_ZN20QGraphicsEllipseItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  QPainterPath QGraphicsEllipseItem::shape();
extern void C_ZNK20QGraphicsEllipseItem5shapeEv(void* qthis); // 4
  // proto:  void QGraphicsEllipseItem::~QGraphicsEllipseItem();
extern void C_ZN20QGraphicsEllipseItemD2Ev(void* qthis); // 4
  // proto:  int QGraphicsEllipseItem::startAngle();
extern void C_ZNK20QGraphicsEllipseItem10startAngleEv(void* qthis); // 4
  // proto:  bool QGraphicsEllipseItem::isObscuredBy(const QGraphicsItem * item);
extern void C_ZNK20QGraphicsEllipseItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsEllipseItem::QGraphicsEllipseItem(QGraphicsItem * parent);
extern void* C_ZN20QGraphicsEllipseItemC2EP13QGraphicsItem(void* arg0); // 3
  // proto:  void QGraphicsEllipseItem::QGraphicsEllipseItem(qreal x, qreal y, qreal w, qreal h, QGraphicsItem * parent);
extern void* C_ZN20QGraphicsEllipseItemC2EddddP13QGraphicsItem(double arg0, double arg1, double arg2, double arg3, void* arg4); // 3
  // proto:  void QGraphicsEllipseItem::QGraphicsEllipseItem(const QRectF & rect, QGraphicsItem * parent);
extern void* C_ZN20QGraphicsEllipseItemC2ERK6QRectFP13QGraphicsItem(void* arg0, void* arg1); // 3
  // proto:  QRectF QGraphicsEllipseItem::rect();
extern void C_ZNK20QGraphicsEllipseItem4rectEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsPolygonItem::opaqueArea();
extern void C_ZNK20QGraphicsPolygonItem10opaqueAreaEv(void* qthis); // 4
  // proto:  QRectF QGraphicsPolygonItem::boundingRect();
extern void C_ZNK20QGraphicsPolygonItem12boundingRectEv(void* qthis); // 4
  // proto:  void QGraphicsPolygonItem::setPolygon(const QPolygonF & polygon);
extern void C_ZN20QGraphicsPolygonItem10setPolygonERK9QPolygonF(void* qthis, void* arg0); // 4
  // proto:  int QGraphicsPolygonItem::type();
extern void C_ZNK20QGraphicsPolygonItem4typeEv(void* qthis); // 4
  // proto:  bool QGraphicsPolygonItem::contains(const QPointF & point);
extern void C_ZNK20QGraphicsPolygonItem8containsERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsPolygonItem::polygon();
extern void C_ZNK20QGraphicsPolygonItem7polygonEv(void* qthis); // 4
  // proto:  void QGraphicsPolygonItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void C_ZN20QGraphicsPolygonItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  QPainterPath QGraphicsPolygonItem::shape();
extern void C_ZNK20QGraphicsPolygonItem5shapeEv(void* qthis); // 4
  // proto:  void QGraphicsPolygonItem::QGraphicsPolygonItem(const QPolygonF & polygon, QGraphicsItem * parent);
extern void* C_ZN20QGraphicsPolygonItemC2ERK9QPolygonFP13QGraphicsItem(void* arg0, void* arg1); // 3
  // proto:  void QGraphicsPolygonItem::QGraphicsPolygonItem(QGraphicsItem * parent);
extern void* C_ZN20QGraphicsPolygonItemC2EP13QGraphicsItem(void* arg0); // 3
  // proto:  void QGraphicsPolygonItem::~QGraphicsPolygonItem();
extern void C_ZN20QGraphicsPolygonItemD2Ev(void* qthis); // 4
  // proto:  bool QGraphicsPolygonItem::isObscuredBy(const QGraphicsItem * item);
extern void C_ZNK20QGraphicsPolygonItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  Qt::FillRule QGraphicsPolygonItem::fillRule();
extern void C_ZNK20QGraphicsPolygonItem8fillRuleEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsPathItem::opaqueArea();
extern void C_ZNK17QGraphicsPathItem10opaqueAreaEv(void* qthis); // 4
  // proto:  void QGraphicsPathItem::~QGraphicsPathItem();
extern void C_ZN17QGraphicsPathItemD2Ev(void* qthis); // 4
  // proto:  QRectF QGraphicsPathItem::boundingRect();
extern void C_ZNK17QGraphicsPathItem12boundingRectEv(void* qthis); // 4
  // proto:  bool QGraphicsPathItem::isObscuredBy(const QGraphicsItem * item);
extern void C_ZNK17QGraphicsPathItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  bool QGraphicsPathItem::contains(const QPointF & point);
extern void C_ZNK17QGraphicsPathItem8containsERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsPathItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void C_ZN17QGraphicsPathItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  QPainterPath QGraphicsPathItem::shape();
extern void C_ZNK17QGraphicsPathItem5shapeEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsPathItem::path();
extern void C_ZNK17QGraphicsPathItem4pathEv(void* qthis); // 4
  // proto:  void QGraphicsPathItem::setPath(const QPainterPath & path);
extern void C_ZN17QGraphicsPathItem7setPathERK12QPainterPath(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsPathItem::QGraphicsPathItem(const QPainterPath & path, QGraphicsItem * parent);
extern void* C_ZN17QGraphicsPathItemC2ERK12QPainterPathP13QGraphicsItem(void* arg0, void* arg1); // 3
  // proto:  void QGraphicsPathItem::QGraphicsPathItem(QGraphicsItem * parent);
extern void* C_ZN17QGraphicsPathItemC2EP13QGraphicsItem(void* arg0); // 3
  // proto:  int QGraphicsPathItem::type();
extern void C_ZNK17QGraphicsPathItem4typeEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsLineItem::opaqueArea();
extern void C_ZNK17QGraphicsLineItem10opaqueAreaEv(void* qthis); // 4
  // proto:  QRectF QGraphicsLineItem::boundingRect();
extern void C_ZNK17QGraphicsLineItem12boundingRectEv(void* qthis); // 4
  // proto:  int QGraphicsLineItem::type();
extern void C_ZNK17QGraphicsLineItem4typeEv(void* qthis); // 4
  // proto:  bool QGraphicsLineItem::contains(const QPointF & point);
extern void C_ZNK17QGraphicsLineItem8containsERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QPainterPath QGraphicsLineItem::shape();
extern void C_ZNK17QGraphicsLineItem5shapeEv(void* qthis); // 4
  // proto:  void QGraphicsLineItem::~QGraphicsLineItem();
extern void C_ZN17QGraphicsLineItemD2Ev(void* qthis); // 4
  // proto:  void QGraphicsLineItem::QGraphicsLineItem(QGraphicsItem * parent);
extern void* C_ZN17QGraphicsLineItemC2EP13QGraphicsItem(void* arg0); // 3
  // proto:  void QGraphicsLineItem::QGraphicsLineItem(const QLineF & line, QGraphicsItem * parent);
extern void* C_ZN17QGraphicsLineItemC2ERK6QLineFP13QGraphicsItem(void* arg0, void* arg1); // 3
  // proto:  void QGraphicsLineItem::QGraphicsLineItem(qreal x1, qreal y1, qreal x2, qreal y2, QGraphicsItem * parent);
extern void* C_ZN17QGraphicsLineItemC2EddddP13QGraphicsItem(double arg0, double arg1, double arg2, double arg3, void* arg4); // 3
  // proto:  void QGraphicsLineItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void C_ZN17QGraphicsLineItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  QPen QGraphicsLineItem::pen();
extern void C_ZNK17QGraphicsLineItem3penEv(void* qthis); // 4
  // proto:  void QGraphicsLineItem::setPen(const QPen & pen);
extern void C_ZN17QGraphicsLineItem6setPenERK4QPen(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsLineItem::setLine(const QLineF & line);
extern void C_ZN17QGraphicsLineItem7setLineERK6QLineF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsLineItem::setLine(qreal x1, qreal y1, qreal x2, qreal y2);
extern void C_ZN17QGraphicsLineItem7setLineEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  QLineF QGraphicsLineItem::line();
extern void C_ZNK17QGraphicsLineItem4lineEv(void* qthis); // 4
  // proto:  bool QGraphicsLineItem::isObscuredBy(const QGraphicsItem * item);
extern void C_ZNK17QGraphicsLineItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  QPainterPath QGraphicsItemGroup::opaqueArea();
extern void C_ZNK18QGraphicsItemGroup10opaqueAreaEv(void* qthis); // 4
  // proto:  void QGraphicsItemGroup::addToGroup(QGraphicsItem * item);
extern void C_ZN18QGraphicsItemGroup10addToGroupEP13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsItemGroup::QGraphicsItemGroup(QGraphicsItem * parent);
extern void* C_ZN18QGraphicsItemGroupC2EP13QGraphicsItem(void* arg0); // 3
  // proto:  QRectF QGraphicsItemGroup::boundingRect();
extern void C_ZNK18QGraphicsItemGroup12boundingRectEv(void* qthis); // 4
  // proto:  void QGraphicsItemGroup::removeFromGroup(QGraphicsItem * item);
extern void C_ZN18QGraphicsItemGroup15removeFromGroupEP13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  bool QGraphicsItemGroup::isObscuredBy(const QGraphicsItem * item);
extern void C_ZNK18QGraphicsItemGroup12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsItemGroup::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void C_ZN18QGraphicsItemGroup5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  int QGraphicsItemGroup::type();
extern void C_ZNK18QGraphicsItemGroup4typeEv(void* qthis); // 4
  // proto:  void QGraphicsItemGroup::~QGraphicsItemGroup();
extern void C_ZN18QGraphicsItemGroupD2Ev(void* qthis); // 4
  // proto:  QPainterPath QAbstractGraphicsShapeItem::opaqueArea();
extern void C_ZNK26QAbstractGraphicsShapeItem10opaqueAreaEv(void* qthis); // 4
  // proto:  bool QAbstractGraphicsShapeItem::isObscuredBy(const QGraphicsItem * item);
extern void C_ZNK26QAbstractGraphicsShapeItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  QPen QAbstractGraphicsShapeItem::pen();
extern void C_ZNK26QAbstractGraphicsShapeItem3penEv(void* qthis); // 4
  // proto:  QBrush QAbstractGraphicsShapeItem::brush();
extern void C_ZNK26QAbstractGraphicsShapeItem5brushEv(void* qthis); // 4
  // proto:  void QAbstractGraphicsShapeItem::setPen(const QPen & pen);
extern void C_ZN26QAbstractGraphicsShapeItem6setPenERK4QPen(void* qthis, void* arg0); // 4
  // proto:  void QAbstractGraphicsShapeItem::setBrush(const QBrush & brush);
extern void C_ZN26QAbstractGraphicsShapeItem8setBrushERK6QBrush(void* qthis, void* arg0); // 4
  // proto:  void QAbstractGraphicsShapeItem::~QAbstractGraphicsShapeItem();
extern void C_ZN26QAbstractGraphicsShapeItemD2Ev(void* qthis); // 4
  // proto:  void QAbstractGraphicsShapeItem::QAbstractGraphicsShapeItem(QGraphicsItem * parent);
extern void* C_ZN26QAbstractGraphicsShapeItemC2EP13QGraphicsItem(void* arg0); // 3
  // proto:  void QGraphicsItem::show();
extern void C_ZN13QGraphicsItem4showEv(void* qthis); // 2
  // proto:  bool QGraphicsItem::hasFocus();
extern void C_ZNK13QGraphicsItem8hasFocusEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setGroup(QGraphicsItemGroup * group);
extern void C_ZN13QGraphicsItem8setGroupEP18QGraphicsItemGroup(void* qthis, void* arg0); // 4
  // proto:  bool QGraphicsItem::handlesChildEvents();
extern void C_ZNK13QGraphicsItem18handlesChildEventsEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsItem::shape();
extern void C_ZNK13QGraphicsItem5shapeEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::isPanel();
extern void C_ZNK13QGraphicsItem7isPanelEv(void* qthis); // 4
  // proto:  QVariant QGraphicsItem::data(int key);
extern void C_ZNK13QGraphicsItem4dataEi(void* qthis, int32_t arg0); // 4
  // proto:  QGraphicsItem * QGraphicsItem::parentItem();
extern void C_ZNK13QGraphicsItem10parentItemEv(void* qthis); // 4
  // proto:  QRectF QGraphicsItem::childrenBoundingRect();
extern void C_ZNK13QGraphicsItem20childrenBoundingRectEv(void* qthis); // 4
  // proto:  QGraphicsItemGroup * QGraphicsItem::group();
extern void C_ZNK13QGraphicsItem5groupEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::isEnabled();
extern void C_ZNK13QGraphicsItem9isEnabledEv(void* qthis); // 4
  // proto:  QGraphicsWidget * QGraphicsItem::window();
extern void C_ZNK13QGraphicsItem6windowEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setBoundingRegionGranularity(qreal granularity);
extern void C_ZN13QGraphicsItem28setBoundingRegionGranularityEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsItem::setRotation(qreal angle);
extern void C_ZN13QGraphicsItem11setRotationEd(void* qthis, double arg0); // 4
  // proto:  QGraphicsItem * QGraphicsItem::commonAncestorItem(const QGraphicsItem * other);
extern void C_ZNK13QGraphicsItem18commonAncestorItemEPKS_(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapToItem(const QGraphicsItem * item, const QPolygonF & polygon);
extern void C_ZNK13QGraphicsItem9mapToItemEPKS_RK9QPolygonF(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QPolygonF QGraphicsItem::mapToItem(const QGraphicsItem * item, const QRectF & rect);
extern void C_ZNK13QGraphicsItem9mapToItemEPKS_RK6QRectF(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QPainterPath QGraphicsItem::mapToItem(const QGraphicsItem * item, const QPainterPath & path);
extern void C_ZNK13QGraphicsItem9mapToItemEPKS_RK12QPainterPath(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QPointF QGraphicsItem::mapToItem(const QGraphicsItem * item, const QPointF & point);
extern void C_ZNK13QGraphicsItem9mapToItemEPKS_RK7QPointF(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QPolygonF QGraphicsItem::mapToItem(const QGraphicsItem * item, qreal x, qreal y, qreal w, qreal h);
extern void C_ZNK13QGraphicsItem9mapToItemEPKS_dddd(void* qthis, void* arg0, double arg1, double arg2, double arg3, double arg4); // 2
  // proto:  QPointF QGraphicsItem::mapToItem(const QGraphicsItem * item, qreal x, qreal y);
extern void C_ZNK13QGraphicsItem9mapToItemEPKS_dd(void* qthis, void* arg0, double arg1, double arg2); // 2
  // proto:  QGraphicsObject * QGraphicsItem::toGraphicsObject();
extern void C_ZN13QGraphicsItem16toGraphicsObjectEv(void* qthis); // 4
  // proto:  Qt::MouseButtons QGraphicsItem::acceptedMouseButtons();
extern void C_ZNK13QGraphicsItem20acceptedMouseButtonsEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setGraphicsEffect(QGraphicsEffect * effect);
extern void C_ZN13QGraphicsItem17setGraphicsEffectEP15QGraphicsEffect(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapToParent(qreal x, qreal y, qreal w, qreal h);
extern void C_ZNK13QGraphicsItem11mapToParentEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  QPointF QGraphicsItem::mapToParent(const QPointF & point);
extern void C_ZNK13QGraphicsItem11mapToParentERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QPointF QGraphicsItem::mapToParent(qreal x, qreal y);
extern void C_ZNK13QGraphicsItem11mapToParentEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  QPainterPath QGraphicsItem::mapToParent(const QPainterPath & path);
extern void C_ZNK13QGraphicsItem11mapToParentERK12QPainterPath(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapToParent(const QRectF & rect);
extern void C_ZNK13QGraphicsItem11mapToParentERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapToParent(const QPolygonF & polygon);
extern void C_ZNK13QGraphicsItem11mapToParentERK9QPolygonF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsItem::setTransformOriginPoint(const QPointF & origin);
extern void C_ZN13QGraphicsItem23setTransformOriginPointERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsItem::setTransformOriginPoint(qreal ax, qreal ay);
extern void C_ZN13QGraphicsItem23setTransformOriginPointEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  QTransform QGraphicsItem::deviceTransform(const QTransform & viewportTransform);
extern void C_ZNK13QGraphicsItem15deviceTransformERK10QTransform(void* qthis, void* arg0); // 4
  // proto:  qreal QGraphicsItem::rotation();
extern void C_ZNK13QGraphicsItem8rotationEv(void* qthis); // 4
  // proto:  qreal QGraphicsItem::effectiveOpacity();
extern void C_ZNK13QGraphicsItem16effectiveOpacityEv(void* qthis); // 4
  // proto:  QGraphicsItem * QGraphicsItem::focusScopeItem();
extern void C_ZNK13QGraphicsItem14focusScopeItemEv(void* qthis); // 4
  // proto:  Qt::InputMethodHints QGraphicsItem::inputMethodHints();
extern void C_ZNK13QGraphicsItem16inputMethodHintsEv(void* qthis); // 4
  // proto:  QGraphicsItem::PanelModality QGraphicsItem::panelModality();
extern void C_ZNK13QGraphicsItem13panelModalityEv(void* qthis); // 4
  // proto:  QGraphicsWidget * QGraphicsItem::topLevelWidget();
extern void C_ZNK13QGraphicsItem14topLevelWidgetEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::isVisibleTo(const QGraphicsItem * parent);
extern void C_ZNK13QGraphicsItem11isVisibleToEPKS_(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsItem::setAcceptDrops(bool on);
extern void C_ZN13QGraphicsItem14setAcceptDropsEb(void* qthis, bool arg0); // 4
  // proto:  void QGraphicsItem::clearFocus();
extern void C_ZN13QGraphicsItem10clearFocusEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::filtersChildEvents();
extern void C_ZNK13QGraphicsItem18filtersChildEventsEv(void* qthis); // 4
  // proto:  qreal QGraphicsItem::x();
extern void C_ZNK13QGraphicsItem1xEv(void* qthis); // 2
  // proto:  void QGraphicsItem::setTransform(const QTransform & matrix, bool combine);
extern void C_ZN13QGraphicsItem12setTransformERK10QTransformb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  void QGraphicsItem::stackBefore(const QGraphicsItem * sibling);
extern void C_ZN13QGraphicsItem11stackBeforeEPKS_(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsItem::setAcceptHoverEvents(bool enabled);
extern void C_ZN13QGraphicsItem20setAcceptHoverEventsEb(void* qthis, bool arg0); // 4
  // proto:  QGraphicsItem * QGraphicsItem::focusProxy();
extern void C_ZNK13QGraphicsItem10focusProxyEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsItem::mapFromParent(const QPainterPath & path);
extern void C_ZNK13QGraphicsItem13mapFromParentERK12QPainterPath(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapFromParent(const QRectF & rect);
extern void C_ZNK13QGraphicsItem13mapFromParentERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapFromParent(qreal x, qreal y, qreal w, qreal h);
extern void C_ZNK13QGraphicsItem13mapFromParentEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  QPointF QGraphicsItem::mapFromParent(qreal x, qreal y);
extern void C_ZNK13QGraphicsItem13mapFromParentEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  QPointF QGraphicsItem::mapFromParent(const QPointF & point);
extern void C_ZNK13QGraphicsItem13mapFromParentERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapFromParent(const QPolygonF & polygon);
extern void C_ZNK13QGraphicsItem13mapFromParentERK9QPolygonF(void* qthis, void* arg0); // 4
  // proto:  QPointF QGraphicsItem::scenePos();
extern void C_ZNK13QGraphicsItem8scenePosEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setPos(qreal x, qreal y);
extern void C_ZN13QGraphicsItem6setPosEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QGraphicsItem::setPos(const QPointF & pos);
extern void C_ZN13QGraphicsItem6setPosERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  qreal QGraphicsItem::scale();
extern void C_ZNK13QGraphicsItem5scaleEv(void* qthis); // 4
  // proto:  void QGraphicsItem::hide();
extern void C_ZN13QGraphicsItem4hideEv(void* qthis); // 2
  // proto:  QMatrix QGraphicsItem::matrix();
extern void C_ZNK13QGraphicsItem6matrixEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::isSelected();
extern void C_ZNK13QGraphicsItem10isSelectedEv(void* qthis); // 4
  // proto:  void QGraphicsItem::moveBy(qreal dx, qreal dy);
extern void C_ZN13QGraphicsItem6moveByEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QGraphicsItem::setFocusProxy(QGraphicsItem * item);
extern void C_ZN13QGraphicsItem13setFocusProxyEPS_(void* qthis, void* arg0); // 4
  // proto:  QTransform QGraphicsItem::itemTransform(const QGraphicsItem * other, bool * ok);
extern void C_ZNK13QGraphicsItem13itemTransformEPKS_Pb(void* qthis, void* arg0, bool* arg1); // 4
  // proto:  QRectF QGraphicsItem::mapRectToParent(const QRectF & rect);
extern void C_ZNK13QGraphicsItem15mapRectToParentERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QRectF QGraphicsItem::mapRectToParent(qreal x, qreal y, qreal w, qreal h);
extern void C_ZNK13QGraphicsItem15mapRectToParentEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  bool QGraphicsItem::isBlockedByModalPanel(QGraphicsItem ** blockingPanel);
extern void C_ZNK13QGraphicsItem21isBlockedByModalPanelEPPS_(void* qthis, void* arg0); // 4
  // proto:  QGraphicsEffect * QGraphicsItem::graphicsEffect();
extern void C_ZNK13QGraphicsItem14graphicsEffectEv(void* qthis); // 4
  // proto:  QList<QGraphicsItem *> QGraphicsItem::childItems();
extern void C_ZNK13QGraphicsItem10childItemsEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::isClipped();
extern void C_ZNK13QGraphicsItem9isClippedEv(void* qthis); // 4
  // proto:  QString QGraphicsItem::toolTip();
extern void C_ZNK13QGraphicsItem7toolTipEv(void* qthis); // 4
  // proto:  void QGraphicsItem::~QGraphicsItem();
extern void C_ZN13QGraphicsItemD2Ev(void* qthis); // 4
  // proto:  bool QGraphicsItem::isActive();
extern void C_ZNK13QGraphicsItem8isActiveEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setFiltersChildEvents(bool enabled);
extern void C_ZN13QGraphicsItem21setFiltersChildEventsEb(void* qthis, bool arg0); // 4
  // proto:  QPainterPath QGraphicsItem::opaqueArea();
extern void C_ZNK13QGraphicsItem10opaqueAreaEv(void* qthis); // 4
  // proto:  void QGraphicsItem::advance(int phase);
extern void C_ZN13QGraphicsItem7advanceEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsItem::setToolTip(const QString & toolTip);
extern void C_ZN13QGraphicsItem10setToolTipERK7QString(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapToScene(const QPolygonF & polygon);
extern void C_ZNK13QGraphicsItem10mapToSceneERK9QPolygonF(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapToScene(const QRectF & rect);
extern void C_ZNK13QGraphicsItem10mapToSceneERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QPointF QGraphicsItem::mapToScene(const QPointF & point);
extern void C_ZNK13QGraphicsItem10mapToSceneERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QPainterPath QGraphicsItem::mapToScene(const QPainterPath & path);
extern void C_ZNK13QGraphicsItem10mapToSceneERK12QPainterPath(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapToScene(qreal x, qreal y, qreal w, qreal h);
extern void C_ZNK13QGraphicsItem10mapToSceneEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  QPointF QGraphicsItem::mapToScene(qreal x, qreal y);
extern void C_ZNK13QGraphicsItem10mapToSceneEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QGraphicsItem::setHandlesChildEvents(bool enabled);
extern void C_ZN13QGraphicsItem21setHandlesChildEventsEb(void* qthis, bool arg0); // 4
  // proto:  QPointF QGraphicsItem::transformOriginPoint();
extern void C_ZNK13QGraphicsItem20transformOriginPointEv(void* qthis); // 4
  // proto:  QPointF QGraphicsItem::pos();
extern void C_ZNK13QGraphicsItem3posEv(void* qthis); // 4
  // proto:  QRectF QGraphicsItem::mapRectToScene(const QRectF & rect);
extern void C_ZNK13QGraphicsItem14mapRectToSceneERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QRectF QGraphicsItem::mapRectToScene(qreal x, qreal y, qreal w, qreal h);
extern void C_ZNK13QGraphicsItem14mapRectToSceneEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  void QGraphicsItem::ungrabMouse();
extern void C_ZN13QGraphicsItem11ungrabMouseEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::isObscured(qreal x, qreal y, qreal w, qreal h);
extern void C_ZNK13QGraphicsItem10isObscuredEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  bool QGraphicsItem::isObscured(const QRectF & rect);
extern void C_ZNK13QGraphicsItem10isObscuredERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsItem::setSelected(bool selected);
extern void C_ZN13QGraphicsItem11setSelectedEb(void* qthis, bool arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapFromScene(const QPolygonF & polygon);
extern void C_ZNK13QGraphicsItem12mapFromSceneERK9QPolygonF(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapFromScene(const QRectF & rect);
extern void C_ZNK13QGraphicsItem12mapFromSceneERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapFromScene(qreal x, qreal y, qreal w, qreal h);
extern void C_ZNK13QGraphicsItem12mapFromSceneEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  QPointF QGraphicsItem::mapFromScene(qreal x, qreal y);
extern void C_ZNK13QGraphicsItem12mapFromSceneEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  QPainterPath QGraphicsItem::mapFromScene(const QPainterPath & path);
extern void C_ZNK13QGraphicsItem12mapFromSceneERK12QPainterPath(void* qthis, void* arg0); // 4
  // proto:  QPointF QGraphicsItem::mapFromScene(const QPointF & point);
extern void C_ZNK13QGraphicsItem12mapFromSceneERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QGraphicsScene * QGraphicsItem::scene();
extern void C_ZNK13QGraphicsItem5sceneEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::isAncestorOf(const QGraphicsItem * child);
extern void C_ZNK13QGraphicsItem12isAncestorOfEPKS_(void* qthis, void* arg0); // 4
  // proto:  QGraphicsItem * QGraphicsItem::topLevelItem();
extern void C_ZNK13QGraphicsItem12topLevelItemEv(void* qthis); // 4
  // proto:  QGraphicsItem * QGraphicsItem::panel();
extern void C_ZNK13QGraphicsItem5panelEv(void* qthis); // 4
  // proto:  void QGraphicsItem::unsetCursor();
extern void C_ZN13QGraphicsItem11unsetCursorEv(void* qthis); // 4
  // proto:  QList<QGraphicsTransform *> QGraphicsItem::transformations();
extern void C_ZNK13QGraphicsItem15transformationsEv(void* qthis); // 4
  // proto:  QRectF QGraphicsItem::mapRectFromScene(const QRectF & rect);
extern void C_ZNK13QGraphicsItem16mapRectFromSceneERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QRectF QGraphicsItem::mapRectFromScene(qreal x, qreal y, qreal w, qreal h);
extern void C_ZNK13QGraphicsItem16mapRectFromSceneEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  bool QGraphicsItem::contains(const QPointF & point);
extern void C_ZNK13QGraphicsItem8containsERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QTransform QGraphicsItem::transform();
extern void C_ZNK13QGraphicsItem9transformEv(void* qthis); // 4
  // proto:  void QGraphicsItem::ungrabKeyboard();
extern void C_ZN13QGraphicsItem14ungrabKeyboardEv(void* qthis); // 4
  // proto:  void QGraphicsItem::installSceneEventFilter(QGraphicsItem * filterItem);
extern void C_ZN13QGraphicsItem23installSceneEventFilterEPS_(void* qthis, void* arg0); // 4
  // proto:  bool QGraphicsItem::acceptHoverEvents();
extern void C_ZNK13QGraphicsItem17acceptHoverEventsEv(void* qthis); // 4
  // proto:  QRectF QGraphicsItem::mapRectToItem(const QGraphicsItem * item, qreal x, qreal y, qreal w, qreal h);
extern void C_ZNK13QGraphicsItem13mapRectToItemEPKS_dddd(void* qthis, void* arg0, double arg1, double arg2, double arg3, double arg4); // 2
  // proto:  QRectF QGraphicsItem::mapRectToItem(const QGraphicsItem * item, const QRectF & rect);
extern void C_ZNK13QGraphicsItem13mapRectToItemEPKS_RK6QRectF(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QPointF QGraphicsItem::mapFromItem(const QGraphicsItem * item, qreal x, qreal y);
extern void C_ZNK13QGraphicsItem11mapFromItemEPKS_dd(void* qthis, void* arg0, double arg1, double arg2); // 2
  // proto:  QPointF QGraphicsItem::mapFromItem(const QGraphicsItem * item, const QPointF & point);
extern void C_ZNK13QGraphicsItem11mapFromItemEPKS_RK7QPointF(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QPolygonF QGraphicsItem::mapFromItem(const QGraphicsItem * item, qreal x, qreal y, qreal w, qreal h);
extern void C_ZNK13QGraphicsItem11mapFromItemEPKS_dddd(void* qthis, void* arg0, double arg1, double arg2, double arg3, double arg4); // 2
  // proto:  QPolygonF QGraphicsItem::mapFromItem(const QGraphicsItem * item, const QRectF & rect);
extern void C_ZNK13QGraphicsItem11mapFromItemEPKS_RK6QRectF(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QPolygonF QGraphicsItem::mapFromItem(const QGraphicsItem * item, const QPolygonF & polygon);
extern void C_ZNK13QGraphicsItem11mapFromItemEPKS_RK9QPolygonF(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QPainterPath QGraphicsItem::mapFromItem(const QGraphicsItem * item, const QPainterPath & path);
extern void C_ZNK13QGraphicsItem11mapFromItemEPKS_RK12QPainterPath(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QGraphicsItem::QGraphicsItem(QGraphicsItem * parent);
extern void* C_ZN13QGraphicsItemC2EPS_(void* arg0); // 3
  // proto:  int QGraphicsItem::type();
extern void C_ZNK13QGraphicsItem4typeEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setEnabled(bool enabled);
extern void C_ZN13QGraphicsItem10setEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QGraphicsItem::grabMouse();
extern void C_ZN13QGraphicsItem9grabMouseEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setActive(bool active);
extern void C_ZN13QGraphicsItem9setActiveEb(void* qthis, bool arg0); // 4
  // proto:  bool QGraphicsItem::isObscuredBy(const QGraphicsItem * item);
extern void C_ZNK13QGraphicsItem12isObscuredByEPKS_(void* qthis, void* arg0); // 4
  // proto:  bool QGraphicsItem::acceptDrops();
extern void C_ZNK13QGraphicsItem11acceptDropsEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setParentItem(QGraphicsItem * parent);
extern void C_ZN13QGraphicsItem13setParentItemEPS_(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsItem::ensureVisible(qreal x, qreal y, qreal w, qreal h, int xmargin, int ymargin);
extern void C_ZN13QGraphicsItem13ensureVisibleEddddii(void* qthis, double arg0, double arg1, double arg2, double arg3, int32_t arg4, int32_t arg5); // 2
  // proto:  void QGraphicsItem::ensureVisible(const QRectF & rect, int xmargin, int ymargin);
extern void C_ZN13QGraphicsItem13ensureVisibleERK6QRectFii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QGraphicsItem::grabKeyboard();
extern void C_ZN13QGraphicsItem12grabKeyboardEv(void* qthis); // 4
  // proto:  QGraphicsObject * QGraphicsItem::parentObject();
extern void C_ZNK13QGraphicsItem12parentObjectEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::isUnderMouse();
extern void C_ZNK13QGraphicsItem12isUnderMouseEv(void* qthis); // 4
  // proto:  QGraphicsWidget * QGraphicsItem::parentWidget();
extern void C_ZNK13QGraphicsItem12parentWidgetEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setZValue(qreal z);
extern void C_ZN13QGraphicsItem9setZValueEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsItem::setMatrix(const QMatrix & matrix, bool combine);
extern void C_ZN13QGraphicsItem9setMatrixERK7QMatrixb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  QCursor QGraphicsItem::cursor();
extern void C_ZNK13QGraphicsItem6cursorEv(void* qthis); // 4
  // proto:  qreal QGraphicsItem::zValue();
extern void C_ZNK13QGraphicsItem6zValueEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setVisible(bool visible);
extern void C_ZN13QGraphicsItem10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  void QGraphicsItem::resetMatrix();
extern void C_ZN13QGraphicsItem11resetMatrixEv(void* qthis); // 4
  // proto:  QTransform QGraphicsItem::sceneTransform();
extern void C_ZNK13QGraphicsItem14sceneTransformEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::isWindow();
extern void C_ZNK13QGraphicsItem8isWindowEv(void* qthis); // 4
  // proto:  QRectF QGraphicsItem::mapRectFromParent(const QRectF & rect);
extern void C_ZNK13QGraphicsItem17mapRectFromParentERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QRectF QGraphicsItem::mapRectFromParent(qreal x, qreal y, qreal w, qreal h);
extern void C_ZNK13QGraphicsItem17mapRectFromParentEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  void QGraphicsItem::setScale(qreal scale);
extern void C_ZN13QGraphicsItem8setScaleEd(void* qthis, double arg0); // 4
  // proto:  bool QGraphicsItem::isWidget();
extern void C_ZNK13QGraphicsItem8isWidgetEv(void* qthis); // 4
  // proto:  void QGraphicsItem::resetTransform();
extern void C_ZN13QGraphicsItem14resetTransformEv(void* qthis); // 4
  // proto:  QRectF QGraphicsItem::sceneBoundingRect();
extern void C_ZNK13QGraphicsItem17sceneBoundingRectEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setOpacity(qreal opacity);
extern void C_ZN13QGraphicsItem10setOpacityEd(void* qthis, double arg0); // 4
  // proto:  QGraphicsItem * QGraphicsItem::focusItem();
extern void C_ZNK13QGraphicsItem9focusItemEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::hasCursor();
extern void C_ZNK13QGraphicsItem9hasCursorEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setAcceptTouchEvents(bool enabled);
extern void C_ZN13QGraphicsItem20setAcceptTouchEventsEb(void* qthis, bool arg0); // 4
  // proto:  void QGraphicsItem::setData(int key, const QVariant & value);
extern void C_ZN13QGraphicsItem7setDataEiRK8QVariant(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  qreal QGraphicsItem::opacity();
extern void C_ZNK13QGraphicsItem7opacityEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::isVisible();
extern void C_ZNK13QGraphicsItem9isVisibleEv(void* qthis); // 4
  // proto:  void QGraphicsItem::update(qreal x, qreal y, qreal width, qreal height);
extern void C_ZN13QGraphicsItem6updateEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  void QGraphicsItem::update(const QRectF & rect);
extern void C_ZN13QGraphicsItem6updateERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QRectF QGraphicsItem::mapRectFromItem(const QGraphicsItem * item, const QRectF & rect);
extern void C_ZNK13QGraphicsItem15mapRectFromItemEPKS_RK6QRectF(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QRectF QGraphicsItem::mapRectFromItem(const QGraphicsItem * item, qreal x, qreal y, qreal w, qreal h);
extern void C_ZNK13QGraphicsItem15mapRectFromItemEPKS_dddd(void* qthis, void* arg0, double arg1, double arg2, double arg3, double arg4); // 2
  // proto:  void QGraphicsItem::setX(qreal x);
extern void C_ZN13QGraphicsItem4setXEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsItem::setY(qreal y);
extern void C_ZN13QGraphicsItem4setYEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsItem::setCursor(const QCursor & cursor);
extern void C_ZN13QGraphicsItem9setCursorERK7QCursor(void* qthis, void* arg0); // 4
  // proto:  qreal QGraphicsItem::boundingRegionGranularity();
extern void C_ZNK13QGraphicsItem25boundingRegionGranularityEv(void* qthis); // 4
  // proto:  void QGraphicsItem::removeSceneEventFilter(QGraphicsItem * filterItem);
extern void C_ZN13QGraphicsItem22removeSceneEventFilterEPS_(void* qthis, void* arg0); // 4
  // proto:  bool QGraphicsItem::acceptTouchEvents();
extern void C_ZNK13QGraphicsItem17acceptTouchEventsEv(void* qthis); // 4
  // proto:  QMatrix QGraphicsItem::sceneMatrix();
extern void C_ZNK13QGraphicsItem11sceneMatrixEv(void* qthis); // 4
  // proto:  QGraphicsItem::CacheMode QGraphicsItem::cacheMode();
extern void C_ZNK13QGraphicsItem9cacheModeEv(void* qthis); // 4
  // proto:  GraphicsItemFlags QGraphicsItem::flags();
extern void C_ZNK13QGraphicsItem5flagsEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsItem::clipPath();
extern void C_ZNK13QGraphicsItem8clipPathEv(void* qthis); // 4
  // proto:  qreal QGraphicsItem::y();
extern void C_ZNK13QGraphicsItem1yEv(void* qthis); // 2
  // proto:  void QGraphicsItem::scroll(qreal dx, qreal dy, const QRectF & rect);
extern void C_ZN13QGraphicsItem6scrollEddRK6QRectF(void* qthis, double arg0, double arg1, void* arg2); // 4
  // proto:  QRegion QGraphicsItem::boundingRegion(const QTransform & itemToDeviceTransform);
extern void C_ZNK13QGraphicsItem14boundingRegionERK10QTransform(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsObject::~QGraphicsObject();
extern void C_ZN15QGraphicsObjectD2Ev(void* qthis); // 4
  // proto:  void QGraphicsObject::QGraphicsObject(QGraphicsItem * parent);
extern void* C_ZN15QGraphicsObjectC2EP13QGraphicsItem(void* arg0); // 3
  // proto:  const QMetaObject * QGraphicsObject::metaObject();
extern void C_ZNK15QGraphicsObject10metaObjectEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsSimpleTextItem::opaqueArea();
extern void C_ZNK23QGraphicsSimpleTextItem10opaqueAreaEv(void* qthis); // 4
  // proto:  void QGraphicsSimpleTextItem::~QGraphicsSimpleTextItem();
extern void C_ZN23QGraphicsSimpleTextItemD2Ev(void* qthis); // 4
  // proto:  QRectF QGraphicsSimpleTextItem::boundingRect();
extern void C_ZNK23QGraphicsSimpleTextItem12boundingRectEv(void* qthis); // 4
  // proto:  QString QGraphicsSimpleTextItem::text();
extern void C_ZNK23QGraphicsSimpleTextItem4textEv(void* qthis); // 4
  // proto:  void QGraphicsSimpleTextItem::setText(const QString & text);
extern void C_ZN23QGraphicsSimpleTextItem7setTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QGraphicsSimpleTextItem::contains(const QPointF & point);
extern void C_ZNK23QGraphicsSimpleTextItem8containsERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  bool QGraphicsSimpleTextItem::isObscuredBy(const QGraphicsItem * item);
extern void C_ZNK23QGraphicsSimpleTextItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsSimpleTextItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void C_ZN23QGraphicsSimpleTextItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  QPainterPath QGraphicsSimpleTextItem::shape();
extern void C_ZNK23QGraphicsSimpleTextItem5shapeEv(void* qthis); // 4
  // proto:  void QGraphicsSimpleTextItem::QGraphicsSimpleTextItem(QGraphicsItem * parent);
extern void* C_ZN23QGraphicsSimpleTextItemC2EP13QGraphicsItem(void* arg0); // 3
  // proto:  void QGraphicsSimpleTextItem::QGraphicsSimpleTextItem(const QString & text, QGraphicsItem * parent);
extern void* C_ZN23QGraphicsSimpleTextItemC2ERK7QStringP13QGraphicsItem(void* arg0, void* arg1); // 3
  // proto:  void QGraphicsSimpleTextItem::setFont(const QFont & font);
extern void C_ZN23QGraphicsSimpleTextItem7setFontERK5QFont(void* qthis, void* arg0); // 4
  // proto:  QFont QGraphicsSimpleTextItem::font();
extern void C_ZNK23QGraphicsSimpleTextItem4fontEv(void* qthis); // 4
  // proto:  int QGraphicsSimpleTextItem::type();
extern void C_ZNK23QGraphicsSimpleTextItem4typeEv(void* qthis); // 4
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

// class sizeof(QGraphicsTextItem)=1
type QGraphicsTextItem struct {
  /*qbase*/ QGraphicsObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _linkActivated QGraphicsTextItem_linkActivated_signal;
//  _linkHovered QGraphicsTextItem_linkHovered_signal;
}

// class sizeof(QGraphicsPixmapItem)=1
type QGraphicsPixmapItem struct {
  /*qbase*/ QGraphicsItem;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsRectItem)=1
type QGraphicsRectItem struct {
  /*qbase*/ QAbstractGraphicsShapeItem;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsEllipseItem)=1
type QGraphicsEllipseItem struct {
  /*qbase*/ QAbstractGraphicsShapeItem;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsPolygonItem)=1
type QGraphicsPolygonItem struct {
  /*qbase*/ QAbstractGraphicsShapeItem;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsPathItem)=1
type QGraphicsPathItem struct {
  /*qbase*/ QAbstractGraphicsShapeItem;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsLineItem)=1
type QGraphicsLineItem struct {
  /*qbase*/ QGraphicsItem;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsItemGroup)=1
type QGraphicsItemGroup struct {
  /*qbase*/ QGraphicsItem;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAbstractGraphicsShapeItem)=1
type QAbstractGraphicsShapeItem struct {
  /*qbase*/ QGraphicsItem;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsItem)=1
type QGraphicsItem struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsObject)=1
type QGraphicsObject struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _childrenChanged QGraphicsObject_childrenChanged_signal;
//  _parentChanged QGraphicsObject_parentChanged_signal;
//  _heightChanged QGraphicsObject_heightChanged_signal;
//  _zChanged QGraphicsObject_zChanged_signal;
//  _visibleChanged QGraphicsObject_visibleChanged_signal;
//  _yChanged QGraphicsObject_yChanged_signal;
//  _widthChanged QGraphicsObject_widthChanged_signal;
//  _opacityChanged QGraphicsObject_opacityChanged_signal;
//  _rotationChanged QGraphicsObject_rotationChanged_signal;
//  _enabledChanged QGraphicsObject_enabledChanged_signal;
//  _xChanged QGraphicsObject_xChanged_signal;
//  _scaleChanged QGraphicsObject_scaleChanged_signal;
}

// class sizeof(QGraphicsSimpleTextItem)=1
type QGraphicsSimpleTextItem struct {
  /*qbase*/ QAbstractGraphicsShapeItem;
  qclsinst unsafe.Pointer /* *C.void */;
}

// contains(const class QPointF &)
func (this *QGraphicsTextItem) contains(args ...interface{}) () {
  // contains(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem8containsERK7QPointF
    // invoke: bool contains(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK17QGraphicsTextItem8containsERK7QPointF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "contains", args)
  }

}

// boundingRect()
func (this *QGraphicsTextItem) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem12boundingRectEv
    // invoke: QRectF boundingRect()
    var ret = C.C_ZNK17QGraphicsTextItem12boundingRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "boundingRect", args)
  }

}

// setPlainText(const class QString &)
func (this *QGraphicsTextItem) setPlainText(args ...interface{}) () {
  // setPlainText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem12setPlainTextERK7QString
    // invoke: void setPlainText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsTextItem12setPlainTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setPlainText", args)
  }

}

// setHtml(const class QString &)
func (this *QGraphicsTextItem) setHtml(args ...interface{}) () {
  // setHtml(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem7setHtmlERK7QString
    // invoke: void setHtml(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsTextItem7setHtmlERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setHtml", args)
  }

}

// setTabChangesFocus(_Bool)
func (this *QGraphicsTextItem) setTabChangesFocus(args ...interface{}) () {
  // setTabChangesFocus(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem18setTabChangesFocusEb
    // invoke: void setTabChangesFocus(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsTextItem18setTabChangesFocusEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setTabChangesFocus", args)
  }

}

// setDocument(class QTextDocument *)
func (this *QGraphicsTextItem) setDocument(args ...interface{}) () {
  // setDocument(class QTextDocument *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem11setDocumentEP13QTextDocument
    // invoke: void setDocument(class QTextDocument *)
    var arg0 = args[0].(QTextDocument).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsTextItem11setDocumentEP13QTextDocument(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setDocument", args)
  }

}

// setOpenExternalLinks(_Bool)
func (this *QGraphicsTextItem) setOpenExternalLinks(args ...interface{}) () {
  // setOpenExternalLinks(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem20setOpenExternalLinksEb
    // invoke: void setOpenExternalLinks(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsTextItem20setOpenExternalLinksEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setOpenExternalLinks", args)
  }

}

// font()
func (this *QGraphicsTextItem) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem4fontEv
    // invoke: QFont font()
    var ret = C.C_ZNK17QGraphicsTextItem4fontEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "font", args)
  }

}

// setDefaultTextColor(const class QColor &)
func (this *QGraphicsTextItem) setDefaultTextColor(args ...interface{}) () {
  // setDefaultTextColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem19setDefaultTextColorERK6QColor
    // invoke: void setDefaultTextColor(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsTextItem19setDefaultTextColorERK6QColor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setDefaultTextColor", args)
  }

}

// openExternalLinks()
func (this *QGraphicsTextItem) openExternalLinks(args ...interface{}) () {
  // openExternalLinks()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem17openExternalLinksEv
    // invoke: bool openExternalLinks()
    var ret = C.C_ZNK17QGraphicsTextItem17openExternalLinksEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "openExternalLinks", args)
  }

}

// setTextCursor(const class QTextCursor &)
func (this *QGraphicsTextItem) setTextCursor(args ...interface{}) () {
  // setTextCursor(const class QTextCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCursor{}) // "const QTextCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem13setTextCursorERK11QTextCursor
    // invoke: void setTextCursor(const class QTextCursor &)
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsTextItem13setTextCursorERK11QTextCursor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setTextCursor", args)
  }

}

// toPlainText()
func (this *QGraphicsTextItem) toPlainText(args ...interface{}) () {
  // toPlainText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem11toPlainTextEv
    // invoke: QString toPlainText()
    var ret = C.C_ZNK17QGraphicsTextItem11toPlainTextEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "toPlainText", args)
  }

}

// tabChangesFocus()
func (this *QGraphicsTextItem) tabChangesFocus(args ...interface{}) () {
  // tabChangesFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem15tabChangesFocusEv
    // invoke: bool tabChangesFocus()
    var ret = C.C_ZNK17QGraphicsTextItem15tabChangesFocusEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "tabChangesFocus", args)
  }

}

// textCursor()
func (this *QGraphicsTextItem) textCursor(args ...interface{}) () {
  // textCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem10textCursorEv
    // invoke: QTextCursor textCursor()
    var ret = C.C_ZNK17QGraphicsTextItem10textCursorEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "textCursor", args)
  }

}

// paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsTextItem) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
    // invoke: void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(QPainter).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStyleOptionGraphicsItem).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN17QGraphicsTextItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "paint", args)
  }

}

// QGraphicsTextItem(class QGraphicsItem *)
func NewQGraphicsTextItem(args ...interface{}) *QGraphicsTextItem {
  // QGraphicsTextItem(class QGraphicsItem *)
  // QGraphicsTextItem(const class QString &, class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItemC1EP13QGraphicsItem
    // invoke: void QGraphicsTextItem(class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QGraphicsTextItemC2EP13QGraphicsItem(arg0)
    return &QGraphicsTextItem{qclsinst:qthis}
  case 1:
    // invoke: _ZN17QGraphicsTextItemC1ERK7QStringP13QGraphicsItem
    // invoke: void QGraphicsTextItem(const class QString &, class QGraphicsItem *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QGraphicsTextItemC2ERK7QStringP13QGraphicsItem(arg0, arg1)
    return &QGraphicsTextItem{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "QGraphicsTextItem", args)
  }

  return nil // QGraphicsTextItem{qclsinst:qthis}
}

// type()
func (this *QGraphicsTextItem) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem4typeEv
    // invoke: int type()
    var ret = C.C_ZNK17QGraphicsTextItem4typeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "type", args)
  }

}

// textInteractionFlags()
func (this *QGraphicsTextItem) textInteractionFlags(args ...interface{}) () {
  // textInteractionFlags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem20textInteractionFlagsEv
    // invoke: Qt::TextInteractionFlags textInteractionFlags()
    C.C_ZNK17QGraphicsTextItem20textInteractionFlagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "textInteractionFlags", args)
  }

}

// toHtml()
func (this *QGraphicsTextItem) toHtml(args ...interface{}) () {
  // toHtml()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem6toHtmlEv
    // invoke: QString toHtml()
    var ret = C.C_ZNK17QGraphicsTextItem6toHtmlEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "toHtml", args)
  }

}

// adjustSize()
func (this *QGraphicsTextItem) adjustSize(args ...interface{}) () {
  // adjustSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem10adjustSizeEv
    // invoke: void adjustSize()
    C.C_ZN17QGraphicsTextItem10adjustSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "adjustSize", args)
  }

}

// ~QGraphicsTextItem()
func (this *QGraphicsTextItem) FreeQGraphicsTextItem(args ...interface{}) () {
  // ~QGraphicsTextItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItemD0Ev
    // invoke: void ~QGraphicsTextItem()
    C.C_ZN17QGraphicsTextItemD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "~QGraphicsTextItem", args)
  }

}

// setTextWidth(qreal)
func (this *QGraphicsTextItem) setTextWidth(args ...interface{}) () {
  // setTextWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem12setTextWidthEd
    // invoke: void setTextWidth(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsTextItem12setTextWidthEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setTextWidth", args)
  }

}

// isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsTextItem) isObscuredBy(args ...interface{}) () {
  // isObscuredBy(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem12isObscuredByEPK13QGraphicsItem
    // invoke: bool isObscuredBy(const class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK17QGraphicsTextItem12isObscuredByEPK13QGraphicsItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "isObscuredBy", args)
  }

}

// document()
func (this *QGraphicsTextItem) document(args ...interface{}) () {
  // document()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem8documentEv
    // invoke: QTextDocument * document()
    var ret = C.C_ZNK17QGraphicsTextItem8documentEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "document", args)
  }

}

// opaqueArea()
func (this *QGraphicsTextItem) opaqueArea(args ...interface{}) () {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem10opaqueAreaEv
    // invoke: QPainterPath opaqueArea()
    var ret = C.C_ZNK17QGraphicsTextItem10opaqueAreaEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "opaqueArea", args)
  }

}

// defaultTextColor()
func (this *QGraphicsTextItem) defaultTextColor(args ...interface{}) () {
  // defaultTextColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem16defaultTextColorEv
    // invoke: QColor defaultTextColor()
    var ret = C.C_ZNK17QGraphicsTextItem16defaultTextColorEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "defaultTextColor", args)
  }

}

// metaObject()
func (this *QGraphicsTextItem) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK17QGraphicsTextItem10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "metaObject", args)
  }

}

// shape()
func (this *QGraphicsTextItem) shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem5shapeEv
    // invoke: QPainterPath shape()
    var ret = C.C_ZNK17QGraphicsTextItem5shapeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "shape", args)
  }

}

// setFont(const class QFont &)
func (this *QGraphicsTextItem) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem7setFontERK5QFont
    // invoke: void setFont(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsTextItem7setFontERK5QFont(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setFont", args)
  }

}

// textWidth()
func (this *QGraphicsTextItem) textWidth(args ...interface{}) () {
  // textWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem9textWidthEv
    // invoke: qreal textWidth()
    var ret = C.C_ZNK17QGraphicsTextItem9textWidthEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "textWidth", args)
  }

}

// boundingRect()
func (this *QGraphicsPixmapItem) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem12boundingRectEv
    // invoke: QRectF boundingRect()
    var ret = C.C_ZNK19QGraphicsPixmapItem12boundingRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "boundingRect", args)
  }

}

// setPixmap(const class QPixmap &)
func (this *QGraphicsPixmapItem) setPixmap(args ...interface{}) () {
  // setPixmap(const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsPixmapItem9setPixmapERK7QPixmap
    // invoke: void setPixmap(const class QPixmap &)
    var arg0 = args[0].(QPixmap).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsPixmapItem9setPixmapERK7QPixmap(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "setPixmap", args)
  }

}

// shape()
func (this *QGraphicsPixmapItem) shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem5shapeEv
    // invoke: QPainterPath shape()
    var ret = C.C_ZNK19QGraphicsPixmapItem5shapeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "shape", args)
  }

}

// transformationMode()
func (this *QGraphicsPixmapItem) transformationMode(args ...interface{}) () {
  // transformationMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem18transformationModeEv
    // invoke: Qt::TransformationMode transformationMode()
    C.C_ZNK19QGraphicsPixmapItem18transformationModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "transformationMode", args)
  }

}

// pixmap()
func (this *QGraphicsPixmapItem) pixmap(args ...interface{}) () {
  // pixmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem6pixmapEv
    // invoke: QPixmap pixmap()
    var ret = C.C_ZNK19QGraphicsPixmapItem6pixmapEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "pixmap", args)
  }

}

// contains(const class QPointF &)
func (this *QGraphicsPixmapItem) contains(args ...interface{}) () {
  // contains(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem8containsERK7QPointF
    // invoke: bool contains(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK19QGraphicsPixmapItem8containsERK7QPointF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "contains", args)
  }

}

// shapeMode()
func (this *QGraphicsPixmapItem) shapeMode(args ...interface{}) () {
  // shapeMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem9shapeModeEv
    // invoke: QGraphicsPixmapItem::ShapeMode shapeMode()
    C.C_ZNK19QGraphicsPixmapItem9shapeModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "shapeMode", args)
  }

}

// paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsPixmapItem) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsPixmapItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
    // invoke: void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(QPainter).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStyleOptionGraphicsItem).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN19QGraphicsPixmapItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "paint", args)
  }

}

// setOffset(const class QPointF &)
func (this *QGraphicsPixmapItem) setOffset(args ...interface{}) () {
  // setOffset(const class QPointF &)
  // setOffset(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsPixmapItem9setOffsetERK7QPointF
    // invoke: void setOffset(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsPixmapItem9setOffsetERK7QPointF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN19QGraphicsPixmapItem9setOffsetEdd
    // invoke: void setOffset(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsPixmapItem9setOffsetEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "setOffset", args)
  }

}

// type()
func (this *QGraphicsPixmapItem) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem4typeEv
    // invoke: int type()
    var ret = C.C_ZNK19QGraphicsPixmapItem4typeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "type", args)
  }

}

// ~QGraphicsPixmapItem()
func (this *QGraphicsPixmapItem) FreeQGraphicsPixmapItem(args ...interface{}) () {
  // ~QGraphicsPixmapItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsPixmapItemD0Ev
    // invoke: void ~QGraphicsPixmapItem()
    C.C_ZN19QGraphicsPixmapItemD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "~QGraphicsPixmapItem", args)
  }

}

// offset()
func (this *QGraphicsPixmapItem) offset(args ...interface{}) () {
  // offset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem6offsetEv
    // invoke: QPointF offset()
    var ret = C.C_ZNK19QGraphicsPixmapItem6offsetEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "offset", args)
  }

}

// QGraphicsPixmapItem(class QGraphicsItem *)
func NewQGraphicsPixmapItem(args ...interface{}) *QGraphicsPixmapItem {
  // QGraphicsPixmapItem(class QGraphicsItem *)
  // QGraphicsPixmapItem(const class QPixmap &, class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[1][1] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsPixmapItemC1EP13QGraphicsItem
    // invoke: void QGraphicsPixmapItem(class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QGraphicsPixmapItemC2EP13QGraphicsItem(arg0)
    return &QGraphicsPixmapItem{qclsinst:qthis}
  case 1:
    // invoke: _ZN19QGraphicsPixmapItemC1ERK7QPixmapP13QGraphicsItem
    // invoke: void QGraphicsPixmapItem(const class QPixmap &, class QGraphicsItem *)
    var arg0 = args[0].(QPixmap).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QGraphicsPixmapItemC2ERK7QPixmapP13QGraphicsItem(arg0, arg1)
    return &QGraphicsPixmapItem{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "QGraphicsPixmapItem", args)
  }

  return nil // QGraphicsPixmapItem{qclsinst:qthis}
}

// isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsPixmapItem) isObscuredBy(args ...interface{}) () {
  // isObscuredBy(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem12isObscuredByEPK13QGraphicsItem
    // invoke: bool isObscuredBy(const class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK19QGraphicsPixmapItem12isObscuredByEPK13QGraphicsItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "isObscuredBy", args)
  }

}

// opaqueArea()
func (this *QGraphicsPixmapItem) opaqueArea(args ...interface{}) () {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem10opaqueAreaEv
    // invoke: QPainterPath opaqueArea()
    var ret = C.C_ZNK19QGraphicsPixmapItem10opaqueAreaEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "opaqueArea", args)
  }

}

// opaqueArea()
func (this *QGraphicsRectItem) opaqueArea(args ...interface{}) () {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRectItem10opaqueAreaEv
    // invoke: QPainterPath opaqueArea()
    var ret = C.C_ZNK17QGraphicsRectItem10opaqueAreaEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "opaqueArea", args)
  }

}

// boundingRect()
func (this *QGraphicsRectItem) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRectItem12boundingRectEv
    // invoke: QRectF boundingRect()
    var ret = C.C_ZNK17QGraphicsRectItem12boundingRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "boundingRect", args)
  }

}

// type()
func (this *QGraphicsRectItem) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRectItem4typeEv
    // invoke: int type()
    var ret = C.C_ZNK17QGraphicsRectItem4typeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "type", args)
  }

}

// contains(const class QPointF &)
func (this *QGraphicsRectItem) contains(args ...interface{}) () {
  // contains(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRectItem8containsERK7QPointF
    // invoke: bool contains(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK17QGraphicsRectItem8containsERK7QPointF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "contains", args)
  }

}

// paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsRectItem) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsRectItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
    // invoke: void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(QPainter).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStyleOptionGraphicsItem).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN17QGraphicsRectItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "paint", args)
  }

}

// shape()
func (this *QGraphicsRectItem) shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRectItem5shapeEv
    // invoke: QPainterPath shape()
    var ret = C.C_ZNK17QGraphicsRectItem5shapeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "shape", args)
  }

}

// QGraphicsRectItem(class QGraphicsItem *)
func NewQGraphicsRectItem(args ...interface{}) *QGraphicsRectItem {
  // QGraphicsRectItem(class QGraphicsItem *)
  // QGraphicsRectItem(const class QRectF &, class QGraphicsItem *)
  // QGraphicsRectItem(qreal, qreal, qreal, qreal, class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][4] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsRectItemC1EP13QGraphicsItem
    // invoke: void QGraphicsRectItem(class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QGraphicsRectItemC2EP13QGraphicsItem(arg0)
    return &QGraphicsRectItem{qclsinst:qthis}
  case 1:
    // invoke: _ZN17QGraphicsRectItemC1ERK6QRectFP13QGraphicsItem
    // invoke: void QGraphicsRectItem(const class QRectF &, class QGraphicsItem *)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QGraphicsRectItemC2ERK6QRectFP13QGraphicsItem(arg0, arg1)
    return &QGraphicsRectItem{qclsinst:qthis}
  case 2:
    // invoke: _ZN17QGraphicsRectItemC1EddddP13QGraphicsItem
    // invoke: void QGraphicsRectItem(qreal, qreal, qreal, qreal, class QGraphicsItem *)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg4)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QGraphicsRectItemC2EddddP13QGraphicsItem(arg0, arg1, arg2, arg3, arg4)
    return &QGraphicsRectItem{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "QGraphicsRectItem", args)
  }

  return nil // QGraphicsRectItem{qclsinst:qthis}
}

// ~QGraphicsRectItem()
func (this *QGraphicsRectItem) FreeQGraphicsRectItem(args ...interface{}) () {
  // ~QGraphicsRectItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsRectItemD0Ev
    // invoke: void ~QGraphicsRectItem()
    C.C_ZN17QGraphicsRectItemD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "~QGraphicsRectItem", args)
  }

}

// isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsRectItem) isObscuredBy(args ...interface{}) () {
  // isObscuredBy(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRectItem12isObscuredByEPK13QGraphicsItem
    // invoke: bool isObscuredBy(const class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK17QGraphicsRectItem12isObscuredByEPK13QGraphicsItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "isObscuredBy", args)
  }

}

// setRect(const class QRectF &)
func (this *QGraphicsRectItem) setRect(args ...interface{}) () {
  // setRect(const class QRectF &)
  // setRect(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsRectItem7setRectERK6QRectF
    // invoke: void setRect(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsRectItem7setRectERK6QRectF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN17QGraphicsRectItem7setRectEdddd
    // invoke: void setRect(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN17QGraphicsRectItem7setRectEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "setRect", args)
  }

}

// rect()
func (this *QGraphicsRectItem) rect(args ...interface{}) () {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRectItem4rectEv
    // invoke: QRectF rect()
    var ret = C.C_ZNK17QGraphicsRectItem4rectEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "rect", args)
  }

}

// opaqueArea()
func (this *QGraphicsEllipseItem) opaqueArea(args ...interface{}) () {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem10opaqueAreaEv
    // invoke: QPainterPath opaqueArea()
    var ret = C.C_ZNK20QGraphicsEllipseItem10opaqueAreaEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "opaqueArea", args)
  }

}

// boundingRect()
func (this *QGraphicsEllipseItem) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem12boundingRectEv
    // invoke: QRectF boundingRect()
    var ret = C.C_ZNK20QGraphicsEllipseItem12boundingRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "boundingRect", args)
  }

}

// spanAngle()
func (this *QGraphicsEllipseItem) spanAngle(args ...interface{}) () {
  // spanAngle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem9spanAngleEv
    // invoke: int spanAngle()
    var ret = C.C_ZNK20QGraphicsEllipseItem9spanAngleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "spanAngle", args)
  }

}

// type()
func (this *QGraphicsEllipseItem) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem4typeEv
    // invoke: int type()
    var ret = C.C_ZNK20QGraphicsEllipseItem4typeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "type", args)
  }

}

// contains(const class QPointF &)
func (this *QGraphicsEllipseItem) contains(args ...interface{}) () {
  // contains(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem8containsERK7QPointF
    // invoke: bool contains(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK20QGraphicsEllipseItem8containsERK7QPointF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "contains", args)
  }

}

// setSpanAngle(int)
func (this *QGraphicsEllipseItem) setSpanAngle(args ...interface{}) () {
  // setSpanAngle(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsEllipseItem12setSpanAngleEi
    // invoke: void setSpanAngle(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN20QGraphicsEllipseItem12setSpanAngleEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "setSpanAngle", args)
  }

}

// setRect(const class QRectF &)
func (this *QGraphicsEllipseItem) setRect(args ...interface{}) () {
  // setRect(const class QRectF &)
  // setRect(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsEllipseItem7setRectERK6QRectF
    // invoke: void setRect(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN20QGraphicsEllipseItem7setRectERK6QRectF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN20QGraphicsEllipseItem7setRectEdddd
    // invoke: void setRect(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN20QGraphicsEllipseItem7setRectEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "setRect", args)
  }

}

// setStartAngle(int)
func (this *QGraphicsEllipseItem) setStartAngle(args ...interface{}) () {
  // setStartAngle(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsEllipseItem13setStartAngleEi
    // invoke: void setStartAngle(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN20QGraphicsEllipseItem13setStartAngleEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "setStartAngle", args)
  }

}

// paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsEllipseItem) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsEllipseItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
    // invoke: void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(QPainter).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStyleOptionGraphicsItem).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN20QGraphicsEllipseItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "paint", args)
  }

}

// shape()
func (this *QGraphicsEllipseItem) shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem5shapeEv
    // invoke: QPainterPath shape()
    var ret = C.C_ZNK20QGraphicsEllipseItem5shapeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "shape", args)
  }

}

// ~QGraphicsEllipseItem()
func (this *QGraphicsEllipseItem) FreeQGraphicsEllipseItem(args ...interface{}) () {
  // ~QGraphicsEllipseItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsEllipseItemD0Ev
    // invoke: void ~QGraphicsEllipseItem()
    C.C_ZN20QGraphicsEllipseItemD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "~QGraphicsEllipseItem", args)
  }

}

// startAngle()
func (this *QGraphicsEllipseItem) startAngle(args ...interface{}) () {
  // startAngle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem10startAngleEv
    // invoke: int startAngle()
    var ret = C.C_ZNK20QGraphicsEllipseItem10startAngleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "startAngle", args)
  }

}

// isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsEllipseItem) isObscuredBy(args ...interface{}) () {
  // isObscuredBy(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem12isObscuredByEPK13QGraphicsItem
    // invoke: bool isObscuredBy(const class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK20QGraphicsEllipseItem12isObscuredByEPK13QGraphicsItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "isObscuredBy", args)
  }

}

// QGraphicsEllipseItem(class QGraphicsItem *)
func NewQGraphicsEllipseItem(args ...interface{}) *QGraphicsEllipseItem {
  // QGraphicsEllipseItem(class QGraphicsItem *)
  // QGraphicsEllipseItem(qreal, qreal, qreal, qreal, class QGraphicsItem *)
  // QGraphicsEllipseItem(const class QRectF &, class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][4] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[2][1] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsEllipseItemC1EP13QGraphicsItem
    // invoke: void QGraphicsEllipseItem(class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QGraphicsEllipseItemC2EP13QGraphicsItem(arg0)
    return &QGraphicsEllipseItem{qclsinst:qthis}
  case 1:
    // invoke: _ZN20QGraphicsEllipseItemC1EddddP13QGraphicsItem
    // invoke: void QGraphicsEllipseItem(qreal, qreal, qreal, qreal, class QGraphicsItem *)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg4)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QGraphicsEllipseItemC2EddddP13QGraphicsItem(arg0, arg1, arg2, arg3, arg4)
    return &QGraphicsEllipseItem{qclsinst:qthis}
  case 2:
    // invoke: _ZN20QGraphicsEllipseItemC1ERK6QRectFP13QGraphicsItem
    // invoke: void QGraphicsEllipseItem(const class QRectF &, class QGraphicsItem *)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QGraphicsEllipseItemC2ERK6QRectFP13QGraphicsItem(arg0, arg1)
    return &QGraphicsEllipseItem{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "QGraphicsEllipseItem", args)
  }

  return nil // QGraphicsEllipseItem{qclsinst:qthis}
}

// rect()
func (this *QGraphicsEllipseItem) rect(args ...interface{}) () {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem4rectEv
    // invoke: QRectF rect()
    var ret = C.C_ZNK20QGraphicsEllipseItem4rectEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "rect", args)
  }

}

// opaqueArea()
func (this *QGraphicsPolygonItem) opaqueArea(args ...interface{}) () {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsPolygonItem10opaqueAreaEv
    // invoke: QPainterPath opaqueArea()
    var ret = C.C_ZNK20QGraphicsPolygonItem10opaqueAreaEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "opaqueArea", args)
  }

}

// boundingRect()
func (this *QGraphicsPolygonItem) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsPolygonItem12boundingRectEv
    // invoke: QRectF boundingRect()
    var ret = C.C_ZNK20QGraphicsPolygonItem12boundingRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "boundingRect", args)
  }

}

// setPolygon(const class QPolygonF &)
func (this *QGraphicsPolygonItem) setPolygon(args ...interface{}) () {
  // setPolygon(const class QPolygonF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsPolygonItem10setPolygonERK9QPolygonF
    // invoke: void setPolygon(const class QPolygonF &)
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN20QGraphicsPolygonItem10setPolygonERK9QPolygonF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "setPolygon", args)
  }

}

// type()
func (this *QGraphicsPolygonItem) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsPolygonItem4typeEv
    // invoke: int type()
    var ret = C.C_ZNK20QGraphicsPolygonItem4typeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "type", args)
  }

}

// contains(const class QPointF &)
func (this *QGraphicsPolygonItem) contains(args ...interface{}) () {
  // contains(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsPolygonItem8containsERK7QPointF
    // invoke: bool contains(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK20QGraphicsPolygonItem8containsERK7QPointF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "contains", args)
  }

}

// polygon()
func (this *QGraphicsPolygonItem) polygon(args ...interface{}) () {
  // polygon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsPolygonItem7polygonEv
    // invoke: QPolygonF polygon()
    var ret = C.C_ZNK20QGraphicsPolygonItem7polygonEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "polygon", args)
  }

}

// paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsPolygonItem) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsPolygonItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
    // invoke: void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(QPainter).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStyleOptionGraphicsItem).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN20QGraphicsPolygonItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "paint", args)
  }

}

// shape()
func (this *QGraphicsPolygonItem) shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsPolygonItem5shapeEv
    // invoke: QPainterPath shape()
    var ret = C.C_ZNK20QGraphicsPolygonItem5shapeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "shape", args)
  }

}

// QGraphicsPolygonItem(const class QPolygonF &, class QGraphicsItem *)
func NewQGraphicsPolygonItem(args ...interface{}) *QGraphicsPolygonItem {
  // QGraphicsPolygonItem(const class QPolygonF &, class QGraphicsItem *)
  // QGraphicsPolygonItem(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[0][1] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsPolygonItemC1ERK9QPolygonFP13QGraphicsItem
    // invoke: void QGraphicsPolygonItem(const class QPolygonF &, class QGraphicsItem *)
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QGraphicsPolygonItemC2ERK9QPolygonFP13QGraphicsItem(arg0, arg1)
    return &QGraphicsPolygonItem{qclsinst:qthis}
  case 1:
    // invoke: _ZN20QGraphicsPolygonItemC1EP13QGraphicsItem
    // invoke: void QGraphicsPolygonItem(class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QGraphicsPolygonItemC2EP13QGraphicsItem(arg0)
    return &QGraphicsPolygonItem{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "QGraphicsPolygonItem", args)
  }

  return nil // QGraphicsPolygonItem{qclsinst:qthis}
}

// ~QGraphicsPolygonItem()
func (this *QGraphicsPolygonItem) FreeQGraphicsPolygonItem(args ...interface{}) () {
  // ~QGraphicsPolygonItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsPolygonItemD0Ev
    // invoke: void ~QGraphicsPolygonItem()
    C.C_ZN20QGraphicsPolygonItemD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "~QGraphicsPolygonItem", args)
  }

}

// isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsPolygonItem) isObscuredBy(args ...interface{}) () {
  // isObscuredBy(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsPolygonItem12isObscuredByEPK13QGraphicsItem
    // invoke: bool isObscuredBy(const class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK20QGraphicsPolygonItem12isObscuredByEPK13QGraphicsItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "isObscuredBy", args)
  }

}

// fillRule()
func (this *QGraphicsPolygonItem) fillRule(args ...interface{}) () {
  // fillRule()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsPolygonItem8fillRuleEv
    // invoke: Qt::FillRule fillRule()
    C.C_ZNK20QGraphicsPolygonItem8fillRuleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "fillRule", args)
  }

}

// opaqueArea()
func (this *QGraphicsPathItem) opaqueArea(args ...interface{}) () {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsPathItem10opaqueAreaEv
    // invoke: QPainterPath opaqueArea()
    var ret = C.C_ZNK17QGraphicsPathItem10opaqueAreaEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "opaqueArea", args)
  }

}

// ~QGraphicsPathItem()
func (this *QGraphicsPathItem) FreeQGraphicsPathItem(args ...interface{}) () {
  // ~QGraphicsPathItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsPathItemD0Ev
    // invoke: void ~QGraphicsPathItem()
    C.C_ZN17QGraphicsPathItemD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "~QGraphicsPathItem", args)
  }

}

// boundingRect()
func (this *QGraphicsPathItem) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsPathItem12boundingRectEv
    // invoke: QRectF boundingRect()
    var ret = C.C_ZNK17QGraphicsPathItem12boundingRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "boundingRect", args)
  }

}

// isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsPathItem) isObscuredBy(args ...interface{}) () {
  // isObscuredBy(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsPathItem12isObscuredByEPK13QGraphicsItem
    // invoke: bool isObscuredBy(const class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK17QGraphicsPathItem12isObscuredByEPK13QGraphicsItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "isObscuredBy", args)
  }

}

// contains(const class QPointF &)
func (this *QGraphicsPathItem) contains(args ...interface{}) () {
  // contains(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsPathItem8containsERK7QPointF
    // invoke: bool contains(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK17QGraphicsPathItem8containsERK7QPointF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "contains", args)
  }

}

// paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsPathItem) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsPathItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
    // invoke: void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(QPainter).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStyleOptionGraphicsItem).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN17QGraphicsPathItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "paint", args)
  }

}

// shape()
func (this *QGraphicsPathItem) shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsPathItem5shapeEv
    // invoke: QPainterPath shape()
    var ret = C.C_ZNK17QGraphicsPathItem5shapeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "shape", args)
  }

}

// path()
func (this *QGraphicsPathItem) path(args ...interface{}) () {
  // path()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsPathItem4pathEv
    // invoke: QPainterPath path()
    var ret = C.C_ZNK17QGraphicsPathItem4pathEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "path", args)
  }

}

// setPath(const class QPainterPath &)
func (this *QGraphicsPathItem) setPath(args ...interface{}) () {
  // setPath(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsPathItem7setPathERK12QPainterPath
    // invoke: void setPath(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsPathItem7setPathERK12QPainterPath(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "setPath", args)
  }

}

// QGraphicsPathItem(const class QPainterPath &, class QGraphicsItem *)
func NewQGraphicsPathItem(args ...interface{}) *QGraphicsPathItem {
  // QGraphicsPathItem(const class QPainterPath &, class QGraphicsItem *)
  // QGraphicsPathItem(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[0][1] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsPathItemC1ERK12QPainterPathP13QGraphicsItem
    // invoke: void QGraphicsPathItem(const class QPainterPath &, class QGraphicsItem *)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QGraphicsPathItemC2ERK12QPainterPathP13QGraphicsItem(arg0, arg1)
    return &QGraphicsPathItem{qclsinst:qthis}
  case 1:
    // invoke: _ZN17QGraphicsPathItemC1EP13QGraphicsItem
    // invoke: void QGraphicsPathItem(class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QGraphicsPathItemC2EP13QGraphicsItem(arg0)
    return &QGraphicsPathItem{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "QGraphicsPathItem", args)
  }

  return nil // QGraphicsPathItem{qclsinst:qthis}
}

// type()
func (this *QGraphicsPathItem) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsPathItem4typeEv
    // invoke: int type()
    var ret = C.C_ZNK17QGraphicsPathItem4typeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "type", args)
  }

}

// opaqueArea()
func (this *QGraphicsLineItem) opaqueArea(args ...interface{}) () {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsLineItem10opaqueAreaEv
    // invoke: QPainterPath opaqueArea()
    var ret = C.C_ZNK17QGraphicsLineItem10opaqueAreaEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "opaqueArea", args)
  }

}

// boundingRect()
func (this *QGraphicsLineItem) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsLineItem12boundingRectEv
    // invoke: QRectF boundingRect()
    var ret = C.C_ZNK17QGraphicsLineItem12boundingRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "boundingRect", args)
  }

}

// type()
func (this *QGraphicsLineItem) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsLineItem4typeEv
    // invoke: int type()
    var ret = C.C_ZNK17QGraphicsLineItem4typeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "type", args)
  }

}

// contains(const class QPointF &)
func (this *QGraphicsLineItem) contains(args ...interface{}) () {
  // contains(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsLineItem8containsERK7QPointF
    // invoke: bool contains(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK17QGraphicsLineItem8containsERK7QPointF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "contains", args)
  }

}

// shape()
func (this *QGraphicsLineItem) shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsLineItem5shapeEv
    // invoke: QPainterPath shape()
    var ret = C.C_ZNK17QGraphicsLineItem5shapeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "shape", args)
  }

}

// ~QGraphicsLineItem()
func (this *QGraphicsLineItem) FreeQGraphicsLineItem(args ...interface{}) () {
  // ~QGraphicsLineItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsLineItemD0Ev
    // invoke: void ~QGraphicsLineItem()
    C.C_ZN17QGraphicsLineItemD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "~QGraphicsLineItem", args)
  }

}

// QGraphicsLineItem(class QGraphicsItem *)
func NewQGraphicsLineItem(args ...interface{}) *QGraphicsLineItem {
  // QGraphicsLineItem(class QGraphicsItem *)
  // QGraphicsLineItem(const class QLineF &, class QGraphicsItem *)
  // QGraphicsLineItem(qreal, qreal, qreal, qreal, class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QLineF{}) // "const QLineF &"
  vtys[1][1] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][4] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsLineItemC1EP13QGraphicsItem
    // invoke: void QGraphicsLineItem(class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QGraphicsLineItemC2EP13QGraphicsItem(arg0)
    return &QGraphicsLineItem{qclsinst:qthis}
  case 1:
    // invoke: _ZN17QGraphicsLineItemC1ERK6QLineFP13QGraphicsItem
    // invoke: void QGraphicsLineItem(const class QLineF &, class QGraphicsItem *)
    var arg0 = args[0].(QLineF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QGraphicsLineItemC2ERK6QLineFP13QGraphicsItem(arg0, arg1)
    return &QGraphicsLineItem{qclsinst:qthis}
  case 2:
    // invoke: _ZN17QGraphicsLineItemC1EddddP13QGraphicsItem
    // invoke: void QGraphicsLineItem(qreal, qreal, qreal, qreal, class QGraphicsItem *)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg4)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QGraphicsLineItemC2EddddP13QGraphicsItem(arg0, arg1, arg2, arg3, arg4)
    return &QGraphicsLineItem{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "QGraphicsLineItem", args)
  }

  return nil // QGraphicsLineItem{qclsinst:qthis}
}

// paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsLineItem) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsLineItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
    // invoke: void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(QPainter).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStyleOptionGraphicsItem).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN17QGraphicsLineItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "paint", args)
  }

}

// pen()
func (this *QGraphicsLineItem) pen(args ...interface{}) () {
  // pen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsLineItem3penEv
    // invoke: QPen pen()
    var ret = C.C_ZNK17QGraphicsLineItem3penEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "pen", args)
  }

}

// setPen(const class QPen &)
func (this *QGraphicsLineItem) setPen(args ...interface{}) () {
  // setPen(const class QPen &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPen{}) // "const QPen &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsLineItem6setPenERK4QPen
    // invoke: void setPen(const class QPen &)
    var arg0 = args[0].(QPen).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsLineItem6setPenERK4QPen(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "setPen", args)
  }

}

// setLine(const class QLineF &)
func (this *QGraphicsLineItem) setLine(args ...interface{}) () {
  // setLine(const class QLineF &)
  // setLine(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLineF{}) // "const QLineF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsLineItem7setLineERK6QLineF
    // invoke: void setLine(const class QLineF &)
    var arg0 = args[0].(QLineF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsLineItem7setLineERK6QLineF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN17QGraphicsLineItem7setLineEdddd
    // invoke: void setLine(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN17QGraphicsLineItem7setLineEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "setLine", args)
  }

}

// line()
func (this *QGraphicsLineItem) line(args ...interface{}) () {
  // line()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsLineItem4lineEv
    // invoke: QLineF line()
    var ret = C.C_ZNK17QGraphicsLineItem4lineEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "line", args)
  }

}

// isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsLineItem) isObscuredBy(args ...interface{}) () {
  // isObscuredBy(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsLineItem12isObscuredByEPK13QGraphicsItem
    // invoke: bool isObscuredBy(const class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK17QGraphicsLineItem12isObscuredByEPK13QGraphicsItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "isObscuredBy", args)
  }

}

// opaqueArea()
func (this *QGraphicsItemGroup) opaqueArea(args ...interface{}) () {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QGraphicsItemGroup10opaqueAreaEv
    // invoke: QPainterPath opaqueArea()
    var ret = C.C_ZNK18QGraphicsItemGroup10opaqueAreaEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "opaqueArea", args)
  }

}

// addToGroup(class QGraphicsItem *)
func (this *QGraphicsItemGroup) addToGroup(args ...interface{}) () {
  // addToGroup(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QGraphicsItemGroup10addToGroupEP13QGraphicsItem
    // invoke: void addToGroup(class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QGraphicsItemGroup10addToGroupEP13QGraphicsItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "addToGroup", args)
  }

}

// QGraphicsItemGroup(class QGraphicsItem *)
func NewQGraphicsItemGroup(args ...interface{}) *QGraphicsItemGroup {
  // QGraphicsItemGroup(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QGraphicsItemGroupC1EP13QGraphicsItem
    // invoke: void QGraphicsItemGroup(class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QGraphicsItemGroupC2EP13QGraphicsItem(arg0)
    return &QGraphicsItemGroup{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "QGraphicsItemGroup", args)
  }

  return nil // QGraphicsItemGroup{qclsinst:qthis}
}

// boundingRect()
func (this *QGraphicsItemGroup) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QGraphicsItemGroup12boundingRectEv
    // invoke: QRectF boundingRect()
    var ret = C.C_ZNK18QGraphicsItemGroup12boundingRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "boundingRect", args)
  }

}

// removeFromGroup(class QGraphicsItem *)
func (this *QGraphicsItemGroup) removeFromGroup(args ...interface{}) () {
  // removeFromGroup(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QGraphicsItemGroup15removeFromGroupEP13QGraphicsItem
    // invoke: void removeFromGroup(class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QGraphicsItemGroup15removeFromGroupEP13QGraphicsItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "removeFromGroup", args)
  }

}

// isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsItemGroup) isObscuredBy(args ...interface{}) () {
  // isObscuredBy(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QGraphicsItemGroup12isObscuredByEPK13QGraphicsItem
    // invoke: bool isObscuredBy(const class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK18QGraphicsItemGroup12isObscuredByEPK13QGraphicsItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "isObscuredBy", args)
  }

}

// paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsItemGroup) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QGraphicsItemGroup5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
    // invoke: void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(QPainter).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStyleOptionGraphicsItem).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN18QGraphicsItemGroup5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "paint", args)
  }

}

// type()
func (this *QGraphicsItemGroup) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QGraphicsItemGroup4typeEv
    // invoke: int type()
    var ret = C.C_ZNK18QGraphicsItemGroup4typeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "type", args)
  }

}

// ~QGraphicsItemGroup()
func (this *QGraphicsItemGroup) FreeQGraphicsItemGroup(args ...interface{}) () {
  // ~QGraphicsItemGroup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QGraphicsItemGroupD0Ev
    // invoke: void ~QGraphicsItemGroup()
    C.C_ZN18QGraphicsItemGroupD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "~QGraphicsItemGroup", args)
  }

}

// opaqueArea()
func (this *QAbstractGraphicsShapeItem) opaqueArea(args ...interface{}) () {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QAbstractGraphicsShapeItem10opaqueAreaEv
    // invoke: QPainterPath opaqueArea()
    var ret = C.C_ZNK26QAbstractGraphicsShapeItem10opaqueAreaEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "opaqueArea", args)
  }

}

// isObscuredBy(const class QGraphicsItem *)
func (this *QAbstractGraphicsShapeItem) isObscuredBy(args ...interface{}) () {
  // isObscuredBy(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QAbstractGraphicsShapeItem12isObscuredByEPK13QGraphicsItem
    // invoke: bool isObscuredBy(const class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK26QAbstractGraphicsShapeItem12isObscuredByEPK13QGraphicsItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "isObscuredBy", args)
  }

}

// pen()
func (this *QAbstractGraphicsShapeItem) pen(args ...interface{}) () {
  // pen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QAbstractGraphicsShapeItem3penEv
    // invoke: QPen pen()
    var ret = C.C_ZNK26QAbstractGraphicsShapeItem3penEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "pen", args)
  }

}

// brush()
func (this *QAbstractGraphicsShapeItem) brush(args ...interface{}) () {
  // brush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QAbstractGraphicsShapeItem5brushEv
    // invoke: QBrush brush()
    var ret = C.C_ZNK26QAbstractGraphicsShapeItem5brushEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "brush", args)
  }

}

// setPen(const class QPen &)
func (this *QAbstractGraphicsShapeItem) setPen(args ...interface{}) () {
  // setPen(const class QPen &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPen{}) // "const QPen &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAbstractGraphicsShapeItem6setPenERK4QPen
    // invoke: void setPen(const class QPen &)
    var arg0 = args[0].(QPen).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN26QAbstractGraphicsShapeItem6setPenERK4QPen(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "setPen", args)
  }

}

// setBrush(const class QBrush &)
func (this *QAbstractGraphicsShapeItem) setBrush(args ...interface{}) () {
  // setBrush(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAbstractGraphicsShapeItem8setBrushERK6QBrush
    // invoke: void setBrush(const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN26QAbstractGraphicsShapeItem8setBrushERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "setBrush", args)
  }

}

// ~QAbstractGraphicsShapeItem()
func (this *QAbstractGraphicsShapeItem) FreeQAbstractGraphicsShapeItem(args ...interface{}) () {
  // ~QAbstractGraphicsShapeItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAbstractGraphicsShapeItemD0Ev
    // invoke: void ~QAbstractGraphicsShapeItem()
    C.C_ZN26QAbstractGraphicsShapeItemD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "~QAbstractGraphicsShapeItem", args)
  }

}

// QAbstractGraphicsShapeItem(class QGraphicsItem *)
func NewQAbstractGraphicsShapeItem(args ...interface{}) *QAbstractGraphicsShapeItem {
  // QAbstractGraphicsShapeItem(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAbstractGraphicsShapeItemC1EP13QGraphicsItem
    // invoke: void QAbstractGraphicsShapeItem(class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN26QAbstractGraphicsShapeItemC2EP13QGraphicsItem(arg0)
    return &QAbstractGraphicsShapeItem{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "QAbstractGraphicsShapeItem", args)
  }

  return nil // QAbstractGraphicsShapeItem{qclsinst:qthis}
}

// show()
func (this *QGraphicsItem) show(args ...interface{}) () {
  // show()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem4showEv
    // invoke: void show()
    C.C_ZN13QGraphicsItem4showEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "show", args)
  }

}

// hasFocus()
func (this *QGraphicsItem) hasFocus(args ...interface{}) () {
  // hasFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8hasFocusEv
    // invoke: bool hasFocus()
    var ret = C.C_ZNK13QGraphicsItem8hasFocusEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "hasFocus", args)
  }

}

// setGroup(class QGraphicsItemGroup *)
func (this *QGraphicsItem) setGroup(args ...interface{}) () {
  // setGroup(class QGraphicsItemGroup *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItemGroup{}) // "QGraphicsItemGroup *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem8setGroupEP18QGraphicsItemGroup
    // invoke: void setGroup(class QGraphicsItemGroup *)
    var arg0 = args[0].(QGraphicsItemGroup).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem8setGroupEP18QGraphicsItemGroup(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setGroup", args)
  }

}

// handlesChildEvents()
func (this *QGraphicsItem) handlesChildEvents(args ...interface{}) () {
  // handlesChildEvents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem18handlesChildEventsEv
    // invoke: bool handlesChildEvents()
    var ret = C.C_ZNK13QGraphicsItem18handlesChildEventsEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "handlesChildEvents", args)
  }

}

// shape()
func (this *QGraphicsItem) shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem5shapeEv
    // invoke: QPainterPath shape()
    var ret = C.C_ZNK13QGraphicsItem5shapeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "shape", args)
  }

}

// isPanel()
func (this *QGraphicsItem) isPanel(args ...interface{}) () {
  // isPanel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem7isPanelEv
    // invoke: bool isPanel()
    var ret = C.C_ZNK13QGraphicsItem7isPanelEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isPanel", args)
  }

}

// data(int)
func (this *QGraphicsItem) data(args ...interface{}) () {
  // data(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem4dataEi
    // invoke: QVariant data(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem4dataEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "data", args)
  }

}

// parentItem()
func (this *QGraphicsItem) parentItem(args ...interface{}) () {
  // parentItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem10parentItemEv
    // invoke: QGraphicsItem * parentItem()
    C.C_ZNK13QGraphicsItem10parentItemEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "parentItem", args)
  }

}

// childrenBoundingRect()
func (this *QGraphicsItem) childrenBoundingRect(args ...interface{}) () {
  // childrenBoundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem20childrenBoundingRectEv
    // invoke: QRectF childrenBoundingRect()
    var ret = C.C_ZNK13QGraphicsItem20childrenBoundingRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "childrenBoundingRect", args)
  }

}

// group()
func (this *QGraphicsItem) group(args ...interface{}) () {
  // group()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem5groupEv
    // invoke: QGraphicsItemGroup * group()
    C.C_ZNK13QGraphicsItem5groupEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "group", args)
  }

}

// isEnabled()
func (this *QGraphicsItem) isEnabled(args ...interface{}) () {
  // isEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9isEnabledEv
    // invoke: bool isEnabled()
    var ret = C.C_ZNK13QGraphicsItem9isEnabledEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isEnabled", args)
  }

}

// window()
func (this *QGraphicsItem) window(args ...interface{}) () {
  // window()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem6windowEv
    // invoke: QGraphicsWidget * window()
    C.C_ZNK13QGraphicsItem6windowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "window", args)
  }

}

// setBoundingRegionGranularity(qreal)
func (this *QGraphicsItem) setBoundingRegionGranularity(args ...interface{}) () {
  // setBoundingRegionGranularity(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem28setBoundingRegionGranularityEd
    // invoke: void setBoundingRegionGranularity(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem28setBoundingRegionGranularityEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setBoundingRegionGranularity", args)
  }

}

// setRotation(qreal)
func (this *QGraphicsItem) setRotation(args ...interface{}) () {
  // setRotation(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem11setRotationEd
    // invoke: void setRotation(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem11setRotationEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setRotation", args)
  }

}

// commonAncestorItem(const class QGraphicsItem *)
func (this *QGraphicsItem) commonAncestorItem(args ...interface{}) () {
  // commonAncestorItem(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem18commonAncestorItemEPKS_
    // invoke: QGraphicsItem * commonAncestorItem(const class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK13QGraphicsItem18commonAncestorItemEPKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "commonAncestorItem", args)
  }

}

// mapToItem(const class QGraphicsItem *, const class QPolygonF &)
func (this *QGraphicsItem) mapToItem(args ...interface{}) () {
  // mapToItem(const class QGraphicsItem *, const class QPolygonF &)
  // mapToItem(const class QGraphicsItem *, const class QRectF &)
  // mapToItem(const class QGraphicsItem *, const class QPainterPath &)
  // mapToItem(const class QGraphicsItem *, const class QPointF &)
  // mapToItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
  // mapToItem(const class QGraphicsItem *, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[0][1] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[1][1] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[2][1] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[3][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[4][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[4][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[4][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[4][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[5][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[5][2] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9mapToItemEPKS_RK9QPolygonF
    // invoke: QPolygonF mapToItem(const class QGraphicsItem *, const class QPolygonF &)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPolygonF).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK13QGraphicsItem9mapToItemEPKS_RK9QPolygonF(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK13QGraphicsItem9mapToItemEPKS_RK6QRectF
    // invoke: QPolygonF mapToItem(const class QGraphicsItem *, const class QRectF &)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QRectF).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK13QGraphicsItem9mapToItemEPKS_RK6QRectF(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 2:
    // invoke: _ZNK13QGraphicsItem9mapToItemEPKS_RK12QPainterPath
    // invoke: QPainterPath mapToItem(const class QGraphicsItem *, const class QPainterPath &)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPainterPath).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK13QGraphicsItem9mapToItemEPKS_RK12QPainterPath(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 3:
    // invoke: _ZNK13QGraphicsItem9mapToItemEPKS_RK7QPointF
    // invoke: QPointF mapToItem(const class QGraphicsItem *, const class QPointF &)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK13QGraphicsItem9mapToItemEPKS_RK7QPointF(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 4:
    // invoke: _ZNK13QGraphicsItem9mapToItemEPKS_dddd
    // invoke: QPolygonF mapToItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.double(args[4].(float64))
    if false {fmt.Println(arg4)}
    var ret = C.C_ZNK13QGraphicsItem9mapToItemEPKS_dddd(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret)}
  case 5:
    // invoke: _ZNK13QGraphicsItem9mapToItemEPKS_dd
    // invoke: QPointF mapToItem(const class QGraphicsItem *, qreal, qreal)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var ret = C.C_ZNK13QGraphicsItem9mapToItemEPKS_dd(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapToItem", args)
  }

}

// toGraphicsObject()
func (this *QGraphicsItem) toGraphicsObject(args ...interface{}) () {
  // toGraphicsObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem16toGraphicsObjectEv
    // invoke: QGraphicsObject * toGraphicsObject()
    C.C_ZN13QGraphicsItem16toGraphicsObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "toGraphicsObject", args)
  }

}

// acceptedMouseButtons()
func (this *QGraphicsItem) acceptedMouseButtons(args ...interface{}) () {
  // acceptedMouseButtons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem20acceptedMouseButtonsEv
    // invoke: Qt::MouseButtons acceptedMouseButtons()
    C.C_ZNK13QGraphicsItem20acceptedMouseButtonsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "acceptedMouseButtons", args)
  }

}

// setGraphicsEffect(class QGraphicsEffect *)
func (this *QGraphicsItem) setGraphicsEffect(args ...interface{}) () {
  // setGraphicsEffect(class QGraphicsEffect *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsEffect{}) // "QGraphicsEffect *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem17setGraphicsEffectEP15QGraphicsEffect
    // invoke: void setGraphicsEffect(class QGraphicsEffect *)
    var arg0 = args[0].(QGraphicsEffect).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem17setGraphicsEffectEP15QGraphicsEffect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setGraphicsEffect", args)
  }

}

// mapToParent(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) mapToParent(args ...interface{}) () {
  // mapToParent(qreal, qreal, qreal, qreal)
  // mapToParent(const class QPointF &)
  // mapToParent(qreal, qreal)
  // mapToParent(const class QPainterPath &)
  // mapToParent(const class QRectF &)
  // mapToParent(const class QPolygonF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem11mapToParentEdddd
    // invoke: QPolygonF mapToParent(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var ret = C.C_ZNK13QGraphicsItem11mapToParentEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK13QGraphicsItem11mapToParentERK7QPointF
    // invoke: QPointF mapToParent(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem11mapToParentERK7QPointF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 2:
    // invoke: _ZNK13QGraphicsItem11mapToParentEdd
    // invoke: QPointF mapToParent(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK13QGraphicsItem11mapToParentEdd(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 3:
    // invoke: _ZNK13QGraphicsItem11mapToParentERK12QPainterPath
    // invoke: QPainterPath mapToParent(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem11mapToParentERK12QPainterPath(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 4:
    // invoke: _ZNK13QGraphicsItem11mapToParentERK6QRectF
    // invoke: QPolygonF mapToParent(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem11mapToParentERK6QRectF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 5:
    // invoke: _ZNK13QGraphicsItem11mapToParentERK9QPolygonF
    // invoke: QPolygonF mapToParent(const class QPolygonF &)
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem11mapToParentERK9QPolygonF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapToParent", args)
  }

}

// setTransformOriginPoint(const class QPointF &)
func (this *QGraphicsItem) setTransformOriginPoint(args ...interface{}) () {
  // setTransformOriginPoint(const class QPointF &)
  // setTransformOriginPoint(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem23setTransformOriginPointERK7QPointF
    // invoke: void setTransformOriginPoint(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem23setTransformOriginPointERK7QPointF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN13QGraphicsItem23setTransformOriginPointEdd
    // invoke: void setTransformOriginPoint(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN13QGraphicsItem23setTransformOriginPointEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setTransformOriginPoint", args)
  }

}

// deviceTransform(const class QTransform &)
func (this *QGraphicsItem) deviceTransform(args ...interface{}) () {
  // deviceTransform(const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem15deviceTransformERK10QTransform
    // invoke: QTransform deviceTransform(const class QTransform &)
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem15deviceTransformERK10QTransform(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "deviceTransform", args)
  }

}

// rotation()
func (this *QGraphicsItem) rotation(args ...interface{}) () {
  // rotation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8rotationEv
    // invoke: qreal rotation()
    var ret = C.C_ZNK13QGraphicsItem8rotationEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "rotation", args)
  }

}

// effectiveOpacity()
func (this *QGraphicsItem) effectiveOpacity(args ...interface{}) () {
  // effectiveOpacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem16effectiveOpacityEv
    // invoke: qreal effectiveOpacity()
    var ret = C.C_ZNK13QGraphicsItem16effectiveOpacityEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "effectiveOpacity", args)
  }

}

// focusScopeItem()
func (this *QGraphicsItem) focusScopeItem(args ...interface{}) () {
  // focusScopeItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem14focusScopeItemEv
    // invoke: QGraphicsItem * focusScopeItem()
    C.C_ZNK13QGraphicsItem14focusScopeItemEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "focusScopeItem", args)
  }

}

// inputMethodHints()
func (this *QGraphicsItem) inputMethodHints(args ...interface{}) () {
  // inputMethodHints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem16inputMethodHintsEv
    // invoke: Qt::InputMethodHints inputMethodHints()
    C.C_ZNK13QGraphicsItem16inputMethodHintsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "inputMethodHints", args)
  }

}

// panelModality()
func (this *QGraphicsItem) panelModality(args ...interface{}) () {
  // panelModality()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem13panelModalityEv
    // invoke: QGraphicsItem::PanelModality panelModality()
    C.C_ZNK13QGraphicsItem13panelModalityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "panelModality", args)
  }

}

// topLevelWidget()
func (this *QGraphicsItem) topLevelWidget(args ...interface{}) () {
  // topLevelWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem14topLevelWidgetEv
    // invoke: QGraphicsWidget * topLevelWidget()
    C.C_ZNK13QGraphicsItem14topLevelWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "topLevelWidget", args)
  }

}

// isVisibleTo(const class QGraphicsItem *)
func (this *QGraphicsItem) isVisibleTo(args ...interface{}) () {
  // isVisibleTo(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem11isVisibleToEPKS_
    // invoke: bool isVisibleTo(const class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem11isVisibleToEPKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isVisibleTo", args)
  }

}

// setAcceptDrops(_Bool)
func (this *QGraphicsItem) setAcceptDrops(args ...interface{}) () {
  // setAcceptDrops(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem14setAcceptDropsEb
    // invoke: void setAcceptDrops(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem14setAcceptDropsEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setAcceptDrops", args)
  }

}

// clearFocus()
func (this *QGraphicsItem) clearFocus(args ...interface{}) () {
  // clearFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem10clearFocusEv
    // invoke: void clearFocus()
    C.C_ZN13QGraphicsItem10clearFocusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "clearFocus", args)
  }

}

// filtersChildEvents()
func (this *QGraphicsItem) filtersChildEvents(args ...interface{}) () {
  // filtersChildEvents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem18filtersChildEventsEv
    // invoke: bool filtersChildEvents()
    var ret = C.C_ZNK13QGraphicsItem18filtersChildEventsEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "filtersChildEvents", args)
  }

}

// x()
func (this *QGraphicsItem) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem1xEv
    // invoke: qreal x()
    C.C_ZNK13QGraphicsItem1xEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "x", args)
  }

}

// setTransform(const class QTransform &, _Bool)
func (this *QGraphicsItem) setTransform(args ...interface{}) () {
  // setTransform(const class QTransform &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem12setTransformERK10QTransformb
    // invoke: void setTransform(const class QTransform &, _Bool)
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN13QGraphicsItem12setTransformERK10QTransformb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setTransform", args)
  }

}

// stackBefore(const class QGraphicsItem *)
func (this *QGraphicsItem) stackBefore(args ...interface{}) () {
  // stackBefore(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem11stackBeforeEPKS_
    // invoke: void stackBefore(const class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem11stackBeforeEPKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "stackBefore", args)
  }

}

// setAcceptHoverEvents(_Bool)
func (this *QGraphicsItem) setAcceptHoverEvents(args ...interface{}) () {
  // setAcceptHoverEvents(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem20setAcceptHoverEventsEb
    // invoke: void setAcceptHoverEvents(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem20setAcceptHoverEventsEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setAcceptHoverEvents", args)
  }

}

// focusProxy()
func (this *QGraphicsItem) focusProxy(args ...interface{}) () {
  // focusProxy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem10focusProxyEv
    // invoke: QGraphicsItem * focusProxy()
    C.C_ZNK13QGraphicsItem10focusProxyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "focusProxy", args)
  }

}

// mapFromParent(const class QPainterPath &)
func (this *QGraphicsItem) mapFromParent(args ...interface{}) () {
  // mapFromParent(const class QPainterPath &)
  // mapFromParent(const class QRectF &)
  // mapFromParent(qreal, qreal, qreal, qreal)
  // mapFromParent(qreal, qreal)
  // mapFromParent(const class QPointF &)
  // mapFromParent(const class QPolygonF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem13mapFromParentERK12QPainterPath
    // invoke: QPainterPath mapFromParent(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem13mapFromParentERK12QPainterPath(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK13QGraphicsItem13mapFromParentERK6QRectF
    // invoke: QPolygonF mapFromParent(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem13mapFromParentERK6QRectF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 2:
    // invoke: _ZNK13QGraphicsItem13mapFromParentEdddd
    // invoke: QPolygonF mapFromParent(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var ret = C.C_ZNK13QGraphicsItem13mapFromParentEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  case 3:
    // invoke: _ZNK13QGraphicsItem13mapFromParentEdd
    // invoke: QPointF mapFromParent(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK13QGraphicsItem13mapFromParentEdd(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 4:
    // invoke: _ZNK13QGraphicsItem13mapFromParentERK7QPointF
    // invoke: QPointF mapFromParent(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem13mapFromParentERK7QPointF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 5:
    // invoke: _ZNK13QGraphicsItem13mapFromParentERK9QPolygonF
    // invoke: QPolygonF mapFromParent(const class QPolygonF &)
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem13mapFromParentERK9QPolygonF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapFromParent", args)
  }

}

// scenePos()
func (this *QGraphicsItem) scenePos(args ...interface{}) () {
  // scenePos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8scenePosEv
    // invoke: QPointF scenePos()
    var ret = C.C_ZNK13QGraphicsItem8scenePosEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "scenePos", args)
  }

}

// setPos(qreal, qreal)
func (this *QGraphicsItem) setPos(args ...interface{}) () {
  // setPos(qreal, qreal)
  // setPos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem6setPosEdd
    // invoke: void setPos(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN13QGraphicsItem6setPosEdd(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN13QGraphicsItem6setPosERK7QPointF
    // invoke: void setPos(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem6setPosERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setPos", args)
  }

}

// scale()
func (this *QGraphicsItem) scale(args ...interface{}) () {
  // scale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem5scaleEv
    // invoke: qreal scale()
    var ret = C.C_ZNK13QGraphicsItem5scaleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "scale", args)
  }

}

// hide()
func (this *QGraphicsItem) hide(args ...interface{}) () {
  // hide()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem4hideEv
    // invoke: void hide()
    C.C_ZN13QGraphicsItem4hideEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "hide", args)
  }

}

// matrix()
func (this *QGraphicsItem) matrix(args ...interface{}) () {
  // matrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem6matrixEv
    // invoke: QMatrix matrix()
    var ret = C.C_ZNK13QGraphicsItem6matrixEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "matrix", args)
  }

}

// isSelected()
func (this *QGraphicsItem) isSelected(args ...interface{}) () {
  // isSelected()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem10isSelectedEv
    // invoke: bool isSelected()
    var ret = C.C_ZNK13QGraphicsItem10isSelectedEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isSelected", args)
  }

}

// moveBy(qreal, qreal)
func (this *QGraphicsItem) moveBy(args ...interface{}) () {
  // moveBy(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem6moveByEdd
    // invoke: void moveBy(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN13QGraphicsItem6moveByEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "moveBy", args)
  }

}

// setFocusProxy(class QGraphicsItem *)
func (this *QGraphicsItem) setFocusProxy(args ...interface{}) () {
  // setFocusProxy(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem13setFocusProxyEPS_
    // invoke: void setFocusProxy(class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem13setFocusProxyEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setFocusProxy", args)
  }

}

// itemTransform(const class QGraphicsItem *, _Bool *)
func (this *QGraphicsItem) itemTransform(args ...interface{}) () {
  // itemTransform(const class QGraphicsItem *, _Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[0][1] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem13itemTransformEPKS_Pb
    // invoke: QTransform itemTransform(const class QGraphicsItem *, _Bool *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.bool)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK13QGraphicsItem13itemTransformEPKS_Pb(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "itemTransform", args)
  }

}

// mapRectToParent(const class QRectF &)
func (this *QGraphicsItem) mapRectToParent(args ...interface{}) () {
  // mapRectToParent(const class QRectF &)
  // mapRectToParent(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem15mapRectToParentERK6QRectF
    // invoke: QRectF mapRectToParent(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem15mapRectToParentERK6QRectF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK13QGraphicsItem15mapRectToParentEdddd
    // invoke: QRectF mapRectToParent(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var ret = C.C_ZNK13QGraphicsItem15mapRectToParentEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapRectToParent", args)
  }

}

// isBlockedByModalPanel(class QGraphicsItem **)
func (this *QGraphicsItem) isBlockedByModalPanel(args ...interface{}) () {
  // isBlockedByModalPanel(class QGraphicsItem **)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem **"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem21isBlockedByModalPanelEPPS_
    // invoke: bool isBlockedByModalPanel(class QGraphicsItem **)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem21isBlockedByModalPanelEPPS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isBlockedByModalPanel", args)
  }

}

// graphicsEffect()
func (this *QGraphicsItem) graphicsEffect(args ...interface{}) () {
  // graphicsEffect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem14graphicsEffectEv
    // invoke: QGraphicsEffect * graphicsEffect()
    C.C_ZNK13QGraphicsItem14graphicsEffectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "graphicsEffect", args)
  }

}

// childItems()
func (this *QGraphicsItem) childItems(args ...interface{}) () {
  // childItems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem10childItemsEv
    // invoke: QList<QGraphicsItem *> childItems()
    C.C_ZNK13QGraphicsItem10childItemsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "childItems", args)
  }

}

// isClipped()
func (this *QGraphicsItem) isClipped(args ...interface{}) () {
  // isClipped()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9isClippedEv
    // invoke: bool isClipped()
    var ret = C.C_ZNK13QGraphicsItem9isClippedEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isClipped", args)
  }

}

// toolTip()
func (this *QGraphicsItem) toolTip(args ...interface{}) () {
  // toolTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem7toolTipEv
    // invoke: QString toolTip()
    var ret = C.C_ZNK13QGraphicsItem7toolTipEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "toolTip", args)
  }

}

// ~QGraphicsItem()
func (this *QGraphicsItem) FreeQGraphicsItem(args ...interface{}) () {
  // ~QGraphicsItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItemD0Ev
    // invoke: void ~QGraphicsItem()
    C.C_ZN13QGraphicsItemD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "~QGraphicsItem", args)
  }

}

// isActive()
func (this *QGraphicsItem) isActive(args ...interface{}) () {
  // isActive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8isActiveEv
    // invoke: bool isActive()
    var ret = C.C_ZNK13QGraphicsItem8isActiveEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isActive", args)
  }

}

// setFiltersChildEvents(_Bool)
func (this *QGraphicsItem) setFiltersChildEvents(args ...interface{}) () {
  // setFiltersChildEvents(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem21setFiltersChildEventsEb
    // invoke: void setFiltersChildEvents(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem21setFiltersChildEventsEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setFiltersChildEvents", args)
  }

}

// opaqueArea()
func (this *QGraphicsItem) opaqueArea(args ...interface{}) () {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem10opaqueAreaEv
    // invoke: QPainterPath opaqueArea()
    var ret = C.C_ZNK13QGraphicsItem10opaqueAreaEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "opaqueArea", args)
  }

}

// advance(int)
func (this *QGraphicsItem) advance(args ...interface{}) () {
  // advance(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem7advanceEi
    // invoke: void advance(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem7advanceEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "advance", args)
  }

}

// setToolTip(const class QString &)
func (this *QGraphicsItem) setToolTip(args ...interface{}) () {
  // setToolTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem10setToolTipERK7QString
    // invoke: void setToolTip(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem10setToolTipERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setToolTip", args)
  }

}

// mapToScene(const class QPolygonF &)
func (this *QGraphicsItem) mapToScene(args ...interface{}) () {
  // mapToScene(const class QPolygonF &)
  // mapToScene(const class QRectF &)
  // mapToScene(const class QPointF &)
  // mapToScene(const class QPainterPath &)
  // mapToScene(qreal, qreal, qreal, qreal)
  // mapToScene(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[4][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[4][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[4][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[5][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem10mapToSceneERK9QPolygonF
    // invoke: QPolygonF mapToScene(const class QPolygonF &)
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem10mapToSceneERK9QPolygonF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK13QGraphicsItem10mapToSceneERK6QRectF
    // invoke: QPolygonF mapToScene(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem10mapToSceneERK6QRectF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 2:
    // invoke: _ZNK13QGraphicsItem10mapToSceneERK7QPointF
    // invoke: QPointF mapToScene(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem10mapToSceneERK7QPointF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 3:
    // invoke: _ZNK13QGraphicsItem10mapToSceneERK12QPainterPath
    // invoke: QPainterPath mapToScene(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem10mapToSceneERK12QPainterPath(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 4:
    // invoke: _ZNK13QGraphicsItem10mapToSceneEdddd
    // invoke: QPolygonF mapToScene(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var ret = C.C_ZNK13QGraphicsItem10mapToSceneEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  case 5:
    // invoke: _ZNK13QGraphicsItem10mapToSceneEdd
    // invoke: QPointF mapToScene(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK13QGraphicsItem10mapToSceneEdd(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapToScene", args)
  }

}

// setHandlesChildEvents(_Bool)
func (this *QGraphicsItem) setHandlesChildEvents(args ...interface{}) () {
  // setHandlesChildEvents(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem21setHandlesChildEventsEb
    // invoke: void setHandlesChildEvents(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem21setHandlesChildEventsEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setHandlesChildEvents", args)
  }

}

// transformOriginPoint()
func (this *QGraphicsItem) transformOriginPoint(args ...interface{}) () {
  // transformOriginPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem20transformOriginPointEv
    // invoke: QPointF transformOriginPoint()
    var ret = C.C_ZNK13QGraphicsItem20transformOriginPointEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "transformOriginPoint", args)
  }

}

// pos()
func (this *QGraphicsItem) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem3posEv
    // invoke: QPointF pos()
    var ret = C.C_ZNK13QGraphicsItem3posEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "pos", args)
  }

}

// mapRectToScene(const class QRectF &)
func (this *QGraphicsItem) mapRectToScene(args ...interface{}) () {
  // mapRectToScene(const class QRectF &)
  // mapRectToScene(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem14mapRectToSceneERK6QRectF
    // invoke: QRectF mapRectToScene(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem14mapRectToSceneERK6QRectF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK13QGraphicsItem14mapRectToSceneEdddd
    // invoke: QRectF mapRectToScene(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var ret = C.C_ZNK13QGraphicsItem14mapRectToSceneEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapRectToScene", args)
  }

}

// ungrabMouse()
func (this *QGraphicsItem) ungrabMouse(args ...interface{}) () {
  // ungrabMouse()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem11ungrabMouseEv
    // invoke: void ungrabMouse()
    C.C_ZN13QGraphicsItem11ungrabMouseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "ungrabMouse", args)
  }

}

// isObscured(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) isObscured(args ...interface{}) () {
  // isObscured(qreal, qreal, qreal, qreal)
  // isObscured(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem10isObscuredEdddd
    // invoke: bool isObscured(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var ret = C.C_ZNK13QGraphicsItem10isObscuredEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK13QGraphicsItem10isObscuredERK6QRectF
    // invoke: bool isObscured(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem10isObscuredERK6QRectF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isObscured", args)
  }

}

// setSelected(_Bool)
func (this *QGraphicsItem) setSelected(args ...interface{}) () {
  // setSelected(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem11setSelectedEb
    // invoke: void setSelected(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem11setSelectedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setSelected", args)
  }

}

// mapFromScene(const class QPolygonF &)
func (this *QGraphicsItem) mapFromScene(args ...interface{}) () {
  // mapFromScene(const class QPolygonF &)
  // mapFromScene(const class QRectF &)
  // mapFromScene(qreal, qreal, qreal, qreal)
  // mapFromScene(qreal, qreal)
  // mapFromScene(const class QPainterPath &)
  // mapFromScene(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem12mapFromSceneERK9QPolygonF
    // invoke: QPolygonF mapFromScene(const class QPolygonF &)
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem12mapFromSceneERK9QPolygonF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK13QGraphicsItem12mapFromSceneERK6QRectF
    // invoke: QPolygonF mapFromScene(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem12mapFromSceneERK6QRectF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 2:
    // invoke: _ZNK13QGraphicsItem12mapFromSceneEdddd
    // invoke: QPolygonF mapFromScene(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var ret = C.C_ZNK13QGraphicsItem12mapFromSceneEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  case 3:
    // invoke: _ZNK13QGraphicsItem12mapFromSceneEdd
    // invoke: QPointF mapFromScene(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK13QGraphicsItem12mapFromSceneEdd(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 4:
    // invoke: _ZNK13QGraphicsItem12mapFromSceneERK12QPainterPath
    // invoke: QPainterPath mapFromScene(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem12mapFromSceneERK12QPainterPath(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 5:
    // invoke: _ZNK13QGraphicsItem12mapFromSceneERK7QPointF
    // invoke: QPointF mapFromScene(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem12mapFromSceneERK7QPointF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapFromScene", args)
  }

}

// scene()
func (this *QGraphicsItem) scene(args ...interface{}) () {
  // scene()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem5sceneEv
    // invoke: QGraphicsScene * scene()
    C.C_ZNK13QGraphicsItem5sceneEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "scene", args)
  }

}

// isAncestorOf(const class QGraphicsItem *)
func (this *QGraphicsItem) isAncestorOf(args ...interface{}) () {
  // isAncestorOf(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem12isAncestorOfEPKS_
    // invoke: bool isAncestorOf(const class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem12isAncestorOfEPKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isAncestorOf", args)
  }

}

// topLevelItem()
func (this *QGraphicsItem) topLevelItem(args ...interface{}) () {
  // topLevelItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem12topLevelItemEv
    // invoke: QGraphicsItem * topLevelItem()
    C.C_ZNK13QGraphicsItem12topLevelItemEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "topLevelItem", args)
  }

}

// panel()
func (this *QGraphicsItem) panel(args ...interface{}) () {
  // panel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem5panelEv
    // invoke: QGraphicsItem * panel()
    C.C_ZNK13QGraphicsItem5panelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "panel", args)
  }

}

// unsetCursor()
func (this *QGraphicsItem) unsetCursor(args ...interface{}) () {
  // unsetCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem11unsetCursorEv
    // invoke: void unsetCursor()
    C.C_ZN13QGraphicsItem11unsetCursorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "unsetCursor", args)
  }

}

// transformations()
func (this *QGraphicsItem) transformations(args ...interface{}) () {
  // transformations()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem15transformationsEv
    // invoke: QList<QGraphicsTransform *> transformations()
    C.C_ZNK13QGraphicsItem15transformationsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "transformations", args)
  }

}

// mapRectFromScene(const class QRectF &)
func (this *QGraphicsItem) mapRectFromScene(args ...interface{}) () {
  // mapRectFromScene(const class QRectF &)
  // mapRectFromScene(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem16mapRectFromSceneERK6QRectF
    // invoke: QRectF mapRectFromScene(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem16mapRectFromSceneERK6QRectF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK13QGraphicsItem16mapRectFromSceneEdddd
    // invoke: QRectF mapRectFromScene(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var ret = C.C_ZNK13QGraphicsItem16mapRectFromSceneEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapRectFromScene", args)
  }

}

// contains(const class QPointF &)
func (this *QGraphicsItem) contains(args ...interface{}) () {
  // contains(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8containsERK7QPointF
    // invoke: bool contains(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem8containsERK7QPointF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "contains", args)
  }

}

// transform()
func (this *QGraphicsItem) transform(args ...interface{}) () {
  // transform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9transformEv
    // invoke: QTransform transform()
    var ret = C.C_ZNK13QGraphicsItem9transformEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "transform", args)
  }

}

// ungrabKeyboard()
func (this *QGraphicsItem) ungrabKeyboard(args ...interface{}) () {
  // ungrabKeyboard()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem14ungrabKeyboardEv
    // invoke: void ungrabKeyboard()
    C.C_ZN13QGraphicsItem14ungrabKeyboardEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "ungrabKeyboard", args)
  }

}

// installSceneEventFilter(class QGraphicsItem *)
func (this *QGraphicsItem) installSceneEventFilter(args ...interface{}) () {
  // installSceneEventFilter(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem23installSceneEventFilterEPS_
    // invoke: void installSceneEventFilter(class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem23installSceneEventFilterEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "installSceneEventFilter", args)
  }

}

// acceptHoverEvents()
func (this *QGraphicsItem) acceptHoverEvents(args ...interface{}) () {
  // acceptHoverEvents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem17acceptHoverEventsEv
    // invoke: bool acceptHoverEvents()
    var ret = C.C_ZNK13QGraphicsItem17acceptHoverEventsEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "acceptHoverEvents", args)
  }

}

// mapRectToItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) mapRectToItem(args ...interface{}) () {
  // mapRectToItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
  // mapRectToItem(const class QGraphicsItem *, const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[1][1] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem13mapRectToItemEPKS_dddd
    // invoke: QRectF mapRectToItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.double(args[4].(float64))
    if false {fmt.Println(arg4)}
    var ret = C.C_ZNK13QGraphicsItem13mapRectToItemEPKS_dddd(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK13QGraphicsItem13mapRectToItemEPKS_RK6QRectF
    // invoke: QRectF mapRectToItem(const class QGraphicsItem *, const class QRectF &)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QRectF).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK13QGraphicsItem13mapRectToItemEPKS_RK6QRectF(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapRectToItem", args)
  }

}

// mapFromItem(const class QGraphicsItem *, qreal, qreal)
func (this *QGraphicsItem) mapFromItem(args ...interface{}) () {
  // mapFromItem(const class QGraphicsItem *, qreal, qreal)
  // mapFromItem(const class QGraphicsItem *, const class QPointF &)
  // mapFromItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
  // mapFromItem(const class QGraphicsItem *, const class QRectF &)
  // mapFromItem(const class QGraphicsItem *, const class QPolygonF &)
  // mapFromItem(const class QGraphicsItem *, const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[1][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[3][1] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[4][1] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[5][1] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem11mapFromItemEPKS_dd
    // invoke: QPointF mapFromItem(const class QGraphicsItem *, qreal, qreal)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var ret = C.C_ZNK13QGraphicsItem11mapFromItemEPKS_dd(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK13QGraphicsItem11mapFromItemEPKS_RK7QPointF
    // invoke: QPointF mapFromItem(const class QGraphicsItem *, const class QPointF &)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK13QGraphicsItem11mapFromItemEPKS_RK7QPointF(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 2:
    // invoke: _ZNK13QGraphicsItem11mapFromItemEPKS_dddd
    // invoke: QPolygonF mapFromItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.double(args[4].(float64))
    if false {fmt.Println(arg4)}
    var ret = C.C_ZNK13QGraphicsItem11mapFromItemEPKS_dddd(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret)}
  case 3:
    // invoke: _ZNK13QGraphicsItem11mapFromItemEPKS_RK6QRectF
    // invoke: QPolygonF mapFromItem(const class QGraphicsItem *, const class QRectF &)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QRectF).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK13QGraphicsItem11mapFromItemEPKS_RK6QRectF(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 4:
    // invoke: _ZNK13QGraphicsItem11mapFromItemEPKS_RK9QPolygonF
    // invoke: QPolygonF mapFromItem(const class QGraphicsItem *, const class QPolygonF &)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPolygonF).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK13QGraphicsItem11mapFromItemEPKS_RK9QPolygonF(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 5:
    // invoke: _ZNK13QGraphicsItem11mapFromItemEPKS_RK12QPainterPath
    // invoke: QPainterPath mapFromItem(const class QGraphicsItem *, const class QPainterPath &)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPainterPath).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK13QGraphicsItem11mapFromItemEPKS_RK12QPainterPath(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapFromItem", args)
  }

}

// QGraphicsItem(class QGraphicsItem *)
func NewQGraphicsItem(args ...interface{}) *QGraphicsItem {
  // QGraphicsItem(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItemC1EPS_
    // invoke: void QGraphicsItem(class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QGraphicsItemC2EPS_(arg0)
    return &QGraphicsItem{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "QGraphicsItem", args)
  }

  return nil // QGraphicsItem{qclsinst:qthis}
}

// type()
func (this *QGraphicsItem) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem4typeEv
    // invoke: int type()
    var ret = C.C_ZNK13QGraphicsItem4typeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "type", args)
  }

}

// setEnabled(_Bool)
func (this *QGraphicsItem) setEnabled(args ...interface{}) () {
  // setEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem10setEnabledEb
    // invoke: void setEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem10setEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setEnabled", args)
  }

}

// grabMouse()
func (this *QGraphicsItem) grabMouse(args ...interface{}) () {
  // grabMouse()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem9grabMouseEv
    // invoke: void grabMouse()
    C.C_ZN13QGraphicsItem9grabMouseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "grabMouse", args)
  }

}

// setActive(_Bool)
func (this *QGraphicsItem) setActive(args ...interface{}) () {
  // setActive(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem9setActiveEb
    // invoke: void setActive(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem9setActiveEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setActive", args)
  }

}

// isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsItem) isObscuredBy(args ...interface{}) () {
  // isObscuredBy(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem12isObscuredByEPKS_
    // invoke: bool isObscuredBy(const class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem12isObscuredByEPKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isObscuredBy", args)
  }

}

// acceptDrops()
func (this *QGraphicsItem) acceptDrops(args ...interface{}) () {
  // acceptDrops()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem11acceptDropsEv
    // invoke: bool acceptDrops()
    var ret = C.C_ZNK13QGraphicsItem11acceptDropsEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "acceptDrops", args)
  }

}

// setParentItem(class QGraphicsItem *)
func (this *QGraphicsItem) setParentItem(args ...interface{}) () {
  // setParentItem(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem13setParentItemEPS_
    // invoke: void setParentItem(class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem13setParentItemEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setParentItem", args)
  }

}

// ensureVisible(qreal, qreal, qreal, qreal, int, int)
func (this *QGraphicsItem) ensureVisible(args ...interface{}) () {
  // ensureVisible(qreal, qreal, qreal, qreal, int, int)
  // ensureVisible(const class QRectF &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"
  vtys[0][5] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem13ensureVisibleEddddii
    // invoke: void ensureVisible(qreal, qreal, qreal, qreal, int, int)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    C.C_ZN13QGraphicsItem13ensureVisibleEddddii(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  case 1:
    // invoke: _ZN13QGraphicsItem13ensureVisibleERK6QRectFii
    // invoke: void ensureVisible(const class QRectF &, int, int)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN13QGraphicsItem13ensureVisibleERK6QRectFii(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "ensureVisible", args)
  }

}

// grabKeyboard()
func (this *QGraphicsItem) grabKeyboard(args ...interface{}) () {
  // grabKeyboard()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem12grabKeyboardEv
    // invoke: void grabKeyboard()
    C.C_ZN13QGraphicsItem12grabKeyboardEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "grabKeyboard", args)
  }

}

// parentObject()
func (this *QGraphicsItem) parentObject(args ...interface{}) () {
  // parentObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem12parentObjectEv
    // invoke: QGraphicsObject * parentObject()
    C.C_ZNK13QGraphicsItem12parentObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "parentObject", args)
  }

}

// isUnderMouse()
func (this *QGraphicsItem) isUnderMouse(args ...interface{}) () {
  // isUnderMouse()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem12isUnderMouseEv
    // invoke: bool isUnderMouse()
    var ret = C.C_ZNK13QGraphicsItem12isUnderMouseEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isUnderMouse", args)
  }

}

// parentWidget()
func (this *QGraphicsItem) parentWidget(args ...interface{}) () {
  // parentWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem12parentWidgetEv
    // invoke: QGraphicsWidget * parentWidget()
    C.C_ZNK13QGraphicsItem12parentWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "parentWidget", args)
  }

}

// setZValue(qreal)
func (this *QGraphicsItem) setZValue(args ...interface{}) () {
  // setZValue(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem9setZValueEd
    // invoke: void setZValue(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem9setZValueEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setZValue", args)
  }

}

// setMatrix(const class QMatrix &, _Bool)
func (this *QGraphicsItem) setMatrix(args ...interface{}) () {
  // setMatrix(const class QMatrix &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem9setMatrixERK7QMatrixb
    // invoke: void setMatrix(const class QMatrix &, _Bool)
    var arg0 = args[0].(QMatrix).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN13QGraphicsItem9setMatrixERK7QMatrixb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setMatrix", args)
  }

}

// cursor()
func (this *QGraphicsItem) cursor(args ...interface{}) () {
  // cursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem6cursorEv
    // invoke: QCursor cursor()
    var ret = C.C_ZNK13QGraphicsItem6cursorEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "cursor", args)
  }

}

// zValue()
func (this *QGraphicsItem) zValue(args ...interface{}) () {
  // zValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem6zValueEv
    // invoke: qreal zValue()
    var ret = C.C_ZNK13QGraphicsItem6zValueEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "zValue", args)
  }

}

// setVisible(_Bool)
func (this *QGraphicsItem) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem10setVisibleEb
    // invoke: void setVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem10setVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setVisible", args)
  }

}

// resetMatrix()
func (this *QGraphicsItem) resetMatrix(args ...interface{}) () {
  // resetMatrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem11resetMatrixEv
    // invoke: void resetMatrix()
    C.C_ZN13QGraphicsItem11resetMatrixEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "resetMatrix", args)
  }

}

// sceneTransform()
func (this *QGraphicsItem) sceneTransform(args ...interface{}) () {
  // sceneTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem14sceneTransformEv
    // invoke: QTransform sceneTransform()
    var ret = C.C_ZNK13QGraphicsItem14sceneTransformEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "sceneTransform", args)
  }

}

// isWindow()
func (this *QGraphicsItem) isWindow(args ...interface{}) () {
  // isWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8isWindowEv
    // invoke: bool isWindow()
    var ret = C.C_ZNK13QGraphicsItem8isWindowEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isWindow", args)
  }

}

// mapRectFromParent(const class QRectF &)
func (this *QGraphicsItem) mapRectFromParent(args ...interface{}) () {
  // mapRectFromParent(const class QRectF &)
  // mapRectFromParent(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem17mapRectFromParentERK6QRectF
    // invoke: QRectF mapRectFromParent(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem17mapRectFromParentERK6QRectF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK13QGraphicsItem17mapRectFromParentEdddd
    // invoke: QRectF mapRectFromParent(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var ret = C.C_ZNK13QGraphicsItem17mapRectFromParentEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapRectFromParent", args)
  }

}

// setScale(qreal)
func (this *QGraphicsItem) setScale(args ...interface{}) () {
  // setScale(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem8setScaleEd
    // invoke: void setScale(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem8setScaleEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setScale", args)
  }

}

// isWidget()
func (this *QGraphicsItem) isWidget(args ...interface{}) () {
  // isWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8isWidgetEv
    // invoke: bool isWidget()
    var ret = C.C_ZNK13QGraphicsItem8isWidgetEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isWidget", args)
  }

}

// resetTransform()
func (this *QGraphicsItem) resetTransform(args ...interface{}) () {
  // resetTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem14resetTransformEv
    // invoke: void resetTransform()
    C.C_ZN13QGraphicsItem14resetTransformEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "resetTransform", args)
  }

}

// sceneBoundingRect()
func (this *QGraphicsItem) sceneBoundingRect(args ...interface{}) () {
  // sceneBoundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem17sceneBoundingRectEv
    // invoke: QRectF sceneBoundingRect()
    var ret = C.C_ZNK13QGraphicsItem17sceneBoundingRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "sceneBoundingRect", args)
  }

}

// setOpacity(qreal)
func (this *QGraphicsItem) setOpacity(args ...interface{}) () {
  // setOpacity(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem10setOpacityEd
    // invoke: void setOpacity(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem10setOpacityEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setOpacity", args)
  }

}

// focusItem()
func (this *QGraphicsItem) focusItem(args ...interface{}) () {
  // focusItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9focusItemEv
    // invoke: QGraphicsItem * focusItem()
    C.C_ZNK13QGraphicsItem9focusItemEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "focusItem", args)
  }

}

// hasCursor()
func (this *QGraphicsItem) hasCursor(args ...interface{}) () {
  // hasCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9hasCursorEv
    // invoke: bool hasCursor()
    var ret = C.C_ZNK13QGraphicsItem9hasCursorEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "hasCursor", args)
  }

}

// setAcceptTouchEvents(_Bool)
func (this *QGraphicsItem) setAcceptTouchEvents(args ...interface{}) () {
  // setAcceptTouchEvents(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem20setAcceptTouchEventsEb
    // invoke: void setAcceptTouchEvents(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem20setAcceptTouchEventsEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setAcceptTouchEvents", args)
  }

}

// setData(int, const class QVariant &)
func (this *QGraphicsItem) setData(args ...interface{}) () {
  // setData(int, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem7setDataEiRK8QVariant
    // invoke: void setData(int, const class QVariant &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QGraphicsItem7setDataEiRK8QVariant(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setData", args)
  }

}

// opacity()
func (this *QGraphicsItem) opacity(args ...interface{}) () {
  // opacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem7opacityEv
    // invoke: qreal opacity()
    var ret = C.C_ZNK13QGraphicsItem7opacityEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "opacity", args)
  }

}

// isVisible()
func (this *QGraphicsItem) isVisible(args ...interface{}) () {
  // isVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9isVisibleEv
    // invoke: bool isVisible()
    var ret = C.C_ZNK13QGraphicsItem9isVisibleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isVisible", args)
  }

}

// update(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) update(args ...interface{}) () {
  // update(qreal, qreal, qreal, qreal)
  // update(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem6updateEdddd
    // invoke: void update(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN13QGraphicsItem6updateEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN13QGraphicsItem6updateERK6QRectF
    // invoke: void update(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem6updateERK6QRectF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "update", args)
  }

}

// mapRectFromItem(const class QGraphicsItem *, const class QRectF &)
func (this *QGraphicsItem) mapRectFromItem(args ...interface{}) () {
  // mapRectFromItem(const class QGraphicsItem *, const class QRectF &)
  // mapRectFromItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[0][1] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][4] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem15mapRectFromItemEPKS_RK6QRectF
    // invoke: QRectF mapRectFromItem(const class QGraphicsItem *, const class QRectF &)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QRectF).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK13QGraphicsItem15mapRectFromItemEPKS_RK6QRectF(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK13QGraphicsItem15mapRectFromItemEPKS_dddd
    // invoke: QRectF mapRectFromItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.double(args[4].(float64))
    if false {fmt.Println(arg4)}
    var ret = C.C_ZNK13QGraphicsItem15mapRectFromItemEPKS_dddd(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapRectFromItem", args)
  }

}

// setX(qreal)
func (this *QGraphicsItem) setX(args ...interface{}) () {
  // setX(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem4setXEd
    // invoke: void setX(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem4setXEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setX", args)
  }

}

// setY(qreal)
func (this *QGraphicsItem) setY(args ...interface{}) () {
  // setY(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem4setYEd
    // invoke: void setY(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem4setYEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setY", args)
  }

}

// setCursor(const class QCursor &)
func (this *QGraphicsItem) setCursor(args ...interface{}) () {
  // setCursor(const class QCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCursor{}) // "const QCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem9setCursorERK7QCursor
    // invoke: void setCursor(const class QCursor &)
    var arg0 = args[0].(QCursor).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem9setCursorERK7QCursor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setCursor", args)
  }

}

// boundingRegionGranularity()
func (this *QGraphicsItem) boundingRegionGranularity(args ...interface{}) () {
  // boundingRegionGranularity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem25boundingRegionGranularityEv
    // invoke: qreal boundingRegionGranularity()
    var ret = C.C_ZNK13QGraphicsItem25boundingRegionGranularityEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "boundingRegionGranularity", args)
  }

}

// removeSceneEventFilter(class QGraphicsItem *)
func (this *QGraphicsItem) removeSceneEventFilter(args ...interface{}) () {
  // removeSceneEventFilter(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem22removeSceneEventFilterEPS_
    // invoke: void removeSceneEventFilter(class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem22removeSceneEventFilterEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "removeSceneEventFilter", args)
  }

}

// acceptTouchEvents()
func (this *QGraphicsItem) acceptTouchEvents(args ...interface{}) () {
  // acceptTouchEvents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem17acceptTouchEventsEv
    // invoke: bool acceptTouchEvents()
    var ret = C.C_ZNK13QGraphicsItem17acceptTouchEventsEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "acceptTouchEvents", args)
  }

}

// sceneMatrix()
func (this *QGraphicsItem) sceneMatrix(args ...interface{}) () {
  // sceneMatrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem11sceneMatrixEv
    // invoke: QMatrix sceneMatrix()
    var ret = C.C_ZNK13QGraphicsItem11sceneMatrixEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "sceneMatrix", args)
  }

}

// cacheMode()
func (this *QGraphicsItem) cacheMode(args ...interface{}) () {
  // cacheMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9cacheModeEv
    // invoke: QGraphicsItem::CacheMode cacheMode()
    C.C_ZNK13QGraphicsItem9cacheModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "cacheMode", args)
  }

}

// flags()
func (this *QGraphicsItem) flags(args ...interface{}) () {
  // flags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem5flagsEv
    // invoke: GraphicsItemFlags flags()
    C.C_ZNK13QGraphicsItem5flagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "flags", args)
  }

}

// clipPath()
func (this *QGraphicsItem) clipPath(args ...interface{}) () {
  // clipPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8clipPathEv
    // invoke: QPainterPath clipPath()
    var ret = C.C_ZNK13QGraphicsItem8clipPathEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "clipPath", args)
  }

}

// y()
func (this *QGraphicsItem) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem1yEv
    // invoke: qreal y()
    C.C_ZNK13QGraphicsItem1yEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "y", args)
  }

}

// scroll(qreal, qreal, const class QRectF &)
func (this *QGraphicsItem) scroll(args ...interface{}) () {
  // scroll(qreal, qreal, const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem6scrollEddRK6QRectF
    // invoke: void scroll(qreal, qreal, const class QRectF &)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QRectF).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN13QGraphicsItem6scrollEddRK6QRectF(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "scroll", args)
  }

}

// boundingRegion(const class QTransform &)
func (this *QGraphicsItem) boundingRegion(args ...interface{}) () {
  // boundingRegion(const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem14boundingRegionERK10QTransform
    // invoke: QRegion boundingRegion(const class QTransform &)
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QGraphicsItem14boundingRegionERK10QTransform(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "boundingRegion", args)
  }

}

// ~QGraphicsObject()
func (this *QGraphicsObject) FreeQGraphicsObject(args ...interface{}) () {
  // ~QGraphicsObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsObjectD0Ev
    // invoke: void ~QGraphicsObject()
    C.C_ZN15QGraphicsObjectD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsObject", "~QGraphicsObject", args)
  }

}

// QGraphicsObject(class QGraphicsItem *)
func NewQGraphicsObject(args ...interface{}) *QGraphicsObject {
  // QGraphicsObject(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsObjectC1EP13QGraphicsItem
    // invoke: void QGraphicsObject(class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QGraphicsObjectC2EP13QGraphicsItem(arg0)
    return &QGraphicsObject{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsObject", "QGraphicsObject", args)
  }

  return nil // QGraphicsObject{qclsinst:qthis}
}

// metaObject()
func (this *QGraphicsObject) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsObject10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK15QGraphicsObject10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsObject", "metaObject", args)
  }

}

// opaqueArea()
func (this *QGraphicsSimpleTextItem) opaqueArea(args ...interface{}) () {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSimpleTextItem10opaqueAreaEv
    // invoke: QPainterPath opaqueArea()
    var ret = C.C_ZNK23QGraphicsSimpleTextItem10opaqueAreaEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "opaqueArea", args)
  }

}

// ~QGraphicsSimpleTextItem()
func (this *QGraphicsSimpleTextItem) FreeQGraphicsSimpleTextItem(args ...interface{}) () {
  // ~QGraphicsSimpleTextItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsSimpleTextItemD0Ev
    // invoke: void ~QGraphicsSimpleTextItem()
    C.C_ZN23QGraphicsSimpleTextItemD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "~QGraphicsSimpleTextItem", args)
  }

}

// boundingRect()
func (this *QGraphicsSimpleTextItem) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSimpleTextItem12boundingRectEv
    // invoke: QRectF boundingRect()
    var ret = C.C_ZNK23QGraphicsSimpleTextItem12boundingRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "boundingRect", args)
  }

}

// text()
func (this *QGraphicsSimpleTextItem) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSimpleTextItem4textEv
    // invoke: QString text()
    var ret = C.C_ZNK23QGraphicsSimpleTextItem4textEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "text", args)
  }

}

// setText(const class QString &)
func (this *QGraphicsSimpleTextItem) setText(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsSimpleTextItem7setTextERK7QString
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN23QGraphicsSimpleTextItem7setTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "setText", args)
  }

}

// contains(const class QPointF &)
func (this *QGraphicsSimpleTextItem) contains(args ...interface{}) () {
  // contains(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSimpleTextItem8containsERK7QPointF
    // invoke: bool contains(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK23QGraphicsSimpleTextItem8containsERK7QPointF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "contains", args)
  }

}

// isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsSimpleTextItem) isObscuredBy(args ...interface{}) () {
  // isObscuredBy(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSimpleTextItem12isObscuredByEPK13QGraphicsItem
    // invoke: bool isObscuredBy(const class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK23QGraphicsSimpleTextItem12isObscuredByEPK13QGraphicsItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "isObscuredBy", args)
  }

}

// paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsSimpleTextItem) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsSimpleTextItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
    // invoke: void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(QPainter).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStyleOptionGraphicsItem).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN23QGraphicsSimpleTextItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "paint", args)
  }

}

// shape()
func (this *QGraphicsSimpleTextItem) shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSimpleTextItem5shapeEv
    // invoke: QPainterPath shape()
    var ret = C.C_ZNK23QGraphicsSimpleTextItem5shapeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "shape", args)
  }

}

// QGraphicsSimpleTextItem(class QGraphicsItem *)
func NewQGraphicsSimpleTextItem(args ...interface{}) *QGraphicsSimpleTextItem {
  // QGraphicsSimpleTextItem(class QGraphicsItem *)
  // QGraphicsSimpleTextItem(const class QString &, class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsSimpleTextItemC1EP13QGraphicsItem
    // invoke: void QGraphicsSimpleTextItem(class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN23QGraphicsSimpleTextItemC2EP13QGraphicsItem(arg0)
    return &QGraphicsSimpleTextItem{qclsinst:qthis}
  case 1:
    // invoke: _ZN23QGraphicsSimpleTextItemC1ERK7QStringP13QGraphicsItem
    // invoke: void QGraphicsSimpleTextItem(const class QString &, class QGraphicsItem *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN23QGraphicsSimpleTextItemC2ERK7QStringP13QGraphicsItem(arg0, arg1)
    return &QGraphicsSimpleTextItem{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "QGraphicsSimpleTextItem", args)
  }

  return nil // QGraphicsSimpleTextItem{qclsinst:qthis}
}

// setFont(const class QFont &)
func (this *QGraphicsSimpleTextItem) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsSimpleTextItem7setFontERK5QFont
    // invoke: void setFont(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN23QGraphicsSimpleTextItem7setFontERK5QFont(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "setFont", args)
  }

}

// font()
func (this *QGraphicsSimpleTextItem) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSimpleTextItem4fontEv
    // invoke: QFont font()
    var ret = C.C_ZNK23QGraphicsSimpleTextItem4fontEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "font", args)
  }

}

// type()
func (this *QGraphicsSimpleTextItem) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSimpleTextItem4typeEv
    // invoke: int type()
    var ret = C.C_ZNK23QGraphicsSimpleTextItem4typeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "type", args)
  }

}

// <= body block end

