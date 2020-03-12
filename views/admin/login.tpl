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

      <form class="layui-form" action="/admin/login" method="post">
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
          <label class="layui-form-label">验证码</label>
          <div class="layui-input-block">
            <div class="layui-card">
              <div class="layui-card-body">
                <div class="img-box">
                  <img id="slider-bg" src="{{.verifyImgBg}}" width="300" height="150">
                  <img id="slider-block" src="{{.verifyImg}}" width="300" height="150">
                </div>

                <div class="slider-bar" style="margin-top: 10px;">
                  <div id="slideTest7" class="demo-slider"></div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="layui-form-item">
          <button class="layui-btn layui-btn-fluid" lay-submit="" lay-filter="LAY-user-login-submit">登 入</button> 
        </div>

        <div class="layui-form-item" style="text-align:center;">
          <a href="register" tyle>注册</a>
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
    elem: '#slideTest7'
    ,tips: false
    ,change: function(value){
      var offLeft = {{.offLeft}};
      
      $("#slider-block").css({left: (offLeft + value * (300 - 30) / 100)})
    }
  });
});

function sliderCall() {
  console.log($("#slider-block").css('left'));
  $.post("", {
    action: 'verify',
    accuracy: $("#slider-block").css('left').replace("px", "")
  }, function (resp) {
    alert(resp.msg);
  });
}

$('#slideTest7').mouseup(function() {
  sliderCall();
});

init();
function init() {
  $("#slider-block").css({left: {{.offLeft}}});
}
</script>
</html>


