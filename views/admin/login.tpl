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
  <div id="main" class="layui-row" style="margin-top: 15%;">
    <div class="layui-col-md4 layui-col-md-offset4">
      <fieldset class="layui-elem-field layui-field-title" style="margin-top: 30px;">
        <legend>GO BLOG</legend>
      </fieldset>

      <div class="layui-form-item">
        <label class="layui-form-label">用户名</label>
        <div class="layui-input-block">
          <input type="text" autocomplete="off" v-model.trim="m.name" placeholder="请输入用户名" class="layui-input">
        </div>
      </div>

      <div class="layui-form-item">
        <label class="layui-form-label">密码</label>
        <div class="layui-input-block">
          <input type="password" name="password" v-model.trim="m.password" placeholder="密码" class="layui-input">
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
        <button class="layui-btn layui-btn-fluid" type="button" @click="login">登 入</button> 
      </div>

      <div class="layui-form-item" style="text-align:center;">
        <a href="register" tyle>注册</a>
      </div>
    </div>
  </div>
</body>

<script src="/static/js/jquery.min.js"></script>
<script src="/static/layui/layui.all.js"></script>
<script src="/static/js/vue.js"></script>
<script src="/static/js/site.js"></script>
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

const vm = new Vue({
  el: "#main",
  data() {
    return {
      m: {
        name: '',
        email: '',
        password: '',
      },
    }
  },
  methods: {
    login() {
      var $btn = $(event.target).loading()
      $.post("/admin/ajaxLogin", this.m, function (resp) {
        if (resp.code == 0) {
          window.location = '/admin'
        } else {
          alert(resp.msg);
          $btn.loading('reset')
        }
      });
    }
  }
});
</script>
</html>


