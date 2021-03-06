package languages

// EnglishTranslation is the English translation of Wrap
var EnglishTranslation = Translation{
	Language: English,

	SceneTags: []string{
		"int ", "ext ", "est ", "int./ext ", "int/ext ", "i/e ", "i./e ",
		"int.", "ext.", "est.", "int./ext.", "int/ext.", "i/e.", "i./e.",
	},

	StageplaySceneTags: []string{
		"scene ",
	},

	TransitionTags: []string{
		" TO:",
	},

	BeginActTags: []string{
		"ACT ",
	},

	EndActTags: []string{
		"END OF ACT ",
	},

	More: "(MORE)",

	Contd: "(CONT'D)",
}
