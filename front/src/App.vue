<template>
  <component :is="currentComponent ? 'AuthLayout' : 'DefaultLayout'" />
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
  
    const currentComponent = () => {
      const isLoggedUser = Object.keys(user.value).length === 0 ? false : true
      
      return isLoggedUser
    }

    return {
      currentComponent
    }
  }
}

</script>
<style>

</style>
