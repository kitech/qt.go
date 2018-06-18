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
func (this *QGlyphRunList) Operator_equal_0() *QGlyphRunList {
	// QGlyphRunList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QGlyphRunList) Operator_equal_1() *QGlyphRunList {
	// QGlyphRunList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QGlyphRunList) Swap_0() {
	// QGlyphRunList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QGlyphRunList) Operator_equal_equal_0() bool {
	// QGlyphRunList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QGlyphRunList) Operator_not_equal_0() bool {
	// QGlyphRunList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QGlyphRunList) Size_0() int {
	// QGlyphRunList_size_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QGlyphRunList) Detach_0() {
	// QGlyphRunList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QGlyphRunList) DetachShared_0() {
	// QGlyphRunList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QGlyphRunList) IsDetached_0() bool {
	// QGlyphRunList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QGlyphRunList) SetSharable_0() {
	// QGlyphRunList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QGlyphRunList) IsSharedWith_0() bool {
	// QGlyphRunList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QGlyphRunList) IsEmpty_0() bool {
	// QGlyphRunList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QGlyphRunList) Clear_0() {
	// QGlyphRunList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QGlyphRunList) At_0() *QGlyphRun {
	// QGlyphRunList_at_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// const T & operator[](int)
func (this *QGlyphRunList) Operator_get_index_0() *QGlyphRun {
	// QGlyphRunList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// T & operator[](int)
func (this *QGlyphRunList) Operator_get_index_1() *QGlyphRun {
	// QGlyphRunList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// void reserve(int)
func (this *QGlyphRunList) Reserve_0() {
	// QGlyphRunList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QGlyphRunList) Append_0() {
	// QGlyphRunList_append_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QGlyphRunList) Append_1() {
	// QGlyphRunList_append_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QGlyphRunList) Prepend_0() {
	// QGlyphRunList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QGlyphRunList) Insert_0() {
	// QGlyphRunList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QGlyphRunList) Replace_0() {
	// QGlyphRunList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QGlyphRunList) RemoveAt_0() {
	// QGlyphRunList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QGlyphRunList) RemoveAll_0() int {
	// QGlyphRunList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QGlyphRunList) RemoveOne_0() bool {
	// QGlyphRunList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QGlyphRunList) TakeAt_0() *QGlyphRun {
	// QGlyphRunList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// T takeFirst()
func (this *QGlyphRunList) TakeFirst_0() *QGlyphRun {
	// QGlyphRunList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// T takeLast()
func (this *QGlyphRunList) TakeLast_0() *QGlyphRun {
	// QGlyphRunList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// void move(int, int)
func (this *QGlyphRunList) Move_0() {
	// QGlyphRunList_move_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QGlyphRunList) Swap_1() {
	// QGlyphRunList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QGlyphRunList) IndexOf_0() int {
	// QGlyphRunList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QGlyphRunList) LastIndexOf_0() int {
	// QGlyphRunList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QGlyphRunList) Contains_0() bool {
	// QGlyphRunList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QGlyphRunList) Count_0() int {
	// QGlyphRunList_count_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QGlyphRunList) Begin_0() {
	// QGlyphRunList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QGlyphRunList) Begin_1() {
	// QGlyphRunList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QGlyphRunList) Cbegin_0() {
	// QGlyphRunList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QGlyphRunList) ConstBegin_0() {
	// QGlyphRunList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QGlyphRunList) End_0() {
	// QGlyphRunList_end_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QGlyphRunList) End_1() {
	// QGlyphRunList_end_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QGlyphRunList) Cend_0() {
	// QGlyphRunList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QGlyphRunList) ConstEnd_0() {
	// QGlyphRunList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QGlyphRunList) Rbegin_0() {
	// QGlyphRunList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QGlyphRunList) Rend_0() {
	// QGlyphRunList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QGlyphRunList) Rbegin_1() {
	// QGlyphRunList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QGlyphRunList) Rend_1() {
	// QGlyphRunList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QGlyphRunList) Crbegin_0() {
	// QGlyphRunList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QGlyphRunList) Crend_0() {
	// QGlyphRunList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QGlyphRunList) Insert_1() {
	// QGlyphRunList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QGlyphRunList) Erase_0() {
	// QGlyphRunList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QGlyphRunList) Erase_1() {
	// QGlyphRunList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QGlyphRunList) Count_1() int {
	// QGlyphRunList_count_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QGlyphRunList) Length_0() int {
	// QGlyphRunList_length_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QGlyphRunList) First_0() *QGlyphRun {
	// QGlyphRunList_first_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// const T & constFirst()
func (this *QGlyphRunList) ConstFirst_0() *QGlyphRun {
	// QGlyphRunList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// const T & first()
func (this *QGlyphRunList) First_1() *QGlyphRun {
	// QGlyphRunList_first_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// T & last()
func (this *QGlyphRunList) Last_0() *QGlyphRun {
	// QGlyphRunList_last_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// const T & last()
func (this *QGlyphRunList) Last_1() *QGlyphRun {
	// QGlyphRunList_last_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// const T & constLast()
func (this *QGlyphRunList) ConstLast_0() *QGlyphRun {
	// QGlyphRunList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// void removeFirst()
func (this *QGlyphRunList) RemoveFirst_0() {
	// QGlyphRunList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QGlyphRunList) RemoveLast_0() {
	// QGlyphRunList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QGlyphRunList) StartsWith_0() bool {
	// QGlyphRunList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QGlyphRunList) EndsWith_0() bool {
	// QGlyphRunList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QGlyphRunList) Mid_0() *QGlyphRunList {
	// QGlyphRunList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QGlyphRunList) Value_0() *QGlyphRun {
	// QGlyphRunList_value_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// T value(int, const T &)
func (this *QGlyphRunList) Value_1() *QGlyphRun {
	// QGlyphRunList_value_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// void push_back(const T &)
func (this *QGlyphRunList) Push_back_0() {
	// QGlyphRunList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QGlyphRunList) Push_front_0() {
	// QGlyphRunList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QGlyphRunList) Front_0() *QGlyphRun {
	// QGlyphRunList_front_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// const T & front()
func (this *QGlyphRunList) Front_1() *QGlyphRun {
	// QGlyphRunList_front_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// T & back()
func (this *QGlyphRunList) Back_0() *QGlyphRun {
	// QGlyphRunList_back_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// const T & back()
func (this *QGlyphRunList) Back_1() *QGlyphRun {
	// QGlyphRunList_back_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGlyphRun{}
}

// void pop_front()
func (this *QGlyphRunList) Pop_front_0() {
	// QGlyphRunList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QGlyphRunList) Pop_back_0() {
	// QGlyphRunList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QGlyphRunList) Empty_0() bool {
	// QGlyphRunList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QGlyphRunList) Operator_add_equal_0() *QGlyphRunList {
	// QGlyphRunList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QGlyphRunList) Operator_add_0() *QGlyphRunList {
	// QGlyphRunList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QGlyphRunList) Operator_add_equal_1() *QGlyphRunList {
	// QGlyphRunList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QGlyphRunList) Operator_left_shift_0() *QGlyphRunList {
	// QGlyphRunList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QGlyphRunList) Operator_left_shift_1() *QGlyphRunList {
	// QGlyphRunList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QGlyphRunList) ToVector_0() {
	// QGlyphRunList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QGlyphRunList) ToSet_0() {
	// QGlyphRunList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QGlyphRunList) FromVector_0() *QGlyphRunList {
	// QGlyphRunList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QGlyphRunList) FromSet_0() *QGlyphRunList {
	// QGlyphRunList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QGlyphRunList) FromStdList_0() *QGlyphRunList {
	// QGlyphRunList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QGlyphRunList) ToStdList_0() {
	// QGlyphRunList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QGlyphRunList) Detach_helper_grow_0() {
	// QGlyphRunList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QGlyphRunList) Detach_helper_0() {
	// QGlyphRunList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QGlyphRunList) Detach_helper_1() {
	// QGlyphRunList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QGlyphRunList) Dealloc_0() {
	// QGlyphRunList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QGlyphRunList) Node_construct_0() {
	// QGlyphRunList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QGlyphRunList) Node_destruct_0() {
	// QGlyphRunList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QGlyphRunList) Node_copy_0() {
	// QGlyphRunList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QGlyphRunList) Node_destruct_1() {
	// QGlyphRunList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QGlyphRunList) IsValidIterator_0() bool {
	// QGlyphRunList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QGlyphRunList) Op_eq_impl_0() bool {
	// QGlyphRunList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QGlyphRunList) Op_eq_impl_1() bool {
	// QGlyphRunList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QGlyphRunList) Contains_impl_0() bool {
	// QGlyphRunList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QGlyphRunList) Contains_impl_1() bool {
	// QGlyphRunList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QGlyphRunList) Count_impl_0() int {
	// QGlyphRunList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QGlyphRunList) Count_impl_1() int {
	// QGlyphRunList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QGlyphRunList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
