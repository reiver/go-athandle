package athndl

import (
	"golang.org/x/net/idna"
)

func Normalize(handle string) string {

	var str string = handle

	{
		var temp []byte = make([]byte, len(str))

		{
			length := len(str)

			for i:=0; i<length; i++ {
				v := str[i]

				switch {
				case 'A' <= v && v <= 'Z':
					const offset byte = 'a' - 'A'
					temp[i] = v + offset
				default:
					temp[i] = v
				}
			}
		}

		str = string(temp)
	}

	var err error
	str, err = idna.ToASCII(str)
	if nil != err {
		return handle
	}

	return str
}
