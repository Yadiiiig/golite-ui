<template>
	<div>
		<Panel header="Tables">
			<Tree v-if="tabledata !== null" :value="tabledata" selectionMode="single" @node-select="onNodeSelect"></Tree>
				<div v-if="tabledata === null" class="p-mb-3 p-text-light">No database selected.</div>
		</Panel>
	</div>
</template>

<script>
export default {
	name: 'TableView',
	props: ["tabledata", "changeTab"],
	data() {
      	return {
          
    	}
  	},
	methods: {
		onNodeSelect(node) {
			window.backend.selectTable(node.label).then(result => {
				this.$emit('headerResult', result.Columns)
				this.$emit('tableResult', JSON.parse(result.Data))
			}) 
			this.$emit('tableName', node.label)
			this.changeTab(1);
        },
	}
}
</script>