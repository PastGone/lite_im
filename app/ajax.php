<?php


//这段代码是PHP语言的，它的作用是从URL中获取参数'c'和'k'的值。
//
//1. `require_once 'app.php';`：这行代码表示引入一个名为'app.php'的文件。如果该文件不存在或者无法被正确加载，那么程序将停止执行并输出错误信息。
//
//2. `$c = isset($_GET['c'])?$_GET['c']:'';`：这行代码首先检查URL中是否存在名为'c'的参数。如果存在，那么就将其值赋给变量$c；如果不存在，那么就将空字符串赋给变量$c。
//
//3. `$k = isset($_GET['k'])?intval($_GET['k']):0;`：
//这行代码与上一行类似，但是它处理的是名为'k'的参数。
//如果这个参数存在，那么就将其值转换为整数并赋给变量$k；如果不存在，那么就将0赋给变量$k。
re quire_once 'app.php';
$c = isset($_GET['c'])?$_GET['c']:'';
$k = isset($_GET['k'])?intval($_GET['k']):0;


switch ($c) {

//    这段代码是一个PHP的switch语句，根据不同的case执行不同的操作。
//
//1. 当case为"login"时，首先获取HTTP_REFERER的值，然后解析这个URL。
//如果URL的查询字符串等于LSTR，那么将$_SESSION[KEY.'login']设置为'admin'。
//接着从GET参数中获取用户名，
//如果cookie中没有存储用户名，
// 那么就调用nick()函数获取用户名，否则直接从cookie中获取用户名。
//最后，调用get_token()函数获取令牌，并将结果以JSON格式输出。
//
//2. 其他情况不涉及任何操作。
	 case "login":
         //这段代码的作用是获取HTTP请求的来源URL，并检查其查询参数是否等于LSTR。
         //如果等于LSTR，则在会话中设置一个名为KEY.'login'的变量值为'admin'。

		  $ref = $_SERVER["HTTP_REFERER"];
		  $url = parse_url($ref);

		  if($url['query']==LSTR){
		     $_SESSION[KEY.'login'] = 'admin';			  
		  }


		  $user = mb_substr(strip_tags($_GET['n']),0,5,'utf-8');
		  if(empty($_COOKIE[KEYS.'_name'])){
		      $arr = nick($user);

		  }else{

			  $arr['name'] =urldecode($_COOKIE[KEYS.'_name']);
		      $arr['key'] = $_COOKIE[KEYS.'_key'];
		  }
		  get_token();

		  echo json_encode($arr);
		  break;

          /*
           case 'send':
    $arr['msg'] = strip_tags($_POST['msg']); // 从POST请求中获取消息内容，并去除HTML标签
    $arr['name'] = strip_tags(urldecode($_COOKIE[KEYS.'_name'])); // 从Cookie中获取用户名，并去除HTML标签和URL解码
    $arr['key'] = strip_tags(urldecode($_COOKIE[KEYS.'_key'])); // 从Cookie中获取密钥，并去除HTML标签和URL解码

    if(check_post($arr) == false){ // 调用check_post函数检查消息内容的合法性
        $arr['type']= 'msg';
        $str = json_encode($arr);
        exit($str);
    }

    $_SESSION[KEY.'time'] = time(); // 更新会话中的当前时间戳
    unset($_SESSION[KEY . 'token']); // 删除会话中的令牌
    get_token(); // 调用get_token函数生成新的令牌

    $arr['msg'] = mb_substr($arr['msg'],0,140,'utf-8'); // 截取消息内容的前140个字符（以UTF-8编码）

    if($arr['msg'] =='clear' && $_SESSION[KEY.'login'] == 'admin'){ // 如果消息内容为"clear"且用户登录状态为管理员
        file_put_contents(ROOT_PATH.MSGFILE, '' , LOCK_EX); // 清空消息文件的内容
        $arr['type'] = 'sys'; // 设置消息类型为系统消息
        $arr['msg'] = 'rebot'; // 设置消息内容为"rebot"
        $str = json_encode($arr); // 将消息对象转换为JSON字符串
    }else{
        $arr['type']= 'msg'; // 设置消息类型为普通消息
        $_SESSION[KEY.'msg'] = $arr['msg']; // 将消息内容存储到会话中
        $str = json_encode($arr); // 将消息对象转换为JSON字符串
        file_put_contents(ROOT_PATH.MSGFILE, $str."\n" , FILE_APPEND|LOCK_EX); // 将消息写入消息文件，追加模式并加锁
    }

    echo $str; // 输出JSON字符串
    break;


           * */
    /*

这段代码的主要功能是处理发送消息的操作。首先，它从POST请求中获取消息内容，并去除HTML标签。
然后，它从Cookie中获取用户名和密钥，并进行相应的处理。
接下来，它调用check_post函数来检查消息内容的合法性。
    如果消息内容合法，它会更新会话中的当前时间戳，并生成新的令牌。
    然后，它截取消息内容的前140个字符（以UTF-8编码）。
    如果消息内容为"clear"且用户登录状态为管理员，它会清空消息文件的内容，
    并将消息类型设置为系统消息，
    消息内容设置为"rebot"。否则，它将消息类型设置为普通消息，并将消息内容存储到会话中。
    最后，它将消息写入消息文件，追加模式并加锁。最终，它输出JSON字符串作为响应。*/
    case 'send':

	$arr['msg'] = strip_tags($_POST['msg']);
	$arr['name'] = strip_tags(urldecode($_COOKIE[KEYS.'_name']));
	$arr['key'] = strip_tags(urldecode($_COOKIE[KEYS.'_key']));

	if(check_post($arr) == false){
	   //logmsg(0);
	   $arr['type']= 'msg';
	   //$arr['msg']= '内容发送失败！';
	   $str = json_encode($arr);
	   exit($str);
	}

	$_SESSION[KEY.'time'] = time();

	unset($_SESSION[KEY . 'token']) ;

    get_token();


    $arr['msg'] = mb_substr($arr['msg'],0,140,'utf-8');


	if($arr['msg'] =='clear' && $_SESSION[KEY.'login'] == 'admin'){
	//
	   file_put_contents(ROOT_PATH.MSGFILE, '' , LOCK_EX);

       $arr['type'] = 'sys';
	   $arr['msg'] = 'rebot';
	   $str = json_encode($arr);
	}else{
	  $arr['type']= 'msg';
	  $_SESSION[KEY.'msg'] = $arr['msg'];
	  //
	  file_put_contents(ROOT_PATH.MSGFILE, '' , LOCK_EX);

	  $str = json_encode($arr);

//这段代码的功能是处理一个名为'msg'的请求
//。它首先从指定的文件路径读取消息内容，并将其存储在变量$str中。然后，使用explode()函数将字符串按换行符分割成一个数组，并将结果存储在变量$arr中
//。接下来，计算数组的长度减一，得到消息的数量，并将结果存储在变量$count中。
//接下来的代码块包含了几个条件判断和操作：
//如果请求中的参数$k等于消息数量$count，则表示已经获取了所有消息，
//此时将$msg['count']设置为消息数量，$msg['list']设置为空数组，并使用json_encode()函数将结果转换为JSON格式输出。
//然后使用exit()函数终止脚本执行。
//如果请求中的参数$k大于消息数量$count，则表示请求超出了消息范围，
//此时将$msg['type']设置为'sys'，$msg['msg']设置为'rebot'，
//   并将结果添加到$res['list']数组中。
//同时，将$res['count']设置为消息数量，并使用json_encode()函数将结果转换为JSON格式输出。
//然后使用exit()函数终止脚本执行。
//如果请求中的参数$k小于消息数量减去50，则表示需要获取更多的消息，此时将$k设置为消息数量减去50。
//使用array_slice()函数从数组$arr中截取从索引$k开始的所有元素，并将结果存储在新的数组中。
//然后使用array_pop()函数移除数组的最后一个元素。
//最后，将消息数量$count和截取后的消息列表$arr存储在$msg['count']和$msg['list']中，
// 并使用json_encode()函数将结果转换为JSON格式输出。
  case 'msg':
	//$sk = $_
    $str = file_get_contents(ROOT_PATH.MSGFILE);
    $arr = explode("\n",$str);

	$count = count($arr)-1;
	//echo ($count);exit();
	if($k==$count){
	  $msg['count'] = $count;
	  $msg['list'] = [];
	  echo json_encode($msg);
	  exit();
	}
	if($k>count($arr)-1){
        $msg['type'] = 'sys';
        $msg['msg'] = 'rebot';
		$res['list'][] = json_encode($msg);
        $res['count'] = $count;
        echo json_encode($res); 
		exit();
	}

	//$k = $k==0?$k:$k+1;
	if($k<($count-50)){
	   $k= $count-50;
	}

	$arr = array_slice($arr,$k);
	array_pop($arr);

	$msg['count'] = $count;
	$msg['list'] = $arr;

	echo json_encode($msg);
	break;

}