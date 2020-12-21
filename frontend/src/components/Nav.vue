<template>
    <div class="center examplex">
        <vs-navbar color="#1c1e26" text-white square center-collapsed>
            <template #left>
                <vs-navbar-item v-on:click="select">
                    Open file
                </vs-navbar-item>
                <vs-navbar-item>
                    New database
                </vs-navbar-item>
                <vs-navbar-item>
                    Save changes
                </vs-navbar-item>
                <vs-navbar-item>
                    Console
                </vs-navbar-item>
            </template>
        </vs-navbar>
    </div>
</template>

<script>
    export default {
        name: 'Nav',
        data() {
            return {
                tables: null
            };
  	    },
    	methods: {
            select() {
                window.backend.selectDatabase().then(result => {this.$emit('tableResult', result)
                console.log("hey")});
            },
            getTables() {
                this.$emit('tables', this.tables);
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