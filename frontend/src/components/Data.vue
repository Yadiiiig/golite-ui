<template>
	<div>
		<Panel header="Data">
			<div v-if="name !== null">
				<h3>Table: {{ name }}</h3>
				<DataTable :value="rows" editMode="row" :resizableColumns="true" class="p-datatable-responsive" :editingRows.sync="editingRows"
        			@row-edit-init="onRowEditInit" @row-edit-cancel="onRowEditCancel" :paginator="true" :rows="10">
					<Column v-for="(item, index) in fields" :key=index :field="item" :header="item" sortable>
						<template #editor="slotProps">
							<InputText v-model="slotProps.data[slotProps.column.field]" autofocus />
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
	props: ["fields", "rows", "name"],
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

		onRowEditInit(event) {
			console.log(event)
            this.originalRows[event.index] = {...this.rows[event.index]};
        },
        onRowEditCancel(event) {
            this.products3, event.index, this.rows[event.index];
        },
	},
}
</script>



