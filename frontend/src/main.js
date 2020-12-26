import 'core-js/stable';
import 'regenerator-runtime/runtime';
import Vue from 'vue';
import App from './App.vue';

import 'primevue/resources/themes/vela-green/theme.css'
import 'primevue/resources/primevue.min.css';
import 'primeicons/primeicons.css';

import Menubar from 'primevue/menubar';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Tooltip from 'primevue/tooltip';
import Panel from 'primevue/panel';
import Tree from 'primevue/tree';
import TabView from 'primevue/tabview';
import TabPanel from 'primevue/tabpanel';
import Textarea from 'primevue/textarea';
import Button from 'primevue/button';
import OverlayPanel from 'primevue/overlaypanel';

Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from '@wailsapp/runtime';

Vue.component('Menubar', Menubar);
Vue.component('DataTable', DataTable);
Vue.component('Column', Column);
Vue.directive('tooltip', Tooltip);
Vue.component('Panel', Panel);
Vue.component('Tree', Tree);
Vue.component('TabView', TabView);
Vue.component('TabPanel', TabPanel);
Vue.component('Textarea', Textarea);
Vue.component('Button', Button);
Vue.component('OverlayPanel', OverlayPanel);

Wails.Init(() => {
	new Vue({
		render: h => h(App)
	}).$mount('#app');
});
