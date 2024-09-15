package logger

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	
	"zoo/domain/domError"
)

func Field(key string, value interface{}) LogField {
	return LogField{Key: key, Value: value}
}

func (l *Logger) getFields(uuid string, fields ...interface{}) (logField []LogField) {
	
	// custom fields logs
	logField = append(logField, Field("uuid", uuid))
	info := l.haveErrorField(fields...)
	
	if info != "" {
		logField = append(logField, Field("stacktrace", info))
	} else {
		// case where service proxier didn't return domError, just logging domError
		// passing all of the logging fields into lrFields
		for i, v := range fields {
			if i%2 == 0 {
				// need to cast to string,
				// because we dont now the type of v (interface) is string or not
				// if not cast to string, it will panic domError
				logField = append(logField, Field(fmt.Sprintf("%v", v), fields[i+1]))
			} else {
				continue
			}
		}
		
		logField = append(logField, Field("file", fileInfo(3)))
	}
	return logField
}

func (l *Logger) haveErrorField(fields ...interface{}) string {
	for _, field := range fields {
		if reflect.TypeOf(field).Kind() == reflect.Map {
			maps, ok := field.(map[string]interface{})
			if !ok {
				return ""
			}
			for _, v := range maps {
				// check if v is an std domError type
				stderr, ok := v.(error)
				if ok {
					// convert to *efserr.Error
					err, ok := stderr.(*domError.Error)
					if ok {
						return err.Stack()
					}
					return ""
				} else { // check posibility if *efserr.Error
					err, ok := v.(*domError.Error)
					if ok {
						return err.Stack()
					}
					return ""
				}
			}
			continue
		}
		stderr, ok := field.(error)
		if ok {
			// convert to *efserr.Error
			err, ok := stderr.(*domError.Error)
			if ok {
				return err.Stack()
			}
			return ""
		} else { // check posibility if *efserr.Error
			err, ok := field.(*domError.Error)
			if ok {
				return err.Stack()
			}
			return ""
		}
	}
	return ""
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = ""
		line = 1
	} else {
		dirList := strings.Split(file, "/")
		if dirList[len(dirList)-1] == "main.go" {
			file = "main.go"
		} else {
			file = fmt.Sprintf("%s/%s", dirList[len(dirList)-2], dirList[len(dirList)-1])
		}
	}
	
	return fmt.Sprintf("%s:%d", file, line)
}
