package data

import (
	"errors"
	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
)

// DataPage is a transfer object that is used to pass results of paginated queries. It contains items of retrieved
// page and optional total number of items.
// Most often this object type is used to send responses to paginated queries.
// Pagination parameters are defined by PagingParams object. The skip parameter in the PagingParams
// there means how many items to skip. The takes parameter sets number of items to return in the page.
// And the optional total parameter tells to return total number of items in the query.
// Remember: not all implementations support the total parameter because its generation may lead to severe
// performance implications.
//	see PagingParams
//
//	Example:
//  err, page = myDataClient.getDataByFilter(
//      "123",
//      FilterParams.fromTuples("completed": true),
//      NewPagingParams(0, 100, true)
//  	};
//
//  	if err != nil {
//  		panic()
//  	}
//  	for item range page.Data {
//          fmt.Println(item);
//      }
//  );
type DataPage[T any] struct {
	total    int
	data     []T
	hasTotal bool
}

const EmptyTotalValue int = -1

// NewEmptyDataPage creates a new empty instance of data page.
//	Returns: *DataPage
func NewEmptyDataPage[T any]() *DataPage[T] {
	return &DataPage[T]{
		total:    EmptyTotalValue,
		hasTotal: false,
		data:     nil,
	}
}

// NewDataPage creates a new instance of data page and assigns its values.
//	Parameters:
//		- value data a list of items from the retrieved page.
//		- total int
//	Returns: *DataPage
func NewDataPage[T any](data []T, total int) *DataPage[T] {
	dataPage := DataPage[T]{data: data}
	if total == EmptyTotalValue || total < len(data) {
		dataPage.hasTotal = false
		dataPage.total = EmptyTotalValue
	} else {
		dataPage.hasTotal = true
		dataPage.total = total
	}

	return &dataPage
}

func (d *DataPage[T]) Data() ([]T, bool) {
	if len(d.data) > 0 {
		return d.data, true
	}
	return nil, false
}

func (d *DataPage[T]) SetData(data []T) bool {
	if len(data) > 0 {
		d.data = data
		return true
	}
	return false
}

func (d *DataPage[T]) HasData() bool {
	return len(d.data) > 0
}

func (d *DataPage[T]) Total() (int, bool) {
	return d.total, d.hasTotal
}

func (d *DataPage[T]) SetTotal(total int) bool {
	if total > 0 {
		d.hasTotal = true
		d.total = total
		return true
	}
	d.hasTotal = false
	d.total = EmptyTotalValue
	return false
}

func (d *DataPage[T]) HasTotal() bool {
	return d.hasTotal
}

func (d *DataPage[T]) MarshalJSON() ([]byte, error) {
	result := map[string]any{
		"data": d.data,
	}
	if d.HasTotal() {
		result["total"] = d.total
	}
	buf, err := convert.JsonConverter.ToJson(result)
	return []byte(buf), err
}

func (d *DataPage[T]) UnmarshalJSON(data []byte) error {
	buf, err := convert.JsonConverter.FromJson(string(data))
	if err != nil {
		return err
	}
	bufMap, ok := buf.(map[string]any)
	if !ok {
		return errors.New("invalid type conversion")
	}
	if _data, ok := bufMap["data"]; ok {
		if val, ok := _data.([]T); ok {
			d.data = val
		}
	}

	if _total, ok := bufMap["total"]; ok {
		if val, ok := convert.LongConverter.ToNullableLong(_total); ok {
			if int(val) == EmptyTotalValue || int(val) < len(d.data) || val == 0 {
				d.hasTotal = false
				d.total = EmptyTotalValue
			} else {
				d.hasTotal = true
				d.total = int(val)
			}
		}
	}
	return nil
}
