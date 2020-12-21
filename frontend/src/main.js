import 'core-js/stable';
import 'regenerator-runtime/runtime';
import Vue from 'vue';
import App from './App.vue';
import 'vuesax/dist/vuesax.css'
import "./assets/css/index.css";
import Vuesax from 'vuesax'

Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from '@wailsapp/runtime';

Vue.use(Vuesax, {
    // options here
  })


Wails.Init(() => {
	new Vue({
		render: h => h(App)
	}).$mount('#app');
});
