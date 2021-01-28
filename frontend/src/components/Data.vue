<template>
	<div>
		<Panel header="Data">
			<div v-if="name !== null">
				<h3>Table: {{ name }}</h3>
				<DataTable :value="rows" editMode="cell" class="editable-cells-table" :resizableColumns="true" 
					@cell-edit-init="onCellEditInit" @cell-edit-complete="onCellEditComplete" @cell-edit-cancel="onCellEditCancel" :paginator="true" :rows="10">
					<Column v-for="(item, index) in fields" :key=index :field="item" :header="item" sortable>
						<template #editor="slotProps">
							<InputText v-model="slotProps.data[slotProps.column.field]" />
						</template>
					</Column>
					<Column :rowEditor="true" headerStyle="width:7rem" bodyStyle="text-align:center"></Column>
				</DataTable>
			</div>
			<div v-else class="p-mb-3 p-text-light">No database selected.</div>
		</Panel>	
	</div>
</template>

<script>
export default {
	name: 'Data',
	props: ["fields", "rows", "name", "primaryKey"],
	data() {
      	return {
          	editingRows: [],
			originalRows: {}
    	}
  	},
	methods: {
		onNodeSelect(node) {
            console.log(node.label);
			window.backend.selectTable(node.label).then(result => {
				console.log(result)
			 }) 
        },

		onCellEditInit(event) {
            this.originalRows[event.index] = {...this.rows[event.index]};
			console.log(event, "#############")
			console.log(this.originalRows)
        },
		
		onCellEditComplete(event) {
			console.log(this.name, this.primaryKey.Name, event.data[this.primaryKey.Name].toString(), event.field, event.data[event.field].toString())
			window.backend.editField(this.name, this.primaryKey.Name, event.data[this.primaryKey.Name].toString(), event.field, event.data[event.field].toString()).then(result => {
				console.log(result)
			}) 
		},

		onCellEditCancel(event) {
			this.$set(this.rows, event.index, this.originalRows[event.index]);
			console.log(event, "canceled")
		},
	},
}
</script>



