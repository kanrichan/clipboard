package main

func main() {

	// var count = [...]int{1, 2}
	// fmt.Println(int(count))
	// r := gin.Default()

	// auth := r.Group("")

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run("127.0.0.1:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

// for i := 0; i < 50; i++ {
// 	是否成功, 奖励 := 判定(i)
// 	fmt.Printf("次数%d 是否成功%v 奖励%d\n", i, 是否成功, 奖励)
// }

// func 判定(次数 int) (是否成功 bool, 奖励 int) {
// 	var 默认概率 = 0.5
// 	var 默认奖励 = 10

// 	var 随机数 = rand.Intn(10000)
// 	var 衰减概率系数 = 0.03
// 	var 当前概率 = math.Max(0.5-float64(次数)*衰减概率系数, 0.1)

// 	if 随机数 < int(10000*当前概率) {
// 		是否成功 = true
// 	}

// 	奖励 = int(float64(默认奖励) * (默认概率 / 当前概率))
// 	return
// }
