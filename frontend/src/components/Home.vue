<template>
  <el-container>
    <el-header height="auto">
      <el-row>
        <el-col :span="24"><h1>自助邮件群发系统</h1></el-col>
      </el-row>
      <topMenu :menuData="menuData" ref="topMenu"></topMenu>
    </el-header>

    <el-main>
      <router-view></router-view>
    </el-main>
  </el-container>
</template>

<script>
import topMenu from './Common/topMenu.vue'

export default {
  data () {
    return {
      menuData: []
    }
  },
  created () {
    let menus = [
      { title: '发送邮件', url: '/' },
      { title: '查看日志', url: '/logs' }
    ]
    this.menuData = menus
  },
  components: {
    topMenu
  }
}
</script>

<style>
  .el-header {
    text-align: center;
  }
</style>
