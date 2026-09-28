package main

import (
	"bmi_app/apperrors"
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// プロファイル用の構造体
type Profile struct {
	name   string
	height float64
	weight float64
}

// BMI取得・判定メソッド
func (p *Profile) JudgeByBMI() (float64, string) {

	// 身長のメートル変換
	height := p.height / 100

	// BMI算出
	bmi := p.weight / (height * height)

	switch  {
	case bmi < 18.5:
		return bmi, "低体重（痩せ型）"
	case bmi < 25.0:
		return bmi, "普通体重"
	default:
		return bmi, "肥満"
	}
}

// スキャナーから値を読み取る関数
func readLine(s *bufio.Scanner, message string) (string, error) {

	// 質問を出力
	fmt.Println(message)

	// 入力が確認できたら空白を除去
	if s.Scan() {
		input := strings.TrimSpace(s.Text())
		return input, nil
	}
	if s.Err() != nil {
		return "", s.Err()
	}
	return "", io.EOF
}

// スキャナーから値を読み取り、数値に変換する関数
func readFloat(s *bufio.Scanner, message string) (float64, error) {
	for {
		input, err := readLine(s, message)
		if err != nil {
			return 0, err
		}

		inputFloat, err := strconv.ParseFloat(input, 64)

		// 変換結果の値にNaNまたはInfinityでないかも含める
		if err != nil || math.IsNaN(inputFloat) || math.IsInf(inputFloat, 0) {
			fmt.Println("半角数字を入力してください")
			continue
		}
		return inputFloat, nil
	}
}


func main() {

	// プロファイル構造体の初期化
	profile := &Profile{}

	// 標準入力を読み取る準備
	scanner := bufio.NewScanner(os.Stdin)

	// ウィザード開始
	fmt.Println("*** 健康診断ウィザード  ***")

	name, err := readLine(scanner, "あなたのお名前を教えてください")
	if err != nil {
		fmt.Println(apperrors.HandleInputError(err))
		return
	}
	if name == "" {
		fmt.Println("空の名前は受付できません。")
		return
	}
	profile.name = name
	fmt.Printf(`%sさん、こんにちは。`, profile.name)

	height, err := readFloat(scanner, "身長(cm)を入力してください。")
	if err != nil {
		fmt.Println(apperrors.HandleInputError(err))
		return
	}

	if height <= 0 {
		fmt.Println("0より大きい数値を入力してください")
		return
	}

	profile.height = height

	weight, err := readFloat(scanner, "体重(kg)を入力してください。")
	if err != nil {
		fmt.Println(apperrors.HandleInputError(err))
		return
	}

	if weight <= 0 {
		fmt.Println("0より大きい数値を入力してください")
		return
	}

	profile.weight = weight

	// BMI算出・判定
	bmi, judge := profile.JudgeByBMI()

	fmt.Println("================================")
	fmt.Printf("【診断結果】 %s 様\n", profile.name)
	fmt.Printf("BMI数値: %.2f\n", bmi)
	fmt.Printf("判定: %s\n ", judge)
	fmt.Println("================================")

}
