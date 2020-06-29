package simple_factory

type translator interface {
	translate() string
}

type EnglishTranslator struct {

}

func (t EnglishTranslator) translate() string {
	return "en"
}

type ChineseTranslator struct {

}

func (t ChineseTranslator) translate () string {
	return "cn"
}

func NewTranslator(lan string) translator {
	switch lan {
	case "en":
		return &EnglishTranslator{}
	case "cn":
		return new(ChineseTranslator)
	default:
		panic("not defined translator")
	}
}



