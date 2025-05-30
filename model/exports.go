package model

import "regexp"

var ParamRegex = regexp.MustCompile(`(\w+)\['([\w{}/:_-]+)'`)
var IndexRegex = regexp.MustCompile(`(\w+)\[(\d+)]`)
