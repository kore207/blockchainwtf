package cli

import (
	"fmt"
	"os"
)

func Usage() {
	fmt.Println("Welcome to Blockchain WTF")
	fmt.Println("다음의 명령어(command)를 입력하세요.")
	fmt.Println("레벨 테스트 : test")
	fmt.Println("레벨 별 학습 : edu")
	fmt.Println("노마드 코인 실행 : nomad")
	os.Exit(0)
}

func EduStarter() {
	fmt.Println("블록체인이 뭐야?\r\n made by Team HyangChon")
	fmt.Println("=======================================")
	//Todo: 상중하로 나누기가 좀 애매함 -> level 테스트와 연계 가능할듯
	fmt.Println("당신의 블록체인 지식을 상, 중, 하로 표헌 한다면?")
	fmt.Println("1. 상(블록체인의 핵심 원리가 무엇인지 알고 직접 활용 가능한 수준")
	fmt.Println("2. 중(블록체인의 핵심 원리는 알지만 어떻게 활용할지는 모름")
	fmt.Println("3. 하(블록체인이라는 단어만 들어봤을 뿐 뭔지 모름")
}

func LevelTestStarter() {
	fmt.Println("다음 XX가지 문항을 통해 당신의 블록체인 지식 LEVEL을 측정합니다.")

}
