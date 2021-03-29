package pkg

import (
	"fmt"

	"github.com/Frug/goqu-table-generator/config"
)

func Generate(conf config.Config) {
	fmt.Printf("\nGenerating: %+v\n", conf)
}
