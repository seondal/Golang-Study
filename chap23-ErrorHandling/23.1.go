package main

import (
	"fmt"
	"os"
	"bufio"
)

func ReadFile(filename String) (string, error) {
	
	//파일 열기
	file err := os.Open(filename)

	// 파일열기 실패시 에러 반환
	if err != nil {
		return "", err
	}

	//파일 닫기 (함수 종료 직전에 반드시 호출해야함)
	defer file.Close()

	//\n 이 나올때까지 파일의 문자열 읽기
	rd := bufio.newReader(file)
	line, _ := rd.ReadString('\n')
	return line, nil
}

func WriteFile(filename string, line string) error {
	//파일 생성
	file, err := os.Create(filename)

	// 파일생성 실패시 에러 반환
	if err != nil {
		return err
	}
	defer file.Close()

	//파일 핸들에 문자열과 작성
	_, err = fmt.Fprintln(file, line)
	return err
}

const filename string = "data.txt"

func main() {
	//파일 읽기 시도
	line, err := ReadFile(filename)
	if err != nil {
		//파일 생성
		err = WriteFile(filename, "This is WriteFile")

		//에러 처리
		if err != nil {
			fmt.Println("파일 생성에 실패했습니다.", err)
			return
		}

		//다시 읽기 시도
		line, err = ReadFile(filename)
		if err != nil {
			fmt.Println("파일 읽기에 실패했습니다.", err)
			return
		}
	}

	fmt.Println("파일 내용: " , line)
}