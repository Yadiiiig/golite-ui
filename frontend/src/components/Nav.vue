<template>
    <nav class="flex items-center justify-between flex-wrap bg-teal-500 p-6">
        <div class="w-full block flex-grow lg:flex lg:items-center lg:w-auto">
        <div class="text-sm lg:flex-grow">
            <div @click="trigger" class="trigger">Click me</div>
            <input type="file" ref="fileInput"/>
        </div>
        </div>
    </nav>
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