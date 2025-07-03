package pkg

import (
	"os"
	"regexp"
	"strings"
)

func ExpandEnvVars(s string) string {
	re := regexp.MustCompile(`\${(\w+)}`)
	return re.ReplaceAllStringFunc(s, func(match string) string {
		varName := strings.TrimPrefix(match, "${")
		varName = strings.TrimSuffix(varName, "}")
		if value, exists := os.LookupEnv(varName); exists {
			return value
		}
		return match
	})
}

func IsEnvVarRaw(s string) bool {
	re := regexp.MustCompile(`^\${\w+}$`)
	return re.MatchString(s)
}
