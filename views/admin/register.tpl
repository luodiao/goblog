<!DOCTYPE html>
<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" href="/static/layui/css/layui.css" media="all">
  <style>
  .img-box {
    position: relative;
    height: 150px;
    width: 300px;
  }
  .img-box img {
    position: absolute;
  }
  </style>
</head>

<body>
  <div class="layui-row" style="margin-top: 15%;">
    <div class="layui-col-md4 layui-col-md-offset4">
      <fieldset class="layui-elem-field layui-field-title" style="margin-top: 30px;">
        <legend>GO BLOG</legend>
      </fieldset>

      <form class="layui-form" action="/admin/register" method="post">
        <div class="layui-form-item">
          <label class="layui-form-label">用户名</label>
          <div class="layui-input-block">
            <input type="text" autocomplete="off" name="username" placeholder="请输入用户名" class="layui-input">
          </div>
        </div>

        <div class="layui-form-item">
          <label class="layui-form-label">密码</label>
          <div class="layui-input-block">
            <input type="password" name="password" placeholder="密码" class="layui-input">
          </div>
        </div>

        <div class="layui-form-item">
          <label class="layui-form-label">确认密码</label>
          <div class="layui-input-block">
            <input type="password" name="password" placeholder="确认密码" class="layui-input">
          </div>
        </div>

        <div class="layui-form-item">
          <button class="layui-btn layui-btn-fluid layui-btn-normal" lay-submit="" lay-filter="LAY-user-login-submit">注 册</button> 
        </div>

        <div class="layui-form-item" style="text-align:center;">
          <a href="login">登录</a>
        </div>
      </form>
    </div>
  </div>
</body>

<script src="https://cdn.bootcss.com/jquery/3.4.1/jquery.min.js"></script>
<script src="/static/layui/layui.all.js"></script>
<script>
layui.use('slider', function(){
  var slider = layui.slider;
  
  //渲染
  slider.render({
    elem: '#slideTest7'  //绑定元素
    ,tips: false //关闭默认提示层
    ,change: function(value){
      var offLeft = {{.offLeft}};

      $("#slider-block").css({left: offLeft + value * (Math.abs({{.offLeft}}) / 100)});
    }
  });
});

init();
function init() {
  $("#slider-block").css({left: {{.offLeft}}});
}
</script>
</html>


