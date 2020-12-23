<template>
	<div id="app" v-bind:style="{ backgroundColor: '#1c1e26'}">
		<Nav @tableResult="tablesChanged($event)" />
		<br>
		<TabView>
			<TabPanel header="Structure" :active.sync="active[0]">
				<TableView :tabledata="tables"  @headerResult="headersChanged($event)" />
			</TabPanel>
			<TabPanel header="Data" :active.sync="active[1]">
				Content II
			</TabPanel>
			<TabPanel header="SQL-Query" :active.sync="active[2]">
				Content III
			</TabPanel>
		</TabView>
	</div>
</template>
<script>
import Nav from "./components/Nav.vue";
import TableView from "./components/TableView.vue";

export default {
	name: "app",
	components: {
		Nav,
		TableView
	},
	data() {
		return {
			selected: null,
			tables: null,
			headers: null,
			active: [true, false, false]
		}
	},
	methods: {
		tablesChanged(e) {
			this.tables = e;
			//console.log(JSON.stringify(this.tables))
		},
		headersChanged(e) {
			this.headers = e;
			console.log(JSON.stringify(this.headers))
		},
		activate(index) {
            let activeArray = [...this.active];
            for (let i = 0 ; i < activeArray.length; i++) {
                activeArray[i] = (i === index);
            }
            this.active = activeArray;
        }
	}
};
</script>

<style>
	html, body {
		height: 100%;
		margin: 0;
		padding: 0;
		background-color: '#1c1e26';
	}
	#app {
		display: flex;
		flex-direction: column;
		height: 100%;
	}
	.content {
		padding-top: 35px;
		padding-left: 25px;
	}

</style>
