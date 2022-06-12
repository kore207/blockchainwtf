package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Welcome to Blockchain WTF")
	fmt.Println("다음의 명령어(command)를 입력하세요.")
	fmt.Println("레벨 테스트 : level")
	fmt.Println("레벨 별 학습 : edu")
	os.Exit(0)
}

func eduStarter() {
	fmt.Println("블록체인이 뭐야?\r\n made by Team HyangChon")
	fmt.Println("=======================================")
	fmt.Println("당신의 블록체인 지식을 상, 중, 하로 표헌 한다면?")
	fmt.Println("1. 상(블록체인의 핵심 원리가 무엇인지 알고 직접 활용 가능한 수준")
	fmt.Println("2. 중(블록체인의 핵심 원리는 알지만 어떻게 활용할지는 모름")
	fmt.Println("3. 하(블록체인이라는 단어만 들어봤을 뿐 뭔지 모름")
}

func main() {
	if len(os.Args) < 2 {
		usage()
	} else if os.Args[1] == "edu" {
		eduStarter()
	}

}
