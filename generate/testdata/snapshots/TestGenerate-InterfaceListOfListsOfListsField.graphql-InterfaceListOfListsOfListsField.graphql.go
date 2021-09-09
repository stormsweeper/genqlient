package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"encoding/json"
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// InterfaceListOfListOfListsFieldListOfListsOfListsOfContent includes the requested fields of the GraphQL interface Content.
//
// InterfaceListOfListOfListsFieldListOfListsOfListsOfContent is implemented by the following types:
// InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle
// InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo
// InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceListOfListOfListsFieldListOfListsOfListsOfContent interface {
	implementsGraphQLInterfaceInterfaceListOfListOfListsFieldListOfListsOfListsOfContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() testutil.ID
	// GetName returns the interface-field "name" from its implementation.
	GetName() string
}

func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle) implementsGraphQLInterfaceInterfaceListOfListOfListsFieldListOfListsOfListsOfContent() {
}

// GetTypename is a part of, and documented with, the interface InterfaceListOfListOfListsFieldListOfListsOfListsOfContent.
func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle) GetTypename() string {
	return v.Typename
}

// GetId is a part of, and documented with, the interface InterfaceListOfListOfListsFieldListOfListsOfListsOfContent.
func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle) GetId() testutil.ID {
	return v.Id
}

// GetName is a part of, and documented with, the interface InterfaceListOfListOfListsFieldListOfListsOfListsOfContent.
func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle) GetName() string {
	return v.Name
}

func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo) implementsGraphQLInterfaceInterfaceListOfListOfListsFieldListOfListsOfListsOfContent() {
}

// GetTypename is a part of, and documented with, the interface InterfaceListOfListOfListsFieldListOfListsOfListsOfContent.
func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo) GetTypename() string {
	return v.Typename
}

// GetId is a part of, and documented with, the interface InterfaceListOfListOfListsFieldListOfListsOfListsOfContent.
func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo) GetId() testutil.ID {
	return v.Id
}

// GetName is a part of, and documented with, the interface InterfaceListOfListOfListsFieldListOfListsOfListsOfContent.
func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo) GetName() string {
	return v.Name
}

func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic) implementsGraphQLInterfaceInterfaceListOfListOfListsFieldListOfListsOfListsOfContent() {
}

// GetTypename is a part of, and documented with, the interface InterfaceListOfListOfListsFieldListOfListsOfListsOfContent.
func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic) GetTypename() string {
	return v.Typename
}

// GetId is a part of, and documented with, the interface InterfaceListOfListOfListsFieldListOfListsOfListsOfContent.
func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic) GetId() testutil.ID {
	return v.Id
}

// GetName is a part of, and documented with, the interface InterfaceListOfListOfListsFieldListOfListsOfListsOfContent.
func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic) GetName() string {
	return v.Name
}

func __unmarshalInterfaceListOfListOfListsFieldListOfListsOfListsOfContent(v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContent, m json.RawMessage) error {
	if string(m) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(m, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle)
		return json.Unmarshal(m, *v)
	case "Video":
		*v = new(InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo)
		return json.Unmarshal(m, *v)
	case "Topic":
		*v = new(InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic)
		return json.Unmarshal(m, *v)
	case "":
		return fmt.Errorf(
			"Response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`Unexpected concrete type for InterfaceListOfListOfListsFieldListOfListsOfListsOfContent: "%v"`, tn.TypeName)
	}
}

// InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle includes the requested fields of the GraphQL type Article.
type InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic includes the requested fields of the GraphQL type Topic.
type InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo includes the requested fields of the GraphQL type Video.
type InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// InterfaceListOfListOfListsFieldResponse is returned by InterfaceListOfListOfListsField on success.
type InterfaceListOfListOfListsFieldResponse struct {
	ListOfListsOfListsOfContent [][][]InterfaceListOfListOfListsFieldListOfListsOfListsOfContent `json:"-"`
	WithPointer                 [][][]*InterfaceListOfListOfListsFieldWithPointerContent         `json:"-"`
}

func (v *InterfaceListOfListOfListsFieldResponse) UnmarshalJSON(b []byte) error {

	var firstPass struct {
		*InterfaceListOfListOfListsFieldResponse
		ListOfListsOfListsOfContent [][][]json.RawMessage `json:"listOfListsOfListsOfContent"`
		WithPointer                 [][][]json.RawMessage `json:"withPointer"`
		graphql.NoUnmarshalJSON
	}
	firstPass.InterfaceListOfListOfListsFieldResponse = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		target := &v.ListOfListsOfListsOfContent
		raw := firstPass.ListOfListsOfListsOfContent
		*target = make(
			[][][]InterfaceListOfListOfListsFieldListOfListsOfListsOfContent,
			len(raw))
		for i, raw := range raw {
			target := &(*target)[i]
			*target = make(
				[][]InterfaceListOfListOfListsFieldListOfListsOfListsOfContent,
				len(raw))
			for i, raw := range raw {
				target := &(*target)[i]
				*target = make(
					[]InterfaceListOfListOfListsFieldListOfListsOfListsOfContent,
					len(raw))
				for i, raw := range raw {
					target := &(*target)[i]
					err = __unmarshalInterfaceListOfListOfListsFieldListOfListsOfListsOfContent(
						target, raw)
					if err != nil {
						return fmt.Errorf(
							"Unable to unmarshal InterfaceListOfListOfListsFieldResponse.ListOfListsOfListsOfContent: %w", err)
					}
				}
			}
		}
	}

	{
		target := &v.WithPointer
		raw := firstPass.WithPointer
		*target = make(
			[][][]*InterfaceListOfListOfListsFieldWithPointerContent,
			len(raw))
		for i, raw := range raw {
			target := &(*target)[i]
			*target = make(
				[][]*InterfaceListOfListOfListsFieldWithPointerContent,
				len(raw))
			for i, raw := range raw {
				target := &(*target)[i]
				*target = make(
					[]*InterfaceListOfListOfListsFieldWithPointerContent,
					len(raw))
				for i, raw := range raw {
					target := &(*target)[i]
					*target = new(InterfaceListOfListOfListsFieldWithPointerContent)
					err = __unmarshalInterfaceListOfListOfListsFieldWithPointerContent(
						*target, raw)
					if err != nil {
						return fmt.Errorf(
							"Unable to unmarshal InterfaceListOfListOfListsFieldResponse.WithPointer: %w", err)
					}
				}
			}
		}
	}
	return nil
}

// InterfaceListOfListOfListsFieldWithPointerArticle includes the requested fields of the GraphQL type Article.
type InterfaceListOfListOfListsFieldWithPointerArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   *testutil.ID `json:"id"`
	Name *string      `json:"name"`
}

// InterfaceListOfListOfListsFieldWithPointerContent includes the requested fields of the GraphQL interface Content.
//
// InterfaceListOfListOfListsFieldWithPointerContent is implemented by the following types:
// InterfaceListOfListOfListsFieldWithPointerArticle
// InterfaceListOfListOfListsFieldWithPointerVideo
// InterfaceListOfListOfListsFieldWithPointerTopic
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceListOfListOfListsFieldWithPointerContent interface {
	implementsGraphQLInterfaceInterfaceListOfListOfListsFieldWithPointerContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() *testutil.ID
	// GetName returns the interface-field "name" from its implementation.
	GetName() *string
}

func (v *InterfaceListOfListOfListsFieldWithPointerArticle) implementsGraphQLInterfaceInterfaceListOfListOfListsFieldWithPointerContent() {
}

// GetTypename is a part of, and documented with, the interface InterfaceListOfListOfListsFieldWithPointerContent.
func (v *InterfaceListOfListOfListsFieldWithPointerArticle) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface InterfaceListOfListOfListsFieldWithPointerContent.
func (v *InterfaceListOfListOfListsFieldWithPointerArticle) GetId() *testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface InterfaceListOfListOfListsFieldWithPointerContent.
func (v *InterfaceListOfListOfListsFieldWithPointerArticle) GetName() *string { return v.Name }

func (v *InterfaceListOfListOfListsFieldWithPointerVideo) implementsGraphQLInterfaceInterfaceListOfListOfListsFieldWithPointerContent() {
}

// GetTypename is a part of, and documented with, the interface InterfaceListOfListOfListsFieldWithPointerContent.
func (v *InterfaceListOfListOfListsFieldWithPointerVideo) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface InterfaceListOfListOfListsFieldWithPointerContent.
func (v *InterfaceListOfListOfListsFieldWithPointerVideo) GetId() *testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface InterfaceListOfListOfListsFieldWithPointerContent.
func (v *InterfaceListOfListOfListsFieldWithPointerVideo) GetName() *string { return v.Name }

func (v *InterfaceListOfListOfListsFieldWithPointerTopic) implementsGraphQLInterfaceInterfaceListOfListOfListsFieldWithPointerContent() {
}

// GetTypename is a part of, and documented with, the interface InterfaceListOfListOfListsFieldWithPointerContent.
func (v *InterfaceListOfListOfListsFieldWithPointerTopic) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface InterfaceListOfListOfListsFieldWithPointerContent.
func (v *InterfaceListOfListOfListsFieldWithPointerTopic) GetId() *testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface InterfaceListOfListOfListsFieldWithPointerContent.
func (v *InterfaceListOfListOfListsFieldWithPointerTopic) GetName() *string { return v.Name }

func __unmarshalInterfaceListOfListOfListsFieldWithPointerContent(v *InterfaceListOfListOfListsFieldWithPointerContent, m json.RawMessage) error {
	if string(m) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(m, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(InterfaceListOfListOfListsFieldWithPointerArticle)
		return json.Unmarshal(m, *v)
	case "Video":
		*v = new(InterfaceListOfListOfListsFieldWithPointerVideo)
		return json.Unmarshal(m, *v)
	case "Topic":
		*v = new(InterfaceListOfListOfListsFieldWithPointerTopic)
		return json.Unmarshal(m, *v)
	case "":
		return fmt.Errorf(
			"Response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`Unexpected concrete type for InterfaceListOfListOfListsFieldWithPointerContent: "%v"`, tn.TypeName)
	}
}

// InterfaceListOfListOfListsFieldWithPointerTopic includes the requested fields of the GraphQL type Topic.
type InterfaceListOfListOfListsFieldWithPointerTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   *testutil.ID `json:"id"`
	Name *string      `json:"name"`
}

// InterfaceListOfListOfListsFieldWithPointerVideo includes the requested fields of the GraphQL type Video.
type InterfaceListOfListOfListsFieldWithPointerVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   *testutil.ID `json:"id"`
	Name *string      `json:"name"`
}

func InterfaceListOfListOfListsField(
	client graphql.Client,
) (*InterfaceListOfListOfListsFieldResponse, error) {
	var err error

	var retval InterfaceListOfListOfListsFieldResponse
	err = client.MakeRequest(
		nil,
		"InterfaceListOfListOfListsField",
		`
query InterfaceListOfListOfListsField {
	listOfListsOfListsOfContent {
		__typename
		id
		name
	}
	withPointer: listOfListsOfListsOfContent {
		__typename
		id
		name
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}

