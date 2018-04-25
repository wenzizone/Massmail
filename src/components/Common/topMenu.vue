<template>
  <el-menu :default-active="activeIndex" mode="horizontal">
    <el-menu-item v-for="menu in menuData" :index='menu.url' @click="routerChange(menu)">{{menu.title}}</el-menu-item>
  </el-menu>
</template>

<script>
export default {
  props: ['menuData'],
  data () {
    return {
      activeIndex: this.$route.path
    }
  }
}
</script>

<style scoped>

</style>
