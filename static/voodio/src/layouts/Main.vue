<template>
  <div>
    <div class="container mx-auto px-4 mt-8 box-border">
      <slot />
    </div>
    <div class="container mx-auto px-4 box-border">
      <Footer />
    </div>
  </div>
</template>

<script>
import Footer from '@/components/Footer';
export default {
  components: {
    Footer,
  }
};
</script>