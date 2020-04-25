
export const getCurrHost = () => {
  return window.location.hostname
}

export const getCurrPort = () => {
  const isDev = process.env.NODE_ENV === 'development'
  const debugPort = process.env.VUE_APP_SERVER_PORT
  return isDev ? window.location.port : debugPort
}

export const getCurrFullHost = () => {
  return `http://${getCurrHost()}:${getCurrPort()}`
}