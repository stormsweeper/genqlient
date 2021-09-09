package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"encoding/json"
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// ComplexNamedFragmentsResponse is returned by ComplexNamedFragments on success.
type ComplexNamedFragmentsResponse struct {
	QueryFragment `json:"-"`
}

func (v *ComplexNamedFragmentsResponse) UnmarshalJSON(b []byte) error {

	var firstPass struct {
		*ComplexNamedFragmentsResponse
		graphql.NoUnmarshalJSON
	}
	firstPass.ComplexNamedFragmentsResponse = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.QueryFragment)
	if err != nil {
		return err
	}
	return nil
}

// ContentFields includes the GraphQL fields of Content requested by the fragment ContentFields.
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
//
// ContentFields is implemented by the following types:
// ContentFieldsArticle
// ContentFieldsVideo
// ContentFieldsTopic
type ContentFields interface {
	implementsGraphQLInterfaceContentFields()
	// GetName returns the interface-field "name" from its implementation.
	GetName() string
	// GetUrl returns the interface-field "url" from its implementation.
	GetUrl() string
}

func (v *ContentFieldsArticle) implementsGraphQLInterfaceContentFields() {}

// GetName is a part of, and documented with, the interface ContentFields.
func (v *ContentFieldsArticle) GetName() string { return v.Name }

// GetUrl is a part of, and documented with, the interface ContentFields.
func (v *ContentFieldsArticle) GetUrl() string { return v.Url }

func (v *ContentFieldsVideo) implementsGraphQLInterfaceContentFields() {}

// GetName is a part of, and documented with, the interface ContentFields.
func (v *ContentFieldsVideo) GetName() string { return v.Name }

// GetUrl is a part of, and documented with, the interface ContentFields.
func (v *ContentFieldsVideo) GetUrl() string { return v.Url }

func (v *ContentFieldsTopic) implementsGraphQLInterfaceContentFields() {}

// GetName is a part of, and documented with, the interface ContentFields.
func (v *ContentFieldsTopic) GetName() string { return v.Name }

// GetUrl is a part of, and documented with, the interface ContentFields.
func (v *ContentFieldsTopic) GetUrl() string { return v.Url }

func __unmarshalContentFields(v *ContentFields, m json.RawMessage) error {
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
		*v = new(ContentFieldsArticle)
		return json.Unmarshal(m, *v)
	case "Video":
		*v = new(ContentFieldsVideo)
		return json.Unmarshal(m, *v)
	case "Topic":
		*v = new(ContentFieldsTopic)
		return json.Unmarshal(m, *v)
	case "":
		return fmt.Errorf(
			"Response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`Unexpected concrete type for ContentFields: "%v"`, tn.TypeName)
	}
}

// ContentFields includes the GraphQL fields of Article requested by the fragment ContentFields.
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type ContentFieldsArticle struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// ContentFields includes the GraphQL fields of Topic requested by the fragment ContentFields.
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type ContentFieldsTopic struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// ContentFields includes the GraphQL fields of Video requested by the fragment ContentFields.
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type ContentFieldsVideo struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// InnerQueryFragment includes the GraphQL fields of Query requested by the fragment InnerQueryFragment.
// The GraphQL type's documentation follows.
//
// Query's description is probably ignored by almost all callers.
type InnerQueryFragment struct {
	RandomItem InnerQueryFragmentRandomItemContent     `json:"-"`
	RandomLeaf InnerQueryFragmentRandomLeafLeafContent `json:"-"`
	OtherLeaf  InnerQueryFragmentOtherLeafLeafContent  `json:"-"`
}

func (v *InnerQueryFragment) UnmarshalJSON(b []byte) error {

	var firstPass struct {
		*InnerQueryFragment
		RandomItem json.RawMessage `json:"randomItem"`
		RandomLeaf json.RawMessage `json:"randomLeaf"`
		OtherLeaf  json.RawMessage `json:"otherLeaf"`
		graphql.NoUnmarshalJSON
	}
	firstPass.InnerQueryFragment = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		target := &v.RandomItem
		raw := firstPass.RandomItem
		err = __unmarshalInnerQueryFragmentRandomItemContent(
			target, raw)
		if err != nil {
			return fmt.Errorf(
				"Unable to unmarshal InnerQueryFragment.RandomItem: %w", err)
		}
	}

	{
		target := &v.RandomLeaf
		raw := firstPass.RandomLeaf
		err = __unmarshalInnerQueryFragmentRandomLeafLeafContent(
			target, raw)
		if err != nil {
			return fmt.Errorf(
				"Unable to unmarshal InnerQueryFragment.RandomLeaf: %w", err)
		}
	}

	{
		target := &v.OtherLeaf
		raw := firstPass.OtherLeaf
		err = __unmarshalInnerQueryFragmentOtherLeafLeafContent(
			target, raw)
		if err != nil {
			return fmt.Errorf(
				"Unable to unmarshal InnerQueryFragment.OtherLeaf: %w", err)
		}
	}
	return nil
}

// InnerQueryFragmentOtherLeafArticle includes the requested fields of the GraphQL type Article.
type InnerQueryFragmentOtherLeafArticle struct {
	Typename             string `json:"__typename"`
	ContentFieldsArticle `json:"-"`
}

func (v *InnerQueryFragmentOtherLeafArticle) UnmarshalJSON(b []byte) error {

	var firstPass struct {
		*InnerQueryFragmentOtherLeafArticle
		graphql.NoUnmarshalJSON
	}
	firstPass.InnerQueryFragmentOtherLeafArticle = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.ContentFieldsArticle)
	if err != nil {
		return err
	}
	return nil
}

// InnerQueryFragmentOtherLeafLeafContent includes the requested fields of the GraphQL interface LeafContent.
//
// InnerQueryFragmentOtherLeafLeafContent is implemented by the following types:
// InnerQueryFragmentOtherLeafArticle
// InnerQueryFragmentOtherLeafVideo
// The GraphQL type's documentation follows.
//
// LeafContent represents content items that can't have child-nodes.
type InnerQueryFragmentOtherLeafLeafContent interface {
	implementsGraphQLInterfaceInnerQueryFragmentOtherLeafLeafContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
}

func (v *InnerQueryFragmentOtherLeafArticle) implementsGraphQLInterfaceInnerQueryFragmentOtherLeafLeafContent() {
}

// GetTypename is a part of, and documented with, the interface InnerQueryFragmentOtherLeafLeafContent.
func (v *InnerQueryFragmentOtherLeafArticle) GetTypename() string { return v.Typename }

func (v *InnerQueryFragmentOtherLeafVideo) implementsGraphQLInterfaceInnerQueryFragmentOtherLeafLeafContent() {
}

// GetTypename is a part of, and documented with, the interface InnerQueryFragmentOtherLeafLeafContent.
func (v *InnerQueryFragmentOtherLeafVideo) GetTypename() string { return v.Typename }

func __unmarshalInnerQueryFragmentOtherLeafLeafContent(v *InnerQueryFragmentOtherLeafLeafContent, m json.RawMessage) error {
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
		*v = new(InnerQueryFragmentOtherLeafArticle)
		return json.Unmarshal(m, *v)
	case "Video":
		*v = new(InnerQueryFragmentOtherLeafVideo)
		return json.Unmarshal(m, *v)
	case "":
		return fmt.Errorf(
			"Response was missing LeafContent.__typename")
	default:
		return fmt.Errorf(
			`Unexpected concrete type for InnerQueryFragmentOtherLeafLeafContent: "%v"`, tn.TypeName)
	}
}

// InnerQueryFragmentOtherLeafVideo includes the requested fields of the GraphQL type Video.
type InnerQueryFragmentOtherLeafVideo struct {
	Typename           string `json:"__typename"`
	MoreVideoFields    `json:"-"`
	ContentFieldsVideo `json:"-"`
}

func (v *InnerQueryFragmentOtherLeafVideo) UnmarshalJSON(b []byte) error {

	var firstPass struct {
		*InnerQueryFragmentOtherLeafVideo
		graphql.NoUnmarshalJSON
	}
	firstPass.InnerQueryFragmentOtherLeafVideo = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.MoreVideoFields)
	if err != nil {
		return err
	}
	err = json.Unmarshal(
		b, &v.ContentFieldsVideo)
	if err != nil {
		return err
	}
	return nil
}

// InnerQueryFragmentRandomItemArticle includes the requested fields of the GraphQL type Article.
type InnerQueryFragmentRandomItemArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id                   testutil.ID `json:"id"`
	Name                 string      `json:"name"`
	ContentFieldsArticle `json:"-"`
}

func (v *InnerQueryFragmentRandomItemArticle) UnmarshalJSON(b []byte) error {

	var firstPass struct {
		*InnerQueryFragmentRandomItemArticle
		graphql.NoUnmarshalJSON
	}
	firstPass.InnerQueryFragmentRandomItemArticle = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.ContentFieldsArticle)
	if err != nil {
		return err
	}
	return nil
}

// InnerQueryFragmentRandomItemContent includes the requested fields of the GraphQL interface Content.
//
// InnerQueryFragmentRandomItemContent is implemented by the following types:
// InnerQueryFragmentRandomItemArticle
// InnerQueryFragmentRandomItemVideo
// InnerQueryFragmentRandomItemTopic
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InnerQueryFragmentRandomItemContent interface {
	implementsGraphQLInterfaceInnerQueryFragmentRandomItemContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() testutil.ID
	// GetName returns the interface-field "name" from its implementation.
	GetName() string
	ContentFields
}

func (v *InnerQueryFragmentRandomItemArticle) implementsGraphQLInterfaceInnerQueryFragmentRandomItemContent() {
}

// GetTypename is a part of, and documented with, the interface InnerQueryFragmentRandomItemContent.
func (v *InnerQueryFragmentRandomItemArticle) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface InnerQueryFragmentRandomItemContent.
func (v *InnerQueryFragmentRandomItemArticle) GetId() testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface InnerQueryFragmentRandomItemContent.
func (v *InnerQueryFragmentRandomItemArticle) GetName() string { return v.Name }

func (v *InnerQueryFragmentRandomItemVideo) implementsGraphQLInterfaceInnerQueryFragmentRandomItemContent() {
}

// GetTypename is a part of, and documented with, the interface InnerQueryFragmentRandomItemContent.
func (v *InnerQueryFragmentRandomItemVideo) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface InnerQueryFragmentRandomItemContent.
func (v *InnerQueryFragmentRandomItemVideo) GetId() testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface InnerQueryFragmentRandomItemContent.
func (v *InnerQueryFragmentRandomItemVideo) GetName() string { return v.Name }

func (v *InnerQueryFragmentRandomItemTopic) implementsGraphQLInterfaceInnerQueryFragmentRandomItemContent() {
}

// GetTypename is a part of, and documented with, the interface InnerQueryFragmentRandomItemContent.
func (v *InnerQueryFragmentRandomItemTopic) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface InnerQueryFragmentRandomItemContent.
func (v *InnerQueryFragmentRandomItemTopic) GetId() testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface InnerQueryFragmentRandomItemContent.
func (v *InnerQueryFragmentRandomItemTopic) GetName() string { return v.Name }

func __unmarshalInnerQueryFragmentRandomItemContent(v *InnerQueryFragmentRandomItemContent, m json.RawMessage) error {
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
		*v = new(InnerQueryFragmentRandomItemArticle)
		return json.Unmarshal(m, *v)
	case "Video":
		*v = new(InnerQueryFragmentRandomItemVideo)
		return json.Unmarshal(m, *v)
	case "Topic":
		*v = new(InnerQueryFragmentRandomItemTopic)
		return json.Unmarshal(m, *v)
	case "":
		return fmt.Errorf(
			"Response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`Unexpected concrete type for InnerQueryFragmentRandomItemContent: "%v"`, tn.TypeName)
	}
}

// InnerQueryFragmentRandomItemTopic includes the requested fields of the GraphQL type Topic.
type InnerQueryFragmentRandomItemTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id                 testutil.ID `json:"id"`
	Name               string      `json:"name"`
	ContentFieldsTopic `json:"-"`
}

func (v *InnerQueryFragmentRandomItemTopic) UnmarshalJSON(b []byte) error {

	var firstPass struct {
		*InnerQueryFragmentRandomItemTopic
		graphql.NoUnmarshalJSON
	}
	firstPass.InnerQueryFragmentRandomItemTopic = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.ContentFieldsTopic)
	if err != nil {
		return err
	}
	return nil
}

// InnerQueryFragmentRandomItemVideo includes the requested fields of the GraphQL type Video.
type InnerQueryFragmentRandomItemVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id                 testutil.ID `json:"id"`
	Name               string      `json:"name"`
	VideoFields        `json:"-"`
	ContentFieldsVideo `json:"-"`
}

func (v *InnerQueryFragmentRandomItemVideo) UnmarshalJSON(b []byte) error {

	var firstPass struct {
		*InnerQueryFragmentRandomItemVideo
		graphql.NoUnmarshalJSON
	}
	firstPass.InnerQueryFragmentRandomItemVideo = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.VideoFields)
	if err != nil {
		return err
	}
	err = json.Unmarshal(
		b, &v.ContentFieldsVideo)
	if err != nil {
		return err
	}
	return nil
}

// InnerQueryFragmentRandomLeafArticle includes the requested fields of the GraphQL type Article.
type InnerQueryFragmentRandomLeafArticle struct {
	Typename             string `json:"__typename"`
	ContentFieldsArticle `json:"-"`
}

func (v *InnerQueryFragmentRandomLeafArticle) UnmarshalJSON(b []byte) error {

	var firstPass struct {
		*InnerQueryFragmentRandomLeafArticle
		graphql.NoUnmarshalJSON
	}
	firstPass.InnerQueryFragmentRandomLeafArticle = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.ContentFieldsArticle)
	if err != nil {
		return err
	}
	return nil
}

// InnerQueryFragmentRandomLeafLeafContent includes the requested fields of the GraphQL interface LeafContent.
//
// InnerQueryFragmentRandomLeafLeafContent is implemented by the following types:
// InnerQueryFragmentRandomLeafArticle
// InnerQueryFragmentRandomLeafVideo
// The GraphQL type's documentation follows.
//
// LeafContent represents content items that can't have child-nodes.
type InnerQueryFragmentRandomLeafLeafContent interface {
	implementsGraphQLInterfaceInnerQueryFragmentRandomLeafLeafContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
}

func (v *InnerQueryFragmentRandomLeafArticle) implementsGraphQLInterfaceInnerQueryFragmentRandomLeafLeafContent() {
}

// GetTypename is a part of, and documented with, the interface InnerQueryFragmentRandomLeafLeafContent.
func (v *InnerQueryFragmentRandomLeafArticle) GetTypename() string { return v.Typename }

func (v *InnerQueryFragmentRandomLeafVideo) implementsGraphQLInterfaceInnerQueryFragmentRandomLeafLeafContent() {
}

// GetTypename is a part of, and documented with, the interface InnerQueryFragmentRandomLeafLeafContent.
func (v *InnerQueryFragmentRandomLeafVideo) GetTypename() string { return v.Typename }

func __unmarshalInnerQueryFragmentRandomLeafLeafContent(v *InnerQueryFragmentRandomLeafLeafContent, m json.RawMessage) error {
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
		*v = new(InnerQueryFragmentRandomLeafArticle)
		return json.Unmarshal(m, *v)
	case "Video":
		*v = new(InnerQueryFragmentRandomLeafVideo)
		return json.Unmarshal(m, *v)
	case "":
		return fmt.Errorf(
			"Response was missing LeafContent.__typename")
	default:
		return fmt.Errorf(
			`Unexpected concrete type for InnerQueryFragmentRandomLeafLeafContent: "%v"`, tn.TypeName)
	}
}

// InnerQueryFragmentRandomLeafVideo includes the requested fields of the GraphQL type Video.
type InnerQueryFragmentRandomLeafVideo struct {
	Typename           string `json:"__typename"`
	VideoFields        `json:"-"`
	MoreVideoFields    `json:"-"`
	ContentFieldsVideo `json:"-"`
}

func (v *InnerQueryFragmentRandomLeafVideo) UnmarshalJSON(b []byte) error {

	var firstPass struct {
		*InnerQueryFragmentRandomLeafVideo
		graphql.NoUnmarshalJSON
	}
	firstPass.InnerQueryFragmentRandomLeafVideo = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.VideoFields)
	if err != nil {
		return err
	}
	err = json.Unmarshal(
		b, &v.MoreVideoFields)
	if err != nil {
		return err
	}
	err = json.Unmarshal(
		b, &v.ContentFieldsVideo)
	if err != nil {
		return err
	}
	return nil
}

// MoreVideoFields includes the GraphQL fields of Video requested by the fragment MoreVideoFields.
type MoreVideoFields struct {
	// ID is documented in the Content interface.
	Id     *testutil.ID                `json:"id"`
	Parent *MoreVideoFieldsParentTopic `json:"parent"`
}

// MoreVideoFieldsParentTopic includes the requested fields of the GraphQL type Topic.
type MoreVideoFieldsParentTopic struct {
	Name               *string `json:"name"`
	Url                *string `json:"url"`
	ContentFieldsTopic `json:"-"`
	Children           []MoreVideoFieldsParentTopicChildrenContent `json:"-"`
}

func (v *MoreVideoFieldsParentTopic) UnmarshalJSON(b []byte) error {

	var firstPass struct {
		*MoreVideoFieldsParentTopic
		Children []json.RawMessage `json:"children"`
		graphql.NoUnmarshalJSON
	}
	firstPass.MoreVideoFieldsParentTopic = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.ContentFieldsTopic)
	if err != nil {
		return err
	}

	{
		target := &v.Children
		raw := firstPass.Children
		*target = make(
			[]MoreVideoFieldsParentTopicChildrenContent,
			len(raw))
		for i, raw := range raw {
			target := &(*target)[i]
			err = __unmarshalMoreVideoFieldsParentTopicChildrenContent(
				target, raw)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal MoreVideoFieldsParentTopic.Children: %w", err)
			}
		}
	}
	return nil
}

// MoreVideoFieldsParentTopicChildrenArticle includes the requested fields of the GraphQL type Article.
type MoreVideoFieldsParentTopicChildrenArticle struct {
	Typename *string `json:"__typename"`
}

// MoreVideoFieldsParentTopicChildrenContent includes the requested fields of the GraphQL interface Content.
//
// MoreVideoFieldsParentTopicChildrenContent is implemented by the following types:
// MoreVideoFieldsParentTopicChildrenArticle
// MoreVideoFieldsParentTopicChildrenVideo
// MoreVideoFieldsParentTopicChildrenTopic
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type MoreVideoFieldsParentTopicChildrenContent interface {
	implementsGraphQLInterfaceMoreVideoFieldsParentTopicChildrenContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() *string
}

func (v *MoreVideoFieldsParentTopicChildrenArticle) implementsGraphQLInterfaceMoreVideoFieldsParentTopicChildrenContent() {
}

// GetTypename is a part of, and documented with, the interface MoreVideoFieldsParentTopicChildrenContent.
func (v *MoreVideoFieldsParentTopicChildrenArticle) GetTypename() *string { return v.Typename }

func (v *MoreVideoFieldsParentTopicChildrenVideo) implementsGraphQLInterfaceMoreVideoFieldsParentTopicChildrenContent() {
}

// GetTypename is a part of, and documented with, the interface MoreVideoFieldsParentTopicChildrenContent.
func (v *MoreVideoFieldsParentTopicChildrenVideo) GetTypename() *string { return v.Typename }

func (v *MoreVideoFieldsParentTopicChildrenTopic) implementsGraphQLInterfaceMoreVideoFieldsParentTopicChildrenContent() {
}

// GetTypename is a part of, and documented with, the interface MoreVideoFieldsParentTopicChildrenContent.
func (v *MoreVideoFieldsParentTopicChildrenTopic) GetTypename() *string { return v.Typename }

func __unmarshalMoreVideoFieldsParentTopicChildrenContent(v *MoreVideoFieldsParentTopicChildrenContent, m json.RawMessage) error {
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
		*v = new(MoreVideoFieldsParentTopicChildrenArticle)
		return json.Unmarshal(m, *v)
	case "Video":
		*v = new(MoreVideoFieldsParentTopicChildrenVideo)
		return json.Unmarshal(m, *v)
	case "Topic":
		*v = new(MoreVideoFieldsParentTopicChildrenTopic)
		return json.Unmarshal(m, *v)
	case "":
		return fmt.Errorf(
			"Response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`Unexpected concrete type for MoreVideoFieldsParentTopicChildrenContent: "%v"`, tn.TypeName)
	}
}

// MoreVideoFieldsParentTopicChildrenTopic includes the requested fields of the GraphQL type Topic.
type MoreVideoFieldsParentTopicChildrenTopic struct {
	Typename *string `json:"__typename"`
}

// MoreVideoFieldsParentTopicChildrenVideo includes the requested fields of the GraphQL type Video.
type MoreVideoFieldsParentTopicChildrenVideo struct {
	Typename    *string `json:"__typename"`
	VideoFields `json:"-"`
}

func (v *MoreVideoFieldsParentTopicChildrenVideo) UnmarshalJSON(b []byte) error {

	var firstPass struct {
		*MoreVideoFieldsParentTopicChildrenVideo
		graphql.NoUnmarshalJSON
	}
	firstPass.MoreVideoFieldsParentTopicChildrenVideo = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.VideoFields)
	if err != nil {
		return err
	}
	return nil
}

// QueryFragment includes the GraphQL fields of Query requested by the fragment QueryFragment.
// The GraphQL type's documentation follows.
//
// Query's description is probably ignored by almost all callers.
type QueryFragment struct {
	InnerQueryFragment `json:"-"`
}

func (v *QueryFragment) UnmarshalJSON(b []byte) error {

	var firstPass struct {
		*QueryFragment
		graphql.NoUnmarshalJSON
	}
	firstPass.QueryFragment = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.InnerQueryFragment)
	if err != nil {
		return err
	}
	return nil
}

// VideoFields includes the GraphQL fields of Video requested by the fragment VideoFields.
type VideoFields struct {
	// ID is documented in the Content interface.
	Id                 testutil.ID          `json:"id"`
	Name               string               `json:"name"`
	Url                string               `json:"url"`
	Duration           int                  `json:"duration"`
	Thumbnail          VideoFieldsThumbnail `json:"thumbnail"`
	ContentFieldsVideo `json:"-"`
}

func (v *VideoFields) UnmarshalJSON(b []byte) error {

	var firstPass struct {
		*VideoFields
		graphql.NoUnmarshalJSON
	}
	firstPass.VideoFields = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.ContentFieldsVideo)
	if err != nil {
		return err
	}
	return nil
}

// VideoFieldsThumbnail includes the requested fields of the GraphQL type Thumbnail.
type VideoFieldsThumbnail struct {
	Id testutil.ID `json:"id"`
}

func ComplexNamedFragments(
	client graphql.Client,
) (*ComplexNamedFragmentsResponse, error) {
	var err error

	var retval ComplexNamedFragmentsResponse
	err = client.MakeRequest(
		nil,
		"ComplexNamedFragments",
		`
query ComplexNamedFragments {
	... on Query {
		... QueryFragment
	}
}
fragment QueryFragment on Query {
	... InnerQueryFragment
}
fragment InnerQueryFragment on Query {
	randomItem {
		__typename
		id
		name
		... VideoFields
		... ContentFields
	}
	randomLeaf {
		__typename
		... VideoFields
		... MoreVideoFields
		... ContentFields
	}
	otherLeaf: randomLeaf {
		__typename
		... on Video {
			... MoreVideoFields
			... ContentFields
		}
		... ContentFields
	}
}
fragment VideoFields on Video {
	id
	name
	url
	duration
	thumbnail {
		id
	}
	... ContentFields
}
fragment ContentFields on Content {
	name
	url
}
fragment MoreVideoFields on Video {
	id
	parent {
		name
		url
		... ContentFields
		children {
			__typename
			... VideoFields
		}
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}

