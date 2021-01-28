<template>
	<div>
		
			<div v-if="tabledata !== null" class="p-grid">
                <div class="p-col-fixed" style="max-width: 300px"> 
                    <Listbox v-if="tabledata !== null" v-model="selectedTable" :options="tabledata" @change="getInformation" optionLabel="name" :filter="true"/>
                </div>
                <div class="p-col">
                    <Panel header="Information">
						<div v-if="informationData !== null">
							<Button label="View data" @click="viewData" v-if="informationData !== null"/>
							<DataTable :value="informationData" style="height: 50%" v-if="informationData !== null">
								<Column field="Cid" header="Cid"></Column>
								<Column field="Name" header="Name"></Column>
								<Column field="Type" header="Type"></Column>
								<Column field="Pk" header="Primary Key"></Column>
							</DataTable>
						</div>
						<div v-if="informationData === null" class="p-mb-3 p-text-light">No table selected.</div>
                    </Panel>
                </div>
            </div>
			<div v-if="tabledata === null" class="p-mb-3 p-text-light">No database selected.</div>
		
	</div>
</template>

<script>
export default {
	name: 'TableView',
	props: ["tabledata", "changeTab"],
	data() {
      	return {
          selectedTable: "",
		  informationData: null,
		  primaryKey: ""
    	}
  	},
	methods: {
		viewData() {
			window.backend.selectTable(this.selectedTable.code).then(result => {
				this.$emit('headerResult', result.Columns)
				this.$emit('tableResult', JSON.parse(result.Data))
				this.$emit('tablePrimary', this.primaryKey)
			}) 
			this.$emit('tableName', this.selectedTable.code)
			this.changeTab(1);
        },

		getInformation() {
			window.backend.getTableInfo(this.selectedTable.code).then(info => {
				this.informationData = info;
				this.primaryKey = info.find(x => x.Pk === 1);
			})
		}
	}
}
</script>
