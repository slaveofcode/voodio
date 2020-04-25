<template>
  <div class="mb-8 border-t-2 border-solid border-red-600 pt-5 text-center">
    <p class="copyright">Copyright &copy; {{ year }} Voodio</p>
  </div>
</template>

<script>
export default {
  computed: {
    year() {
      return new Date().getFullYear()
    }
  }
}
</script>