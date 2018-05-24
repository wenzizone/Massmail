<template>
  <el-form :inline="true" :model="form" label-width="190px" :rules="rules" ref="form">
    <fieldset>
      <legend>邮件相关信息</legend>
      <el-form-item label="发送邮件显示名称:">
        <el-input v-model="form.aliasName" placeholder="xxx营销组"></el-input>
      </el-form-item>
      <el-form-item label="发送人地址:" prop="fromEmail">
        <el-input v-model="form.fromEmail" placeholder="no-reply@xxx.com"></el-input>
      </el-form-item>
    </fieldset>

    <fieldset>
      <legend>发送邮件服务器信息</legend>
      <el-form-item label="发送邮件服务器地址:" prop="smtpServer">
        <el-input v-model="form.smtpServer" placeholder="smtp.xxx.com"></el-input>
      </el-form-item>
      <el-form-item label="发送邮件服务器端口:">
        <el-input v-model="form.port" placeholder="465"></el-input>
      </el-form-item>
      <el-form-item label="发送邮件服务器登录密码:" prop="password">
        <el-input type="password" v-model="form.password"></el-input>
      </el-form-item>
    </fieldset>

    <fieldset>
      <legend>邮件发送配置</legend>
      <el-form-item label="邮件变量在文件中的列:" prop="varField">
        <el-input v-model="form.varField" placeholder="1,3,5,6,7"></el-input>
      </el-form-item>
      <el-form-item label="每封邮件发送间隔时间:">
        <el-col :span="11">
          <el-input placeholder="最小时间" v-model="form.minTime" style="width: 100%;"></el-input>
        </el-col>
        <el-col class="line" :span="2">-</el-col>
        <el-col :span="11">
          <el-input placeholder="最大时间" v-model="form.maxTime" style="width: 100%;"></el-input>
        </el-col>
      </el-form-item>
    </fieldset>

    <fieldset>
      <legend>发送邮件所需文件</legend>
      <el-form-item label="邮件变量文件:">
        <el-upload
          :limit="1"
          :on-exceed="handleExceed"
          name="varFile"
          accept=".txt,.csv"
          :disabled="isVarfileDisabled"
          action="/api/upload/varfile"
          :on-success="handleSuccess"
          :file-list="varfileList">
          <el-button size="small" type="primary">点击上传</el-button>
          <div slot="tip" class="el-upload__tip">只能上传csv/txt文件，且不超过500kb</div>
        </el-upload>
      </el-form-item>
      <el-form-item label="邮件主题模板文件:">
        <el-upload
          :limit="1"
          name="subjectFile"
          accept=".txt,.csv"
          :on-exceed="handleExceed"
          :on-success="handleSuccess"
          action="/api/upload/subjectile"
          :file-list="titlefileList">
          <el-button size="small" type="primary">点击上传</el-button>
          <div slot="tip" class="el-upload__tip">只能上传csv/txt文件，且不超过100kb</div>
        </el-upload>
      </el-form-item>
      <el-form-item label="邮件正文模板文件:">
        <el-upload
          :limit="1"
          name="msgFile"
          accept=".txt,.csv"
          :on-exceed="handleExceed"
          :on-success="handleSuccess"
          action="/api/upload/msgfile"
          :file-list="msgfileList">
          <el-button size="small" type="primary">点击上传</el-button>
          <div slot="tip" class="el-upload__tip">只能上传csv/txt文件，且不超过100kb</div>
        </el-upload>
      </el-form-item>
    </fieldset>

    <el-row>
      <el-button type="primary" @click="resetForm('form')">取消</el-button>
      <el-button type="primary" @click="submitForm('form')">发送</el-button>
    </el-row>
  </el-form>
</template>

<script>
export default {
  data () {
    return {
      form: {
        aliasName: '',
        fromEmail: '',
        smtpServer: '',
        port: 465,
        password: '',
        varField: '',
        minTime: null,
        maxTime: null,
        delayTime: [],
        emailFiles: {
          varFileName: null,
          subjectFile: null,
          messageFile: null
        }
      },
      isVarfileDisabled: false,
      varfileList: [],
      titlefileList: [],
      msgfileList: [],
      rules: {
        fromEmail: [
          { required: true, message: '请输入发送人邮件地址', trigger: 'blur' },
          { type: 'email', message: '请输入正确的邮件地址', trigger: 'change' }
        ],
        smtpServer: [
          { required: true, message: '请输入发送邮件服务器地址', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入邮件服务器密码', trigger: 'blur' }
        ],
        varField: [
          { required: true, message: '请输入邮件变量在文件中的列', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    handleExceed (files, varfileList) {
      this.$message.warning('只能上传一个文件')
    },
    handleSuccess (response, files, fileList) {
      // console.log('response = ', response);
      switch (response.fileType) {
        case 'varfile':
          this.form.emailFiles.varFileName = response.fileName
          this.varfileList = fileList
          break
        case 'subjectfile':
          this.form.emailFiles.subjectFile = response.fileName
          this.titlefileList = fileList
          break
        case 'msgfile':
          this.form.emailFiles.messageFile = response.fileName
          this.msgfileList = fileList
          break
      }
      console.log('files = ', files)
      console.log('fileList = ', fileList)
    },
    // 检查需要上传文件的部分是否都已经有了文件
    checkUploadFile () {
      if ((this.varfileList.length === 1) && (this.titlefileList.length === 1) && (this.msgfileList.length === 1)) {
        return true
      } else {
        return false
      }
    },
    // 处理邮件发送间隔时间
    checkDelayTime () {
      if (typeof this.form.minTime === 'object') {
        this.form.minTime = 1
      }
      if (typeof this.form.maxTime === 'object') {
        this.form.maxTime = 5
      }
      if (this.form.minTime === this.form.maxTime) {
        this.from.maxTime += 1
      }

      if (this.form.maxTime < this.form.minTime) {
        this.form.delayTime = [this.form.maxTime, this.form.minTime]
      } else {
        this.form.delayTime = [this.form.minTime, this.form.maxTime]
      }
    },
    // 提交群发邮件
    submitForm (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.checkDelayTime()
          if (this.checkUploadFile()) {
            axios.post('/api/sendmail', QS.stringify(this.form, { arrayFormat: 'brackets' }))
              .then((response) => {
                this.$message.success(`邮件发送提交成功，请耐心等待。。`)
                this.$refs[formName].resetFields()
              })
              .catch((error) => {
                this.$message.error(`邮件发送提交失败，请检查后重试。。`)
              })
          } else {
            this.$message.warning(`
            请上传 邮件变量文件，邮件主题模板文件，邮件正文模板文件。
            `)
            return false
          }
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    // 取消，初始化表单
    resetForm (formName) {
      this.$refs[formName].resetFields()
    }
  }
}
</script>

<style>
  .el-row, el-button {
    text-align: center;
    margin-top: 15px;
  }
  .line {
    text-align: center;
  }
  fieldset {
    min-width: 0;
    padding: 0;
    margin: 0;
    border: 0;
  }
  legend {
    display: block;
    width: 100%;
    padding: 0;
    margin-bottom: 20px;
    font-size: 21px;
    line-height: inherit;
    color: #333;
    border: 0;
    border-bottom: 1px solid #e5e5e5;
  }
</style>
