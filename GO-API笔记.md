# **1.数字(int)与字符串(string)的相互转换 strconv包(String Conversions)**<br>
int --→ string `strconv.Itoa(n)` Integer to ASCII<br>
string -> int `strconv.Atoi(s)` ASCII to Integer<br>

如果处理int64或者非十进制数 需要用`strconv.ParseInt``和``strconv.FormatInt`


# **2.字符串``string``和字节切片`[]byte`的互转**<br>
- string-->[]byte: `b := []byte(s)`<br>
- []byte-->string: `s:=string(b)` <br>


## 示例：将int整型数N转为换字节切片
第一步：int--> string <br>
第二部：string --> []byte <br>
`[]byte(strconv.Itoa(n))`


