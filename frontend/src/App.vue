<template>
	<div id="app" v-bind:style="{ backgroundColor: '#1c1e26'}">
		<Nav @tableResult="tablesChanged($event)" />
		<br>
		<TabView>
			<TabPanel header="Structure" :active.sync="active[0]">
				<TableView :tabledata="tables"  @headerResult="headersChanged($event)" @tableResult="tablesResultChanged($event)" @tableName="getTableName($event)" @tablePrimary="primaryKeyChanged($event)" :changeTab="activate" />
			</TabPanel>
			<TabPanel header="Data" :active.sync="active[1]">
				<Data :fields="headers" :rows="tablesValues" :name="tableName" :primaryKey="primaryKey" />
			</TabPanel>
			<TabPanel header="SQL-Query" :active.sync="active[2]">
				<SQLEditor />
			</TabPanel>
			<TabPanel header="Create Table" :active.sync="active[3]">
				<CreateTable />
			</TabPanel>
		</TabView>
	</div>
</template>
<script>
import Nav from "./components/Nav.vue";
import TableView from "./components/TableView.vue";
import Data from "./components/Data.vue";
import SQLEditor from "./components/SQLEditor.vue";
import CreateTable from "./components/CreateTable.vue";

export default {
	name: "app",
	components: {
		Nav,
		TableView,
		Data,
		SQLEditor,
		CreateTable
	},
	data() {
		return {
			selected: null,
			tables: null,
			headers: null,
			tablesValues: null,
			tableName: null,
			primaryKey: null,
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
		primaryKeyChanged(e) {
			this.primaryKey = e;
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
	html,
	body {
	height: 100%;
	width: 100%;
	}

	body {
	padding: 0px;
	margin: 0px;
	background-color: '#1c1e26';
	font-size: 1px;
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
