package qtcore

// /usr/include/qt/QtCore/qvariant.h
// #include <qvariant.h>
// #include <QtCore>

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
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"

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
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QVariantList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QVariantList) Operator_equal_0() *QVariantList {
	// QVariantList_operator_equal_0()
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QVariantList) Operator_equal_1() *QVariantList {
	// QVariantList_operator_equal_1()
	return this
}

// void swap(QList<T> &)
func (this *QVariantList) Swap_0() {
	// QVariantList_swap_0()
}

// bool operator==(const QList<T> &)
func (this *QVariantList) Operator_equal_equal_0() bool {
	// QVariantList_operator_equal_equal_0()
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QVariantList) Operator_not_equal_0() bool {
	// QVariantList_operator_not_equal_0()
	return 0 == 0
}

// int size()
func (this *QVariantList) Size_0() int {
	// QVariantList_size_0()
	return 0
}

// void detach()
func (this *QVariantList) Detach_0() {
	// QVariantList_detach_0()
}

// void detachShared()
func (this *QVariantList) DetachShared_0() {
	// QVariantList_detachShared_0()
}

// bool isDetached()
func (this *QVariantList) IsDetached_0() bool {
	// QVariantList_isDetached_0()
	return 0 == 0
}

// void setSharable(_Bool)
func (this *QVariantList) SetSharable_0() {
	// QVariantList_setSharable_0()
}

// bool isSharedWith(const QList<T> &)
func (this *QVariantList) IsSharedWith_0() bool {
	// QVariantList_isSharedWith_0()
	return 0 == 0
}

// bool isEmpty()
func (this *QVariantList) IsEmpty_0() bool {
	// QVariantList_isEmpty_0()
	return 0 == 0
}

// void clear()
func (this *QVariantList) Clear_0() {
	// QVariantList_clear_0()
}

// const T & at(int)
func (this *QVariantList) At_0() *QVariant {
	// QVariantList_at_0()
	return &QVariant{}
}

// const T & operator[](int)
func (this *QVariantList) Operator_get_index_0() *QVariant {
	// QVariantList_operator_get_index_0()
	return &QVariant{}
}

// T & operator[](int)
func (this *QVariantList) Operator_get_index_1() *QVariant {
	// QVariantList_operator_get_index_1()
	return &QVariant{}
}

// void reserve(int)
func (this *QVariantList) Reserve_0() {
	// QVariantList_reserve_0()
}

// void append(const T &)
func (this *QVariantList) Append_0() {
	// QVariantList_append_0()
}

// void append(const QList<T> &)
func (this *QVariantList) Append_1() {
	// QVariantList_append_1()
}

// void prepend(const T &)
func (this *QVariantList) Prepend_0() {
	// QVariantList_prepend_0()
}

// void insert(int, const T &)
func (this *QVariantList) Insert_0() {
	// QVariantList_insert_0()
}

// void replace(int, const T &)
func (this *QVariantList) Replace_0() {
	// QVariantList_replace_0()
}

// void removeAt(int)
func (this *QVariantList) RemoveAt_0() {
	// QVariantList_removeAt_0()
}

// int removeAll(const T &)
func (this *QVariantList) RemoveAll_0() int {
	// QVariantList_removeAll_0()
	return 0
}

// bool removeOne(const T &)
func (this *QVariantList) RemoveOne_0() bool {
	// QVariantList_removeOne_0()
	return 0 == 0
}

// T takeAt(int)
func (this *QVariantList) TakeAt_0() *QVariant {
	// QVariantList_takeAt_0()
	return &QVariant{}
}

// T takeFirst()
func (this *QVariantList) TakeFirst_0() *QVariant {
	// QVariantList_takeFirst_0()
	return &QVariant{}
}

// T takeLast()
func (this *QVariantList) TakeLast_0() *QVariant {
	// QVariantList_takeLast_0()
	return &QVariant{}
}

// void move(int, int)
func (this *QVariantList) Move_0() {
	// QVariantList_move_0()
}

// void swap(int, int)
func (this *QVariantList) Swap_1() {
	// QVariantList_swap_1()
}

// int indexOf(const T &, int)
func (this *QVariantList) IndexOf_0() int {
	// QVariantList_indexOf_0()
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QVariantList) LastIndexOf_0() int {
	// QVariantList_lastIndexOf_0()
	return 0
}

// bool contains(const T &)
func (this *QVariantList) Contains_0() bool {
	// QVariantList_contains_0()
	return 0 == 0
}

// int count(const T &)
func (this *QVariantList) Count_0() int {
	// QVariantList_count_0()
	return 0
}

// QList::iterator begin()
func (this *QVariantList) Begin_0() {
	// QVariantList_begin_0()
}

// QList::const_iterator begin()
func (this *QVariantList) Begin_1() {
	// QVariantList_begin_1()
}

// QList::const_iterator cbegin()
func (this *QVariantList) Cbegin_0() {
	// QVariantList_cbegin_0()
}

// QList::const_iterator constBegin()
func (this *QVariantList) ConstBegin_0() {
	// QVariantList_constBegin_0()
}

// QList::iterator end()
func (this *QVariantList) End_0() {
	// QVariantList_end_0()
}

// QList::const_iterator end()
func (this *QVariantList) End_1() {
	// QVariantList_end_1()
}

// QList::const_iterator cend()
func (this *QVariantList) Cend_0() {
	// QVariantList_cend_0()
}

// QList::const_iterator constEnd()
func (this *QVariantList) ConstEnd_0() {
	// QVariantList_constEnd_0()
}

// QList::reverse_iterator rbegin()
func (this *QVariantList) Rbegin_0() {
	// QVariantList_rbegin_0()
}

// QList::reverse_iterator rend()
func (this *QVariantList) Rend_0() {
	// QVariantList_rend_0()
}

// QList::const_reverse_iterator rbegin()
func (this *QVariantList) Rbegin_1() {
	// QVariantList_rbegin_1()
}

// QList::const_reverse_iterator rend()
func (this *QVariantList) Rend_1() {
	// QVariantList_rend_1()
}

// QList::const_reverse_iterator crbegin()
func (this *QVariantList) Crbegin_0() {
	// QVariantList_crbegin_0()
}

// QList::const_reverse_iterator crend()
func (this *QVariantList) Crend_0() {
	// QVariantList_crend_0()
}

// QList::iterator insert(class QList::iterator, const T &)
func (this *QVariantList) Insert_1() {
	// QVariantList_insert_1()
}

// QList::iterator erase(class QList::iterator)
func (this *QVariantList) Erase_0() {
	// QVariantList_erase_0()
}

// QList::iterator erase(class QList::iterator, class QList::iterator)
func (this *QVariantList) Erase_1() {
	// QVariantList_erase_1()
}

// int count()
func (this *QVariantList) Count_1() int {
	// QVariantList_count_1()
	return 0
}

// int length()
func (this *QVariantList) Length_0() int {
	// QVariantList_length_0()
	return 0
}

// T & first()
func (this *QVariantList) First_0() *QVariant {
	// QVariantList_first_0()
	return &QVariant{}
}

// const T & constFirst()
func (this *QVariantList) ConstFirst_0() *QVariant {
	// QVariantList_constFirst_0()
	return &QVariant{}
}

// const T & first()
func (this *QVariantList) First_1() *QVariant {
	// QVariantList_first_1()
	return &QVariant{}
}

// T & last()
func (this *QVariantList) Last_0() *QVariant {
	// QVariantList_last_0()
	return &QVariant{}
}

// const T & last()
func (this *QVariantList) Last_1() *QVariant {
	// QVariantList_last_1()
	return &QVariant{}
}

// const T & constLast()
func (this *QVariantList) ConstLast_0() *QVariant {
	// QVariantList_constLast_0()
	return &QVariant{}
}

// void removeFirst()
func (this *QVariantList) RemoveFirst_0() {
	// QVariantList_removeFirst_0()
}

// void removeLast()
func (this *QVariantList) RemoveLast_0() {
	// QVariantList_removeLast_0()
}

// bool startsWith(const T &)
func (this *QVariantList) StartsWith_0() bool {
	// QVariantList_startsWith_0()
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QVariantList) EndsWith_0() bool {
	// QVariantList_endsWith_0()
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QVariantList) Mid_0() *QVariantList {
	// QVariantList_mid_0()
	return this
}

// T value(int)
func (this *QVariantList) Value_0() *QVariant {
	// QVariantList_value_0()
	return &QVariant{}
}

// T value(int, const T &)
func (this *QVariantList) Value_1() *QVariant {
	// QVariantList_value_1()
	return &QVariant{}
}

// void push_back(const T &)
func (this *QVariantList) Push_back_0() {
	// QVariantList_push_back_0()
}

// void push_front(const T &)
func (this *QVariantList) Push_front_0() {
	// QVariantList_push_front_0()
}

// T & front()
func (this *QVariantList) Front_0() *QVariant {
	// QVariantList_front_0()
	return &QVariant{}
}

// const T & front()
func (this *QVariantList) Front_1() *QVariant {
	// QVariantList_front_1()
	return &QVariant{}
}

// T & back()
func (this *QVariantList) Back_0() *QVariant {
	// QVariantList_back_0()
	return &QVariant{}
}

// const T & back()
func (this *QVariantList) Back_1() *QVariant {
	// QVariantList_back_1()
	return &QVariant{}
}

// void pop_front()
func (this *QVariantList) Pop_front_0() {
	// QVariantList_pop_front_0()
}

// void pop_back()
func (this *QVariantList) Pop_back_0() {
	// QVariantList_pop_back_0()
}

// bool empty()
func (this *QVariantList) Empty_0() bool {
	// QVariantList_empty_0()
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QVariantList) Operator_add_equal_0() *QVariantList {
	// QVariantList_operator_add_equal_0()
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QVariantList) Operator_add_0() *QVariantList {
	// QVariantList_operator_add_0()
	return this
}

// QList<T> & operator+=(const T &)
func (this *QVariantList) Operator_add_equal_1() *QVariantList {
	// QVariantList_operator_add_equal_1()
	return this
}

// QList<T> & operator<<(const T &)
func (this *QVariantList) Operator_left_shift_0() *QVariantList {
	// QVariantList_operator_left_shift_0()
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QVariantList) Operator_left_shift_1() *QVariantList {
	// QVariantList_operator_left_shift_1()
	return this
}

// QVector<T> toVector()
func (this *QVariantList) ToVector_0() {
	// QVariantList_toVector_0()
}

// QSet<T> toSet()
func (this *QVariantList) ToSet_0() {
	// QVariantList_toSet_0()
}

// QList<T> fromVector(const QVector<T> &)
func (this *QVariantList) FromVector_0() *QVariantList {
	// QVariantList_fromVector_0()
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QVariantList) FromSet_0() *QVariantList {
	// QVariantList_fromSet_0()
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QVariantList) FromStdList_0() *QVariantList {
	// QVariantList_fromStdList_0()
	return this
}

// std::list<T> toStdList()
func (this *QVariantList) ToStdList_0() {
	// QVariantList_toStdList_0()
}

// QList::Node * detach_helper_grow(int, int)
func (this *QVariantList) Detach_helper_grow_0() {
	// QVariantList_detach_helper_grow_0()
}

// void detach_helper(int)
func (this *QVariantList) Detach_helper_0() {
	// QVariantList_detach_helper_0()
}

// void detach_helper()
func (this *QVariantList) Detach_helper_1() {
	// QVariantList_detach_helper_1()
}

// void dealloc(struct QListData::Data *)
func (this *QVariantList) Dealloc_0() {
	// QVariantList_dealloc_0()
}

// void node_construct(struct QList::Node *, const T &)
func (this *QVariantList) Node_construct_0() {
	// QVariantList_node_construct_0()
}

// void node_destruct(struct QList::Node *)
func (this *QVariantList) Node_destruct_0() {
	// QVariantList_node_destruct_0()
}

// void node_copy(struct QList::Node *, struct QList::Node *, struct QList::Node *)
func (this *QVariantList) Node_copy_0() {
	// QVariantList_node_copy_0()
}

// void node_destruct(struct QList::Node *, struct QList::Node *)
func (this *QVariantList) Node_destruct_1() {
	// QVariantList_node_destruct_1()
}

// bool isValidIterator(const class QList::iterator &)
func (this *QVariantList) IsValidIterator_0() bool {
	// QVariantList_isValidIterator_0()
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, struct QListData::NotArrayCompatibleLayout)
func (this *QVariantList) Op_eq_impl_0() bool {
	// QVariantList_op_eq_impl_0()
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, struct QListData::ArrayCompatibleLayout)
func (this *QVariantList) Op_eq_impl_1() bool {
	// QVariantList_op_eq_impl_1()
	return 0 == 0
}

// bool contains_impl(const T &, struct QListData::NotArrayCompatibleLayout)
func (this *QVariantList) Contains_impl_0() bool {
	// QVariantList_contains_impl_0()
	return 0 == 0
}

// bool contains_impl(const T &, struct QListData::ArrayCompatibleLayout)
func (this *QVariantList) Contains_impl_1() bool {
	// QVariantList_contains_impl_1()
	return 0 == 0
}

// int count_impl(const T &, struct QListData::NotArrayCompatibleLayout)
func (this *QVariantList) Count_impl_0() int {
	// QVariantList_count_impl_0()
	return 0
}

// int count_impl(const T &, struct QListData::ArrayCompatibleLayout)
func (this *QVariantList) Count_impl_1() int {
	// QVariantList_count_impl_1()
	return 0
}

//  body block end
