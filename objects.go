package util

type Result struct {
	Items     []*Item           `json:"items"`
	Variables map[string]string `json:"variables"`
}

type Item struct {
	Uid          string            `json:"uid"`
	Type         string            `json:"type"`
	Title        string            `json:"title"`
	Subtitle     string            `json:"subtitle"`
	Arg          string            `json:"arg"` // 相当于query
	Autocomplete string            `json:"autocomplete"`
	Icon         *Icon             `json:"icon"`
	Variables    map[string]string `json:"variables"`
	QuickLookUrl string            `json:"quicklookurl"`
	Text         *Text             `json:"text"` // Defines the text the user will get when copying the selected result row with ⌘C or displaying large type with ⌘L.
}

type Text struct {
	Copy      string `json:"copy"`
	Largetype string `json:"largetype"`
}

type Icon struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

type AlfredWorkflowRsp struct {
	AlfredWorkflow *AlfredWorkflow `json:"alfredworkflow"`
}

type AlfredWorkflow struct {
	Arg       string            `json:"arg"`
	Variables map[string]string `json:"variables"`
}
