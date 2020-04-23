const extensionsUsingCSS = ['vue', 'html'];
module.exports = {
  plugins: {
    tailwindcss: {},
     // see:  https://www.purgecss.com/configuration#options
    'vue-cli-plugin-tailwind/purgecss': {
      content: [
        `./@(public|src)/**/*.@(${extensionsUsingCSS.join('|')})`, 
        './node_modules/video.js/**/*.css'
      ]
    }
  }
}
