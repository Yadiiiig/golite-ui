<template>
    <div>
		<Menubar :model="items" />
    </div>
</template>


<script>
    export default {
        name: 'Nav',
        data() {
            return {
                nodes: [],
				items: [
                {
                   label:'Database',
                   icon:'pi pi-fw pi-folder-open',
                   items:[{
                        label:'Open',
                        icon:'pi pi-fw pi-file',
						command: () => {
							this.select();
						}
                    },
                    {
                         label:'Create',
                         icon:'pi pi-fw pi-file'
                    }]
                },
                {
                   label:'Edit',
                   icon:'pi pi-fw pi-pencil',
                   items:[
                   ]
                },
                {
                   label:'Quit',
                   icon:'pi pi-fw pi-power-off'
                }
             ]
            };
  	    },
    	methods: {
            select() {
                window.backend.selectDatabase().then(result => {
                    for (var i = 0; i < result.length; i++) {
                        this.nodes.push({
                            key: String(i),
                            label: result[i].Name,
                            data: 'Name',
                            icon: 'pi pi-fw pi-file',
                            style: "margin:-3px; margin-left: 1px; margin-top: 1px; padding:0; text-decoration: none;",
                            children: [
                                {
                                key: String(i) + '-0',
                                label: result[i].Rootpage,
                                data: 'Rootpage',
                                icon: 'pi pi-fw pi-cog',
                                style: "margin:-4px; padding:0;"
                                },
                                {
                                key: String(i) + '-1',
                                label: result[i].TblName,
                                data: 'TblName',
                                icon: 'pi pi-fw pi-cog',
                                style: "margin:-4px; padding:0;"
                                },
                                {
                                key: String(i) + '-2',
                                label: 'SQL',
                                data: 'SQL',
                                icon: 'pi pi-fw pi-cog',
                                style: "margin:-4px; padding:0;",
                                children: [
                                    {
                                    key: String(i) + '-1-0', 
                                    label: result[i].SQL,
                                    data: 'Query',
                                    style: "margin:-4px; padding:0;font-size:12px;"
                                    }
                                ]
                                }
                            ]
                        })
                    }
                this.$emit('tableResult', this.nodes)
                })
            },

            test() {
                console.log("hey")
            }
        }
    }
</script>

<style scoped>
input[type="file"] {
    display: none;
}
.custom-file-upload {
    cursor: pointer;
}
</style>