package main

import "fmt"

//에러 구조체 선언
type PasswordError struct {
	Len        int
	RequireLen int
}

//Error() 메서드 : error 인터페이스로 사용될 수 있음
func (err PasswordError) Error() string {
	return "암호 길이가 짧다!"
}

//암호길이가 짧을때 PasswordError 구조체 정보 반환
func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8}
	}

	return nil
}

func main() {
	//ID PW 입력
	err := RegisterAccount("myID", "myPw")

	//에러 발생한 경우
	if err != nil {
		//에러를 PasswordError 타입으로 변환 -> 에러메세지와 구조체 필드에 접근하여 인터페이스 변환 성공 여부 검사
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len:%d ReauireLen:%d\n", errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("회원 가입됐습니다.")
	}
}
