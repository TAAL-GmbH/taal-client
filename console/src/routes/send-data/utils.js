import { copyTextToClipboard } from '../../lib/utils/clipboard'
import { failure } from '../../lib/utils/notifications'
import * as api from '../../lib/api'

const pageKey = 'page.send-data'

export const getModeItems = (t) => {
  return [
    { label: t(`${pageKey}.mode-item-label.raw`), value: 'raw' },
    { label: t(`${pageKey}.mode-item-label.hash`), value: 'hash' },
    { label: t(`${pageKey}.mode-item-label.encrypt`), value: 'encrypt' },
  ]
}

export function getStoreValue(key, defaultValue) {
  return localStorage.getItem(key) || defaultValue
}

export function setStoreValue(key, value) {
  localStorage.setItem(key, value)
}

export function updateStore(key, value, loading) {
  if (value) {
    localStorage.setItem(key, value)
  } else if (!loading) {
    localStorage.removeItem(key)
  }
}

export function copyCurl(curlCommand) {
  copyTextToClipboard(curlCommand)
}

export function getCorrectURL(url) {
  try {
    const urlParts = new URL(url)
    return `http://${url}`
  } catch {
    return `http://localhost${url}`
  }
}

export function getTimeoutMillis(timeout) {
  let t = timeout.trim()
  let factor = 1
  let numStr = '1'
  if (t.includes('ms')) {
    factor = 1
    numStr = t.split('ms').join('')
  } else if (t.includes('s')) {
    factor = 1000
    numStr = t.split('s').join('')
  } else if (t.includes('m')) {
    factor = 60000
    numStr = t.split('m').join('')
  }
  return parseInt(numStr) * factor
}

export function getApiKeys(t) {
  return api.getApiKeys(
    (data) => {
      return data
    },
    (error) =>
      failure(t(`${pageKey}.notifications.api-keys-load-failure`, { error }), {
        duration: 2000,
      })
  )
}

export function getSettings(t) {
  return api.getSettings(
    (data) => {
      return data
    },
    (error) =>
      failure(t(`${pageKey}.notifications.settings-load-failure`, { error }), {
        duration: 2000,
      })
  )
}

export function getCurlCommand(key, type, tag, mode, secret, data, url, file) {
  let curl = 'curl \\\n  -X POST \\\n'
  if (key) curl += `  -H 'Authorization: Bearer ${key}' \\\n`
  if (type) curl += `  -H 'Content-Type: ${type}' \\\n`
  if (tag) curl += `  -H 'X-Tag: ${tag}' \\\n`
  switch (mode) {
    case 'hash':
      curl += `  -H 'X-Mode: ${mode}' \\\n`
      break

    case 'encrypt':
      curl += `  -H 'X-Mode: ${mode}' \\\n`
      curl += `  -H 'X-Key: ${secret}' \\\n`
      break
  }

  if (file) {
    curl += `  --data-binary @'${file.name}' \\\n`
  } else if (data) {
    curl += `  -d '${data}' \\\n`
  }
  curl += `${url}/api/v1/transactions`

  return curl
}
