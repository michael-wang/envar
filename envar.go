package envar

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var (
	defVals = map[string]interface{}{}
)

// SetDef set default value for key.
// When env var is not found, value set here is returned.
func SetDef(key string, value interface{}) {
	defVals[key] = value
}

// Load filenames, or .env if called by Load().
//
// Notice existed env var won't be replaced by loading these files.
func Load(filenames ...string) (err error) {
	return godotenv.Load(filenames...)
}

// String returns var from env based on key.
func String(key string) string {
	v := Get(key)

	str, ok := v.(string)
	if ok {
		return str
	}

	log.Printf("envar.String: no string value for key: %s\n", key)
	return ""
}

// Int returns integer value based on key.
func Int(key string) int {
	v := Get(key)

	i, ok := v.(int)
	if ok {
		return i
	}

	str, ok := v.(string)
	if ok {
		i, err := strconv.Atoi(str)
		if err == nil {
			return i
		}
	}

	log.Printf("envar.Int: no integer value for key: %s\n", key)
	return 0
}

// Bool returns boolean value based on key.
func Bool(key string) bool {
	v := Get(key)

	b, ok := v.(bool)
	if ok {
		return b
	}

	str, ok := v.(string)
	if ok {
		b, err := strconv.ParseBool(str)
		if err == nil {
			return b
		}
	}

	log.Printf("envar.Bool: no bool value for key: %s\n", key)
	return false
}

// Float returns float value based on key.
func Float(key string) float64 {
	v := Get(key)

	f, ok := v.(float64)
	if ok {
		return f
	}

	str, ok := v.(string)
	if ok {
		f, err := strconv.ParseFloat(str, 64)
		if err == nil {
			return f
		}
	}

	log.Printf("envar.Float: no float value for key: %s\n", key)
	return 0
}

// MilliSeconds assumes key mapped to value of millisecond (e.g. 1000 for one second),
// and returns it as time.Duration type.
func MilliSeconds(key string) time.Duration {
	v := Get(key)

	i, ok := v.(int)
	if ok {
		return time.Duration(i) * time.Millisecond
	}

	s, ok := v.(string)
	if ok {
		i, err := strconv.Atoi(s)
		if err == nil {
			return time.Duration(i) * time.Millisecond
		}
	}

	log.Printf("envar.MilliSeconds: no proper value for key: %s\n", key)
	return time.Duration(0)
}

// Get returns var from env based on key.
func Get(key string) interface{} {
	v, ok := os.LookupEnv(key)
	if !ok {
		v, ok := defVals[key]
		if !ok {
			log.Fatalf("envar.Get variable not found, key: %s\n", key)
		}
		return v
	}
	return v
}
