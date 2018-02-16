<!DOCTYPE html>

<html>
<head>
  <title>Massmail</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <!--<link href="https://cdn.bootcss.com/flat-ui/2.3.0/css/flat-ui.min.css" rel="stylesheet">-->
  <link href="https://cdn.bootcss.com/flat-ui/2.3.0/css/vendor/bootstrap/css/bootstrap.min.css" rel="stylesheet">
  <link href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap-theme.min.css" rel="stylesheet">
</head>

<body>
<div class="container">
  <h1 class="text-center">自助邮件群发系统</h1>
  <form>
    {{ .xsrfdata }}
    <fieldset>
      <legend>邮件相关信息</legend>
      <div class="row">
        <div class = "form-group col-sm-3">
          <label for="aliasName" class="control-label">发送邮件显示名称：</label>
          <input type="text" class="form-control" id="aliasName" placeholder="发送邮件显示名称" name="aliasName">
        </div>
        <div class="form-group col-sm-3">
          <label for="fromEmail" class="control-label">发送人地址：</label>
          <input type="email" class="form-control" id="fromEmail" placeholder="Email" name="fromEmail">
        </div>
      </div>
    </fieldset>
    <fieldset>
      <legend>发送邮件服务器信息</legend>
      <div class="row">
        <div class="form-group col-sm-3">
          <label for="smtpServer" class="control-label">发送邮件服务器地址：</label>
          <input type="text" class="form-control" id="smtpServer" placeholder="smtp.xxx.com" name="smtpServer">
        </div>
        <div class="form-group col-sm-3">
          <label for="port" class="control-label">发送邮件服务器端口：</label>
          <input type="text" class="form-control" id="port" placeholder="465" name="port">
        </div>
        <div class="form-group col-sm-3">
          <label for="password" class="control-label">发送邮件服务器登录密码：</label>
          <input type="password" class="form-control" id="password" placeholder="******" name="password">
        </div>
      </div>
    </fieldset>
    <fieldset>
      <legend>邮件发送配置</legend>
      <div class="row">
        <div class="form-group col-sm-3">
          <label for="varField" class="control-label">邮件变量在文件中的列</label>
          <input type="text" class="form-control" id="varField" placeholder="2,3,5" name="varField">
        </div>
        <div class="form-group col-sm-9">
          <label for="delayTime[]" class="control-label">每封邮件发送间隔时间</label>
          <div class="form-inline">
            <span>最小时间:</span>
            <input type="text" class="form-control" placeholder="1" name="delayTime[]">
            <span>最大时间:</span>
            <input type="text" class="form-control" placeholder="5" name="delayTime[]">
          </div>
        </div>
      </div>
    </fieldset>
    <fieldset>
      <legend>发送邮件所需文件</legend>
      <div class="row">
        <div class="form-group col-sm-3">
          <label for="varFile" class="control-label">邮件变量文件</label>
          <input type="file" id="varFile" name="varFile">
        </div>
        <div class="form-group col-sm-3">
          <label for="subjectFile" class="control-label">邮件主题模板文件</label>
          <input type="file" id="subjectFile" name="subjectFile">
        </div>
        <div class="form-group col-sm-3">
          <label for="messageFile" class="control-label">邮件正文模板文件</label>
          <input type="file" id="messageFile" name="messageFile">
        </div>
      </div>
    </fieldset>
    <div class="row">
      <button class="btn btn-default" type="submit">开始发送</button>
    </div>
  </form>
</div>
</body>

<script src="https://cdn.bootcss.com/jquery/3.3.1/jquery.min.js"></script>
<script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
<!--<script src="https://cdn.bootcss.com/flat-ui/2.3.0/js/flat-ui.min.js"></script>-->

</html>