<template>
	<div>
        <Button type="button" label="Add row" @click="addRow()" />
        <br>
        <br>
        <DataTable :value="fields" class="p-datatable-responsive-demo p-datatable-sm p-datatable-gridlines" :paginator="true" :rows="10">
            <Column header="Name">
                <template #body="slotProps">
                    <InputText type="text" v-model="slotProps.data.name" class="p-inputtext-sm" placeholder="name" />
                </template>
            </Column>
            <Column header="Type">
                <template #body="slotProps">
                    <InputText type="text" v-model="slotProps.data.type" class="p-inputtext-sm" placeholder="type" />
                </template>
            </Column>
            <Column header="PK">
                <template #body="slotProps" style="text-align: center;">
                    <Checkbox v-model="slotProps.data.PK" @change="saveValue()" :binary="true" />
                </template>
            </Column>
            <Column header="NN">
                <template #body="slotProps">
                    <Checkbox v-model="slotProps.data.NN" @change="saveValue()" :binary="true" />
                </template>
            </Column>
            <Column header="AI">
                <template #body="slotProps">
                    <Checkbox v-model="slotProps.data.AI" @change="saveValue()" :binary="true" />
                </template>
            </Column>
            <Column header="UN">
                <template #body="slotProps">
                    <Checkbox v-model="slotProps.data.UN" @change="saveValue()" :binary="true" />
                </template>
            </Column>
            <Column header="Delete">
                <template #body="slotProps">
                    <Button icon="pi pi-trash" iconPos="right"  @click="deleteRow(slotProps.data)" />
                </template>
            </Column>
        </DataTable>
        <Button label="Create table"  @click="create()" />
	</div>
</template>

<script>
export default {
	name: 'CreateTable',
	data() {
      	return {
            field_name: null,
            selected_type: null,
		    database_types: [
                {name: 'Integer', value: 'INTEGER'},
                {name: 'Text', value: 'TEXT'},
                {name: 'Blob', value: 'BLOB'},
                {name: 'Real', value: 'REAL'},
                {name: 'Numeric', value: 'NUMERIC'}
            ],
            selected_options: null,
		    database_options: [
                {name: 'NN', value: 'NOT NULL'},
                {name: 'PK', value: 'PRIMARY KEY'},
                {name: 'AI', value: 'AUTOINCREMENT'},
                {name: 'UN', value: 'UNIQUE'},
            ],
            // Don't forget to clean up those variables
            has_primary_key: false,
            primary_key: null,
            //fields: [],
            //fields: [{name: "", type: "", NN: false, PK: false, AI: false, UN: false}],
            // For testing only
            fields: [
                {name: 'ArtistId', type: 'int', NN: true, PK: true, AI: true, UN: true}, 
                {name: 'ArtistName', type: 'string', NN: true, PK: false, AI: true, UN: true}
            ]
    	}
  	},
	methods: {	
        saveValue() {
            console.log(JSON.stringify(this.fields))
        },

        addRow() {
            this.fields.push({name: "", type: "", NN: false, PK: false, AI: false, UN: false});
        },
        
        deleteRow(prop) {
            this.fields.splice(this.fields.indexOf(prop), 1);
            //this.fields.splice(this.fields[prop], 1);
            console.log(JSON.stringify(this.fields))
        },

        create() {
            window.backend.createTable(JSON.stringify(this.fields)).then(result => {
				console.log(result)
			}) 
        }
	},
}
</script>
<style>
InputText { 
    max-width: 200px; 
}
label {
    margin-top: 5px;
    display: inline-block;
}
</style>
