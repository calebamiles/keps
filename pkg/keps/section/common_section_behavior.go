package section

type section interface {
	Filename() string
	Name() string
	Content() []byte
	Persist() error
	Erase() error
}

type commonSectionInfo struct {
	filename   string
	name       string
	content    []byte
	contentDir string
}

func (i *commonSectionInfo) Filename() string { return i.filename }
func (i *commonSectionInfo) Name() string     { return i.name }
func (i *commonSectionInfo) Content() []byte  { return i.content }
