import './styles/index.css'
import 'vue-toastification/dist/index.css'

import { createPinia } from 'pinia'
import { createApp } from 'vue'
import { vfmPlugin } from 'vue-final-modal'
import type { PluginOptions } from 'vue-toastification'
import Toast, { POSITION } from 'vue-toastification'

import App from '@/App.vue'
import router from '@/router'

const app = createApp(App)

app.use(createPinia())
app.use(router)

const options: PluginOptions = {
  position: POSITION.TOP_RIGHT,
  timeout: 8000,
}

app.use(Toast, options)

app.use(
  vfmPlugin({
    key: '$vfm',
    componentName: 'VueFinalModal',
    dynamicContainerName: 'ModalsContainer',
  })
)

app.mount('#app')
