<template>
	<div id="app" v-bind:style="{ backgroundColor: '#1c1e26'}">
		<Nav />
		<Test />
	</div>
</template>
<script>
import Test from "./components/Test.vue";
import Nav from "./components/Nav.vue";
import 'vuesax/dist/vuesax.css'
import "./assets/css/index.css";

export default {
	name: "app",
	components: {
		Test,
		Nav
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
</style>
