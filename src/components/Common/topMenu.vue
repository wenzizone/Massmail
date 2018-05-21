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
  },
  methods: {
    routerChange (item) {
      if (item.url !== this.$route.path) {
        router.push(item.url)
      }
    }
  }
}
</script>

<style scoped>

</style>
