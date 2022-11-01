// 这是一个动态起床时间计算库，根据不同时节太阳的升起时间，给出对应的偏移量获取起床时间
package getup

import (
	"time"
)

// day：计算一年中的哪一天 例： 03-23（必须是两位数，用0补位）
// sTime: 设置最早的起床时间，日出过早防止过早起床
// dTime: 设置最晚的起床时间，日出过晚，防止过晚起床
// offset: 日出后的偏移量时间，在此偏移量后起床，例：+30m：30分钟后、 -30m：三十分钟前
func Calc(day, sTime, dTime, offset string) (string, error) {
	// 加载配置文件json
	config := map[string][]string{
		"range": {"06:00", "07:00"},
		"01-01": {"07:12", "16:11"},
		"01-02": {"07:12", "16:12"},
		"01-03": {"07:12", "16:13"},
		"01-04": {"07:12", "16:14"},
		"01-05": {"07:12", "16:15"},
		"01-06": {"07:12", "16:16"},
		"01-07": {"07:12", "16:17"},
		"01-08": {"07:12", "16:18"},
		"01-09": {"07:11", "16:19"},
		"01-10": {"07:11", "16:20"},
		"01-11": {"07:11", "16:21"},
		"01-12": {"07:11", "16:22"},
		"01-13": {"07:10", "16:23"},
		"01-14": {"07:10", "16:24"},
		"01-15": {"07:09", "16:26"},
		"01-16": {"07:09", "16:27"},
		"01-17": {"07:08", "16:28"},
		"01-18": {"07:08", "16:29"},
		"01-19": {"07:07", "16:31"},
		"01-20": {"07:06", "16:32"},
		"01-21": {"07:06", "16:33"},
		"01-22": {"07:05", "16:34"},
		"01-23": {"07:04", "16:36"},
		"01-24": {"07:04", "16:37"},
		"01-25": {"07:03", "16:38"},
		"01-26": {"07:02", "16:40"},
		"01-27": {"07:01", "16:41"},
		"01-28": {"07:00", "16:42"},
		"01-29": {"06:59", "16:44"},
		"01-30": {"06:58", "16:45"},
		"01-31": {"06:57", "16:47"},
		"02-01": {"06:56", "16:48"},
		"02-02": {"06:55", "16:49"},
		"02-03": {"06:54", "16:51"},
		"02-04": {"06:53", "16:52"},
		"02-05": {"06:51", "16:53"},
		"02-06": {"06:50", "16:55"},
		"02-07": {"06:49", "16:56"},
		"02-08": {"06:48", "16:58"},
		"02-09": {"06:46", "16:59"},
		"02-10": {"06:45", "17:00"},
		"02-11": {"06:44", "17:02"},
		"02-12": {"06:42", "17:03"},
		"02-13": {"06:41", "17:04"},
		"02-14": {"06:40", "17:06"},
		"02-15": {"06:38", "17:07"},
		"02-16": {"06:37", "17:08"},
		"02-17": {"06:35", "17:10"},
		"02-18": {"06:34", "17:11"},
		"02-19": {"06:32", "17:13"},
		"02-20": {"06:31", "17:14"},
		"02-21": {"06:29", "17:15"},
		"02-22": {"06:28", "17:17"},
		"02-23": {"06:26", "17:18"},
		"02-24": {"06:25", "17:19"},
		"02-25": {"06:23", "17:21"},
		"02-26": {"06:21", "17:22"},
		"02-27": {"06:20", "17:23"},
		"02-28": {"06:18", "17:24"},
		"02-29": {"06:16", "17:26"},
		"03-01": {"06:15", "17:27"},
		"03-02": {"06:13", "17:28"},
		"03-03": {"06:11", "17:30"},
		"03-04": {"06:10", "17:31"},
		"03-05": {"06:08", "17:32"},
		"03-06": {"06:06", "17:33"},
		"03-07": {"06:04", "17:35"},
		"03-08": {"06:03", "17:36"},
		"03-09": {"06:01", "17:37"},
		"03-10": {"05:59", "17:38"},
		"03-11": {"05:57", "17:40"},
		"03-12": {"05:56", "17:41"},
		"03-13": {"05:54", "17:42"},
		"03-14": {"05:52", "17:43"},
		"03-15": {"05:50", "17:45"},
		"03-16": {"05:48", "17:46"},
		"03-17": {"05:47", "17:47"},
		"03-18": {"05:45", "17:48"},
		"03-19": {"05:43", "17:50"},
		"03-20": {"05:41", "17:51"},
		"03-21": {"05:39", "17:52"},
		"03-22": {"05:38", "17:53"},
		"03-23": {"05:36", "17:55"},
		"03-24": {"05:34", "17:56"},
		"03-25": {"05:32", "17:57"},
		"03-26": {"05:30", "17:58"},
		"03-27": {"05:29", "17:59"},
		"03-28": {"05:27", "18:01"},
		"03-29": {"05:25", "18:02"},
		"03-30": {"05:23", "18:03"},
		"03-31": {"05:21", "18:04"},
		"04-01": {"05:20", "18:05"},
		"04-02": {"05:18", "18:07"},
		"04-03": {"05:16", "18:08"},
		"04-04": {"05:14", "18:09"},
		"04-05": {"05:12", "18:10"},
		"04-06": {"05:11", "18:11"},
		"04-07": {"05:09", "18:13"},
		"04-08": {"05:07", "18:14"},
		"04-09": {"05:05", "18:15"},
		"04-10": {"05:04", "18:16"},
		"04-11": {"05:02", "18:17"},
		"04-12": {"05:00", "18:19"},
		"04-13": {"04:58", "18:20"},
		"04-14": {"04:57", "18:21"},
		"04-15": {"04:55", "18:22"},
		"04-16": {"04:53", "18:23"},
		"04-17": {"04:52", "18:25"},
		"04-18": {"04:50", "18:26"},
		"04-19": {"04:48", "18:27"},
		"04-20": {"04:47", "18:28"},
		"04-21": {"04:45", "18:29"},
		"04-22": {"04:44", "18:31"},
		"04-23": {"04:42", "18:32"},
		"04-24": {"04:40", "18:33"},
		"04-25": {"04:39", "18:34"},
		"04-26": {"04:37", "18:35"},
		"04-27": {"04:36", "18:37"},
		"04-28": {"04:34", "18:38"},
		"04-29": {"04:33", "18:39"},
		"04-30": {"04:31", "18:40"},
		"05-01": {"04:30", "18:41"},
		"05-02": {"04:29", "18:43"},
		"05-03": {"04:27", "18:44"},
		"05-04": {"04:26", "18:45"},
		"05-05": {"04:25", "18:46"},
		"05-06": {"04:23", "18:47"},
		"05-07": {"04:22", "18:48"},
		"05-08": {"04:21", "18:50"},
		"05-09": {"04:19", "18:51"},
		"05-10": {"04:18", "18:52"},
		"05-11": {"04:17", "18:53"},
		"05-12": {"04:16", "18:54"},
		"05-13": {"04:15", "18:55"},
		"05-14": {"04:14", "18:56"},
		"05-15": {"04:12", "18:57"},
		"05-16": {"04:11", "18:58"},
		"05-17": {"04:10", "19:00"},
		"05-18": {"04:09", "19:01"},
		"05-19": {"04:08", "19:02"},
		"05-20": {"04:07", "19:03"},
		"05-21": {"04:06", "19:04"},
		"05-22": {"04:06", "19:05"},
		"05-23": {"04:05", "19:06"},
		"05-24": {"04:04", "19:07"},
		"05-25": {"04:03", "19:08"},
		"05-26": {"04:02", "19:09"},
		"05-27": {"04:02", "19:10"},
		"05-28": {"04:01", "19:10"},
		"05-29": {"04:00", "19:11"},
		"05-30": {"04:00", "19:12"},
		"05-31": {"03:59", "19:13"},
		"06-01": {"03:59", "19:14"},
		"06-02": {"03:58", "19:15"},
		"06-03": {"03:58", "19:15"},
		"06-04": {"03:57", "19:16"},
		"06-05": {"03:57", "19:17"},
		"06-06": {"03:57", "19:18"},
		"06-07": {"03:56", "19:18"},
		"06-08": {"03:56", "19:19"},
		"06-09": {"03:56", "19:20"},
		"06-10": {"03:56", "19:20"},
		"06-11": {"03:55", "19:21"},
		"06-12": {"03:55", "19:21"},
		"06-13": {"03:55", "19:22"},
		"06-14": {"03:55", "19:22"},
		"06-15": {"03:55", "19:23"},
		"06-16": {"03:55", "19:23"},
		"06-17": {"03:55", "19:23"},
		"06-18": {"03:55", "19:24"},
		"06-19": {"03:55", "19:24"},
		"06-20": {"03:55", "19:24"},
		"06-21": {"03:56", "19:24"},
		"06-22": {"03:56", "19:25"},
		"06-23": {"03:56", "19:25"},
		"06-24": {"03:56", "19:25"},
		"06-25": {"03:57", "19:25"},
		"06-26": {"03:57", "19:25"},
		"06-27": {"03:58", "19:25"},
		"06-28": {"03:58", "19:25"},
		"06-29": {"03:58", "19:25"},
		"06-30": {"03:59", "19:25"},
		"07-01": {"03:59", "19:25"},
		"07-02": {"04:00", "19:24"},
		"07-03": {"04:01", "19:24"},
		"07-04": {"04:01", "19:24"},
		"07-05": {"04:02", "19:24"},
		"07-06": {"04:02", "19:23"},
		"07-07": {"04:03", "19:23"},
		"07-08": {"04:04", "19:22"},
		"07-09": {"04:05", "19:22"},
		"07-10": {"04:05", "19:21"},
		"07-11": {"04:06", "19:21"},
		"07-12": {"04:07", "19:20"},
		"07-13": {"04:08", "19:20"},
		"07-14": {"04:09", "19:19"},
		"07-15": {"04:10", "19:18"},
		"07-16": {"04:10", "19:18"},
		"07-17": {"04:11", "19:17"},
		"07-18": {"04:12", "19:16"},
		"07-19": {"04:13", "19:15"},
		"07-20": {"04:14", "19:15"},
		"07-21": {"04:15", "19:14"},
		"07-22": {"04:16", "19:13"},
		"07-23": {"04:17", "19:12"},
		"07-24": {"04:18", "19:11"},
		"07-25": {"04:19", "19:10"},
		"07-26": {"04:20", "19:09"},
		"07-27": {"04:21", "19:08"},
		"07-28": {"04:22", "19:07"},
		"07-29": {"04:23", "19:06"},
		"07-30": {"04:24", "19:04"},
		"07-31": {"04:25", "19:03"},
		"08-01": {"04:26", "19:02"},
		"08-02": {"04:28", "19:01"},
		"08-03": {"04:29", "19:00"},
		"08-04": {"04:30", "18:58"},
		"08-05": {"04:31", "18:57"},
		"08-06": {"04:32", "18:56"},
		"08-07": {"04:33", "18:54"},
		"08-08": {"04:34", "18:53"},
		"08-09": {"04:35", "18:51"},
		"08-10": {"04:36", "18:50"},
		"08-11": {"04:37", "18:49"},
		"08-12": {"04:39", "18:47"},
		"08-13": {"04:40", "18:46"},
		"08-14": {"04:41", "18:44"},
		"08-15": {"04:42", "18:43"},
		"08-16": {"04:43", "18:41"},
		"08-17": {"04:44", "18:40"},
		"08-18": {"04:45", "18:38"},
		"08-19": {"04:47", "18:36"},
		"08-20": {"04:48", "18:35"},
		"08-21": {"04:49", "18:33"},
		"08-22": {"04:50", "18:31"},
		"08-23": {"04:51", "18:30"},
		"08-24": {"04:52", "18:28"},
		"08-25": {"04:53", "18:26"},
		"08-26": {"04:54", "18:25"},
		"08-27": {"04:56", "18:23"},
		"08-28": {"04:57", "18:21"},
		"08-29": {"04:58", "18:20"},
		"08-30": {"04:59", "18:18"},
		"08-31": {"05:00", "18:16"},
		"09-01": {"05:01", "18:14"},
		"09-02": {"05:02", "18:13"},
		"09-03": {"05:03", "18:11"},
		"09-04": {"05:05", "18:09"},
		"09-05": {"05:06", "18:07"},
		"09-06": {"05:07", "18:05"},
		"09-07": {"05:08", "18:04"},
		"09-08": {"05:09", "18:02"},
		"09-09": {"05:10", "18:00"},
		"09-10": {"05:11", "17:58"},
		"09-11": {"05:13", "17:56"},
		"09-12": {"05:14", "17:55"},
		"09-13": {"05:15", "17:53"},
		"09-14": {"05:16", "17:51"},
		"09-15": {"05:17", "17:49"},
		"09-16": {"05:18", "17:47"},
		"09-17": {"05:19", "17:45"},
		"09-18": {"05:20", "17:44"},
		"09-19": {"05:22", "17:42"},
		"09-20": {"05:23", "17:40"},
		"09-21": {"05:24", "17:38"},
		"09-22": {"05:25", "17:36"},
		"09-23": {"05:26", "17:34"},
		"09-24": {"05:27", "17:32"},
		"09-25": {"05:28", "17:31"},
		"09-26": {"05:30", "17:29"},
		"09-27": {"05:31", "17:27"},
		"09-28": {"05:32", "17:25"},
		"09-29": {"05:33", "17:23"},
		"09-30": {"05:34", "17:21"},
		"10-01": {"05:35", "17:20"},
		"10-02": {"05:37", "17:18"},
		"10-03": {"05:38", "17:16"},
		"10-04": {"05:39", "17:14"},
		"10-05": {"05:40", "17:12"},
		"10-06": {"05:41", "17:11"},
		"10-07": {"05:42", "17:09"},
		"10-08": {"05:44", "17:07"},
		"10-09": {"05:45", "17:05"},
		"10-10": {"05:46", "17:04"},
		"10-11": {"05:47", "17:02"},
		"10-12": {"05:48", "17:00"},
		"10-13": {"05:50", "16:59"},
		"10-14": {"05:51", "16:57"},
		"10-15": {"05:52", "16:55"},
		"10-16": {"05:53", "16:53"},
		"10-17": {"05:55", "16:52"},
		"10-18": {"05:56", "16:50"},
		"10-19": {"05:57", "16:49"},
		"10-20": {"05:58", "16:47"},
		"10-21": {"06:00", "16:45"},
		"10-22": {"06:01", "16:44"},
		"10-23": {"06:02", "16:42"},
		"10-24": {"06:04", "16:41"},
		"10-25": {"06:05", "16:39"},
		"10-26": {"06:06", "16:38"},
		"10-27": {"06:07", "16:36"},
		"10-28": {"06:09", "16:35"},
		"10-29": {"06:10", "16:33"},
		"10-30": {"06:11", "16:32"},
		"10-31": {"06:13", "16:31"},
		"11-01": {"06:14", "16:29"},
		"11-02": {"06:15", "16:28"},
		"11-03": {"06:16", "16:27"},
		"11-04": {"06:18", "16:25"},
		"11-05": {"06:19", "16:24"},
		"11-06": {"06:20", "16:23"},
		"11-07": {"06:22", "16:22"},
		"11-08": {"06:23", "16:20"},
		"11-09": {"06:24", "16:19"},
		"11-10": {"06:26", "16:18"},
		"11-11": {"06:27", "16:17"},
		"11-12": {"06:28", "16:16"},
		"11-13": {"06:30", "16:15"},
		"11-14": {"06:31", "16:14"},
		"11-15": {"06:32", "16:13"},
		"11-16": {"06:34", "16:12"},
		"11-17": {"06:35", "16:11"},
		"11-18": {"06:36", "16:10"},
		"11-19": {"06:37", "16:09"},
		"11-20": {"06:39", "16:09"},
		"11-21": {"06:40", "16:08"},
		"11-22": {"06:41", "16:07"},
		"11-23": {"06:42", "16:07"},
		"11-24": {"06:44", "16:06"},
		"11-25": {"06:45", "16:05"},
		"11-26": {"06:46", "16:05"},
		"11-27": {"06:47", "16:04"},
		"11-28": {"06:48", "16:04"},
		"11-29": {"06:50", "16:03"},
		"11-30": {"06:51", "16:03"},
		"12-01": {"06:52", "16:02"},
		"12-02": {"06:53", "16:02"},
		"12-03": {"06:54", "16:02"},
		"12-04": {"06:55", "16:02"},
		"12-05": {"06:56", "16:01"},
		"12-06": {"06:57", "16:01"},
		"12-07": {"06:58", "16:01"},
		"12-08": {"06:59", "16:01"},
		"12-09": {"07:00", "16:01"},
		"12-10": {"07:01", "16:01"},
		"12-11": {"07:02", "16:01"},
		"12-12": {"07:03", "16:01"},
		"12-13": {"07:03", "16:01"},
		"12-14": {"07:04", "16:02"},
		"12-15": {"07:05", "16:02"},
		"12-16": {"07:06", "16:02"},
		"12-17": {"07:06", "16:02"},
		"12-18": {"07:07", "16:03"},
		"12-19": {"07:08", "16:03"},
		"12-20": {"07:08", "16:04"},
		"12-21": {"07:09", "16:04"},
		"12-22": {"07:09", "16:05"},
		"12-23": {"07:10", "16:05"},
		"12-24": {"07:10", "16:06"},
		"12-25": {"07:10", "16:06"},
		"12-26": {"07:11", "16:07"},
		"12-27": {"07:11", "16:08"},
		"12-28": {"07:11", "16:08"},
		"12-29": {"07:12", "16:09"},
		"12-30": {"07:12", "16:10"},
		"12-31": {"07:12", "16:11"},
	}

	// 获取当天日期，并根据当天时间获取日出时间
	// nowDay := string(time.Now().Format("01-02"))
	uptime, err := time.Parse("15:04", config[day][0])
	if err != nil {
		return "", err
	}

	// 从配置文件获取最早起床时间
	stime, err := time.Parse("15:04", sTime)
	if err != nil {
		return "", err
	}

	// 从配置文件获取最晚起床时间
	dtime, err := time.Parse("15:04", dTime)
	if err != nil {
		return "", err
	}

	// 日出后半小时初步认定为起床时间
	var timeOver time.Time
	m, _ := time.ParseDuration(offset)
	timeOver = uptime.Add(m)

	// 如果日出半小时后依然早于最早起床时间，那么就判定为起床时间为 最早起床时间
	// 如果日出半小时后时间晚于最早起床时间，初步判定日出半小时后时间为 起床时间
	if timeOver.Before(stime) {
		timeOver = stime
	}
	// 如果日出后半小时超过最晚起床时间，那么设定起床时间为最晚起床时间
	if timeOver.After(dtime) {
		timeOver = dtime
	}

	return timeOver.Format("15:04"), nil
}
