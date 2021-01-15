<template>
	<div>
        <Panel header="SQL Editor">
            <div class="p-d-inline">
                <Button class="p-mx-auto" icon="pi pi-play" @click="run" />
            </div>
            <br>
            <Textarea v-model="query" rows="5" />
            <br>
            <Button type="button" label="Builder" @click="toggle" />
            <hr>
            <Data v-if="queryDone !== false" :fields="headers" :rows="tablesValues" />
            <OverlayPanel ref="op">
                <div class="p-p-1">
                    <Button class="p-button-text" label="CREATE" @click="build('CREATE')" />
                    <Button class="p-button-text" label="INSERT INTO" @click="build('INSERT INTO')" />
                    <Button class="p-button-text" label="SELECT" @click="build('SELECT')" />
                    <Button class="p-button-text" label="FROM" @click="build('FROM')" />
                    <Button class="p-button-text" label="*" @click="build('*')" />
                </div>
                <div class="p-p-1">
                    <Button class="p-button-text" label="ALTER" @click="build('ALTER')" />
                    <Button class="p-button-text" label="ADD" @click="build('ADD')" />
                    <Button class="p-button-text" label="DISTINCT" @click="build('DISTINCT')" />
                    <Button class="p-button-text" label="UPDATE" @click="build('UPDATE')" />
                    <Button class="p-button-text" label="SET" @click="build('SET')" />
                </div>
                <div class="p-p-1">
                    <Button class="p-button-text" label="DELETE" @click="build('DELETE')" />
                    <Button class="p-button-text" label="TRUNCATE" @click="build('TRUNCATE')" />
                    <Button class="p-button-text" label="AS" @click="build('AS')" />
                    <Button class="p-button-text" label="ORDER BY" @click="build('ORDER BY')" />
                    <Button class="p-button-text" label="ASC" @click="build('ASC')" />
                </div>
                <div class="p-p-1">
                    <Button class="p-button-text" label="DESC" @click="build('DESC')" />
                    <Button class="p-button-text" label="BETWEEN" @click="build('BETWEEN')" />
                    <Button class="p-button-text" label="WHERE" @click="build('WHERE')" />
                    <Button class="p-button-text" label="AND" @click="build('AND')" />
                    <Button class="p-button-text" label="OR" @click="build('OR')" />
                </div>
                <div class="p-p-1">
                    <Button class="p-button-text" label="NOT" @click="build('NOT')" />
                    <Button class="p-button-text" label="LIMIT" @click="build('LIMIT')" />
                    <Button class="p-button-text" label="IS NULL" @click="build('IS NULL')" />
                    <Button class="p-button-text" label="DROP" @click="build('DROP')" />
                    <Button class="p-button-text" label="DROP COLUMN" @click="build('DROP COLUMN')" />
                </div>
                <div class="p-p-1">
                    <Button class="p-button-text" label="DROP DATABASE" @click="build('DROP DATABASE')" />
                    <Button class="p-button-text" label="DROP TABLE" @click="build('DROP TABLE')" />
                    <Button class="p-button-text" label="DISTINCT" @click="build('DISTINCT')" />
                    <Button class="p-button-text" label="GROUP BY" @click="build('GROUP BY')" />
                    <Button class="p-button-text" label="HAVING" @click="build('HAVING')" />
                </div>
                <div class="p-p-1">
                    <Button class="p-button-text" label="IN" @click="build('IN')" />
                    <Button class="p-button-text" label="JOIN" @click="build('JOIN')" />
                    <Button class="p-button-text" label="UNION" @click="build('UNION')" />
                    <Button class="p-button-text" label="UNION ALL" @click="build('UNION ALL')" />
                    <Button class="p-button-text" label="EXISTS" @click="build('EXISTS')" />
                </div>
                <div class="p-p-1">
                    <Button class="p-button-text" label="LIKE" @click="build('LIKE')" />
                    <Button class="p-button-text" label="CASE" @click="build('CASE')" />
                </div>
            </OverlayPanel>
        </Panel>
	</div>
</template>

<script>
import Data from "./Data.vue";
export default {
	name: 'SQLEditor',
    components: {
		Data
	},
	data() {
      	return {
            query: "",
            headers: null,
            tablesValues: null,
            queryDone: false
    	}
  	},
	methods: {
        run() {
            window.backend.runQuery(this.query).then(result => {
				this.headers = result.Columns;
                this.tablesValues = JSON.parse(result.Data);
                this.queryDone = true;
			}) 
        },
        build(keyword) {
            keyword += " ";
            this.query += keyword;
        },
        toggle(event) {
            this.$refs.op.toggle(event);
        },
		onNodeSelect(node) {
            console.log(node.label);
			 window.backend.selectTable(node.label).then(result => {
				console.log(result)
			 }) 
        },
	},
}
</script>

<style>
textarea {
    width: 50%;
    max-width:95%;
    }
</style>


