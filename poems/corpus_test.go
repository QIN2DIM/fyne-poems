package poems

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"strings"
	"testing"
)

func TestSelector_GetCorpus(t *testing.T) {
	s := NewSelector()
	cursorBounds := []int{0, len(s.corpus) - 1, 1, len(s.corpus), len(s.corpus) - 2, (len(s.corpus) - 1) / 2}
	for _, cb := range cursorBounds {
		s.cursor = cb
		for i := 0; i <= 500; i++ {
			s.GetCurrentCorpus()
			s.GetNextCorpus()
			assert.NotEqual(t, -1, s.cursor, "OutOfBounds")
			assert.NotEqual(t, len(s.corpus), s.cursor, "OutOfBounds")
		}
	}
}

func TestI18n(t *testing.T) {
	cmd := exec.Command("powershell", "Get-Culture | select -exp Name")
	output, err := cmd.Output()
	if err == nil {
		langLocRaw := strings.TrimSpace(string(output))
		langLoc := strings.Split(langLocRaw, "-")
		lang := langLoc[0]
		loc := langLoc[1]
		fmt.Println(lang, loc)
	}
}
