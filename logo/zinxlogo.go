package logo

import (
	"fmt"
)

var zinxLogo = `                                        
.oPYo. o     .oPYo. o      o .oPYo.  .oPYo. 
8    8 8     8    8 8      8 8.      8    8 
8      8     8    8 8      8  boo   o8YooP 
8   oo 8     8    8 8  db  8 .P      8    b 
8    8 8     8    8  b.PY.d  8       8    8 
 YooP8 8oooo  YooP    8  8    YooP  8    8 
:....8 ......:.....:::..:..:::.....::..:::..
:::::8 :::::::::::::::::::::::::::::::::::::
:::::..:::::::::::::::::::::::::::::::::::::

							请享受无法回避的痛苦
                                        `
var topLine = `┌──────────────────────────────────────────────────────┐`
var borderLine = `│`

var bottomLine = `└──────────────────────────────────────────────────────┘`

func PrintLogo() {
	fmt.Println(zinxLogo)
	fmt.Println(topLine)
	fmt.Println(fmt.Sprintf("%s [Github] https://github.com/6902140                    %s", borderLine, borderLine))
	fmt.Println(bottomLine)
}
