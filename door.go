package main

type Door struct {
	usualTR          *TimeRule
	daylightSavingTR *TimeRule
}

func (d *Door) IsOpenToEnter(t string, isDaylightSaving bool) error {
	if isDaylightSaving {
		return d.daylightSavingTR.IsOpenToEnter(t)
	}

	return d.usualTR.IsOpenToEnter(t)
}
