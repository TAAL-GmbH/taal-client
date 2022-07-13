import { copyTextToClipboard } from '../../lib/utils/clipboard'
import * as api from '../../lib/api'

export const modeItems = [
  { label: 'Raw', value: 'raw' },
  { label: 'Hash', value: 'hash' },
  { label: 'Encrypt', value: 'encrypt' },
]

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

export async function copyCurl(curlCommand, addNotification) {
  const { ok, error } = await copyTextToClipboard(curlCommand)

  addNotification({
    text: ok ? `Successfully copied.` : `Failed to copy.`,
    position: 'bottom-left',
    type: ok ? 'success' : 'danger',
    removeAfter: 2000,
  })
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

export function getApiKeys(addNotification) {
  return api.getApiKeys(
    (data) => {
      return data
    },
    (error) => {
      addNotification({
        text: `Failed to load api keys: ${error}`,
        position: 'bottom-left',
        type: 'danger',
        removeAfter: 2000,
      })
    }
  )
}

export function getSettings(addNotification) {
  return api.getSettings(
    (data) => {
      return data
    },
    (error) => {
      addNotification({
        text: `Failed to load settings: ${error}`,
        position: 'bottom-left',
        type: 'danger',
        removeAfter: 2000,
      })
    }
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

// function uploadFile(file) {
//   const url = `${BASE_URL}/api/v1/transactions`
//   const xhr = new XMLHttpRequest()
//   const formData = new FormData()
//   const fileKey = getFileKey(file)
//   xhr.open('POST', url, true)

//   xhr.upload.addEventListener('progress', function (e) {
//     console.log('progress = ', e.loaded / e.total)
//     // TODO: right now the backend does not give proper progress updates
//     //       basically (e.loaded / e.total) evaluates to 1 right at the start already
//     //       we don't want to indicate completed progress when upload is still in progress, so
//     //       not doing an update here for now
//     // fileProgressData[fileKey] = {
//     //   state: 'progress',
//     //   progress: e.loaded / e.total,
//     // }
//   })

//   xhr.addEventListener('readystatechange', function (e) {
//     console.log(
//       'readystatechange: xhr.readyState = ',
//       xhr.readyState,
//       ' xhr.status = ',
//       xhr.status
//     )

//     // we reached a final state
//     if (xhr.readyState === 4) {
//       if (xhr.status >= 200 && xhr.status < 300) {
//         fileProgressData[fileKey] = { state: 'success', progress: 1 }

//         addNotification({
//           text: `Successfully uploaded file: ${file.name}`,
//           position: 'bottom-left',
//           type: 'success',
//           removeAfter: 2000,
//         })
//       } else {
//         fileProgressData[fileKey] = { state: 'failure', progress: 1 }

//         addNotification({
//           text: `Upload failed for: ${file.name}`,
//           position: 'bottom-left',
//           type: 'danger',
//           removeAfter: 2000,
//         })
//       }
//     }
//   })

//   // initialise start state
//   fileProgressData[fileKey] = { state: 'progress', progress: 0 }

//   formData.append('file', file)

//   // headers
//   xhr.setRequestHeader('Authorization', 'Bearer ' + apiKey)
//   xhr.setRequestHeader('Content-Type', file.type)
//   xhr.setRequestHeader('X-Tag', tag)
//   xhr.setRequestHeader('X-Filename', file.name)
//   xhr.setRequestHeader('X-Mode', mode)
//   xhr.setRequestHeader('X-Key', secret)

//   // set timeout
//   xhr.timeout = getTimeoutMillis(timeout)

//   xhr.send(formData)
// }
