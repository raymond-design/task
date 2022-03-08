package cli

import "fmt"

func wrapError(custom string, original error) error {
	return fmt.Errorf("%s: %v", custom, original)
}
