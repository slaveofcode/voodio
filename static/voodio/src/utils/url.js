
export const getCurrHost = () => {
  return window.location.hostname
}

export const getCurrPort = () => {
  const isDev = process.env.NODE_ENV === 'development'
  const debugPort = process.env.VUE_APP_SERVER_PORT
  return isDev ? debugPort : window.location.port
}

export const getCurrFullHost = () => {
  return `http://${getCurrHost()}:${getCurrPort()}`
}