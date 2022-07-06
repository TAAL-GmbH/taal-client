import { spinCount } from '../stores'

function incSpinCount() {
  spinCount.update((n) => n + 1)
}

function decSpinCount() {
  spinCount.update((n) => n - 1)
}

function checkInitialResponse(res) {
  if (!res.ok) {
    return res.text().then((text) => {
      throw new Error(text)
    })
  }
  return res.json()
}

function getErrorMessage(err) {
  let msg = ''
  try {
    const errJson = JSON.parse(err.message)
    msg = errJson.error
  } catch (error) {
    msg = err
  }
  return msg
}

function callApi(url, data, options = {}, done, fail) {
  if (!options.method) {
    options.method = 'GET'
  }
  if (data) {
    options.body = JSON.stringify(data)
  }
  incSpinCount()

  return fetch(url, options)
    .then((res) => {
      return checkInitialResponse(res)
    })
    .then((data) => {
      if (done) {
        done(data)
      }
      return data
    })
    .catch((err) => {
      const msg = getErrorMessage(err)
      console.log(msg)
      if (fail) {
        fail(msg)
      }
    })
    .finally(decSpinCount)
}

function get(url, done, fail) {
  return callApi(url, null, { method: 'GET' }, done, fail)
}

function post(url, data, done, fail) {
  return callApi(url, data, { method: 'POST' }, done, fail)
}

function del(url, done, fail) {
  return callApi(url, null, { method: 'DELETE' }, done, fail)
}

// api methods

export function getApiKeys(done, fail) {
  return get(`${BASE_URL}/api/v1/apikeys`, done, fail)
}

export function getApiKeysUsage(done, fail) {
  return get(`${BASE_URL}/api/v1/apikeys/usage`, done, fail)
}

export function registerKey(apiKey, done, fail) {
  return post(`${BASE_URL}/api/v1/apikeys/${apiKey}`, null, done, fail)
}

export function deleteKey(apiKey, done, fail) {
  return del(`${BASE_URL}/api/v1/apikeys/${apiKey}`, done, fail)
}

export function getTransactions(hours, done, fail) {
  return get(`${BASE_URL}/api/v1/transactions/?hours_back=${hours}`, done, fail)
}
