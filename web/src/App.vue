<template>
    <div id="app">
        <AppHeader/>
        <vs-row vs-w="12">
            <vs-col class="show-sidebar" vs-type="flex" vs-justify="center" vs-align="center" vs-w="2">
                <AppSidebar/>
            </vs-col>
            <vs-col vs-type="flex" vs-justify="center" vs-align="center" vs-w="10">
                <router-view></router-view>
            </vs-col>
        </vs-row>
    </div>
</template>

<script>
    import AppHeader from './components/AppHeader'
    import AppSidebar from './components/AppSidebar'

    export default {
        name: 'App',
        components: {
            AppHeader,
            AppSidebar,
        }
    }
</script>

<style>
    #app {
        font-family: Avenir, Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: center;
        color: #2c3e50;
    }

    .show-sidebar {
        display: block !important;
    }
</style>
