<template>
	<div id="app" v-bind:style="{ backgroundColor: '#1c1e26'}">
		<Nav @tableResult="tablesChanged($event)" />
		<div class="content">		
			<div class="center">
				<vs-table stripe :data="users">
					<template slot="header">
						<h3>
						Users
						</h3>
					</template>
					<template slot="thead">
						<vs-th>Type</vs-th>
						<vs-th>Name</vs-th>
						<vs-th>Table name</vs-th>
						<vs-th>Rootpage</vs-th>
					</template>

					<template slot-scope="{tables}">
						<vs-tr :key="indextr" v-for="(tr, indextr) in tables" >
						<vs-td :data="tables[indextr].TypeStr">
							{{tables[indextr].TypeStr}}
						</vs-td>

						<vs-td :data="tables[indextr].Name">
							{{tables[indextr].Name}}
						</vs-td>

						<vs-td :data="tables[indextr].TblName">
							{{tables[indextr].TblName}}
						</vs-td>

						<vs-td :data="tables[indextr].Rootpage">
							{{tables[indextr].Rootpage}}
						</vs-td>
						</vs-tr>
					</template>
					</vs-table>
			<vs-table v-model="selected">
				<template #thead>
				<vs-tr>
					<vs-th>Type</vs-th>
					<vs-th>Name</vs-th>
					<vs-th>Table name</vs-th>
					<vs-th>Rootpage</vs-th>
				</vs-tr>
				</template>
				<template #tbody>
				<vs-tr
					:key="i"
					v-for="(tr, i) in tables"
					:data="tr"
					:is-selected="selected == tr"
				>
					<vs-td>{{ tr.TypeStr }}</vs-td>
					<vs-td>{{ tr.Name }}</vs-td>
					<vs-td>{{ tr.TblName }}</vs-td>
					<vs-td>{{ tr.Rootpage }}</vs-td>
				</vs-tr>
				</template>
			</vs-table>

			<span class="data">
				<pre>
				{{ selected ? selected.SQL : 'Select an item in the table' }}
				</pre>
			</span>
			</div>
		</div>
	</div>
</template>
<script>
import Nav from "./components/Nav.vue";
import 'vuesax/dist/vuesax.css'


export default {
	name: "app",
	components: {
		Nav
	},
	data() {
		return {
			selected: null,
			tables: null
		}
	},
	methods: {
		tablesChanged(e) {
			this.tables = e;
			console.log(this.tables)
		}
	}
};
</script>

<style>
	html, body {
		height: 100%;
		margin: 0;
		padding: 0;
	}
	#app {
		display: flex;
		flex-direction: column;
		height: 100%;
	}
	.content {
		padding-top: 35px;
		padding-left: 25px;
	}
	p {
		color: white;
	}
</style>
