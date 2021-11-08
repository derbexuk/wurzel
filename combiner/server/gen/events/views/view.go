// Code generated by goa v3.5.2, DO NOT EDIT.
//
// events views
//
// Command:
// $ goa gen github.com/derbexuk/poieventservice/server/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// CspaceEvent is the viewed result type that is projected based on a view.
type CspaceEvent struct {
	// Type to project
	Projected *CspaceEventView
	// View to render
	View string
}

// CspaceEventCollection is the viewed result type that is projected based on a
// view.
type CspaceEventCollection struct {
	// Type to project
	Projected CspaceEventCollectionView
	// View to render
	View string
}

// CspaceEventView is a type that runs validations on a projected type.
type CspaceEventView struct {
	// ID of Event
	ID *string
	// Event title/name
	Title       *string
	Description *string
	Deactivated *bool
	Start       *string
	End         *string
	// Hash of application specific properties
	Properties map[string]string
	// refs
	Refs []string
	Path *string
}

// CspaceEventCollectionView is a type that runs validations on a projected
// type.
type CspaceEventCollectionView []*CspaceEventView

var (
	// CspaceEventMap is a map indexing the attribute names of CspaceEvent by view
	// name.
	CspaceEventMap = map[string][]string{
		"default": {
			"id",
			"title",
			"description",
			"deactivated",
			"start",
			"end",
			"properties",
			"refs",
			"path",
		},
		"summary": {
			"id",
			"title",
			"start",
			"end",
			"path",
		},
	}
	// CspaceEventCollectionMap is a map indexing the attribute names of
	// CspaceEventCollection by view name.
	CspaceEventCollectionMap = map[string][]string{
		"default": {
			"id",
			"title",
			"description",
			"deactivated",
			"start",
			"end",
			"properties",
			"refs",
			"path",
		},
		"summary": {
			"id",
			"title",
			"start",
			"end",
			"path",
		},
	}
)

// ValidateCspaceEvent runs the validations defined on the viewed result type
// CspaceEvent.
func ValidateCspaceEvent(result *CspaceEvent) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateCspaceEventView(result.Projected)
	case "summary":
		err = ValidateCspaceEventViewSummary(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "summary"})
	}
	return
}

// ValidateCspaceEventCollection runs the validations defined on the viewed
// result type CspaceEventCollection.
func ValidateCspaceEventCollection(result CspaceEventCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateCspaceEventCollectionView(result.Projected)
	case "summary":
		err = ValidateCspaceEventCollectionViewSummary(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "summary"})
	}
	return
}

// ValidateCspaceEventView runs the validations defined on CspaceEventView
// using the "default" view.
func ValidateCspaceEventView(result *CspaceEventView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "result"))
	}
	if result.Start == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("start", "result"))
	}
	if result.Start != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.start", *result.Start, goa.FormatDateTime))
	}
	if result.End != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.end", *result.End, goa.FormatDateTime))
	}
	return
}

// ValidateCspaceEventViewSummary runs the validations defined on
// CspaceEventView using the "summary" view.
func ValidateCspaceEventViewSummary(result *CspaceEventView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "result"))
	}
	if result.Start == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("start", "result"))
	}
	if result.Start != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.start", *result.Start, goa.FormatDateTime))
	}
	if result.End != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.end", *result.End, goa.FormatDateTime))
	}
	return
}

// ValidateCspaceEventCollectionView runs the validations defined on
// CspaceEventCollectionView using the "default" view.
func ValidateCspaceEventCollectionView(result CspaceEventCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateCspaceEventView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateCspaceEventCollectionViewSummary runs the validations defined on
// CspaceEventCollectionView using the "summary" view.
func ValidateCspaceEventCollectionViewSummary(result CspaceEventCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateCspaceEventViewSummary(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}
