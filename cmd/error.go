package cmd

type TextError struct {
	Text string
}

func (err *TextError) Error() string {
	return err.Text
}
