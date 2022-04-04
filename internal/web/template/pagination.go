package template

import (
	"fmt"
	"math"
)

// Pagination is a pagination element that can be added to a webpage
type Pagination []PaginationNode

// PaginationNode is an element in a pagination element
type PaginationNode struct {
	Text        string
	DisplayHTML string
	HRef        string

	Active   bool
	Disabled bool
}

// PaginationConfig
type PaginationConfig struct {
	Count         int    // item count
	DisplayCount  int    // how many items to display per page
	HRef          string // href to add query to
	HRefCount     int    // count to include in the href, if 0 no count is added
	MaxPagination int    // the max number of pages to show
	Page          int    // current page
}

// MakePagination creates a pagination element from the provided parameters
func MakePagination(page, count, displayCount int, href string, hrefCount int) Pagination {
	displayItems := 5
	pages := int(math.Ceil(float64(count) / float64(displayCount)))
	startingNumber := 1

	if pages < displayItems {
		// less than
		displayItems = pages
	} else if page > pages-displayItems/2 {
		// end of the
		startingNumber = pages - displayItems + 1
	} else if page > displayItems/2 {
		// center active
		startingNumber = page - displayItems/2
	}

	var items Pagination

	// previous button
	prevItem := PaginationNode{
		Text:        "Previous",
		DisplayHTML: "<i class=\"fas fa-caret-left\"></i>",
	}
	if page == 1 {
		prevItem.Disabled = true
	} else if hrefCount > 0 {
		prevItem.HRef = fmt.Sprintf("%s?page=%d&count=%d", href, page-1, hrefCount)
	} else {
		prevItem.HRef = fmt.Sprintf("%s?page=%d", href, page-1)
	}
	items = append(items, prevItem)

	// add pages
	for i := 0; i < displayItems; i++ {
		newItem := PaginationNode{
			Text: fmt.Sprintf("%d", startingNumber+i),
		}

		if page == startingNumber+i {
			newItem.Active = true
		} else if hrefCount > 0 {
			newItem.HRef = fmt.Sprintf("%s?page=%d&count=%d", href, startingNumber+i, hrefCount)
		} else {
			newItem.HRef = fmt.Sprintf("%s?page=%d", href, startingNumber+i)
		}

		items = append(items, newItem)
	}

	// next button
	nextItem := PaginationNode{
		Text:        "Next",
		DisplayHTML: "<i class=\"fas fa-caret-right\"></i>",
	}
	if page == count {
		nextItem.Disabled = true
	} else if hrefCount > 0 {
		nextItem.HRef = fmt.Sprintf("%s?page=%d&count=%d", href, page+1, hrefCount)
	} else {
		nextItem.HRef = fmt.Sprintf("%s?page=%d", href, page+1)
	}
	items = append(items, nextItem)

	return items
}
