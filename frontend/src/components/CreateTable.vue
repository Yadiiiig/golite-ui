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
                        <InputText id="fieldn" type="text" />
                        <br>
                        <div class="p-formgroup-inline">
                            
                            <div class="p-field">
                                <p>Types: </p>
                                <SelectButton id="types" v-model="selected_type" :options="database_types" optionLabel="name" />
                            </div>
                            <div class="p-field">
                                <p>Constraints:</p>
                                <SelectButton id="options" v-model="selected_option" :options="database_options" optionLabel="name" :multiple="true" />
                            </div>
                        </div>

                        <Button type="button" label="Add Field" @click="addField" />
                    </Panel>
                </div>
                <div class="p-col">
                    <Panel header="Preview">
                        
                    </Panel>
                </div>
                <Button type="button" label="test" @click="test" />
            </div>
        
            <br>
		</Panel>	
	</div>
</template>

<script>
export default {
	name: 'CreateTable',
	data() {
      	return {
            selected_type: null,
		    database_types: [
                {name: 'Integer', value: 'INTEGER'},
                {name: 'Text', value: 'TEXT'},
                {name: 'Blob', value: 'BLOB'},
                {name: 'Real', value: 'REAL'},
                {name: 'Numeric', value: 'NUMERIC'}
            ],
            selected_option: null,
		    database_options: [
                {name: 'Not Null', value: 'NOT NULL'},
                {name: 'Primary key', value: 'PRIMARY KEY'},
                {name: 'Auto Increment', value: 'AUTOINCREMENT'},
                {name: 'Unique', value: 'UNIQUE'},
                {name: 'Numeric', value: 'NUMERIC'}
            ],
            // Don't forget to clean up those variables
            has_primary_key: false,
            fields: []
    	}
  	},
	methods: {	
        addField() {

            this.has_primary_key = true;
            let result = this.database_options.map(a => a.value);
            console.log(result);
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