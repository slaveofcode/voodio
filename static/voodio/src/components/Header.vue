<template>
  <div class="mb-8 border-b-2 border-solid border-red-600 py-2">
    <h1 class="text-5xl text-center">{{ title }}</h1>
  </div>
</template>

<script>
export default {
  props: {
    title: {
      type: String,
      default: ''
    }
  }
}
</script>