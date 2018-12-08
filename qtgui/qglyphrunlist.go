package qtgui

// /usr/include/qt/QtGui/qglyphrun.h
// #include <qglyphrun.h>
// #include <QtGui>

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
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end

//  body block begin
type QGlyphRunList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QGlyphRunList) Operator_equal0() *QGlyphRunList {
	// QGlyphRunList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QGlyphRunList) Operator_equal1() *QGlyphRunList {
	// QGlyphRunList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QGlyphRunList) Swap0() {
	// QGlyphRunList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QGlyphRunList) Operator_equal_equal0() bool {
	// QGlyphRunList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QGlyphRunList) Operator_not_equal0() bool {
	// QGlyphRunList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QGlyphRunList) Size0() int {
	// QGlyphRunList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QGlyphRunList) Detach0() {
	// QGlyphRunList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QGlyphRunList) DetachShared0() {
	// QGlyphRunList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QGlyphRunList) IsDetached0() bool {
	// QGlyphRunList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QGlyphRunList) SetSharable0() {
	// QGlyphRunList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QGlyphRunList) IsSharedWith0() bool {
	// QGlyphRunList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QGlyphRunList) IsEmpty0() bool {
	// QGlyphRunList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QGlyphRunList) Clear0() {
	// QGlyphRunList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QGlyphRunList) At0() *QGlyphRun {
	// QGlyphRunList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// const T & operator[](int)
func (this *QGlyphRunList) Operator_get_index0() *QGlyphRun {
	// QGlyphRunList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// T & operator[](int)
func (this *QGlyphRunList) Operator_get_index1() *QGlyphRun {
	// QGlyphRunList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// void reserve(int)
func (this *QGlyphRunList) Reserve0() {
	// QGlyphRunList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QGlyphRunList) Append0() {
	// QGlyphRunList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QGlyphRunList) Append1() {
	// QGlyphRunList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QGlyphRunList) Prepend0() {
	// QGlyphRunList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QGlyphRunList) Insert0() {
	// QGlyphRunList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QGlyphRunList) Replace0() {
	// QGlyphRunList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QGlyphRunList) RemoveAt0() {
	// QGlyphRunList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QGlyphRunList) RemoveAll0() int {
	// QGlyphRunList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QGlyphRunList) RemoveOne0() bool {
	// QGlyphRunList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QGlyphRunList) TakeAt0() *QGlyphRun {
	// QGlyphRunList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// T takeFirst()
func (this *QGlyphRunList) TakeFirst0() *QGlyphRun {
	// QGlyphRunList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// T takeLast()
func (this *QGlyphRunList) TakeLast0() *QGlyphRun {
	// QGlyphRunList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// void move(int, int)
func (this *QGlyphRunList) Move0() {
	// QGlyphRunList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QGlyphRunList) Swap1() {
	// QGlyphRunList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QGlyphRunList) IndexOf0() int {
	// QGlyphRunList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QGlyphRunList) LastIndexOf0() int {
	// QGlyphRunList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QGlyphRunList) Contains0() bool {
	// QGlyphRunList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QGlyphRunList) Count0() int {
	// QGlyphRunList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QGlyphRunList) Begin0() {
	// QGlyphRunList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QGlyphRunList) Begin1() {
	// QGlyphRunList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QGlyphRunList) Cbegin0() {
	// QGlyphRunList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QGlyphRunList) ConstBegin0() {
	// QGlyphRunList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QGlyphRunList) End0() {
	// QGlyphRunList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QGlyphRunList) End1() {
	// QGlyphRunList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QGlyphRunList) Cend0() {
	// QGlyphRunList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QGlyphRunList) ConstEnd0() {
	// QGlyphRunList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QGlyphRunList) Rbegin0() {
	// QGlyphRunList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QGlyphRunList) Rend0() {
	// QGlyphRunList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QGlyphRunList) Rbegin1() {
	// QGlyphRunList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QGlyphRunList) Rend1() {
	// QGlyphRunList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QGlyphRunList) Crbegin0() {
	// QGlyphRunList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QGlyphRunList) Crend0() {
	// QGlyphRunList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QGlyphRunList) Insert1() {
	// QGlyphRunList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QGlyphRunList) Erase0() {
	// QGlyphRunList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QGlyphRunList) Erase1() {
	// QGlyphRunList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QGlyphRunList) Count1() int {
	// QGlyphRunList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QGlyphRunList) Length0() int {
	// QGlyphRunList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QGlyphRunList) First0() *QGlyphRun {
	// QGlyphRunList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// const T & constFirst()
func (this *QGlyphRunList) ConstFirst0() *QGlyphRun {
	// QGlyphRunList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// const T & first()
func (this *QGlyphRunList) First1() *QGlyphRun {
	// QGlyphRunList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// T & last()
func (this *QGlyphRunList) Last0() *QGlyphRun {
	// QGlyphRunList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// const T & last()
func (this *QGlyphRunList) Last1() *QGlyphRun {
	// QGlyphRunList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// const T & constLast()
func (this *QGlyphRunList) ConstLast0() *QGlyphRun {
	// QGlyphRunList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// void removeFirst()
func (this *QGlyphRunList) RemoveFirst0() {
	// QGlyphRunList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QGlyphRunList) RemoveLast0() {
	// QGlyphRunList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QGlyphRunList) StartsWith0() bool {
	// QGlyphRunList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QGlyphRunList) EndsWith0() bool {
	// QGlyphRunList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QGlyphRunList) Mid0() *QGlyphRunList {
	// QGlyphRunList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QGlyphRunList) Value0() *QGlyphRun {
	// QGlyphRunList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// T value(int, const T &)
func (this *QGlyphRunList) Value1() *QGlyphRun {
	// QGlyphRunList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// void push_back(const T &)
func (this *QGlyphRunList) Push_back0() {
	// QGlyphRunList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QGlyphRunList) Push_front0() {
	// QGlyphRunList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QGlyphRunList) Front0() *QGlyphRun {
	// QGlyphRunList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// const T & front()
func (this *QGlyphRunList) Front1() *QGlyphRun {
	// QGlyphRunList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// T & back()
func (this *QGlyphRunList) Back0() *QGlyphRun {
	// QGlyphRunList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// const T & back()
func (this *QGlyphRunList) Back1() *QGlyphRun {
	// QGlyphRunList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// void pop_front()
func (this *QGlyphRunList) Pop_front0() {
	// QGlyphRunList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QGlyphRunList) Pop_back0() {
	// QGlyphRunList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QGlyphRunList) Empty0() bool {
	// QGlyphRunList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QGlyphRunList) Operator_add_equal0() *QGlyphRunList {
	// QGlyphRunList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QGlyphRunList) Operator_add0() *QGlyphRunList {
	// QGlyphRunList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QGlyphRunList) Operator_add_equal1() *QGlyphRunList {
	// QGlyphRunList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QGlyphRunList) Operator_left_shift0() *QGlyphRunList {
	// QGlyphRunList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QGlyphRunList) Operator_left_shift1() *QGlyphRunList {
	// QGlyphRunList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QGlyphRunList) ToVector0() {
	// QGlyphRunList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QGlyphRunList) ToSet0() {
	// QGlyphRunList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QGlyphRunList) FromVector0() *QGlyphRunList {
	// QGlyphRunList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QGlyphRunList) FromSet0() *QGlyphRunList {
	// QGlyphRunList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QGlyphRunList) FromStdList0() *QGlyphRunList {
	// QGlyphRunList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QGlyphRunList) ToStdList0() {
	// QGlyphRunList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QGlyphRunList) Detach_helper_grow0() {
	// QGlyphRunList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QGlyphRunList) Detach_helper0() {
	// QGlyphRunList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QGlyphRunList) Detach_helper1() {
	// QGlyphRunList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QGlyphRunList) Dealloc0() {
	// QGlyphRunList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QGlyphRunList) Node_construct0() {
	// QGlyphRunList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QGlyphRunList) Node_destruct0() {
	// QGlyphRunList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QGlyphRunList) Node_copy0() {
	// QGlyphRunList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QGlyphRunList) Node_destruct1() {
	// QGlyphRunList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QGlyphRunList) IsValidIterator0() bool {
	// QGlyphRunList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QGlyphRunList) Op_eq_impl0() bool {
	// QGlyphRunList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QGlyphRunList) Op_eq_impl1() bool {
	// QGlyphRunList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QGlyphRunList) Contains_impl0() bool {
	// QGlyphRunList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QGlyphRunList) Contains_impl1() bool {
	// QGlyphRunList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QGlyphRunList) Count_impl0() int {
	// QGlyphRunList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QGlyphRunList) Count_impl1() int {
	// QGlyphRunList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGlyphRunList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
