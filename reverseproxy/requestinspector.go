package reverseproxy

import (
	"net/http"
	"strings"
	"sort"
	"errors"
	"fmt"
)

type RequestInspector interface {
	Inspect(req *http.Request) error
}

type AllowQueryTypesInspector struct {
	allowedArgs []string
}

func NewAllowQueryTypesInspector(allowedArgs []string) RequestInspector {
	return &AllowQueryTypesInspector{
		allowedArgs: allowedArgs,
	}
}

func (insp AllowQueryTypesInspector) Inspect(req *http.Request) error {
	if rawQuery := req.URL.RawQuery; rawQuery != "" {	
		args := strings.Split(rawQuery, "&")
		for _, arg  := range args {
			argkey := strings.Split(arg, "=")[0]
			if !insp.contains(argkey, insp.allowedArgs) {
				return errors.New(fmt.Sprintf("invalid request, %s is not allowed", argkey))
			}
		}
	}

	return nil
}

func (AllowQueryTypesInspector) contains(val string, arr []string) bool {
	sort.Strings(arr)
	i := sort.SearchStrings(arr, val)
	return i < len(arr) && arr[i] == val
}
