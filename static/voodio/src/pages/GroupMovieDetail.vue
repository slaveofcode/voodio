<template>
  <Layout>
    <Header title="Voodio Media Server" />
    <div v-for="(detail, idx) in movies" :key="idx" class="flex flex-row flex-wrap justify-center border-solid border-b-2 border-brown-800 pb-12 mb-12">
      <div class="text-center w-full">
        <h1 class="text-orange-400 mb-5">{{ detail.cleanBaseName }}</h1>
        <div class="mb-16">
          <video
            v-if="detail.videoSource"
            :ref="`vplayer_${detail.ID}`"
            class="video-js"
            controls
            preload="none"
            poster="../assets/movie.svg">
            <source :src="detail.videoSource" type="application/x-mpegURL" />
          </video>
        </div>
        <div class="my-4 mt-12">
          <button v-if="vplayerMounted[`${detail.ID}`] && isMoviePrepared(detail)" class="bg-green-700 text-3xl px-6 py-1" @click="playMovie(detail)">
            <fa-icon :icon="['fas', 'film']" />
            Play Movie
          </button>
          <button v-if="!isMoviePrepared(detail)" class="bg-orange-500 text-lg px-3 py-2" @click="prepareTheMovie(detail.ID)">
            <fa-icon :icon="['fas', 'cogs']" />
            {{ !isOnPrepareMovie[detail.ID] ? 'Prepare the Movie' : 'Movie Preparing...' }}
          </button>
          <p v-if="!isMoviePrepared(detail)" class="mt-4 text-orange-500">This movie isn't prepared, click to start video transcoding</p>
          <p v-if="!isMoviePrepared(detail)" class="text-green-500 text-sm">* You can start watching streamly while the video is processed in the background</p>
          <p v-if="!isMoviePrepared(detail)" class="text-yellow-500 text-sm">* This operation might drain you Disk Space and CPU resources for a temporary</p>
        </div>
      </div>
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
</style>

<script>
import { mapGetters } from 'vuex';
import videojs from 'video.js';
import 'video.js/dist/video-js.css';
import Layout from "@/layouts/Main";
import Header from '@/components/Header';
import { getMovieGroups, prepareMovie } from '@/utils/voodio_request';
// import tmdbApi from '@/utils/tmdb_request';
import { getCurrFullHost } from '@/utils/url';

export default {
  components: {
    Layout,
    Header,
  },
  data() {
    return {
      movies: [],
      vplayerMounted: {},
      videoJsInst: {},
      isOnPrepareMovie: {},
    }
  },
  computed: {
    ...mapGetters({
      tmdbApiKey: 'tmdb/tmdb_api_key'
    }),
  },
  created() {
    const { id } = this.$route.params

    getMovieGroups(id).then(({ movies }) => {
      this.movies = movies
      for (const m of this.movies) {
        m.videoSource = `${getCurrFullHost()}/hls/${m.ID}/playlist.m3u8`
        if (this.isMoviePrepared(m)) {
          this.$nextTick(() => {
            if (this.$refs[`vplayer_${m.ID}`] && !this.vplayerMounted[m.ID]) {
              const t = this
              videojs(this.$refs[`vplayer_${m.ID}`][0], {
                controls: true,
                autoplay: false,
                loop: false,
                preload: 'metadata',
                liveui: true,
              }, function() {
                t.$set(t.vplayerMounted, m.ID, true)
                t.$set(t.videoJsInst, m.ID, this)
                // this.addRemoteTextTrack({src: `${getCurrFullHost()}/hls/${t.detail.cleanDirName}/subs.vtt`}, false)
              })
            }
          })
        }
      }
    })
  },
  methods: {
    isMoviePrepared(movie) {
      return movie.isPrepared || movie.isInPrepare
    },
    async prepareTheMovie(movieId) {
      await prepareMovie(movieId)
      this.isOnPrepareMovie[movieId] = true
      setTimeout(() => this.$router.go(), 4500)
    },
    playMovie(movie) {
      if (this.vplayerMounted[movie.ID] && this.videoJsInst[movie.ID]) {
        this.videoJsInst[movie.ID].play()
      }
    }
  }
}
</script>