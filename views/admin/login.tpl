<!DOCTYPE html>
<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" href="/static/layui/css/layui.css" media="all">
</head>

<body>
  <div class="layui-row" style="margin-top: 20%;">
    <div class="layui-col-md4 layui-col-md-offset4">
      <fieldset class="layui-elem-field layui-field-title" style="margin-top: 30px;">
        <legend>GO BLOG</legend>
      </fieldset>
      <form class="layui-form" action="/admin/login" method="post">
        <div class="layui-form-item">
          <label class="layui-form-label">用户名</label>
          <div class="layui-input-block">
            <input type="text" autocomplete="off" placeholder="请输入用户名" class="layui-input">
          </div>
        </div>

        <div class="layui-form-item">
          <label class="layui-form-label">密码</label>
          <div class="layui-input-block">
            <input type="password" placeholder="密码" class="layui-input">
          </div>
        </div>

        <div class="layui-form-item">
          <button class="layui-btn layui-btn-fluid" lay-submit="" lay-filter="LAY-user-login-submit">登 入</button>
        </div>
      </form>
    </div>
  </div>
</body>
</html>


