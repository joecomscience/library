<template>
    <div class="parentx-static">

        <vs-sidebar static-position default-index="1" color="primary" class="sidebarx" spacer v-model="active" hidden-background>

            <div class="header-sidebar" slot="header">
                <vs-avatar size="70px" src="https://randomuser.me/api/portraits/men/85.jpg"/>

                <h4>
                    My Name
                    <vs-button color="primary" icon="more_horiz" type="flat"></vs-button>
                </h4>

            </div>

            <vs-sidebar-item index="1" icon="question_answer">
                Dashboard
            </vs-sidebar-item>

            <vs-sidebar-item index="2" icon="gavel">
                History
            </vs-sidebar-item>

            <vs-divider icon="person" position="left">
                User
            </vs-divider>

            <vs-sidebar-item index="3" icon="verified_user">
                Configurations
            </vs-sidebar-item>
            <vs-sidebar-item index="4" icon="account_box">
                Profile
            </vs-sidebar-item>
            <vs-sidebar-item index="5">
                Card
            </vs-sidebar-item>

            <div class="footer-sidebar" slot="footer">
                <vs-button icon="reply" color="danger" type="flat">log out</vs-button>
                <vs-button icon="settings" color="primary" type="border"></vs-button>
            </div>

        </vs-sidebar>
    </div>
</template>

<script>
    export default {
        name: "AppSidebar",
        data() {
            return {
                active: false,
            }
        }
    }
</script>

<style scoped>
    .parentx-static {
        overflow: hidden;
        height: 500px;
        position: relative;
    }
    .header-sidebar {
        display: flex;
        align-items: center;
        justify-content: center;
        flex-direction: column;
        width: 100%;
    }
    .header-sidebar h4 {
        display: flex;
        align-items: center;
        justify-content: center;
        width: 100%;
    }
    .header-sidebar h4 > button {
        margin-left: 10px;
    }
    .footer-sidebar {
        display: flex;
        align-items: center;
        justify-content: space-between;
        width: 100%;
    }
    .footer-sidebar > button {
        border: 0px solid rgba(0,0,0,0) !important;
        border-left: 1px solid rgba(0,0,0,0.07) !important;
        border-radius: 0px !important;
    }
</style>

