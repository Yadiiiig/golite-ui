<template>
	<div>
		<Panel header="Create table">
            <label for="tablen">Table name </label>
            <br>
            <InputText id="tablen" type="text" />
            <hr>
            <div class="p-grid">
                <div class="p-col"> 
                    <Panel header="Add Fields">
                        <label for="fieldn">Field name </label>
                        <br>
                        <InputText v-model="field_name" type="text" />
                        <br>
                        <div class="p-formgroup-inline">
                            
                            <div class="p-field">
                                <p>Types: </p>
                                <SelectButton id="types" v-model="selected_type" :options="database_types" optionLabel="name" />
                            </div>

                            <div class="p-field">
                                <p>Constraints:</p>
                                <SelectButton id="options" v-model="selected_options" :options="database_options" optionLabel="name" :multiple="true" />
                            </div>
                        </div>

                        <Button type="button" label="Add Field" @click="addField" />
                    </Panel>
                </div>
                <div class="p-col">
                    <Panel header="Preview">

                        <DataTable :value="fields" class="p-datatable-responsive-demo p-datatable-sm p-datatable-gridlines" :paginator="true" :rows="10">
                            <Column header="Name">
                                <template #body="slotProps">
                                    {{slotProps.data.name}}
                                </template>
                            </Column>
                            <Column header="Type">
                                <template #body="slotProps">
                                    {{slotProps.data.type}}
                                </template>
                            </Column>

                            <Column header="PK">
                                <template #body="slotProps">
                                    <div v-if="slotProps.data.constraints.some(e => e.name === 'PK')">
                                        Yes
                                    </div>
                                </template>
                            </Column>
                            <Column header="NN">
                                <template #body="slotProps">
                                    <div v-if="slotProps.data.constraints.some(e => e.name === 'NN')">
                                        Yes
                                    </div>
                                </template>
                            </Column>
                            <Column header="AI">
                                <template #body="slotProps">
                                    <div v-if="slotProps.data.constraints.some(e => e.name === 'AI')">
                                        Yes
                                    </div>
                                </template>
                            </Column>
                            <Column header="UN">
                                <template #body="slotProps">
                                    <div v-if="slotProps.data.constraints.some(e => e.name === 'UN')">
                                        Yes
                                    </div>
                                </template>
                            </Column>
                        </DataTable>
                    </Panel>
                </div>
            </div>
            <Button type="button" label="test" @click="test" />
            <br>
		</Panel>	
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
            fields: []
    	}
  	},
	methods: {	
        addField() {
            var constrains_array = [];
            // for (var i = 0; i < this.selected_options.length; i++) {
            //     this.fields.push({this.selected_options[i].value});
            // }

            this.fields.push({
                "name": this.field_name,
                "type": this.selected_type.value,
                "constraints": this.selected_options
            })

            // if (constrains_array.includes("PRIMARY KEY")) {
                
            // }
            console.log(constrains_array.includes("PRIMARY KEY"))
            // this.has_primary_key = true;
            // let result = this.database_options.map(a => a.value);
            console.log(this.fields);
        },
        test() {
            console.log(this.has_primary_key);
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
