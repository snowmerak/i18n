package i18n

type I18N[T any] struct {
	inner map[string]map[string]T
}

func New[T any]() *I18N[T] {
	return &I18N[T]{
		inner: make(map[string]map[string]T),
	}
}

func (i *I18N[T]) Set(nation, lang string, value T) {
	if _, ok := i.inner[nation]; !ok {
		i.inner[nation] = make(map[string]T)
	}
	i.inner[nation][lang] = value
}

func (i *I18N[T]) Get(nation, lang string) (t T, ok bool) {
	m, ok := i.inner[nation]
	if !ok {
		return t, false
	}
	t, ok = m[lang]
	if !ok {
		return t, false
	}
	return t, true
}
