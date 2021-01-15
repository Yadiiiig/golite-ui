<template>
	<div>
		<Panel header="Data">
			<div v-if="name !== null">
				<h3>Table: {{ name }}</h3>
				<DataTable :value="rows" editMode="cell" :resizableColumns="true" class="p-datatable-responsive editable-cells-table" :paginator="true" :rows="10">
					<Column v-for="(item, index) in fields" :key=index :field="item" :header="item" sortable>
						<template #editor="slotProps">
            				<InputText v-model="slotProps.data[slotProps.column.props.field]" />
        				</template>
					</Column>
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
          
    	}
  	},
	methods: {
		onNodeSelect(node) {
            console.log(node.label);
			 window.backend.selectTable(node.label).then(result => {
				console.log(result)
			 }) 
        },
	},
}
</script>
