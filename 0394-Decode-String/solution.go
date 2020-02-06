/*
this one is not original work 
*/
//nnn array slice 



type item struct {
    prestr string
    loopNum int 
}

//实现判断第一个为数字功能
func helper(s string) string {
    var res string // 移动的单位
    var loopNum int 
        // 通过item stack 实现
    for step := 0 ; step < len(s); step++{
        v := s[step]
        if v >= '0' && v <= '9' {
            loopNum = loopNum*10 + int(v-'0')
        } else if v == '[' {
            stack = append(stack,item{prestr: res, loopNum: loopNum})
            res = ""      //dynamic programmin
            loopNum = 0
        } else if v == ']' {
            item := stack[len(stack)-1] // pop
            for i := 0; i < item.loopNum; i++ {
                item.prestr += res
            }
            res = item.prestr
            stack = stack[:len(stack)-1] // 剪栈 没匹配的还留着比如第一个"["
        } else {
            res += string(v)
        }
    }
    return res 
}

func decodeString(s string) string {
    return helper(s)
}


