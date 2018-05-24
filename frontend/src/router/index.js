import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import sendmail from '@/components/SendMail'
import showlogs from '@/components/ShowLogs'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      component: Home,
      children: [
        { path: '', component: sendmail, name: 'sendmail' }
      ]
    },
    {
      path: '/logs',
      component: Home,
      children: [
        { path: '', component: showlogs, name: 'showlogs' }
      ]
    }
  ]
})
