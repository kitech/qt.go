package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtGui/qtextformat.h
// dst-file: /src/gui/qtextformat.go
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
  // proto:  qreal QTextLength::value(qreal maximumLength);
extern void demth_ZNK11QTextLength5valueEd(void* qthis, double arg0);
  // proto:  void QTextLength::QTextLength();
extern void* dector_ZN11QTextLengthC1Ev();
extern void demth_ZN11QTextLengthC1Ev(void* qthis);
  // proto:  qreal QTextLength::rawValue();
extern void demth_ZNK11QTextLength8rawValueEv(void* qthis);
  // proto:  void QTextImageFormat::QTextImageFormat();
extern void* dector_ZN16QTextImageFormatC1Ev();
extern void _ZN16QTextImageFormatC1Ev(void* qthis);
  // proto:  bool QTextImageFormat::isValid();
extern void demth_ZNK16QTextImageFormat7isValidEv(void* qthis);
  // proto:  qreal QTextImageFormat::width();
extern void demth_ZNK16QTextImageFormat5widthEv(void* qthis);
  // proto:  void QTextImageFormat::QTextImageFormat(const QTextFormat & format);
extern void* dector_ZN16QTextImageFormatC1ERK11QTextFormat(void* arg0);
extern void _ZN16QTextImageFormatC1ERK11QTextFormat(void* qthis, void* arg0);
  // proto:  void QTextImageFormat::setHeight(qreal height);
extern void demth_ZN16QTextImageFormat9setHeightEd(void* qthis, double arg0);
  // proto:  void QTextImageFormat::setWidth(qreal width);
extern void demth_ZN16QTextImageFormat8setWidthEd(void* qthis, double arg0);
  // proto:  void QTextImageFormat::setName(const QString & name);
extern void demth_ZN16QTextImageFormat7setNameERK7QString(void* qthis, void* arg0);
  // proto:  QString QTextImageFormat::name();
extern void demth_ZNK16QTextImageFormat4nameEv(void* qthis);
  // proto:  qreal QTextImageFormat::height();
extern void demth_ZNK16QTextImageFormat6heightEv(void* qthis);
  // proto:  QTextBlockFormat QTextFormat::toBlockFormat();
extern void _ZNK11QTextFormat13toBlockFormatEv(void* qthis);
  // proto:  QString QTextFormat::stringProperty(int propertyId);
extern void _ZNK11QTextFormat14stringPropertyEi(void* qthis, int arg0);
  // proto:  QVector<QTextLength> QTextFormat::lengthVectorProperty(int propertyId);
extern void _ZNK11QTextFormat20lengthVectorPropertyEi(void* qthis, int arg0);
  // proto:  int QTextFormat::objectIndex();
extern void _ZNK11QTextFormat11objectIndexEv(void* qthis);
  // proto:  void QTextFormat::setObjectIndex(int object);
extern void _ZN11QTextFormat14setObjectIndexEi(void* qthis, int arg0);
  // proto:  void QTextFormat::clearForeground();
extern void demth_ZN11QTextFormat15clearForegroundEv(void* qthis);
  // proto:  bool QTextFormat::isTableCellFormat();
extern void demth_ZNK11QTextFormat17isTableCellFormatEv(void* qthis);
  // proto:  void QTextFormat::~QTextFormat();
extern void _ZN11QTextFormatD0Ev(void* qthis);
  // proto:  bool QTextFormat::isValid();
extern void demth_ZNK11QTextFormat7isValidEv(void* qthis);
  // proto:  void QTextFormat::QTextFormat(const QTextFormat & rhs);
extern void* dector_ZN11QTextFormatC1ERKS_(void* arg0);
extern void _ZN11QTextFormatC1ERKS_(void* qthis, void* arg0);
  // proto:  QTextLength QTextFormat::lengthProperty(int propertyId);
extern void _ZNK11QTextFormat14lengthPropertyEi(void* qthis, int arg0);
  // proto:  void QTextFormat::merge(const QTextFormat & other);
extern void _ZN11QTextFormat5mergeERKS_(void* qthis, void* arg0);
  // proto:  QColor QTextFormat::colorProperty(int propertyId);
extern void _ZNK11QTextFormat13colorPropertyEi(void* qthis, int arg0);
  // proto:  void QTextFormat::QTextFormat();
extern void* dector_ZN11QTextFormatC1Ev();
extern void _ZN11QTextFormatC1Ev(void* qthis);
  // proto:  void QTextFormat::setForeground(const QBrush & brush);
extern void demth_ZN11QTextFormat13setForegroundERK6QBrush(void* qthis, void* arg0);
  // proto:  bool QTextFormat::boolProperty(int propertyId);
extern void _ZNK11QTextFormat12boolPropertyEi(void* qthis, int arg0);
  // proto:  bool QTextFormat::isListFormat();
extern void demth_ZNK11QTextFormat12isListFormatEv(void* qthis);
  // proto:  void QTextFormat::QTextFormat(int type);
extern void* dector_ZN11QTextFormatC1Ei(int arg0);
extern void _ZN11QTextFormatC1Ei(void* qthis, int arg0);
  // proto:  bool QTextFormat::isImageFormat();
extern void demth_ZNK11QTextFormat13isImageFormatEv(void* qthis);
  // proto:  void QTextFormat::clearProperty(int propertyId);
extern void _ZN11QTextFormat13clearPropertyEi(void* qthis, int arg0);
  // proto:  QTextFrameFormat QTextFormat::toFrameFormat();
extern void _ZNK11QTextFormat13toFrameFormatEv(void* qthis);
  // proto:  QBrush QTextFormat::brushProperty(int propertyId);
extern void _ZNK11QTextFormat13brushPropertyEi(void* qthis, int arg0);
  // proto:  int QTextFormat::propertyCount();
extern void _ZNK11QTextFormat13propertyCountEv(void* qthis);
  // proto:  QPen QTextFormat::penProperty(int propertyId);
extern void _ZNK11QTextFormat11penPropertyEi(void* qthis, int arg0);
  // proto:  QVariant QTextFormat::property(int propertyId);
extern void _ZNK11QTextFormat8propertyEi(void* qthis, int arg0);
  // proto:  bool QTextFormat::isTableFormat();
extern void demth_ZNK11QTextFormat13isTableFormatEv(void* qthis);
  // proto:  void QTextFormat::setProperty(int propertyId, const QVariant & value);
extern void _ZN11QTextFormat11setPropertyEiRK8QVariant(void* qthis, int arg0, void* arg1);
  // proto:  int QTextFormat::type();
extern void _ZNK11QTextFormat4typeEv(void* qthis);
  // proto:  bool QTextFormat::isCharFormat();
extern void demth_ZNK11QTextFormat12isCharFormatEv(void* qthis);
  // proto:  void QTextFormat::clearBackground();
extern void demth_ZN11QTextFormat15clearBackgroundEv(void* qthis);
  // proto:  bool QTextFormat::isBlockFormat();
extern void demth_ZNK11QTextFormat13isBlockFormatEv(void* qthis);
  // proto:  QBrush QTextFormat::background();
extern void demth_ZNK11QTextFormat10backgroundEv(void* qthis);
  // proto:  qreal QTextFormat::doubleProperty(int propertyId);
extern void _ZNK11QTextFormat14doublePropertyEi(void* qthis, int arg0);
  // proto:  void QTextFormat::swap(QTextFormat & other);
extern void _ZN11QTextFormat4swapERS_(void* qthis, void* arg0);
  // proto:  QTextImageFormat QTextFormat::toImageFormat();
extern void _ZNK11QTextFormat13toImageFormatEv(void* qthis);
  // proto:  bool QTextFormat::hasProperty(int propertyId);
extern void _ZNK11QTextFormat11hasPropertyEi(void* qthis, int arg0);
  // proto:  QBrush QTextFormat::foreground();
extern void demth_ZNK11QTextFormat10foregroundEv(void* qthis);
  // proto:  void QTextFormat::setObjectType(int type);
extern void demth_ZN11QTextFormat13setObjectTypeEi(void* qthis, int arg0);
  // proto:  void QTextFormat::setBackground(const QBrush & brush);
extern void demth_ZN11QTextFormat13setBackgroundERK6QBrush(void* qthis, void* arg0);
  // proto:  QTextTableFormat QTextFormat::toTableFormat();
extern void _ZNK11QTextFormat13toTableFormatEv(void* qthis);
  // proto:  bool QTextFormat::isFrameFormat();
extern void demth_ZNK11QTextFormat13isFrameFormatEv(void* qthis);
  // proto:  int QTextFormat::intProperty(int propertyId);
extern void _ZNK11QTextFormat11intPropertyEi(void* qthis, int arg0);
  // proto:  QTextCharFormat QTextFormat::toCharFormat();
extern void _ZNK11QTextFormat12toCharFormatEv(void* qthis);
  // proto:  bool QTextFormat::isEmpty();
extern void demth_ZNK11QTextFormat7isEmptyEv(void* qthis);
  // proto:  QTextTableCellFormat QTextFormat::toTableCellFormat();
extern void _ZNK11QTextFormat17toTableCellFormatEv(void* qthis);
  // proto:  int QTextFormat::objectType();
extern void demth_ZNK11QTextFormat10objectTypeEv(void* qthis);
  // proto:  QTextListFormat QTextFormat::toListFormat();
extern void _ZNK11QTextFormat12toListFormatEv(void* qthis);
  // proto:  QMap<int, QVariant> QTextFormat::properties();
extern void _ZNK11QTextFormat10propertiesEv(void* qthis);
  // proto:  int QTextBlockFormat::indent();
extern void demth_ZNK16QTextBlockFormat6indentEv(void* qthis);
  // proto:  void QTextBlockFormat::setTextIndent(qreal aindent);
extern void demth_ZN16QTextBlockFormat13setTextIndentEd(void* qthis, double arg0);
  // proto:  void QTextBlockFormat::setNonBreakableLines(bool b);
extern void demth_ZN16QTextBlockFormat20setNonBreakableLinesEb(void* qthis, bool arg0);
  // proto:  void QTextBlockFormat::setIndent(int indent);
extern void demth_ZN16QTextBlockFormat9setIndentEi(void* qthis, int arg0);
  // proto:  qreal QTextBlockFormat::textIndent();
extern void demth_ZNK16QTextBlockFormat10textIndentEv(void* qthis);
  // proto:  qreal QTextBlockFormat::lineHeight();
extern void demth_ZNK16QTextBlockFormat10lineHeightEv(void* qthis);
  // proto:  void QTextBlockFormat::QTextBlockFormat(const QTextFormat & fmt);
extern void* dector_ZN16QTextBlockFormatC1ERK11QTextFormat(void* arg0);
extern void _ZN16QTextBlockFormatC1ERK11QTextFormat(void* qthis, void* arg0);
  // proto:  qreal QTextBlockFormat::lineHeight(qreal scriptLineHeight, qreal scaling);
extern void demth_ZNK16QTextBlockFormat10lineHeightEdd(void* qthis, double arg0, double arg1);
  // proto:  void QTextBlockFormat::setRightMargin(qreal margin);
extern void demth_ZN16QTextBlockFormat14setRightMarginEd(void* qthis, double arg0);
  // proto:  qreal QTextBlockFormat::topMargin();
extern void demth_ZNK16QTextBlockFormat9topMarginEv(void* qthis);
  // proto:  void QTextBlockFormat::QTextBlockFormat();
extern void* dector_ZN16QTextBlockFormatC1Ev();
extern void _ZN16QTextBlockFormatC1Ev(void* qthis);
  // proto:  qreal QTextBlockFormat::rightMargin();
extern void demth_ZNK16QTextBlockFormat11rightMarginEv(void* qthis);
  // proto:  qreal QTextBlockFormat::bottomMargin();
extern void demth_ZNK16QTextBlockFormat12bottomMarginEv(void* qthis);
  // proto:  void QTextBlockFormat::setTopMargin(qreal margin);
extern void demth_ZN16QTextBlockFormat12setTopMarginEd(void* qthis, double arg0);
  // proto:  qreal QTextBlockFormat::leftMargin();
extern void demth_ZNK16QTextBlockFormat10leftMarginEv(void* qthis);
  // proto:  void QTextBlockFormat::setLineHeight(qreal height, int heightType);
extern void demth_ZN16QTextBlockFormat13setLineHeightEdi(void* qthis, double arg0, int arg1);
  // proto:  void QTextBlockFormat::setBottomMargin(qreal margin);
extern void demth_ZN16QTextBlockFormat15setBottomMarginEd(void* qthis, double arg0);
  // proto:  int QTextBlockFormat::lineHeightType();
extern void demth_ZNK16QTextBlockFormat14lineHeightTypeEv(void* qthis);
  // proto:  void QTextBlockFormat::setLeftMargin(qreal margin);
extern void demth_ZN16QTextBlockFormat13setLeftMarginEd(void* qthis, double arg0);
  // proto:  bool QTextBlockFormat::isValid();
extern void demth_ZNK16QTextBlockFormat7isValidEv(void* qthis);
  // proto:  bool QTextBlockFormat::nonBreakableLines();
extern void demth_ZNK16QTextBlockFormat17nonBreakableLinesEv(void* qthis);
  // proto:  void QTextCharFormat::setFontLetterSpacing(qreal spacing);
extern void demth_ZN15QTextCharFormat20setFontLetterSpacingEd(void* qthis, double arg0);
  // proto:  bool QTextCharFormat::isAnchor();
extern void demth_ZNK15QTextCharFormat8isAnchorEv(void* qthis);
  // proto:  void QTextCharFormat::setFont(const QFont & font);
extern void _ZN15QTextCharFormat7setFontERK5QFont(void* qthis, void* arg0);
  // proto:  bool QTextCharFormat::fontOverline();
extern void demth_ZNK15QTextCharFormat12fontOverlineEv(void* qthis);
  // proto:  QFont QTextCharFormat::font();
extern void _ZNK15QTextCharFormat4fontEv(void* qthis);
  // proto:  QString QTextCharFormat::fontFamily();
extern void demth_ZNK15QTextCharFormat10fontFamilyEv(void* qthis);
  // proto:  bool QTextCharFormat::fontStrikeOut();
extern void demth_ZNK15QTextCharFormat13fontStrikeOutEv(void* qthis);
  // proto:  void QTextCharFormat::setFontPointSize(qreal size);
extern void demth_ZN15QTextCharFormat16setFontPointSizeEd(void* qthis, double arg0);
  // proto:  void QTextCharFormat::setUnderlineColor(const QColor & color);
extern void demth_ZN15QTextCharFormat17setUnderlineColorERK6QColor(void* qthis, void* arg0);
  // proto:  int QTextCharFormat::tableCellRowSpan();
extern void demth_ZNK15QTextCharFormat16tableCellRowSpanEv(void* qthis);
  // proto:  void QTextCharFormat::setFontUnderline(bool underline);
extern void demth_ZN15QTextCharFormat16setFontUnderlineEb(void* qthis, bool arg0);
  // proto:  bool QTextCharFormat::isValid();
extern void _ZNK15QTextCharFormat7isValidEv(void* qthis);
  // proto:  bool QTextCharFormat::fontItalic();
extern void demth_ZNK15QTextCharFormat10fontItalicEv(void* qthis);
  // proto:  void QTextCharFormat::setToolTip(const QString & tip);
extern void demth_ZN15QTextCharFormat10setToolTipERK7QString(void* qthis, void* arg0);
  // proto:  void QTextCharFormat::setTextOutline(const QPen & pen);
extern void demth_ZN15QTextCharFormat14setTextOutlineERK4QPen(void* qthis, void* arg0);
  // proto:  void QTextCharFormat::setTableCellRowSpan(int tableCellRowSpan);
extern void demth_ZN15QTextCharFormat19setTableCellRowSpanEi(void* qthis, int arg0);
  // proto:  void QTextCharFormat::setAnchor(bool anchor);
extern void demth_ZN15QTextCharFormat9setAnchorEb(void* qthis, bool arg0);
  // proto:  qreal QTextCharFormat::fontPointSize();
extern void demth_ZNK15QTextCharFormat13fontPointSizeEv(void* qthis);
  // proto:  void QTextCharFormat::QTextCharFormat(const QTextFormat & fmt);
extern void* dector_ZN15QTextCharFormatC1ERK11QTextFormat(void* arg0);
extern void _ZN15QTextCharFormatC1ERK11QTextFormat(void* qthis, void* arg0);
  // proto:  void QTextCharFormat::setFontStrikeOut(bool strikeOut);
extern void demth_ZN15QTextCharFormat16setFontStrikeOutEb(void* qthis, bool arg0);
  // proto:  qreal QTextCharFormat::fontWordSpacing();
extern void demth_ZNK15QTextCharFormat15fontWordSpacingEv(void* qthis);
  // proto:  QString QTextCharFormat::toolTip();
extern void demth_ZNK15QTextCharFormat7toolTipEv(void* qthis);
  // proto:  void QTextCharFormat::setAnchorNames(const QStringList & names);
extern void demth_ZN15QTextCharFormat14setAnchorNamesERK11QStringList(void* qthis, void* arg0);
  // proto:  QStringList QTextCharFormat::anchorNames();
extern void _ZNK15QTextCharFormat11anchorNamesEv(void* qthis);
  // proto:  void QTextCharFormat::setFontFixedPitch(bool fixedPitch);
extern void demth_ZN15QTextCharFormat17setFontFixedPitchEb(void* qthis, bool arg0);
  // proto:  void QTextCharFormat::setFontItalic(bool italic);
extern void demth_ZN15QTextCharFormat13setFontItalicEb(void* qthis, bool arg0);
  // proto:  void QTextCharFormat::setFontFamily(const QString & family);
extern void demth_ZN15QTextCharFormat13setFontFamilyERK7QString(void* qthis, void* arg0);
  // proto:  bool QTextCharFormat::fontFixedPitch();
extern void demth_ZNK15QTextCharFormat14fontFixedPitchEv(void* qthis);
  // proto:  void QTextCharFormat::setAnchorHref(const QString & value);
extern void demth_ZN15QTextCharFormat13setAnchorHrefERK7QString(void* qthis, void* arg0);
  // proto:  int QTextCharFormat::fontStretch();
extern void demth_ZNK15QTextCharFormat11fontStretchEv(void* qthis);
  // proto:  void QTextCharFormat::setFontKerning(bool enable);
extern void demth_ZN15QTextCharFormat14setFontKerningEb(void* qthis, bool arg0);
  // proto:  int QTextCharFormat::tableCellColumnSpan();
extern void demth_ZNK15QTextCharFormat19tableCellColumnSpanEv(void* qthis);
  // proto:  void QTextCharFormat::QTextCharFormat();
extern void* dector_ZN15QTextCharFormatC1Ev();
extern void _ZN15QTextCharFormatC1Ev(void* qthis);
  // proto:  qreal QTextCharFormat::fontLetterSpacing();
extern void demth_ZNK15QTextCharFormat17fontLetterSpacingEv(void* qthis);
  // proto:  QString QTextCharFormat::anchorHref();
extern void demth_ZNK15QTextCharFormat10anchorHrefEv(void* qthis);
  // proto:  QString QTextCharFormat::anchorName();
extern void _ZNK15QTextCharFormat10anchorNameEv(void* qthis);
  // proto:  void QTextCharFormat::setFontStretch(int factor);
extern void demth_ZN15QTextCharFormat14setFontStretchEi(void* qthis, int arg0);
  // proto:  void QTextCharFormat::setAnchorName(const QString & name);
extern void demth_ZN15QTextCharFormat13setAnchorNameERK7QString(void* qthis, void* arg0);
  // proto:  bool QTextCharFormat::fontKerning();
extern void demth_ZNK15QTextCharFormat11fontKerningEv(void* qthis);
  // proto:  void QTextCharFormat::setFontWeight(int weight);
extern void demth_ZN15QTextCharFormat13setFontWeightEi(void* qthis, int arg0);
  // proto:  bool QTextCharFormat::fontUnderline();
extern void _ZNK15QTextCharFormat13fontUnderlineEv(void* qthis);
  // proto:  void QTextCharFormat::setFontWordSpacing(qreal spacing);
extern void demth_ZN15QTextCharFormat18setFontWordSpacingEd(void* qthis, double arg0);
  // proto:  QColor QTextCharFormat::underlineColor();
extern void demth_ZNK15QTextCharFormat14underlineColorEv(void* qthis);
  // proto:  int QTextCharFormat::fontWeight();
extern void demth_ZNK15QTextCharFormat10fontWeightEv(void* qthis);
  // proto:  void QTextCharFormat::setFontOverline(bool overline);
extern void demth_ZN15QTextCharFormat15setFontOverlineEb(void* qthis, bool arg0);
  // proto:  void QTextCharFormat::setTableCellColumnSpan(int tableCellColumnSpan);
extern void demth_ZN15QTextCharFormat22setTableCellColumnSpanEi(void* qthis, int arg0);
  // proto:  QPen QTextCharFormat::textOutline();
extern void demth_ZNK15QTextCharFormat11textOutlineEv(void* qthis);
  // proto:  void QTextTableFormat::QTextTableFormat();
extern void* dector_ZN16QTextTableFormatC1Ev();
extern void _ZN16QTextTableFormatC1Ev(void* qthis);
  // proto:  bool QTextTableFormat::isValid();
extern void demth_ZNK16QTextTableFormat7isValidEv(void* qthis);
  // proto:  int QTextTableFormat::headerRowCount();
extern void demth_ZNK16QTextTableFormat14headerRowCountEv(void* qthis);
  // proto:  int QTextTableFormat::columns();
extern void demth_ZNK16QTextTableFormat7columnsEv(void* qthis);
  // proto:  QVector<QTextLength> QTextTableFormat::columnWidthConstraints();
extern void demth_ZNK16QTextTableFormat22columnWidthConstraintsEv(void* qthis);
  // proto:  void QTextTableFormat::setCellPadding(qreal padding);
extern void demth_ZN16QTextTableFormat14setCellPaddingEd(void* qthis, double arg0);
  // proto:  qreal QTextTableFormat::cellPadding();
extern void demth_ZNK16QTextTableFormat11cellPaddingEv(void* qthis);
  // proto:  void QTextTableFormat::setCellSpacing(qreal spacing);
extern void demth_ZN16QTextTableFormat14setCellSpacingEd(void* qthis, double arg0);
  // proto:  void QTextTableFormat::setColumns(int columns);
extern void demth_ZN16QTextTableFormat10setColumnsEi(void* qthis, int arg0);
  // proto:  void QTextTableFormat::QTextTableFormat(const QTextFormat & fmt);
extern void* dector_ZN16QTextTableFormatC1ERK11QTextFormat(void* arg0);
extern void _ZN16QTextTableFormatC1ERK11QTextFormat(void* qthis, void* arg0);
  // proto:  void QTextTableFormat::clearColumnWidthConstraints();
extern void demth_ZN16QTextTableFormat27clearColumnWidthConstraintsEv(void* qthis);
  // proto:  void QTextTableFormat::setHeaderRowCount(int count);
extern void demth_ZN16QTextTableFormat17setHeaderRowCountEi(void* qthis, int arg0);
  // proto:  qreal QTextTableFormat::cellSpacing();
extern void demth_ZNK16QTextTableFormat11cellSpacingEv(void* qthis);
  // proto:  void QTextTableCellFormat::QTextTableCellFormat();
extern void* dector_ZN20QTextTableCellFormatC1Ev();
extern void _ZN20QTextTableCellFormatC1Ev(void* qthis);
  // proto:  void QTextTableCellFormat::setLeftPadding(qreal padding);
extern void demth_ZN20QTextTableCellFormat14setLeftPaddingEd(void* qthis, double arg0);
  // proto:  bool QTextTableCellFormat::isValid();
extern void demth_ZNK20QTextTableCellFormat7isValidEv(void* qthis);
  // proto:  void QTextTableCellFormat::setTopPadding(qreal padding);
extern void demth_ZN20QTextTableCellFormat13setTopPaddingEd(void* qthis, double arg0);
  // proto:  qreal QTextTableCellFormat::leftPadding();
extern void demth_ZNK20QTextTableCellFormat11leftPaddingEv(void* qthis);
  // proto:  void QTextTableCellFormat::setPadding(qreal padding);
extern void demth_ZN20QTextTableCellFormat10setPaddingEd(void* qthis, double arg0);
  // proto:  qreal QTextTableCellFormat::topPadding();
extern void demth_ZNK20QTextTableCellFormat10topPaddingEv(void* qthis);
  // proto:  qreal QTextTableCellFormat::rightPadding();
extern void demth_ZNK20QTextTableCellFormat12rightPaddingEv(void* qthis);
  // proto:  void QTextTableCellFormat::QTextTableCellFormat(const QTextFormat & fmt);
extern void* dector_ZN20QTextTableCellFormatC1ERK11QTextFormat(void* arg0);
extern void _ZN20QTextTableCellFormatC1ERK11QTextFormat(void* qthis, void* arg0);
  // proto:  qreal QTextTableCellFormat::bottomPadding();
extern void demth_ZNK20QTextTableCellFormat13bottomPaddingEv(void* qthis);
  // proto:  void QTextTableCellFormat::setRightPadding(qreal padding);
extern void demth_ZN20QTextTableCellFormat15setRightPaddingEd(void* qthis, double arg0);
  // proto:  void QTextTableCellFormat::setBottomPadding(qreal padding);
extern void demth_ZN20QTextTableCellFormat16setBottomPaddingEd(void* qthis, double arg0);
  // proto:  int QTextListFormat::indent();
extern void demth_ZNK15QTextListFormat6indentEv(void* qthis);
  // proto:  void QTextListFormat::QTextListFormat(const QTextFormat & fmt);
extern void* dector_ZN15QTextListFormatC1ERK11QTextFormat(void* arg0);
extern void _ZN15QTextListFormatC1ERK11QTextFormat(void* qthis, void* arg0);
  // proto:  void QTextListFormat::setIndent(int indent);
extern void demth_ZN15QTextListFormat9setIndentEi(void* qthis, int arg0);
  // proto:  QString QTextListFormat::numberSuffix();
extern void demth_ZNK15QTextListFormat12numberSuffixEv(void* qthis);
  // proto:  void QTextListFormat::QTextListFormat();
extern void* dector_ZN15QTextListFormatC1Ev();
extern void _ZN15QTextListFormatC1Ev(void* qthis);
  // proto:  QString QTextListFormat::numberPrefix();
extern void demth_ZNK15QTextListFormat12numberPrefixEv(void* qthis);
  // proto:  bool QTextListFormat::isValid();
extern void _ZNK15QTextListFormat7isValidEv(void* qthis);
  // proto:  void QTextListFormat::setNumberSuffix(const QString & numberSuffix);
extern void demth_ZN15QTextListFormat15setNumberSuffixERK7QString(void* qthis, void* arg0);
  // proto:  void QTextListFormat::setNumberPrefix(const QString & numberPrefix);
extern void demth_ZN15QTextListFormat15setNumberPrefixERK7QString(void* qthis, void* arg0);
  // proto:  bool QTextFrameFormat::isValid();
extern void _ZNK16QTextFrameFormat7isValidEv(void* qthis);
  // proto:  void QTextFrameFormat::setHeight(qreal height);
extern void demth_ZN16QTextFrameFormat9setHeightEd(void* qthis, double arg0);
  // proto:  void QTextFrameFormat::setBorderBrush(const QBrush & brush);
extern void demth_ZN16QTextFrameFormat14setBorderBrushERK6QBrush(void* qthis, void* arg0);
  // proto:  qreal QTextFrameFormat::margin();
extern void demth_ZNK16QTextFrameFormat6marginEv(void* qthis);
  // proto:  QBrush QTextFrameFormat::borderBrush();
extern void demth_ZNK16QTextFrameFormat11borderBrushEv(void* qthis);
  // proto:  void QTextFrameFormat::setRightMargin(qreal margin);
extern void demth_ZN16QTextFrameFormat14setRightMarginEd(void* qthis, double arg0);
  // proto:  void QTextFrameFormat::setMargin(qreal margin);
extern void _ZN16QTextFrameFormat9setMarginEd(void* qthis, double arg0);
  // proto:  void QTextFrameFormat::setBorder(qreal border);
extern void demth_ZN16QTextFrameFormat9setBorderEd(void* qthis, double arg0);
  // proto:  void QTextFrameFormat::setHeight(const QTextLength & height);
extern void demth_ZN16QTextFrameFormat9setHeightERK11QTextLength(void* qthis, void* arg0);
  // proto:  void QTextFrameFormat::setWidth(const QTextLength & length);
extern void demth_ZN16QTextFrameFormat8setWidthERK11QTextLength(void* qthis, void* arg0);
  // proto:  qreal QTextFrameFormat::bottomMargin();
extern void _ZNK16QTextFrameFormat12bottomMarginEv(void* qthis);
  // proto:  void QTextFrameFormat::setBottomMargin(qreal margin);
extern void demth_ZN16QTextFrameFormat15setBottomMarginEd(void* qthis, double arg0);
  // proto:  QTextLength QTextFrameFormat::height();
extern void demth_ZNK16QTextFrameFormat6heightEv(void* qthis);
  // proto:  void QTextFrameFormat::setWidth(qreal width);
extern void demth_ZN16QTextFrameFormat8setWidthEd(void* qthis, double arg0);
  // proto:  qreal QTextFrameFormat::rightMargin();
extern void _ZNK16QTextFrameFormat11rightMarginEv(void* qthis);
  // proto:  void QTextFrameFormat::setPadding(qreal padding);
extern void demth_ZN16QTextFrameFormat10setPaddingEd(void* qthis, double arg0);
  // proto:  void QTextFrameFormat::setTopMargin(qreal margin);
extern void demth_ZN16QTextFrameFormat12setTopMarginEd(void* qthis, double arg0);
  // proto:  qreal QTextFrameFormat::topMargin();
extern void _ZNK16QTextFrameFormat9topMarginEv(void* qthis);
  // proto:  QTextLength QTextFrameFormat::width();
extern void demth_ZNK16QTextFrameFormat5widthEv(void* qthis);
  // proto:  void QTextFrameFormat::QTextFrameFormat(const QTextFormat & fmt);
extern void* dector_ZN16QTextFrameFormatC1ERK11QTextFormat(void* arg0);
extern void _ZN16QTextFrameFormatC1ERK11QTextFormat(void* qthis, void* arg0);
  // proto:  qreal QTextFrameFormat::padding();
extern void demth_ZNK16QTextFrameFormat7paddingEv(void* qthis);
  // proto:  void QTextFrameFormat::setLeftMargin(qreal margin);
extern void demth_ZN16QTextFrameFormat13setLeftMarginEd(void* qthis, double arg0);
  // proto:  qreal QTextFrameFormat::border();
extern void demth_ZNK16QTextFrameFormat6borderEv(void* qthis);
  // proto:  void QTextFrameFormat::QTextFrameFormat();
extern void* dector_ZN16QTextFrameFormatC1Ev();
extern void _ZN16QTextFrameFormatC1Ev(void* qthis);
  // proto:  qreal QTextFrameFormat::leftMargin();
extern void _ZNK16QTextFrameFormat10leftMarginEv(void* qthis);
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

// class sizeof(QTextLength)=16
type QTextLength struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextImageFormat)=1
type QTextImageFormat struct {
  /*qbase*/ QTextCharFormat;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextFormat)=1
type QTextFormat struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextBlockFormat)=1
type QTextBlockFormat struct {
  /*qbase*/ QTextFormat;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextCharFormat)=1
type QTextCharFormat struct {
  /*qbase*/ QTextFormat;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextTableFormat)=1
type QTextTableFormat struct {
  /*qbase*/ QTextFrameFormat;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextTableCellFormat)=1
type QTextTableCellFormat struct {
  /*qbase*/ QTextCharFormat;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextListFormat)=1
type QTextListFormat struct {
  /*qbase*/ QTextFormat;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextFrameFormat)=1
type QTextFrameFormat struct {
  /*qbase*/ QTextFormat;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  qreal QTextLength::value(qreal maximumLength);
func (this *QTextLength) value(args ...interface{}) () {
  // value(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLength5valueEd
  default:
    qtrt.ErrorResolve("QTextLength", "value", args)
  }

}

  // proto:  void QTextLength::QTextLength();
func NewQTextLength(args ...interface{}) QTextLength {
  return QTextLength{}
}

  // proto:  qreal QTextLength::rawValue();
func (this *QTextLength) rawValue(args ...interface{}) () {
  // rawValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLength8rawValueEv
  default:
    qtrt.ErrorResolve("QTextLength", "rawValue", args)
  }

}

  // proto:  void QTextImageFormat::QTextImageFormat();
func NewQTextImageFormat(args ...interface{}) QTextImageFormat {
  return QTextImageFormat{}
}

  // proto:  bool QTextImageFormat::isValid();
func (this *QTextImageFormat) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextImageFormat7isValidEv
  default:
    qtrt.ErrorResolve("QTextImageFormat", "isValid", args)
  }

}

  // proto:  qreal QTextImageFormat::width();
func (this *QTextImageFormat) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextImageFormat5widthEv
  default:
    qtrt.ErrorResolve("QTextImageFormat", "width", args)
  }

}

  // proto:  void QTextImageFormat::setHeight(qreal height);
func (this *QTextImageFormat) setHeight(args ...interface{}) () {
  // setHeight(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextImageFormat9setHeightEd
  default:
    qtrt.ErrorResolve("QTextImageFormat", "setHeight", args)
  }

}

  // proto:  void QTextImageFormat::setWidth(qreal width);
func (this *QTextImageFormat) setWidth(args ...interface{}) () {
  // setWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextImageFormat8setWidthEd
  default:
    qtrt.ErrorResolve("QTextImageFormat", "setWidth", args)
  }

}

  // proto:  void QTextImageFormat::setName(const QString & name);
func (this *QTextImageFormat) setName(args ...interface{}) () {
  // setName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextImageFormat7setNameERK7QString
  default:
    qtrt.ErrorResolve("QTextImageFormat", "setName", args)
  }

}

  // proto:  QString QTextImageFormat::name();
func (this *QTextImageFormat) name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextImageFormat4nameEv
  default:
    qtrt.ErrorResolve("QTextImageFormat", "name", args)
  }

}

  // proto:  qreal QTextImageFormat::height();
func (this *QTextImageFormat) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextImageFormat6heightEv
  default:
    qtrt.ErrorResolve("QTextImageFormat", "height", args)
  }

}

  // proto:  QTextBlockFormat QTextFormat::toBlockFormat();
func (this *QTextFormat) toBlockFormat(args ...interface{}) () {
  // toBlockFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13toBlockFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "toBlockFormat", args)
  }

}

  // proto:  QString QTextFormat::stringProperty(int propertyId);
func (this *QTextFormat) stringProperty(args ...interface{}) () {
  // stringProperty(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat14stringPropertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "stringProperty", args)
  }

}

  // proto:  QVector<QTextLength> QTextFormat::lengthVectorProperty(int propertyId);
func (this *QTextFormat) lengthVectorProperty(args ...interface{}) () {
  // lengthVectorProperty(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat20lengthVectorPropertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "lengthVectorProperty", args)
  }

}

  // proto:  int QTextFormat::objectIndex();
func (this *QTextFormat) objectIndex(args ...interface{}) () {
  // objectIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat11objectIndexEv
  default:
    qtrt.ErrorResolve("QTextFormat", "objectIndex", args)
  }

}

  // proto:  void QTextFormat::setObjectIndex(int object);
func (this *QTextFormat) setObjectIndex(args ...interface{}) () {
  // setObjectIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat14setObjectIndexEi
  default:
    qtrt.ErrorResolve("QTextFormat", "setObjectIndex", args)
  }

}

  // proto:  void QTextFormat::clearForeground();
func (this *QTextFormat) clearForeground(args ...interface{}) () {
  // clearForeground()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat15clearForegroundEv
  default:
    qtrt.ErrorResolve("QTextFormat", "clearForeground", args)
  }

}

  // proto:  bool QTextFormat::isTableCellFormat();
func (this *QTextFormat) isTableCellFormat(args ...interface{}) () {
  // isTableCellFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat17isTableCellFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "isTableCellFormat", args)
  }

}

  // proto:  void QTextFormat::~QTextFormat();
func (this *QTextFormat) FreeQTextFormat(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextFormat", "~QTextFormat", args)
  }

}

  // proto:  bool QTextFormat::isValid();
func (this *QTextFormat) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat7isValidEv
  default:
    qtrt.ErrorResolve("QTextFormat", "isValid", args)
  }

}

  // proto:  void QTextFormat::QTextFormat(const QTextFormat & rhs);
func NewQTextFormat(args ...interface{}) QTextFormat {
  return QTextFormat{}
}

  // proto:  QTextLength QTextFormat::lengthProperty(int propertyId);
func (this *QTextFormat) lengthProperty(args ...interface{}) () {
  // lengthProperty(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat14lengthPropertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "lengthProperty", args)
  }

}

  // proto:  void QTextFormat::merge(const QTextFormat & other);
func (this *QTextFormat) merge(args ...interface{}) () {
  // merge(const class QTextFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextFormat{}) // "const QTextFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat5mergeERKS_
  default:
    qtrt.ErrorResolve("QTextFormat", "merge", args)
  }

}

  // proto:  QColor QTextFormat::colorProperty(int propertyId);
func (this *QTextFormat) colorProperty(args ...interface{}) () {
  // colorProperty(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13colorPropertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "colorProperty", args)
  }

}

  // proto:  void QTextFormat::setForeground(const QBrush & brush);
func (this *QTextFormat) setForeground(args ...interface{}) () {
  // setForeground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat13setForegroundERK6QBrush
  default:
    qtrt.ErrorResolve("QTextFormat", "setForeground", args)
  }

}

  // proto:  bool QTextFormat::boolProperty(int propertyId);
func (this *QTextFormat) boolProperty(args ...interface{}) () {
  // boolProperty(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat12boolPropertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "boolProperty", args)
  }

}

  // proto:  bool QTextFormat::isListFormat();
func (this *QTextFormat) isListFormat(args ...interface{}) () {
  // isListFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat12isListFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "isListFormat", args)
  }

}

  // proto:  bool QTextFormat::isImageFormat();
func (this *QTextFormat) isImageFormat(args ...interface{}) () {
  // isImageFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13isImageFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "isImageFormat", args)
  }

}

  // proto:  void QTextFormat::clearProperty(int propertyId);
func (this *QTextFormat) clearProperty(args ...interface{}) () {
  // clearProperty(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat13clearPropertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "clearProperty", args)
  }

}

  // proto:  QTextFrameFormat QTextFormat::toFrameFormat();
func (this *QTextFormat) toFrameFormat(args ...interface{}) () {
  // toFrameFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13toFrameFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "toFrameFormat", args)
  }

}

  // proto:  QBrush QTextFormat::brushProperty(int propertyId);
func (this *QTextFormat) brushProperty(args ...interface{}) () {
  // brushProperty(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13brushPropertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "brushProperty", args)
  }

}

  // proto:  int QTextFormat::propertyCount();
func (this *QTextFormat) propertyCount(args ...interface{}) () {
  // propertyCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13propertyCountEv
  default:
    qtrt.ErrorResolve("QTextFormat", "propertyCount", args)
  }

}

  // proto:  QPen QTextFormat::penProperty(int propertyId);
func (this *QTextFormat) penProperty(args ...interface{}) () {
  // penProperty(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat11penPropertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "penProperty", args)
  }

}

  // proto:  QVariant QTextFormat::property(int propertyId);
func (this *QTextFormat) property(args ...interface{}) () {
  // property(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat8propertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "property", args)
  }

}

  // proto:  bool QTextFormat::isTableFormat();
func (this *QTextFormat) isTableFormat(args ...interface{}) () {
  // isTableFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13isTableFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "isTableFormat", args)
  }

}

  // proto:  void QTextFormat::setProperty(int propertyId, const QVariant & value);
func (this *QTextFormat) setProperty(args ...interface{}) () {
  // setProperty(int, const QVector<class QTextLength> &)
  // setProperty(int, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  // vtys[0][1] = reflect.TypeOf(QVector<QTextLength>{}) // "const QVector<QTextLength> &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat11setPropertyEiRK7QVectorI11QTextLengthE
  case 1:
    // invoke: _ZN11QTextFormat11setPropertyEiRK8QVariant
  default:
    qtrt.ErrorResolve("QTextFormat", "setProperty", args)
  }

}

  // proto:  int QTextFormat::type();
func (this *QTextFormat) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextFormat", "type", args)
  }

}

  // proto:  bool QTextFormat::isCharFormat();
func (this *QTextFormat) isCharFormat(args ...interface{}) () {
  // isCharFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat12isCharFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "isCharFormat", args)
  }

}

  // proto:  void QTextFormat::clearBackground();
func (this *QTextFormat) clearBackground(args ...interface{}) () {
  // clearBackground()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat15clearBackgroundEv
  default:
    qtrt.ErrorResolve("QTextFormat", "clearBackground", args)
  }

}

  // proto:  bool QTextFormat::isBlockFormat();
func (this *QTextFormat) isBlockFormat(args ...interface{}) () {
  // isBlockFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13isBlockFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "isBlockFormat", args)
  }

}

  // proto:  QBrush QTextFormat::background();
func (this *QTextFormat) background(args ...interface{}) () {
  // background()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat10backgroundEv
  default:
    qtrt.ErrorResolve("QTextFormat", "background", args)
  }

}

  // proto:  qreal QTextFormat::doubleProperty(int propertyId);
func (this *QTextFormat) doubleProperty(args ...interface{}) () {
  // doubleProperty(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat14doublePropertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "doubleProperty", args)
  }

}

  // proto:  void QTextFormat::swap(QTextFormat & other);
func (this *QTextFormat) swap(args ...interface{}) () {
  // swap(class QTextFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextFormat{}) // "QTextFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat4swapERS_
  default:
    qtrt.ErrorResolve("QTextFormat", "swap", args)
  }

}

  // proto:  QTextImageFormat QTextFormat::toImageFormat();
func (this *QTextFormat) toImageFormat(args ...interface{}) () {
  // toImageFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13toImageFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "toImageFormat", args)
  }

}

  // proto:  bool QTextFormat::hasProperty(int propertyId);
func (this *QTextFormat) hasProperty(args ...interface{}) () {
  // hasProperty(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat11hasPropertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "hasProperty", args)
  }

}

  // proto:  QBrush QTextFormat::foreground();
func (this *QTextFormat) foreground(args ...interface{}) () {
  // foreground()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat10foregroundEv
  default:
    qtrt.ErrorResolve("QTextFormat", "foreground", args)
  }

}

  // proto:  void QTextFormat::setObjectType(int type);
func (this *QTextFormat) setObjectType(args ...interface{}) () {
  // setObjectType(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat13setObjectTypeEi
  default:
    qtrt.ErrorResolve("QTextFormat", "setObjectType", args)
  }

}

  // proto:  void QTextFormat::setBackground(const QBrush & brush);
func (this *QTextFormat) setBackground(args ...interface{}) () {
  // setBackground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat13setBackgroundERK6QBrush
  default:
    qtrt.ErrorResolve("QTextFormat", "setBackground", args)
  }

}

  // proto:  QTextTableFormat QTextFormat::toTableFormat();
func (this *QTextFormat) toTableFormat(args ...interface{}) () {
  // toTableFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13toTableFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "toTableFormat", args)
  }

}

  // proto:  bool QTextFormat::isFrameFormat();
func (this *QTextFormat) isFrameFormat(args ...interface{}) () {
  // isFrameFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13isFrameFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "isFrameFormat", args)
  }

}

  // proto:  int QTextFormat::intProperty(int propertyId);
func (this *QTextFormat) intProperty(args ...interface{}) () {
  // intProperty(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat11intPropertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "intProperty", args)
  }

}

  // proto:  QTextCharFormat QTextFormat::toCharFormat();
func (this *QTextFormat) toCharFormat(args ...interface{}) () {
  // toCharFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat12toCharFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "toCharFormat", args)
  }

}

  // proto:  bool QTextFormat::isEmpty();
func (this *QTextFormat) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat7isEmptyEv
  default:
    qtrt.ErrorResolve("QTextFormat", "isEmpty", args)
  }

}

  // proto:  QTextTableCellFormat QTextFormat::toTableCellFormat();
func (this *QTextFormat) toTableCellFormat(args ...interface{}) () {
  // toTableCellFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat17toTableCellFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "toTableCellFormat", args)
  }

}

  // proto:  int QTextFormat::objectType();
func (this *QTextFormat) objectType(args ...interface{}) () {
  // objectType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat10objectTypeEv
  default:
    qtrt.ErrorResolve("QTextFormat", "objectType", args)
  }

}

  // proto:  QTextListFormat QTextFormat::toListFormat();
func (this *QTextFormat) toListFormat(args ...interface{}) () {
  // toListFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat12toListFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "toListFormat", args)
  }

}

  // proto:  QMap<int, QVariant> QTextFormat::properties();
func (this *QTextFormat) properties(args ...interface{}) () {
  // properties()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat10propertiesEv
  default:
    qtrt.ErrorResolve("QTextFormat", "properties", args)
  }

}

  // proto:  int QTextBlockFormat::indent();
func (this *QTextBlockFormat) indent(args ...interface{}) () {
  // indent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat6indentEv
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "indent", args)
  }

}

  // proto:  void QTextBlockFormat::setTextIndent(qreal aindent);
func (this *QTextBlockFormat) setTextIndent(args ...interface{}) () {
  // setTextIndent(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextBlockFormat13setTextIndentEd
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setTextIndent", args)
  }

}

  // proto:  void QTextBlockFormat::setNonBreakableLines(bool b);
func (this *QTextBlockFormat) setNonBreakableLines(args ...interface{}) () {
  // setNonBreakableLines(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextBlockFormat20setNonBreakableLinesEb
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setNonBreakableLines", args)
  }

}

  // proto:  void QTextBlockFormat::setIndent(int indent);
func (this *QTextBlockFormat) setIndent(args ...interface{}) () {
  // setIndent(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextBlockFormat9setIndentEi
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setIndent", args)
  }

}

  // proto:  qreal QTextBlockFormat::textIndent();
func (this *QTextBlockFormat) textIndent(args ...interface{}) () {
  // textIndent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat10textIndentEv
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "textIndent", args)
  }

}

  // proto:  qreal QTextBlockFormat::lineHeight();
func (this *QTextBlockFormat) lineHeight(args ...interface{}) () {
  // lineHeight()
  // lineHeight(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat10lineHeightEv
  case 1:
    // invoke: _ZNK16QTextBlockFormat10lineHeightEdd
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "lineHeight", args)
  }

}

  // proto:  void QTextBlockFormat::QTextBlockFormat(const QTextFormat & fmt);
func NewQTextBlockFormat(args ...interface{}) QTextBlockFormat {
  return QTextBlockFormat{}
}

  // proto:  void QTextBlockFormat::setRightMargin(qreal margin);
func (this *QTextBlockFormat) setRightMargin(args ...interface{}) () {
  // setRightMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextBlockFormat14setRightMarginEd
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setRightMargin", args)
  }

}

  // proto:  qreal QTextBlockFormat::topMargin();
func (this *QTextBlockFormat) topMargin(args ...interface{}) () {
  // topMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat9topMarginEv
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "topMargin", args)
  }

}

  // proto:  qreal QTextBlockFormat::rightMargin();
func (this *QTextBlockFormat) rightMargin(args ...interface{}) () {
  // rightMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat11rightMarginEv
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "rightMargin", args)
  }

}

  // proto:  qreal QTextBlockFormat::bottomMargin();
func (this *QTextBlockFormat) bottomMargin(args ...interface{}) () {
  // bottomMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat12bottomMarginEv
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "bottomMargin", args)
  }

}

  // proto:  void QTextBlockFormat::setTopMargin(qreal margin);
func (this *QTextBlockFormat) setTopMargin(args ...interface{}) () {
  // setTopMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextBlockFormat12setTopMarginEd
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setTopMargin", args)
  }

}

  // proto:  qreal QTextBlockFormat::leftMargin();
func (this *QTextBlockFormat) leftMargin(args ...interface{}) () {
  // leftMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat10leftMarginEv
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "leftMargin", args)
  }

}

  // proto:  void QTextBlockFormat::setLineHeight(qreal height, int heightType);
func (this *QTextBlockFormat) setLineHeight(args ...interface{}) () {
  // setLineHeight(qreal, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextBlockFormat13setLineHeightEdi
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setLineHeight", args)
  }

}

  // proto:  void QTextBlockFormat::setBottomMargin(qreal margin);
func (this *QTextBlockFormat) setBottomMargin(args ...interface{}) () {
  // setBottomMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextBlockFormat15setBottomMarginEd
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setBottomMargin", args)
  }

}

  // proto:  int QTextBlockFormat::lineHeightType();
func (this *QTextBlockFormat) lineHeightType(args ...interface{}) () {
  // lineHeightType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat14lineHeightTypeEv
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "lineHeightType", args)
  }

}

  // proto:  void QTextBlockFormat::setLeftMargin(qreal margin);
func (this *QTextBlockFormat) setLeftMargin(args ...interface{}) () {
  // setLeftMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextBlockFormat13setLeftMarginEd
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setLeftMargin", args)
  }

}

  // proto:  bool QTextBlockFormat::isValid();
func (this *QTextBlockFormat) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat7isValidEv
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "isValid", args)
  }

}

  // proto:  bool QTextBlockFormat::nonBreakableLines();
func (this *QTextBlockFormat) nonBreakableLines(args ...interface{}) () {
  // nonBreakableLines()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat17nonBreakableLinesEv
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "nonBreakableLines", args)
  }

}

  // proto:  void QTextCharFormat::setFontLetterSpacing(qreal spacing);
func (this *QTextCharFormat) setFontLetterSpacing(args ...interface{}) () {
  // setFontLetterSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat20setFontLetterSpacingEd
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontLetterSpacing", args)
  }

}

  // proto:  bool QTextCharFormat::isAnchor();
func (this *QTextCharFormat) isAnchor(args ...interface{}) () {
  // isAnchor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat8isAnchorEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "isAnchor", args)
  }

}

  // proto:  void QTextCharFormat::setFont(const QFont & font);
func (this *QTextCharFormat) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  // setFont(const class QFont &, enum QTextCharFormat::FontPropertiesInheritanceBehavior)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QFont{}) // "const QFont &"
  vtys[1][1] = qtrt.Int32Ty(false) // "QTextCharFormat::FontPropertiesInheritanceBehavior"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat7setFontERK5QFont
  case 1:
    // invoke: _ZN15QTextCharFormat7setFontERK5QFontNS_33FontPropertiesInheritanceBehaviorE
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFont", args)
  }

}

  // proto:  bool QTextCharFormat::fontOverline();
func (this *QTextCharFormat) fontOverline(args ...interface{}) () {
  // fontOverline()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat12fontOverlineEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontOverline", args)
  }

}

  // proto:  QFont QTextCharFormat::font();
func (this *QTextCharFormat) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat4fontEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "font", args)
  }

}

  // proto:  QString QTextCharFormat::fontFamily();
func (this *QTextCharFormat) fontFamily(args ...interface{}) () {
  // fontFamily()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat10fontFamilyEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontFamily", args)
  }

}

  // proto:  bool QTextCharFormat::fontStrikeOut();
func (this *QTextCharFormat) fontStrikeOut(args ...interface{}) () {
  // fontStrikeOut()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat13fontStrikeOutEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontStrikeOut", args)
  }

}

  // proto:  void QTextCharFormat::setFontPointSize(qreal size);
func (this *QTextCharFormat) setFontPointSize(args ...interface{}) () {
  // setFontPointSize(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat16setFontPointSizeEd
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontPointSize", args)
  }

}

  // proto:  void QTextCharFormat::setUnderlineColor(const QColor & color);
func (this *QTextCharFormat) setUnderlineColor(args ...interface{}) () {
  // setUnderlineColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat17setUnderlineColorERK6QColor
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setUnderlineColor", args)
  }

}

  // proto:  int QTextCharFormat::tableCellRowSpan();
func (this *QTextCharFormat) tableCellRowSpan(args ...interface{}) () {
  // tableCellRowSpan()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat16tableCellRowSpanEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "tableCellRowSpan", args)
  }

}

  // proto:  void QTextCharFormat::setFontUnderline(bool underline);
func (this *QTextCharFormat) setFontUnderline(args ...interface{}) () {
  // setFontUnderline(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat16setFontUnderlineEb
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontUnderline", args)
  }

}

  // proto:  bool QTextCharFormat::isValid();
func (this *QTextCharFormat) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat7isValidEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "isValid", args)
  }

}

  // proto:  bool QTextCharFormat::fontItalic();
func (this *QTextCharFormat) fontItalic(args ...interface{}) () {
  // fontItalic()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat10fontItalicEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontItalic", args)
  }

}

  // proto:  void QTextCharFormat::setToolTip(const QString & tip);
func (this *QTextCharFormat) setToolTip(args ...interface{}) () {
  // setToolTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat10setToolTipERK7QString
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setToolTip", args)
  }

}

  // proto:  void QTextCharFormat::setTextOutline(const QPen & pen);
func (this *QTextCharFormat) setTextOutline(args ...interface{}) () {
  // setTextOutline(const class QPen &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPen{}) // "const QPen &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat14setTextOutlineERK4QPen
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setTextOutline", args)
  }

}

  // proto:  void QTextCharFormat::setTableCellRowSpan(int tableCellRowSpan);
func (this *QTextCharFormat) setTableCellRowSpan(args ...interface{}) () {
  // setTableCellRowSpan(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat19setTableCellRowSpanEi
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setTableCellRowSpan", args)
  }

}

  // proto:  void QTextCharFormat::setAnchor(bool anchor);
func (this *QTextCharFormat) setAnchor(args ...interface{}) () {
  // setAnchor(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat9setAnchorEb
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setAnchor", args)
  }

}

  // proto:  qreal QTextCharFormat::fontPointSize();
func (this *QTextCharFormat) fontPointSize(args ...interface{}) () {
  // fontPointSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat13fontPointSizeEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontPointSize", args)
  }

}

  // proto:  void QTextCharFormat::QTextCharFormat(const QTextFormat & fmt);
func NewQTextCharFormat(args ...interface{}) QTextCharFormat {
  return QTextCharFormat{}
}

  // proto:  void QTextCharFormat::setFontStrikeOut(bool strikeOut);
func (this *QTextCharFormat) setFontStrikeOut(args ...interface{}) () {
  // setFontStrikeOut(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat16setFontStrikeOutEb
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontStrikeOut", args)
  }

}

  // proto:  qreal QTextCharFormat::fontWordSpacing();
func (this *QTextCharFormat) fontWordSpacing(args ...interface{}) () {
  // fontWordSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat15fontWordSpacingEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontWordSpacing", args)
  }

}

  // proto:  QString QTextCharFormat::toolTip();
func (this *QTextCharFormat) toolTip(args ...interface{}) () {
  // toolTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat7toolTipEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "toolTip", args)
  }

}

  // proto:  void QTextCharFormat::setAnchorNames(const QStringList & names);
func (this *QTextCharFormat) setAnchorNames(args ...interface{}) () {
  // setAnchorNames(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat14setAnchorNamesERK11QStringList
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setAnchorNames", args)
  }

}

  // proto:  QStringList QTextCharFormat::anchorNames();
func (this *QTextCharFormat) anchorNames(args ...interface{}) () {
  // anchorNames()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat11anchorNamesEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "anchorNames", args)
  }

}

  // proto:  void QTextCharFormat::setFontFixedPitch(bool fixedPitch);
func (this *QTextCharFormat) setFontFixedPitch(args ...interface{}) () {
  // setFontFixedPitch(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat17setFontFixedPitchEb
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontFixedPitch", args)
  }

}

  // proto:  void QTextCharFormat::setFontItalic(bool italic);
func (this *QTextCharFormat) setFontItalic(args ...interface{}) () {
  // setFontItalic(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat13setFontItalicEb
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontItalic", args)
  }

}

  // proto:  void QTextCharFormat::setFontFamily(const QString & family);
func (this *QTextCharFormat) setFontFamily(args ...interface{}) () {
  // setFontFamily(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat13setFontFamilyERK7QString
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontFamily", args)
  }

}

  // proto:  bool QTextCharFormat::fontFixedPitch();
func (this *QTextCharFormat) fontFixedPitch(args ...interface{}) () {
  // fontFixedPitch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat14fontFixedPitchEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontFixedPitch", args)
  }

}

  // proto:  void QTextCharFormat::setAnchorHref(const QString & value);
func (this *QTextCharFormat) setAnchorHref(args ...interface{}) () {
  // setAnchorHref(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat13setAnchorHrefERK7QString
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setAnchorHref", args)
  }

}

  // proto:  int QTextCharFormat::fontStretch();
func (this *QTextCharFormat) fontStretch(args ...interface{}) () {
  // fontStretch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat11fontStretchEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontStretch", args)
  }

}

  // proto:  void QTextCharFormat::setFontKerning(bool enable);
func (this *QTextCharFormat) setFontKerning(args ...interface{}) () {
  // setFontKerning(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat14setFontKerningEb
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontKerning", args)
  }

}

  // proto:  int QTextCharFormat::tableCellColumnSpan();
func (this *QTextCharFormat) tableCellColumnSpan(args ...interface{}) () {
  // tableCellColumnSpan()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat19tableCellColumnSpanEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "tableCellColumnSpan", args)
  }

}

  // proto:  qreal QTextCharFormat::fontLetterSpacing();
func (this *QTextCharFormat) fontLetterSpacing(args ...interface{}) () {
  // fontLetterSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat17fontLetterSpacingEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontLetterSpacing", args)
  }

}

  // proto:  QString QTextCharFormat::anchorHref();
func (this *QTextCharFormat) anchorHref(args ...interface{}) () {
  // anchorHref()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat10anchorHrefEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "anchorHref", args)
  }

}

  // proto:  QString QTextCharFormat::anchorName();
func (this *QTextCharFormat) anchorName(args ...interface{}) () {
  // anchorName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat10anchorNameEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "anchorName", args)
  }

}

  // proto:  void QTextCharFormat::setFontStretch(int factor);
func (this *QTextCharFormat) setFontStretch(args ...interface{}) () {
  // setFontStretch(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat14setFontStretchEi
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontStretch", args)
  }

}

  // proto:  void QTextCharFormat::setAnchorName(const QString & name);
func (this *QTextCharFormat) setAnchorName(args ...interface{}) () {
  // setAnchorName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat13setAnchorNameERK7QString
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setAnchorName", args)
  }

}

  // proto:  bool QTextCharFormat::fontKerning();
func (this *QTextCharFormat) fontKerning(args ...interface{}) () {
  // fontKerning()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat11fontKerningEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontKerning", args)
  }

}

  // proto:  void QTextCharFormat::setFontWeight(int weight);
func (this *QTextCharFormat) setFontWeight(args ...interface{}) () {
  // setFontWeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat13setFontWeightEi
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontWeight", args)
  }

}

  // proto:  bool QTextCharFormat::fontUnderline();
func (this *QTextCharFormat) fontUnderline(args ...interface{}) () {
  // fontUnderline()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat13fontUnderlineEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontUnderline", args)
  }

}

  // proto:  void QTextCharFormat::setFontWordSpacing(qreal spacing);
func (this *QTextCharFormat) setFontWordSpacing(args ...interface{}) () {
  // setFontWordSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat18setFontWordSpacingEd
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontWordSpacing", args)
  }

}

  // proto:  QColor QTextCharFormat::underlineColor();
func (this *QTextCharFormat) underlineColor(args ...interface{}) () {
  // underlineColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat14underlineColorEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "underlineColor", args)
  }

}

  // proto:  int QTextCharFormat::fontWeight();
func (this *QTextCharFormat) fontWeight(args ...interface{}) () {
  // fontWeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat10fontWeightEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontWeight", args)
  }

}

  // proto:  void QTextCharFormat::setFontOverline(bool overline);
func (this *QTextCharFormat) setFontOverline(args ...interface{}) () {
  // setFontOverline(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat15setFontOverlineEb
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontOverline", args)
  }

}

  // proto:  void QTextCharFormat::setTableCellColumnSpan(int tableCellColumnSpan);
func (this *QTextCharFormat) setTableCellColumnSpan(args ...interface{}) () {
  // setTableCellColumnSpan(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat22setTableCellColumnSpanEi
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setTableCellColumnSpan", args)
  }

}

  // proto:  QPen QTextCharFormat::textOutline();
func (this *QTextCharFormat) textOutline(args ...interface{}) () {
  // textOutline()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat11textOutlineEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "textOutline", args)
  }

}

  // proto:  void QTextTableFormat::QTextTableFormat();
func NewQTextTableFormat(args ...interface{}) QTextTableFormat {
  return QTextTableFormat{}
}

  // proto:  bool QTextTableFormat::isValid();
func (this *QTextTableFormat) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextTableFormat7isValidEv
  default:
    qtrt.ErrorResolve("QTextTableFormat", "isValid", args)
  }

}

  // proto:  int QTextTableFormat::headerRowCount();
func (this *QTextTableFormat) headerRowCount(args ...interface{}) () {
  // headerRowCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextTableFormat14headerRowCountEv
  default:
    qtrt.ErrorResolve("QTextTableFormat", "headerRowCount", args)
  }

}

  // proto:  int QTextTableFormat::columns();
func (this *QTextTableFormat) columns(args ...interface{}) () {
  // columns()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextTableFormat7columnsEv
  default:
    qtrt.ErrorResolve("QTextTableFormat", "columns", args)
  }

}

  // proto:  QVector<QTextLength> QTextTableFormat::columnWidthConstraints();
func (this *QTextTableFormat) columnWidthConstraints(args ...interface{}) () {
  // columnWidthConstraints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextTableFormat22columnWidthConstraintsEv
  default:
    qtrt.ErrorResolve("QTextTableFormat", "columnWidthConstraints", args)
  }

}

  // proto:  void QTextTableFormat::setCellPadding(qreal padding);
func (this *QTextTableFormat) setCellPadding(args ...interface{}) () {
  // setCellPadding(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextTableFormat14setCellPaddingEd
  default:
    qtrt.ErrorResolve("QTextTableFormat", "setCellPadding", args)
  }

}

  // proto:  qreal QTextTableFormat::cellPadding();
func (this *QTextTableFormat) cellPadding(args ...interface{}) () {
  // cellPadding()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextTableFormat11cellPaddingEv
  default:
    qtrt.ErrorResolve("QTextTableFormat", "cellPadding", args)
  }

}

  // proto:  void QTextTableFormat::setCellSpacing(qreal spacing);
func (this *QTextTableFormat) setCellSpacing(args ...interface{}) () {
  // setCellSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextTableFormat14setCellSpacingEd
  default:
    qtrt.ErrorResolve("QTextTableFormat", "setCellSpacing", args)
  }

}

  // proto:  void QTextTableFormat::setColumns(int columns);
func (this *QTextTableFormat) setColumns(args ...interface{}) () {
  // setColumns(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextTableFormat10setColumnsEi
  default:
    qtrt.ErrorResolve("QTextTableFormat", "setColumns", args)
  }

}

  // proto:  void QTextTableFormat::clearColumnWidthConstraints();
func (this *QTextTableFormat) clearColumnWidthConstraints(args ...interface{}) () {
  // clearColumnWidthConstraints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextTableFormat27clearColumnWidthConstraintsEv
  default:
    qtrt.ErrorResolve("QTextTableFormat", "clearColumnWidthConstraints", args)
  }

}

  // proto:  void QTextTableFormat::setHeaderRowCount(int count);
func (this *QTextTableFormat) setHeaderRowCount(args ...interface{}) () {
  // setHeaderRowCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextTableFormat17setHeaderRowCountEi
  default:
    qtrt.ErrorResolve("QTextTableFormat", "setHeaderRowCount", args)
  }

}

  // proto:  qreal QTextTableFormat::cellSpacing();
func (this *QTextTableFormat) cellSpacing(args ...interface{}) () {
  // cellSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextTableFormat11cellSpacingEv
  default:
    qtrt.ErrorResolve("QTextTableFormat", "cellSpacing", args)
  }

}

  // proto:  void QTextTableCellFormat::QTextTableCellFormat();
func NewQTextTableCellFormat(args ...interface{}) QTextTableCellFormat {
  return QTextTableCellFormat{}
}

  // proto:  void QTextTableCellFormat::setLeftPadding(qreal padding);
func (this *QTextTableCellFormat) setLeftPadding(args ...interface{}) () {
  // setLeftPadding(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QTextTableCellFormat14setLeftPaddingEd
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "setLeftPadding", args)
  }

}

  // proto:  bool QTextTableCellFormat::isValid();
func (this *QTextTableCellFormat) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QTextTableCellFormat7isValidEv
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "isValid", args)
  }

}

  // proto:  void QTextTableCellFormat::setTopPadding(qreal padding);
func (this *QTextTableCellFormat) setTopPadding(args ...interface{}) () {
  // setTopPadding(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QTextTableCellFormat13setTopPaddingEd
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "setTopPadding", args)
  }

}

  // proto:  qreal QTextTableCellFormat::leftPadding();
func (this *QTextTableCellFormat) leftPadding(args ...interface{}) () {
  // leftPadding()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QTextTableCellFormat11leftPaddingEv
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "leftPadding", args)
  }

}

  // proto:  void QTextTableCellFormat::setPadding(qreal padding);
func (this *QTextTableCellFormat) setPadding(args ...interface{}) () {
  // setPadding(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QTextTableCellFormat10setPaddingEd
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "setPadding", args)
  }

}

  // proto:  qreal QTextTableCellFormat::topPadding();
func (this *QTextTableCellFormat) topPadding(args ...interface{}) () {
  // topPadding()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QTextTableCellFormat10topPaddingEv
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "topPadding", args)
  }

}

  // proto:  qreal QTextTableCellFormat::rightPadding();
func (this *QTextTableCellFormat) rightPadding(args ...interface{}) () {
  // rightPadding()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QTextTableCellFormat12rightPaddingEv
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "rightPadding", args)
  }

}

  // proto:  qreal QTextTableCellFormat::bottomPadding();
func (this *QTextTableCellFormat) bottomPadding(args ...interface{}) () {
  // bottomPadding()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QTextTableCellFormat13bottomPaddingEv
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "bottomPadding", args)
  }

}

  // proto:  void QTextTableCellFormat::setRightPadding(qreal padding);
func (this *QTextTableCellFormat) setRightPadding(args ...interface{}) () {
  // setRightPadding(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QTextTableCellFormat15setRightPaddingEd
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "setRightPadding", args)
  }

}

  // proto:  void QTextTableCellFormat::setBottomPadding(qreal padding);
func (this *QTextTableCellFormat) setBottomPadding(args ...interface{}) () {
  // setBottomPadding(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QTextTableCellFormat16setBottomPaddingEd
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "setBottomPadding", args)
  }

}

  // proto:  int QTextListFormat::indent();
func (this *QTextListFormat) indent(args ...interface{}) () {
  // indent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextListFormat6indentEv
  default:
    qtrt.ErrorResolve("QTextListFormat", "indent", args)
  }

}

  // proto:  void QTextListFormat::QTextListFormat(const QTextFormat & fmt);
func NewQTextListFormat(args ...interface{}) QTextListFormat {
  return QTextListFormat{}
}

  // proto:  void QTextListFormat::setIndent(int indent);
func (this *QTextListFormat) setIndent(args ...interface{}) () {
  // setIndent(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextListFormat9setIndentEi
  default:
    qtrt.ErrorResolve("QTextListFormat", "setIndent", args)
  }

}

  // proto:  QString QTextListFormat::numberSuffix();
func (this *QTextListFormat) numberSuffix(args ...interface{}) () {
  // numberSuffix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextListFormat12numberSuffixEv
  default:
    qtrt.ErrorResolve("QTextListFormat", "numberSuffix", args)
  }

}

  // proto:  QString QTextListFormat::numberPrefix();
func (this *QTextListFormat) numberPrefix(args ...interface{}) () {
  // numberPrefix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextListFormat12numberPrefixEv
  default:
    qtrt.ErrorResolve("QTextListFormat", "numberPrefix", args)
  }

}

  // proto:  bool QTextListFormat::isValid();
func (this *QTextListFormat) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextListFormat7isValidEv
  default:
    qtrt.ErrorResolve("QTextListFormat", "isValid", args)
  }

}

  // proto:  void QTextListFormat::setNumberSuffix(const QString & numberSuffix);
func (this *QTextListFormat) setNumberSuffix(args ...interface{}) () {
  // setNumberSuffix(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextListFormat15setNumberSuffixERK7QString
  default:
    qtrt.ErrorResolve("QTextListFormat", "setNumberSuffix", args)
  }

}

  // proto:  void QTextListFormat::setNumberPrefix(const QString & numberPrefix);
func (this *QTextListFormat) setNumberPrefix(args ...interface{}) () {
  // setNumberPrefix(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextListFormat15setNumberPrefixERK7QString
  default:
    qtrt.ErrorResolve("QTextListFormat", "setNumberPrefix", args)
  }

}

  // proto:  bool QTextFrameFormat::isValid();
func (this *QTextFrameFormat) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat7isValidEv
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "isValid", args)
  }

}

  // proto:  void QTextFrameFormat::setHeight(qreal height);
func (this *QTextFrameFormat) setHeight(args ...interface{}) () {
  // setHeight(qreal)
  // setHeight(const class QTextLength &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTextLength{}) // "const QTextLength &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextFrameFormat9setHeightEd
  case 1:
    // invoke: _ZN16QTextFrameFormat9setHeightERK11QTextLength
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setHeight", args)
  }

}

  // proto:  void QTextFrameFormat::setBorderBrush(const QBrush & brush);
func (this *QTextFrameFormat) setBorderBrush(args ...interface{}) () {
  // setBorderBrush(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextFrameFormat14setBorderBrushERK6QBrush
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setBorderBrush", args)
  }

}

  // proto:  qreal QTextFrameFormat::margin();
func (this *QTextFrameFormat) margin(args ...interface{}) () {
  // margin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat6marginEv
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "margin", args)
  }

}

  // proto:  QBrush QTextFrameFormat::borderBrush();
func (this *QTextFrameFormat) borderBrush(args ...interface{}) () {
  // borderBrush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat11borderBrushEv
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "borderBrush", args)
  }

}

  // proto:  void QTextFrameFormat::setRightMargin(qreal margin);
func (this *QTextFrameFormat) setRightMargin(args ...interface{}) () {
  // setRightMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextFrameFormat14setRightMarginEd
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setRightMargin", args)
  }

}

  // proto:  void QTextFrameFormat::setMargin(qreal margin);
func (this *QTextFrameFormat) setMargin(args ...interface{}) () {
  // setMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextFrameFormat9setMarginEd
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setMargin", args)
  }

}

  // proto:  void QTextFrameFormat::setBorder(qreal border);
func (this *QTextFrameFormat) setBorder(args ...interface{}) () {
  // setBorder(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextFrameFormat9setBorderEd
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setBorder", args)
  }

}

  // proto:  void QTextFrameFormat::setWidth(const QTextLength & length);
func (this *QTextFrameFormat) setWidth(args ...interface{}) () {
  // setWidth(const class QTextLength &)
  // setWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextLength{}) // "const QTextLength &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextFrameFormat8setWidthERK11QTextLength
  case 1:
    // invoke: _ZN16QTextFrameFormat8setWidthEd
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setWidth", args)
  }

}

  // proto:  qreal QTextFrameFormat::bottomMargin();
func (this *QTextFrameFormat) bottomMargin(args ...interface{}) () {
  // bottomMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat12bottomMarginEv
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "bottomMargin", args)
  }

}

  // proto:  void QTextFrameFormat::setBottomMargin(qreal margin);
func (this *QTextFrameFormat) setBottomMargin(args ...interface{}) () {
  // setBottomMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextFrameFormat15setBottomMarginEd
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setBottomMargin", args)
  }

}

  // proto:  QTextLength QTextFrameFormat::height();
func (this *QTextFrameFormat) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat6heightEv
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "height", args)
  }

}

  // proto:  qreal QTextFrameFormat::rightMargin();
func (this *QTextFrameFormat) rightMargin(args ...interface{}) () {
  // rightMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat11rightMarginEv
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "rightMargin", args)
  }

}

  // proto:  void QTextFrameFormat::setPadding(qreal padding);
func (this *QTextFrameFormat) setPadding(args ...interface{}) () {
  // setPadding(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextFrameFormat10setPaddingEd
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setPadding", args)
  }

}

  // proto:  void QTextFrameFormat::setTopMargin(qreal margin);
func (this *QTextFrameFormat) setTopMargin(args ...interface{}) () {
  // setTopMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextFrameFormat12setTopMarginEd
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setTopMargin", args)
  }

}

  // proto:  qreal QTextFrameFormat::topMargin();
func (this *QTextFrameFormat) topMargin(args ...interface{}) () {
  // topMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat9topMarginEv
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "topMargin", args)
  }

}

  // proto:  QTextLength QTextFrameFormat::width();
func (this *QTextFrameFormat) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat5widthEv
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "width", args)
  }

}

  // proto:  void QTextFrameFormat::QTextFrameFormat(const QTextFormat & fmt);
func NewQTextFrameFormat(args ...interface{}) QTextFrameFormat {
  return QTextFrameFormat{}
}

  // proto:  qreal QTextFrameFormat::padding();
func (this *QTextFrameFormat) padding(args ...interface{}) () {
  // padding()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat7paddingEv
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "padding", args)
  }

}

  // proto:  void QTextFrameFormat::setLeftMargin(qreal margin);
func (this *QTextFrameFormat) setLeftMargin(args ...interface{}) () {
  // setLeftMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextFrameFormat13setLeftMarginEd
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setLeftMargin", args)
  }

}

  // proto:  qreal QTextFrameFormat::border();
func (this *QTextFrameFormat) border(args ...interface{}) () {
  // border()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat6borderEv
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "border", args)
  }

}

  // proto:  qreal QTextFrameFormat::leftMargin();
func (this *QTextFrameFormat) leftMargin(args ...interface{}) () {
  // leftMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat10leftMarginEv
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "leftMargin", args)
  }

}

// <= body block end

