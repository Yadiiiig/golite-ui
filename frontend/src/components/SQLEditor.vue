<template>
	<div>
		<h3>SQL Editor</h3>
        <div class="p-d-inline">
            <Button class="p-mx-auto" icon="pi pi-play" @click="run" />
        </div>
        <br>
        <Textarea v-model="query" :autoResize="true" rows="5" cols="100" />
        <hr>
        <Data v-if="queryDone !== false" :fields="headers" :rows="tablesValues" />
	</div>
</template>

<script>
import Data from "./Data.vue";
export default {
	name: 'SQLEditor',
    components: {
		Data
	},
	data() {
      	return {
            query:"",
            headers: null,
            tablesValues: null,
            queryDone: false
    	}
  	},
	methods: {
        run() {
            window.backend.runQuery(this.query).then(result => {
				this.headers = result.Columns;
                this.tablesValues = JSON.parse(result.Data);
                this.queryDone = true;
			}) 
        },

		onNodeSelect(node) {
            console.log(node.label);
			 window.backend.selectTable(node.label).then(result => {
				console.log(result)
			 }) 
        },
	},
}
</script>
