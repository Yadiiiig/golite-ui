<template>
	<div>
		<h3>Expand Mode</h3>
        <DataTable :value="products" :resizableColumns="true" columnResizeMode="expand" class="p-datatable-gridlines">
            <Column field="code" header="Code"></Column>
            <Column field="name" header="Name"></Column>
            <Column field="category" header="Category"></Column>
            <Column field="quantity" header="Quantity"></Column>
        </DataTable>	
	</div>
</template>

<script>
export default {
	name: 'Data',
	props: ["fields", "rows"],
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
    mounted
}
</script>