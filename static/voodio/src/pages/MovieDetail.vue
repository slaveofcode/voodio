<template>
  <Layout>
    <Header title="Voodio Media Server" />
    <div class="flex flex-row flex-wrap justify-center border-solid border-b-2 border-brown-800 pb-12 mb-12">
      <div class="md:w-1/5 md:mr-12">
        <div v-if="detail.tmdbInfo" class="movie-cover w-full">
          <img class="bg-cover" :src="parseCover(detail.tmdbInfo.poster_path)" />
        </div>
      </div>
      <div class="text-center md:w-3/6 md:text-left">
        <h1 v-if="detail.tmdbInfo" class="text-orange-400">{{ detail.tmdbInfo.title }}</h1>
        <p v-if="detail.tmdbInfo" class="text-brown-400">{{ detail.tmdbInfo.tagline }}</p>
        <div v-if="detail.tmdbInfo" class="my-4 flex flex-row justify-center md:justify-start">
          <p class="font-bold mr-4 bg-brown-800 px-3 py-1">Rating: <span class="text-yellow-500">{{ detail.tmdbInfo.vote_average }}</span></p>
          <p class="font-bold bg-brown-800 px-3 py-1">Release Date: <span class="text-yellow-500">{{ detail.tmdbInfo.release_date }}</span></p>
        </div>
        <div class="my-4 mt-6">
          <p class="font-bold mb-4">Overview</p>
          <p v-if="detail.tmdbInfo" class="text-yellow-300">{{ detail.tmdbInfo.overview }}</p>
        </div>
        <div v-if="detail.tmdbInfo">
          <p class="font-bold mb-4">Genre</p>
          <div class="flex flex-row flex-wrap justify-center md:justify-start">
            <div class="genre-pills border border-solid border-red-700 font-bold ml-4 px-3 py-1 bg-red-700 text-xs" v-for="(genre, idx) in detail.tmdbInfo.genres" :key="idx">
              {{ genre.name }}
            </div>
          </div>
        </div>
        <div v-if="vplayerMounted" class="my-4 mt-12">
          <button class="bg-green-700 text-3xl px-6 py-1" @click="playMovie()">
            <fa-icon :icon="['fas', 'film']" />
            Play Movie
          </button>
        </div>
      </div>
    </div>
    <div class="mb-16">
      <video
        v-if="videoSource"
        ref="vplayer"
        class="video-js"
        controls
        preload="none"
        :poster="videoPoster">
        <source :src="videoSource" type="application/x-mpegURL" />
      </video>
    </div>
  </Layout>
</template>

<style lang="scss" scoped>
.movie-cover {
  width: 250px;
  height: 370px;
}

.video-js {
  width: 720px;
  height: 400px;
}

.genre-pills:first-child{
  margin-left: 0;
}
</style>

<script>
// import Hls from 'hls.js';
// import 'video.js/dist/video'
import videojs from 'video.js/dist/video'
import 'video.js/dist/video-js.css'
import Layout from "@/layouts/Main";
import Header from '@/components/Header';
import { getMovieDetail } from '@/utils/voodio_request';
import { getDetailMovieById } from '@/utils/tmdb_request';

export default {
  components: {
    Layout,
    Header,
  },
  data() {
    return {
      detail: {},
      videoPoster: undefined,
      videoSource: undefined,
      vplayerMounted: false,
      videoJsInst: undefined,
    }
  },
  created() {
    const { id } = this.$route.params
    const { tmdbId } = this.$route.query

    getMovieDetail(id).then((detail) => {
      // this.detail = detail
      this.detail = detail
      this.videoSource = `http://192.168.8.102:1818/hls/${this.detail.cleanDirName}/playlist.m3u8`

      this.$nextTick(() => {
        console.log("mounted -> this.$refs.vplayer", this.$refs.vplayer)
        if (this.$refs.vplayer && !this.vplayerMounted) {
          const t = this
          videojs(this.$refs.vplayer, {}, function() {
            t.vplayerMounted = true
            t.videoJsInst = this
          })
        }
      })
    })

    if (tmdbId) {
      getDetailMovieById(tmdbId).then((tmdbInfo) => {
        this.videoPoster = this.parseBackdrop(tmdbInfo.backdrop_path)
        this.$set(this.detail, 'tmdbInfo', tmdbInfo)
      })
    }
  },
  methods: {
    parseCover(coverFile) {
      return `https://image.tmdb.org/t/p/w300${coverFile}`
    },
    parseBackdrop(backdropFile) {
      return `https://image.tmdb.org/t/p/w500${backdropFile}`
    },
    playMovie() {
      if (this.vplayerMounted && this.videoJsInst) {
        this.videoJsInst.play()
      }
    }
  }
}
</script>