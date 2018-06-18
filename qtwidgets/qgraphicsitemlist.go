package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>

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
import "github.com/kitech/qt.go/qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end

//  body block begin
type QGraphicsItemList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QGraphicsItemList) Operator_equal_0() *QGraphicsItemList {
	// QGraphicsItemList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QGraphicsItemList) Operator_equal_1() *QGraphicsItemList {
	// QGraphicsItemList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QGraphicsItemList) Swap_0() {
	// QGraphicsItemList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QGraphicsItemList) Operator_equal_equal_0() bool {
	// QGraphicsItemList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QGraphicsItemList) Operator_not_equal_0() bool {
	// QGraphicsItemList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QGraphicsItemList) Size_0() int {
	// QGraphicsItemList_size_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QGraphicsItemList) Detach_0() {
	// QGraphicsItemList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QGraphicsItemList) DetachShared_0() {
	// QGraphicsItemList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QGraphicsItemList) IsDetached_0() bool {
	// QGraphicsItemList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QGraphicsItemList) SetSharable_0() {
	// QGraphicsItemList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QGraphicsItemList) IsSharedWith_0() bool {
	// QGraphicsItemList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QGraphicsItemList) IsEmpty_0() bool {
	// QGraphicsItemList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QGraphicsItemList) Clear_0() {
	// QGraphicsItemList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QGraphicsItemList) At_0() *QGraphicsItem {
	// QGraphicsItemList_at_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsItem{}
}

// const T & operator[](int)
func (this *QGraphicsItemList) Operator_get_index_0() *QGraphicsItem {
	// QGraphicsItemList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsItem{}
}

// T & operator[](int)
func (this *QGraphicsItemList) Operator_get_index_1() *QGraphicsItem {
	// QGraphicsItemList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsItem{}
}

// void reserve(int)
func (this *QGraphicsItemList) Reserve_0() {
	// QGraphicsItemList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QGraphicsItemList) Append_0() {
	// QGraphicsItemList_append_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QGraphicsItemList) Append_1() {
	// QGraphicsItemList_append_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QGraphicsItemList) Prepend_0() {
	// QGraphicsItemList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QGraphicsItemList) Insert_0() {
	// QGraphicsItemList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QGraphicsItemList) Replace_0() {
	// QGraphicsItemList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QGraphicsItemList) RemoveAt_0() {
	// QGraphicsItemList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QGraphicsItemList) RemoveAll_0() int {
	// QGraphicsItemList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QGraphicsItemList) RemoveOne_0() bool {
	// QGraphicsItemList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QGraphicsItemList) TakeAt_0() *QGraphicsItem {
	// QGraphicsItemList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsItem{}
}

// T takeFirst()
func (this *QGraphicsItemList) TakeFirst_0() *QGraphicsItem {
	// QGraphicsItemList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsItem{}
}

// T takeLast()
func (this *QGraphicsItemList) TakeLast_0() *QGraphicsItem {
	// QGraphicsItemList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsItem{}
}

// void move(int, int)
func (this *QGraphicsItemList) Move_0() {
	// QGraphicsItemList_move_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QGraphicsItemList) Swap_1() {
	// QGraphicsItemList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QGraphicsItemList) IndexOf_0() int {
	// QGraphicsItemList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QGraphicsItemList) LastIndexOf_0() int {
	// QGraphicsItemList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QGraphicsItemList) Contains_0() bool {
	// QGraphicsItemList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QGraphicsItemList) Count_0() int {
	// QGraphicsItemList_count_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QGraphicsItemList) Begin_0() {
	// QGraphicsItemList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QGraphicsItemList) Begin_1() {
	// QGraphicsItemList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QGraphicsItemList) Cbegin_0() {
	// QGraphicsItemList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QGraphicsItemList) ConstBegin_0() {
	// QGraphicsItemList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QGraphicsItemList) End_0() {
	// QGraphicsItemList_end_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QGraphicsItemList) End_1() {
	// QGraphicsItemList_end_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QGraphicsItemList) Cend_0() {
	// QGraphicsItemList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QGraphicsItemList) ConstEnd_0() {
	// QGraphicsItemList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QGraphicsItemList) Rbegin_0() {
	// QGraphicsItemList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QGraphicsItemList) Rend_0() {
	// QGraphicsItemList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QGraphicsItemList) Rbegin_1() {
	// QGraphicsItemList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QGraphicsItemList) Rend_1() {
	// QGraphicsItemList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QGraphicsItemList) Crbegin_0() {
	// QGraphicsItemList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QGraphicsItemList) Crend_0() {
	// QGraphicsItemList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QGraphicsItemList) Insert_1() {
	// QGraphicsItemList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QGraphicsItemList) Erase_0() {
	// QGraphicsItemList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QGraphicsItemList) Erase_1() {
	// QGraphicsItemList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QGraphicsItemList) Count_1() int {
	// QGraphicsItemList_count_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QGraphicsItemList) Length_0() int {
	// QGraphicsItemList_length_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QGraphicsItemList) First_0() *QGraphicsItem {
	// QGraphicsItemList_first_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsItem{}
}

// const T & constFirst()
func (this *QGraphicsItemList) ConstFirst_0() *QGraphicsItem {
	// QGraphicsItemList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsItem{}
}

// const T & first()
func (this *QGraphicsItemList) First_1() *QGraphicsItem {
	// QGraphicsItemList_first_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsItem{}
}

// T & last()
func (this *QGraphicsItemList) Last_0() *QGraphicsItem {
	// QGraphicsItemList_last_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsItem{}
}

// const T & last()
func (this *QGraphicsItemList) Last_1() *QGraphicsItem {
	// QGraphicsItemList_last_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsItem{}
}

// const T & constLast()
func (this *QGraphicsItemList) ConstLast_0() *QGraphicsItem {
	// QGraphicsItemList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsItem{}
}

// void removeFirst()
func (this *QGraphicsItemList) RemoveFirst_0() {
	// QGraphicsItemList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QGraphicsItemList) RemoveLast_0() {
	// QGraphicsItemList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QGraphicsItemList) StartsWith_0() bool {
	// QGraphicsItemList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QGraphicsItemList) EndsWith_0() bool {
	// QGraphicsItemList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QGraphicsItemList) Mid_0() *QGraphicsItemList {
	// QGraphicsItemList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QGraphicsItemList) Value_0() *QGraphicsItem {
	// QGraphicsItemList_value_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsItem{}
}

// T value(int, const T &)
func (this *QGraphicsItemList) Value_1() *QGraphicsItem {
	// QGraphicsItemList_value_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsItem{}
}

// void push_back(const T &)
func (this *QGraphicsItemList) Push_back_0() {
	// QGraphicsItemList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QGraphicsItemList) Push_front_0() {
	// QGraphicsItemList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QGraphicsItemList) Front_0() *QGraphicsItem {
	// QGraphicsItemList_front_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsItem{}
}

// const T & front()
func (this *QGraphicsItemList) Front_1() *QGraphicsItem {
	// QGraphicsItemList_front_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsItem{}
}

// T & back()
func (this *QGraphicsItemList) Back_0() *QGraphicsItem {
	// QGraphicsItemList_back_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsItem{}
}

// const T & back()
func (this *QGraphicsItemList) Back_1() *QGraphicsItem {
	// QGraphicsItemList_back_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsItem{}
}

// void pop_front()
func (this *QGraphicsItemList) Pop_front_0() {
	// QGraphicsItemList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QGraphicsItemList) Pop_back_0() {
	// QGraphicsItemList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QGraphicsItemList) Empty_0() bool {
	// QGraphicsItemList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QGraphicsItemList) Operator_add_equal_0() *QGraphicsItemList {
	// QGraphicsItemList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QGraphicsItemList) Operator_add_0() *QGraphicsItemList {
	// QGraphicsItemList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QGraphicsItemList) Operator_add_equal_1() *QGraphicsItemList {
	// QGraphicsItemList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QGraphicsItemList) Operator_left_shift_0() *QGraphicsItemList {
	// QGraphicsItemList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QGraphicsItemList) Operator_left_shift_1() *QGraphicsItemList {
	// QGraphicsItemList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QGraphicsItemList) ToVector_0() {
	// QGraphicsItemList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QGraphicsItemList) ToSet_0() {
	// QGraphicsItemList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QGraphicsItemList) FromVector_0() *QGraphicsItemList {
	// QGraphicsItemList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QGraphicsItemList) FromSet_0() *QGraphicsItemList {
	// QGraphicsItemList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QGraphicsItemList) FromStdList_0() *QGraphicsItemList {
	// QGraphicsItemList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QGraphicsItemList) ToStdList_0() {
	// QGraphicsItemList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QGraphicsItemList) Detach_helper_grow_0() {
	// QGraphicsItemList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QGraphicsItemList) Detach_helper_0() {
	// QGraphicsItemList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QGraphicsItemList) Detach_helper_1() {
	// QGraphicsItemList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QGraphicsItemList) Dealloc_0() {
	// QGraphicsItemList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QGraphicsItemList) Node_construct_0() {
	// QGraphicsItemList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QGraphicsItemList) Node_destruct_0() {
	// QGraphicsItemList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QGraphicsItemList) Node_copy_0() {
	// QGraphicsItemList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QGraphicsItemList) Node_destruct_1() {
	// QGraphicsItemList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QGraphicsItemList) IsValidIterator_0() bool {
	// QGraphicsItemList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QGraphicsItemList) Op_eq_impl_0() bool {
	// QGraphicsItemList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QGraphicsItemList) Op_eq_impl_1() bool {
	// QGraphicsItemList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QGraphicsItemList) Contains_impl_0() bool {
	// QGraphicsItemList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QGraphicsItemList) Contains_impl_1() bool {
	// QGraphicsItemList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QGraphicsItemList) Count_impl_0() int {
	// QGraphicsItemList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QGraphicsItemList) Count_impl_1() int {
	// QGraphicsItemList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QGraphicsItemList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
