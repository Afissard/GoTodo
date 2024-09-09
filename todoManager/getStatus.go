package todomanager

func (t task) GetStatus() (status string) {
	switch t.Status {
	case NotStarted:
		return "Not started"
	case Started:
		return "Started"
	case Finished:
		return "Finished"
	default:
		return "Unknown"
	}
}
