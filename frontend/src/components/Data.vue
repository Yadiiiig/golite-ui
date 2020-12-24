<template>
	<div>
		<h3>Table: {{ name }}</h3>
        <DataTable :value="rows" :resizableColumns="true" class="p-datatable-responsive" :paginator="true" :rows="10">
            <Column v-for="(item, index) in fields" :key=index :field="item.Name" :header="item.Name" sortable></Column>
        </DataTable>	
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
