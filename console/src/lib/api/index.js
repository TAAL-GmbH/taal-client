import { spinCount } from '../stores'

function incSpinCount() {
  spinCount.update((n) => n + 1)
}

function decSpinCount() {
  spinCount.update((n) => n - 1)
}

function checkInitialResponse(response) {
  return new Promise(async (resolve, reject) => {
    if (response.ok) {
      let data = null

      const mimeType = response.headers.get('Content-Type') || 'text/plain'
      if (mimeType.toLowerCase().startsWith('application/json')) {
        try {
          data = await response.json()
        } catch (e) {
          data = null
        }
      } else if (
        mimeType.toLowerCase().startsWith('application/octet-stream')
      ) {
        try {
          data = await response.blob()
        } catch (e) {
          data = null
        }
      } else {
        data = await response.text()
      }

      resolve({ data })
    } else {
      let errorBody = null
      try {
        errorBody = await response.json()
      } catch (e) {
        errorBody = null
      }

      reject({
        code: response.status,
        message:
          errorBody?.error || response.statusText || 'Unspecified error.',
      })
    }
  })
}

function callApi(url, options = {}, done, fail) {
  if (!options.method) {
    options.method = 'GET'
  }
  incSpinCount()

  return fetch(url, options)
    .then(async (res) => {
      const { data } = await checkInitialResponse(res)
      return data
    })
    .then((data) => {
      if (done) {
        done(data)
      }
      return data
    })
    .catch((err) => {
      console.log(err)
      if (fail) {
        fail(err.message)
      }
    })
    .finally(decSpinCount)
}

function get(url, options = {}, done, fail) {
  return callApi(url, { ...options, method: 'GET' }, done, fail)
}

function post(url, options = {}, done, fail) {
  return callApi(url, { ...options, method: 'POST' }, done, fail)
}

function put(url, options = {}, done, fail) {
  return callApi(url, { ...options, method: 'PUT' }, done, fail)
}

function del(url, options = {}, done, fail) {
  return callApi(url, { ...options, method: 'DELETE' }, done, fail)
}

// api methods

export function getApiKeys(done, fail) {
  return get(`${BASE_URL}/api/v1/apikeys`, {}, done, fail)
}

export function getApiKeysUsage(done, fail) {
  return get(`${BASE_URL}/api/v1/apikeys/usage`, {}, done, fail)
}

export function registerKey(apiKey, done, fail) {
  return post(`${BASE_URL}/api/v1/apikeys/${apiKey}`, {}, done, fail)
}

export function deleteKey(apiKey, done, fail) {
  return del(`${BASE_URL}/api/v1/apikeys/${apiKey}`, {}, done, fail)
}

export function getTransactions(hours, done, fail) {
  return get(
    `${BASE_URL}/api/v1/transactions/?hours_back=${hours}`,
    {},
    done,
    fail
  )
}

export function writeTransaction(options, done, fail) {
  return post(`${BASE_URL}/api/v1/transactions`, options, done, fail)
}

export function getSettings(done, fail) {
  return get(`${BASE_URL}/api/v1/settings`, {}, done, fail)
}

export function updateSetting(options, done, fail) {
  return put(`${BASE_URL}/api/v1/settings`, options, done, fail)
}
