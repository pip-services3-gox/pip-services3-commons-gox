package data

import (
	"errors"
	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
)

// TokenizedDataPage is a data transfer object that is used to pass results of paginated queries.
// It contains items of retrieved page and optional total number of items.
// Most often this object type is used to send responses to paginated queries.
// Pagination parameters are defined by TokenizedPagingParams object. The token parameter in the TokenizedPagingParams
// there determines a starting point for a new search. It is received in the TokenizedDataPage from the previous search.
// The takes parameter sets number of items to return in the page.
// And the optional total parameter tells to return total number of items in the query.
// Remember: not all implementations support the total parameter because its generation may lead to severe
// performance implications.
//	see TokenizedPagingParams
//
//	Example:
//		err, page = myDataClient.getDataByFilter(
//			"123",
//			FilterParams.fromTuples("completed": true),
//			NewTokenizedPagingParams("", 100, true)
//		};
//
//		if err != nil {
//			panic()
//		}
//		for item range page.Data {
//			fmt.Println(item);
//		}
type TokenizedDataPage[T any] struct {
	token string `json:"token"`
	data  []T    `json:"data"`
}

const EmptyTokenValue string = ""

// NewEmptyTokenizedDataPage creates a new empty instance of data page.
//	Returns: *TokenizedDataPage[T]
func NewEmptyTokenizedDataPage[T any]() *TokenizedDataPage[T] {
	return &TokenizedDataPage[T]{
		token: EmptyTokenValue,
	}
}

// NewTokenizedDataPage creates a new instance of data page and assigns its values.
//	Parameters:
//		- token a token that defines a starting point for next search
//		- data []T a list of items from the retrieved page.
//	Returns: *TokenizedDataPage[T]
func NewTokenizedDataPage[T any](token string, data []T) *TokenizedDataPage[T] {
	return &TokenizedDataPage[T]{token: token, data: data}
}

func (d *TokenizedDataPage[T]) Data() ([]T, bool) {
	if len(d.data) > 0 {
		return d.data, true
	}
	return nil, false
}

func (d *TokenizedDataPage[T]) SetData(data []T) bool {
	if len(data) > 0 {
		d.data = data
		return true
	}
	return false
}

func (d *TokenizedDataPage[T]) HasData() bool {
	return len(d.data) > 0
}

func (d *TokenizedDataPage[T]) Token() (string, bool) {
	if len(d.token) > 0 {
		return d.token, true
	}
	return d.token, false
}

func (d *TokenizedDataPage[T]) SetToken(token string) bool {
	if len(token) > 0 {
		d.token = token
		return true
	}
	return false
}

func (d *TokenizedDataPage[T]) HasToken() bool {
	return len(d.token) > 0
}

func (d *TokenizedDataPage[T]) MarshalJSON() ([]byte, error) {
	result := map[string]any{
		"data":  d.data,
		"token": d.token,
	}
	buf, err := convert.JsonConverter.ToJson(result)
	return []byte(buf), err
}

func (d *TokenizedDataPage[T]) UnmarshalJSON(data []byte) error {
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

	if _token, ok := bufMap["token"]; ok {
		if val, ok := _token.(string); ok {
			d.token = val
		}
	}

	return nil
}
