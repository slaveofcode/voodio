<template>
  <div class="mb-8 border-b-2 border-solid border-red-600 py-2 mb-12">
    <div class="flex flex-row justify-between items-end">
      <router-link :to="{ name: 'home' }">
        <img src="../assets/logo.svg" class="logo" />
      </router-link>
      <a href="https://www.themoviedb.org/" class="tmdb">
        <span class="tmdb-text font-bold">API via</span> <img src="../assets/tmdb.svg" class="tmdb" />
      </a>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.logo {
  width: 200px;
}
.tmdb {
  .tmdb-text {
    color: #01b4e4;
    width: 100px;
  }

  img {
    width: 100px;
    height: 50px;
  }
}
</style>

<script>
export default {
  props: {
    title: {
      type: String,
      default: ""
    }
  }
};
</script>