package platform

import (
	"encoding/csv"
	"github.com/wyw14/cry038/internal/domain"
	"io"
	"strconv"
)

func ExportSession(w io.Writer, s domain.Session) error {
	cw := csv.NewWriter(w)
	defer cw.Flush()
	if err := cw.Write([]string{"item", "owner", "state", "consumed"}); err != nil {
		return err
	}
	for _, i := range s.Items {
		if err := cw.Write([]string{i.Name, i.Owner, string(i.State), strconv.Itoa(i.Quantity + i.Consumed)}); err != nil {
			return err
		}
	}
	return cw.Error()
}
