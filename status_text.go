package healthcheck

type StatusText string

func (s StatusText) String() string {
	return string(s)
}

func StatusTextFromStatus(in Status) StatusText {
	return StatusText(in.String())
}

type ReportStatusText string

func (s ReportStatusText) String() string {
	return string(s)
}

func ReportStatusTextFromStatus(in ReportStatus) ReportStatusText {
	return ReportStatusText(in.String())
}
