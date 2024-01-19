package main

import "fmt"

type ASensitiveWordsFilter struct {
}

func (a ASensitiveWordsFilter) FilterSexyWords(text string) string {
	return text
}

func (a ASensitiveWordsFilter) FilterPoliticalWords(text string) string {
	return text
}

type BSensitiveWordsFilter struct {
}

func (b BSensitiveWordsFilter) Filter(text string) string {
	return text
}

type CSensitiveWordsFilter struct {
}

func (c CSensitiveWordsFilter) Filter(text string, mask string) string {
	return mask
}

// If there's no Adaptor
type RiskManagement struct {
	a *ASensitiveWordsFilter
	b *BSensitiveWordsFilter
	c *CSensitiveWordsFilter
}

func (r *RiskManagement) filterSensitiveWords(text string) string {
	maskedText := r.a.FilterSexyWords(text)
	maskedText = r.a.FilterPoliticalWords(maskedText)
	maskedText = r.b.Filter(maskedText)
	maskedText = r.c.Filter(maskedText, "***")
	return maskedText
}

// Leveraging Adaptor
type SensitiveWordsFilter interface {
	filter(text string) string
}

func (a ASensitiveWordsFilter) filter(text string) string {
	maskedText := a.FilterSexyWords(text)
	return a.FilterPoliticalWords(maskedText)
}

func (b BSensitiveWordsFilter) filter(text string) string {
	return b.Filter(text)
}

func (c CSensitiveWordsFilter) filter(text string) string {
	return c.Filter(text, "***")
}

type RiskManagementWithAdaptor struct {
	Filters []SensitiveWordsFilter
}

func (r *RiskManagementWithAdaptor) RegisterFilter(filter SensitiveWordsFilter) {
	r.Filters = append(r.Filters, filter)
}

// 把不同的function都统一成一个接口的调用
func (r *RiskManagementWithAdaptor) FilterSensitiveWords(text string) string {
	maskedText := text
	for _, filter := range r.Filters {
		maskedText = filter.filter(maskedText)
	}
	return maskedText
}

func main() {
	r := RiskManagementWithAdaptor{Filters: []SensitiveWordsFilter{}}
	r.RegisterFilter(&ASensitiveWordsFilter{})
	r.RegisterFilter(&BSensitiveWordsFilter{})
	r.RegisterFilter(&CSensitiveWordsFilter{})

	fmt.Println(r.FilterSensitiveWords("abc"))
}
