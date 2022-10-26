<template>
  <section class="AuthLayout">
    <div class="AuthLayout__SideBar">
      AUTH LAYOUT
    </div>
    <div>
      <router-view />
    </div>
  </section>
</template>

<script>
  export default {
    name: 'AuthLayout',

    components: {},

    data: () => ({})
  }
</script>

<style lang="scss" scoped>
.AuthLayout {
  display: flex;
  flex-direction: row;
  &__SideBar {
    background-color: var(--color-primary-login);
  }
 }
</style>