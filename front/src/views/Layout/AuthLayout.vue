<template>
  <div class="FlexContainer">
    <div class="FlexContainer__FlexItem">
      <div class="FlexContainer__FlexItem__Banner">
        <img :src="bannerLogin" alt="Banner Login">
      </div>
    </div>
    <div class="FlexContainer__FlexItem">
      <router-view />
    </div>
  </div>
</template>

<script>
  export default {
    name: 'AuthLayout',

    components: {},

    props: {},

    data: () => ({
      bannerLogin: require('@/assets/images/login_banner.png')
    })
  }
</script>

<style lang="scss" scoped>
.FlexContainer {
  display: flex;
  flex-direction: row wrap;

  max-width: 100%;
  height: 100vh;
  overflow: hidden;
  &__FlexItem {
    display: flex;
    
    background-color: var(--color-primary-login);
  
    &__Banner img {
      max-width: 100%;
      height: 100vh;
    }
  }
}
/* MAX WIDTH = TELA MENORES QUE ESSA RESOLUÇÃO */
@media screen and (max-width: 1023px) {
  .FlexContainer {
    overflow: scroll;
  }
  .FlexContainer__FlexItem__Banner img {
    display: none;
  }
}
</style>