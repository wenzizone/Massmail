
new Vue({
  el: '#app'
})

Vue.component('my-checkbox', {
  template: '#checkbox-template',
  data() {
    return { checked: false, title: 'Check me' }
  },
  methods: {
    check() { this.checked = !this.checked; }
  }
});

<script type="text/x-template" id="checkbox-template">
  <div class="checkbox-wrapper" @click="check">
    <div :class="{ checkbox: true, checked: checked }"></div>
    <div class="title"></div>
  </div>
</script>