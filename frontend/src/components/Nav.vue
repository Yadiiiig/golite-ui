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
                        label:'Open database',
                        icon:'pi pi-fw pi-file',
						command: () => {
							this.select();
						}
                    },
                    {
                        label:'Close database',
                        icon:'pi pi-fw pi-file',
                        command: () => {
							this.close();
						}
                    },
                    {
                        label:'Create database',
                        icon:'pi pi-fw pi-file'
                    }]
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
                    this.$emit('tableResult', null);
                    this.nodes = [];
                    for (var i = 0; i < result.length; i++) {
                        // var info_nodes = [];
                        // this.nodes.push({
                        //     key: String(i),
                        //     label: result[i].Name,
                        //     data: 'Name',
                        //     icon: 'pi pi-fw pi-file',
                        //     style: "margin:-3px; margin-left: 1px; margin-top: 1px; padding:0; text-decoration: none;",
                        //     children: info_nodes
                        // })
                        this.nodes.push({
                            "name": result[i].Name,
                            "code": result[i].Name
                        })
                    }
                this.$emit('tableResult', this.nodes)
                })
            },
            close() {
                window.backend.closeDatabase().then(() => {
                    this.$emit('tableResult', null);
                    this.$emit('headerResult', null)
				    this.$emit('tableResult', null)
			        this.$emit('tableName', null)
                    this.nodes = [];
                })
            },
            // selectInfo(result, i) {
            //     var info_nodes = [];
            //             console.log("first " + i)
            //             // window.backend.getTableInfo(result[i].Name).then(info => {
            //             //     console.log("second " + i)
            //             //     for (var x = 0; x < info.length; x++) {
            //             //         console.log("third " + i)
            //             //         var label_text = "Name: " + info[x].Name + " Type: " + info[x].Type + " Not Null: " + info[x].NotNull 
            //             //         info_nodes.push({
            //             //             key: String(i) + '-0',
            //             //             label: label_text,
            //             //             data: 'Rootpage',
            //             //             icon: 'pi pi-fw pi-cog',
            //             //             style: "margin:-4px; padding:0;"
            //             //         })
            //             //     }
            //             // })

            //             info_nodes.push({
            //                 key: String(i) + '-2',
            //                 label: 'SQL',
            //                 data: 'SQL',
            //                 icon: 'pi pi-fw pi-cog',
            //                 style: "margin:-4px; padding:0;",
            //                 children: [
            //                     {
            //                     key: String(i) + '-1-0', 
            //                     label: result[i].SQL,
            //                     data: 'Query',
            //                     style: "margin:-4px; padding:0;font-size:12px;"
            //                     }
            //                 ]
            //             })
            //             this.nodes.push({
            //                 key: String(i),
            //                 label: result[i].Name,
            //                 data: 'Name',
            //                 icon: 'pi pi-fw pi-file',
            //                 style: "margin:-3px; margin-left: 1px; margin-top: 1px; padding:0; text-decoration: none;",
            //                 children: info_nodes
            //             })
            // },

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