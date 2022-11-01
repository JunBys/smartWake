#### 这是一个Golang的动态起床时间计算库，根据不同时节太阳的升起时间，给出对应的偏移量获取起床时间

通过写一个简单的demon程序可以演示出 如下效果
``` go
func main() {
	day := os.Args[1]	
	sTime := os.Args[2]
	dTime := os.Args[3]
	offset := os.Args[4]
	
	// day：计算一年中的哪一天 例： 03-23（必须是两位数，用0补位）
	// sTime: 设置最早的起床时间，日出过早防止过早起床
	// dTime: 设置最晚的起床时间，日出过晚，防止过晚起床
	// offset: 日出后的偏移量时间，在此偏移量后起床，例：+30m：30分钟后、 -30m：三十分钟前

	time, err := getup.Calc(day, sTime, dTime, offset)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(time)
}
```

<img width="606" alt="image" src="https://user-images.githubusercontent.com/86999978/199165929-fee86740-8c6f-4135-be3d-712f82dea5d5.png">
