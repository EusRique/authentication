<template>
  <component :is='currentComponent' />
</template>

<script>
import { mapGetters } from 'vuex';
import AuthLayout from '@/views/Layout/AuthLayout.vue';
import DefaultLayout from '@/views/Layout/DefaultLayout.vue';
export default {
  name: 'App',

  components: { AuthLayout, DefaultLayout },

  data: () => ({}),

  computed: {
    ...mapGetters("User", {
      user: "getUser"
    }),

    currentComponent() {
      if (this.user) {
        return "DefaultLayout"
      }
      return 'AuthLayout'
    }
  }
}

</script>
<style>

</style>
