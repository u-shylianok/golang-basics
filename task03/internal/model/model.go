package model

type Config struct {
	Tests []Test `xml:"Tests>Test" json:"Тесты"`
}

type Test struct {
	Name            string           `xml:"Name,attr" json:"Название теста"`
	Path            string           `xml:"Path" json:"Путь к файлу"`
	AnswerOptions   []AnswerOption   `xml:"AnswerOptions>Answer" json:"Варианты ответов"`
	Interpretations []Interpretation `xml:"Interpretations>Interpretation" json:"Интерпретация"`
}

type AnswerOption struct {
	Option       string `xml:"Option" json:"Вариант ответа"`
	DirectValue  int    `xml:"DirectValue" json:"Прямое значение"`
	ReverseValue int    `xml:"ReverseValue" json:"Обратное значение"`
}

type Interpretation struct {
	Scale             string `xml:"Scale,attr" json:"Шкала"`
	DirectStatements  []int  `xml:"DirectStatements>State" json:"Прямые утверждения"`
	ReverseStatements []int  `xml:"ReverseStatements>State" json:"Обратные утверждения"`
}
