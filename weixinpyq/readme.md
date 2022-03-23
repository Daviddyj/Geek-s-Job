本次作业的设计思路
在老师原有代码上进行修改，增加  DeletePersonPyq() 和  ShowPersonPyq()  完成功能需求

1.首先注册pyq状态信息，如下图      通过ShowPersonPyq()查看
[
{
"id": 20,
"person_id": 100,
"personal_name": "戴一杰",
"content": "开心",
"show_time": "2022-03-20 19:49:17",
"by_time_age": 26,
"by_time_height": 1.76,
"by_time_weight": 74,
"by_time_fatrate": 0.184474,
"personal_sex": "男"
},
{
"id": 21,
"person_id": 101,
"personal_name": "王瑞",
"content": "快乐",
"show_time": "2022-03-20 20:49:17",
"by_time_age": 26,
"by_time_height": 1.58,
"by_time_weight": 55,
"by_time_fatrate": 0.270181,
"personal_sex": "女"
},
{
"id": 22,
"person_id": 102,
"personal_name": "高佳楠",
"content": "今天又泡到一个",
"show_time": "2022-03-20 23:49:17",
"by_time_age": 26,
"by_time_height": 1.75,
"by_time_weight": 60,
"by_time_fatrate": 0.132902,
"personal_sex": "男"
},
{
"id": 23,
"person_id": 103,
"personal_name": "笨蛋",
"content": "可可爱爱，没有脑袋",
"show_time": "2022-03-21 23:49:17",
"by_time_age": 18,
"by_time_height": 1.88,
"by_time_weight": 70,
"by_time_fatrate": 0.117064,
"personal_sex": "男"
}
]

2.删除id=22 是用户的状态    数据库信息还在，通过ShowPersonPyq()不显示该用户状态
{
"id": 20,
"person_id": 100,
"personal_name": "戴一杰",
"content": "开心",
"show_time": "2022-03-20 19:49:17",
"by_time_age": 26,
"by_time_height": 1.76,
"by_time_weight": 74,
"by_time_fatrate": 0.184474,
"personal_sex": "男"
},
{
"id": 21,
"person_id": 101,
"personal_name": "王瑞",
"content": "快乐",
"show_time": "2022-03-20 20:49:17",
"by_time_age": 26,
"by_time_height": 1.58,
"by_time_weight": 55,
"by_time_fatrate": 0.270181,
"personal_sex": "女"
},
{
"id": 23,
"person_id": 103,
"personal_name": "笨蛋",
"content": "可可爱爱，没有脑袋",
"show_time": "2022-03-21 23:49:17",
"by_time_age": 18,
"by_time_height": 1.88,
"by_time_weight": 70,
"by_time_fatrate": 0.117064,
"personal_sex": "男"
}