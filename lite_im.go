package main

import (
	"bufio"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var webId = "myLiteIM"

var psd = ""

type DataLine struct {
	Msg  string `json:"msg" `
	Name string `json:"name,omitempty"`
	Key  string `json:"key,omitempty"`
	Typ  string `json:"type"`
}

// SetNick
/*function nick($user=''){
if(empty($user)){
    $name=rand_nick();
}else{
    $name = $user;
}
    $arr['msg'] = '<span>'.date('Y-m-d H:i').'</span>';
	$arr['type']= 'sys';
	$str = json_encode($arr);
	$arr['msg'] = '<span class="tips-warning">系统消息：<strong>'.$name.'</strong>进入聊天室</span>';
	$arr['type']= 'sys';

	$str = $str."\n".json_encode($arr);
	file_put_contents(ROOT_PATH.MSGFILE, $str."\n" , FILE_APPEND|LOCK_EX);


	$key = uniqid();

    setcookie(KEYS.'_key',$key,time()+3600*24*90,'/');
    setcookie(KEYS.'_name',urlencode($name),time()+3600*24*30,'/');

    return array('name'=>$name,'key'=>$key); //输出生成的昵称
}
*/
func SetNick(n string, context **gin.Context) (string, string) {
	//设置name
	name := ""

	if n == "" {
		name = RandomNIck()
	} else {
		name = n

	}

	key := uuid.New().String()
	fmt.Println("uuid: ", key)

	(*context).SetCookie(fmt.Sprint(webId, "_name"), url.QueryEscape(name), 3600*24*90, "/", "localhost", false, true)

	(*context).SetCookie(fmt.Sprint(webId, "_key"), url.QueryEscape(key), 3600*24*90, "/", "localhost", false, true)

	// $arr['msg'] = '<span>'.date('Y-m-d H:i').'</span>';
	//	$arr['type']= 'sys';

	//	$str = json_encode($arr);

	//	$arr['msg'] = '<span class="tips-warning">系统消息：<strong>'.$name.'</strong>进入聊天室</span>';
	//	$arr['type']= 'sys';
	//
	//	$str = $str."\n".json_encode($arr);

	//	file_put_contents(ROOT_PATH.MSGFILE, $str."\n" , FILE_APPEND|LOCK_EX);
	now := time.Now()
	var dataLine = DataLine{
		Msg: fmt.Sprint("<span>", now.Format("2006-01-02 15:03"), "</span>"),
		Typ: fmt.Sprint("sys"),
	}

	//
	println(dataLine.Msg, dataLine.Typ)

	jsonByte, err := json.Marshal(dataLine)
	if err != nil {
		println("失败")
	}

	putStr := string(jsonByte)

	println(putStr)

	//
	var dataLine2 = DataLine{
		Msg: fmt.Sprint("<span class='tips-warning'>系统消息：<strong>", name, "</strong>进入聊天室</span>"),
		Typ: fmt.Sprint("sys"),
	}

	jsonByte, err = json.Marshal(dataLine2)
	if err != nil {
		println("失败")
	}

	putStr = fmt.Sprint(putStr, "\n", string(jsonByte), "\n")

	filePath := "./app/MSG.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(putStr)
	//Flush将缓存的文件真正写入到文件中
	write.Flush()

	println(putStr)
	return name, key

}

func RandomNIck() string {

	nameTou := []string{"快乐的", "冷静的", "醉熏的", "潇洒的", "糊涂的", "积极的", "冷酷的", "深情的", "粗暴的", "温柔的", "可爱的", "愉快的", "义气的", "认真的", "威武的", "帅气的", "传统的", "潇洒的", "漂亮的", "自然的", "专一的", "听话的", "昏睡的", "狂野的", "等待的", "搞怪的", "幽默的", "魁梧的", "活泼的", "开心的", "高兴的", "超帅的", "坦率的", "直率的", "轻松的", "痴情的", "完美的", "精明的", "无聊的", "丰富的", "繁荣的", "饱满的", "炙热的", "暴躁的", "碧蓝的", "俊逸的", "英勇的", "健忘的", "故意的", "无心的", "土豪的", "朴实的", "兴奋的", "幸福的", "淡定的", "不安的", "阔达的", "孤独的", "独特的", "疯狂的", "时尚的", "落后的", "风趣的", "忧伤的", "大胆的", "爱笑的", "矮小的", "健康的", "合适的", "玩命的", "沉默的", "斯文的", "任性的", "细心的", "粗心的", "大意的", "甜甜的", "酷酷的", "健壮的", "英俊的", "霸气的", "阳光的", "默默的", "大力的", "孝顺的", "忧虑的", "着急的", "紧张的", "善良的", "凶狠的", "害怕的", "重要的", "危机的", "欢喜的", "欣慰的", "满意的", "跳跃的", "诚心的", "称心的", "如意的", "怡然的", "娇气的", "无奈的", "无语的", "激动的", "愤怒的", "美好的", "感动的", "激情的", "激昂的", "震动的", "虚拟的", "超级的", "寒冷的", "精明的", "明理的", "犹豫的", "忧郁的", "寂寞的", "奋斗的", "勤奋的", "现代的", "过时的", "稳重的", "热情的", "含蓄的", "开放的", "无辜的", "多情的", "纯真的", "拉长的", "热心的", "从容的", "体贴的", "风中的", "曾经的", "追寻的", "儒雅的", "优雅的", "开朗的", "外向的", "内向的", "清爽的", "文艺的", "长情的", "平常的", "单身的", "伶俐的", "高大的", "懦弱的", "柔弱的", "爱笑的", "乐观的", "耍酷的", "酷炫的", "神勇的", "年轻的", "唠叨的", "瘦瘦的", "无情的", "包容的", "顺心的", "畅快的", "舒适的", "靓丽的", "负责的", "背后的", "简单的", "谦让的", "彩色的", "缥缈的", "欢呼的", "生动的", "复杂的", "慈祥的", "仁爱的", "魔幻的", "虚幻的", "淡然的", "受伤的", "雪白的", "高高的", "糟糕的", "顺利的", "闪闪的", "羞涩的", "缓慢的", "迅速的", "优秀的", "聪明的", "含糊的", "俏皮的", "淡淡的", "坚强的", "平淡的", "欣喜的", "能干的", "灵巧的", "友好的", "机智的"}
	nameWei := []string{"嚓茶", "凉面", "便当", "毛豆", "花生", "可乐", "灯泡", "野狼", "背包", "眼神", "缘分", "雪碧", "人生", "牛排", "蚂蚁", "飞鸟", "灰狼", "斑马", "汉堡", "悟空", "巨人", "绿茶", "大碗", "墨镜", "魔镜", "煎饼", "月饼", "月亮", "星星", "芝麻", "啤酒", "玫瑰", "大叔", "小伙", "太阳", "树叶", "芹菜", "黄蜂", "蜜粉", "蜜蜂", "信封", "西装", "外套", "裙子", "大象", "猫咪", "母鸡", "路灯", "蓝天", "白云", "星月", "彩虹", "微笑", "摩托", "板栗", "高山", "大地", "大树", "砖头", "楼房", "水池", "鸡翅", "蜻蜓", "红牛", "咖啡", "枕头", "大船", "诺言", "钢笔", "刺猬", "天空", "飞机", "大炮", "冬天", "洋葱", "春天", "夏天", "秋天", "冬日", "航空", "毛衣", "豌豆", "黑米", "玉米", "眼睛", "老鼠", "白羊", "帅哥", "美女", "季节", "鲜花", "服饰", "裙子", "秀发", "大山", "火车", "汽车", "歌曲", "舞蹈", "老师", "导师", "方盒", "大米", "麦片", "水杯", "水壶", "手套", "鞋子", "鼠标", "手机", "电脑", "书本", "奇迹", "身影", "香烟", "夕阳", "台灯", "宝贝", "未来", "皮带", "钥匙", "心锁", "故事", "花瓣", "滑板", "画笔", "画板", "学姐", "店员", "电源", "饼干", "宝马", "过客", "大白", "时光", "石头", "钻石", "河马", "犀牛", "西牛", "绿草", "抽屉", "柜子", "往事", "寒风", "路人", "橘子", "耳机", "鸵鸟", "朋友", "苗条", "铅笔", "钢笔", "硬币", "热狗", "大侠", "御姐", "萝莉", "毛巾", "期待", "盼望", "白昼", "黑夜", "大门", "黑裤", "哑铃", "板凳", "枫叶", "荷花", "乌龟", "衬衫", "大神", "草丛", "早晨", "心情", "茉莉", "流沙", "蜗牛", "猎豹", "棒球", "篮球", "乐曲", "电话", "网络", "世界", "中心", "老虎", "鸭子", "羽毛", "翅膀", "外套", "书包", "钢笔", "冷风", "烤鸡", "大雁", "音响", "招牌", "冰棍", "帽子"}
	s := rand.NewSource(time.Now().Unix())
	rand.New(s)
	name := nameTou[rand.Intn(len(nameTou))] + nameWei[rand.Intn(len(nameWei))]

	return name
}
func CheckPost(context **gin.Context, session *sessions.Session) bool {

	//$now = time();
	//    $token = $_COOKIE[md5(KEY.'token')];
	//	if(empty($token) or $token !=  $_SESSION[KEY . 'token']){
	//	   return false;
	//	}
	//	return true;
	token, err := (*context).Cookie(MD5(fmt.Sprint(webId, "_token")))
	if token == "" || err != nil || token == (*session).Get(fmt.Sprint(webId, "_token")) {
		return false
	}

	return true
}

// UpdateToken
/*  function get_token(){

  if(empty($_SESSION[KEY.'token'])){
      $token = md5(uniqid(rand(), true));
      $_SESSION[KEY.'token'] = $token;
  }else{
     $token = $_SESSION[KEY.'token'];
  }
  setcookie(md5(KEY.'token'),$token,time()+3600*24,'/');
  return $token;

}

*/

func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

func GetToken(context **gin.Context, session *sessions.Session) string {
	var token string
	if (*session).Get(fmt.Sprint(webId, "_token")) == nil {
		token = MD5(uuid.New().String())
		(*session).Set(fmt.Sprint(webId, "_token"), token)
		(*session).Save()
	} else {
		token = fmt.Sprint((*session).Get(fmt.Sprint(webId, "_token")))

	}
	(*context).SetCookie(MD5(fmt.Sprint(webId, "_token")), token, 3600*24, "/", "localhost", false, true)
	return token
}

func readAllContent(r io.Reader) ([]string, error) {
	var b = make([]byte, 4096)
	_, err := r.Read(b)
	if err != nil {
		return nil, err
	}

	l := strings.Split(string(b), "\n")
	return l, nil
}

func main() {

	ginServer := gin.Default()
	var addr = ":8080"
	store := cookie.NewStore([]byte("secret"))

	//路由上加入session中间件
	ginServer.Use(sessions.Sessions("mysession", store))

	ginServer.LoadHTMLGlob("./html/*")
	ginServer.StaticFS("/js", gin.Dir("static/js", false))
	ginServer.StaticFS("/css", gin.Dir("static/css", false))

	ginServer.GET("/", func(context *gin.Context) {
		context.HTML(200, "index.html", gin.H{
			"nick": RandomNIck(),
		})
	})

	// "/app/login?n=%E5%A4%A7%E7%AC%A8%E8%9B%8B"
	//Error #01: html/template: "addon.html" is undefined
	//json name key
	ginServer.GET("/app/login", func(context *gin.Context) {
		//将带有三个参数，标记为管理员，或者有单独的页面
		// n admin psd
		/*
			$ref = $_SERVER["HTTP_REFERER"];
				  $url = parse_url($ref);

				  if($url['query']==LSTR){
				     $_SESSION[KEY.'login'] = 'admin';
				  }
		*/
		/*
			//这段代码的作用是获取HTTP请求的来源URL，并检查其查询参数是否等于LSTR。
			         //如果等于LSTR，则在会话中设置一个名为KEY.'login'的变量值为'admin'。

		*/

		session := sessions.Default(context)

		cook, _ := context.Cookie(fmt.Sprint(webId, "_name"))

		if context.Query("admin") != "" && context.Query("psd") == psd {

			session.Set(fmt.Sprint(webId, "_login"), "admin")
			err := session.Save()

			if err != nil {
				return
			}
		}

		/*
			 $user = mb_substr(strip_tags($_GET['n']),0,5,'utf-8');
				  if(empty($_COOKIE[KEYS.'_name'])){
				      $arr = nick($user);

				  }else{

					  $arr['name'] =urldecode($_COOKIE[KEYS.'_name']);
				      $arr['key'] = $_COOKIE[KEYS.'_key'];
				  }
				  get_token();
				  echo json_encode($arr);


		*/

		var name = ""
		var key = ""
		println(context)
		println(context.Query("n"))
		n := context.Query("n")
		n = n[0:5]
		if cook == "" {
			name, key = SetNick(n, &context)
		} else {

			name, _ = context.Cookie(fmt.Sprint(webId, "_name"))
			key, _ = context.Cookie(fmt.Sprint(webId, "_key"))
		}

		//生成信息
		//追加信息
		//cookie
		//令牌
		GetToken(&context, &session)
		//返回值
		context.JSON(200, gin.H{
			"name": name,
			"key":  key,
		})

	},
	)

	//json name key type msg
	ginServer.POST("/app/send", func(context *gin.Context) {

		//获取信息

		//检查

		//更新令牌

		//判断是否是命令

		//追加
		type content struct {
			msg string
			//`json:"username"`
		}
		var myContent content
		if err := context.BindJSON(&myContent); err != nil {
			// 返回错误信息
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//var dataLine2 = DataLine{
		//		Msg: fmt.Sprint("<span class='tips-warning'>系统消息：<strong>", name, "</strong>进入聊天室</span>"),
		//		Typ: fmt.Sprint("sys"),
		//	}
		//
		//	jsonByte, err = json.Marshal(dataLine2)
		//	if err != nil {
		//		println("失败")
		//	}
		//
		//	putStr = fmt.Sprint(putStr, "\n", string(jsonByte), "\n")

		var msg = myContent.msg[0:140]

		println("随便打印：", myContent.msg)

		var name, _ = context.Cookie(fmt.Sprint(webId, "_name"))
		var key, _ = context.Cookie(fmt.Sprint(webId, "_name"))

		var dataLine = DataLine{
			Msg:  msg,
			Name: name,
			Key:  key,
		}

		////logmsg(0);
		//	   $arr['type']= 'msg';
		//	   //$arr['msg']= '内容发送失败！';
		//	   $str = json_encode($arr);
		//	   exit($str);

		session := sessions.Default(context)

		if !CheckPost(&context, &session) {
			println("发送失败")

			dataLine.Msg = "发送失败"
			dataLine.Typ = "msg"
			context.JSON(200, dataLine)
		}
		session.Set(fmt.Sprint(webId, "_time"), time.Now().Format("2006-01-02 15:03"))

		session.Delete(fmt.Sprint(webId, "_token"))

		session.Save()
		GetToken(&context, &session)

		//		if($arr['msg'] =='clear' && $_SESSION[KEY.'login'] == 'admin'){
		//file_put_contents(ROOT_PATH.MSGFILE, '' , LOCK_EX);
		//
		//$arr['type'] = 'sys';
		//$arr['msg'] = 'rebot';
		//$str = json_encode($arr);
		//}else{
		//$arr['type']= 'msg';
		//$_SESSION[KEY.'msg'] = $arr['msg'];
		//file_put_contents(ROOT_PATH.MSGFILE, '' , LOCK_EX);
		//
		//$str = json_encode($arr);
		//
		if dataLine.Msg == "clear" && session.Get(fmt.Sprint(webId, "_login")) == "admin" {
			dataLine.Typ = "sys"
			dataLine.Msg = "rebot"
			//
			context.JSON(200, dataLine)

		} else {
			dataLine.Typ = "msg"
			session.Set(fmt.Sprint(webId, "_msg"), dataLine.Msg)
			session.Save()
		}

		//filePath := "./app/MSG.txt"
		//
		//var dataLine2 = DataLine{
		//	Msg: fmt.Sprint("<span class='tips-warning'>系统消息：<strong>", name, "</strong>进入聊天室</span>"),
		//	Typ: fmt.Sprint("sys"),
		//}
		//
		//
		//
		//file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
		//if err != nil {
		//	fmt.Println("文件打开失败", err)
		//}
		//defer file.Close()
		////写入文件时，使用带缓存的 *Writer
		//write := bufio.NewWriter(file)
		//write.WriteString(putStr)
		////Flush将缓存的文件真正写入到文件中
		//write.Flush()
		//

		println(dataLine.Msg)
		println("send out")
		//返回值
		//context.JSON(200, dataLine)

		context.JSON(200, gin.H{
			"msg": dataLine.Msg,
			name:  dataLine.Name,
		})
		//调用get_msg
		//写入

	})
	//json msg :
	//count int
	//AList list:(json_string )
	//          type
	//          msg

	ginServer.GET("/app/msg", func(context *gin.Context) {
		var k = context.Query("k")
		var cd int
		var err error
		if k == "" {
			cd = 0
		} else {
			cd, err = strconv.Atoi(k)
			if err != nil {
				cd = 0
			}

		}
		//读取信息
		f, err := os.Open("./app/MSG.txt")
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		defer f.Close()

		content, err := readAllContent(f)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		//fmt.Println("content:", content)
		count := len(content) - 1

		println(content[0])
		println(count)
		//if($k==$count){
		//	  $msg['count'] = $count;
		//	  $msg['list'] = [];
		//	  echo json_encode($msg);
		//	  exit();
		//	}
		//	if($k>count($arr)-1){
		//        $msg['type'] = 'sys';
		//        $msg['msg'] = 'rebot';
		//		$res['list'][] = json_encode($msg);
		//        $res['count'] = $count;
		//        echo json_encode($res);
		//		exit();
		//	}
		type msgList struct {
			count int
			list  []DataLine
		}

		if cd == count {
			context.JSON(200, msgList{
				count: count,
				list:  nil,
			})
		}

		if cd > count {
			var list = [1]DataLine{{
				Msg: "renbot",
				Typ: "sys",
			}}

			context.JSON(200, msgList{
				count: count,
				list:  list[0:0],
			})

		}
		if cd < count-50 {
			cd = count - 50
		}

		//$arr = array_slice($arr,$k);
		//	array_pop($arr);

		//	$msgList['count'] = $count;
		//	$msgList['list'] = $arr;
		//	echo json_encode($msgList);
		//
		arr := content[cd : len(content)-1]

		//lastElement := arr[len(arr)-1]
		//arr = arr[:len(arr)-1]

		//分割转换，得到信标
		//比较大小，做出选择
		//返回值
		println("msg over")
		context.JSON(200, gin.H{
			"count": count,
			"list":  arr,
		})

	})
	{

		println(fmt.Sprint("http://localhost" + addr + "/"))

	}

	ginServer.Run(addr)

}
