package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  qreal QTextLength::rawValue();
extern void C_ZNK11QTextLength8rawValueEv(void* qthis); // 2
  // proto:  qreal QTextLength::value(qreal maximumLength);
extern void C_ZNK11QTextLength5valueEd(void* qthis, double arg0); // 2
  // proto:  void QTextLength::QTextLength();
extern void C_ZN11QTextLengthC2Ev(void* qthis); // 1
  // proto:  QTextLength::Type QTextLength::type();
extern void C_ZNK11QTextLength4typeEv(void* qthis); // 2
  // proto:  void QTextImageFormat::setHeight(qreal height);
extern void C_ZN16QTextImageFormat9setHeightEd(void* qthis, double arg0); // 2
  // proto:  QString QTextImageFormat::name();
extern void C_ZNK16QTextImageFormat4nameEv(void* qthis); // 2
  // proto:  void QTextImageFormat::setName(const QString & name);
extern void C_ZN16QTextImageFormat7setNameERK7QString(void* qthis, void* arg0); // 2
  // proto:  bool QTextImageFormat::isValid();
extern void C_ZNK16QTextImageFormat7isValidEv(void* qthis); // 2
  // proto:  qreal QTextImageFormat::height();
extern void C_ZNK16QTextImageFormat6heightEv(void* qthis); // 2
  // proto:  qreal QTextImageFormat::width();
extern void C_ZNK16QTextImageFormat5widthEv(void* qthis); // 2
  // proto:  void QTextImageFormat::QTextImageFormat();
extern void C_ZN16QTextImageFormatC2Ev(void* qthis); // 3
  // proto:  void QTextImageFormat::setWidth(qreal width);
extern void C_ZN16QTextImageFormat8setWidthEd(void* qthis, double arg0); // 2
  // proto:  QTextBlockFormat QTextFormat::toBlockFormat();
extern void C_ZNK11QTextFormat13toBlockFormatEv(void* qthis); // 4
  // proto:  int QTextFormat::intProperty(int propertyId);
extern void C_ZNK11QTextFormat11intPropertyEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QTextFormat::isFrameFormat();
extern void C_ZNK11QTextFormat13isFrameFormatEv(void* qthis); // 2
  // proto:  QTextImageFormat QTextFormat::toImageFormat();
extern void C_ZNK11QTextFormat13toImageFormatEv(void* qthis); // 4
  // proto:  QColor QTextFormat::colorProperty(int propertyId);
extern void C_ZNK11QTextFormat13colorPropertyEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QTextFormat::stringProperty(int propertyId);
extern void C_ZNK11QTextFormat14stringPropertyEi(void* qthis, int32_t arg0); // 4
  // proto:  qreal QTextFormat::doubleProperty(int propertyId);
extern void C_ZNK11QTextFormat14doublePropertyEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextFormat::clearForeground();
extern void C_ZN11QTextFormat15clearForegroundEv(void* qthis); // 2
  // proto:  int QTextFormat::objectIndex();
extern void C_ZNK11QTextFormat11objectIndexEv(void* qthis); // 4
  // proto:  bool QTextFormat::isImageFormat();
extern void C_ZNK11QTextFormat13isImageFormatEv(void* qthis); // 2
  // proto:  int QTextFormat::objectType();
extern void C_ZNK11QTextFormat10objectTypeEv(void* qthis); // 2
  // proto:  bool QTextFormat::hasProperty(int propertyId);
extern void C_ZNK11QTextFormat11hasPropertyEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextFormat::setForeground(const QBrush & brush);
extern void C_ZN11QTextFormat13setForegroundERK6QBrush(void* qthis, void* arg0); // 2
  // proto:  bool QTextFormat::isCharFormat();
extern void C_ZNK11QTextFormat12isCharFormatEv(void* qthis); // 2
  // proto:  QTextTableCellFormat QTextFormat::toTableCellFormat();
extern void C_ZNK11QTextFormat17toTableCellFormatEv(void* qthis); // 4
  // proto:  void QTextFormat::setBackground(const QBrush & brush);
extern void C_ZN11QTextFormat13setBackgroundERK6QBrush(void* qthis, void* arg0); // 2
  // proto:  bool QTextFormat::isEmpty();
extern void C_ZNK11QTextFormat7isEmptyEv(void* qthis); // 2
  // proto:  void QTextFormat::swap(QTextFormat & other);
extern void C_ZN11QTextFormat4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QTextFormat::QTextFormat();
extern void C_ZN11QTextFormatC2Ev(void* qthis); // 3
  // proto:  void QTextFormat::QTextFormat(int type);
extern void C_ZN11QTextFormatC2Ei(void* qthis, int32_t arg0); // 3
  // proto:  void QTextFormat::QTextFormat(const QTextFormat & rhs);
extern void C_ZN11QTextFormatC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  bool QTextFormat::isBlockFormat();
extern void C_ZNK11QTextFormat13isBlockFormatEv(void* qthis); // 2
  // proto:  void QTextFormat::clearBackground();
extern void C_ZN11QTextFormat15clearBackgroundEv(void* qthis); // 2
  // proto:  int QTextFormat::type();
extern void C_ZNK11QTextFormat4typeEv(void* qthis); // 4
  // proto:  QBrush QTextFormat::foreground();
extern void C_ZNK11QTextFormat10foregroundEv(void* qthis); // 2
  // proto:  QTextCharFormat QTextFormat::toCharFormat();
extern void C_ZNK11QTextFormat12toCharFormatEv(void* qthis); // 4
  // proto:  void QTextFormat::~QTextFormat();
extern void C_ZN11QTextFormatD2Ev(void* qthis); // 4
  // proto:  bool QTextFormat::isValid();
extern void C_ZNK11QTextFormat7isValidEv(void* qthis); // 2
  // proto:  bool QTextFormat::boolProperty(int propertyId);
extern void C_ZNK11QTextFormat12boolPropertyEi(void* qthis, int32_t arg0); // 4
  // proto:  QTextLength QTextFormat::lengthProperty(int propertyId);
extern void C_ZNK11QTextFormat14lengthPropertyEi(void* qthis, int32_t arg0); // 4
  // proto:  QBrush QTextFormat::background();
extern void C_ZNK11QTextFormat10backgroundEv(void* qthis); // 2
  // proto:  void QTextFormat::setProperty(int propertyId, const QVariant & value);
extern void C_ZN11QTextFormat11setPropertyEiRK8QVariant(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QTextListFormat QTextFormat::toListFormat();
extern void C_ZNK11QTextFormat12toListFormatEv(void* qthis); // 4
  // proto:  QMap<int, QVariant> QTextFormat::properties();
extern void C_ZNK11QTextFormat10propertiesEv(void* qthis); // 4
  // proto:  QTextFrameFormat QTextFormat::toFrameFormat();
extern void C_ZNK11QTextFormat13toFrameFormatEv(void* qthis); // 4
  // proto:  QPen QTextFormat::penProperty(int propertyId);
extern void C_ZNK11QTextFormat11penPropertyEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QTextFormat::isTableCellFormat();
extern void C_ZNK11QTextFormat17isTableCellFormatEv(void* qthis); // 2
  // proto:  void QTextFormat::setObjectIndex(int object);
extern void C_ZN11QTextFormat14setObjectIndexEi(void* qthis, int32_t arg0); // 4
  // proto:  Qt::LayoutDirection QTextFormat::layoutDirection();
extern void C_ZNK11QTextFormat15layoutDirectionEv(void* qthis); // 2
  // proto:  void QTextFormat::clearProperty(int propertyId);
extern void C_ZN11QTextFormat13clearPropertyEi(void* qthis, int32_t arg0); // 4
  // proto:  QBrush QTextFormat::brushProperty(int propertyId);
extern void C_ZNK11QTextFormat13brushPropertyEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTextFormat::propertyCount();
extern void C_ZNK11QTextFormat13propertyCountEv(void* qthis); // 4
  // proto:  void QTextFormat::merge(const QTextFormat & other);
extern void C_ZN11QTextFormat5mergeERKS_(void* qthis, void* arg0); // 4
  // proto:  QVector<QTextLength> QTextFormat::lengthVectorProperty(int propertyId);
extern void C_ZNK11QTextFormat20lengthVectorPropertyEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QTextFormat::isTableFormat();
extern void C_ZNK11QTextFormat13isTableFormatEv(void* qthis); // 2
  // proto:  QTextTableFormat QTextFormat::toTableFormat();
extern void C_ZNK11QTextFormat13toTableFormatEv(void* qthis); // 4
  // proto:  QVariant QTextFormat::property(int propertyId);
extern void C_ZNK11QTextFormat8propertyEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QTextFormat::isListFormat();
extern void C_ZNK11QTextFormat12isListFormatEv(void* qthis); // 2
  // proto:  void QTextFormat::setObjectType(int type);
extern void C_ZN11QTextFormat13setObjectTypeEi(void* qthis, int32_t arg0); // 2
  // proto:  void QTextBlockFormat::setTextIndent(qreal aindent);
extern void C_ZN16QTextBlockFormat13setTextIndentEd(void* qthis, double arg0); // 2
  // proto:  int QTextBlockFormat::lineHeightType();
extern void C_ZNK16QTextBlockFormat14lineHeightTypeEv(void* qthis); // 2
  // proto:  Qt::Alignment QTextBlockFormat::alignment();
extern void C_ZNK16QTextBlockFormat9alignmentEv(void* qthis); // 2
  // proto:  qreal QTextBlockFormat::leftMargin();
extern void C_ZNK16QTextBlockFormat10leftMarginEv(void* qthis); // 2
  // proto:  qreal QTextBlockFormat::topMargin();
extern void C_ZNK16QTextBlockFormat9topMarginEv(void* qthis); // 2
  // proto:  void QTextBlockFormat::setLeftMargin(qreal margin);
extern void C_ZN16QTextBlockFormat13setLeftMarginEd(void* qthis, double arg0); // 2
  // proto:  qreal QTextBlockFormat::lineHeight();
extern void C_ZNK16QTextBlockFormat10lineHeightEv(void* qthis); // 2
  // proto:  qreal QTextBlockFormat::lineHeight(qreal scriptLineHeight, qreal scaling);
extern void C_ZNK16QTextBlockFormat10lineHeightEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QTextBlockFormat::setTopMargin(qreal margin);
extern void C_ZN16QTextBlockFormat12setTopMarginEd(void* qthis, double arg0); // 2
  // proto:  qreal QTextBlockFormat::rightMargin();
extern void C_ZNK16QTextBlockFormat11rightMarginEv(void* qthis); // 2
  // proto:  bool QTextBlockFormat::isValid();
extern void C_ZNK16QTextBlockFormat7isValidEv(void* qthis); // 2
  // proto:  void QTextBlockFormat::setIndent(int indent);
extern void C_ZN16QTextBlockFormat9setIndentEi(void* qthis, int32_t arg0); // 2
  // proto:  PageBreakFlags QTextBlockFormat::pageBreakPolicy();
extern void C_ZNK16QTextBlockFormat15pageBreakPolicyEv(void* qthis); // 2
  // proto:  qreal QTextBlockFormat::bottomMargin();
extern void C_ZNK16QTextBlockFormat12bottomMarginEv(void* qthis); // 2
  // proto:  void QTextBlockFormat::setLineHeight(qreal height, int heightType);
extern void C_ZN16QTextBlockFormat13setLineHeightEdi(void* qthis, double arg0, int32_t arg1); // 2
  // proto:  void QTextBlockFormat::setNonBreakableLines(bool b);
extern void C_ZN16QTextBlockFormat20setNonBreakableLinesEb(void* qthis, bool arg0); // 2
  // proto:  void QTextBlockFormat::setRightMargin(qreal margin);
extern void C_ZN16QTextBlockFormat14setRightMarginEd(void* qthis, double arg0); // 2
  // proto:  int QTextBlockFormat::indent();
extern void C_ZNK16QTextBlockFormat6indentEv(void* qthis); // 2
  // proto:  QList<QTextOption::Tab> QTextBlockFormat::tabPositions();
extern void C_ZNK16QTextBlockFormat12tabPositionsEv(void* qthis); // 4
  // proto:  qreal QTextBlockFormat::textIndent();
extern void C_ZNK16QTextBlockFormat10textIndentEv(void* qthis); // 2
  // proto:  void QTextBlockFormat::setBottomMargin(qreal margin);
extern void C_ZN16QTextBlockFormat15setBottomMarginEd(void* qthis, double arg0); // 2
  // proto:  bool QTextBlockFormat::nonBreakableLines();
extern void C_ZNK16QTextBlockFormat17nonBreakableLinesEv(void* qthis); // 2
  // proto:  void QTextBlockFormat::QTextBlockFormat();
extern void C_ZN16QTextBlockFormatC2Ev(void* qthis); // 3
  // proto:  void QTextCharFormat::setTableCellColumnSpan(int tableCellColumnSpan);
extern void C_ZN15QTextCharFormat22setTableCellColumnSpanEi(void* qthis, int32_t arg0); // 2
  // proto:  void QTextCharFormat::setAnchorHref(const QString & value);
extern void C_ZN15QTextCharFormat13setAnchorHrefERK7QString(void* qthis, void* arg0); // 2
  // proto:  void QTextCharFormat::setFontOverline(bool overline);
extern void C_ZN15QTextCharFormat15setFontOverlineEb(void* qthis, bool arg0); // 2
  // proto:  void QTextCharFormat::setTextOutline(const QPen & pen);
extern void C_ZN15QTextCharFormat14setTextOutlineERK4QPen(void* qthis, void* arg0); // 2
  // proto:  QString QTextCharFormat::fontFamily();
extern void C_ZNK15QTextCharFormat10fontFamilyEv(void* qthis); // 2
  // proto:  QStringList QTextCharFormat::anchorNames();
extern void C_ZNK15QTextCharFormat11anchorNamesEv(void* qthis); // 4
  // proto:  QFont::StyleHint QTextCharFormat::fontStyleHint();
extern void C_ZNK15QTextCharFormat13fontStyleHintEv(void* qthis); // 2
  // proto:  void QTextCharFormat::setFontStretch(int factor);
extern void C_ZN15QTextCharFormat14setFontStretchEi(void* qthis, int32_t arg0); // 2
  // proto:  QFont QTextCharFormat::font();
extern void C_ZNK15QTextCharFormat4fontEv(void* qthis); // 4
  // proto:  bool QTextCharFormat::fontItalic();
extern void C_ZNK15QTextCharFormat10fontItalicEv(void* qthis); // 2
  // proto:  bool QTextCharFormat::fontKerning();
extern void C_ZNK15QTextCharFormat11fontKerningEv(void* qthis); // 2
  // proto:  qreal QTextCharFormat::fontWordSpacing();
extern void C_ZNK15QTextCharFormat15fontWordSpacingEv(void* qthis); // 2
  // proto:  bool QTextCharFormat::fontOverline();
extern void C_ZNK15QTextCharFormat12fontOverlineEv(void* qthis); // 2
  // proto:  void QTextCharFormat::setFontUnderline(bool underline);
extern void C_ZN15QTextCharFormat16setFontUnderlineEb(void* qthis, bool arg0); // 2
  // proto:  QTextCharFormat::UnderlineStyle QTextCharFormat::underlineStyle();
extern void C_ZNK15QTextCharFormat14underlineStyleEv(void* qthis); // 2
  // proto:  QFont::StyleStrategy QTextCharFormat::fontStyleStrategy();
extern void C_ZNK15QTextCharFormat17fontStyleStrategyEv(void* qthis); // 2
  // proto:  void QTextCharFormat::setFontWordSpacing(qreal spacing);
extern void C_ZN15QTextCharFormat18setFontWordSpacingEd(void* qthis, double arg0); // 2
  // proto:  void QTextCharFormat::setFontWeight(int weight);
extern void C_ZN15QTextCharFormat13setFontWeightEi(void* qthis, int32_t arg0); // 2
  // proto:  QTextCharFormat::VerticalAlignment QTextCharFormat::verticalAlignment();
extern void C_ZNK15QTextCharFormat17verticalAlignmentEv(void* qthis); // 2
  // proto:  void QTextCharFormat::setFontFamily(const QString & family);
extern void C_ZN15QTextCharFormat13setFontFamilyERK7QString(void* qthis, void* arg0); // 2
  // proto:  void QTextCharFormat::setAnchorNames(const QStringList & names);
extern void C_ZN15QTextCharFormat14setAnchorNamesERK11QStringList(void* qthis, void* arg0); // 2
  // proto:  void QTextCharFormat::setTableCellRowSpan(int tableCellRowSpan);
extern void C_ZN15QTextCharFormat19setTableCellRowSpanEi(void* qthis, int32_t arg0); // 2
  // proto:  void QTextCharFormat::setFontKerning(bool enable);
extern void C_ZN15QTextCharFormat14setFontKerningEb(void* qthis, bool arg0); // 2
  // proto:  void QTextCharFormat::setFontLetterSpacing(qreal spacing);
extern void C_ZN15QTextCharFormat20setFontLetterSpacingEd(void* qthis, double arg0); // 2
  // proto:  QFont::SpacingType QTextCharFormat::fontLetterSpacingType();
extern void C_ZNK15QTextCharFormat21fontLetterSpacingTypeEv(void* qthis); // 2
  // proto:  bool QTextCharFormat::fontUnderline();
extern void C_ZNK15QTextCharFormat13fontUnderlineEv(void* qthis); // 4
  // proto:  QFont::HintingPreference QTextCharFormat::fontHintingPreference();
extern void C_ZNK15QTextCharFormat21fontHintingPreferenceEv(void* qthis); // 2
  // proto:  void QTextCharFormat::setFontStrikeOut(bool strikeOut);
extern void C_ZN15QTextCharFormat16setFontStrikeOutEb(void* qthis, bool arg0); // 2
  // proto:  QString QTextCharFormat::anchorName();
extern void C_ZNK15QTextCharFormat10anchorNameEv(void* qthis); // 4
  // proto:  qreal QTextCharFormat::fontPointSize();
extern void C_ZNK15QTextCharFormat13fontPointSizeEv(void* qthis); // 2
  // proto:  bool QTextCharFormat::fontFixedPitch();
extern void C_ZNK15QTextCharFormat14fontFixedPitchEv(void* qthis); // 2
  // proto:  bool QTextCharFormat::isValid();
extern void C_ZNK15QTextCharFormat7isValidEv(void* qthis); // 2
  // proto:  QString QTextCharFormat::toolTip();
extern void C_ZNK15QTextCharFormat7toolTipEv(void* qthis); // 2
  // proto:  void QTextCharFormat::setFontPointSize(qreal size);
extern void C_ZN15QTextCharFormat16setFontPointSizeEd(void* qthis, double arg0); // 2
  // proto:  void QTextCharFormat::setAnchor(bool anchor);
extern void C_ZN15QTextCharFormat9setAnchorEb(void* qthis, bool arg0); // 2
  // proto:  void QTextCharFormat::setAnchorName(const QString & name);
extern void C_ZN15QTextCharFormat13setAnchorNameERK7QString(void* qthis, void* arg0); // 2
  // proto:  qreal QTextCharFormat::fontLetterSpacing();
extern void C_ZNK15QTextCharFormat17fontLetterSpacingEv(void* qthis); // 2
  // proto:  bool QTextCharFormat::isAnchor();
extern void C_ZNK15QTextCharFormat8isAnchorEv(void* qthis); // 2
  // proto:  void QTextCharFormat::setFontItalic(bool italic);
extern void C_ZN15QTextCharFormat13setFontItalicEb(void* qthis, bool arg0); // 2
  // proto:  void QTextCharFormat::QTextCharFormat();
extern void C_ZN15QTextCharFormatC2Ev(void* qthis); // 3
  // proto:  QColor QTextCharFormat::underlineColor();
extern void C_ZNK15QTextCharFormat14underlineColorEv(void* qthis); // 2
  // proto:  void QTextCharFormat::setUnderlineColor(const QColor & color);
extern void C_ZN15QTextCharFormat17setUnderlineColorERK6QColor(void* qthis, void* arg0); // 2
  // proto:  int QTextCharFormat::fontWeight();
extern void C_ZNK15QTextCharFormat10fontWeightEv(void* qthis); // 2
  // proto:  void QTextCharFormat::setToolTip(const QString & tip);
extern void C_ZN15QTextCharFormat10setToolTipERK7QString(void* qthis, void* arg0); // 2
  // proto:  int QTextCharFormat::tableCellColumnSpan();
extern void C_ZNK15QTextCharFormat19tableCellColumnSpanEv(void* qthis); // 2
  // proto:  QString QTextCharFormat::anchorHref();
extern void C_ZNK15QTextCharFormat10anchorHrefEv(void* qthis); // 2
  // proto:  QFont::Capitalization QTextCharFormat::fontCapitalization();
extern void C_ZNK15QTextCharFormat18fontCapitalizationEv(void* qthis); // 2
  // proto:  QPen QTextCharFormat::textOutline();
extern void C_ZNK15QTextCharFormat11textOutlineEv(void* qthis); // 2
  // proto:  int QTextCharFormat::tableCellRowSpan();
extern void C_ZNK15QTextCharFormat16tableCellRowSpanEv(void* qthis); // 2
  // proto:  void QTextCharFormat::setFontFixedPitch(bool fixedPitch);
extern void C_ZN15QTextCharFormat17setFontFixedPitchEb(void* qthis, bool arg0); // 2
  // proto:  void QTextCharFormat::setFont(const QFont & font);
extern void C_ZN15QTextCharFormat7setFontERK5QFont(void* qthis, void* arg0); // 4
  // proto:  int QTextCharFormat::fontStretch();
extern void C_ZNK15QTextCharFormat11fontStretchEv(void* qthis); // 2
  // proto:  bool QTextCharFormat::fontStrikeOut();
extern void C_ZNK15QTextCharFormat13fontStrikeOutEv(void* qthis); // 2
  // proto:  void QTextTableFormat::clearColumnWidthConstraints();
extern void C_ZN16QTextTableFormat27clearColumnWidthConstraintsEv(void* qthis); // 2
  // proto:  QVector<QTextLength> QTextTableFormat::columnWidthConstraints();
extern void C_ZNK16QTextTableFormat22columnWidthConstraintsEv(void* qthis); // 2
  // proto:  void QTextTableFormat::QTextTableFormat();
extern void C_ZN16QTextTableFormatC2Ev(void* qthis); // 3
  // proto:  void QTextTableFormat::setHeaderRowCount(int count);
extern void C_ZN16QTextTableFormat17setHeaderRowCountEi(void* qthis, int32_t arg0); // 2
  // proto:  int QTextTableFormat::headerRowCount();
extern void C_ZNK16QTextTableFormat14headerRowCountEv(void* qthis); // 2
  // proto:  void QTextTableFormat::setColumns(int columns);
extern void C_ZN16QTextTableFormat10setColumnsEi(void* qthis, int32_t arg0); // 2
  // proto:  Qt::Alignment QTextTableFormat::alignment();
extern void C_ZNK16QTextTableFormat9alignmentEv(void* qthis); // 2
  // proto:  bool QTextTableFormat::isValid();
extern void C_ZNK16QTextTableFormat7isValidEv(void* qthis); // 2
  // proto:  qreal QTextTableFormat::cellPadding();
extern void C_ZNK16QTextTableFormat11cellPaddingEv(void* qthis); // 2
  // proto:  void QTextTableFormat::setCellSpacing(qreal spacing);
extern void C_ZN16QTextTableFormat14setCellSpacingEd(void* qthis, double arg0); // 2
  // proto:  qreal QTextTableFormat::cellSpacing();
extern void C_ZNK16QTextTableFormat11cellSpacingEv(void* qthis); // 2
  // proto:  void QTextTableFormat::setCellPadding(qreal padding);
extern void C_ZN16QTextTableFormat14setCellPaddingEd(void* qthis, double arg0); // 2
  // proto:  int QTextTableFormat::columns();
extern void C_ZNK16QTextTableFormat7columnsEv(void* qthis); // 2
  // proto:  void QTextTableCellFormat::setLeftPadding(qreal padding);
extern void C_ZN20QTextTableCellFormat14setLeftPaddingEd(void* qthis, double arg0); // 2
  // proto:  qreal QTextTableCellFormat::leftPadding();
extern void C_ZNK20QTextTableCellFormat11leftPaddingEv(void* qthis); // 2
  // proto:  void QTextTableCellFormat::setBottomPadding(qreal padding);
extern void C_ZN20QTextTableCellFormat16setBottomPaddingEd(void* qthis, double arg0); // 2
  // proto:  void QTextTableCellFormat::QTextTableCellFormat();
extern void C_ZN20QTextTableCellFormatC2Ev(void* qthis); // 3
  // proto:  void QTextTableCellFormat::setTopPadding(qreal padding);
extern void C_ZN20QTextTableCellFormat13setTopPaddingEd(void* qthis, double arg0); // 2
  // proto:  qreal QTextTableCellFormat::rightPadding();
extern void C_ZNK20QTextTableCellFormat12rightPaddingEv(void* qthis); // 2
  // proto:  void QTextTableCellFormat::setRightPadding(qreal padding);
extern void C_ZN20QTextTableCellFormat15setRightPaddingEd(void* qthis, double arg0); // 2
  // proto:  bool QTextTableCellFormat::isValid();
extern void C_ZNK20QTextTableCellFormat7isValidEv(void* qthis); // 2
  // proto:  qreal QTextTableCellFormat::topPadding();
extern void C_ZNK20QTextTableCellFormat10topPaddingEv(void* qthis); // 2
  // proto:  void QTextTableCellFormat::setPadding(qreal padding);
extern void C_ZN20QTextTableCellFormat10setPaddingEd(void* qthis, double arg0); // 2
  // proto:  qreal QTextTableCellFormat::bottomPadding();
extern void C_ZNK20QTextTableCellFormat13bottomPaddingEv(void* qthis); // 2
  // proto:  QTextListFormat::Style QTextListFormat::style();
extern void C_ZNK15QTextListFormat5styleEv(void* qthis); // 2
  // proto:  QString QTextListFormat::numberPrefix();
extern void C_ZNK15QTextListFormat12numberPrefixEv(void* qthis); // 2
  // proto:  int QTextListFormat::indent();
extern void C_ZNK15QTextListFormat6indentEv(void* qthis); // 2
  // proto:  void QTextListFormat::setNumberPrefix(const QString & numberPrefix);
extern void C_ZN15QTextListFormat15setNumberPrefixERK7QString(void* qthis, void* arg0); // 2
  // proto:  bool QTextListFormat::isValid();
extern void C_ZNK15QTextListFormat7isValidEv(void* qthis); // 2
  // proto:  void QTextListFormat::QTextListFormat();
extern void C_ZN15QTextListFormatC2Ev(void* qthis); // 3
  // proto:  void QTextListFormat::setIndent(int indent);
extern void C_ZN15QTextListFormat9setIndentEi(void* qthis, int32_t arg0); // 2
  // proto:  QString QTextListFormat::numberSuffix();
extern void C_ZNK15QTextListFormat12numberSuffixEv(void* qthis); // 2
  // proto:  void QTextListFormat::setNumberSuffix(const QString & numberSuffix);
extern void C_ZN15QTextListFormat15setNumberSuffixERK7QString(void* qthis, void* arg0); // 2
  // proto:  void QTextFrameFormat::setBorderBrush(const QBrush & brush);
extern void C_ZN16QTextFrameFormat14setBorderBrushERK6QBrush(void* qthis, void* arg0); // 2
  // proto:  QTextLength QTextFrameFormat::height();
extern void C_ZNK16QTextFrameFormat6heightEv(void* qthis); // 2
  // proto:  void QTextFrameFormat::setWidth(qreal width);
extern void C_ZN16QTextFrameFormat8setWidthEd(void* qthis, double arg0); // 2
  // proto:  void QTextFrameFormat::setWidth(const QTextLength & length);
extern void C_ZN16QTextFrameFormat8setWidthERK11QTextLength(void* qthis, void* arg0); // 2
  // proto:  qreal QTextFrameFormat::topMargin();
extern void C_ZNK16QTextFrameFormat9topMarginEv(void* qthis); // 4
  // proto:  qreal QTextFrameFormat::border();
extern void C_ZNK16QTextFrameFormat6borderEv(void* qthis); // 2
  // proto:  QTextLength QTextFrameFormat::width();
extern void C_ZNK16QTextFrameFormat5widthEv(void* qthis); // 2
  // proto:  void QTextFrameFormat::setLeftMargin(qreal margin);
extern void C_ZN16QTextFrameFormat13setLeftMarginEd(void* qthis, double arg0); // 2
  // proto:  qreal QTextFrameFormat::leftMargin();
extern void C_ZNK16QTextFrameFormat10leftMarginEv(void* qthis); // 4
  // proto:  void QTextFrameFormat::setPadding(qreal padding);
extern void C_ZN16QTextFrameFormat10setPaddingEd(void* qthis, double arg0); // 2
  // proto:  void QTextFrameFormat::setTopMargin(qreal margin);
extern void C_ZN16QTextFrameFormat12setTopMarginEd(void* qthis, double arg0); // 2
  // proto:  qreal QTextFrameFormat::rightMargin();
extern void C_ZNK16QTextFrameFormat11rightMarginEv(void* qthis); // 4
  // proto:  bool QTextFrameFormat::isValid();
extern void C_ZNK16QTextFrameFormat7isValidEv(void* qthis); // 2
  // proto:  PageBreakFlags QTextFrameFormat::pageBreakPolicy();
extern void C_ZNK16QTextFrameFormat15pageBreakPolicyEv(void* qthis); // 2
  // proto:  qreal QTextFrameFormat::padding();
extern void C_ZNK16QTextFrameFormat7paddingEv(void* qthis); // 2
  // proto:  void QTextFrameFormat::setMargin(qreal margin);
extern void C_ZN16QTextFrameFormat9setMarginEd(void* qthis, double arg0); // 4
  // proto:  qreal QTextFrameFormat::bottomMargin();
extern void C_ZNK16QTextFrameFormat12bottomMarginEv(void* qthis); // 4
  // proto:  void QTextFrameFormat::QTextFrameFormat();
extern void C_ZN16QTextFrameFormatC2Ev(void* qthis); // 3
  // proto:  void QTextFrameFormat::setHeight(qreal height);
extern void C_ZN16QTextFrameFormat9setHeightEd(void* qthis, double arg0); // 2
  // proto:  void QTextFrameFormat::setHeight(const QTextLength & height);
extern void C_ZN16QTextFrameFormat9setHeightERK11QTextLength(void* qthis, void* arg0); // 2
  // proto:  QBrush QTextFrameFormat::borderBrush();
extern void C_ZNK16QTextFrameFormat11borderBrushEv(void* qthis); // 2
  // proto:  void QTextFrameFormat::setRightMargin(qreal margin);
extern void C_ZN16QTextFrameFormat14setRightMarginEd(void* qthis, double arg0); // 2
  // proto:  void QTextFrameFormat::setBorder(qreal border);
extern void C_ZN16QTextFrameFormat9setBorderEd(void* qthis, double arg0); // 2
  // proto:  void QTextFrameFormat::setBottomMargin(qreal margin);
extern void C_ZN16QTextFrameFormat15setBottomMarginEd(void* qthis, double arg0); // 2
  // proto:  QTextFrameFormat::Position QTextFrameFormat::position();
extern void C_ZNK16QTextFrameFormat8positionEv(void* qthis); // 2
  // proto:  QTextFrameFormat::BorderStyle QTextFrameFormat::borderStyle();
extern void C_ZNK16QTextFrameFormat11borderStyleEv(void* qthis); // 2
  // proto:  qreal QTextFrameFormat::margin();
extern void C_ZNK16QTextFrameFormat6marginEv(void* qthis); // 2
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextImageFormat)=1
type QTextImageFormat struct {
  /*qbase*/ QTextCharFormat;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextFormat)=1
type QTextFormat struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextBlockFormat)=1
type QTextBlockFormat struct {
  /*qbase*/ QTextFormat;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextCharFormat)=1
type QTextCharFormat struct {
  /*qbase*/ QTextFormat;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextTableFormat)=1
type QTextTableFormat struct {
  /*qbase*/ QTextFrameFormat;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextTableCellFormat)=1
type QTextTableCellFormat struct {
  /*qbase*/ QTextCharFormat;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextListFormat)=1
type QTextListFormat struct {
  /*qbase*/ QTextFormat;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextFrameFormat)=1
type QTextFrameFormat struct {
  /*qbase*/ QTextFormat;
  qclsinst unsafe.Pointer /* *C.void */;
}

// rawValue()
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
    // invoke: qreal rawValue()
    C.C_ZNK11QTextLength8rawValueEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLength", "rawValue", args)
  }

}

// value(qreal)
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
    // invoke: qreal value(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QTextLength5valueEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLength", "value", args)
  }

}

// QTextLength()
func NewQTextLength(args ...interface{}) QTextLength {
  // QTextLength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextLengthC1Ev
    // invoke: void QTextLength()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QTextLengthC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QTextLength", "QTextLength", args)
  }

  return QTextLength{}
}

// type()
func (this *QTextLength) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLength4typeEv
    // invoke: QTextLength::Type type()
    C.C_ZNK11QTextLength4typeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLength", "type", args)
  }

}

// setHeight(qreal)
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
    // invoke: void setHeight(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextImageFormat9setHeightEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextImageFormat", "setHeight", args)
  }

}

// name()
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
    // invoke: QString name()
    C.C_ZNK16QTextImageFormat4nameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextImageFormat", "name", args)
  }

}

// setName(const class QString &)
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
    // invoke: void setName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextImageFormat7setNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextImageFormat", "setName", args)
  }

}

// isValid()
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
    // invoke: bool isValid()
    C.C_ZNK16QTextImageFormat7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextImageFormat", "isValid", args)
  }

}

// height()
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
    // invoke: qreal height()
    C.C_ZNK16QTextImageFormat6heightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextImageFormat", "height", args)
  }

}

// width()
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
    // invoke: qreal width()
    C.C_ZNK16QTextImageFormat5widthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextImageFormat", "width", args)
  }

}

// QTextImageFormat()
func NewQTextImageFormat(args ...interface{}) QTextImageFormat {
  // QTextImageFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextImageFormatC1Ev
    // invoke: void QTextImageFormat()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN16QTextImageFormatC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QTextImageFormat", "QTextImageFormat", args)
  }

  return QTextImageFormat{}
}

// setWidth(qreal)
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
    // invoke: void setWidth(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextImageFormat8setWidthEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextImageFormat", "setWidth", args)
  }

}

// toBlockFormat()
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
    // invoke: QTextBlockFormat toBlockFormat()
    C.C_ZNK11QTextFormat13toBlockFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "toBlockFormat", args)
  }

}

// intProperty(int)
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
    // invoke: int intProperty(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QTextFormat11intPropertyEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "intProperty", args)
  }

}

// isFrameFormat()
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
    // invoke: bool isFrameFormat()
    C.C_ZNK11QTextFormat13isFrameFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "isFrameFormat", args)
  }

}

// toImageFormat()
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
    // invoke: QTextImageFormat toImageFormat()
    C.C_ZNK11QTextFormat13toImageFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "toImageFormat", args)
  }

}

// colorProperty(int)
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
    // invoke: QColor colorProperty(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QTextFormat13colorPropertyEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "colorProperty", args)
  }

}

// stringProperty(int)
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
    // invoke: QString stringProperty(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QTextFormat14stringPropertyEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "stringProperty", args)
  }

}

// doubleProperty(int)
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
    // invoke: qreal doubleProperty(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QTextFormat14doublePropertyEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "doubleProperty", args)
  }

}

// clearForeground()
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
    // invoke: void clearForeground()
    C.C_ZN11QTextFormat15clearForegroundEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "clearForeground", args)
  }

}

// objectIndex()
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
    // invoke: int objectIndex()
    C.C_ZNK11QTextFormat11objectIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "objectIndex", args)
  }

}

// isImageFormat()
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
    // invoke: bool isImageFormat()
    C.C_ZNK11QTextFormat13isImageFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "isImageFormat", args)
  }

}

// objectType()
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
    // invoke: int objectType()
    C.C_ZNK11QTextFormat10objectTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "objectType", args)
  }

}

// hasProperty(int)
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
    // invoke: bool hasProperty(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QTextFormat11hasPropertyEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "hasProperty", args)
  }

}

// setForeground(const class QBrush &)
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
    // invoke: void setForeground(const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextFormat13setForegroundERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "setForeground", args)
  }

}

// isCharFormat()
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
    // invoke: bool isCharFormat()
    C.C_ZNK11QTextFormat12isCharFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "isCharFormat", args)
  }

}

// toTableCellFormat()
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
    // invoke: QTextTableCellFormat toTableCellFormat()
    C.C_ZNK11QTextFormat17toTableCellFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "toTableCellFormat", args)
  }

}

// setBackground(const class QBrush &)
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
    // invoke: void setBackground(const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextFormat13setBackgroundERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "setBackground", args)
  }

}

// isEmpty()
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
    // invoke: bool isEmpty()
    C.C_ZNK11QTextFormat7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "isEmpty", args)
  }

}

// swap(class QTextFormat &)
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
    // invoke: void swap(class QTextFormat &)
    var arg0 = args[0].(QTextFormat).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextFormat4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "swap", args)
  }

}

// QTextFormat()
func NewQTextFormat(args ...interface{}) QTextFormat {
  // QTextFormat()
  // QTextFormat(int)
  // QTextFormat(const class QTextFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QTextFormat{}) // "const QTextFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormatC1Ev
    // invoke: void QTextFormat()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QTextFormatC2Ev(qthis)
  case 1:
    // invoke: _ZN11QTextFormatC1Ei
    // invoke: void QTextFormat(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QTextFormatC2Ei(qthis, arg0)
  case 2:
    // invoke: _ZN11QTextFormatC1ERKS_
    // invoke: void QTextFormat(const class QTextFormat &)
    var arg0 = args[0].(QTextFormat).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QTextFormatC2ERKS_(qthis, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "QTextFormat", args)
  }

  return QTextFormat{}
}

// isBlockFormat()
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
    // invoke: bool isBlockFormat()
    C.C_ZNK11QTextFormat13isBlockFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "isBlockFormat", args)
  }

}

// clearBackground()
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
    // invoke: void clearBackground()
    C.C_ZN11QTextFormat15clearBackgroundEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "clearBackground", args)
  }

}

// type()
func (this *QTextFormat) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat4typeEv
    // invoke: int type()
    C.C_ZNK11QTextFormat4typeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "type", args)
  }

}

// foreground()
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
    // invoke: QBrush foreground()
    C.C_ZNK11QTextFormat10foregroundEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "foreground", args)
  }

}

// toCharFormat()
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
    // invoke: QTextCharFormat toCharFormat()
    C.C_ZNK11QTextFormat12toCharFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "toCharFormat", args)
  }

}

// ~QTextFormat()
func (this *QTextFormat) FreeQTextFormat(args ...interface{}) () {
  // ~QTextFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormatD0Ev
    // invoke: void ~QTextFormat()
    C.C_ZN11QTextFormatD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "~QTextFormat", args)
  }

}

// isValid()
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
    // invoke: bool isValid()
    C.C_ZNK11QTextFormat7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "isValid", args)
  }

}

// boolProperty(int)
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
    // invoke: bool boolProperty(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QTextFormat12boolPropertyEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "boolProperty", args)
  }

}

// lengthProperty(int)
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
    // invoke: QTextLength lengthProperty(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QTextFormat14lengthPropertyEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "lengthProperty", args)
  }

}

// background()
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
    // invoke: QBrush background()
    C.C_ZNK11QTextFormat10backgroundEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "background", args)
  }

}

// setProperty(int, const class QVariant &)
func (this *QTextFormat) setProperty(args ...interface{}) () {
  // setProperty(int, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat11setPropertyEiRK8QVariant
    // invoke: void setProperty(int, const class QVariant &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QTextFormat11setPropertyEiRK8QVariant(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextFormat", "setProperty", args)
  }

}

// toListFormat()
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
    // invoke: QTextListFormat toListFormat()
    C.C_ZNK11QTextFormat12toListFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "toListFormat", args)
  }

}

// properties()
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
    // invoke: QMap<int, QVariant> properties()
    C.C_ZNK11QTextFormat10propertiesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "properties", args)
  }

}

// toFrameFormat()
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
    // invoke: QTextFrameFormat toFrameFormat()
    C.C_ZNK11QTextFormat13toFrameFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "toFrameFormat", args)
  }

}

// penProperty(int)
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
    // invoke: QPen penProperty(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QTextFormat11penPropertyEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "penProperty", args)
  }

}

// isTableCellFormat()
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
    // invoke: bool isTableCellFormat()
    C.C_ZNK11QTextFormat17isTableCellFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "isTableCellFormat", args)
  }

}

// setObjectIndex(int)
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
    // invoke: void setObjectIndex(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextFormat14setObjectIndexEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "setObjectIndex", args)
  }

}

// layoutDirection()
func (this *QTextFormat) layoutDirection(args ...interface{}) () {
  // layoutDirection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat15layoutDirectionEv
    // invoke: Qt::LayoutDirection layoutDirection()
    C.C_ZNK11QTextFormat15layoutDirectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "layoutDirection", args)
  }

}

// clearProperty(int)
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
    // invoke: void clearProperty(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextFormat13clearPropertyEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "clearProperty", args)
  }

}

// brushProperty(int)
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
    // invoke: QBrush brushProperty(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QTextFormat13brushPropertyEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "brushProperty", args)
  }

}

// propertyCount()
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
    // invoke: int propertyCount()
    C.C_ZNK11QTextFormat13propertyCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "propertyCount", args)
  }

}

// merge(const class QTextFormat &)
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
    // invoke: void merge(const class QTextFormat &)
    var arg0 = args[0].(QTextFormat).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextFormat5mergeERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "merge", args)
  }

}

// lengthVectorProperty(int)
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
    // invoke: QVector<QTextLength> lengthVectorProperty(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QTextFormat20lengthVectorPropertyEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "lengthVectorProperty", args)
  }

}

// isTableFormat()
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
    // invoke: bool isTableFormat()
    C.C_ZNK11QTextFormat13isTableFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "isTableFormat", args)
  }

}

// toTableFormat()
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
    // invoke: QTextTableFormat toTableFormat()
    C.C_ZNK11QTextFormat13toTableFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "toTableFormat", args)
  }

}

// property(int)
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
    // invoke: QVariant property(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QTextFormat8propertyEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "property", args)
  }

}

// isListFormat()
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
    // invoke: bool isListFormat()
    C.C_ZNK11QTextFormat12isListFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "isListFormat", args)
  }

}

// setObjectType(int)
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
    // invoke: void setObjectType(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextFormat13setObjectTypeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "setObjectType", args)
  }

}

// setTextIndent(qreal)
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
    // invoke: void setTextIndent(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextBlockFormat13setTextIndentEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setTextIndent", args)
  }

}

// lineHeightType()
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
    // invoke: int lineHeightType()
    C.C_ZNK16QTextBlockFormat14lineHeightTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "lineHeightType", args)
  }

}

// alignment()
func (this *QTextBlockFormat) alignment(args ...interface{}) () {
  // alignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat9alignmentEv
    // invoke: Qt::Alignment alignment()
    C.C_ZNK16QTextBlockFormat9alignmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "alignment", args)
  }

}

// leftMargin()
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
    // invoke: qreal leftMargin()
    C.C_ZNK16QTextBlockFormat10leftMarginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "leftMargin", args)
  }

}

// topMargin()
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
    // invoke: qreal topMargin()
    C.C_ZNK16QTextBlockFormat9topMarginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "topMargin", args)
  }

}

// setLeftMargin(qreal)
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
    // invoke: void setLeftMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextBlockFormat13setLeftMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setLeftMargin", args)
  }

}

// lineHeight()
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
    // invoke: qreal lineHeight()
    C.C_ZNK16QTextBlockFormat10lineHeightEv(this.qclsinst)
  case 1:
    // invoke: _ZNK16QTextBlockFormat10lineHeightEdd
    // invoke: qreal lineHeight(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZNK16QTextBlockFormat10lineHeightEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "lineHeight", args)
  }

}

// setTopMargin(qreal)
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
    // invoke: void setTopMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextBlockFormat12setTopMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setTopMargin", args)
  }

}

// rightMargin()
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
    // invoke: qreal rightMargin()
    C.C_ZNK16QTextBlockFormat11rightMarginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "rightMargin", args)
  }

}

// isValid()
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
    // invoke: bool isValid()
    C.C_ZNK16QTextBlockFormat7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "isValid", args)
  }

}

// setIndent(int)
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
    // invoke: void setIndent(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextBlockFormat9setIndentEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setIndent", args)
  }

}

// pageBreakPolicy()
func (this *QTextBlockFormat) pageBreakPolicy(args ...interface{}) () {
  // pageBreakPolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat15pageBreakPolicyEv
    // invoke: PageBreakFlags pageBreakPolicy()
    C.C_ZNK16QTextBlockFormat15pageBreakPolicyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "pageBreakPolicy", args)
  }

}

// bottomMargin()
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
    // invoke: qreal bottomMargin()
    C.C_ZNK16QTextBlockFormat12bottomMarginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "bottomMargin", args)
  }

}

// setLineHeight(qreal, int)
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
    // invoke: void setLineHeight(qreal, int)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN16QTextBlockFormat13setLineHeightEdi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setLineHeight", args)
  }

}

// setNonBreakableLines(_Bool)
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
    // invoke: void setNonBreakableLines(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextBlockFormat20setNonBreakableLinesEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setNonBreakableLines", args)
  }

}

// setRightMargin(qreal)
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
    // invoke: void setRightMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextBlockFormat14setRightMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setRightMargin", args)
  }

}

// indent()
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
    // invoke: int indent()
    C.C_ZNK16QTextBlockFormat6indentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "indent", args)
  }

}

// tabPositions()
func (this *QTextBlockFormat) tabPositions(args ...interface{}) () {
  // tabPositions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat12tabPositionsEv
    // invoke: QList<QTextOption::Tab> tabPositions()
    C.C_ZNK16QTextBlockFormat12tabPositionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "tabPositions", args)
  }

}

// textIndent()
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
    // invoke: qreal textIndent()
    C.C_ZNK16QTextBlockFormat10textIndentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "textIndent", args)
  }

}

// setBottomMargin(qreal)
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
    // invoke: void setBottomMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextBlockFormat15setBottomMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setBottomMargin", args)
  }

}

// nonBreakableLines()
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
    // invoke: bool nonBreakableLines()
    C.C_ZNK16QTextBlockFormat17nonBreakableLinesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "nonBreakableLines", args)
  }

}

// QTextBlockFormat()
func NewQTextBlockFormat(args ...interface{}) QTextBlockFormat {
  // QTextBlockFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextBlockFormatC1Ev
    // invoke: void QTextBlockFormat()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN16QTextBlockFormatC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "QTextBlockFormat", args)
  }

  return QTextBlockFormat{}
}

// setTableCellColumnSpan(int)
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
    // invoke: void setTableCellColumnSpan(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat22setTableCellColumnSpanEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setTableCellColumnSpan", args)
  }

}

// setAnchorHref(const class QString &)
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
    // invoke: void setAnchorHref(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat13setAnchorHrefERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setAnchorHref", args)
  }

}

// setFontOverline(_Bool)
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
    // invoke: void setFontOverline(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat15setFontOverlineEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontOverline", args)
  }

}

// setTextOutline(const class QPen &)
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
    // invoke: void setTextOutline(const class QPen &)
    var arg0 = args[0].(QPen).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat14setTextOutlineERK4QPen(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setTextOutline", args)
  }

}

// fontFamily()
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
    // invoke: QString fontFamily()
    C.C_ZNK15QTextCharFormat10fontFamilyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontFamily", args)
  }

}

// anchorNames()
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
    // invoke: QStringList anchorNames()
    C.C_ZNK15QTextCharFormat11anchorNamesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "anchorNames", args)
  }

}

// fontStyleHint()
func (this *QTextCharFormat) fontStyleHint(args ...interface{}) () {
  // fontStyleHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat13fontStyleHintEv
    // invoke: QFont::StyleHint fontStyleHint()
    C.C_ZNK15QTextCharFormat13fontStyleHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontStyleHint", args)
  }

}

// setFontStretch(int)
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
    // invoke: void setFontStretch(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat14setFontStretchEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontStretch", args)
  }

}

// font()
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
    // invoke: QFont font()
    C.C_ZNK15QTextCharFormat4fontEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "font", args)
  }

}

// fontItalic()
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
    // invoke: bool fontItalic()
    C.C_ZNK15QTextCharFormat10fontItalicEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontItalic", args)
  }

}

// fontKerning()
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
    // invoke: bool fontKerning()
    C.C_ZNK15QTextCharFormat11fontKerningEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontKerning", args)
  }

}

// fontWordSpacing()
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
    // invoke: qreal fontWordSpacing()
    C.C_ZNK15QTextCharFormat15fontWordSpacingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontWordSpacing", args)
  }

}

// fontOverline()
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
    // invoke: bool fontOverline()
    C.C_ZNK15QTextCharFormat12fontOverlineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontOverline", args)
  }

}

// setFontUnderline(_Bool)
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
    // invoke: void setFontUnderline(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat16setFontUnderlineEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontUnderline", args)
  }

}

// underlineStyle()
func (this *QTextCharFormat) underlineStyle(args ...interface{}) () {
  // underlineStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat14underlineStyleEv
    // invoke: QTextCharFormat::UnderlineStyle underlineStyle()
    C.C_ZNK15QTextCharFormat14underlineStyleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "underlineStyle", args)
  }

}

// fontStyleStrategy()
func (this *QTextCharFormat) fontStyleStrategy(args ...interface{}) () {
  // fontStyleStrategy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat17fontStyleStrategyEv
    // invoke: QFont::StyleStrategy fontStyleStrategy()
    C.C_ZNK15QTextCharFormat17fontStyleStrategyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontStyleStrategy", args)
  }

}

// setFontWordSpacing(qreal)
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
    // invoke: void setFontWordSpacing(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat18setFontWordSpacingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontWordSpacing", args)
  }

}

// setFontWeight(int)
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
    // invoke: void setFontWeight(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat13setFontWeightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontWeight", args)
  }

}

// verticalAlignment()
func (this *QTextCharFormat) verticalAlignment(args ...interface{}) () {
  // verticalAlignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat17verticalAlignmentEv
    // invoke: QTextCharFormat::VerticalAlignment verticalAlignment()
    C.C_ZNK15QTextCharFormat17verticalAlignmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "verticalAlignment", args)
  }

}

// setFontFamily(const class QString &)
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
    // invoke: void setFontFamily(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat13setFontFamilyERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontFamily", args)
  }

}

// setAnchorNames(const class QStringList &)
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
    // invoke: void setAnchorNames(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat14setAnchorNamesERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setAnchorNames", args)
  }

}

// setTableCellRowSpan(int)
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
    // invoke: void setTableCellRowSpan(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat19setTableCellRowSpanEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setTableCellRowSpan", args)
  }

}

// setFontKerning(_Bool)
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
    // invoke: void setFontKerning(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat14setFontKerningEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontKerning", args)
  }

}

// setFontLetterSpacing(qreal)
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
    // invoke: void setFontLetterSpacing(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat20setFontLetterSpacingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontLetterSpacing", args)
  }

}

// fontLetterSpacingType()
func (this *QTextCharFormat) fontLetterSpacingType(args ...interface{}) () {
  // fontLetterSpacingType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat21fontLetterSpacingTypeEv
    // invoke: QFont::SpacingType fontLetterSpacingType()
    C.C_ZNK15QTextCharFormat21fontLetterSpacingTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontLetterSpacingType", args)
  }

}

// fontUnderline()
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
    // invoke: bool fontUnderline()
    C.C_ZNK15QTextCharFormat13fontUnderlineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontUnderline", args)
  }

}

// fontHintingPreference()
func (this *QTextCharFormat) fontHintingPreference(args ...interface{}) () {
  // fontHintingPreference()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat21fontHintingPreferenceEv
    // invoke: QFont::HintingPreference fontHintingPreference()
    C.C_ZNK15QTextCharFormat21fontHintingPreferenceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontHintingPreference", args)
  }

}

// setFontStrikeOut(_Bool)
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
    // invoke: void setFontStrikeOut(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat16setFontStrikeOutEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontStrikeOut", args)
  }

}

// anchorName()
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
    // invoke: QString anchorName()
    C.C_ZNK15QTextCharFormat10anchorNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "anchorName", args)
  }

}

// fontPointSize()
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
    // invoke: qreal fontPointSize()
    C.C_ZNK15QTextCharFormat13fontPointSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontPointSize", args)
  }

}

// fontFixedPitch()
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
    // invoke: bool fontFixedPitch()
    C.C_ZNK15QTextCharFormat14fontFixedPitchEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontFixedPitch", args)
  }

}

// isValid()
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
    // invoke: bool isValid()
    C.C_ZNK15QTextCharFormat7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "isValid", args)
  }

}

// toolTip()
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
    // invoke: QString toolTip()
    C.C_ZNK15QTextCharFormat7toolTipEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "toolTip", args)
  }

}

// setFontPointSize(qreal)
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
    // invoke: void setFontPointSize(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat16setFontPointSizeEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontPointSize", args)
  }

}

// setAnchor(_Bool)
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
    // invoke: void setAnchor(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat9setAnchorEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setAnchor", args)
  }

}

// setAnchorName(const class QString &)
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
    // invoke: void setAnchorName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat13setAnchorNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setAnchorName", args)
  }

}

// fontLetterSpacing()
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
    // invoke: qreal fontLetterSpacing()
    C.C_ZNK15QTextCharFormat17fontLetterSpacingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontLetterSpacing", args)
  }

}

// isAnchor()
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
    // invoke: bool isAnchor()
    C.C_ZNK15QTextCharFormat8isAnchorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "isAnchor", args)
  }

}

// setFontItalic(_Bool)
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
    // invoke: void setFontItalic(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat13setFontItalicEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontItalic", args)
  }

}

// QTextCharFormat()
func NewQTextCharFormat(args ...interface{}) QTextCharFormat {
  // QTextCharFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormatC1Ev
    // invoke: void QTextCharFormat()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN15QTextCharFormatC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "QTextCharFormat", args)
  }

  return QTextCharFormat{}
}

// underlineColor()
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
    // invoke: QColor underlineColor()
    C.C_ZNK15QTextCharFormat14underlineColorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "underlineColor", args)
  }

}

// setUnderlineColor(const class QColor &)
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
    // invoke: void setUnderlineColor(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat17setUnderlineColorERK6QColor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setUnderlineColor", args)
  }

}

// fontWeight()
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
    // invoke: int fontWeight()
    C.C_ZNK15QTextCharFormat10fontWeightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontWeight", args)
  }

}

// setToolTip(const class QString &)
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
    // invoke: void setToolTip(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat10setToolTipERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setToolTip", args)
  }

}

// tableCellColumnSpan()
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
    // invoke: int tableCellColumnSpan()
    C.C_ZNK15QTextCharFormat19tableCellColumnSpanEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "tableCellColumnSpan", args)
  }

}

// anchorHref()
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
    // invoke: QString anchorHref()
    C.C_ZNK15QTextCharFormat10anchorHrefEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "anchorHref", args)
  }

}

// fontCapitalization()
func (this *QTextCharFormat) fontCapitalization(args ...interface{}) () {
  // fontCapitalization()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat18fontCapitalizationEv
    // invoke: QFont::Capitalization fontCapitalization()
    C.C_ZNK15QTextCharFormat18fontCapitalizationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontCapitalization", args)
  }

}

// textOutline()
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
    // invoke: QPen textOutline()
    C.C_ZNK15QTextCharFormat11textOutlineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "textOutline", args)
  }

}

// tableCellRowSpan()
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
    // invoke: int tableCellRowSpan()
    C.C_ZNK15QTextCharFormat16tableCellRowSpanEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "tableCellRowSpan", args)
  }

}

// setFontFixedPitch(_Bool)
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
    // invoke: void setFontFixedPitch(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat17setFontFixedPitchEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontFixedPitch", args)
  }

}

// setFont(const class QFont &)
func (this *QTextCharFormat) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat7setFontERK5QFont
    // invoke: void setFont(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat7setFontERK5QFont(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFont", args)
  }

}

// fontStretch()
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
    // invoke: int fontStretch()
    C.C_ZNK15QTextCharFormat11fontStretchEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontStretch", args)
  }

}

// fontStrikeOut()
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
    // invoke: bool fontStrikeOut()
    C.C_ZNK15QTextCharFormat13fontStrikeOutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontStrikeOut", args)
  }

}

// clearColumnWidthConstraints()
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
    // invoke: void clearColumnWidthConstraints()
    C.C_ZN16QTextTableFormat27clearColumnWidthConstraintsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableFormat", "clearColumnWidthConstraints", args)
  }

}

// columnWidthConstraints()
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
    // invoke: QVector<QTextLength> columnWidthConstraints()
    C.C_ZNK16QTextTableFormat22columnWidthConstraintsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableFormat", "columnWidthConstraints", args)
  }

}

// QTextTableFormat()
func NewQTextTableFormat(args ...interface{}) QTextTableFormat {
  // QTextTableFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextTableFormatC1Ev
    // invoke: void QTextTableFormat()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN16QTextTableFormatC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QTextTableFormat", "QTextTableFormat", args)
  }

  return QTextTableFormat{}
}

// setHeaderRowCount(int)
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
    // invoke: void setHeaderRowCount(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextTableFormat17setHeaderRowCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTableFormat", "setHeaderRowCount", args)
  }

}

// headerRowCount()
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
    // invoke: int headerRowCount()
    C.C_ZNK16QTextTableFormat14headerRowCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableFormat", "headerRowCount", args)
  }

}

// setColumns(int)
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
    // invoke: void setColumns(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextTableFormat10setColumnsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTableFormat", "setColumns", args)
  }

}

// alignment()
func (this *QTextTableFormat) alignment(args ...interface{}) () {
  // alignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextTableFormat9alignmentEv
    // invoke: Qt::Alignment alignment()
    C.C_ZNK16QTextTableFormat9alignmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableFormat", "alignment", args)
  }

}

// isValid()
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
    // invoke: bool isValid()
    C.C_ZNK16QTextTableFormat7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableFormat", "isValid", args)
  }

}

// cellPadding()
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
    // invoke: qreal cellPadding()
    C.C_ZNK16QTextTableFormat11cellPaddingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableFormat", "cellPadding", args)
  }

}

// setCellSpacing(qreal)
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
    // invoke: void setCellSpacing(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextTableFormat14setCellSpacingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTableFormat", "setCellSpacing", args)
  }

}

// cellSpacing()
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
    // invoke: qreal cellSpacing()
    C.C_ZNK16QTextTableFormat11cellSpacingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableFormat", "cellSpacing", args)
  }

}

// setCellPadding(qreal)
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
    // invoke: void setCellPadding(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextTableFormat14setCellPaddingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTableFormat", "setCellPadding", args)
  }

}

// columns()
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
    // invoke: int columns()
    C.C_ZNK16QTextTableFormat7columnsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableFormat", "columns", args)
  }

}

// setLeftPadding(qreal)
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
    // invoke: void setLeftPadding(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN20QTextTableCellFormat14setLeftPaddingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "setLeftPadding", args)
  }

}

// leftPadding()
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
    // invoke: qreal leftPadding()
    C.C_ZNK20QTextTableCellFormat11leftPaddingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "leftPadding", args)
  }

}

// setBottomPadding(qreal)
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
    // invoke: void setBottomPadding(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN20QTextTableCellFormat16setBottomPaddingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "setBottomPadding", args)
  }

}

// QTextTableCellFormat()
func NewQTextTableCellFormat(args ...interface{}) QTextTableCellFormat {
  // QTextTableCellFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QTextTableCellFormatC1Ev
    // invoke: void QTextTableCellFormat()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN20QTextTableCellFormatC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "QTextTableCellFormat", args)
  }

  return QTextTableCellFormat{}
}

// setTopPadding(qreal)
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
    // invoke: void setTopPadding(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN20QTextTableCellFormat13setTopPaddingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "setTopPadding", args)
  }

}

// rightPadding()
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
    // invoke: qreal rightPadding()
    C.C_ZNK20QTextTableCellFormat12rightPaddingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "rightPadding", args)
  }

}

// setRightPadding(qreal)
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
    // invoke: void setRightPadding(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN20QTextTableCellFormat15setRightPaddingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "setRightPadding", args)
  }

}

// isValid()
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
    // invoke: bool isValid()
    C.C_ZNK20QTextTableCellFormat7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "isValid", args)
  }

}

// topPadding()
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
    // invoke: qreal topPadding()
    C.C_ZNK20QTextTableCellFormat10topPaddingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "topPadding", args)
  }

}

// setPadding(qreal)
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
    // invoke: void setPadding(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN20QTextTableCellFormat10setPaddingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "setPadding", args)
  }

}

// bottomPadding()
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
    // invoke: qreal bottomPadding()
    C.C_ZNK20QTextTableCellFormat13bottomPaddingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "bottomPadding", args)
  }

}

// style()
func (this *QTextListFormat) style(args ...interface{}) () {
  // style()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextListFormat5styleEv
    // invoke: QTextListFormat::Style style()
    C.C_ZNK15QTextListFormat5styleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextListFormat", "style", args)
  }

}

// numberPrefix()
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
    // invoke: QString numberPrefix()
    C.C_ZNK15QTextListFormat12numberPrefixEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextListFormat", "numberPrefix", args)
  }

}

// indent()
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
    // invoke: int indent()
    C.C_ZNK15QTextListFormat6indentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextListFormat", "indent", args)
  }

}

// setNumberPrefix(const class QString &)
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
    // invoke: void setNumberPrefix(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextListFormat15setNumberPrefixERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextListFormat", "setNumberPrefix", args)
  }

}

// isValid()
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
    // invoke: bool isValid()
    C.C_ZNK15QTextListFormat7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextListFormat", "isValid", args)
  }

}

// QTextListFormat()
func NewQTextListFormat(args ...interface{}) QTextListFormat {
  // QTextListFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextListFormatC1Ev
    // invoke: void QTextListFormat()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN15QTextListFormatC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QTextListFormat", "QTextListFormat", args)
  }

  return QTextListFormat{}
}

// setIndent(int)
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
    // invoke: void setIndent(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextListFormat9setIndentEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextListFormat", "setIndent", args)
  }

}

// numberSuffix()
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
    // invoke: QString numberSuffix()
    C.C_ZNK15QTextListFormat12numberSuffixEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextListFormat", "numberSuffix", args)
  }

}

// setNumberSuffix(const class QString &)
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
    // invoke: void setNumberSuffix(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextListFormat15setNumberSuffixERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextListFormat", "setNumberSuffix", args)
  }

}

// setBorderBrush(const class QBrush &)
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
    // invoke: void setBorderBrush(const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat14setBorderBrushERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setBorderBrush", args)
  }

}

// height()
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
    // invoke: QTextLength height()
    C.C_ZNK16QTextFrameFormat6heightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "height", args)
  }

}

// setWidth(qreal)
func (this *QTextFrameFormat) setWidth(args ...interface{}) () {
  // setWidth(qreal)
  // setWidth(const class QTextLength &)
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
    // invoke: _ZN16QTextFrameFormat8setWidthEd
    // invoke: void setWidth(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat8setWidthEd(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN16QTextFrameFormat8setWidthERK11QTextLength
    // invoke: void setWidth(const class QTextLength &)
    var arg0 = args[0].(QTextLength).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat8setWidthERK11QTextLength(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setWidth", args)
  }

}

// topMargin()
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
    // invoke: qreal topMargin()
    C.C_ZNK16QTextFrameFormat9topMarginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "topMargin", args)
  }

}

// border()
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
    // invoke: qreal border()
    C.C_ZNK16QTextFrameFormat6borderEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "border", args)
  }

}

// width()
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
    // invoke: QTextLength width()
    C.C_ZNK16QTextFrameFormat5widthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "width", args)
  }

}

// setLeftMargin(qreal)
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
    // invoke: void setLeftMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat13setLeftMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setLeftMargin", args)
  }

}

// leftMargin()
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
    // invoke: qreal leftMargin()
    C.C_ZNK16QTextFrameFormat10leftMarginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "leftMargin", args)
  }

}

// setPadding(qreal)
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
    // invoke: void setPadding(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat10setPaddingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setPadding", args)
  }

}

// setTopMargin(qreal)
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
    // invoke: void setTopMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat12setTopMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setTopMargin", args)
  }

}

// rightMargin()
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
    // invoke: qreal rightMargin()
    C.C_ZNK16QTextFrameFormat11rightMarginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "rightMargin", args)
  }

}

// isValid()
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
    // invoke: bool isValid()
    C.C_ZNK16QTextFrameFormat7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "isValid", args)
  }

}

// pageBreakPolicy()
func (this *QTextFrameFormat) pageBreakPolicy(args ...interface{}) () {
  // pageBreakPolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat15pageBreakPolicyEv
    // invoke: PageBreakFlags pageBreakPolicy()
    C.C_ZNK16QTextFrameFormat15pageBreakPolicyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "pageBreakPolicy", args)
  }

}

// padding()
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
    // invoke: qreal padding()
    C.C_ZNK16QTextFrameFormat7paddingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "padding", args)
  }

}

// setMargin(qreal)
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
    // invoke: void setMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat9setMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setMargin", args)
  }

}

// bottomMargin()
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
    // invoke: qreal bottomMargin()
    C.C_ZNK16QTextFrameFormat12bottomMarginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "bottomMargin", args)
  }

}

// QTextFrameFormat()
func NewQTextFrameFormat(args ...interface{}) QTextFrameFormat {
  // QTextFrameFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextFrameFormatC1Ev
    // invoke: void QTextFrameFormat()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN16QTextFrameFormatC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "QTextFrameFormat", args)
  }

  return QTextFrameFormat{}
}

// setHeight(qreal)
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
    // invoke: void setHeight(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat9setHeightEd(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN16QTextFrameFormat9setHeightERK11QTextLength
    // invoke: void setHeight(const class QTextLength &)
    var arg0 = args[0].(QTextLength).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat9setHeightERK11QTextLength(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setHeight", args)
  }

}

// borderBrush()
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
    // invoke: QBrush borderBrush()
    C.C_ZNK16QTextFrameFormat11borderBrushEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "borderBrush", args)
  }

}

// setRightMargin(qreal)
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
    // invoke: void setRightMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat14setRightMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setRightMargin", args)
  }

}

// setBorder(qreal)
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
    // invoke: void setBorder(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat9setBorderEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setBorder", args)
  }

}

// setBottomMargin(qreal)
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
    // invoke: void setBottomMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat15setBottomMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setBottomMargin", args)
  }

}

// position()
func (this *QTextFrameFormat) position(args ...interface{}) () {
  // position()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat8positionEv
    // invoke: QTextFrameFormat::Position position()
    C.C_ZNK16QTextFrameFormat8positionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "position", args)
  }

}

// borderStyle()
func (this *QTextFrameFormat) borderStyle(args ...interface{}) () {
  // borderStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat11borderStyleEv
    // invoke: QTextFrameFormat::BorderStyle borderStyle()
    C.C_ZNK16QTextFrameFormat11borderStyleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "borderStyle", args)
  }

}

// margin()
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
    // invoke: qreal margin()
    C.C_ZNK16QTextFrameFormat6marginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "margin", args)
  }

}

// <= body block end

