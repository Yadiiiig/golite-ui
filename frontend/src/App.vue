<template>
	<div id="app" v-bind:style="{ backgroundColor: '#1c1e26'}">
		<Nav @tableResult="tablesChanged($event)" />
		<br>
		<TabView>
			<TabPanel header="Structure" :active.sync="active[0]">
				<TableView :tabledata="tables"  @headerResult="headersChanged($event)" @tableResult="tablesResultChanged($event)" @tableName="getTableName($event)" :changeTab="activate" />
			</TabPanel>
			<TabPanel header="Data" :active.sync="active[1]">
				<Data :fields="headers" :rows="tablesValues" :name="tableName" />
			</TabPanel>
			<TabPanel header="SQL-Query" :active.sync="active[2]">
				<SQLEditor />
			</TabPanel>
		</TabView>
	</div>
</template>
<script>
import Nav from "./components/Nav.vue";
import TableView from "./components/TableView.vue";
import Data from "./components/Data.vue";
import SQLEditor from "./components/SQLEditor.vue";

export default {
	name: "app",
	components: {
		Nav,
		TableView,
		Data,
		SQLEditor
	},
	data() {
		return {
			selected: null,
			tables: null,
			headers: null,
			tablesValues: null,
			tableName: null,
			active: [true, false, false]
		}
	},
	methods: {
		tablesChanged(e) {
			this.tables = e;
		},
		headersChanged(e) {
			this.headers = e;
		},
		tablesResultChanged(e) {
			this.tablesValues = e;
		},
		getTableName(e) {
			this.tableName = e;
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
