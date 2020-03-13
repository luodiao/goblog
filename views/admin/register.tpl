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

      <form class="layui-form">
        <div class="layui-form-item">
          <label class="layui-form-label">用户名</label>
          <div class="layui-input-block">
            <input type="text" autocomplete="off" v-model="m.name" name="username" placeholder="请输入用户名" class="layui-input" :class="{'layui-form-danger':isValid && m.name==''}">
          </div>
        </div>

        <div class="layui-form-item">
          <label class="layui-form-label">邮箱</label>
          <div class="layui-input-block">
            <input type="email" name="email" v-model="m.email" placeholder="邮箱" class="layui-input" :class="{'layui-form-danger':isValid && m.emaill==''}">
          </div>
        </div>

        <div class="layui-form-item">
          <label class="layui-form-label">密码</label>
          <div class="layui-input-block">
            <input type="password" name="password" v-model="m.password" placeholder="密码" class="layui-input" :class="{'layui-form-danger':isValid && m.password==''}">
          </div>
        </div>

        <div class="layui-form-item">
          <label class="layui-form-label">确认密码</label>
          <div class="layui-input-block">
            <input type="password" name="password" v-model="m.passwordConfirm" placeholder="确认密码" class="layui-input" :class="{'layui-form-danger':isValid && m.passwordConfirm==''}">
          </div>
        </div>

        <div class="layui-form-item">
          <button type="button" class="layui-btn layui-btn-fluid layui-btn-normal" @click="save">注 册</button> 
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
<script src="/static/js/vue.js"></script>
<script src="/static/js/site.js"></script>
<script>
const vm = new Vue({
  el: '#main',
  data() {
    return {
      m: {
        name: '',
        email: '',
        password: '',
        passwordConfirm: ''
      },
      isValid: false,
    };
  },
  methods: {
    save(event) {
      /*this.isValid = true;
      if (this.m.name == ''
        || this.m.email == ''
        || this.m.password == ''
      ) {
        layer.msg('请填写完整信息!!', {icon: 4});
        return false;
      }
      if (this.m.password != this.m.passwordConfirm) {
        layer.msg('两次密码不一致!!', {icon: 4});
        return false;
      }*/
      
      var $btn = $(event.target).loading()
      $.post("", {
        action: 'add',
        data: {a:1,b:2}
      }, function (resp) {
        console.log(resp)
        $btn.loading('reset')
      });
    }
  }
});
</script>
</html>


