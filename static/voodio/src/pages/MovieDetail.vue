<template>
  <Layout>
    <Header title="Voodio Media Server" />
    <div class="flex flex-row justify-evenly">
      <div class="w-1/5 mr-12">
        <div v-if="detail.tmdbInfo" class="movie-cover w-full">
          <img class="bg-cover" :src="parseCover(detail.tmdbInfo.poster_path)" />
        </div>
      </div>
      <div class="w-4/5">
        <h1 v-if="detail.tmdbInfo">{{ detail.tmdbInfo.title }}</h1>
        <p v-if="detail.tmdbInfo">{{ detail.tmdbInfo.tagline }}</p>
        <p v-if="detail.tmdbInfo">{{ detail.tmdbInfo.overview }}</p>
        <p v-if="detail.tmdbInfo">{{ detail.tmdbInfo.vote_average }}</p>
        <p v-if="detail.tmdbInfo">{{ detail.tmdbInfo.release_date }}</p>
        <div v-if="detail.tmdbInfo" class="flex flex-row flex-wrap">
          <div class="border border-solid border-red-600 text-base font-bold ml-4 px-2 py-1 rounded bg-red-500" v-for="(genre, idx) in detail.tmdbInfo.genres" :key="idx">
            {{ genre.name }}
          </div>
        </div>
        <button>Play Movie</button>
      </div>
    </div>
  </Layout>
</template>

<style lang="scss" scoped>
.movie-cover {
  width: 250px;
  height: 370px;
}
</style>

<script>
import Layout from "@/layouts/Main";
import Header from '@/components/Header';
import { getMovieDetail } from '@/utils/voodio_request';
import { getDetailMovieById } from '@/utils/tmdb_request';

export default {
  components: {
    Layout,
    Header
  },
  data() {
    return {
      detail: {}
    }
  },
  created() {
    const { id } = this.$route.params
    const { tmdbId } = this.$route.query

    getMovieDetail(id).then((detail) => {
      this.detail = detail
      this.$forceUpdate()
    })

    if (tmdbId) {
      getDetailMovieById(tmdbId).then((tmdbInfo) => {
        this.detail.tmdbInfo = tmdbInfo
        this.$forceUpdate()
      })
    }
  },
  methods: {
    parseCover(coverFile) {
      return `https://image.tmdb.org/t/p/w300${coverFile}`
    }
  }
}
</script>