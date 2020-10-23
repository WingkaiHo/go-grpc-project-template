package manager

import (
	"fmt"
	"os"
)

type detailsManager struct {
}

// DetailsManager to export function
var DetailsManager detailsManager

func (m *detailsManager) Get(id uint64) (uint64, string) {

	hostName, _ := os.Hostname()
	name := fmt.Sprintf("%s-%d", hostName, id)

	return id, name
}
