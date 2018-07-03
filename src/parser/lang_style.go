package parser

import "github.com/YuriyLisovskiy/sloc/src/models"

func singleComment(c string) models.Comment {
	return models.Comment{Start: c, End: "\n"}
}

func language(lang string, slc, mlc models.Comment) models.AvailableLang {
	return models.AvailableLang{Name: lang, SingleLineComment: slc, MultiLineComment: mlc}
}

func cStyle(lang string) models.AvailableLang {
	return language(lang, clangSComment, clangMComment)
}

func shStyle(lang string) models.AvailableLang {
	return language(lang, shSComment, noneComment)
}

func noComments(lang string) models.AvailableLang {
	return language(lang, noneComment, noneComment)
}

func htmlStyle(lang string) models.AvailableLang {
	return language(lang, noneComment, xmlComment)
}

func mlStyle(lang string) models.AvailableLang {
	return language(lang, noneComment, mlMComment)
}

func prologStyle(lang string) models.AvailableLang {
	return language(lang, singleComment("%"), clangMComment)
}
