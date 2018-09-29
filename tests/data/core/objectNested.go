// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package test
import "encoding/json"
import "fmt"

type 421ArrayMyObjectArrayElem map[string]interface{}
type 421Array struct {
	// MyBooleanArray corresponds to the JSON schema field "myBooleanArray".
	MyBooleanArray []bool `json:"myBooleanArray,omitempty"`

	// MyNullArray corresponds to the JSON schema field "myNullArray".
	MyNullArray []interface{} `json:"myNullArray,omitempty"`

	// MyNumberArray corresponds to the JSON schema field "myNumberArray".
	MyNumberArray []float64 `json:"myNumberArray,omitempty"`

	// MyObjectArray corresponds to the JSON schema field "myObjectArray".
	MyObjectArray []421ArrayMyObjectArrayElem `json:"myObjectArray,omitempty"`

	// MyStringArray corresponds to the JSON schema field "myStringArray".
	MyStringArray []string `json:"myStringArray,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *421Array) UnmarshalJSON(b []byte) error {
	var v struct {
		421Array
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*j = v.421Array
	return nil
}


type ObjectMyObject struct {
	// MyString corresponds to the JSON schema field "myString".
	MyString string `json:"myString"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ObjectMyObject) UnmarshalJSON(b []byte) error {
	var v struct {
		ObjectMyObject
		__synthetic_MyString *string `json:"myString"`
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	if v.__synthetic_MyString == nil {
		return fmt.Errorf("field myString: must be set")
	}
	*j = v.ObjectMyObject
	return nil
}


type Object struct {
	// MyObject corresponds to the JSON schema field "myObject".
	MyObject *ObjectMyObject `json:"myObject,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Object) UnmarshalJSON(b []byte) error {
	var v struct {
		Object
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*j = v.Object
	return nil
}


type ObjectEmptyFoo map[string]interface{}
type ObjectEmpty struct {
	// Foo corresponds to the JSON schema field "foo".
	Foo ObjectEmptyFoo `json:"foo,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ObjectEmpty) UnmarshalJSON(b []byte) error {
	var v struct {
		ObjectEmpty
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*j = v.ObjectEmpty
	return nil
}


type ObjectNestedMyObjectMyObject struct {
	// MyString corresponds to the JSON schema field "myString".
	MyString *string `json:"myString,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ObjectNestedMyObjectMyObject) UnmarshalJSON(b []byte) error {
	var v struct {
		ObjectNestedMyObjectMyObject
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*j = v.ObjectNestedMyObjectMyObject
	return nil
}


type ObjectNestedMyObject struct {
	// MyObject corresponds to the JSON schema field "myObject".
	MyObject *ObjectNestedMyObjectMyObject `json:"myObject,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ObjectNestedMyObject) UnmarshalJSON(b []byte) error {
	var v struct {
		ObjectNestedMyObject
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*j = v.ObjectNestedMyObject
	return nil
}


type ObjectNested struct {
	// MyObject corresponds to the JSON schema field "myObject".
	MyObject *ObjectNestedMyObject `json:"myObject,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ObjectNested) UnmarshalJSON(b []byte) error {
	var v struct {
		ObjectNested
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*j = v.ObjectNested
	return nil
}

