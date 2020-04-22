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
          <p class="font-bold text-xs md:text-base mr-4 bg-brown-800 px-3 py-1">Rating: <span class="text-yellow-500">{{ detail.tmdbInfo.vote_average }}</span></p>
          <p class="font-bold text-xs md:text-base bg-brown-800 px-3 py-1">Release Date: <span class="text-yellow-500">{{ detail.tmdbInfo.release_date }}</span></p>
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
        <div class="my-4 mt-12">
          <button v-if="vplayerMounted && isMoviePrepared" class="bg-green-700 text-3xl px-6 py-1" @click="playMovie()">
            <fa-icon :icon="['fas', 'film']" />
            Play Movie
          </button>
          <button v-if="!isMoviePrepared" class="bg-orange-500 text-lg px-3 py-2" @click="prepareTheMovie()">
            <fa-icon :icon="['fas', 'cogs']" />
            {{ !isOnPrepareMovie ? 'Prepare the Movie' : 'Movie Preparing...' }}
          </button>
          <p v-if="!isMoviePrepared" class="mt-4 text-orange-500">This movie isn't prepared, click to start video transcoding</p>
          <p class="text-yellow-500">* This operation might need more Disk Space and CPU resources</p>
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
  width: 100% !important;
  height: auto !important;
}

.video-js .vjs-tech {
  position: static;
  max-width: 100%;
  height: auto;
}

.genre-pills:first-child{
  margin-left: 0;
}
</style>

<script>
import videojs from 'video.js'
import 'video.js/dist/video-js.css'
import Layout from "@/layouts/Main";
import Header from '@/components/Header';
import { getMovieDetail, prepareMovie } from '@/utils/voodio_request';
import { getDetailMovieById } from '@/utils/tmdb_request';
import { getCurrHost } from '@/utils/url';

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
      isOnPrepareMovie: false
    }
  },
  computed: {
    isMoviePrepared() {
      return this.detail.isPrepared || this.detail.isInPrepare
    }
  },
  created() {
    const { id } = this.$route.params
    const { tmdbId } = this.$route.query

    getMovieDetail(id).then((detail) => {
      // this.detail = detail
      this.detail = detail
      this.videoSource = `http://${getCurrHost()}:1818/hls/${this.detail.cleanDirName}/playlist.m3u8`

      if (this.isMoviePrepared) {
        this.$nextTick(() => {
          if (this.$refs.vplayer && !this.vplayerMounted) {
            const t = this
            videojs(this.$refs.vplayer, {
              controls: true,
              autoplay: false,
              loop: false,
              preload: 'metadata',
              fluid: true,
              liveui: true,
            }, function() {
              t.vplayerMounted = true
              t.videoJsInst = this
              // this.addRemoteTextTrack({src: `http://${getCurrHost()}:1818/hls/${t.detail.cleanDirName}/subs.vtt`}, false)
            })
          }
        })
      }
    })

    if (tmdbId) {
      getDetailMovieById(tmdbId).then((tmdbInfo) => {
        this.videoPoster = this.parseBackdrop(tmdbInfo.backdrop_path)
        this.$set(this.detail, 'tmdbInfo', tmdbInfo)

        setTimeout(() => {
          if (this.videoJsInst) {
            this.videoJsInst.poster(this.videoPoster)
          }
        }, 2000)
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
    },
    prepareTheMovie() {
      prepareMovie(this.detail.ID).then(() => {
        this.isOnPrepareMovie = true
        setTimeout(() => {
          this.$router.go() // reload
        }, 4500)
      })
    }
  }
}
</script>