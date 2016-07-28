package algorithms

import (
	"bytes"
)

//将字符串的指定开始的字符转移到尾部
func LeftMoveRight(str string,n int) string {
	if(len(str) <= n){
		return str;
	}
	left := bytes.Buffer{};
	right := bytes.Buffer{};

	for i,s := range str{
		if(i >= n){
			right.WriteRune(s);
		}else{
			left.WriteRune(s);
		}

	}
	return right.String()+left.String();
}

//判断B中的所有字符串是否都在A中存在
func StringContain(a string,b string) bool {
	for _,s1:=range b {
		var has bool = false;
		for _,s2:= range a{
			if( s1 == s2){
				has = true;
			}
		}
		if(has == false){
			return false;
		}
	}
	return  true;
}

//如果连个字符串中字符一样，出现册书也一样，只是顺序不一样，则认为这两个字符串是兄弟字符串。如果从一个数组中快速找出一个字符串的所有兄弟字符串
func FindBrothers(dic []string,str string) []string {
	newDic := make([]string,1);

	for _,s1 := range dic {
		if(len(s1) != len(str)){
			continue;
		}
		num := 0;
		for _,s2:=range s1{
			for _,s3 := range str{
				if(s2 == s3){
					num = num +1;
				}
			}
		}

		if num == len(str){
			newDic = append(newDic,s1);
		}
	}
	return newDic;
}