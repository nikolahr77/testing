package testpagination

import "fmt"

// Pages contains a slice of strings that will hold the information
type Pages struct {
	Elements []string
	// PageSize specifies the number of elements that a single page will hold
	PageSize int
	current int
}

// FirstPage returns the elements on the first page and sets the current page to first
func (p *Pages) FirstPage() []string {
	p.current = 1
	s := p.Elements[:p.PageSize]
	return s
}

// LastPage returns the elements on the last page and sets the current page to last
func (p *Pages) LastPage() []string {
	p.current = len(p.Elements) / p.PageSize
	if len(p.Elements)%p.PageSize != 0 {
		p.current++
	}
	s := p.Elements[len(p.Elements)-len(p.Elements)%p.PageSize:]
	fmt.Println(s)
	return s
}

// GetCurrentPageNumber returns the current page that we are on
func (p *Pages) GetCurrentPageNumber() int {
	return p.current
}

// HasNext checks if there is a page after the one we are currently on
func (p Pages) HasNext() bool {
	totalPages := len(p.Elements) / p.PageSize
	if len(p.Elements)%p.PageSize != 0 {
		totalPages++
	}
	if p.current == totalPages {
		return false
	}
	return true
}

// HasPrevious checks if there is a page before the one we are currently on
func (p Pages) HasPrevious() bool {
	if p.current <= 1 {
		return false
	}
	return true
}

// Next returns the page after the one we are currently on
func (p *Pages) Next() []string {
	var s []string
	if !p.HasNext() {
		return s
	}
	for i := p.current * p.PageSize; i < p.current*p.PageSize+p.PageSize; i++ {
		s = append(s, p.Elements[i])
		if i == len(p.Elements)-1 {
			break
		}
	}
	return s
}

// Previous returns the page before the one we are currently on
func (p *Pages) Previous() []string {
	var s []string
	if !p.HasPrevious() {
		return s
	}
	for i := p.current*p.PageSize - p.PageSize*2; i < p.current*p.PageSize-p.PageSize; i++ {
		s = append(s, p.Elements[i])
		if i == len(p.Elements)+1 {
			break
		}
	}
	return s
}
