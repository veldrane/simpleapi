// Code generated by goa v3.2.5, DO NOT EDIT.
//
// stuff views
//
// Command:
// $ goa gen simpleapi/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// StoredArticleCollection is the viewed result type that is projected based on
// a view.
type StoredArticleCollection struct {
	// Type to project
	Projected StoredArticleCollectionView
	// View to render
	View string
}

// StoredArticle is the viewed result type that is projected based on a view.
type StoredArticle struct {
	// Type to project
	Projected *StoredArticleView
	// View to render
	View string
}

// StoredArticleCollectionView is a type that runs validations on a projected
// type.
type StoredArticleCollectionView []*StoredArticleView

// StoredArticleView is a type that runs validations on a projected type.
type StoredArticleView struct {
	// ID is the unique id of the article
	ID      *int
	Title   *string
	Author  *string
	Desc    *string
	Content *string
	Admin   *string
}

var (
	// StoredArticleCollectionMap is a map of attribute names in result type
	// StoredArticleCollection indexed by view name.
	StoredArticleCollectionMap = map[string][]string{
		"default": []string{
			"id",
			"title",
			"author",
			"desc",
			"content",
			"admin",
		},
	}
	// StoredArticleMap is a map of attribute names in result type StoredArticle
	// indexed by view name.
	StoredArticleMap = map[string][]string{
		"default": []string{
			"id",
			"title",
			"author",
			"desc",
			"content",
			"admin",
		},
	}
)

// ValidateStoredArticleCollection runs the validations defined on the viewed
// result type StoredArticleCollection.
func ValidateStoredArticleCollection(result StoredArticleCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStoredArticleCollectionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateStoredArticle runs the validations defined on the viewed result type
// StoredArticle.
func ValidateStoredArticle(result *StoredArticle) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStoredArticleView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateStoredArticleCollectionView runs the validations defined on
// StoredArticleCollectionView using the "default" view.
func ValidateStoredArticleCollectionView(result StoredArticleCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateStoredArticleView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStoredArticleView runs the validations defined on StoredArticleView
// using the "default" view.
func ValidateStoredArticleView(result *StoredArticleView) (err error) {

	return
}