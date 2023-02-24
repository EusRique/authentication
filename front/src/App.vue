<template>
  <component :is="currentComponent"/>
</template>

<script>
import { useStore } from 'vuex';
import { computed } from 'vue';
import AuthLayout from '@/views/Layout/AuthLayout.vue';
import DefaultLayout from '@/views/Layout/DefaultLayout.vue';
export default {
  name: 'App',

  components: { AuthLayout, DefaultLayout },

  setup (props) {
    const store = useStore()
    const user = computed(() => store.getters.getUser)

    const currentComponent = computed(() => {
      const isLoggedUser = Object.keys(user.value)
      return isLoggedUser.length ? 'DefaultLayout' : 'AuthLayout'
    })

    return {
      currentComponent,
    }
  }
}

</script>
<style>

</style>
