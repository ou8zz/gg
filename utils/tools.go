package utils

import (
	"log"
	"errors"
)


// 在后面程序出现异常的时候就会捕获
func RecoverException() (err error) {
	defer func() {
		if r := recover(); r != nil {
			// 这里可以对异常进行一些处理和捕获
			log.Println("Exception:", r)

			//check exactly what the panic was and create error.
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknow panic")
			}
		}
	}()
	return err
}

// 错误打印
func CheckErr(errMasg error) {
	if errMasg != nil {
		log.Print(errMasg)
	}
}

