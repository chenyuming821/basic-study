package main	


import (
	"fmt"
	"math"
)

// Sqrt 使用牛顿迭代法计算平方根
// 参数 x: 要求平方根的数
// 返回值: x 的平方根近似值
func Sqrt(x float64) float64 {
	// 方案1：固定迭代10次
	fmt.Println("=== 固定迭代10次 ===")
	z := 1.0 // 初始猜测值
	for i := 0; i < 10; i++ {
		fmt.Printf("第%d次迭代: z = %.10f\n", i+1, z)
		z -= (z*z - x) / (2 * z)
	}

	// 方案2：直到收敛（值不再变化）
	fmt.Println("\n=== 直到收敛 ===")
	z = 1.0 // 重置初始值
	iterations := 0
	for {
		fmt.Printf("第%d次迭代: z = %.10f\n", iterations+1, z)
		oldZ := z
		z -= (z*z - x) / (2 * z)
		iterations++
		// 当变化量小于很小的数时停止
		if math.Abs(z-oldZ) < 1e-10 {
			break
		}
	}
	fmt.Printf("收敛所需迭代次数: %d\n", iterations)

	return z
}

// SqrtWithGuess 使用不同初始猜测值计算平方根
// 参数:
//   x    - 要求平方根的数
//   guess - 初始猜测值
// 返回值:
//   result - 计算结果
//   iterations - 迭代次数
func SqrtWithGuess(x, guess float64) (result float64, iterations int) {
	z := guess
	for {
		oldZ := z
		z -= (z*z - x) / (2 * z)
		iterations++
		if math.Abs(z-oldZ) < 1e-10 {
			break
		}
	}
	return z, iterations
}

func main() {
	// 测试不同的值
	for _, x := range []float64{1, 2, 3, 10, 100} {
		fmt.Printf("\n=====================================\n")
		fmt.Printf("计算 sqrt(%g)\n", x)
		fmt.Printf("=====================================\n")
		
		result := Sqrt(x)
		
		fmt.Printf("\n【结果对比】\n")
		fmt.Printf("自定义 Sqrt:  %.15f\n", result)
		fmt.Printf("标准库 math.Sqrt: %.15f\n", math.Sqrt(x))
		fmt.Printf("两者差值:    %.15f\n", math.Abs(result-math.Sqrt(x)))
	}

	// 测试不同初始猜测值的影响
	fmt.Printf("\n\n=====================================\n")
	fmt.Printf("测试不同初始猜测值的影响\n")
	fmt.Printf("=====================================\n")
	
	testValues := []float64{2, 10, 100}
	
	for _, x := range testValues {
		fmt.Printf("\n--- 计算 sqrt(%g) ---\n", x)
		
		// 测试不同的初始猜测值
		guesses := []float64{1.0, x, x / 2}
		guessNames := []string{"z=1.0", "z=x", "z=x/2"}
		
		for i, guess := range guesses {
			result, iterations := SqrtWithGuess(x, guess)
			fmt.Printf("初始值 %s: 迭代%d次, 结果=%.15f, 与标准库差值=%.15f\n",
				guessNames[i], iterations, result, math.Abs(result-math.Sqrt(x)))
		}
	}
}