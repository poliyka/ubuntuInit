package core

import (
	"sync"
)

var RcPath string
var Wg sync.WaitGroup
var Lock sync.Mutex
var Ic = NewInstallChoices()
