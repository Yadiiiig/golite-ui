<template>
    <div class="center examplex">
        <vs-navbar color="#1c1e26" text-white square center-collapsed v-model="active">
            <template #left>
                <vs-navbar-item :active="active == 'guide'" id="guide">
                    Open database
                </vs-navbar-item>
                <vs-navbar-item :active="active == 'docs'" id="docs">
                    New database
                </vs-navbar-item>
                <vs-navbar-item :active="active == 'components'" id="components">
                    Save changes
                </vs-navbar-item>
                <vs-navbar-item :active="active == 'license'" id="license">
                    Console
                </vs-navbar-item>
            </template>
        </vs-navbar>
    </div>
</template>

<script>
    export default {
        name: 'Nav',
        methods: {
            trigger() {
                this.$refs.fileInput.click()
            }
        }
    }
</script>